package cli

import (
	"database/sql"
	"strings"
	"testing"
)

func TestGlobalMatrixClassifyCellMatchesAvailabilityColors(t *testing.T) {
	tests := []struct {
		name string
		pkg  *PkgInfo
		want string
	}{
		{
			name: "available from pgdg is blue code",
			pkg: &PkgInfo{
				State:   sql.NullString{Valid: true, String: "AVAIL"},
				Org:     sql.NullString{Valid: true, String: "pgdg"},
				Version: sql.NullString{Valid: true, String: "1.0.0"},
			},
			want: "B",
		},
		{
			name: "available from pigsty is green code",
			pkg: &PkgInfo{
				State:   sql.NullString{Valid: true, String: "AVAIL"},
				Org:     sql.NullString{Valid: true, String: "pigsty"},
				Version: sql.NullString{Valid: true, String: "1.0.0"},
			},
			want: "G",
		},
		{
			name: "missing package is red code",
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "MISS"}},
			want: "R",
		},
		{
			name: "not applicable package is gray code",
			pkg:  &PkgInfo{State: sql.NullString{Valid: true, String: "N/A"}},
			want: ".",
		},
		{
			name: "missing row defaults to not applicable",
			pkg:  nil,
			want: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classifyGlobalMatrixCell(tt.pkg)
			if got.Code != tt.want {
				t.Fatalf("Code = %q, want %q; cell = %+v", got.Code, tt.want, got)
			}
		})
	}
}

func TestPackageStatusRenderersUseTriStateColors(t *testing.T) {
	colors := []struct {
		name, state, org, want string
	}{
		{"pgdg available", "AVAIL", "PGDG", "blue"},
		{"pigsty available", "AVAIL", "PIGSTY", "green"},
		{"missing", "MISS", "", "red"},
		{"not applicable", "N/A", "", "gray"},
	}
	for _, tt := range colors {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBadgeColor(tt.state, tt.org); got != tt.want {
				t.Fatalf("getBadgeColor(%q, %q) = %q, want %q", tt.state, tt.org, got, tt.want)
			}
			if got := getOSBadgeColor(tt.state, tt.org); got != tt.want {
				t.Fatalf("getOSBadgeColor(%q, %q) = %q, want %q", tt.state, tt.org, got, tt.want)
			}
		})
	}

	if got := CCPkgStateBadge("N/A", ""); !strings.Contains(got, "ext-badge--na") || !strings.Contains(got, ">N/A<") {
		t.Fatalf("CCPkgStateBadge(N/A) = %q", got)
	}

	ext := &Extension{RpmRepo: sql.NullString{Valid: true, String: "PIGSTY"}}
	pkg := &PkgInfo{State: sql.NullString{Valid: true, String: "N/A"}, Count: sql.NullInt64{Valid: true, Int64: 0}}
	if got := (&IOPageGenerator{}).formatAvailCell(pkg, ext, "el9.x86_64"); got != "N/A PIGSTY - 0" {
		t.Fatalf("IO N/A cell = %q", got)
	}
	if got := (&CCPageGenerator{}).formatAvailCell(pkg, ext, "el9.x86_64"); got != "N/A PIGSTY - 0" {
		t.Fatalf("CC N/A cell = %q", got)
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
				15: {Code: "R"},
			},
		},
	}
	osVersions := []OSVersion{{OS: "el9.x86_64"}, {OS: "u24.x86_64"}}
	pgVersions := []int{14, 15}

	got := formatGlobalMatrixLetterLine(row, osVersions, pgVersions)
	want := "[pg_duckdb](https://pigsty.io/ext/e/pg_duckdb) | BG | .R"
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
				16: {Code: ".", Class: "gm-na", Label: "N/A", Title: "el9.x86_64 PG16 | N/A"},
			},
		},
		PGDGCells: map[string]map[int]GlobalMatrixCell{
			"el9.x86_64": {
				14: {Code: "B", Class: "gm-pgdg", Label: "PGDG", Title: "el9.x86_64 PG14 | PGDG"},
				15: {Code: "R", Class: "gm-missing", Label: "Missing", Title: "el9.x86_64 PG15 | Missing"},
				16: {Code: ".", Class: "gm-na", Label: "N/A", Title: "el9.x86_64 PG16 | N/A"},
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
		Counts:     map[string]int{"B": 1, "G": 1, ".": 1},
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
		`<b>N/A</b><em data-full-count="1" data-pgdg-count="1">1</em>`,
		`<span class="mx-dot gm-na"></span>`,
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
