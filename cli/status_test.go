package cli

import (
	"testing"
	"time"
)

func TestPackageCatalogStale(t *testing.T) {
	base := time.Date(2026, time.July, 11, 12, 0, 0, 0, time.UTC)
	earlier := base.Add(-time.Minute)
	later := base.Add(time.Minute)

	tests := []struct {
		name        string
		times       map[string]*time.Time
		wantCommand string
	}{
		{name: "never fetched", times: map[string]*time.Time{}, wantCommand: ""},
		{name: "fetch invalidated", times: map[string]*time.Time{"parse": &base, "recap": &base}, wantCommand: "pgext fetch"},
		{name: "fetched but never parsed", times: map[string]*time.Time{"fetch": &base}, wantCommand: "pgext parse"},
		{name: "fetch newer", times: map[string]*time.Time{"fetch": &base, "parse": &earlier, "recap": &later}, wantCommand: "pgext parse"},
		{name: "never recapped", times: map[string]*time.Time{"fetch": &earlier, "parse": &base}, wantCommand: "pgext recap"},
		{name: "parse newer", times: map[string]*time.Time{"fetch": &earlier, "parse": &base, "recap": &earlier}, wantCommand: "pgext recap"},
		{name: "same transaction", times: map[string]*time.Time{"fetch": &base, "parse": &base, "recap": &base}, wantCommand: ""},
		{name: "recap newer", times: map[string]*time.Time{"fetch": &earlier, "parse": &base, "recap": &later}, wantCommand: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := packageCatalogRefreshCommand(tt.times); got != tt.wantCommand {
				t.Fatalf("packageCatalogRefreshCommand() = %q, want %q", got, tt.wantCommand)
			}
			if got := packageCatalogStale(tt.times); got != (tt.wantCommand != "") {
				t.Fatalf("packageCatalogStale() = %v, want %v", got, tt.wantCommand != "")
			}
		})
	}
}
