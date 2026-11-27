/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed assets/matrix_standalone.css
var matrixStandaloneCSS string

//go:embed assets/matrix_standalone.js
var matrixStandaloneJS string

const globalMatrixBaseURL = "https://pigsty.io/ext/e"
const globalMatrixOSBaseURL = "https://pigsty.io/ext/os"

// GlobalMatrixGenerator creates the cross-OS, cross-PG package matrix.
type GlobalMatrixGenerator struct {
	Cache      *ExtensionCache
	ContentDir string
	StaticDir  string
}

// GlobalMatrixRow is one package row in the global matrix.
type GlobalMatrixRow struct {
	ID        int
	Pkg       string
	Ext       string
	Cells     map[string]map[int]GlobalMatrixCell
	PGDGCells map[string]map[int]GlobalMatrixCell
}

// GlobalMatrixCell is the compact status rendered for one OS/PG slot.
type GlobalMatrixCell struct {
	Code    string `json:"code"`
	Class   string `json:"class"`
	Label   string `json:"label"`
	State   string `json:"state"`
	Org     string `json:"org,omitempty"`
	Version string `json:"version,omitempty"`
	Count   int64  `json:"count,omitempty"`
	Title   string `json:"title"`
}

type globalMatrixStats struct {
	Rows       int
	OS         int
	PG         int
	Cells      int
	Counts     map[string]int
	PGDGCounts map[string]int `json:"pgdg_counts,omitempty"`
}

type globalMatrixLegendItem struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	Class string `json:"class"`
	Color string `json:"color"`
}

var globalMatrixLegend = []globalMatrixLegendItem{
	{Code: "B", Label: "PGDG", Class: "gm-pgdg", Color: "#5b9cd5"},
	{Code: "G", Label: "PIGSTY", Class: "gm-pigsty", Color: "#60be59"},
	{Code: "R", Label: "Missing", Class: "gm-missing", Color: "#cc4637"},
	{Code: ".", Label: "N/A", Class: "gm-na", Color: "#6c6c6c"},
}

// NewGlobalMatrixGenerator creates a global matrix generator.
func NewGlobalMatrixGenerator(cache *ExtensionCache, contentDir, staticDir string) *GlobalMatrixGenerator {
	return &GlobalMatrixGenerator{
		Cache:      cache,
		ContentDir: contentDir,
		StaticDir:  staticDir,
	}
}

// Generate writes the content pages and reusable data files.
func (g *GlobalMatrixGenerator) Generate(ctx context.Context) error {
	rows, osVersions, pgVersions, stats, err := g.buildMatrix(ctx)
	if err != nil {
		return err
	}

	osDir := filepath.Join(g.ContentDir, "os")
	if err := WriteMarkdownFile(filepath.Join(osDir, "matrix.md"), g.renderPage("en")); err != nil {
		return fmt.Errorf("write English matrix page: %w", err)
	}
	if err := WriteMarkdownFile(filepath.Join(osDir, "matrix.zh.md"), g.renderPage("zh")); err != nil {
		return fmt.Errorf("write Chinese matrix page: %w", err)
	}

	dataDir := filepath.Join(g.StaticDir, "matrix")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("create matrix data directory: %w", err)
	}
	if err := g.writeCSV(filepath.Join(dataDir, "pgext-global-matrix.csv"), rows, osVersions, pgVersions); err != nil {
		return err
	}
	if err := g.writeJSON(filepath.Join(dataDir, "pgext-global-matrix.json"), rows, osVersions, pgVersions, stats); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dataDir, "index.html"), []byte(g.renderStandaloneHTML(rows, osVersions, pgVersions, stats)), 0644); err != nil {
		return fmt.Errorf("write matrix presentation page: %w", err)
	}

	return nil
}

func (g *GlobalMatrixGenerator) buildMatrix(ctx context.Context) ([]*GlobalMatrixRow, []OSVersion, []int, globalMatrixStats, error) {
	rows, err := g.loadGlobalMatrixRows(ctx)
	if err != nil {
		return nil, nil, nil, globalMatrixStats{}, err
	}

	pkgData, err := g.loadGlobalMatrixPackageData(ctx)
	if err != nil {
		return nil, nil, nil, globalMatrixStats{}, err
	}
	pgdgOnlyData, err := g.loadGlobalMatrixPGDGOnlyPackageData(ctx)
	if err != nil {
		return nil, nil, nil, globalMatrixStats{}, err
	}

	osVersions := append([]OSVersion(nil), g.Cache.OSVersions...)
	pgVersions := ascendingPGVersions(g.Cache.PGVersions)
	stats := globalMatrixStats{
		Rows:       len(rows),
		OS:         len(osVersions),
		PG:         len(pgVersions),
		Cells:      len(rows) * len(osVersions) * len(pgVersions),
		Counts:     make(map[string]int),
		PGDGCounts: make(map[string]int),
	}

	for _, row := range rows {
		row.Cells = make(map[string]map[int]GlobalMatrixCell, len(osVersions))
		row.PGDGCells = make(map[string]map[int]GlobalMatrixCell, len(osVersions))
		for _, osv := range osVersions {
			row.Cells[osv.OS] = make(map[int]GlobalMatrixCell, len(pgVersions))
			row.PGDGCells[osv.OS] = make(map[int]GlobalMatrixCell, len(pgVersions))
			for _, pg := range pgVersions {
				var pkg *PkgInfo
				if byOS, ok := pkgData[row.Pkg]; ok {
					if byPG, ok := byOS[osv.OS]; ok {
						pkg = byPG[pg]
					}
				}
				cell := classifyGlobalMatrixCell(pkg)
				cell.Title = globalMatrixCellTitle(osv.OS, pg, cell)
				row.Cells[osv.OS][pg] = cell
				stats.Counts[cell.Code]++

				var pgdgPkg *PkgInfo
				if byOS, ok := pgdgOnlyData[row.Pkg]; ok {
					if byPG, ok := byOS[osv.OS]; ok {
						pgdgPkg = byPG[pg]
					}
				}
				pgdgCell := classifyGlobalMatrixCell(pgdgPkg)
				pgdgCell.Title = globalMatrixCellTitle(osv.OS, pg, pgdgCell)
				row.PGDGCells[osv.OS][pg] = pgdgCell
				stats.PGDGCounts[pgdgCell.Code]++
			}
		}
	}

	return rows, osVersions, pgVersions, stats, nil
}

func (g *GlobalMatrixGenerator) loadGlobalMatrixRows(ctx context.Context) ([]*GlobalMatrixRow, error) {
	query := `
		WITH lead_packages AS (
			SELECT DISTINCT ON (e.pkg) e.id, e.pkg, e.name AS ext
			FROM pgext.extension e
			WHERE e.lead = true
			  AND (e.state IS NULL OR e.state <> 'not-ready')
			  AND EXISTS (SELECT 1 FROM pgext.pkg p WHERE p.pkg = e.pkg)
			ORDER BY e.pkg, e.id
		)
		SELECT id, pkg, ext
		FROM lead_packages
		ORDER BY id
	`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("load matrix rows: %w", err)
	}
	defer rows.Close()

	result := make([]*GlobalMatrixRow, 0, 384)
	for rows.Next() {
		row := &GlobalMatrixRow{}
		if err := rows.Scan(&row.ID, &row.Pkg, &row.Ext); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func (g *GlobalMatrixGenerator) loadGlobalMatrixPackageData(ctx context.Context) (map[string]map[string]map[int]*PkgInfo, error) {
	query := `
		WITH lead_packages AS (
			SELECT DISTINCT e.pkg
			FROM pgext.extension e
			WHERE e.lead = true
			  AND (e.state IS NULL OR e.state <> 'not-ready')
		)
		SELECT p.pg, p.os, p.name, p.pkg, p.ext, p.state, p.org, p.version, p.count
		FROM pgext.pkg p
		JOIN lead_packages lp ON lp.pkg = p.pkg
		JOIN pgext.active_pg ap ON ap.pg = p.pg
		JOIN pgext.os o ON o.os = p.os AND o.active
		ORDER BY p.pkg, p.os, p.pg
	`
	return scanGlobalMatrixPkgRows(ctx, query, "load package matrix data")
}

func (g *GlobalMatrixGenerator) loadGlobalMatrixPGDGOnlyPackageData(ctx context.Context) (map[string]map[string]map[int]*PkgInfo, error) {
	query := `
		WITH lead_packages AS (
			SELECT DISTINCT e.pkg
			FROM pgext.extension e
			WHERE e.lead = true
			  AND (e.state IS NULL OR e.state <> 'not-ready')
		),
		base AS (
			SELECT p.pg, p.os, p.name, p.pkg, p.ext, p.state::text AS state, o.os_type
			FROM pgext.pkg p
			JOIN lead_packages lp ON lp.pkg = p.pkg
			JOIN pgext.active_pg ap ON ap.pg = p.pg
			JOIN pgext.os o ON o.os = p.os AND o.active
		),
		candidates AS (
			SELECT pg, os, pkg, ext, name, 1 AS pri FROM base
			UNION ALL
			SELECT pg, os, pkg, ext, regexp_replace(name, '-tsl$', ''), 2
			FROM base
			WHERE os_type = 'deb' AND name LIKE '%-tsl'
			UNION ALL
			SELECT pg, os, pkg, ext, format('postgresql-%s-%s', pg, replace(pkg, '_', '-')), 3
			FROM base
			WHERE os_type = 'deb'
			UNION ALL
			SELECT pg, os, pkg, ext, format('postgresql-%s-%s', pg, regexp_replace(replace(pkg, '_', '-'), '^pg-', '')), 4
			FROM base
			WHERE os_type = 'deb'
			UNION ALL
			SELECT pg, os, pkg, ext, format('%s_%s', replace(pkg, '-', '_'), pg), 3
			FROM base
			WHERE os_type = 'rpm'
			UNION ALL
			SELECT pg, os, pkg, ext, format('%s_%s', regexp_replace(replace(pkg, '-', '_'), '^pg_', ''), pg), 4
			FROM base
			WHERE os_type = 'rpm'
		),
		pgdg_bin AS (
			SELECT b.pg, b.os, b.name, b.version, b.ver,
			       COUNT(*) OVER (PARTITION BY b.pg, b.os, b.name) AS count
			FROM pgext.bin b
			JOIN pgext.repository r ON r.id = b.repo
			WHERE r.org = 'pgdg'
		),
		pgdg_best AS (
			SELECT DISTINCT ON (pg, os, name)
			       pg, os, name, 'PGDG' AS org, version, ver, count
			FROM pgdg_bin
			ORDER BY pg, os, name, ver::pgext.version USING OPERATOR (pgext.>)
		),
		matches AS (
			SELECT DISTINCT ON (c.pg, c.os, c.pkg)
			       c.pg, c.os, c.pkg, c.name, c.pri, p.org, p.version, p.ver, p.count
			FROM candidates c
			JOIN pgdg_best p ON p.pg = c.pg AND p.os = c.os AND p.name = c.name
			ORDER BY c.pg, c.os, c.pkg, c.pri, p.ver::pgext.version USING OPERATOR (pgext.>)
		)
		SELECT b.pg, b.os, COALESCE(m.name, b.name) AS name, b.pkg, b.ext,
		       CASE
		         WHEN COALESCE(m.count, 0) > 0 THEN 'AVAIL'
		         WHEN b.state = 'N/A' THEN 'N/A'
		         ELSE 'MISS'
		       END AS state,
		       CASE WHEN COALESCE(m.count, 0) > 0 THEN m.org ELSE NULL END AS org,
		       CASE WHEN COALESCE(m.count, 0) > 0 THEN m.version ELSE NULL END AS version,
		       COALESCE(m.count, 0) AS count
		FROM base b
		LEFT JOIN matches m ON m.pg = b.pg AND m.os = b.os AND m.pkg = b.pkg
		ORDER BY b.pkg, b.os, b.pg
	`
	return scanGlobalMatrixPkgRows(ctx, query, "load PGDG-only package matrix data")
}

func scanGlobalMatrixPkgRows(ctx context.Context, query, label string) (map[string]map[string]map[int]*PkgInfo, error) {
	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}
	defer rows.Close()

	result := make(map[string]map[string]map[int]*PkgInfo)
	for rows.Next() {
		info := &PkgInfo{}
		if err := rows.Scan(&info.PG, &info.OS, &info.Name, &info.Pkg, &info.Ext, &info.State, &info.Org, &info.Version, &info.Count); err != nil {
			return nil, err
		}
		if result[info.Pkg] == nil {
			result[info.Pkg] = make(map[string]map[int]*PkgInfo)
		}
		if result[info.Pkg][info.OS] == nil {
			result[info.Pkg][info.OS] = make(map[int]*PkgInfo)
		}
		result[info.Pkg][info.OS][info.PG] = info
	}
	return result, rows.Err()
}

func classifyGlobalMatrixCell(pkg *PkgInfo) GlobalMatrixCell {
	if pkg == nil {
		return GlobalMatrixCell{Code: ".", Class: "gm-na", Label: "N/A", State: "N/A"}
	}

	state := "N/A"
	if pkg.State.Valid && pkg.State.String != "" {
		state = strings.ToUpper(pkg.State.String)
	}

	org := ""
	if pkg.Org.Valid {
		org = strings.ToUpper(pkg.Org.String)
	}

	version := ""
	if pkg.Version.Valid {
		version = pkg.Version.String
	}

	count := int64(0)
	if pkg.Count.Valid {
		count = pkg.Count.Int64
	}

	cell := GlobalMatrixCell{
		State:   state,
		Org:     org,
		Version: version,
		Count:   count,
	}

	switch state {
	case "AVAIL":
		if org == "PGDG" {
			cell.Code, cell.Class, cell.Label = "B", "gm-pgdg", "PGDG"
		} else {
			cell.Code, cell.Class, cell.Label = "G", "gm-pigsty", "Pigsty"
		}
	case "MISS":
		cell.Code, cell.Class, cell.Label = "R", "gm-missing", "Missing"
	case "N/A":
		cell.Code, cell.Class, cell.Label = ".", "gm-na", "N/A"
	default:
		cell.Code, cell.Class, cell.Label, cell.State = ".", "gm-na", "N/A", "N/A"
	}

	return cell
}

func globalMatrixCellTitle(osName string, pg int, cell GlobalMatrixCell) string {
	parts := []string{fmt.Sprintf("%s PG%d", osName, pg), cell.Label}
	if cell.State != "" {
		parts = append(parts, "state="+cell.State)
	}
	if cell.Org != "" {
		parts = append(parts, "repo="+cell.Org)
	}
	if cell.Version != "" {
		parts = append(parts, "version="+cell.Version)
	}
	if cell.Count > 0 {
		parts = append(parts, fmt.Sprintf("packages=%d", cell.Count))
	}
	return strings.Join(parts, " | ")
}

func ascendingPGVersions(values []int) []int {
	result := append([]int(nil), values...)
	sort.Ints(result)
	return result
}

// renderPage writes the global matrix page, which is front matter only: the
// grid itself is drawn by the site's `matrix` layout from the JSON export
// written alongside it. That keeps a page which would otherwise be eight
// megabytes of table markup under a kilobyte, and keeps one copy of the rows
// instead of three.
func (g *GlobalMatrixGenerator) renderPage(locale string) string {
	title := "Global Matrix"
	description := "Global PostgreSQL extension package availability matrix across OS and PG versions"
	if locale == "zh" {
		title = "全局矩阵"
		description = "跨操作系统与 PostgreSQL 大版本的扩展包可用性全局矩阵"
	}

	return fmt.Sprintf(`---
title: "%s"
linkTitle: "%s"
description: "%s"
weight: 699
layout: matrix
page_width: full
sidebar_enabled: false
breadcrumb: false
search_exclude: true
---
`, title, title, description)
}

func (g *GlobalMatrixGenerator) writeCSV(path string, rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create matrix csv: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	header := []string{"pkg", "ext", "url"}
	for _, osv := range osVersions {
		for _, pg := range pgVersions {
			header = append(header, fmt.Sprintf("%s_pg%d", osv.OS, pg))
		}
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("write matrix csv header: %w", err)
	}

	for _, row := range rows {
		record := []string{row.Pkg, row.Ext, fmt.Sprintf("%s/%s", globalMatrixBaseURL, row.Ext)}
		for _, osv := range osVersions {
			for _, pg := range pgVersions {
				record = append(record, row.Cells[osv.OS][pg].Code)
			}
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("write matrix csv row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("flush matrix csv: %w", err)
	}
	return nil
}

func (g *GlobalMatrixGenerator) writeJSON(path string, rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int, stats globalMatrixStats) error {
	type exportRow struct {
		Pkg        string            `json:"pkg"`
		Ext        string            `json:"ext"`
		URL        string            `json:"url"`
		Groups     map[string]string `json:"groups"`
		PGDGGroups map[string]string `json:"pgdg_groups"`
	}
	type exportData struct {
		PGVersions []int                    `json:"pg_versions"`
		OSVersions []string                 `json:"os_versions"`
		Legend     []globalMatrixLegendItem `json:"legend"`
		Stats      globalMatrixStats        `json:"stats"`
		Rows       []exportRow              `json:"rows"`
	}

	data := exportData{
		PGVersions: pgVersions,
		OSVersions: make([]string, 0, len(osVersions)),
		Legend:     globalMatrixLegend,
		Stats:      stats,
		Rows:       make([]exportRow, 0, len(rows)),
	}
	for _, osv := range osVersions {
		data.OSVersions = append(data.OSVersions, osv.OS)
	}
	for _, row := range rows {
		exported := exportRow{
			Pkg:        row.Pkg,
			Ext:        row.Ext,
			URL:        fmt.Sprintf("%s/%s", globalMatrixBaseURL, row.Ext),
			Groups:     make(map[string]string, len(osVersions)),
			PGDGGroups: make(map[string]string, len(osVersions)),
		}
		for _, osv := range osVersions {
			var group strings.Builder
			var pgdgGroup strings.Builder
			for _, pg := range pgVersions {
				group.WriteString(row.Cells[osv.OS][pg].Code)
				pgdgGroup.WriteString(row.PGDGCells[osv.OS][pg].Code)
			}
			exported.Groups[osv.OS] = group.String()
			exported.PGDGGroups[osv.OS] = pgdgGroup.String()
		}
		data.Rows = append(data.Rows, exported)
	}

	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal matrix json: %w", err)
	}
	payload = append(payload, '\n')
	if err := os.WriteFile(path, payload, 0644); err != nil {
		return fmt.Errorf("write matrix json: %w", err)
	}
	return nil
}

func (g *GlobalMatrixGenerator) renderStandaloneHTML(rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int, stats globalMatrixStats) string {
	var b strings.Builder
	b.WriteString(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>PG Extension Matrix</title>
`)
	b.WriteString(globalMatrixStandaloneCSS())
	b.WriteString(`</head>
<body data-view="pgdg">
`)
	b.WriteString(fmt.Sprintf("<main class=\"mx-page\" style=\"--mx-columns: %d;\">\n", max(1, len(osVersions)*len(pgVersions))))
	b.WriteString(`
  <header class="mx-hero">
    <h1>PG Extension Matrix</h1>
  </header>
`)
	b.WriteString(g.renderStandaloneStats(stats))
	b.WriteString(g.renderStandaloneToolbar(stats))
	b.WriteString(g.renderStandaloneTable(rows, osVersions, pgVersions))
	b.WriteString(`</main>
`)
	b.WriteString(globalMatrixStandaloneScript())
	b.WriteString(`
</body>
</html>
`)
	return b.String()
}

func (g *GlobalMatrixGenerator) renderStandaloneStats(stats globalMatrixStats) string {
	items := []struct {
		Label string
		Value string
	}{
		{"Non-core Extension Packages", fmt.Sprintf("%d", stats.Rows)},
		{"Linux OS Platforms", fmt.Sprintf("%d", stats.OS)},
		{"PG Major Versions", fmt.Sprintf("%d", stats.PG)},
		{"Package Cells", fmt.Sprintf("%d", stats.Cells)},
	}

	var b strings.Builder
	b.WriteString(`<section class="mx-stats" aria-label="Matrix summary">` + "\n")
	for _, item := range items {
		b.WriteString(fmt.Sprintf(`  <div class="mx-stat"><span>%s</span><strong>%s</strong></div>`+"\n", html.EscapeString(item.Label), html.EscapeString(item.Value)))
	}
	b.WriteString(`</section>` + "\n")
	return b.String()
}

func (g *GlobalMatrixGenerator) renderStandaloneToolbar(stats globalMatrixStats) string {
	var b strings.Builder
	b.WriteString(`<section class="mx-toolbar" aria-label="Legend and view controls">` + "\n")
	b.WriteString(`  <div class="mx-legend" aria-label="Legend">` + "\n")
	for _, item := range globalMatrixLegend {
		count := stats.Counts[item.Code]
		pgdgCount := stats.PGDGCounts[item.Code]
		b.WriteString(fmt.Sprintf(`  <span class="mx-legend-item" data-code="%s"><i class="mx-dot %s"></i><b>%s</b><em data-full-count="%d" data-pgdg-count="%d">%d</em></span>`+"\n",
			html.EscapeString(item.Code),
			html.EscapeString(item.Class),
			html.EscapeString(item.Label),
			count,
			pgdgCount,
			pgdgCount,
		))
	}
	b.WriteString(`  </div>` + "\n")
	b.WriteString(`  <div class="mx-tabs" role="tablist" aria-label="Matrix view">
    <button class="mx-tab is-active" type="button" role="tab" aria-selected="true" data-mode="pgdg">PGDG</button>
    <button class="mx-tab" type="button" role="tab" aria-selected="false" data-mode="full">FULL</button>
  </div>
</section>
`)
	return b.String()
}

func (g *GlobalMatrixGenerator) renderStandaloneTable(rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int) string {
	displayPGVersions := standalonePGVersions(pgVersions)

	var b strings.Builder
	b.WriteString(`<section class="mx-matrix" aria-label="Global package matrix">` + "\n")
	b.WriteString(`<table>` + "\n<colgroup>\n")
	b.WriteString(`  <col class="mx-col-pkg">` + "\n")
	for range osVersions {
		for range displayPGVersions {
			b.WriteString(`  <col class="mx-col-cell">` + "\n")
		}
	}
	b.WriteString(`</colgroup>` + "\n<thead>\n<tr>")
	b.WriteString(`<th class="mx-corner" rowspan="2">PKG</th>`)
	for _, osv := range osVersions {
		b.WriteString(fmt.Sprintf(`<th class="mx-os mx-group-start" colspan="%d"><a href="%s/%s">%s</a></th>`,
			len(displayPGVersions),
			globalMatrixOSBaseURL,
			html.EscapeString(osv.OS),
			html.EscapeString(osv.OS),
		))
	}
	b.WriteString("</tr>\n<tr>")
	for _, osv := range osVersions {
		for i, pg := range displayPGVersions {
			groupClass := ""
			if i == 0 {
				groupClass = " mx-group-start"
			}
			b.WriteString(fmt.Sprintf(`<th class="mx-pg%s"><a href="%s/%s">%d</a></th>`,
				groupClass,
				globalMatrixOSBaseURL,
				html.EscapeString(osv.OS),
				pg,
			))
		}
	}
	b.WriteString("</tr>\n</thead>\n<tbody>\n")

	for _, row := range rows {
		b.WriteString("<tr>")
		b.WriteString(fmt.Sprintf(`<th class="mx-pkg"><a href="%s/%s">%s</a></th>`, globalMatrixBaseURL, html.EscapeString(row.Ext), html.EscapeString(row.Pkg)))
		for _, osv := range osVersions {
			for i, pg := range displayPGVersions {
				cell := row.Cells[osv.OS][pg]
				pgdgCell := row.PGDGCells[osv.OS][pg]
				groupClass := ""
				if i == 0 {
					groupClass = " mx-group-start"
				}
				title := fmt.Sprintf("Full: %s\nPGDG-only: %s", cell.Title, pgdgCell.Title)
				b.WriteString(fmt.Sprintf(`<td class="mx-cell%s" data-code="%s" data-pgdg-code="%s" title="%s"><span class="mx-dot %s"></span></td>`,
					groupClass,
					html.EscapeString(cell.Code),
					html.EscapeString(pgdgCell.Code),
					html.EscapeString(title),
					html.EscapeString(cell.Class),
				))
			}
		}
		b.WriteString("</tr>\n")
	}

	b.WriteString("</tbody>\n</table>\n</section>\n")
	return b.String()
}

func standalonePGVersions(pgVersions []int) []int {
	display := append([]int(nil), pgVersions...)
	sort.Sort(sort.Reverse(sort.IntSlice(display)))
	return display
}

func globalMatrixStandaloneCSS() string {
	return "<style>\n" + matrixStandaloneCSS + "</style>\n"
}

func globalMatrixStandaloneScript() string {
	return "<script>\n" + strings.TrimRight(matrixStandaloneJS, "\n") + "\n</script>"
}
