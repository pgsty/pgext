/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
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

// Ext is one extension record from pgext.universe, enriched with doc metadata.
type Ext struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Pkg      string `json:"pkg"`
	LeadExt  string `json:"lead_ext,omitempty"`
	Category string `json:"category"`
	State    string `json:"state"`
	URL      string `json:"url,omitempty"`
	License  string `json:"license,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Repo     string `json:"repo,omitempty"`
	Version  string `json:"version,omitempty"`

	Tags        []string `json:"tags,omitempty"`
	Contrib     bool     `json:"contrib"`
	Lead        bool     `json:"lead"`
	HasBin      bool     `json:"has_bin"`
	HasLib      bool     `json:"has_lib"`
	NeedDDL     bool     `json:"need_ddl"`
	NeedLoad    bool     `json:"need_load"`
	Trusted     bool     `json:"trusted"`
	Relocatable bool     `json:"relocatable"`

	Schemas   []string `json:"schemas,omitempty"`
	PG        []int    `json:"pg_ver,omitempty"`
	Requires  []string `json:"requires,omitempty"`
	RequireBy []string `json:"require_by,omitempty"`
	SeeAlso   []string `json:"see_also,omitempty"`
	Siblings  []string `json:"siblings,omitempty"`

	RPMVer  string   `json:"rpm_ver,omitempty"`
	RPMRepo string   `json:"rpm_repo,omitempty"`
	RPMPkg  string   `json:"rpm_pkg,omitempty"`
	RPMPG   []int    `json:"rpm_pg,omitempty"`
	RPMDeps []string `json:"rpm_deps,omitempty"`
	DEBVer  string   `json:"deb_ver,omitempty"`
	DEBRepo string   `json:"deb_repo,omitempty"`
	DEBPkg  string   `json:"deb_pkg,omitempty"`
	DEBPG   []int    `json:"deb_pg,omitempty"`
	DEBDeps []string `json:"deb_deps,omitempty"`

	Source    string `json:"source,omitempty"`
	ExtType   string `json:"ext_type,omitempty"`
	ExtKernel string `json:"ext_kernel,omitempty"`
	ExtVendor string `json:"ext_vendor,omitempty"`

	Stars *int `json:"star_cnt,omitempty"`
	Watch *int `json:"watch_cnt,omitempty"`
	Fork  *int `json:"fork_cnt,omitempty"`

	LastCommit  string `json:"last_commit,omitempty"`
	LastRelease string `json:"last_release,omitempty"`
	LastUpdate  string `json:"last_update,omitempty"`

	EnDesc string `json:"en_desc,omitempty"`
	ZhDesc string `json:"zh_desc,omitempty"`

	HasEnDoc bool      `json:"has_en_doc"`
	HasZhDoc bool      `json:"has_zh_doc"`
	DocLinks *DocLinks `json:"doc_links,omitempty"`
}

// DocLinks carries the reference URLs curated in pgext.doc.
type DocLinks struct {
	Repo    string `json:"repo_url,omitempty"`
	License string `json:"license_url,omitempty"`
	Control string `json:"control_url,omitempty"`
	Author  string `json:"author_url,omitempty"`
	Home    string `json:"home_url,omitempty"`
	Cargo   string `json:"cargo_url,omitempty"`
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
	Version  string // strong ETag for bootstrap/list payloads

	Exts   []*Ext
	ByName map[string]*Ext
	Cats   []Category
	PGs    []int    // active PostgreSQL majors, descending
	OSs    []string // active OS targets, canonical order

	CountAvail  int
	CountVendor int
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
	logrus.Infof("snapshot loaded: %d extensions (%d packaged, %d vendor), %d categories, pg=%v, %d os targets",
		len(snap.Exts), snap.CountAvail, snap.CountVendor, len(snap.Cats), snap.PGs, len(snap.OSs))
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
SELECT id, name, pkg, lead_ext, category, state, url, license, tags, version, repo, lang,
       contrib, lead, has_bin, has_lib, need_ddl, need_load, trusted, relocatable,
       schemas, pg_ver, requires, require_by, see_also,
       rpm_ver, rpm_repo, rpm_pkg, rpm_pg, rpm_deps,
       deb_ver, deb_repo, deb_pkg, deb_pg, deb_deps,
       source, ext_type, ext_kernel, ext_vendor,
       star_cnt, watch_cnt, fork_cnt,
       last_commit_date::text, last_release_date::text, last_update_date::text,
       en_desc, zh_desc
FROM pgext.universe`

func loadSnapshot(ctx context.Context, pool *pgxpool.Pool) (*Snapshot, error) {
	snap := &Snapshot{LoadedAt: time.Now(), ByName: map[string]*Ext{}}

	rows, err := pool.Query(ctx, universeQuery)
	if err != nil {
		return nil, fmt.Errorf("query pgext.universe: %w", err)
	}
	defer rows.Close()
	byPkg := map[string][]*Ext{}
	for rows.Next() {
		var e Ext
		var pkg, leadExt, url, license, version, repo, lang *string
		var rpmVer, rpmRepo, rpmPkg, debVer, debRepo, debPkg *string
		var source, extType, extKernel, extVendor *string
		var lastCommit, lastRelease, lastUpdate, enDesc, zhDesc *string
		var contrib, lead, hasBin, hasLib, needDDL, needLoad, trusted, reloc *bool
		var tags, schemas, pgVer, requires, requireBy, seeAlso, rpmPG, rpmDeps, debPG, debDeps []string
		if err := rows.Scan(
			&e.ID, &e.Name, &pkg, &leadExt, &e.Category, &e.State, &url, &license, &tags, &version, &repo, &lang,
			&contrib, &lead, &hasBin, &hasLib, &needDDL, &needLoad, &trusted, &reloc,
			&schemas, &pgVer, &requires, &requireBy, &seeAlso,
			&rpmVer, &rpmRepo, &rpmPkg, &rpmPG, &rpmDeps,
			&debVer, &debRepo, &debPkg, &debPG, &debDeps,
			&source, &extType, &extKernel, &extVendor,
			&e.Stars, &e.Watch, &e.Fork,
			&lastCommit, &lastRelease, &lastUpdate,
			&enDesc, &zhDesc,
		); err != nil {
			return nil, fmt.Errorf("scan pgext.universe: %w", err)
		}
		e.Contrib, e.Lead, e.HasBin, e.HasLib = boolOf(contrib), boolOf(lead), boolOf(hasBin), boolOf(hasLib)
		e.NeedDDL, e.NeedLoad, e.Trusted, e.Relocatable = boolOf(needDDL), boolOf(needLoad), boolOf(trusted), boolOf(reloc)
		e.Pkg = strOr(pkg, e.Name)
		e.LeadExt = strOr(leadExt, e.Name)
		e.URL = deref(url)
		e.License = deref(license)
		e.Version = deref(version)
		e.Repo = deref(repo)
		e.Lang = deref(lang)
		e.Tags = tags
		e.Schemas = schemas
		e.PG = atois(pgVer)
		e.Requires = requires
		e.RequireBy = requireBy
		e.SeeAlso = seeAlso
		e.RPMVer, e.RPMRepo, e.RPMPkg = deref(rpmVer), deref(rpmRepo), deref(rpmPkg)
		e.RPMPG, e.RPMDeps = atois(rpmPG), rpmDeps
		e.DEBVer, e.DEBRepo, e.DEBPkg = deref(debVer), deref(debRepo), deref(debPkg)
		e.DEBPG, e.DEBDeps = atois(debPG), debDeps
		e.Source = deref(source)
		e.ExtType = deref(extType)
		e.ExtKernel = deref(extKernel)
		e.ExtVendor = deref(extVendor)
		e.LastCommit, e.LastRelease, e.LastUpdate = deref(lastCommit), deref(lastRelease), deref(lastUpdate)
		e.EnDesc, e.ZhDesc = deref(enDesc), deref(zhDesc)
		snap.Exts = append(snap.Exts, &e)
		snap.ByName[e.Name] = &e
		byPkg[e.Pkg] = append(byPkg[e.Pkg], &e)
		if e.State == "available" {
			snap.CountAvail++
		}
		if e.ExtKernel != "" || e.ExtVendor != "" {
			snap.CountVendor++
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate pgext.universe: %w", err)
	}
	if len(snap.Exts) == 0 {
		return nil, fmt.Errorf("pgext.universe is empty — is the pgext catalog initialized? (try `pgext init`)")
	}

	// siblings: extensions delivered by the same package
	for _, e := range snap.Exts {
		if group := byPkg[e.Pkg]; len(group) > 1 {
			for _, g := range group {
				if g.Name != e.Name {
					e.Siblings = append(e.Siblings, g.Name)
				}
			}
			sort.Strings(e.Siblings)
		}
	}

	// default order: stars desc, then name — the canonical browse order
	sort.SliceStable(snap.Exts, func(i, j int) bool {
		si, sj := starOf(snap.Exts[i]), starOf(snap.Exts[j])
		if si != sj {
			return si > sj
		}
		return snap.Exts[i].Name < snap.Exts[j].Name
	})

	// categories with member counts
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
		c.ZhDesc, c.EnDesc = deref(zh), deref(en)
		c.Count = catCount[c.Name]
		snap.Cats = append(snap.Cats, c)
	}
	if err := crows.Err(); err != nil {
		return nil, err
	}

	// active dimensions
	if err := pool.QueryRow(ctx, `SELECT array_agg(pg ORDER BY pg DESC) FROM pgext.active_pg`).Scan(&snap.PGs); err != nil {
		return nil, fmt.Errorf("query pgext.active_pg: %w", err)
	}
	if err := pool.QueryRow(ctx, `SELECT array_agg(os) FROM pgext.active_os`).Scan(&snap.OSs); err != nil {
		return nil, fmt.Errorf("query pgext.active_os: %w", err)
	}
	sortOS(snap.OSs)

	// doc metadata (text bodies are fetched on demand, links live in memory)
	drows, err := pool.Query(ctx, `
		SELECT ext, repo_url, license_url, control_url, author_url, home_url, cargo_url,
		       en_doc IS NOT NULL AND en_doc != '', zh_doc IS NOT NULL AND zh_doc != ''
		FROM pgext.doc`)
	if err != nil {
		logrus.Warnf("pgext.doc unavailable, usage docs disabled: %v", err)
	} else {
		defer drows.Close()
		for drows.Next() {
			var ext string
			var l DocLinks
			var repo, license, control, author, home, cargo *string
			var hasEn, hasZh bool
			if err := drows.Scan(&ext, &repo, &license, &control, &author, &home, &cargo, &hasEn, &hasZh); err != nil {
				return nil, fmt.Errorf("scan pgext.doc: %w", err)
			}
			l.Repo, l.License, l.Control = deref(repo), deref(license), deref(control)
			l.Author, l.Home, l.Cargo = deref(author), deref(home), deref(cargo)
			if e, ok := snap.ByName[ext]; ok {
				e.HasEnDoc, e.HasZhDoc = hasEn, hasZh
				if l != (DocLinks{}) {
					e.DocLinks = &l
				}
			}
		}
		if err := drows.Err(); err != nil {
			return nil, err
		}
	}

	// version stamp: content-addressed enough for ETag purposes
	h := sha1.New()
	fmt.Fprintf(h, "%d/%d/%d/%s", len(snap.Exts), snap.CountAvail, snap.CountVendor, snap.LoadedAt.Format(time.RFC3339))
	snap.Version = `"` + hex.EncodeToString(h.Sum(nil))[:16] + `"`
	return snap, nil
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

func strOr(s *string, def string) string {
	if s == nil || *s == "" {
		return def
	}
	return *s
}

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

func starOf(e *Ext) int {
	if e.Stars == nil {
		return -1
	}
	return *e.Stars
}
