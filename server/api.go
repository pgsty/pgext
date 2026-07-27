/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"cmp"
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

/* ----------------------------------------------------------------
   API surface (all JSON, read-only except /reload):

     GET  /api/v1/meta                  server + snapshot + dimensions
     GET  /api/v1/ext                   list/search: q + catalog/runtime/build/doc/relation dimensions
     GET  /api/v1/ext/{name}            one extension, full record
     GET  /api/v1/ext/{name}/matrix     package availability: pg × os cells
     GET  /api/v1/matrix                global package availability: package × os × pg
     GET  /api/v1/ext/{name}/files      binary artifacts with download URLs (?pg= &os=)
     GET  /api/v1/ext/{name}/doc        usage doc markdown (?lang=en|zh)
     GET  /api/v1/ext/{name}/changelog  release history rows for the package
     GET  /api/v1/dim/{key}             aggregate: 19 universe/doc/pkg dimensions
     GET  /api/v1/bootstrap             compact positional payload for the SPA
     POST /api/v1/reload                refresh the snapshot from the database
     GET  /healthz                      liveness + snapshot age
   ---------------------------------------------------------------- */

// Filter mirrors the SPA search semantics: free words plus key:value operators.
type Filter struct {
	Words      []string
	Cat        string
	Repo       string
	License    string
	Lang       string
	PG         []int
	OS         string
	Kind       string
	Lifecycle  string
	Kernel     string
	Vendor     string
	Scope      string // "", "packaged", "source", "kernel", "vendor", "contrib"
	Tag        string
	Pkg        string
	Capability string
	Build      string
	Docs       string
	Relation   string
	PGRX       string
	Active     string
}

// ParseFilter reads filters from query params; `q` may embed operators
// like `cat:GIS tag:vector pg:18 build:pgrx is:packaged`.
func ParseFilter(v url.Values) Filter {
	f := Filter{
		Cat:        strings.ToUpper(cmp.Or(v.Get("cat"), v.Get("cate"))),
		Repo:       strings.ToUpper(v.Get("repo")),
		License:    v.Get("license"),
		Lang:       cmp.Or(v.Get("lang"), v.Get("lng")),
		OS:         v.Get("os"),
		Kind:       strings.ToLower(cmp.Or(v.Get("kind"), v.Get("type"))),
		Lifecycle:  strings.ToLower(v.Get("lifecycle")),
		Kernel:     v.Get("kernel"),
		Vendor:     v.Get("vendor"),
		Scope:      strings.ToLower(v.Get("scope")),
		Tag:        v.Get("tag"),
		Pkg:        cmp.Or(v.Get("pkg"), v.Get("package")),
		Capability: strings.ToLower(cmp.Or(v.Get("capability"), v.Get("cap"))),
		Build:      strings.ToLower(v.Get("build")),
		Docs:       strings.ToLower(cmp.Or(v.Get("docs"), v.Get("doc"))),
		Relation:   strings.ToLower(cmp.Or(v.Get("relation"), v.Get("rel"))),
		PGRX:       v.Get("pgrx"),
		Active:     strings.ToLower(cmp.Or(v.Get("active"), v.Get("year"))),
	}
	if f.Scope == "cloud" { // legacy query alias
		f.Scope = "vendor"
	}
	f.PG = parsePGValues(v.Get("pg"))
	for _, tok := range strings.Fields(v.Get("q")) {
		key, val, ok := strings.Cut(tok, ":")
		if !ok || val == "" {
			f.Words = append(f.Words, strings.ToLower(tok))
			continue
		}
		switch strings.ToLower(key) {
		case "cat", "cate", "category":
			f.Cat = strings.ToUpper(val)
		case "repo":
			f.Repo = strings.ToUpper(val)
		case "license":
			f.License = val
		case "lang", "lng":
			f.Lang = val
		case "pg":
			f.PG = parsePGValues(val)
		case "type", "kind":
			f.Kind = strings.ToLower(val)
		case "lifecycle", "life":
			f.Lifecycle = strings.ToLower(val)
		case "kernel":
			f.Kernel = val
		case "os", "target":
			f.OS = val
		case "vendor":
			f.Vendor = val
		case "tag", "tags":
			f.Tag = val
		case "pkg", "package", "project":
			f.Pkg = val
		case "cap", "capability", "feature":
			f.Capability = strings.ToLower(val)
		case "build", "builder":
			f.Build = strings.ToLower(val)
		case "doc", "docs":
			f.Docs = strings.ToLower(val)
		case "rel", "relation", "dependency":
			f.Relation = strings.ToLower(val)
		case "pgrx":
			f.PGRX = val
		case "active", "year":
			f.Active = strings.ToLower(val)
		case "is":
			switch strings.ToLower(val) {
			case "packaged":
				f.Scope = "packaged"
			case "unpacked", "unpackaged":
				f.Scope = "unpacked"
			case "source", "source-only":
				f.Scope = "source"
			case "kernel":
				f.Scope = "kernel"
			case "vendor", "cloud":
				f.Scope = "vendor"
			case "contrib":
				f.Scope = "contrib"
			case "binary", "library", "create-extension", "preload", "trusted", "relocatable":
				f.Capability = strings.ToLower(val)
			}
		default:
			f.Words = append(f.Words, strings.ToLower(tok))
		}
	}
	return f
}

func (f Filter) match(e *Ext, snap *Snapshot) (int, bool) {
	if f.Cat != "" && e.Category != f.Cat {
		return 0, false
	}
	if f.Repo != "" && !strings.EqualFold(e.Repo, f.Repo) {
		return 0, false
	}
	if f.License != "" {
		if strings.EqualFold(f.License, "Unknown") {
			if e.License != "" && !strings.EqualFold(e.License, "Unknown") {
				return 0, false
			}
		} else if !strings.EqualFold(e.License, f.License) {
			return 0, false
		}
	}
	if f.Lang != "" && !strings.EqualFold(e.Lang, f.Lang) {
		return 0, false
	}
	if len(f.PG) > 0 && !containsAllInts(e.PG, f.PG) {
		return 0, false
	}
	if f.OS != "" && !hasTargets(e, snap, f.PG, f.OS) {
		return 0, false
	}
	if f.Kind != "" && !strings.EqualFold(e.Kind, f.Kind) {
		return 0, false
	}
	if f.Lifecycle != "" && !strings.EqualFold(e.Lifecycle, f.Lifecycle) {
		return 0, false
	}
	if f.Tag != "" && !containsEqualFold(e.Tags, f.Tag) {
		return 0, false
	}
	if f.Pkg != "" && !strings.EqualFold(e.Pkg, f.Pkg) {
		return 0, false
	}
	if f.Capability != "" && !matchesCapability(e, f.Capability) {
		return 0, false
	}
	if f.Build != "" && !matchesBuild(e, f.Build) {
		return 0, false
	}
	if f.Docs != "" && docClass(e) != f.Docs {
		return 0, false
	}
	if f.Relation != "" && !matchesRelation(e, f.Relation) {
		return 0, false
	}
	if f.PGRX != "" && !strings.EqualFold(e.PGRXVer, f.PGRX) {
		return 0, false
	}
	if f.Active != "" && activityYear(e) != f.Active {
		return 0, false
	}
	if !matchesScope(e, f.Scope) {
		return 0, false
	}
	if f.Kernel != "" && !strings.Contains(strings.ToLower(e.Kernel), strings.ToLower(f.Kernel)) {
		return 0, false
	}
	if f.Vendor != "" {
		v := strings.ToLower(f.Vendor)
		if !strings.Contains(strings.ToLower(e.Vendor), v) {
			return 0, false
		}
	}
	score := 0
	for _, w := range f.Words {
		name, pkg := strings.ToLower(e.Name), strings.ToLower(e.Pkg)
		switch {
		case name == w:
			score += 120
		case strings.HasPrefix(name, w):
			score += 60
		case strings.Contains(name, w) || strings.Contains(pkg, w):
			score += 30
		case containsFold(e.Tags, w):
			score += 18
		case strings.Contains(strings.ToLower(e.EnDesc), w) || strings.Contains(e.ZhDesc, w) || strings.Contains(strings.ToLower(e.Comment), w):
			score += 10
		case strings.EqualFold(e.Category, w),
			strings.Contains(strings.ToLower(e.Vendor), w),
			strings.Contains(strings.ToLower(e.Kernel), w),
			strings.Contains(strings.ToLower(e.Lifecycle), w):
			score += 8
		default:
			return 0, false
		}
	}
	return score, true
}

// Apply filters and ranks the snapshot: relevance first when free words are
// present, then the requested sort key (stars | name | updated).
func (f Filter) Apply(snap *Snapshot, sortKey string) []*Ext {
	type hit struct {
		e     *Ext
		score int
	}
	hits := make([]hit, 0, 64)
	for _, e := range snap.Exts {
		if score, ok := f.match(e, snap); ok {
			hits = append(hits, hit{e, score})
		}
	}
	byQ := len(f.Words) > 0
	sort.SliceStable(hits, func(i, j int) bool {
		if byQ && hits[i].score != hits[j].score {
			return hits[i].score > hits[j].score
		}
		a, b := hits[i].e, hits[j].e
		switch sortKey {
		case "name":
			return a.Name < b.Name
		case "updated":
			if a.LastActive != b.LastActive {
				return a.LastActive > b.LastActive
			}
			return starOf(a) > starOf(b)
		default: // stars
			if sa, sb := starOf(a), starOf(b); sa != sb {
				return sa > sb
			}
			return a.Name < b.Name
		}
	})
	out := make([]*Ext, len(hits))
	for i, h := range hits {
		out[i] = h.e
	}
	return out
}

func containsFold(values []string, needle string) bool {
	for _, value := range values {
		if strings.Contains(strings.ToLower(value), needle) {
			return true
		}
	}
	return false
}

func containsEqualFold(values []string, needle string) bool {
	for _, value := range values {
		if strings.EqualFold(value, needle) {
			return true
		}
	}
	return false
}

// matchesScope narrows by distribution scope; an empty or unknown scope value
// filters nothing, mirroring the other dimension parameters.
func matchesScope(e *Ext, scope string) bool {
	switch scope {
	case "packaged":
		return e.Packaged
	case "unpacked":
		return !e.Packaged
	case "source":
		return !e.Packaged && (e.RepoURL != "" || e.Tarball != "")
	case "kernel":
		return e.Kernel != ""
	case "vendor":
		return e.Vendor != ""
	case "contrib":
		return e.Contrib
	default:
		return true
	}
}

func matchesCapability(e *Ext, value string) bool {
	switch strings.ToLower(value) {
	case "binary", "bin":
		return e.HasBin
	case "library", "lib":
		return e.HasLib
	case "create-extension", "ddl":
		return e.NeedDDL
	case "preload", "load":
		return e.NeedLoad
	case "trusted":
		return e.Trusted
	case "relocatable":
		return e.Relocatable
	default:
		return false
	}
}

func matchesBuild(e *Ext, value string) bool {
	switch strings.ToLower(value) {
	case "rpm":
		return e.RPMBuild
	case "deb":
		return e.DEBBuild
	case "pgrx":
		return e.PGRXVer != ""
	case "source":
		return e.RepoURL != "" || e.Tarball != ""
	case "packaged", "binary-package":
		return e.Packaged
	default:
		return false
	}
}

func docClass(e *Ext) string {
	switch {
	case e.HasEnDoc && e.HasZhDoc:
		return "bilingual"
	case e.HasEnDoc:
		return "english-only"
	case e.HasZhDoc:
		return "chinese-only"
	default:
		return "undocumented"
	}
}

func matchesRelation(e *Ext, value string) bool {
	switch strings.ToLower(value) {
	case "requires":
		return len(e.Requires) > 0
	case "required-by":
		return len(e.RequiredBy) > 0
	case "see-also":
		return len(e.SeeAlso) > 0
	case "no-relations":
		return len(e.Requires) == 0 && len(e.RequiredBy) == 0 && len(e.SeeAlso) == 0
	default:
		return false
	}
}

func activityYear(e *Ext) string {
	if len(e.LastActive) >= 4 {
		year := e.LastActive[:4]
		if _, err := strconv.Atoi(year); err == nil {
			return year
		}
	}
	return "unknown"
}

func hasTarget(e *Ext, snap *Snapshot, pg int, os string) bool {
	if len(e.TargetIdx) == 0 || len(snap.OSs) == 0 {
		return false
	}
	osIndex := slices.Index(snap.OSs, os)
	if osIndex < 0 {
		return false
	}
	for _, idx := range e.TargetIdx {
		pi, oi := idx/len(snap.OSs), idx%len(snap.OSs)
		if oi == osIndex && pi >= 0 && pi < len(snap.PGs) && (pg == 0 || snap.PGs[pi] == pg) {
			return true
		}
	}
	return false
}

func hasTargets(e *Ext, snap *Snapshot, pgs []int, os string) bool {
	if len(pgs) == 0 {
		return hasTarget(e, snap, 0, os)
	}
	for _, pg := range pgs {
		if !hasTarget(e, snap, pg, os) {
			return false
		}
	}
	return true
}

func parsePGValues(raw string) []int {
	var values []int
	for _, part := range strings.FieldsFunc(raw, func(r rune) bool { return r == ',' || r == ';' || r == ' ' }) {
		if pg, err := strconv.Atoi(part); err == nil && pg > 0 && pg < 100 && !slices.Contains(values, pg) {
			values = append(values, pg)
		}
	}
	sortDesc(values)
	return values
}

func containsAllInts(xs, wanted []int) bool {
	for _, value := range wanted {
		if !slices.Contains(xs, value) {
			return false
		}
	}
	return true
}

/* ---------------- handlers ---------------- */

type api struct {
	store       *Store
	cache       *ttlCache
	pool        *pgxpool.Pool
	reloadToken string
	counter     *counterStore
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		logrus.Debugf("response encode failed: %v", err)
	}
}

func writeErr(w http.ResponseWriter, code int, format string, args ...any) {
	writeJSON(w, code, map[string]string{"error": fmt.Sprintf(format, args...)})
}

// extOf resolves the {name} path segment against the snapshot; on a miss it
// writes the JSON 404 and returns a nil extension.
func (a *api) extOf(w http.ResponseWriter, r *http.Request) (*Snapshot, *Ext) {
	snap := a.store.Get()
	if e, ok := snap.ByName[r.PathValue("name")]; ok {
		return snap, e
	}
	writeErr(w, http.StatusNotFound, "extension %q does not exist", r.PathValue("name"))
	return snap, nil
}

func (a *api) handleMeta(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	writeJSON(w, 200, map[string]any{
		"server":     "pgext",
		"generated":  snap.LoadedAt.Format(time.RFC3339),
		"dimensions": dimKeys,
		"counts":     snap.counts(),
		"pg":         snap.PGs,
		"os":         snap.OSs,
		"categories": snap.Cats,
	})
}

func (a *api) handleList(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	q := r.URL.Query()
	f := ParseFilter(q)
	sortKey := q.Get("sort")

	limit := 50
	if s := q.Get("limit"); s != "" {
		if s == "all" {
			limit = 5000
		} else if n, err := strconv.Atoi(s); err == nil && n > 0 {
			limit = min(n, 5000)
		}
	}
	offset := 0
	if n, err := strconv.Atoi(q.Get("offset")); err == nil && n > 0 {
		offset = n
	}

	list := f.Apply(snap, sortKey)
	total := len(list)
	if offset > total {
		offset = total
	}
	end := min(offset+limit, total)

	w.Header().Set("Cache-Control", "public, max-age=60")
	writeJSON(w, 200, map[string]any{
		"total":  total,
		"limit":  limit,
		"offset": offset,
		"data":   list[offset:end],
	})
}

func (a *api) handleExt(w http.ResponseWriter, r *http.Request) {
	_, e := a.extOf(w, r)
	if e == nil {
		return
	}
	w.Header().Set("Cache-Control", "public, max-age=60")
	writeJSON(w, 200, e)
}

// MatrixCell is one (pg, os) slot of the availability matrix.
type MatrixCell struct {
	PG      int    `json:"pg"`
	OS      string `json:"os"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state"`
	Org     string `json:"org,omitempty"`
	Version string `json:"version,omitempty"`
	Count   int    `json:"count"`
}

func (a *api) handleMatrix(w http.ResponseWriter, r *http.Request) {
	snap, e := a.extOf(w, r)
	if e == nil {
		return
	}
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		rows, _ := a.pool.Query(ctx, `
			SELECT pg, os, coalesce(name,''), state::text, coalesce(org,''), coalesce(version,''), coalesce(count,0)
			FROM pgext.pkg WHERE pkg = $1 ORDER BY pg DESC, os`, e.Pkg)
		cells, err := pgx.CollectRows(rows, pgx.RowToStructByPos[MatrixCell])
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"ext": e.Name, "pkg": e.Pkg,
			"pg": snap.PGs, "os": snap.OSs,
			"cells": cells,
		}, nil
	})
}

// handleGlobalMatrix serves the write-time-computed matrix payload (see
// server/matrix.go). The request path is one primary-key lookup on
// pgext.matrix_cache behind the in-memory TTL cache; the payload is compact
// and positional: p=package, e=lead extension, c=status bytes (lowercase =
// hidden), v/i=row-local version dictionary and per-cell index.
// ?repo=pgdg|pigsty selects the single-repository what-if view.
func (a *api) handleGlobalMatrix(w http.ResponseWriter, r *http.Request) {
	view := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("repo")))
	mv, ok := matrixViewOf(view)
	if !ok {
		writeErr(w, http.StatusBadRequest, "unknown matrix repo view %q (want pgdg or pigsty)", view)
		return
	}
	snap := a.store.Get()
	// LoadedAt participates in this cache key because package version/count
	// changes can be meaningful even when the catalog content hash is stable.
	cacheVersion := snap.Version + "|" + snap.LoadedAt.Format(time.RFC3339Nano)
	a.cached(w, r, cacheVersion, func(ctx context.Context) (any, error) {
		payload, err := a.matrixPayload(ctx, false, mv)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(payload), nil
	})
}

// FileRow is one downloadable binary artifact.
type FileRow struct {
	PG      int    `json:"pg"`
	OS      string `json:"os"`
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Org     string `json:"org,omitempty"`
	Ver     string `json:"ver"`
	Version string `json:"version,omitempty"`
	File    string `json:"file"`
	SHA256  string `json:"sha256,omitempty"`
	Size    int    `json:"size"`
	URL     string `json:"url"`
}

func (a *api) handleFiles(w http.ResponseWriter, r *http.Request) {
	snap, e := a.extOf(w, r)
	if e == nil {
		return
	}
	q := r.URL.Query()
	pg, _ := strconv.Atoi(q.Get("pg"))
	osFilter := q.Get("os")
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		rows, _ := a.pool.Query(ctx, `
			SELECT b.pg, b.os, b.name, b.repo, coalesce(r.org,''), coalesce(b.ver,''), coalesce(b.version,''),
			       coalesce(b.file,''), coalesce(b.sha256,''), coalesce(b.size,0),
			       CASE WHEN b.href ~ '^https?://' THEN b.href
			            ELSE coalesce(r.default_url,'') || '/' || coalesce(b.href,'') END
			FROM pgext.bin b JOIN pgext.repository r ON b.repo = r.id
			WHERE b.name IN (SELECT DISTINCT name FROM pgext.pkg WHERE pkg = $1 AND name IS NOT NULL)
			  AND ($2 = 0 OR b.pg = $2) AND ($3 = '' OR b.os = $3)
			ORDER BY b.pg DESC, b.os, b.name, b.ver DESC
			LIMIT 5000`, e.Pkg, pg, osFilter)
		files, err := pgx.CollectRows(rows, pgx.RowToStructByPos[FileRow])
		if err != nil {
			return nil, err
		}
		return map[string]any{"ext": e.Name, "pkg": e.Pkg, "total": len(files), "files": files}, nil
	})
}

func (a *api) handleDoc(w http.ResponseWriter, r *http.Request) {
	snap, e := a.extOf(w, r)
	if e == nil {
		return
	}
	lang, col := "en", "en_doc"
	if r.URL.Query().Get("lang") == "zh" {
		lang, col = "zh", "zh_doc"
	}
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		var doc *string
		err := a.pool.QueryRow(ctx, `SELECT `+col+` FROM pgext.doc WHERE ext = $1`, e.Name).Scan(&doc)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		if doc == nil || *doc == "" {
			return nil, errNotFound{fmt.Sprintf("no %s doc for %q", lang, e.Name)}
		}
		return map[string]any{"ext": e.Name, "lang": lang, "content": *doc}, nil
	})
}

var dimKeys = []string{
	"category", "tag", "package", "kind", "lifecycle", "license", "lang",
	"distribution", "repo", "pg", "os", "build", "pgrx",
	"capability", "docs", "relation", "vendor", "kernel", "activity",
}

func (a *api) handleDim(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	key := r.PathValue("key")
	type val struct {
		Value   string   `json:"value"`
		Count   int      `json:"count"`
		Samples []string `json:"samples,omitempty"`
	}
	counts := map[string]*val{}
	add := func(k string, e *Ext) {
		if k == "" {
			return
		}
		v := counts[k]
		if v == nil {
			v = &val{Value: k}
			counts[k] = v
		}
		v.Count++
		if len(v.Samples) < 4 && starOf(e) > 0 {
			v.Samples = append(v.Samples, e.Name)
		}
	}
	for _, e := range snap.Exts { // snapshot is stars-desc, so samples are top-starred
		switch key {
		case "category":
			add(e.Category, e)
		case "tag", "tags":
			seen := map[string]bool{}
			for _, tag := range e.Tags {
				if !seen[strings.ToLower(tag)] {
					add(tag, e)
					seen[strings.ToLower(tag)] = true
				}
			}
		case "package", "pkg":
			add(e.Pkg, e)
		case "license":
			if e.License == "" {
				add("Unknown", e)
			} else {
				add(e.License, e)
			}
		case "lang":
			add(e.Lang, e)
		case "repo":
			if e.Repo != "n/a" {
				add(e.Repo, e)
			}
		case "vendor":
			add(e.Vendor, e)
		case "kernel":
			add(e.Kernel, e)
		case "kind", "type":
			add(e.Kind, e)
		case "lifecycle":
			add(e.Lifecycle, e)
		case "pg":
			for _, v := range e.PG {
				add(strconv.Itoa(v), e)
			}
		case "os":
			if len(snap.OSs) == 0 {
				continue
			}
			seen := map[int]bool{}
			for _, idx := range e.TargetIdx {
				oi := idx % len(snap.OSs)
				if oi >= 0 && oi < len(snap.OSs) && !seen[oi] {
					add(snap.OSs[oi], e)
					seen[oi] = true
				}
			}
		case "distribution":
			if e.Packaged {
				add("packaged", e)
			}
			if !e.Packaged && (e.RepoURL != "" || e.Tarball != "") {
				add("source", e)
			}
			if e.Vendor != "" {
				add("vendor", e)
			}
			if e.Kernel != "" {
				add("kernel", e)
			}
			if e.Contrib {
				add("contrib", e)
			}
		case "capability", "cap":
			for _, capability := range []string{"binary", "library", "create-extension", "preload", "trusted", "relocatable"} {
				if matchesCapability(e, capability) {
					add(capability, e)
				}
			}
		case "build":
			for _, build := range []string{"rpm", "deb", "pgrx", "source", "packaged"} {
				if matchesBuild(e, build) {
					add(build, e)
				}
			}
		case "docs", "doc":
			add(docClass(e), e)
		case "relation", "rel":
			added := false
			for _, relation := range []string{"requires", "required-by", "see-also"} {
				if matchesRelation(e, relation) {
					add(relation, e)
					added = true
				}
			}
			if !added {
				add("no-relations", e)
			}
		case "pgrx":
			add(e.PGRXVer, e)
		case "activity":
			add(activityYear(e), e)
		default:
			writeErr(w, 400, "unknown dimension %q (valid: %s)", key, strings.Join(dimKeys, " "))
			return
		}
	}
	vals := make([]*val, 0, len(counts))
	for _, v := range counts {
		vals = append(vals, v)
	}
	sort.Slice(vals, func(i, j int) bool {
		if vals[i].Count != vals[j].Count {
			return vals[i].Count > vals[j].Count
		}
		return vals[i].Value < vals[j].Value
	})
	w.Header().Set("Cache-Control", "public, max-age=60")
	writeJSON(w, 200, map[string]any{"key": key, "values": vals})
}

/*
bootstrap: compact positional rows for the SPA.

	Column order — keep in sync with decodeBoot() in web/app.js:
	  0 name  1 category  2 avail  3 repo  4 license  5 lang  6 version  7 stars
	  8 en_desc  9 zh_desc  10 kind  11 vendor  12 kernel  13 pg_ver[]
	  14 flags(lead1 contrib2 bin4 lib8 ddl16 load32 trusted64 reloc128)
	  15 docbits(en1 zh2)  16 last_commit  17 pkg  18 lead_ext  19 lifecycle
	  20 tags[]  21 last_active  22 checked_at
	  23 buildbits(rpm1 deb2 pgrx4 source8)  24 target_idx[]  25 family_size  26 comment
	  27 relationbits(requires1 required_by2 see_also4)  28 pgrx_ver  29 repo_url  30 url  31 license_url  32 rpm_pkg  33 deb_pkg
*/
// bootFormat versions the positional payload layout. It participates in the
// ETag (and the SPA's request URL) so a redeployed binary can never pair its
// new app.js with an older cached payload whose columns no longer line up.
const bootFormat = "b34"

func bootstrapETag(snap *Snapshot) string {
	return `"` + bootFormat + "-" + strings.Trim(snap.Version, `"`) + `"`
}

// bits packs booleans LSB-first, so the call site reads as the bit layout.
func bits(bs ...bool) int {
	n := 0
	for i, b := range bs {
		if b {
			n |= 1 << i
		}
	}
	return n
}

func (a *api) handleBootstrap(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	etag := bootstrapETag(snap)
	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	rows := make([]any, len(snap.Exts))
	for i, e := range snap.Exts {
		flags := bits(e.Lead, e.Contrib, e.HasBin, e.HasLib, e.NeedDDL, e.NeedLoad, e.Trusted, e.Relocatable)
		docbits := bits(e.HasEnDoc, e.HasZhDoc)
		buildbits := bits(e.RPMBuild, e.DEBBuild, e.PGRXVer != "", e.RepoURL != "" || e.Tarball != "")
		relationbits := bits(len(e.Requires) > 0, len(e.RequiredBy) > 0, len(e.SeeAlso) > 0)
		rows[i] = []any{
			e.Name, e.Category, bits(e.Packaged), e.Repo, e.License, e.Lang, e.Version, e.Stars,
			e.EnDesc, e.ZhDesc, e.Kind, e.Vendor, e.Kernel, e.PG,
			flags, docbits, e.LastCommit, e.Pkg, e.LeadExt, e.Lifecycle, e.Tags,
			e.LastActive, e.CheckedAt, buildbits, e.TargetIdx, e.FamilySize, e.Comment,
			relationbits, e.PGRXVer, e.RepoURL, e.URL, e.LicenseURL, e.RPMPkg, e.DEBPkg,
		}
	}
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", "public, max-age=60")
	writeJSON(w, 200, map[string]any{
		"generated": snap.LoadedAt.Format(time.RFC3339),
		"format":    bootFormat,
		"version":   strings.Trim(snap.Version, `"`),
		"pg":        snap.PGs,
		"os":        snap.OSs,
		"counts":    snap.counts(),
		"cats":      snap.Cats,
		"rows":      rows,
	})
}

/* visit counter: in-memory counts persisted additively to pgext.counter.
   POST increments, GET reads. The download counter is a placeholder that
   reports the stored value until artifact download tracking lands. */

func (a *api) handleVisit(w http.ResponseWriter, r *http.Request) {
	_, e := a.extOf(w, r)
	if e == nil {
		return
	}
	var visits, downloads int64
	if r.Method == http.MethodPost {
		visits, downloads = a.counter.hit(e.Name)
	} else {
		visits, downloads = a.counter.get(e.Name)
	}
	w.Header().Set("Cache-Control", "no-store")
	writeJSON(w, 200, map[string]any{"ext": e.Name, "visits": visits, "downloads": downloads})
}

// ChangelogRow is one release-notes entry of the package this extension
// belongs to. Versions may be absent: a new package has no old version, an
// unchanged rebuild records neither.
type ChangelogRow struct {
	DS     string `json:"ds"`
	Old    string `json:"old,omitempty"`
	New    string `json:"new,omitempty"`
	NoteEn string `json:"note_en,omitempty"`
	NoteZh string `json:"note_zh,omitempty"`
}

// handleChangelog serves the per-package release history distilled from the
// RPM/DEB release notes (pgext.changelog). A catalog database that predates
// the changelog table simply reports an empty history.
func (a *api) handleChangelog(w http.ResponseWriter, r *http.Request) {
	snap, e := a.extOf(w, r)
	if e == nil {
		return
	}
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		rows, _ := a.pool.Query(ctx, `
			SELECT ds::text, coalesce(old_ver,''), coalesce(new_ver,''), coalesce(note_en,''), coalesce(note_zh,'')
			FROM pgext.changelog WHERE pkg = $1 ORDER BY ds DESC`, e.Pkg)
		entries, err := pgx.CollectRows(rows, pgx.RowToStructByPos[ChangelogRow])
		if err != nil {
			if !isUndefinedTable(err) {
				return nil, err
			}
			entries = []ChangelogRow{} // catalog database predates the changelog table
		}
		return map[string]any{"ext": e.Name, "pkg": e.Pkg, "total": len(entries), "entries": entries}, nil
	})
}

// isUndefinedTable reports whether err is PostgreSQL 42P01 (relation does not
// exist) — the serving database has not loaded the current schema yet.
func isUndefinedTable(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "42P01"
}

func (a *api) handleReload(w http.ResponseWriter, r *http.Request) {
	if a.reloadToken == "" {
		writeErr(w, http.StatusNotFound, "catalog reload is disabled")
		return
	}
	got := strings.TrimSpace(r.Header.Get("X-PGEXT-Reload-Token"))
	if auth := r.Header.Get("Authorization"); got == "" && strings.HasPrefix(auth, "Bearer ") {
		got = strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
	}
	if subtle.ConstantTimeCompare([]byte(got), []byte(a.reloadToken)) != 1 {
		writeErr(w, http.StatusUnauthorized, "invalid reload token")
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	snap, err := a.store.Reload(ctx)
	if err != nil {
		logrus.Errorf("catalog reload failed: %v", err)
		writeErr(w, 500, "catalog reload failed")
		return
	}
	// recompute the matrix caches at write time so no visitor pays for them
	a.refreshMatrixCaches(ctx, true)
	writeJSON(w, 200, map[string]any{
		"generated": snap.LoadedAt.Format(time.RFC3339),
		"total":     len(snap.Exts),
	})
}

func (a *api) handleHealth(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	dbOK := a.pool.Ping(ctx) == nil
	code := 200
	if snap == nil || !dbOK {
		code = 503
	}
	body := map[string]any{"db": dbOK}
	if snap != nil {
		body["extensions"] = len(snap.Exts)
		body["snapshot_age_seconds"] = int(time.Since(snap.LoadedAt).Seconds())
	}
	writeJSON(w, code, body)
}

/* cached wraps a per-extension query with the TTL response cache; entries are
   keyed by snapshot version + URL, so a snapshot reload invalidates everything. */

type errNotFound struct{ msg string }

func (e errNotFound) Error() string { return e.msg }

func (a *api) cached(w http.ResponseWriter, r *http.Request, ver string, fn func(ctx context.Context) (any, error)) {
	key := ver + "|" + r.URL.RequestURI()
	body, ok := a.cache.Get(key)
	if !ok {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		v, err := fn(ctx)
		if err != nil {
			if nf, ok := err.(errNotFound); ok {
				writeErr(w, 404, "%s", nf.msg)
				return
			}
			logrus.Errorf("query failed for %s: %v", r.URL.Path, err)
			writeErr(w, 500, "catalog query failed")
			return
		}
		if body, err = json.Marshal(v); err != nil {
			logrus.Errorf("encode failed for %s: %v", r.URL.Path, err)
			writeErr(w, 500, "response encoding failed")
			return
		}
		a.cache.Set(key, body)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.WriteHeader(200)
	w.Write(body)
}
