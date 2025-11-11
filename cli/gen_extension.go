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

// GenerateExtensionList generates extension list pages (ext.md and ext.zh.md)
// This combines statistics from catalog, category organization, and full extension index
func (g *ExtensionGenerator) GenerateExtensionList(locale, outputPath string) error {
	// Fetch statistics (reuse from catalog generation)
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	// Fetch categories with extensions (sorted by category ID)
	categories, err := g.fetchCategoriesWithExtensionsSorted()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	// Filter out extensions that are not ready
	var extensions []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			extensions = append(extensions, ext)
		}
	}

	content := g.generateExtensionListContent(stats, categories, extensions, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// fetchCategoriesWithExtensionsSorted fetches categories sorted by their ID in the category table
func (g *ExtensionGenerator) fetchCategoriesWithExtensionsSorted() ([]CategoryExtensions, error) {
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

	// Then get extensions grouped by category
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
		}

		categoryMap[category].Extensions = append(categoryMap[category].Extensions, ext)
	}

	// Convert map to slice in the correct order
	result := make([]CategoryExtensions, 0, len(categoryOrder))
	for _, cat := range categoryOrder {
		if catData, exists := categoryMap[cat]; exists {
			result = append(result, *catData)
		}
	}

	return result, nil
}

func (g *ExtensionGenerator) generateExtensionListContent(stats *CatalogStats, categories []CategoryExtensions, extensions []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	// Calculate total extensions
	totalExts := 0
	if len(stats.ExtensionStats) > 0 {
		totalExts = stats.ExtensionStats[0].Total // ALL row
	}

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "扩展清单"
weight: 10
excludeSearch: true
comments: false
---

`)
	} else {
		b.WriteString(`---
title: "Extensions"
weight: 10
excludeSearch: true
comments: false
---

`)
	}

	// Statistics section - Extract only Extension Stats part
	if isZh {
		b.WriteString("## 统计\n\n")
	} else {
		b.WriteString("## Statistics\n\n")
	}
	b.WriteString(g.generateExtensionStatsTable(stats.ExtensionStats, stats.PGVersions, isZh))
	b.WriteString("\n")

	// Categories section with extensions
	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}
	b.WriteString(g.generateCategoriesTable(categories, isZh))
	b.WriteString("\n")

	// Full extension list section
	if isZh {
		b.WriteString(fmt.Sprintf("## 扩展列表\n\n共有 %d 个可用的 PostgreSQL 扩展：\n\n", totalExts))
	} else {
		b.WriteString(fmt.Sprintf("## Extensions\n\nThere are %d available PostgreSQL extensions:\n\n", totalExts))
	}
	b.WriteString(g.generateFullExtensionTable(extensions, isZh))

	return b.String()
}

// generateExtensionStatsTable generates the statistics table for extensions only
func (g *ExtensionGenerator) generateExtensionStatsTable(rows []StatsRow, pgVersions []int, isZh bool) string {
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

// generateCategoriesTable generates the categories table with extensions and count
func (g *ExtensionGenerator) generateCategoriesTable(categories []CategoryExtensions, isZh bool) string {
	var b strings.Builder

	// Table header
	if isZh {
		b.WriteString("| **类别** | **数** | **扩展** |\n")
		b.WriteString("|:------------:|:--------:|:---------------|\n")
	} else {
		b.WriteString("| **Category** | **Count** | **Extensions** |\n")
		b.WriteString("|:------------:|:---------:|:---------------|\n")
	}

	// Table rows - one row per category
	for _, cat := range categories {
		// Category column with shortcode
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} |", categoryLower))

		// Count column
		b.WriteString(fmt.Sprintf(" %d |", len(cat.Extensions)))

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

// generateFullExtensionTable generates the full extension list table
func (g *ExtensionGenerator) generateFullExtensionTable(extensions []*Extension, isZh bool) string {
	var b strings.Builder

	// Sort extensions by ID for consistency
	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	// Table header
	if isZh {
		b.WriteString("| 扩展 | PG 版本 | 属性 | 描述 |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------------|\n")
	} else {
		b.WriteString("| Extension | PG Versions | Attribute | Description |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------------|\n")
	}

	// Table rows
	for _, ext := range extensions {
		// Extension link - use single parameter (just extension name)
		extLink := fmt.Sprintf(`{{< ext "%s" >}}`, ext.Name)

		// Parse PG versions
		pgVersions := ParsePGVersions(ext.PgVer)
		pgBadges := PGVerShortcode(g.Cache.PGVersions, pgVersions)

		// Attribute badge
		attrBadge := Badge(ext.GetAttributeBadge(), "blue", "", "", "")

		// Description
		desc := ""
		if isZh {
			if ext.ZhDesc.Valid {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
		} else {
			if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
			extLink, pgBadges, attrBadge, desc))
	}

	return b.String()
}

// GenerateExtensionIndexPages generates extension index pages (_index.md and _index.zh.md)
// This generates only the extension table with category column, no package table
func (g *ExtensionGenerator) GenerateExtensionIndexPages(enPath, zhPath string) error {
	// Filter out extensions that are not ready
	var extensions []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			extensions = append(extensions, ext)
		}
	}

	// Generate English version
	enContent := g.generateExtensionIndexContent(extensions, "en")
	if err := os.WriteFile(enPath, []byte(enContent), 0644); err != nil {
		return fmt.Errorf("failed to write English index: %w", err)
	}

	// Generate Chinese version
	zhContent := g.generateExtensionIndexContent(extensions, "zh")
	if err := os.WriteFile(zhPath, []byte(zhContent), 0644); err != nil {
		return fmt.Errorf("failed to write Chinese index: %w", err)
	}

	return nil
}

// generateExtensionIndexContent generates the index page content with only extension table
func (g *ExtensionGenerator) generateExtensionIndexContent(extensions []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	extCount := len(extensions)

	// Frontmatter
	if isZh {
		b.WriteString(fmt.Sprintf(`---
title: "扩展列表"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

共有 %d 个可用的 PostgreSQL 扩展：

`, extCount))
	} else {
		b.WriteString(fmt.Sprintf(`---
title: "Extensions"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

There are %d available PostgreSQL extensions:

`, extCount))
	}

	// Generate extensions table with category column
	b.WriteString(g.generateExtensionIndexTable(extensions, isZh))

	// Attribute legend
	if isZh {
		b.WriteString("\n**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位")
	} else {
		b.WriteString("\n**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable")
	}

	return b.String()
}

// generateExtensionIndexTable generates the extension table for index pages with category column
func (g *ExtensionGenerator) generateExtensionIndexTable(extensions []*Extension, isZh bool) string {
	var b strings.Builder

	// Sort extensions by ID for consistency
	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	// Table header with category column
	if isZh {
		b.WriteString("| 扩展 | PG 版本 | 属性 | 分类 | 描述 |\n")
		b.WriteString("|:----------|:------------|:---------:|:---------:|:--------------|\n")
	} else {
		b.WriteString("| Extension | PG Versions | Attribute | Category | Description |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------:|:--------------|\n")
	}

	// Table rows
	for _, ext := range extensions {
		// Extension link - use two-parameter shortcode format
		extLink := ExtShortcode(ext.Name, ext.Pkg)

		// Parse PG versions
		pgVersions := ParsePGVersions(ext.PgVer)
		pgBadges := PGVerShortcode(g.Cache.PGVersions, pgVersions)

		// Attribute badge
		attrBadge := Badge(ext.GetAttributeBadge(), "blue", "", "", "")

		// Category badge
		cateBadge := "-"
		if ext.Category.Valid {
			cateBadge = CategoryShortcode(ext.Category.String)
		}

		// Description
		desc := ""
		if isZh {
			if ext.ZhDesc.Valid {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
		} else {
			if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
		}

		// Truncate description to 100 chars
		if len(desc) > 100 {
			desc = desc[:100]
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			extLink, pgBadges, attrBadge, cateBadge, desc))
	}

	return b.String()
}
