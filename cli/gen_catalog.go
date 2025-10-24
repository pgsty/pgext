/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// CatalogStats holds statistics for the catalog
type CatalogStats struct {
	ExtensionStats []StatsRow
	PackageStats   []StatsRow
	PGVersions     []int
}

// StatsRow represents a single row in the statistics table
type StatsRow struct {
	Label    string
	All      int
	PGDG     int
	PIGSTY   int
	CONTRIB  int
	MISS     int
	PGCounts map[int]int // Dynamic PG version counts
}

// CategoryExtensions holds category and its extensions
type CategoryExtensions struct {
	Category   string
	Extensions []struct {
		ID   int
		Name string
		Pkg  string
	}
}

// GenerateCatalogPage generates catalog list pages (_index.md and _index.zh.md)
func (g *ExtensionGenerator) GenerateCatalogPage(locale, outputPath string) error {
	// Fetch statistics
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	// Fetch categories with extensions
	categories, err := g.fetchCategoriesWithExtensions()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	content := g.generateCatalogContent(stats, categories, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) fetchCatalogStats() (*CatalogStats, error) {
	stats := &CatalogStats{}
	ctx := context.Background()

	// Get active PostgreSQL versions
	pgVersions, err := g.getActivePGVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get active PG versions: %w", err)
	}
	stats.PGVersions = pgVersions

	// Build dynamic SQL for PG version columns
	pgColumns := g.buildPGColumnSQL(pgVersions)

	// Fetch extension statistics (3 rows: ALL, EL, Debian)
	extStats, err := g.fetchExtensionStats(ctx, pgVersions, pgColumns)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch extension stats: %w", err)
	}
	stats.ExtensionStats = extStats

	// Fetch package statistics (3 rows: ALL, EL, Debian)
	pkgStats, err := g.fetchPackageStats(ctx, pgVersions, pgColumns)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch package stats: %w", err)
	}
	stats.PackageStats = pkgStats

	return stats, nil
}

func (g *ExtensionGenerator) getActivePGVersions(ctx context.Context) ([]int, error) {
	query := `SELECT pg FROM pgext.active_pg ORDER BY pg DESC`
	rows, err := g.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []int
	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err != nil {
			return nil, err
		}
		versions = append(versions, pg)
	}
	return versions, rows.Err()
}

func (g *ExtensionGenerator) buildPGColumnSQL(pgVersions []int) string {
	var columns []string
	for _, pg := range pgVersions {
		columns = append(columns, fmt.Sprintf(`count(*) FILTER ( WHERE pg_ver @> '{%d}') AS pg%d`, pg, pg))
	}
	if len(columns) == 0 {
		return ""
	}
	return ",\n       " + strings.Join(columns, ",\n       ")
}

func (g *ExtensionGenerator) fetchExtensionStats(ctx context.Context, pgVersions []int, pgColumns string) ([]StatsRow, error) {
	// SQL 1: ALL extensions
	allQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       0 AS miss%s
FROM pgext.extension`, pgColumns)

	allRow, err := g.queryStatsRow(ctx, allQuery, "ALL", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query ALL extension stats: %w", err)
	}

	// SQL 2: EL extensions
	elPGColumns := g.buildPGColumnsForEL(pgVersions)
	elQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       (SELECT count(*) FROM pgext.extension) - count(*) AS miss%s
FROM pgext.extension WHERE rpm_repo IS NOT NULL`, elPGColumns)

	elRow, err := g.queryStatsRow(ctx, elQuery, "EL", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query EL extension stats: %w", err)
	}

	// SQL 3: Debian extensions
	debPGColumns := g.buildPGColumnsForDeb(pgVersions)
	debianQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE deb_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE deb_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       (SELECT count(*) FROM pgext.extension) - count(*) AS miss%s
FROM pgext.extension WHERE deb_repo IS NOT NULL`, debPGColumns)

	debianRow, err := g.queryStatsRow(ctx, debianQuery, "Debian", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query Debian extension stats: %w", err)
	}

	return []StatsRow{allRow, elRow, debianRow}, nil
}

func (g *ExtensionGenerator) buildPGColumnsForEL(pgVersions []int) string {
	var columns []string
	for _, pg := range pgVersions {
		columns = append(columns, fmt.Sprintf(`count(*) FILTER ( WHERE rpm_pg @> '{%d}') AS pg%d`, pg, pg))
	}
	if len(columns) == 0 {
		return ""
	}
	return ",\n       " + strings.Join(columns, ",\n       ")
}

func (g *ExtensionGenerator) buildPGColumnsForDeb(pgVersions []int) string {
	var columns []string
	for _, pg := range pgVersions {
		columns = append(columns, fmt.Sprintf(`count(*) FILTER ( WHERE deb_pg @> '{%d}') AS pg%d`, pg, pg))
	}
	if len(columns) == 0 {
		return ""
	}
	return ",\n       " + strings.Join(columns, ",\n       ")
}

func (g *ExtensionGenerator) fetchPackageStats(ctx context.Context, pgVersions []int, pgColumns string) ([]StatsRow, error) {
	// SQL 4: ALL packages (lead extensions only)
	allQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       0 AS miss%s
FROM pgext.extension WHERE lead`, pgColumns)

	allRow, err := g.queryStatsRow(ctx, allQuery, "ALL", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query ALL package stats: %w", err)
	}

	// SQL 5: EL packages (lead extensions only)
	elPGColumns := g.buildPGColumnsForEL(pgVersions)
	elQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       (SELECT count(*) FROM pgext.extension WHERE lead) - count(*) AS miss%s
FROM pgext.extension WHERE lead AND rpm_repo IS NOT NULL`, elPGColumns)

	elRow, err := g.queryStatsRow(ctx, elQuery, "EL", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query EL package stats: %w", err)
	}

	// SQL 6: Debian packages (lead extensions only)
	debPGColumns := g.buildPGColumnsForDeb(pgVersions)
	debianQuery := fmt.Sprintf(`SELECT count(*) AS total,
       count(*) FILTER ( WHERE deb_repo = 'PGDG' ) AS pgdg,
       count(*) FILTER ( WHERE deb_repo = 'PIGSTY' ) AS pigsty,
       count(*) FILTER ( WHERE contrib ) AS contrib,
       (SELECT count(*) FROM pgext.extension WHERE lead) - count(*) AS miss%s
FROM pgext.extension WHERE lead AND deb_repo IS NOT NULL`, debPGColumns)

	debianRow, err := g.queryStatsRow(ctx, debianQuery, "Debian", pgVersions)
	if err != nil {
		return nil, fmt.Errorf("failed to query Debian package stats: %w", err)
	}

	return []StatsRow{allRow, elRow, debianRow}, nil
}

func (g *ExtensionGenerator) queryStatsRow(ctx context.Context, query, label string, pgVersions []int) (StatsRow, error) {
	row := StatsRow{
		Label:    label,
		PGCounts: make(map[int]int),
	}

	// Build scan destinations
	scanDest := []interface{}{&row.All, &row.PGDG, &row.PIGSTY, &row.CONTRIB, &row.MISS}

	// Create variables to hold PG version counts
	pgCounts := make([]int, len(pgVersions))
	for i := range pgVersions {
		scanDest = append(scanDest, &pgCounts[i])
	}

	// Execute query and scan results
	err := g.DB.QueryRow(ctx, query).Scan(scanDest...)
	if err != nil {
		return row, err
	}

	// Copy PG counts to map
	for i, pg := range pgVersions {
		row.PGCounts[pg] = pgCounts[i]
	}

	return row, nil
}

func (g *ExtensionGenerator) fetchCategoriesWithExtensions() ([]CategoryExtensions, error) {
	ctx := context.Background()
	query := `
		SELECT e.category, e.id, e.name, COALESCE(e.pkg, e.name) as pkg
		FROM pgext.extension e
		WHERE e.category IS NOT NULL
		ORDER BY e.category, e.id
	`

	rows, err := g.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	categoryMap := make(map[string]*CategoryExtensions)
	var categories []string

	for rows.Next() {
		var category string
		var ext struct {
			ID   int
			Name string
			Pkg  string
		}

		err := rows.Scan(&category, &ext.ID, &ext.Name, &ext.Pkg)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if _, exists := categoryMap[category]; !exists {
			categoryMap[category] = &CategoryExtensions{
				Category: category,
				Extensions: []struct {
					ID   int
					Name string
					Pkg  string
				}{},
			}
			categories = append(categories, category)
		}

		categoryMap[category].Extensions = append(categoryMap[category].Extensions, ext)
	}

	// Convert map to slice in order
	result := make([]CategoryExtensions, 0, len(categories))
	for _, cat := range categories {
		result = append(result, *categoryMap[cat])
	}

	return result, nil
}

func (g *ExtensionGenerator) generateCatalogContent(stats *CatalogStats, categories []CategoryExtensions, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	// Calculate totals from stats
	totalExts := 0
	totalPkgs := 0
	if len(stats.ExtensionStats) > 0 {
		totalExts = stats.ExtensionStats[0].All // ALL row
	}
	if len(stats.PackageStats) > 0 {
		totalPkgs = stats.PackageStats[0].All // ALL row
	}

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "目录"
weight: 200
excludeSearch: true
comments: false
---

`)
		b.WriteString(fmt.Sprintf("PostgreSQL 扩展目录包含了 **%d** 个扩展和 **%d** 个包。\n\n",
			totalExts, totalPkgs))
	} else {
		b.WriteString(`---
title: "Catalog"
weight: 200
excludeSearch: true
comments: false
---

`)
		b.WriteString(fmt.Sprintf("The PostgreSQL Extension Catalog contains **%d** extensions and **%d** packages.\n\n",
			totalExts, totalPkgs))
	}

	// Statistics Tables
	b.WriteString(g.generateStatisticsTables(stats, isZh))

	// Categories section
	b.WriteString(g.generateCategoriesSection(categories, isZh))

	return b.String()
}

func (g *ExtensionGenerator) generateStatisticsTables(stats *CatalogStats, isZh bool) string {
	var b strings.Builder

	// Extensions table
	if isZh {
		b.WriteString("## 扩展统计\n\n")
	} else {
		b.WriteString("## Extension Stat\n\n")
	}
	b.WriteString(g.generateStatsTable(stats.ExtensionStats, stats.PGVersions, isZh, true))
	b.WriteString("\n")

	// Packages table
	if isZh {
		b.WriteString("## 扩展包统计\n\n")
	} else {
		b.WriteString("## Package Stat\n\n")
	}
	b.WriteString(g.generateStatsTable(stats.PackageStats, stats.PGVersions, isZh, false))
	b.WriteString("\n")

	return b.String()
}

func (g *ExtensionGenerator) generateStatsTable(rows []StatsRow, pgVersions []int, isZh, isExtension bool) string {
	var b strings.Builder

	// Table header
	b.WriteString("|")
	if isZh {
		b.WriteString(" **分类** |")
	} else {
		b.WriteString(" **Category** |")
	}
	b.WriteString(" **All** | **PGDG** | **PIGSTY** | **CONTRIB** | **MISS** |")

	// Add dynamic PG version columns
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")

	// Table separator
	b.WriteString("|")
	b.WriteString(":---------|")
	b.WriteString("--------:|")    // All
	b.WriteString("--------:|")    // PGDG
	b.WriteString("----------:|")  // PIGSTY
	b.WriteString("-----------:|") // CONTRIB
	b.WriteString("--------:|")    // MISS
	for range pgVersions {
		b.WriteString("--------:|")
	}
	b.WriteString("\n")

	// Table rows
	for _, row := range rows {
		b.WriteString(fmt.Sprintf("| **%s** |", row.Label))
		b.WriteString(fmt.Sprintf(" %d |", row.All))
		b.WriteString(fmt.Sprintf(" %d |", row.PGDG))
		b.WriteString(fmt.Sprintf(" %d |", row.PIGSTY))
		b.WriteString(fmt.Sprintf(" %d |", row.CONTRIB))
		b.WriteString(fmt.Sprintf(" %d |", row.MISS))

		// Add PG version counts
		for _, pg := range pgVersions {
			count := 0
			if val, ok := row.PGCounts[pg]; ok {
				count = val
			}
			b.WriteString(fmt.Sprintf(" %d |", count))
		}
		b.WriteString("\n")
	}

	return b.String()
}

func (g *ExtensionGenerator) generateCategoriesSection(categories []CategoryExtensions, isZh bool) string {
	var b strings.Builder

	// Section header
	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}

	// Table header
	if isZh {
		b.WriteString("| **类别** | **扩展** |\n")
		b.WriteString("|:--------:|:---------|")
	} else {
		b.WriteString("| **Category** | **Extensions** |\n")
		b.WriteString("|:------------:|:---------------|")
	}
	b.WriteString("\n")

	// Table rows - one row per category
	for _, cat := range categories {
		// Category column with shortcode
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} |", categoryLower))

		// Extensions column - all extensions as shortcodes
		b.WriteString(" ")
		for i, ext := range cat.Extensions {
			// Use ext shortcode format
			b.WriteString(fmt.Sprintf(`{{< ext "%s" >}}`, ext.Name))

			// Add space between shortcodes
			if i < len(cat.Extensions)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString(" |\n")
	}

	return b.String()
}

func getCategoryTitle(category string, isZh bool) string {
	if isZh {
		switch category {
		case "TIME":
			return "时间与日期"
		case "GIS":
			return "地理空间"
		case "RAG":
			return "RAG向量"
		case "FTS":
			return "全文搜索"
		case "OLAP":
			return "数据分析"
		case "FEAT":
			return "功能扩展"
		case "LANG":
			return "编程语言"
		case "TYPE":
			return "数据类型"
		case "FUNC":
			return "函数与操作符"
		case "INDEX":
			return "索引访问"
		case "FDW":
			return "外部数据源"
		case "SIM":
			return "仿真与兼容性"
		case "ETL":
			return "ETL与迁移"
		case "TOOL":
			return "工具与实用程序"
		case "STAT":
			return "统计与监控"
		case "ADMIN":
			return "管理与安全"
		case "AUDIT":
			return "审计与日志"
		case "SEC":
			return "安全与合规"
		case "SHARD":
			return "分片与MPP"
		case "REP":
			return "流式复制"
		case "EXTRA":
			return "其他扩展"
		default:
			return category
		}
	}
	// English categories (already in proper case in database)
	return category
}
