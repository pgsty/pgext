/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// testMux builds the full route table without a database; redirect handlers
// and the SPA shell never touch the api receiver.
func testMux() *http.ServeMux { return newMux(&api{}) }

func get(t *testing.T, mux *http.ServeMux, target string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, target, nil))
	return rec
}

func TestLegacyRedirects(t *testing.T) {
	mux := testMux()
	cases := []struct {
		path     string
		code     int
		location string
	}{
		// language mirrors: strip /zh (and /en), keep path + query
		{"/zh", http.StatusFound, "/"},
		{"/zh/", http.StatusFound, "/"},
		{"/zh/e/postgis", http.StatusFound, "/e/postgis"},
		{"/zh/ext/postgis?lang=zh", http.StatusFound, "/ext/postgis?lang=zh"},
		{"/zh/list/cate", http.StatusFound, "/list/cate"},
		{"/zh/os/matrix/", http.StatusFound, "/os/matrix/"},
		{"/en", http.StatusFound, "/"},
		{"/en/", http.StatusFound, "/"},
		{"/en/e/postgis", http.StatusFound, "/e/postgis"},

		// short deep-link forms stay permanent, now also with trailing slash
		{"/e/postgis", http.StatusMovedPermanently, "/ext/postgis"},
		{"/e/postgis/", http.StatusMovedPermanently, "/ext/postgis"},
		{"/e/postgis?pg=18", http.StatusMovedPermanently, "/ext/postgis?pg=18"},
		{"/p/timescaledb/", http.StatusMovedPermanently, "/pkg/timescaledb"},
		{"/c/TIME/", http.StatusMovedPermanently, "/cate/TIME"},

		// Hugo-era section pages
		{"/e", http.StatusFound, "/"},
		{"/e/", http.StatusFound, "/"},
		{"/list/ext", http.StatusFound, "/"},
		{"/list/ext/", http.StatusFound, "/"},
		{"/list/pkg", http.StatusFound, "/list/package"},
		{"/list/pkg/", http.StatusFound, "/list/package"},
		{"/list/cate", http.StatusFound, "/list/category"},
		{"/list/cate?x=1", http.StatusFound, "/list/category?x=1"},
		{"/os", http.StatusFound, "/list/os"},
		{"/os/", http.StatusFound, "/list/os"},
		{"/os/matrix", http.StatusFound, "/matrix"},
		{"/os/matrix/", http.StatusFound, "/matrix"},
		{"/categories", http.StatusFound, "/list/category"},
		{"/categories/", http.StatusFound, "/list/category"},
		{"/categories/time", http.StatusFound, "/cate/time"},
		{"/categories/time/", http.StatusFound, "/cate/time"},
		{"/tags", http.StatusFound, "/list/tag"},
		{"/tags/", http.StatusFound, "/list/tag"},
		{"/tags/vector", http.StatusFound, "/?tag=vector"},
		{"/tags/vector/", http.StatusFound, "/?tag=vector"},
		{"/tags/vector?lang=zh", http.StatusFound, "/?tag=vector&lang=zh"},
		{"/repo", http.StatusFound, "/list/repo"},
		{"/repo/", http.StatusFound, "/list/repo"},
		{"/repo/pgsql", http.StatusFound, "/repo/PIGSTY"},
		{"/repo/pgsql/", http.StatusFound, "/repo/PIGSTY"},

		// retired pig CLI handbook moved to the Pigsty docs
		{"/pig", http.StatusFound, pigDocsURL},
		{"/pig/", http.StatusFound, pigDocsURL},
		{"/pig/install", http.StatusFound, pigDocsURL},
		{"/pig/cmd/repo/", http.StatusFound, pigDocsURL},
	}
	for _, tc := range cases {
		rec := get(t, mux, tc.path)
		if rec.Code != tc.code {
			t.Errorf("GET %s: status = %d, want %d", tc.path, rec.Code, tc.code)
			continue
		}
		if loc := rec.Header().Get("Location"); loc != tc.location {
			t.Errorf("GET %s: Location = %q, want %q", tc.path, loc, tc.location)
		}
	}
}

// Legacy chains (e.g. /zh/list/cate → /list/cate → /list/category) must reach
// a 200 in a handful of hops and never loop.
func TestLegacyRedirectChainsTerminate(t *testing.T) {
	mux := testMux()
	for _, start := range []string{
		"/zh/zh/e/postgis", "/zh/list/cate/", "/zh/e/", "/zh/os/matrix",
		"/en/tags/vector/", "/zh/categories/time/",
	} {
		path, seen := start, map[string]bool{}
		for hop := 0; ; hop++ {
			if hop > 5 {
				t.Fatalf("%s: more than 5 redirect hops, stuck at %s", start, path)
			}
			if seen[path] {
				t.Fatalf("%s: redirect loop via %s", start, path)
			}
			seen[path] = true
			rec := get(t, mux, path)
			if rec.Code == http.StatusOK {
				break
			}
			if rec.Code < 300 || rec.Code > 399 {
				t.Fatalf("%s: unexpected status %d at %s", start, rec.Code, path)
			}
			loc := rec.Header().Get("Location")
			if loc == "" || !strings.HasPrefix(loc, "/") || strings.HasPrefix(loc, "//") {
				t.Fatalf("%s: non-local redirect %q at %s", start, loc, path)
			}
			path = loc
		}
	}
}

// Live SPA and API routes must not be shadowed by the compatibility table.
func TestLegacyRedirectsDoNotShadowLiveRoutes(t *testing.T) {
	mux := testMux()
	for _, path := range []string{
		"/", "/ext/postgis", "/pkg/timescaledb", "/cate/TIME", "/matrix",
		"/list", "/list/lang", "/list/license", "/list/category", "/about",
		"/repo/pgdg", "/repo/PGDG", "/repo/PIGSTY", "/os/el9.x86_64", "/pg/18",
	} {
		rec := get(t, mux, path)
		if rec.Code != http.StatusOK {
			t.Errorf("GET %s: status = %d, want 200 (SPA shell)", path, rec.Code)
			continue
		}
		if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
			t.Errorf("GET %s: Content-Type = %q, want text/html shell", path, ct)
		}
	}
}
