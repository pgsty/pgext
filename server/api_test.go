/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func fixtureSnapshot() *Snapshot {
	stars := func(n int) *int { return &n }
	exts := []*Ext{
		{Name: "postgis", Pkg: "postgis", Category: "GIS", Packaged: true, State: "available", Repo: "PGDG", TargetIdx: []int{0, 1, 2, 3},
			License: "GPL-2.0", Lang: "C", PG: []int{18, 17, 16, 15, 14}, Stars: stars(9),
			EnDesc: "spatial types and functions", Kind: "standard", ExtType: "standard", Lifecycle: "active", HasLib: true, NeedDDL: true,
			LastCommit: "2026-05-11", LastActive: "2026-05-11"},
		{Name: "timescaledb", Pkg: "timescaledb", Category: "TIME", Packaged: true, State: "available", Repo: "PIGSTY", TargetIdx: []int{0, 2},
			License: "Timescale", Lang: "C", PG: []int{18, 17, 16, 15}, Stars: stars(22656),
			EnDesc: "time-series database", ZhDesc: "时序数据库扩展插件", Kind: "preload", ExtType: "preload", Lifecycle: "active", HasLib: true, NeedDDL: true, NeedLoad: true,
			RPMBuild: true, DEBBuild: true, HasEnDoc: true, LastCommit: "2026-05-18", LastActive: "2026-05-18"},
		{Name: "vector", Pkg: "pgvector", Category: "RAG", Packaged: true, State: "available", Repo: "PGDG", Tags: []string{"embedding", "hnsw"}, TargetIdx: []int{0, 1, 2, 3},
			License: "PostgreSQL", Lang: "C", PG: []int{18, 17, 16, 15, 14}, Stars: stars(21351),
			EnDesc: "vector data type and ivfflat and hnsw", Kind: "standard", ExtType: "standard", Lifecycle: "active", HasLib: true, NeedDDL: true, Trusted: true,
			PGRXVer: "0.12.0", RepoURL: "https://github.com/pgvector/pgvector", Requires: []string{"btree_gist"}, HasEnDoc: true, HasZhDoc: true,
			LastCommit: "2026-04-01", LastActive: "2026-04-01"},
		{Name: "aws_s3", Pkg: "aws_s3", Category: "ETL", State: "n/a", Repo: "n/a",
			License: "Unknown", Lang: "SQL", Kernel: "aws-rds-pg", Vendor: "AWS", ExtKernel: "aws-rds-pg", ExtVendor: "AWS",
			EnDesc: "S3 import export for RDS", Kind: "puresql", ExtType: "puresql", Lifecycle: "active"},
	}
	snap := &Snapshot{Exts: exts, ByName: map[string]*Ext{}, LoadedAt: time.Now(), PGs: []int{18, 17}, OSs: []string{"el9.x86_64", "d12.x86_64"}, Version: `"fixture"`}
	for _, e := range exts {
		snap.ByName[e.Name] = e
		if e.State == "available" {
			snap.CountAvail++
		}
	}
	return snap
}

func TestParseFilterOperators(t *testing.T) {
	tests := []struct {
		name string
		q    string
		want Filter
	}{
		{"free words", "q=vector+index", Filter{Words: []string{"vector", "index"}}},
		{"cat operator", "q=cat:gis", Filter{Cat: "GIS"}},
		{"mixed", "q=cat:RAG+hnsw", Filter{Cat: "RAG", Words: []string{"hnsw"}}},
		{"pg operator", "q=pg:18", Filter{PG: []int{18}}},
		{"multi pg operator", "q=pg:17,18", Filter{PG: []int{18, 17}}},
		{"is packaged", "q=is:packaged", Filter{Scope: "packaged"}},
		{"vendor", "q=vendor:aws", Filter{Vendor: "aws"}},
		{"kind lifecycle target", "kind=preload&lifecycle=active&os=el9.x86_64", Filter{Kind: "preload", Lifecycle: "active", OS: "el9.x86_64"}},
		{"universe params", "tag=hnsw&pkg=pgvector&capability=trusted&build=pgrx&docs=bilingual&relation=requires&pgrx=0.12.0&active=2026",
			Filter{Tag: "hnsw", Pkg: "pgvector", Capability: "trusted", Build: "pgrx", Docs: "bilingual", Relation: "requires", PGRX: "0.12.0", Active: "2026"}},
		{"universe operators", "q=tag:hnsw+package:pgvector+build:pgrx+doc:bilingual+rel:requires+year:2026",
			Filter{Tag: "hnsw", Pkg: "pgvector", Build: "pgrx", Docs: "bilingual", Relation: "requires", Active: "2026"}},
		{"legacy cloud scope", "scope=cloud", Filter{Scope: "vendor"}},
		{"explicit params", "cat=time&repo=pigsty&pg=17", Filter{Cat: "TIME", Repo: "PIGSTY", PG: []int{17}}},
		{"multi pg params", "pg=14,18,14", Filter{PG: []int{18, 14}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := url.ParseQuery(tt.q)
			if err != nil {
				t.Fatal(err)
			}
			got := ParseFilter(v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFilter(%q) = %+v, want %+v", tt.q, got, tt.want)
			}
		})
	}
}

func TestFilterApply(t *testing.T) {
	snap := fixtureSnapshot()
	tests := []struct {
		name  string
		query string
		sort  string
		want  []string
	}{
		{"all sorted by stars", "", "", []string{"timescaledb", "vector", "postgis", "aws_s3"}},
		{"by name", "", "name", []string{"aws_s3", "postgis", "timescaledb", "vector"}},
		{"by updated", "", "updated", []string{"timescaledb", "postgis", "vector", "aws_s3"}},
		{"category", "cat=GIS", "", []string{"postgis"}},
		{"pg 14 support", "pg=14", "", []string{"vector", "postgis"}},
		{"all selected pg majors", "pg=18,17", "", []string{"timescaledb", "vector", "postgis"}},
		{"all selected including pg14", "pg=18,14", "", []string{"vector", "postgis"}},
		{"packaged only", "scope=packaged", "", []string{"timescaledb", "vector", "postgis"}},
		{"vendor catalog", "scope=cloud", "", []string{"aws_s3"}},
		{"exact name outranks stars", "q=postgis", "", []string{"postgis"}},
		{"desc match", "q=hnsw", "", []string{"vector"}},
		{"tag match", "q=embedding", "", []string{"vector"}},
		{"exact binary target", "pg=18&os=d12.x86_64", "", []string{"vector", "postgis"}},
		{"all selected exact binary targets", "pg=18,17&os=d12.x86_64", "", []string{"vector", "postgis"}},
		{"missing selected binary target", "pg=18,14&os=d12.x86_64", "", []string{}},
		{"zh desc match", "q=时序", "", []string{"timescaledb"}},
		{"operator in q", "q=cat:RAG", "", []string{"vector"}},
		{"exact tag", "tag=hnsw", "", []string{"vector"}},
		{"package family", "pkg=pgvector", "", []string{"vector"}},
		{"runtime capability", "capability=preload", "", []string{"timescaledb"}},
		{"pgrx build", "build=pgrx", "", []string{"vector"}},
		{"bilingual docs", "docs=bilingual", "", []string{"vector"}},
		{"dependency signal", "relation=requires", "", []string{"vector"}},
		{"pgrx version", "pgrx=0.12.0", "", []string{"vector"}},
		{"activity unknown", "active=unknown", "", []string{"aws_s3"}},
		{"no match", "q=nosuchthing", "", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, _ := url.ParseQuery(tt.query)
			got := ParseFilter(v).Apply(snap, tt.sort)
			names := make([]string, len(got))
			for i, e := range got {
				names[i] = e.Name
			}
			if len(names) != len(tt.want) {
				t.Fatalf("Apply(%q) = %v, want %v", tt.query, names, tt.want)
			}
			for i := range names {
				if names[i] != tt.want[i] {
					t.Fatalf("Apply(%q) = %v, want %v", tt.query, names, tt.want)
				}
			}
		})
	}
}

func TestPGTargetAlsoRequiresUniverseCoverage(t *testing.T) {
	snap := &Snapshot{PGs: []int{18, 17}, OSs: []string{"el9.x86_64"}}
	ext := &Ext{Name: "legacy_member", PG: []int{18}, TargetIdx: []int{0, 1}}
	filter := Filter{PG: []int{18, 17}, OS: "el9.x86_64"}
	if _, ok := filter.match(ext, snap); ok {
		t.Fatal("package-family targets must not bypass the extension's pg_ver coverage")
	}
}

func TestBootstrapIncludesNewUniverseFields(t *testing.T) {
	snap := fixtureSnapshot()
	snap.CountPackaged = 3
	snap.CountProjects = 4
	store := NewStore(nil)
	store.snap.Store(snap)
	a := &api{store: store}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/bootstrap", nil)
	w := httptest.NewRecorder()
	a.handleBootstrap(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("bootstrap status = %d", w.Code)
	}
	var body struct {
		Rows   [][]any        `json:"rows"`
		Counts map[string]int `json:"counts"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Rows) != 4 || len(body.Rows[0]) != 29 {
		t.Fatalf("bootstrap shape = %d rows x %d cols", len(body.Rows), len(body.Rows[0]))
	}
	if body.Counts["projects"] != 4 || body.Counts["packaged"] != 3 {
		t.Fatalf("unexpected counts: %#v", body.Counts)
	}
}

func TestUniverseDimensions(t *testing.T) {
	snap := fixtureSnapshot()
	store := NewStore(nil)
	store.snap.Store(snap)
	a := &api{store: store}

	for _, key := range dimKeys {
		t.Run(key, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/dim/"+key, nil)
			req.SetPathValue("key", key)
			w := httptest.NewRecorder()
			a.handleDim(w, req)
			if w.Code != http.StatusOK {
				t.Fatalf("dimension %s status = %d: %s", key, w.Code, w.Body.String())
			}
			var body struct {
				Key    string `json:"key"`
				Values []struct {
					Value string `json:"value"`
					Count int    `json:"count"`
				} `json:"values"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if body.Key != key || len(body.Values) == 0 {
				t.Fatalf("dimension %s returned %#v", key, body)
			}
		})
	}
}

func TestPreloadLibrariesIncludesDependencies(t *testing.T) {
	deps := map[string]*Ext{
		"dep_a": {Name: "dep_a", NeedLoad: true, Libs: []string{"dep_a, shared"}},
		"dep_b": {Name: "dep_b", NeedLoad: false, Libs: []string{"ignored"}},
	}
	ext := &Ext{Name: "demo", NeedLoad: true, Requires: []string{"dep_a", "dep_b"}, Libs: []string{"shared", "demo_lib"}}
	got := preloadLibraries(ext, deps)
	want := []string{"dep_a", "shared", "demo_lib"}
	if len(got) != len(want) {
		t.Fatalf("preloadLibraries = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("preloadLibraries = %v, want %v", got, want)
		}
	}
	if got := preloadLibraries(&Ext{Name: "headless"}, deps); got != nil {
		t.Fatalf("non-preload extension returned %v", got)
	}
}

func TestReloadDisabledAndUnauthorized(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  int
	}{
		{"disabled", "", http.StatusNotFound},
		{"wrong token", "secret", http.StatusUnauthorized},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &api{reloadToken: tt.token}
			req := httptest.NewRequest(http.MethodPost, "/api/v1/reload", strings.NewReader(""))
			w := httptest.NewRecorder()
			a.handleReload(w, req)
			if w.Code != tt.want {
				t.Fatalf("status = %d, want %d", w.Code, tt.want)
			}
		})
	}
}

func TestPublicMiddlewareIsReadOnlyAndHardened(t *testing.T) {
	h := withSecurityHeaders(withCORS(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/ext", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("preflight status = %d", w.Code)
	}
	if methods := w.Header().Get("Access-Control-Allow-Methods"); strings.Contains(methods, "POST") {
		t.Fatalf("public CORS unexpectedly allows POST: %s", methods)
	}
	if w.Header().Get("Content-Security-Policy") == "" || w.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatalf("missing security headers: %#v", w.Header())
	}
}

func TestServeIndexFingerprintsAssets(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	serveIndex(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("index status = %d", w.Code)
	}
	if got := w.Header().Get("Cache-Control"); got != "no-cache" {
		t.Fatalf("index cache control = %q", got)
	}
	etag := w.Header().Get("ETag")
	if etag == "" {
		t.Fatal("index is missing ETag")
	}
	body := w.Body.String()
	for _, asset := range []string{"style.css", "app.js"} {
		version := strings.Trim(assetETags[asset], `"`)
		want := "/assets/" + asset + "?v=" + version
		if !strings.Contains(body, want) {
			t.Fatalf("index is missing fingerprinted asset %q", want)
		}
	}

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", etag)
	w = httptest.NewRecorder()
	serveIndex(w, req)
	if w.Code != http.StatusNotModified {
		t.Fatalf("conditional index status = %d", w.Code)
	}
	if got := w.Header().Get("ETag"); got != etag {
		t.Fatalf("304 ETag = %q, want %q", got, etag)
	}
	if got := w.Header().Get("Cache-Control"); got != "no-cache" {
		t.Fatalf("304 cache control = %q", got)
	}
}

func TestSPAHistoryRoutesServeShell(t *testing.T) {
	for _, path := range []string{"/e/vector", "/p/pgvector", "/c/RAG", "/browse", "/dim/license", "/about"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			w := httptest.NewRecorder()
			handleSPA(w, req)
			if w.Code != http.StatusOK {
				t.Fatalf("GET %s status = %d", path, w.Code)
			}
			if got := w.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
				t.Fatalf("GET %s content type = %q", path, got)
			}
			if got := w.Header().Get("Cache-Control"); got != "no-cache" {
				t.Fatalf("GET %s cache control = %q", path, got)
			}
			if !strings.Contains(w.Body.String(), `id="app"`) {
				t.Fatalf("GET %s did not return the SPA shell", path)
			}
		})
	}
}

func TestEmbeddedDetailManualContract(t *testing.T) {
	app, err := webFS.ReadFile("web/app.js")
	if err != nil {
		t.Fatal(err)
	}
	appText := string(app)
	for _, want := range []string{
		`id="ext-overview"`, `id="ext-metadata"`, `id="ext-relations"`,
		`id="ext-packages"`, `id="ext-build"`, `id="ext-install"`, `id="ext-usage"`,
		`id="pkg-overview"`, `id="pkg-version"`, `id="pkg-availability"`, `id="pkg-downloads"`,
		`id="pkg-build"`, `id="pkg-install"`, `id="pkg-extensions"`,
		"metadataHTML", "packageVersionsHTML", "packageInstallHTML", "preload_libs", "pig build pkg",
		"file-browser", "usage-prose", "function navigateTo", "const extHref", "const pkgHref",
		"DIM_GROUPS", "relationbits", "migrateLegacyHash", "data-dim-search", "capabilityMatches",
	} {
		if !strings.Contains(appText, want) {
			t.Errorf("embedded detail app is missing %q", want)
		}
	}
	if strings.Contains(appText, `href="#/`) {
		t.Error("embedded app still emits hash-prefixed links")
	}

	style, err := webFS.ReadFile("web/style.css")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{".manual-outline", ".metadata-grid", ".install-steps", ".pkg-facts", ".pkg-install-panel", ".doc-layout", ".usage-prose", ".md-table", ".dimension-group", ".dim-toolbar", ".active-filters"} {
		if !strings.Contains(string(style), want) {
			t.Errorf("embedded detail styles are missing %q", want)
		}
	}
}

func TestSortOS(t *testing.T) {
	oss := []string{
		"u24.x86_64", "el8.aarch64", "d12.x86_64", "el10.x86_64", "u22.aarch64",
		"el8.x86_64", "d13.aarch64", "el9.x86_64", "u26.x86_64", "d12.aarch64",
	}
	sortOS(oss)
	want := []string{
		"el8.x86_64", "el8.aarch64", "el9.x86_64", "el10.x86_64",
		"d12.x86_64", "d12.aarch64", "d13.aarch64",
		"u22.aarch64", "u24.x86_64", "u26.x86_64",
	}
	for i := range want {
		if oss[i] != want[i] {
			t.Fatalf("sortOS = %v, want %v", oss, want)
		}
	}
}

func TestAtois(t *testing.T) {
	got := atois([]string{"14", "18", "junk", "16", ""})
	want := []int{18, 16, 14}
	if len(got) != len(want) {
		t.Fatalf("atois = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("atois = %v, want %v", got, want)
		}
	}
}

func TestTTLCache(t *testing.T) {
	c := newTTLCache(50*time.Millisecond, 2)
	c.Set("a", []byte("1"))
	if v, ok := c.Get("a"); !ok || string(v) != "1" {
		t.Fatal("expected cache hit for a")
	}
	c.Set("b", []byte("2"))
	c.Set("c", []byte("3")) // exceeds maxSize, evicts something
	if len(c.items) > 2 {
		t.Fatalf("cache exceeded max size: %d", len(c.items))
	}
	time.Sleep(60 * time.Millisecond)
	if _, ok := c.Get("a"); ok {
		t.Fatal("expected ttl expiry for a")
	}
}
