/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

/* ----------------------------------------------------------------
   API surface (all JSON, read-only except /reload):

     GET  /api/v1/meta                  server + snapshot + dimensions
     GET  /api/v1/ext                   list/search: q cat repo license lang pg type vendor scope sort limit offset
     GET  /api/v1/ext/{name}            one extension, full record
     GET  /api/v1/ext/{name}/matrix     package availability: pg × os cells
     GET  /api/v1/ext/{name}/files      binary artifacts with download URLs (?pg= &os=)
     GET  /api/v1/ext/{name}/doc        usage doc markdown (?lang=en|zh)
     GET  /api/v1/dim/{key}             aggregate: category license lang repo vendor kernel type pg
     GET  /api/v1/bootstrap             compact positional payload for the SPA
     POST /api/v1/reload                refresh the snapshot from the database
     GET  /healthz                      liveness + snapshot age
   ---------------------------------------------------------------- */

// Filter mirrors the SPA search semantics: free words plus key:value operators.
type Filter struct {
	Words   []string
	Cat     string
	Repo    string
	License string
	Lang    string
	PG      int
	Type    string
	Vendor  string
	Scope   string // "", "packaged", "cloud"
}

// ParseFilter reads filters from query params; `q` may embed operators
// like `cat:GIS lang:rust pg:18 vendor:aws is:packaged`.
func ParseFilter(v url.Values) Filter {
	f := Filter{
		Cat:     strings.ToUpper(v.Get("cat")),
		Repo:    strings.ToUpper(v.Get("repo")),
		License: v.Get("license"),
		Lang:    v.Get("lang"),
		Type:    strings.ToLower(v.Get("type")),
		Vendor:  v.Get("vendor"),
		Scope:   strings.ToLower(v.Get("scope")),
	}
	if n, err := strconv.Atoi(v.Get("pg")); err == nil {
		f.PG = n
	}
	for _, tok := range strings.Fields(v.Get("q")) {
		key, val, ok := strings.Cut(tok, ":")
		if !ok || val == "" {
			f.Words = append(f.Words, strings.ToLower(tok))
			continue
		}
		switch strings.ToLower(key) {
		case "cat", "category":
			f.Cat = strings.ToUpper(val)
		case "repo":
			f.Repo = strings.ToUpper(val)
		case "license":
			f.License = val
		case "lang", "lng":
			f.Lang = val
		case "pg":
			if n, err := strconv.Atoi(val); err == nil {
				f.PG = n
			}
		case "type":
			f.Type = strings.ToLower(val)
		case "vendor":
			f.Vendor = val
		case "is":
			switch strings.ToLower(val) {
			case "packaged":
				f.Scope = "packaged"
			case "cloud":
				f.Scope = "cloud"
			case "contrib":
				f.Repo = "CONTRIB"
			}
		default:
			f.Words = append(f.Words, strings.ToLower(tok))
		}
	}
	return f
}

func (f Filter) match(e *Ext) (int, bool) {
	if f.Cat != "" && e.Category != f.Cat {
		return 0, false
	}
	if f.Repo != "" && !strings.EqualFold(e.Repo, f.Repo) {
		return 0, false
	}
	if f.License != "" && !strings.EqualFold(e.License, f.License) {
		return 0, false
	}
	if f.Lang != "" && !strings.EqualFold(e.Lang, f.Lang) {
		return 0, false
	}
	if f.PG != 0 && !containsInt(e.PG, f.PG) {
		return 0, false
	}
	if f.Type != "" && e.ExtType != f.Type {
		return 0, false
	}
	if f.Scope == "packaged" && e.State != "available" {
		return 0, false
	}
	if f.Scope == "cloud" && e.ExtKernel == "" && e.ExtVendor == "" {
		return 0, false
	}
	if f.Vendor != "" {
		v := strings.ToLower(f.Vendor)
		if !strings.Contains(strings.ToLower(e.ExtVendor), v) && !strings.Contains(strings.ToLower(e.ExtKernel), v) {
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
		case strings.Contains(strings.ToLower(e.EnDesc), w) || strings.Contains(e.ZhDesc, w):
			score += 10
		case strings.EqualFold(e.Category, w),
			strings.Contains(strings.ToLower(e.ExtVendor), w),
			strings.Contains(strings.ToLower(e.ExtKernel), w):
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
		if score, ok := f.match(e); ok {
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
			if a.LastCommit != b.LastCommit {
				return a.LastCommit > b.LastCommit
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

func containsInt(xs []int, x int) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

/* ---------------- handlers ---------------- */

type api struct {
	store *Store
	cache *ttlCache
	pool  *pgxpool.Pool
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

func (a *api) handleMeta(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	writeJSON(w, 200, map[string]any{
		"server":    "pgext",
		"generated": snap.LoadedAt.Format(time.RFC3339),
		"counts": map[string]int{
			"total":    len(snap.Exts),
			"packaged": snap.CountAvail,
			"vendor":   snap.CountVendor,
		},
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
	snap := a.store.Get()
	e, ok := snap.ByName[r.PathValue("name")]
	if !ok {
		writeErr(w, 404, "extension %q does not exist", r.PathValue("name"))
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
	snap := a.store.Get()
	e, ok := snap.ByName[r.PathValue("name")]
	if !ok {
		writeErr(w, 404, "extension %q does not exist", r.PathValue("name"))
		return
	}
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		rows, err := a.pool.Query(ctx, `
			SELECT pg, os, coalesce(name,''), state::text, coalesce(org,''), coalesce(version,''), coalesce(count,0)
			FROM pgext.pkg WHERE pkg = $1 ORDER BY pg DESC, os`, e.Pkg)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		cells := []MatrixCell{}
		for rows.Next() {
			var c MatrixCell
			var cnt int64
			if err := rows.Scan(&c.PG, &c.OS, &c.Name, &c.State, &c.Org, &c.Version, &cnt); err != nil {
				return nil, err
			}
			c.Count = int(cnt)
			cells = append(cells, c)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return map[string]any{
			"ext": e.Name, "pkg": e.Pkg,
			"pg": snap.PGs, "os": snap.OSs,
			"cells": cells,
		}, nil
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
	Size    int    `json:"size"`
	URL     string `json:"url"`
}

func (a *api) handleFiles(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	e, ok := snap.ByName[r.PathValue("name")]
	if !ok {
		writeErr(w, 404, "extension %q does not exist", r.PathValue("name"))
		return
	}
	q := r.URL.Query()
	pg, _ := strconv.Atoi(q.Get("pg"))
	osFilter := q.Get("os")
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		rows, err := a.pool.Query(ctx, `
			SELECT b.pg, b.os, b.name, b.repo, coalesce(r.org,''), coalesce(b.ver,''), coalesce(b.version,''),
			       coalesce(b.file,''), coalesce(b.size,0),
			       CASE WHEN b.href ~ '^https?://' THEN b.href
			            ELSE coalesce(r.default_url,'') || '/' || coalesce(b.href,'') END
			FROM pgext.bin b JOIN pgext.repository r ON b.repo = r.id
			WHERE b.name IN (SELECT DISTINCT name FROM pgext.pkg WHERE pkg = $1 AND name IS NOT NULL)
			  AND ($2 = 0 OR b.pg = $2) AND ($3 = '' OR b.os = $3)
			ORDER BY b.pg DESC, b.os, b.name, b.ver DESC
			LIMIT 5000`, e.Pkg, pg, osFilter)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		files := []FileRow{}
		for rows.Next() {
			var f FileRow
			if err := rows.Scan(&f.PG, &f.OS, &f.Name, &f.Repo, &f.Org, &f.Ver, &f.Version, &f.File, &f.Size, &f.URL); err != nil {
				return nil, err
			}
			files = append(files, f)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return map[string]any{"ext": e.Name, "pkg": e.Pkg, "total": len(files), "files": files}, nil
	})
}

func (a *api) handleDoc(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	e, ok := snap.ByName[r.PathValue("name")]
	if !ok {
		writeErr(w, 404, "extension %q does not exist", r.PathValue("name"))
		return
	}
	lang := "en"
	if r.URL.Query().Get("lang") == "zh" {
		lang = "zh"
	}
	col := "en_doc"
	if lang == "zh" {
		col = "zh_doc"
	}
	a.cached(w, r, snap.Version, func(ctx context.Context) (any, error) {
		var doc *string
		row, err := a.pool.Query(ctx, `SELECT `+col+` FROM pgext.doc WHERE ext = $1`, e.Name)
		if err != nil {
			return nil, err
		}
		defer row.Close()
		if row.Next() {
			if err := row.Scan(&doc); err != nil {
				return nil, err
			}
		}
		if err := row.Err(); err != nil {
			return nil, err
		}
		if doc == nil || *doc == "" {
			return nil, errNotFound{fmt.Sprintf("no %s doc for %q", lang, e.Name)}
		}
		return map[string]any{"ext": e.Name, "lang": lang, "content": *doc}, nil
	})
}

var dimKeys = []string{"category", "license", "lang", "repo", "vendor", "kernel", "type", "pg"}

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
		case "license":
			add(e.License, e)
		case "lang":
			add(e.Lang, e)
		case "repo":
			if e.Repo != "n/a" {
				add(e.Repo, e)
			}
		case "vendor":
			add(e.ExtVendor, e)
		case "kernel":
			add(e.ExtKernel, e)
		case "type":
			add(e.ExtType, e)
		case "pg":
			for _, v := range e.PG {
				add(strconv.Itoa(v), e)
			}
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

	Column order — keep in sync with decode() in web/app.js:
	  0 name  1 category  2 avail  3 repo  4 license  5 lang  6 version  7 stars
	  8 en_desc  9 zh_desc  10 ext_type  11 vendor  12 kernel  13 pg_ver[]
	  14 flags(lead1 contrib2 bin4 lib8 ddl16 load32 trusted64 reloc128)
	  15 docbits(en1 zh2)  16 last_commit
*/
func (a *api) handleBootstrap(w http.ResponseWriter, r *http.Request) {
	snap := a.store.Get()
	if match := r.Header.Get("If-None-Match"); match != "" && match == snap.Version {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	rows := make([]any, len(snap.Exts))
	for i, e := range snap.Exts {
		flags := 0
		for bit, on := range []bool{e.Lead, e.Contrib, e.HasBin, e.HasLib, e.NeedDDL, e.NeedLoad, e.Trusted, e.Relocatable} {
			if on {
				flags |= 1 << bit
			}
		}
		docbits := 0
		if e.HasEnDoc {
			docbits |= 1
		}
		if e.HasZhDoc {
			docbits |= 2
		}
		avail := 0
		if e.State == "available" {
			avail = 1
		}
		rows[i] = []any{
			e.Name, e.Category, avail, e.Repo, e.License, e.Lang, e.Version, e.Stars,
			e.EnDesc, e.ZhDesc, e.ExtType, e.ExtVendor, e.ExtKernel, e.PG,
			flags, docbits, e.LastCommit,
		}
	}
	w.Header().Set("ETag", snap.Version)
	w.Header().Set("Cache-Control", "public, max-age=60")
	writeJSON(w, 200, map[string]any{
		"generated": snap.LoadedAt.Format(time.RFC3339),
		"version":   strings.Trim(snap.Version, `"`),
		"pg":        snap.PGs,
		"os":        snap.OSs,
		"counts": map[string]int{
			"total":    len(snap.Exts),
			"packaged": snap.CountAvail,
			"vendor":   snap.CountVendor,
		},
		"cats": snap.Cats,
		"rows": rows,
	})
}

func (a *api) handleReload(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	snap, err := a.store.Reload(ctx)
	if err != nil {
		writeErr(w, 500, "reload failed: %v", err)
		return
	}
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
	if body, ok := a.cache.Get(key); ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=60")
		w.WriteHeader(200)
		w.Write(body)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	v, err := fn(ctx)
	if err != nil {
		if nf, ok := err.(errNotFound); ok {
			writeErr(w, 404, "%s", nf.msg)
			return
		}
		logrus.Errorf("query failed for %s: %v", r.URL.Path, err)
		writeErr(w, 500, "query failed: %v", err)
		return
	}
	body, err := json.Marshal(v)
	if err != nil {
		writeErr(w, 500, "encode failed: %v", err)
		return
	}
	a.cache.Set(key, body)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.WriteHeader(200)
	w.Write(body)
}
