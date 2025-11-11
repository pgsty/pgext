/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
)

// GeneratePackageList generates package list pages (pkg.md and pkg.zh.md)
// This focuses on packages (lead extensions) with statistics, categories, and full package index
func (g *ExtensionGenerator) GeneratePackageList(locale, outputPath string) error {
	// Fetch statistics (reuse from catalog generation)
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	// Fetch categories with packages (sorted by category ID)
	categories, err := g.fetchCategoriesWithPackagesSorted()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	// Filter to get lead packages only
	var packages []*Extension
	for _, ext := range g.Cache.Extensions {
		// Skip non-ready extensions
		if ext.State.Valid && ext.State.String == "not-ready" {
			continue
		}
		// Only include lead extensions (packages)
		if ext.Lead {
			packages = append(packages, ext)
		}
	}

	// Sort packages by ID
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].ID < packages[j].ID
	})

	content := g.generatePackageListContent(stats, categories, packages, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// CategoryPackages holds category and its packages (lead extensions)
type CategoryPackages struct {
	Category string
	Packages []struct {
		ID   int
		Name string
		Pkg  string
	}
}

// fetchCategoriesWithPackagesSorted fetches categories with packages sorted by category ID
func (g *ExtensionGenerator) fetchCategoriesWithPackagesSorted() ([]CategoryPackages, error) {
	ctx := context.Background()

	// First get the category order from category table
	categoryQuery := `
		SELECT name, id
		FROM pgext.category
		ORDER BY id
	`

	categoryRows, err := g.DB.Query(ctx, categoryQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query category order: %w", err)
	}
	defer categoryRows.Close()

	var categoryOrder []string
	categoryOrderMap := make(map[string]int)
	idx := 0
	for categoryRows.Next() {
		var name string
		var id int
		if err := categoryRows.Scan(&name, &id); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categoryOrder = append(categoryOrder, name)
		categoryOrderMap[name] = idx
		idx++
	}

	// Then get packages grouped by category (only lead extensions)
	query := `
		SELECT e.category, e.id, e.name, COALESCE(e.pkg, e.name) as pkg
		FROM pgext.extension e
		WHERE e.category IS NOT NULL AND e.lead = true
		ORDER BY e.category, e.id
	`

	rows, err := g.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query packages: %w", err)
	}
	defer rows.Close()

	categoryMap := make(map[string]*CategoryPackages)

	for rows.Next() {
		var category string
		var pkg struct {
			ID   int
			Name string
			Pkg  string
		}

		err := rows.Scan(&category, &pkg.ID, &pkg.Name, &pkg.Pkg)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if _, exists := categoryMap[category]; !exists {
			categoryMap[category] = &CategoryPackages{
				Category: category,
				Packages: []struct {
					ID   int
					Name string
					Pkg  string
				}{},
			}
		}

		categoryMap[category].Packages = append(categoryMap[category].Packages, pkg)
	}

	// Convert map to slice in the correct order
	result := make([]CategoryPackages, 0, len(categoryOrder))
	for _, cat := range categoryOrder {
		if catData, exists := categoryMap[cat]; exists {
			result = append(result, *catData)
		}
	}

	return result, nil
}

func (g *ExtensionGenerator) generatePackageListContent(stats *CatalogStats, categories []CategoryPackages, packages []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	// Calculate total packages
	totalPkgs := 0
	if len(stats.PackageStats) > 0 {
		totalPkgs = stats.PackageStats[0].Total // ALL row
	}

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "扩展包清单"
weight: 20
excludeSearch: true
comments: false
---

`)
	} else {
		b.WriteString(`---
title: "Packages"
weight: 20
excludeSearch: true
comments: false
---

`)
	}

	// Package Statistics section
	if isZh {
		b.WriteString("## 统计\n\n")
		b.WriteString("※ 一个扩展软件包可能同时包含多个 PG 扩展，因此按软件包统计的数量会少于扩展数量。\n\n")
	} else {
		b.WriteString("## Statistics\n\n")
		b.WriteString("※ One extension package may consist of multiple extensions\n\n")
	}
	b.WriteString(g.generatePackageStatsTable(stats.PackageStats, stats.PGVersions, isZh))
	b.WriteString("\n")

	// Categories section with packages
	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}
	b.WriteString(g.generatePackageCategoriesTable(categories, isZh))
	b.WriteString("\n")

	// Full package list section
	if isZh {
		b.WriteString(fmt.Sprintf("## 扩展包列表\n\n共有 %d 个可用的 PostgreSQL 扩展包：\n\n", totalPkgs))
	} else {
		b.WriteString(fmt.Sprintf("## Packages\n\nThere are %d available PostgreSQL packages:\n\n", totalPkgs))
	}
	b.WriteString(g.generateFullPackageTable(packages, isZh))

	return b.String()
}

// generatePackageStatsTable generates the statistics table for packages
func (g *ExtensionGenerator) generatePackageStatsTable(rows []StatsRow, pgVersions []int, isZh bool) string {
	var b strings.Builder

	// Table header
	b.WriteString("|")
	if isZh {
		b.WriteString(" **类别** |")
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
	b.WriteString(":-------------|")
	b.WriteString("--------:|")     // All
	b.WriteString("---------:|")    // PGDG
	b.WriteString("-----------:|")  // PIGSTY
	b.WriteString("------------:|") // CONTRIB
	b.WriteString("---------:|")    // MISS
	for range pgVersions {
		b.WriteString("---------:|")
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

// generatePackageCategoriesTable generates the categories table with packages and count
func (g *ExtensionGenerator) generatePackageCategoriesTable(categories []CategoryPackages, isZh bool) string {
	var b strings.Builder

	// Table header
	if isZh {
		b.WriteString("| **类别** | **数** | **扩展包** |\n")
		b.WriteString("|:------------:|:--------:|:---------------|\n")
	} else {
		b.WriteString("| **Category** | **Count** | **Packages** |\n")
		b.WriteString("|:------------:|:---------:|:---------------|\n")
	}

	// Table rows - one row per category
	for _, cat := range categories {
		// Category column with shortcode
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} |", categoryLower))

		// Count column
		b.WriteString(fmt.Sprintf(" %d |", len(cat.Packages)))

		// Packages column - all packages as two-parameter shortcodes
		b.WriteString(" ")
		for i, pkg := range cat.Packages {
			// Use ext shortcode with two parameters (name, pkg)
			b.WriteString(fmt.Sprintf(`{{< ext "%s" "%s" >}}`, pkg.Name, pkg.Pkg))

			// Add space between shortcodes
			if i < len(cat.Packages)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString(" |\n")
	}

	return b.String()
}

// generateFullPackageTable generates the full package list table
func (g *ExtensionGenerator) generateFullPackageTable(packages []*Extension, isZh bool) string {
	var b strings.Builder

	// Table header
	if isZh {
		b.WriteString("| 包 | 版本 | 仓库 | 分类 | RPM | DEB |\n")
		b.WriteString("|:---|:-----|:-----|:-----|:-----|:-----|\n")
	} else {
		b.WriteString("| Package | Version | Repo | Category | RPM | DEB |\n")
		b.WriteString("|:--------|:--------|:-----|:---------|:-----|:-----|\n")
	}

	// Table rows
	for _, pkg := range packages {
		// Package link - use two-parameter shortcode
		pkgLink := ExtShortcode(pkg.Name, pkg.Pkg)

		// Version
		version := "-"
		if pkg.Version.Valid {
			version = fmt.Sprintf("`%s`", pkg.Version.String)
		}

		// Repo/URL badge
		urlBadge := Badge("N/A", "gray", "", "", "")
		if pkg.URL.Valid && pkg.URL.String != "" {
			linkText := "Link"
			urlBadge = Badge(linkText, "", "", pkg.URL.String, "")
		}

		// Category badge
		cateBadge := "-"
		if pkg.Category.Valid {
			cateBadge = CategoryShortcode(pkg.Category.String)
		}

		// RPM package name
		rpmDisplay := "-"
		if pkg.RpmPkg.Valid {
			rpmDisplay = fmt.Sprintf("`%s`", pkg.RpmPkg.String)
		}

		// DEB package name
		debDisplay := "-"
		if pkg.DebPkg.Valid {
			debDisplay = fmt.Sprintf("`%s`", pkg.DebPkg.String)
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			pkgLink, version, urlBadge, cateBadge, rpmDisplay, debDisplay))
	}

	return b.String()
}
