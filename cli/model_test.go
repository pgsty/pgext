package cli

import (
	"database/sql"
	"testing"
)

func TestGetPkgPageName(t *testing.T) {
	tests := []struct {
		name string
		ext  Extension
		want string
	}{
		{
			name: "lead page uses current extension name",
			ext: Extension{
				Name:    "biscuit",
				Pkg:     "pg_biscuit",
				Lead:    true,
				LeadExt: sql.NullString{String: "pg_biscuit", Valid: true},
			},
			want: "biscuit",
		},
		{
			name: "non-lead page uses lead extension page",
			ext: Extension{
				Name:    "postgis_topology",
				Pkg:     "postgis",
				Lead:    false,
				LeadExt: sql.NullString{String: "postgis", Valid: true},
			},
			want: "postgis",
		},
		{
			name: "fallback to current page when lead extension is missing",
			ext: Extension{
				Name: "pgml",
				Pkg:  "pgml",
				Lead: false,
			},
			want: "pgml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ext.GetPkgPageName(); got != tt.want {
				t.Fatalf("GetPkgPageName() = %q, want %q", got, tt.want)
			}
		})
	}
}
