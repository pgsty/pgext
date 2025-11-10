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
	ID       int
	Title    string
	Total    int
	PGDG     int
	PIGSTY   int
	CONTRIB  int
	MISS     int
	PGCounts map[int]int // Dynamic PG version counts (e.g., 18->395, 17->419, etc.)
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

	// Get active PostgreSQL versions first
	pgVersions, err := g.getActivePGVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get active PG versions: %w", err)
	}
	stats.PGVersions = pgVersions

	// Build column list for dynamic PG versions
	pgColumnList := make([]string, 0)
	for _, pg := range pgVersions {
		pgColumnList = append(pgColumnList, fmt.Sprintf("pg%d", pg))
	}

	// Query the summary view directly
	query := fmt.Sprintf(`
		SELECT id, title, total, pgdg, pigsty, contrib, miss, %s
		FROM pgext.summary
		ORDER BY id
	`, strings.Join(pgColumnList, ", "))

	rows, err := g.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query summary: %w", err)
	}
	defer rows.Close()

	allRows := make([]StatsRow, 0, 6)

	for rows.Next() {
		row := StatsRow{
			PGCounts: make(map[int]int),
		}

		// Build scan destinations
		scanDest := []interface{}{&row.ID, &row.Title, &row.Total, &row.PGDG, &row.PIGSTY, &row.CONTRIB, &row.MISS}

		// Create variables to hold PG version counts
		pgCounts := make([]int, len(pgVersions))
		for i := range pgVersions {
			scanDest = append(scanDest, &pgCounts[i])
		}

		// Scan the row
		if err := rows.Scan(scanDest...); err != nil {
			return nil, fmt.Errorf("failed to scan summary row: %w", err)
		}

		// Map PG counts to version numbers
		for i, pg := range pgVersions {
			row.PGCounts[pg] = pgCounts[i]
		}

		allRows = append(allRows, row)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading summary rows: %w", err)
	}

	// Split results into Extension Stats (rows 1-3) and Package Stats (rows 4-6)
	if len(allRows) >= 6 {
		stats.ExtensionStats = allRows[0:3]
		stats.PackageStats = allRows[3:6]
	} else {
		return nil, fmt.Errorf("expected at least 6 rows from summary, got %d", len(allRows))
	}

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

// Removed obsolete functions - now using pgext.summary view directly

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
		totalExts = stats.ExtensionStats[0].Total // ALL row
	}
	if len(stats.PackageStats) > 0 {
		totalPkgs = stats.PackageStats[0].Total // ALL row
	}

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "扩展目录"
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
		b.WriteString("※ 一个扩展软件包可能同时包含多个 PG 扩展，因此按软件包统计的数量会少于扩展数量。\n\n")
	} else {
		b.WriteString("## Package Stat\n\n")
		b.WriteString("※ One extension package may consist of multiple extension\n\n")
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
		// Get label from title - extract just the first word (ALL/EL/Debian)
		label := strings.Split(row.Title, " ")[0]
		b.WriteString(fmt.Sprintf("| **%s** |", label))
		b.WriteString(fmt.Sprintf(" %d |", row.Total))
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

// Removed unused getCategoryTitle function - categories are handled by shortcodes
