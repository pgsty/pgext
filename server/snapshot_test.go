package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
