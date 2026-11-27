package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLiveQueryCacheExpiresOnIdenticalCatalogReload(t *testing.T) {
	for _, key := range []string{"doc|demo|en", "changelog|demo", "matrix|demo", "files|demo|0|"} {
		t.Run(key, func(t *testing.T) {
			snap := fixtureSnapshot()
			a := &api{cache: newTTLCache(time.Hour, 10)}
			calls := 0
			query := func(context.Context) (any, error) { calls++; return map[string]int{"revision": calls}, nil }
			request := func(extraQuery string) string {
				w := httptest.NewRecorder()
				a.cached(w, httptest.NewRequest(http.MethodGet, "/api/v1/ext/demo/matrix"+extraQuery, nil), snap.queryVersion()+"|"+key, query)
				return w.Body.String()
			}
			before := request("")
			// Canonical key ensures arbitrary query strings cannot bust the cache.
			if request("?bust=12345") != before || calls != 1 {
				t.Fatal("same snapshot did not reuse its canonical query cache")
			}
			snap.LoadedAt = snap.LoadedAt.Add(time.Minute)
			if request("") == before || calls != 2 {
				t.Fatal("live content stayed cached after a metadata-identical reload")
			}
		})
	}
}

func TestBootstrapRevalidatesMatrixCounts(t *testing.T) {
	for _, tt := range []struct {
		name   string
		change func(*Snapshot)
		status int
	}{
		{"same content reloaded", func(s *Snapshot) { s.LoadedAt = s.LoadedAt.Add(time.Minute) }, http.StatusNotModified},
		{"package count only", func(s *Snapshot) { s.CountPackages++ }, http.StatusOK},
		{"slot count only", func(s *Snapshot) { s.CountBuildSlots++ }, http.StatusOK},
	} {
		t.Run(tt.name, func(t *testing.T) {
			snap := fixtureSnapshot()
			snap.CountPackages, snap.CountBuildSlots = 3, 12
			var err error
			snap.Version, err = snap.contentVersion()
			if err != nil {
				t.Fatal(err)
			}
			oldETag := bootstrapETag(snap)
			tt.change(snap)
			snap.Version, err = snap.contentVersion()
			if err != nil {
				t.Fatal(err)
			}
			store := NewStore(nil)
			store.snap.Store(snap)
			a := &api{store: store}
			req := httptest.NewRequest(http.MethodGet, "/api/v1/bootstrap", nil)
			req.Header.Set("If-None-Match", oldETag)
			w := httptest.NewRecorder()
			a.handleBootstrap(w, req)
			if w.Code != tt.status {
				t.Fatalf("bootstrap status = %d, want %d", w.Code, tt.status)
			}
			if tt.status == http.StatusOK && w.Header().Get("ETag") == oldETag {
				t.Fatal("changed counts retained the old bootstrap ETag")
			}
		})
	}
}
