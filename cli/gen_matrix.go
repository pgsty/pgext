/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

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
	if err := WriteMarkdownFile(filepath.Join(osDir, "matrix.md"), g.renderPage(rows, osVersions, pgVersions, stats, "en")); err != nil {
		return fmt.Errorf("write English matrix page: %w", err)
	}
	if err := WriteMarkdownFile(filepath.Join(osDir, "matrix.zh.md"), g.renderPage(rows, osVersions, pgVersions, stats, "zh")); err != nil {
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

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("load package matrix data: %w", err)
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

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("load PGDG-only package matrix data: %w", err)
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

func (g *GlobalMatrixGenerator) renderPage(rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int, stats globalMatrixStats, locale string) string {
	isZh := locale == "zh"
	title := "Global Matrix"
	description := "Global PostgreSQL extension package availability matrix across OS and PG versions"
	intro := fmt.Sprintf("This matrix compresses <strong>%d</strong> package rows across <strong>%d</strong> operating systems and <strong>%d</strong> PostgreSQL majors into <strong>%d</strong> colored slots.", stats.Rows, stats.OS, stats.PG, stats.Cells)
	if isZh {
		title = "全局矩阵"
		description = "跨操作系统与 PostgreSQL 大版本的扩展包可用性全局矩阵"
		intro = fmt.Sprintf("这个矩阵将 <strong>%d</strong> 个扩展包、<strong>%d</strong> 个操作系统、<strong>%d</strong> 个 PostgreSQL 大版本压缩为 <strong>%d</strong> 个彩色槽位。", stats.Rows, stats.OS, stats.PG, stats.Cells)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf(`---
title: "%s"
linkTitle: "%s"
description: "%s"
weight: 699
width: full
---

`, title, title, description))

	b.WriteString(globalMatrixCSS())
	b.WriteString(`<div class="gm-page">` + "\n")
	b.WriteString(fmt.Sprintf("<p>%s</p>\n", intro))
	b.WriteString(g.renderSummary(stats, isZh))
	b.WriteString(g.renderLegend(stats, isZh))
	b.WriteString(`<p class="gm-links"><a href="/matrix/pgext-global-matrix.csv">CSV</a> <a href="/matrix/pgext-global-matrix.json">JSON</a></p>` + "\n")
	b.WriteString(g.renderHTMLMatrix(rows, osVersions, pgVersions))
	b.WriteString(g.renderLetterMatrix(rows, osVersions, pgVersions, isZh))
	b.WriteString("</div>\n")

	return b.String()
}

func (g *GlobalMatrixGenerator) renderSummary(stats globalMatrixStats, isZh bool) string {
	labels := []string{"Packages", "OS", "PG", "Cells"}
	if isZh {
		labels = []string{"扩展包", "操作系统", "PG版本", "槽位"}
	}
	values := []int{stats.Rows, stats.OS, stats.PG, stats.Cells}

	var b strings.Builder
	b.WriteString(`<div class="gm-summary">`)
	for i, label := range labels {
		b.WriteString(fmt.Sprintf(`<div class="gm-stat"><strong>%d</strong><span>%s</span></div>`, values[i], label))
	}
	b.WriteString(`</div>` + "\n")
	return b.String()
}

func (g *GlobalMatrixGenerator) renderLegend(stats globalMatrixStats, isZh bool) string {
	labels := map[string]string{
		"B": "PGDG",
		"G": "Pigsty",
		"R": "Missing",
		".": "N/A",
	}
	if isZh {
		labels = map[string]string{
			"B": "PGDG",
			"G": "Pigsty",
			"R": "缺失",
			".": "N/A",
		}
	}

	var b strings.Builder
	b.WriteString(`<div class="gm-legend">`)
	for _, item := range globalMatrixLegend {
		count := stats.Counts[item.Code]
		b.WriteString(fmt.Sprintf(`<span><i class="gm-dot %s"></i><b>%s</b><em>%d</em></span>`, item.Class, labels[item.Code], count))
	}
	b.WriteString(`</div>` + "\n")
	return b.String()
}

func (g *GlobalMatrixGenerator) renderHTMLMatrix(rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int) string {
	var b strings.Builder
	b.WriteString(`<div class="gm-table-wrap">` + "\n")
	b.WriteString(`<table class="gm-table">` + "\n<thead>\n<tr>")
	b.WriteString(`<th class="gm-pkg-head" rowspan="2">PKG</th>`)
	for _, osv := range osVersions {
		b.WriteString(fmt.Sprintf(`<th class="gm-os-head gm-group-start" colspan="%d"><a href="/os/%s">%s</a></th>`, len(pgVersions), html.EscapeString(osv.OS), html.EscapeString(osv.OS)))
	}
	b.WriteString("</tr>\n<tr>")
	for range osVersions {
		for i, pg := range pgVersions {
			groupClass := ""
			if i == 0 {
				groupClass = " gm-group-start"
			}
			b.WriteString(fmt.Sprintf(`<th class="gm-pg-head%s">%d</th>`, groupClass, pg))
		}
	}
	b.WriteString("</tr>\n</thead>\n<tbody>\n")

	for _, row := range rows {
		b.WriteString("<tr>")
		b.WriteString(fmt.Sprintf(`<th class="gm-pkg-cell"><a href="%s/%s">%s</a></th>`, globalMatrixBaseURL, html.EscapeString(row.Ext), html.EscapeString(row.Pkg)))
		for _, osv := range osVersions {
			for i, pg := range pgVersions {
				cell := row.Cells[osv.OS][pg]
				groupClass := ""
				if i == 0 {
					groupClass = " gm-group-start"
				}
				b.WriteString(fmt.Sprintf(`<td class="gm-cell%s" title="%s" data-code="%s"><span class="gm-dot %s" aria-label="%s">%s</span></td>`,
					groupClass,
					html.EscapeString(cell.Title),
					html.EscapeString(cell.Code),
					html.EscapeString(cell.Class),
					html.EscapeString(cell.Title),
					html.EscapeString(cell.Code),
				))
			}
		}
		b.WriteString("</tr>\n")
	}

	b.WriteString("</tbody>\n</table>\n</div>\n")
	return b.String()
}

func (g *GlobalMatrixGenerator) renderLetterMatrix(rows []*GlobalMatrixRow, osVersions []OSVersion, pgVersions []int, isZh bool) string {
	summary := "Letter matrix data"
	if isZh {
		summary = "字母矩阵数据"
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf(`<details class="gm-letter"><summary>%s</summary><pre>`, summary))
	for i, row := range rows {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(html.EscapeString(formatGlobalMatrixLetterLine(row, osVersions, pgVersions)))
	}
	b.WriteString(`</pre></details>` + "\n")
	return b.String()
}

func formatGlobalMatrixLetterLine(row *GlobalMatrixRow, osVersions []OSVersion, pgVersions []int) string {
	groups := make([]string, 0, len(osVersions))
	for _, osv := range osVersions {
		var group strings.Builder
		for _, pg := range pgVersions {
			cell := row.Cells[osv.OS][pg]
			if cell.Code == "" {
				group.WriteByte('?')
			} else {
				group.WriteString(cell.Code)
			}
		}
		groups = append(groups, group.String())
	}
	return fmt.Sprintf("[%s](%s/%s) | %s", row.Pkg, globalMatrixBaseURL, row.Ext, strings.Join(groups, " | "))
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
<main class="mx-page">
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
	items := []globalMatrixLegendItem{
		{Code: "B", Label: "PGDG", Class: "gm-pgdg"},
		{Code: "G", Label: "PIGSTY", Class: "gm-pigsty"},
		{Code: "R", Label: "Missing", Class: "gm-missing"},
		{Code: ".", Label: "N/A", Class: "gm-na"},
	}
	for _, item := range items {
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
					html.EscapeString(standaloneCellClass(cell)),
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

func standaloneCellClass(cell GlobalMatrixCell) string { return cell.Class }

func globalMatrixStandaloneCSS() string {
	return `<style>
:root {
  --bg: #f4f6f9;
  --paper: #ffffff;
  --ink: #111827;
  --muted: #667085;
  --line: #d9dee7;
  --line-soft: #edf0f4;
  --pgdg: rgb(91, 156, 213);
  --pigsty: rgb(96, 190, 89);
  --na: rgb(108, 108, 108);
  --missing: rgb(204, 70, 55);
}
* { box-sizing: border-box; }
body {
  margin: 0;
  background: var(--bg);
  color: var(--ink);
  font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
}
a { color: inherit; text-decoration: none; }
a:hover { text-decoration: underline; }
.mx-page {
  container-type: inline-size;
  --pkg-col: clamp(142px, 12cqw, 184px);
  --cell-size: max(7px, calc((100cqw - var(--pkg-col)) / 80));
  width: 100%;
  padding: 24px 32px 48px;
}
.mx-hero {
  margin-bottom: 16px;
}
h1 {
  margin: 0;
  font-size: clamp(40px, 4.8vw, 64px);
  line-height: 1;
  letter-spacing: 0;
}
.mx-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  margin-bottom: 14px;
}
.mx-tabs {
  display: inline-flex;
  flex: 0 0 auto;
  padding: 3px;
  border: 1px solid var(--line);
  border-radius: 9px;
  background: var(--paper);
  box-shadow: 0 8px 20px rgba(15, 23, 42, .05);
}
.mx-tab {
  min-width: 72px;
  border: 0;
  border-radius: 6px;
  padding: 8px 11px;
  background: transparent;
  color: var(--muted);
  font: inherit;
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 0;
  cursor: pointer;
}
.mx-tab.is-active {
  background: #111827;
  color: #fff;
}
.mx-stats {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 12px;
  margin-bottom: 12px;
}
.mx-stat {
  min-height: 78px;
  border: 1px solid var(--line);
  border-radius: 8px;
  padding: 13px 15px;
  background: var(--paper);
  box-shadow: 0 8px 18px rgba(15, 23, 42, .04);
}
.mx-stat span {
  display: block;
  color: var(--muted);
  font-size: clamp(11px, 1.1cqw, 14px);
  font-weight: 700;
  line-height: 1.2;
}
.mx-stat strong {
  display: block;
  margin-top: 8px;
  font-size: clamp(30px, 3.2cqw, 44px);
  line-height: 1;
  letter-spacing: 0;
}
.mx-legend {
  display: flex;
  flex-wrap: wrap;
  flex: 1 1 auto;
  gap: 8px;
  align-items: center;
}
.mx-legend-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  min-height: 25px;
  border: 1px solid var(--line);
  border-radius: 999px;
  padding: 5px 8px;
  background: var(--paper);
  font-size: 12px;
  line-height: 1;
}
.mx-legend-item em {
  color: var(--muted);
  font-style: normal;
}
body[data-view="pgdg"] .mx-legend-item:not([data-code="B"]) { display: none; }
.mx-matrix {
  width: 100%;
  overflow: visible;
  border: 1px solid var(--line);
  border-radius: 8px;
  background: var(--paper);
  box-shadow: 0 12px 30px rgba(15, 23, 42, .08);
}
.mx-matrix table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}
.mx-col-pkg { width: var(--pkg-col); }
.mx-col-cell { width: var(--cell-size); }
.mx-matrix th,
.mx-matrix td {
  border-right: 1px solid var(--line-soft);
  border-bottom: 1px solid var(--line-soft);
  padding: 0;
  text-align: center;
  vertical-align: middle;
}
.mx-matrix thead th {
  position: sticky;
  top: 0;
  z-index: 4;
  background: #f1f3f7;
}
.mx-matrix thead tr:nth-child(2) th {
  top: 24px;
}
.mx-corner,
.mx-pkg {
  position: sticky;
  left: 0;
  z-index: 5;
  width: var(--pkg-col);
  min-width: var(--pkg-col);
  max-width: var(--pkg-col);
  background: #fff !important;
  text-align: left !important;
}
.mx-corner {
  z-index: 8 !important;
  height: 40px;
  padding-left: 8px !important;
  background: #f1f3f7 !important;
  font-size: 11px;
}
.mx-os {
  height: 24px;
  color: #344054;
  font-size: clamp(8px, calc(var(--cell-size) * .76), 11px);
  font-weight: 800;
}
.mx-pg {
  width: var(--cell-size);
  min-width: var(--cell-size);
  max-width: var(--cell-size);
  height: 16px;
  color: #667085;
  font-size: clamp(6px, calc(var(--cell-size) * .54), 9px);
  font-weight: 700;
  line-height: 1;
}
.mx-pkg {
  height: var(--cell-size);
  padding: 0 8px !important;
  font-size: clamp(8px, calc(var(--cell-size) * .72), 11px);
  line-height: var(--cell-size);
  font-weight: 700;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.mx-cell {
  width: var(--cell-size);
  min-width: var(--cell-size);
  max-width: var(--cell-size);
  height: var(--cell-size);
  line-height: 0;
}
.mx-group-start {
  border-left: 2px solid #c3cad6 !important;
}
.mx-dot {
  display: inline-block;
  width: max(4px, calc(var(--cell-size) - 4px));
  height: max(4px, calc(var(--cell-size) - 4px));
  margin: auto;
  border-radius: 2px;
  box-shadow: inset 0 0 0 1px rgba(0,0,0,.08);
  vertical-align: middle;
}
.gm-pgdg { background: var(--pgdg); }
.gm-pigsty { background: var(--pigsty); }
.gm-na { background: var(--na); }
.gm-missing { background: var(--missing); }
.mx-dimmed { background: rgb(108, 108, 108); }
body[data-view="pgdg"] .mx-cell .mx-dot {
  background: #d8dde6;
  opacity: .24;
  box-shadow: none;
}
body[data-view="pgdg"] .mx-cell[data-pgdg-code="B"] .mx-dot {
  background: var(--pgdg);
  opacity: 1;
  box-shadow: 0 0 0 1px rgba(91,156,213,.30), 0 0 8px rgba(91,156,213,.20);
}
@media (max-width: 900px) {
  .mx-page {
    --pkg-col: 126px;
    padding: 18px 12px 32px;
  }
  .mx-stats { grid-template-columns: repeat(2, minmax(0, 1fr)); }
  .mx-toolbar { align-items: stretch; flex-direction: column; }
  .mx-tabs { width: max-content; }
}
@media print {
  body { background: #fff; }
  .mx-page { padding: 16px; }
  .mx-tabs { display: none; }
  .mx-matrix { box-shadow: none; }
}
</style>
`
}

func globalMatrixStandaloneScript() string {
	return `<script>
const tabs = document.querySelectorAll(".mx-tab");
const legendCounts = document.querySelectorAll(".mx-legend-item em[data-full-count]");
function applyMatrixMode(mode) {
  document.body.dataset.view = mode;
  for (const count of legendCounts) {
    count.textContent = mode === "pgdg" ? count.dataset.pgdgCount : count.dataset.fullCount;
  }
}
for (const tab of tabs) {
  tab.addEventListener("click", () => {
    applyMatrixMode(tab.dataset.mode);
    for (const item of tabs) {
      const active = item === tab;
      item.classList.toggle("is-active", active);
      item.setAttribute("aria-selected", active ? "true" : "false");
    }
  });
}
applyMatrixMode(document.body.dataset.view || "pgdg");
</script>`
}

func globalMatrixCSS() string {
	return `<style>
.gm-page {
  --gm-bg: #ffffff;
  --gm-text: #111827;
  --gm-muted: #6b7280;
  --gm-border: #d1d5db;
  --gm-soft: #f8fafc;
  --gm-pgdg: rgb(91, 156, 213);
  --gm-pigsty: rgb(96, 190, 89);
  --gm-na: rgb(108, 108, 108);
  --gm-missing: rgb(204, 70, 55);
  color: var(--gm-text);
}
.gm-summary {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 10px;
  margin: 18px 0 14px;
}
.gm-stat {
  border: 1px solid var(--gm-border);
  border-radius: 8px;
  padding: 10px 12px;
  background: var(--gm-soft);
}
.gm-stat strong {
  display: block;
  font-size: 24px;
  line-height: 1.1;
}
.gm-stat span {
  display: block;
  color: var(--gm-muted);
  font-size: 12px;
  margin-top: 2px;
}
.gm-legend {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin: 10px 0 8px;
}
.gm-legend span {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border: 1px solid var(--gm-border);
  border-radius: 999px;
  padding: 5px 8px;
  font-size: 12px;
  line-height: 1;
  background: #fff;
}
.gm-legend em {
  color: var(--gm-muted);
  font-style: normal;
}
.gm-links {
  font-size: 13px;
  margin: 8px 0 14px;
}
.gm-links a {
  margin-right: 12px;
}
.gm-table-wrap {
  max-height: 78vh;
  overflow: auto;
  border: 1px solid var(--gm-border);
  border-radius: 8px;
  background: #fff;
}
.gm-table {
  border-collapse: separate;
  border-spacing: 0;
  table-layout: fixed;
  font-size: 11px;
  line-height: 1;
  width: max-content;
  min-width: 100%;
}
.gm-table th,
.gm-table td {
  border-right: 1px solid #edf0f4;
  border-bottom: 1px solid #edf0f4;
  padding: 0;
  text-align: center;
  vertical-align: middle;
}
.gm-table thead th {
  position: sticky;
  top: 0;
  z-index: 5;
  background: #f3f4f6;
  color: #374151;
}
.gm-table thead tr:nth-child(2) th {
  top: 24px;
}
.gm-pkg-head,
.gm-pkg-cell {
  position: sticky;
  left: 0;
  z-index: 6;
  width: 190px;
  min-width: 190px;
  max-width: 190px;
  background: #fff;
  text-align: left !important;
}
.gm-pkg-head {
  top: 0;
  z-index: 9 !important;
  padding-left: 8px !important;
  background: #f3f4f6 !important;
}
.gm-pkg-cell {
  padding: 0 8px !important;
  height: 20px;
  font-weight: 500;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.gm-pkg-cell a {
  color: #111827;
  text-decoration: none;
}
.gm-pkg-cell a:hover {
  text-decoration: underline;
}
.gm-os-head {
  height: 24px;
  min-width: 76px;
  font-size: 11px;
  font-weight: 700;
}
.gm-os-head a {
  color: #374151;
  text-decoration: none;
}
.gm-pg-head {
  width: 15px;
  min-width: 15px;
  height: 18px;
  font-size: 9px;
  font-weight: 600;
  color: #6b7280;
}
.gm-cell {
  width: 15px;
  min-width: 15px;
  height: 20px;
}
.gm-group-start {
  border-left: 2px solid #cbd5e1 !important;
}
.gm-dot {
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 2px;
  font-size: 0;
  color: transparent;
  box-shadow: inset 0 0 0 1px rgba(0,0,0,.08);
}
.gm-pgdg { background: var(--gm-pgdg); }
.gm-pigsty { background: var(--gm-pigsty); }
.gm-na { background: var(--gm-na); }
.gm-missing { background: var(--gm-missing); }
.gm-letter {
  margin-top: 18px;
}
.gm-letter summary {
  cursor: pointer;
  font-weight: 700;
}
.gm-letter pre {
  max-height: 420px;
  overflow: auto;
  border: 1px solid var(--gm-border);
  border-radius: 8px;
  padding: 12px;
  font-size: 11px;
  line-height: 1.45;
  background: #0f172a;
  color: #e5e7eb;
}
@media (max-width: 760px) {
  .gm-summary {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
  .gm-pkg-head,
  .gm-pkg-cell {
    width: 132px;
    min-width: 132px;
    max-width: 132px;
  }
}
</style>
`
}
