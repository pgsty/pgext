/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Ext is one canonical extension record from pgext.universe. New universe
// columns are exposed directly; the state/repo/ext_* fields remain as v1 API
// compatibility aliases for existing clients.
type Ext struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Pkg      string `json:"pkg"`
	LeadExt  string `json:"lead_ext"`
	Category string `json:"category"`

	Packaged  bool     `json:"packaged"`
	Lifecycle string   `json:"lifecycle,omitempty"`
	Kind      string   `json:"kind"`
	Kernel    string   `json:"kernel,omitempty"`
	Vendor    string   `json:"vendor,omitempty"`
	Contrib   bool     `json:"contrib"`
	Lang      string   `json:"lang"`
	License   string   `json:"license,omitempty"`
	Tags      []string `json:"tags,omitempty"`

	Version    string `json:"version,omitempty"`
	URL        string `json:"url,omitempty"`
	RepoURL    string `json:"repo_url,omitempty"`
	HomeURL    string `json:"home_url,omitempty"`
	DocURL     string `json:"doc_url,omitempty"`
	LicenseURL string `json:"license_url,omitempty"`
	ControlURL string `json:"control_url,omitempty"`
	AuthorURL  string `json:"author_url,omitempty"`
	CargoURL   string `json:"cargo_url,omitempty"`
	PGXNURL    string `json:"pgxn_url,omitempty"`

	HasBin      bool `json:"has_bin"`
	HasLib      bool `json:"has_lib"`
	NeedDDL     bool `json:"need_ddl"`
	NeedLoad    bool `json:"need_load"`
	Trusted     bool `json:"trusted"`
	Relocatable bool `json:"relocatable"`
	Lead        bool `json:"lead"`

	Libs        []string `json:"libs,omitempty"`
	PreloadLibs []string `json:"preload_libs,omitempty"`
	Schemas     []string `json:"schemas,omitempty"`
	PG          []int    `json:"pg_ver,omitempty"`
	Requires    []string `json:"requires,omitempty"`
	RequiredBy  []string `json:"required_by,omitempty"`
	RequireBy   []string `json:"require_by,omitempty"` // v1 compatibility
	SeeAlso     []string `json:"see_also,omitempty"`
	Siblings    []string `json:"siblings,omitempty"`

	Tarball  string   `json:"tarball,omitempty"`
	PGRXVer  string   `json:"pgrx_ver,omitempty"`
	RPMVer   string   `json:"rpm_ver,omitempty"`
	RPMRepo  string   `json:"rpm_repo,omitempty"`
	RPMPkg   string   `json:"rpm_pkg,omitempty"`
	RPMPG    []int    `json:"rpm_pg,omitempty"`
	RPMDeps  []string `json:"rpm_deps,omitempty"`
	RPMBuild bool     `json:"rpm_build"`
	DEBVer   string   `json:"deb_ver,omitempty"`
	DEBRepo  string   `json:"deb_repo,omitempty"`
	DEBPkg   string   `json:"deb_pkg,omitempty"`
	DEBPG    []int    `json:"deb_pg,omitempty"`
	DEBDeps  []string `json:"deb_deps,omitempty"`
	DEBBuild bool     `json:"deb_build"`

	Stars    *int `json:"stars,omitempty"`
	Watchers *int `json:"watchers,omitempty"`
	Forks    *int `json:"forks,omitempty"`
	StarCnt  *int `json:"star_cnt,omitempty"`  // v1 compatibility
	WatchCnt *int `json:"watch_cnt,omitempty"` // v1 compatibility
	ForkCnt  *int `json:"fork_cnt,omitempty"`  // v1 compatibility

	RepoCreatedAt string `json:"repo_created_at,omitempty"`
	LastCommit    string `json:"last_commit,omitempty"`
	LastRelease   string `json:"last_release,omitempty"`
	LastActive    string `json:"last_active,omitempty"`
	CheckedAt     string `json:"checked_at,omitempty"`
	MTime         string `json:"mtime,omitempty"`

	EnDesc  string          `json:"en_desc,omitempty"`
	ZhDesc  string          `json:"zh_desc,omitempty"`
	Comment string          `json:"comment,omitempty"`
	Extra   json.RawMessage `json:"extra,omitempty"`

	// Compact package availability indices. Each value is
	// pgIndex*len(snapshot.OSs)+osIndex and is shared by a package family.
	TargetIdx  []int `json:"target_idx,omitempty"`
	FamilySize int   `json:"family_size"`

	HasEnDoc bool      `json:"has_en_doc"`
	HasZhDoc bool      `json:"has_zh_doc"`
	DocLinks *DocLinks `json:"doc_links,omitempty"`

	// Legacy aliases retained for API v1 clients.
	State      string `json:"state"`
	Repo       string `json:"repo,omitempty"`
	Source     string `json:"source,omitempty"`
	ExtType    string `json:"ext_type,omitempty"`
	ExtKernel  string `json:"ext_kernel,omitempty"`
	ExtVendor  string `json:"ext_vendor,omitempty"`
	LastUpdate string `json:"last_update,omitempty"`
}

// DocLinks carries the most useful upstream references as a compatibility
// bundle. The same URLs are also first-class Ext fields in the new schema.
type DocLinks struct {
	Repo    string `json:"repo_url,omitempty"`
	License string `json:"license_url,omitempty"`
	Control string `json:"control_url,omitempty"`
	Author  string `json:"author_url,omitempty"`
	Home    string `json:"home_url,omitempty"`
	Docs    string `json:"doc_url,omitempty"`
	Cargo   string `json:"cargo_url,omitempty"`
	PGXN    string `json:"pgxn_url,omitempty"`
}

// Category is one row of pgext.category with a member count.
type Category struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ZhDesc string `json:"zh_desc"`
	EnDesc string `json:"en_desc"`
	Count  int    `json:"count"`
}

// Snapshot is an immutable in-memory copy of the slowly-changing catalog.
// It doubles as the query cache: list, detail, and dimension endpoints are
// answered from here; only per-extension package/doc queries hit the pool.
type Snapshot struct {
	LoadedAt time.Time
	Version  string // stable content ETag for bootstrap/list payloads

	Exts   []*Ext
	ByName map[string]*Ext
	Cats   []Category
	PGs    []int    // active PostgreSQL majors, descending
	OSs    []string // active OS targets, canonical order

	CountPackaged        int
	CountProjects        int
	CountPackagedProject int
	CountVendor          int
	CountKernel          int
	CountContrib         int
	CountDocs            int
	CountSourceOnly      int
	CountPackages        int // distinct package families in pgext.pkg
	CountBuildSlots      int // pkg × pg × os combinations in pgext.pkg

	// CountAvail is kept internally for older tests/callers.
	CountAvail int
}

// Store holds the current snapshot and refreshes it in the background.
type Store struct {
	pool *pgxpool.Pool
	snap atomic.Pointer[Snapshot]

	mu         sync.Mutex // guards Reload
	lastReload time.Time
}

func NewStore(pool *pgxpool.Pool) *Store { return &Store{pool: pool} }

// Get returns the current snapshot (never nil after successful boot).
func (s *Store) Get() *Snapshot { return s.snap.Load() }

// Reload loads a fresh snapshot from the database and swaps it in.
// Calls within 5 seconds of the previous reload are coalesced into a no-op.
func (s *Store) Reload(ctx context.Context) (*Snapshot, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if cur := s.snap.Load(); cur != nil && time.Since(s.lastReload) < 5*time.Second {
		return cur, nil
	}
	snap, err := loadSnapshot(ctx, s.pool)
	if err != nil {
		return nil, err
	}
	s.snap.Store(snap)
	s.lastReload = time.Now()
	logrus.Infof("snapshot loaded: %d extensions, %d projects (%d packaged extensions), %d categories, pg=%v, %d os targets",
		len(snap.Exts), snap.CountProjects, snap.CountPackaged, len(snap.Cats), snap.PGs, len(snap.OSs))
	return snap, nil
}

// StartRefresher reloads the snapshot every ttl until ctx is done.
func (s *Store) StartRefresher(ctx context.Context, ttl time.Duration) {
	if ttl <= 0 {
		return
	}
	go func() {
		t := time.NewTicker(ttl)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				rctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				if _, err := s.Reload(rctx); err != nil {
					logrus.Warnf("snapshot refresh failed (serving stale data): %v", err)
				}
				cancel()
			}
		}
	}()
}

const universeQuery = `
SELECT id, name, pkg, lead_ext, category,
       packaged, lifecycle, kind, kernel, vendor, contrib, lang, license, tags,
       version, url, repo_url, home_url, doc_url, license_url, control_url, author_url, cargo_url, pgxn_url,
       has_bin, has_lib, need_ddl, need_load, trusted, relocatable,
       libs, schemas, pg_ver, requires, required_by, see_also,
       tarball, pgrx_ver,
       rpm_ver, rpm_repo, rpm_pkg, rpm_pg, rpm_deps, rpm_build,
       deb_ver, deb_repo, deb_pkg, deb_pg, deb_deps, deb_build,
       stars, watchers, forks,
       repo_created_at::text, last_commit::text, last_release::text, last_active::text, checked_at::text,
       en_desc, zh_desc, comment, extra, mtime::text,
       COALESCE(extra->>'repo',
         CASE WHEN contrib THEN 'CONTRIB'
              WHEN rpm_repo = deb_repo THEN rpm_repo
              WHEN rpm_repo IS NULL THEN deb_repo
              WHEN deb_repo IS NULL THEN rpm_repo
              ELSE 'MIXED' END,
         'n/a') AS repo
FROM pgext.universe`

func verifyUniverseSchema(ctx context.Context, pool *pgxpool.Pool) error {
	const q = `
	SELECT coalesce(string_agg(required, ', ' ORDER BY required), '')
	FROM unnest(ARRAY['packaged','lifecycle','kind','kernel','vendor','repo_url','required_by','rpm_build','deb_build','last_active','checked_at']) AS required
	WHERE NOT EXISTS (
		SELECT 1 FROM information_schema.columns
		WHERE table_schema = 'pgext' AND table_name = 'universe' AND column_name = required
	)`
	var missing string
	if err := pool.QueryRow(ctx, q).Scan(&missing); err != nil {
		return fmt.Errorf("inspect pgext.universe schema: %w", err)
	}
	if missing != "" {
		return fmt.Errorf("pgext.universe uses an incompatible schema (missing: %s); load the current db/schema.sql before starting pgext serve", missing)
	}
	return nil
}

func loadSnapshot(ctx context.Context, pool *pgxpool.Pool) (*Snapshot, error) {
	if err := verifyUniverseSchema(ctx, pool); err != nil {
		return nil, err
	}
	snap := &Snapshot{LoadedAt: time.Now().UTC(), ByName: map[string]*Ext{}}

	// Dimensions are loaded first because compact target indices depend on
	// their canonical order.
	if err := pool.QueryRow(ctx, `SELECT array_agg(pg ORDER BY pg DESC) FROM pgext.active_pg`).Scan(&snap.PGs); err != nil {
		return nil, fmt.Errorf("query pgext.active_pg: %w", err)
	}
	if err := pool.QueryRow(ctx, `SELECT array_agg(os) FROM pgext.active_os`).Scan(&snap.OSs); err != nil {
		return nil, fmt.Errorf("query pgext.active_os: %w", err)
	}
	sortOS(snap.OSs)

	rows, err := pool.Query(ctx, universeQuery)
	if err != nil {
		return nil, fmt.Errorf("query pgext.universe: %w", err)
	}
	defer rows.Close()
	byPkg := map[string][]*Ext{}
	packagedProjects := map[string]struct{}{}
	for rows.Next() {
		var e Ext
		var lifecycle, kernel, vendor, license, version *string
		var repoURL, homeURL, docURL, licenseURL, controlURL, authorURL, cargoURL, pgxnURL *string
		var hasBin, hasLib, needLoad, trusted, relocatable *bool
		var tarball, pgrxVer, rpmVer, rpmRepo, rpmPkg, debVer, debRepo, debPkg *string
		var repoCreated, lastCommit, lastRelease, lastActive, checkedAt *string
		var enDesc, zhDesc, comment *string
		var tags, libs, schemas, pgVer, requires, requiredBy, seeAlso, rpmDeps, debDeps []string
		var rpmPG, debPG []int16
		var extra []byte
		if err := rows.Scan(
			&e.ID, &e.Name, &e.Pkg, &e.LeadExt, &e.Category,
			&e.Packaged, &lifecycle, &e.Kind, &kernel, &vendor, &e.Contrib, &e.Lang, &license, &tags,
			&version, &e.URL, &repoURL, &homeURL, &docURL, &licenseURL, &controlURL, &authorURL, &cargoURL, &pgxnURL,
			&hasBin, &hasLib, &e.NeedDDL, &needLoad, &trusted, &relocatable,
			&libs, &schemas, &pgVer, &requires, &requiredBy, &seeAlso,
			&tarball, &pgrxVer,
			&rpmVer, &rpmRepo, &rpmPkg, &rpmPG, &rpmDeps, &e.RPMBuild,
			&debVer, &debRepo, &debPkg, &debPG, &debDeps, &e.DEBBuild,
			&e.Stars, &e.Watchers, &e.Forks,
			&repoCreated, &lastCommit, &lastRelease, &lastActive, &checkedAt,
			&enDesc, &zhDesc, &comment, &extra, &e.MTime, &e.Repo,
		); err != nil {
			return nil, fmt.Errorf("scan pgext.universe: %w", err)
		}
		e.Lifecycle, e.Kernel, e.Vendor, e.License = deref(lifecycle), deref(kernel), deref(vendor), deref(license)
		e.Version = deref(version)
		e.RepoURL, e.HomeURL, e.DocURL = deref(repoURL), deref(homeURL), deref(docURL)
		e.LicenseURL, e.ControlURL = deref(licenseURL), deref(controlURL)
		e.AuthorURL, e.CargoURL, e.PGXNURL = deref(authorURL), deref(cargoURL), deref(pgxnURL)
		e.HasBin, e.HasLib, e.NeedLoad = boolOf(hasBin), boolOf(hasLib), boolOf(needLoad)
		e.Trusted, e.Relocatable = boolOf(trusted), boolOf(relocatable)
		e.Lead = e.Name == e.LeadExt
		e.Tags, e.Libs, e.Schemas = tags, libs, schemas
		e.PG, e.Requires, e.RequiredBy, e.RequireBy, e.SeeAlso = atois(pgVer), requires, requiredBy, requiredBy, seeAlso
		e.Tarball, e.PGRXVer = deref(tarball), deref(pgrxVer)
		e.RPMVer, e.RPMRepo, e.RPMPkg, e.RPMPG, e.RPMDeps = deref(rpmVer), deref(rpmRepo), deref(rpmPkg), ints16(rpmPG), rpmDeps
		e.DEBVer, e.DEBRepo, e.DEBPkg, e.DEBPG, e.DEBDeps = deref(debVer), deref(debRepo), deref(debPkg), ints16(debPG), debDeps
		e.RepoCreatedAt, e.LastCommit, e.LastRelease = deref(repoCreated), deref(lastCommit), deref(lastRelease)
		e.LastActive, e.CheckedAt = deref(lastActive), deref(checkedAt)
		e.EnDesc, e.ZhDesc, e.Comment = deref(enDesc), deref(zhDesc), deref(comment)
		e.Extra = json.RawMessage(extra)

		// v1 aliases.
		if e.Packaged {
			e.State = "available"
		} else {
			e.State = "n/a"
		}
		e.Source, e.ExtType, e.ExtKernel, e.ExtVendor = e.Tarball, e.Kind, e.Kernel, e.Vendor
		e.LastUpdate = e.LastActive
		e.StarCnt, e.WatchCnt, e.ForkCnt = e.Stars, e.Watchers, e.Forks

		links := DocLinks{Repo: e.RepoURL, License: e.LicenseURL, Control: e.ControlURL, Author: e.AuthorURL, Home: e.HomeURL, Docs: e.DocURL, Cargo: e.CargoURL, PGXN: e.PGXNURL}
		if links != (DocLinks{}) {
			e.DocLinks = &links
		}

		snap.Exts = append(snap.Exts, &e)
		snap.ByName[e.Name] = &e
		byPkg[e.Pkg] = append(byPkg[e.Pkg], &e)
		if e.Packaged {
			snap.CountPackaged++
			snap.CountAvail++
			packagedProjects[e.Pkg] = struct{}{}
		} else if e.RepoURL != "" || e.Tarball != "" {
			snap.CountSourceOnly++
		}
		if e.Vendor != "" {
			snap.CountVendor++
		}
		if e.Kernel != "" {
			snap.CountKernel++
		}
		if e.Contrib {
			snap.CountContrib++
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate pgext.universe: %w", err)
	}
	if len(snap.Exts) == 0 {
		return nil, fmt.Errorf("pgext.universe is empty — is the pgext catalog initialized? (try `pgext init`)")
	}
	snap.CountProjects = len(byPkg)
	snap.CountPackagedProject = len(packagedProjects)

	// Package families and exact active target availability.
	for _, group := range byPkg {
		for _, e := range group {
			e.FamilySize = len(group)
			if len(group) > 1 {
				for _, sibling := range group {
					if sibling.Name != e.Name {
						e.Siblings = append(e.Siblings, sibling.Name)
					}
				}
				sort.Strings(e.Siblings)
			}
		}
	}
	for _, e := range snap.Exts {
		e.PreloadLibs = preloadLibraries(e, snap.ByName)
	}
	if err := loadTargets(ctx, pool, snap, byPkg); err != nil {
		return nil, err
	}

	// Default browse order: stars, credible activity, then name.
	sort.SliceStable(snap.Exts, func(i, j int) bool {
		si, sj := starOf(snap.Exts[i]), starOf(snap.Exts[j])
		if si != sj {
			return si > sj
		}
		if snap.Exts[i].LastActive != snap.Exts[j].LastActive {
			return snap.Exts[i].LastActive > snap.Exts[j].LastActive
		}
		return snap.Exts[i].Name < snap.Exts[j].Name
	})

	catCount := map[string]int{}
	for _, e := range snap.Exts {
		catCount[e.Category]++
	}
	crows, err := pool.Query(ctx, `SELECT id, name, zh_desc, en_desc FROM pgext.category ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("query pgext.category: %w", err)
	}
	defer crows.Close()
	for crows.Next() {
		var c Category
		var zh, en *string
		if err := crows.Scan(&c.ID, &c.Name, &zh, &en); err != nil {
			return nil, fmt.Errorf("scan pgext.category: %w", err)
		}
		c.ZhDesc, c.EnDesc, c.Count = deref(zh), deref(en), catCount[c.Name]
		snap.Cats = append(snap.Cats, c)
	}
	if err := crows.Err(); err != nil {
		return nil, fmt.Errorf("iterate pgext.category: %w", err)
	}

	// Only doc presence is fetched here; URLs already live in universe.
	drows, err := pool.Query(ctx, `
		SELECT ext, en_doc IS NOT NULL AND en_doc != '', zh_doc IS NOT NULL AND zh_doc != ''
		FROM pgext.doc`)
	if err != nil {
		logrus.Warnf("pgext.doc unavailable, usage docs disabled: %v", err)
	} else {
		defer drows.Close()
		for drows.Next() {
			var ext string
			var hasEn, hasZh bool
			if err := drows.Scan(&ext, &hasEn, &hasZh); err != nil {
				return nil, fmt.Errorf("scan pgext.doc: %w", err)
			}
			if e, ok := snap.ByName[ext]; ok {
				e.HasEnDoc, e.HasZhDoc = hasEn, hasZh
				if hasEn || hasZh {
					snap.CountDocs++
				}
			}
		}
		if err := drows.Err(); err != nil {
			return nil, fmt.Errorf("iterate pgext.doc: %w", err)
		}
	}

	// Stable content hash: refreshing identical data keeps browser and query
	// caches warm. LoadedAt is deliberately excluded.
	h := sha1.New()
	if err := json.NewEncoder(h).Encode(struct {
		Exts []*Ext
		Cats []Category
		PGs  []int
		OSs  []string
	}{snap.Exts, snap.Cats, snap.PGs, snap.OSs}); err != nil {
		return nil, fmt.Errorf("hash catalog snapshot: %w", err)
	}
	snap.Version = `"` + hex.EncodeToString(h.Sum(nil))[:16] + `"`
	return snap, nil
}

// preloadLibraries mirrors the cc/io page generators: preload dependencies
// first, then the extension itself, preserving order and removing duplicates.
func preloadLibraries(e *Ext, byName map[string]*Ext) []string {
	if e == nil || !e.NeedLoad {
		return nil
	}
	seen := map[string]bool{}
	libs := []string{}
	add := func(values []string, fallback string) {
		if len(values) == 0 && fallback != "" {
			values = []string{fallback}
		}
		for _, value := range values {
			for _, lib := range strings.Split(value, ",") {
				lib = strings.TrimSpace(lib)
				if lib != "" && !seen[lib] {
					seen[lib] = true
					libs = append(libs, lib)
				}
			}
		}
	}
	for _, name := range e.Requires {
		if dep := byName[name]; dep != nil && dep.NeedLoad {
			add(dep.Libs, dep.Name)
		}
	}
	add(e.Libs, e.Name)
	return libs
}

func loadTargets(ctx context.Context, pool *pgxpool.Pool, snap *Snapshot, byPkg map[string][]*Ext) error {
	pgIdx := make(map[int]int, len(snap.PGs))
	osIdx := make(map[string]int, len(snap.OSs))
	for i, pg := range snap.PGs {
		pgIdx[pg] = i
	}
	for i, os := range snap.OSs {
		osIdx[os] = i
	}
	if err := pool.QueryRow(ctx, `SELECT count(DISTINCT pkg), count(*) FROM pgext.pkg`).Scan(&snap.CountPackages, &snap.CountBuildSlots); err != nil {
		return fmt.Errorf("count pgext.pkg: %w", err)
	}
	targets := map[string]map[int]struct{}{}
	rows, err := pool.Query(ctx, `SELECT pkg, pg, os FROM pgext.pkg WHERE state::text = 'AVAIL'`)
	if err != nil {
		return fmt.Errorf("query pgext.pkg targets: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var pkg, os string
		var pg int
		if err := rows.Scan(&pkg, &pg, &os); err != nil {
			return fmt.Errorf("scan pgext.pkg targets: %w", err)
		}
		pi, pok := pgIdx[pg]
		oi, ook := osIdx[os]
		if !pok || !ook {
			continue
		}
		if targets[pkg] == nil {
			targets[pkg] = map[int]struct{}{}
		}
		targets[pkg][pi*len(snap.OSs)+oi] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate pgext.pkg targets: %w", err)
	}
	for pkg, set := range targets {
		idx := make([]int, 0, len(set))
		for n := range set {
			idx = append(idx, n)
		}
		sort.Ints(idx)
		for _, e := range byPkg[pkg] {
			e.TargetIdx = idx
		}
	}
	return nil
}

// sortOS orders OS targets canonically: EL before Debian before Ubuntu,
// then by release number, with x86_64 ahead of aarch64.
func sortOS(oss []string) {
	rank := func(os string) (int, int, int) {
		family, num := 0, 0
		s := os
		switch {
		case strings.HasPrefix(os, "el"):
			family, s = 0, strings.TrimPrefix(os, "el")
		case strings.HasPrefix(os, "d"):
			family, s = 1, strings.TrimPrefix(os, "d")
		case strings.HasPrefix(os, "u"):
			family, s = 2, strings.TrimPrefix(os, "u")
		}
		numStr, arch, _ := strings.Cut(s, ".")
		num, _ = strconv.Atoi(numStr)
		archRank := 0
		if arch == "aarch64" {
			archRank = 1
		}
		return family, num, archRank
	}
	sort.SliceStable(oss, func(i, j int) bool {
		fi, ni, ai := rank(oss[i])
		fj, nj, aj := rank(oss[j])
		if fi != fj {
			return fi < fj
		}
		if ni != nj {
			return ni < nj
		}
		return ai < aj
	})
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func boolOf(b *bool) bool { return b != nil && *b }

func atois(ss []string) []int {
	if len(ss) == 0 {
		return nil
	}
	out := make([]int, 0, len(ss))
	for _, s := range ss {
		if n, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			out = append(out, n)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	return out
}

func ints16(ss []int16) []int {
	if len(ss) == 0 {
		return nil
	}
	out := make([]int, len(ss))
	for i, n := range ss {
		out[i] = int(n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	return out
}

func starOf(e *Ext) int {
	if e.Stars == nil {
		return -1
	}
	return *e.Stars
}
