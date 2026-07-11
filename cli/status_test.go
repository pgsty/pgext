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
		name  string
		times map[string]*time.Time
		want  bool
	}{
		{name: "never parsed", times: map[string]*time.Time{}, want: false},
		{name: "never recapped", times: map[string]*time.Time{"parse": &base}, want: true},
		{name: "parse newer", times: map[string]*time.Time{"parse": &base, "recap": &earlier}, want: true},
		{name: "same transaction", times: map[string]*time.Time{"parse": &base, "recap": &base}, want: false},
		{name: "recap newer", times: map[string]*time.Time{"parse": &base, "recap": &later}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := packageCatalogStale(tt.times); got != tt.want {
				t.Fatalf("packageCatalogStale() = %v, want %v", got, tt.want)
			}
		})
	}
}
