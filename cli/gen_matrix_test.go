package cli

import (
	"database/sql"
	"strings"
	"testing"
)

func TestGlobalMatrixClassifyCellMatchesAvailabilityColors(t *testing.T) {
	ext := &Extension{
		PgVer:   []string{"18", "17", "16"},
		RpmRepo: sql.NullString{Valid: true, String: "PIGSTY"},
		DebRepo: sql.NullString{Valid: true, String: "PIGSTY"},
	}

	tests := []struct {
		name string
		os   string
		pg   int
		pkg  *PkgInfo
		want string
	}{
		{
			name: "available from pgdg is blue code",
			os:   "el9.x86_64",
			pg:   18,
			pkg: &PkgInfo{
				State:   sql.NullString{Valid: true, String: "AVAIL"},
				Org:     sql.NullString{Valid: true, String: "pgdg"},
				Version: sql.NullString{Valid: true, String: "1.0.0"},
			},
			want: "B",
		},
		{
			name: "available from pigsty is green code",
			os:   "u24.x86_64",
			pg:   18,
			pkg: &PkgInfo{
				State:   sql.NullString{Valid: true, String: "AVAIL"},
				Org:     sql.NullString{Valid: true, String: "pigsty"},
				Version: sql.NullString{Valid: true, String: "1.0.0"},
			},
			want: "G",
		},
		{
			name: "missing supported package is red code",
			os:   "el9.x86_64",
			pg:   17,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "MISS"}},
			want: "R",
		},
		{
			name: "missing unsupported pg is gray code",
			os:   "el9.x86_64",
			pg:   15,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "MISS"}},
			want: ".",
		},
		{
			name: "missing unsupported platform is amber code",
			os:   "el9.x86_64",
			pg:   17,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "MISS"}},
			want: "A",
		},
		{
			name: "fork is purple code",
			os:   "u24.x86_64",
			pg:   18,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "FORK"}},
			want: "P",
		},
		{
			name: "throw is purple code",
			os:   "u24.x86_64",
			pg:   18,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "THROW"}},
			want: "P",
		},
		{
			name: "break is orange code",
			os:   "u24.x86_64",
			pg:   18,
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "BREAK"}},
			want: "O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testExt := *ext
			if tt.want == "A" {
				testExt.RpmRepo = sql.NullString{}
			}

			got := classifyGlobalMatrixCell(&testExt, tt.pkg, tt.os, tt.pg)
			if got.Code != tt.want {
				t.Fatalf("Code = %q, want %q; cell = %+v", got.Code, tt.want, got)
			}
		})
	}
}

func TestGlobalMatrixLetterLineUsesPackageNameAndLeadLink(t *testing.T) {
	row := &GlobalMatrixRow{
		Pkg: "pg_duckdb",
		Ext: "pg_duckdb",
		Cells: map[string]map[int]GlobalMatrixCell{
			"el9.x86_64": {
				14: {Code: "B"},
				15: {Code: "G"},
			},
			"u24.x86_64": {
				14: {Code: "."},
				15: {Code: "P"},
			},
		},
	}
	osVersions := []OSVersion{{OS: "el9.x86_64"}, {OS: "u24.x86_64"}}
	pgVersions := []int{14, 15}

	got := formatGlobalMatrixLetterLine(row, osVersions, pgVersions)
	want := "[pg_duckdb](https://pigsty.io/ext/e/pg_duckdb) | BG | .P"
	if got != want {
		t.Fatalf("letter line = %q, want %q", got, want)
	}
}

func TestGlobalMatrixStandaloneHTMLIsPresentationPage(t *testing.T) {
	row := &GlobalMatrixRow{
		Pkg: "pg_duckdb",
		Ext: "pg_duckdb",
		Cells: map[string]map[int]GlobalMatrixCell{
			"el9.x86_64": {
				14: {Code: "B", Class: "gm-pgdg", Label: "PGDG", Title: "el9.x86_64 PG14 | PGDG"},
				15: {Code: "G", Class: "gm-pigsty", Label: "Pigsty", Title: "el9.x86_64 PG15 | Pigsty"},
				16: {Code: "A", Class: "gm-platform", Label: "No platform repo", Title: "el9.x86_64 PG16 | No platform repo"},
			},
		},
		PGDGCells: map[string]map[int]GlobalMatrixCell{
			"el9.x86_64": {
				14: {Code: "B", Class: "gm-pgdg", Label: "PGDG", Title: "el9.x86_64 PG14 | PGDG"},
				15: {Code: "R", Class: "gm-missing", Label: "Missing", Title: "el9.x86_64 PG15 | Missing"},
				16: {Code: ".", Class: "gm-unavailable", Label: "Unavailable", Title: "el9.x86_64 PG16 | Unavailable"},
			},
		},
	}
	osVersions := []OSVersion{{OS: "el9.x86_64"}}
	pgVersions := []int{14, 15, 16}
	stats := globalMatrixStats{
		Rows:       1,
		OS:         1,
		PG:         3,
		Cells:      3,
		Counts:     map[string]int{"B": 1, "G": 1, "A": 1},
		PGDGCounts: map[string]int{"B": 1, "R": 1, ".": 1},
	}

	got := NewGlobalMatrixGenerator(nil, "", "").renderStandaloneHTML([]*GlobalMatrixRow{row}, osVersions, pgVersions, stats)

	required := []string{
		`<body data-view="pgdg">`,
		`<button class="mx-tab is-active" type="button" role="tab" aria-selected="true" data-mode="pgdg">PGDG</button>`,
		`<button class="mx-tab" type="button" role="tab" aria-selected="false" data-mode="full">FULL</button>`,
		`data-mode="full"`,
		`data-mode="pgdg"`,
		`>FULL</button>`,
		`>PGDG</button>`,
		`data-code="B"`,
		`data-code="G"`,
		`data-pgdg-code="B"`,
		`data-pgdg-count="1"`,
		`href="https://pigsty.io/ext/e/pg_duckdb"`,
		`href="https://pigsty.io/ext/os/el9.x86_64"`,
		`--cell-size: max(7px, calc((100cqw - var(--pkg-col)) / 80))`,
		`Non-core Extension Packages`,
		`Package Cells`,
		`PG Extension Matrix`,
		`<b>PIGSTY</b><em data-full-count="1"`,
		`<b>Unavailable</b><em data-full-count="1" data-pgdg-count="1">1</em>`,
		`<span class="mx-dot gm-unavailable"></span>`,
	}
	for _, needle := range required {
		if !strings.Contains(got, needle) {
			t.Fatalf("standalone HTML missing %q", needle)
		}
	}
	if strings.Contains(got, "gm-table-wrap") {
		t.Fatalf("standalone HTML should not use the fixed Hugo table scroll wrapper")
	}
	if strings.Contains(got, ">PG14</a>") || strings.Contains(got, "PGDG Coverage") || strings.Contains(got, "No platform repo</b>") || strings.Contains(got, "<b>PXD</b>") {
		t.Fatalf("standalone HTML still contains old presentation labels")
	}
	if strings.Index(got, ">16</a>") > strings.Index(got, ">14</a>") {
		t.Fatalf("standalone HTML should render PG headers from newest to oldest")
	}
	if strings.Index(got, `data-mode="pgdg"`) > strings.Index(got, `data-mode="full"`) {
		t.Fatalf("standalone HTML should put PGDG before FULL")
	}
}
