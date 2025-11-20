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

//============================================================================
// Unified List Generator
//============================================================================

// ListGenerator generates all list pages for the extension catalog.
// It consolidates generation for extensions, packages, categories, languages,
// licenses, and catalog pages.
type ListGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	DB        interface {
		Query(ctx context.Context, sql string, args ...interface{}) (interface{ Next() bool; Scan(dest ...interface{}) error; Close(); Err() error }, error)
	}
}

// NewListGenerator creates a new unified list generator
func NewListGenerator(cache *ExtensionCache, outputDir string) *ListGenerator {
	return &ListGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

//============================================================================
// Statistics and Data Types
//============================================================================

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

// CategoryPackages holds category and its packages (lead extensions)
type CategoryPackages struct {
	Category string
	Packages []struct {
		ID   int
		Name string
		Pkg  string
	}
}

// LicenseInfo contains metadata about a license
type LicenseInfo struct {
	URL         string
	Description string
	Order       int
}

type licenseItem struct {
	name  string
	exts  []*Extension
	order int
}

//============================================================================
// Extension List Generation
//============================================================================

// GenerateExtensionList generates extension list pages (ext.md and ext.zh.md)
func (g *ExtensionGenerator) GenerateExtensionList(locale, outputPath string) error {
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	categories, err := g.fetchCategoriesWithExtensionsSorted()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	// Filter out not-ready extensions
	var extensions []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			extensions = append(extensions, ext)
		}
	}

	content := g.generateExtensionListContent(stats, categories, extensions, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) generateExtensionListContent(stats *CatalogStats, categories []CategoryExtensions, extensions []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	totalExts := 0
	if len(stats.ExtensionStats) > 0 {
		totalExts = stats.ExtensionStats[0].Total
	}

	// Frontmatter
	if isZh {
		b.WriteString("---\ntitle: \"扩展清单\"\nweight: 10\nexcludeSearch: true\ncomments: false\n---\n\n")
	} else {
		b.WriteString("---\ntitle: \"Extensions\"\nweight: 10\nexcludeSearch: true\ncomments: false\n---\n\n")
	}

	// Statistics section
	if isZh {
		b.WriteString("## 统计\n\n")
	} else {
		b.WriteString("## Statistics\n\n")
	}
	b.WriteString(g.generateStatsTable(stats.ExtensionStats, stats.PGVersions, isZh, true))
	b.WriteString("\n")

	// Categories section
	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}
	b.WriteString(g.generateCategoriesTable(categories, isZh))
	b.WriteString("\n")

	// Full extension list
	if isZh {
		b.WriteString(fmt.Sprintf("## 扩展列表\n\n共有 %d 个可用的 PostgreSQL 扩展：\n\n", totalExts))
	} else {
		b.WriteString(fmt.Sprintf("## Extensions\n\nThere are %d available PostgreSQL extensions:\n\n", totalExts))
	}
	b.WriteString(g.generateFullExtensionTable(extensions, isZh))

	return b.String()
}

// GenerateExtensionIndexPages generates extension index pages (_index.md and _index.zh.md)
func (g *ExtensionGenerator) GenerateExtensionIndexPages(enPath, zhPath string) error {
	var extensions []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			extensions = append(extensions, ext)
		}
	}

	enContent := g.generateExtensionIndexContent(extensions, "en")
	if err := os.WriteFile(enPath, []byte(enContent), 0644); err != nil {
		return fmt.Errorf("failed to write English index: %w", err)
	}

	zhContent := g.generateExtensionIndexContent(extensions, "zh")
	if err := os.WriteFile(zhPath, []byte(zhContent), 0644); err != nil {
		return fmt.Errorf("failed to write Chinese index: %w", err)
	}

	return nil
}

func (g *ExtensionGenerator) generateExtensionIndexContent(extensions []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	extCount := len(extensions)

	if isZh {
		b.WriteString(fmt.Sprintf("---\ntitle: \"扩展列表\"\nbreadcrumbs: false\nexcludeSearch: true\ncomments: false\nweight: 900\n---\n\n共有 %d 个可用的 PostgreSQL 扩展：\n\n", extCount))
	} else {
		b.WriteString(fmt.Sprintf("---\ntitle: \"Extensions\"\nbreadcrumbs: false\nexcludeSearch: true\ncomments: false\nweight: 900\n---\n\nThere are %d available PostgreSQL extensions:\n\n", extCount))
	}

	b.WriteString(g.generateExtensionIndexTable(extensions, isZh))

	if isZh {
		b.WriteString("\n**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位")
	} else {
		b.WriteString("\n**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable")
	}

	return b.String()
}

func (g *ExtensionGenerator) generateExtensionIndexTable(extensions []*Extension, isZh bool) string {
	var b strings.Builder

	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	if isZh {
		b.WriteString("| 扩展 | PG 版本 | 属性 | 分类 | 描述 |\n")
		b.WriteString("|:----------|:------------|:---------:|:---------:|:--------------|\n")
	} else {
		b.WriteString("| Extension | PG Versions | Attribute | Category | Description |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------:|:--------------|\n")
	}

	for _, ext := range extensions {
		extLink := ExtShortcode(ext.Name, ext.Pkg)
		pgVersions := ParsePGVersions(ext.PgVer)
		pgBadges := PGVerShortcode(g.Cache.PGVersions, pgVersions)
		attrBadge := Badge(ext.GetAttributeBadge(), "blue", "", "", "")

		cateBadge := "-"
		if ext.Category.Valid {
			cateBadge = CategoryShortcode(ext.Category.String)
		}

		desc := getDescription(ext, isZh)
		if len(desc) > 100 {
			desc = desc[:100]
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			extLink, pgBadges, attrBadge, cateBadge, desc))
	}

	return b.String()
}

func (g *ExtensionGenerator) generateFullExtensionTable(extensions []*Extension, isZh bool) string {
	var b strings.Builder

	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	if isZh {
		b.WriteString("| 扩展 | PG 版本 | 属性 | 描述 |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------------|\n")
	} else {
		b.WriteString("| Extension | PG Versions | Attribute | Description |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------------|\n")
	}

	for _, ext := range extensions {
		extLink := fmt.Sprintf(`{{< ext "%s" >}}`, ext.Name)
		pgVersions := ParsePGVersions(ext.PgVer)
		pgBadges := PGVerShortcode(g.Cache.PGVersions, pgVersions)
		attrBadge := Badge(ext.GetAttributeBadge(), "blue", "", "", "")
		desc := getDescription(ext, isZh)

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
			extLink, pgBadges, attrBadge, desc))
	}

	return b.String()
}

func (g *ExtensionGenerator) generateCategoriesTable(categories []CategoryExtensions, isZh bool) string {
	var b strings.Builder

	if isZh {
		b.WriteString("| **类别** | **数** | **扩展** |\n")
		b.WriteString("|:------------:|:--------:|:---------------|\n")
	} else {
		b.WriteString("| **Category** | **Count** | **Extensions** |\n")
		b.WriteString("|:------------:|:---------:|:---------------|\n")
	}

	for _, cat := range categories {
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} | %d | ", categoryLower, len(cat.Extensions)))

		for i, ext := range cat.Extensions {
			b.WriteString(fmt.Sprintf(`{{< ext "%s" >}}`, ext.Name))
			if i < len(cat.Extensions)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString(" |\n")
	}

	return b.String()
}

//============================================================================
// Package List Generation
//============================================================================

// GeneratePackageList generates package list pages (pkg.md and pkg.zh.md)
func (g *ExtensionGenerator) GeneratePackageList(locale, outputPath string) error {
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	categories, err := g.fetchCategoriesWithPackagesSorted()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	var packages []*Extension
	for _, ext := range g.Cache.Extensions {
		if ext.State.Valid && ext.State.String == "not-ready" {
			continue
		}
		if ext.Lead {
			packages = append(packages, ext)
		}
	}

	sort.Slice(packages, func(i, j int) bool {
		return packages[i].ID < packages[j].ID
	})

	content := g.generatePackageListContent(stats, categories, packages, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) generatePackageListContent(stats *CatalogStats, categories []CategoryPackages, packages []*Extension, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	totalPkgs := 0
	if len(stats.PackageStats) > 0 {
		totalPkgs = stats.PackageStats[0].Total
	}

	if isZh {
		b.WriteString("---\ntitle: \"扩展包清单\"\nweight: 20\nexcludeSearch: true\ncomments: false\n---\n\n")
		b.WriteString("## 统计\n\n")
		b.WriteString("※ 一个扩展软件包可能同时包含多个 PG 扩展，因此按软件包统计的数量会少于扩展数量。\n\n")
	} else {
		b.WriteString("---\ntitle: \"Packages\"\nweight: 20\nexcludeSearch: true\ncomments: false\n---\n\n")
		b.WriteString("## Statistics\n\n")
		b.WriteString("※ One extension package may consist of multiple extensions\n\n")
	}

	b.WriteString(g.generateStatsTable(stats.PackageStats, stats.PGVersions, isZh, false))
	b.WriteString("\n")

	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}
	b.WriteString(g.generatePackageCategoriesTable(categories, isZh))
	b.WriteString("\n")

	if isZh {
		b.WriteString(fmt.Sprintf("## 扩展包列表\n\n共有 %d 个可用的 PostgreSQL 扩展包：\n\n", totalPkgs))
	} else {
		b.WriteString(fmt.Sprintf("## Packages\n\nThere are %d available PostgreSQL packages:\n\n", totalPkgs))
	}
	b.WriteString(g.generateFullPackageTable(packages, isZh))

	return b.String()
}

func (g *ExtensionGenerator) generatePackageCategoriesTable(categories []CategoryPackages, isZh bool) string {
	var b strings.Builder

	if isZh {
		b.WriteString("| **类别** | **数** | **扩展包** |\n")
		b.WriteString("|:------------:|:--------:|:---------------|\n")
	} else {
		b.WriteString("| **Category** | **Count** | **Packages** |\n")
		b.WriteString("|:------------:|:---------:|:---------------|\n")
	}

	for _, cat := range categories {
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} | %d | ", categoryLower, len(cat.Packages)))

		for i, pkg := range cat.Packages {
			b.WriteString(fmt.Sprintf(`{{< ext "%s" "%s" >}}`, pkg.Name, pkg.Pkg))
			if i < len(cat.Packages)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString(" |\n")
	}

	return b.String()
}

func (g *ExtensionGenerator) generateFullPackageTable(packages []*Extension, isZh bool) string {
	var b strings.Builder

	if isZh {
		b.WriteString("| 包 | 版本 | 仓库 | 分类 | RPM | DEB |\n")
		b.WriteString("|:---|:-----|:-----|:-----|:-----|:-----|\n")
	} else {
		b.WriteString("| Package | Version | Repo | Category | RPM | DEB |\n")
		b.WriteString("|:--------|:--------|:-----|:---------|:-----|:-----|\n")
	}

	for _, pkg := range packages {
		pkgLink := ExtShortcode(pkg.Name, pkg.Pkg)

		version := "-"
		if pkg.Version.Valid {
			version = fmt.Sprintf("`%s`", pkg.Version.String)
		}

		urlBadge := Badge("N/A", "gray", "", "", "")
		if pkg.URL.Valid && pkg.URL.String != "" {
			urlBadge = Badge("Link", "", "", pkg.URL.String, "")
		}

		cateBadge := "-"
		if pkg.Category.Valid {
			cateBadge = CategoryShortcode(pkg.Category.String)
		}

		rpmDisplay := "-"
		if pkg.RpmPkg.Valid {
			rpmDisplay = fmt.Sprintf("`%s`", pkg.RpmPkg.String)
		}

		debDisplay := "-"
		if pkg.DebPkg.Valid {
			debDisplay = fmt.Sprintf("`%s`", pkg.DebPkg.String)
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			pkgLink, version, urlBadge, cateBadge, rpmDisplay, debDisplay))
	}

	return b.String()
}

//============================================================================
// Catalog Page Generation
//============================================================================

// GenerateCatalogPage generates catalog list pages (_index.md and _index.zh.md)
func (g *ExtensionGenerator) GenerateCatalogPage(locale, outputPath string) error {
	stats, err := g.fetchCatalogStats()
	if err != nil {
		return fmt.Errorf("failed to fetch catalog stats: %w", err)
	}

	categories, err := g.fetchCategoriesWithExtensions()
	if err != nil {
		return fmt.Errorf("failed to fetch categories: %w", err)
	}

	content := g.generateCatalogContent(stats, categories, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) generateCatalogContent(stats *CatalogStats, categories []CategoryExtensions, locale string) string {
	isZh := locale == "zh"
	var b strings.Builder

	totalExts := 0
	totalPkgs := 0
	if len(stats.ExtensionStats) > 0 {
		totalExts = stats.ExtensionStats[0].Total
	}
	if len(stats.PackageStats) > 0 {
		totalPkgs = stats.PackageStats[0].Total
	}

	if isZh {
		b.WriteString("---\ntitle: \"扩展目录\"\nweight: 200\nexcludeSearch: true\ncomments: false\n---\n\n")
		b.WriteString(fmt.Sprintf("PostgreSQL 扩展目录包含了 **%d** 个扩展和 **%d** 个包。\n\n", totalExts, totalPkgs))
		b.WriteString("## 扩展统计\n\n")
	} else {
		b.WriteString("---\ntitle: \"Catalog\"\nweight: 200\nexcludeSearch: true\ncomments: false\n---\n\n")
		b.WriteString(fmt.Sprintf("The PostgreSQL Extension Catalog contains **%d** extensions and **%d** packages.\n\n", totalExts, totalPkgs))
		b.WriteString("## Extension Stat\n\n")
	}

	b.WriteString(g.generateStatsTable(stats.ExtensionStats, stats.PGVersions, isZh, true))
	b.WriteString("\n")

	if isZh {
		b.WriteString("## 扩展包统计\n\n")
		b.WriteString("※ 一个扩展软件包可能同时包含多个 PG 扩展，因此按软件包统计的数量会少于扩展数量。\n\n")
	} else {
		b.WriteString("## Package Stat\n\n")
		b.WriteString("※ One extension package may consist of multiple extension\n\n")
	}

	b.WriteString(g.generateStatsTable(stats.PackageStats, stats.PGVersions, isZh, false))
	b.WriteString("\n")

	if isZh {
		b.WriteString("## 分类\n\n")
	} else {
		b.WriteString("## Categories\n\n")
	}

	if isZh {
		b.WriteString("| **类别** | **扩展** |\n")
		b.WriteString("|:--------:|:---------|\n")
	} else {
		b.WriteString("| **Category** | **Extensions** |\n")
		b.WriteString("|:------------:|:---------------|\n")
	}

	for _, cat := range categories {
		categoryLower := strings.ToLower(cat.Category)
		b.WriteString(fmt.Sprintf("| {{< category %s >}} | ", categoryLower))

		for i, ext := range cat.Extensions {
			b.WriteString(fmt.Sprintf(`{{< ext "%s" >}}`, ext.Name))
			if i < len(cat.Extensions)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString(" |\n")
	}

	return b.String()
}

//============================================================================
// Category List Generation
//============================================================================

// GenerateCategoryList generates category-based extension list pages
func (g *ListGenerator) GenerateCategoryList(locale, outputPath string) error {
	isZh := locale == "zh"

	pkgMap := make(map[string]bool)
	for _, ext := range g.Cache.Extensions {
		pkgMap[ext.Pkg] = true
	}
	pkgCount := len(pkgMap)

	var b strings.Builder

	if isZh {
		b.WriteString(fmt.Sprintf("---\ntitle: \"按分类\"\nweight: 100\n---\n\nPostgreSQL 扩展（%d ext / %d pkg）归属 %d 个分类。\n\n", len(g.Cache.Extensions), pkgCount, len(g.Cache.Categories)))
	} else {
		b.WriteString(fmt.Sprintf("---\ntitle: \"By Category\"\nweight: 100\n---\n\nPostgreSQL Extensions (%d ext in %d pkg) categorized into %d categories.\n\n", len(g.Cache.Extensions), pkgCount, len(g.Cache.Categories)))
	}

	b.WriteString(`

| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------|
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |


`)

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		exts := g.Cache.CateExtMap[catKey]
		b.WriteString(g.generateCategorySection(cat, exts, isZh))
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func (g *ListGenerator) generateCategorySection(category *Category, extensions []*Extension, isZh bool) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n\n", category.Name))

	desc := ""
	if isZh && category.ZhDesc.Valid {
		desc = category.ZhDesc.String
	} else if category.EnDesc.Valid {
		desc = category.EnDesc.String
	}

	if desc != "" {
		prefix := fmt.Sprintf("%s:", category.Name)
		if strings.HasPrefix(desc, prefix) {
			desc = strings.TrimSpace(desc[len(prefix):])
		}
		b.WriteString(SanitizeText(desc))
		b.WriteString("\n\n")
	}

	if isZh {
		b.WriteString("| ID | 扩展/包 | 版本 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension / Package | Version | Description |\n")
		b.WriteString("|:---:|:---|:---|:---|\n")
	}

	if len(extensions) > 0 {
		sort.Slice(extensions, func(i, j int) bool {
			return extensions[i].ID < extensions[j].ID
		})

		for _, ext := range extensions {
			desc := getDescription(ext, isZh)
			if desc == "" {
				desc = "-"
			}

			version := "-"
			if ext.Version.Valid {
				version = ext.Version.String
			}

			extPkgCell := AliasShortcode(ext.Name, ext.Pkg)
			b.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n", ext.ID, extPkgCell, version, desc))
		}
	} else {
		emptyMsg := "此分类暂无扩展"
		if !isZh {
			emptyMsg = "No extensions available in this category"
		}
		b.WriteString(fmt.Sprintf("| - | - | - | %s |\n", emptyMsg))
	}

	b.WriteString("\n")
	return b.String()
}

//============================================================================
// Language List Generation
//============================================================================

// GenerateLanguageList generates language-based extension list pages
func (g *ListGenerator) GenerateLanguageList(locale, outputPath string) error {
	isZh := locale == "zh"

	langMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.Lang.Valid && ext.Lang.String != "" {
			lang := ext.Lang.String
			langMap[lang] = append(langMap[lang], ext)
		}
	}

	type langCount struct {
		lang  string
		count int
	}
	var langs []langCount
	for lang, exts := range langMap {
		langs = append(langs, langCount{lang, len(exts)})
	}
	sort.Slice(langs, func(i, j int) bool {
		if langs[i].count != langs[j].count {
			return langs[i].count > langs[j].count
		}
		return strings.ToLower(langs[i].lang) < strings.ToLower(langs[j].lang)
	})

	var b strings.Builder

	if isZh {
		b.WriteString("---\ntitle: \"按语言\"\ndescription: \"按实现语言组织的 PostgreSQL 扩展\"\nexcludeSearch: true\nweight: 200\n---\n\n")
	} else {
		b.WriteString("---\ntitle: \"By Language\"\ndescription: \"PostgreSQL extensions organized by implementation language\"\nexcludeSearch: true\nweight: 200\n---\n\n")
	}

	b.WriteString(`

| {{< language "c" >}}       | {{< language "c++" >}}       | {{< language "rust" >}}      | {{< language "java" >}}        | {{< language "python" >}}      | {{< language "sql" >}}         | {{< language "data" >}} |
|----------------------------|------------------------------|------------------------------|--------------------------------|--------------------------------|--------------------------------|-------------------------|

`)

	// Summary section
	if isZh {
		b.WriteString("## 概览\n\n")
		b.WriteString("| 语言 | 数量 | 描述 |\n")
		b.WriteString("|:-------:|:-----:|:--------------|\n")
	} else {
		b.WriteString("## Summary\n\n")
		b.WriteString("| Language | Count | Description |\n")
		b.WriteString("|:-------:|:-----:|:--------------|\n")
	}

	langDescriptions := getLanguageDescriptions(isZh)
	for _, lc := range langs {
		desc := langDescriptions[lc.lang]
		if desc == "" {
			if isZh {
				desc = fmt.Sprintf("使用 %s 编写的扩展", lc.lang)
			} else {
				desc = fmt.Sprintf("Extensions written in %s", lc.lang)
			}
		}
		b.WriteString(fmt.Sprintf("| %s | %d | %s |\n", LanguageShortcode(lc.lang), lc.count, desc))
	}
	b.WriteString("\n\n")

	for _, lc := range langs {
		b.WriteString(g.generateLanguageSection(lc.lang, langMap[lc.lang], isZh))
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func (g *ListGenerator) generateLanguageSection(lang string, extensions []*Extension, isZh bool) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n\n", lang))

	countText := fmt.Sprintf("%d Extensions", len(extensions))
	if isZh {
		countText = fmt.Sprintf("%d 个扩展", len(extensions))
	}
	b.WriteString(fmt.Sprintf("%s %s\n\n", LanguageShortcode(lang), Badge(countText, "gray", "", "", "cube")))

	langDesc := getLanguageDescriptions(isZh)[lang]
	if langDesc == "" {
		if isZh {
			langDesc = fmt.Sprintf("使用 %s 编写的扩展", lang)
		} else {
			langDesc = fmt.Sprintf("Extensions written in %s", lang)
		}
	}
	b.WriteString(langDesc)
	b.WriteString("\n\n")

	if isZh {
		b.WriteString("| ID | 扩展 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension | Description |\n")
		b.WriteString("|:---:|:---|:---|\n")
	}

	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	for _, ext := range extensions {
		extCell := AliasShortcode(ext.Name, ext.Pkg)
		desc := getDescription(ext, isZh)
		if desc == "" {
			if isZh {
				desc = "暂无描述"
			} else {
				desc = "No description"
			}
		}
		b.WriteString(fmt.Sprintf("| %d | %s | %s |\n", ext.ID, extCell, desc))
	}
	b.WriteString("\n")

	return b.String()
}

//============================================================================
// License List Generation
//============================================================================

// GenerateLicenseList generates license-based extension list pages
func (g *ListGenerator) GenerateLicenseList(locale, outputPath string) error {
	isZh := locale == "zh"

	licenseMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.License.Valid && ext.License.String != "" {
			license := ext.License.String
			licenseMap[license] = append(licenseMap[license], ext)
		}
	}

	var licenses []licenseItem
	for name, exts := range licenseMap {
		order := getLicenseOrder(name)
		licenses = append(licenses, licenseItem{name, exts, order})
	}

	sort.Slice(licenses, func(i, j int) bool {
		if len(licenses[i].exts) != len(licenses[j].exts) {
			return len(licenses[i].exts) > len(licenses[j].exts)
		}
		return licenses[i].order < licenses[j].order
	})

	var b strings.Builder

	if isZh {
		b.WriteString("---\ntitle: \"按许可证\"\ndescription: \"按开源许可证组织的 PostgreSQL 扩展\"\nweight: 300\n---\n\n按照所使用开源许可证，对 PostgreSQL 扩展进行分类。\n\n")
	} else {
		b.WriteString("---\ntitle: \"By License\"\ndescription: \"PostgreSQL extensions organized by open source license\"\nweight: 300\n---\n\nPostgreSQL extension categorized by license.\n\n")
	}

	b.WriteString(`

| {{< license "MIT" >}}      | {{< license "ISC" >}}        | {{< license "PostgreSQL" >}} | {{< license "BSD 0-Clause" >}} | {{< license "BSD 2-Clause" >}} | {{< license "BSD 3-Clause" >}} |
|:---------------------------|:-----------------------------|:-----------------------------|:-------------------------------|:-------------------------------|:-------------------------------|
| {{< license "Artistic" >}} | {{< license "Apache-2.0" >}} | {{< license "MPL-2.0" >}}    |                                |                                |                                |
| {{< license "GPL-2.0" >}}  | {{< license "GPL-3.0" >}}    | {{< license "LGPL-2.1" >}}   | {{< license "LGPL-3.0" >}}     | {{< license "AGPL-3.0" >}}     | {{< license "Timescale" >}}    |

`)

	// Summary section
	if isZh {
		b.WriteString("## 概览\n\n")
		b.WriteString("| 许可证 | 数量 | 参考 | 描述 |\n")
		b.WriteString("|:--------|:-----:|:-------:|:-----------|\n")
	} else {
		b.WriteString("## Summary\n\n")
		b.WriteString("| License | Count | Reference | Description |\n")
		b.WriteString("|:--------|:-----:|:-------:|:-----------|\n")
	}

	for _, lic := range licenses {
		info := getLicenseInfo(lic.name)
		refText := "License Text"
		if isZh {
			refText = "许可证文本"
		}
		b.WriteString(fmt.Sprintf("| %s | %d | [%s](%s) | %s |\n",
			LicenseShortcode(lic.name), len(lic.exts), refText, info.URL, info.Description))
	}

	b.WriteString("\n---------\n\n")

	for _, lic := range licenses {
		b.WriteString(g.generateLicenseSection(lic.name, lic.exts, isZh))
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func (g *ListGenerator) generateLicenseSection(license string, extensions []*Extension, isZh bool) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n\n", license))

	countText := fmt.Sprintf("%d Extensions", len(extensions))
	if isZh {
		countText = fmt.Sprintf("%d 个扩展", len(extensions))
	}
	info := getLicenseInfo(license)
	refText := "License Text"
	if isZh {
		refText = "许可证文本"
	}
	b.WriteString(fmt.Sprintf("\n\n| %s | %s  |\n|:----|:---|\n| %s | %s |\n\n",
		LicenseShortcode(license), Badge(countText, "gray", "", "", "cube"),
		Badge(refText, "gray", "", info.URL, "scale"), info.Description))

	if isZh {
		b.WriteString("| ID | 扩展 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension | Description |\n")
		b.WriteString("|:---:|:---|:---|\n")
	}

	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	for _, ext := range extensions {
		desc := getDescription(ext, isZh)
		b.WriteString(fmt.Sprintf("| %d | %s | %s |\n",
			ext.ID, AliasShortcode(ext.Name, ext.Pkg), desc))
	}

	b.WriteString("\n")
	return b.String()
}

//============================================================================
// Database Query Methods
//============================================================================

func (g *ExtensionGenerator) fetchCatalogStats() (*CatalogStats, error) {
	stats := &CatalogStats{}
	ctx := context.Background()

	pgVersions, err := g.getActivePGVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get active PG versions: %w", err)
	}
	stats.PGVersions = pgVersions

	pgColumnList := make([]string, 0)
	for _, pg := range pgVersions {
		pgColumnList = append(pgColumnList, fmt.Sprintf("pg%d", pg))
	}

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
		row := StatsRow{PGCounts: make(map[int]int)}
		scanDest := []interface{}{&row.ID, &row.Title, &row.Total, &row.PGDG, &row.PIGSTY, &row.CONTRIB, &row.MISS}

		pgCounts := make([]int, len(pgVersions))
		for i := range pgVersions {
			scanDest = append(scanDest, &pgCounts[i])
		}

		if err := rows.Scan(scanDest...); err != nil {
			return nil, fmt.Errorf("failed to scan summary row: %w", err)
		}

		for i, pg := range pgVersions {
			row.PGCounts[pg] = pgCounts[i]
		}
		allRows = append(allRows, row)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading summary rows: %w", err)
	}

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

	result := make([]CategoryExtensions, 0, len(categories))
	for _, cat := range categories {
		result = append(result, *categoryMap[cat])
	}

	return result, nil
}

func (g *ExtensionGenerator) fetchCategoriesWithExtensionsSorted() ([]CategoryExtensions, error) {
	ctx := context.Background()

	categoryQuery := `SELECT name, id FROM pgext.category ORDER BY id`
	categoryRows, err := g.DB.Query(ctx, categoryQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query category order: %w", err)
	}
	defer categoryRows.Close()

	var categoryOrder []string
	for categoryRows.Next() {
		var name string
		var id int
		if err := categoryRows.Scan(&name, &id); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categoryOrder = append(categoryOrder, name)
	}

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

	result := make([]CategoryExtensions, 0, len(categoryOrder))
	for _, cat := range categoryOrder {
		if catData, exists := categoryMap[cat]; exists {
			result = append(result, *catData)
		}
	}

	return result, nil
}

func (g *ExtensionGenerator) fetchCategoriesWithPackagesSorted() ([]CategoryPackages, error) {
	ctx := context.Background()

	categoryQuery := `SELECT name, id FROM pgext.category ORDER BY id`
	categoryRows, err := g.DB.Query(ctx, categoryQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query category order: %w", err)
	}
	defer categoryRows.Close()

	var categoryOrder []string
	for categoryRows.Next() {
		var name string
		var id int
		if err := categoryRows.Scan(&name, &id); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categoryOrder = append(categoryOrder, name)
	}

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

	result := make([]CategoryPackages, 0, len(categoryOrder))
	for _, cat := range categoryOrder {
		if catData, exists := categoryMap[cat]; exists {
			result = append(result, *catData)
		}
	}

	return result, nil
}

func (g *ExtensionGenerator) generateStatsTable(rows []StatsRow, pgVersions []int, isZh, isExtension bool) string {
	var b strings.Builder

	b.WriteString("|")
	if isZh {
		b.WriteString(" **分类** |")
	} else {
		b.WriteString(" **Category** |")
	}
	b.WriteString(" **All** | **PGDG** | **PIGSTY** | **CONTRIB** | **MISS** |")

	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")

	b.WriteString("|:---------|--------:|--------:|----------:|-----------:|--------:|")
	for range pgVersions {
		b.WriteString("--------:|")
	}
	b.WriteString("\n")

	for _, row := range rows {
		label := strings.Split(row.Title, " ")[0]
		b.WriteString(fmt.Sprintf("| **%s** | %d | %d | %d | %d | %d |",
			label, row.Total, row.PGDG, row.PIGSTY, row.CONTRIB, row.MISS))

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

//============================================================================
// Helper Functions
//============================================================================

func getDescription(ext *Extension, isZh bool) string {
	if isZh {
		if ext.ZhDesc.Valid {
			return SanitizeText(ext.ZhDesc.String)
		} else if ext.EnDesc.Valid {
			return SanitizeText(ext.EnDesc.String)
		}
	} else {
		if ext.EnDesc.Valid {
			return SanitizeText(ext.EnDesc.String)
		}
	}
	return ""
}

func getLanguageDescriptions(isZh bool) map[string]string {
	if isZh {
		return map[string]string{
			"C":      "传统的 PostgreSQL 扩展开发语言",
			"C++":    "使用 C++ 特性和库的扩展",
			"Rust":   "使用 pgrx 框架用 Rust 编写的扩展",
			"Python": "使用 Python 编写的扩展",
			"SQL":    "纯 SQL 扩展和函数",
			"Java":   "在 JVM 上运行的扩展",
			"Data":   "仅包含数据的扩展",
		}
	}
	return map[string]string{
		"C":      "The traditional PostgreSQL extension language",
		"C++":    "Extensions leveraging C++ features and libraries",
		"Rust":   "Extensions written in Rust with the pgrx framework",
		"Python": "Extensions written in Python",
		"SQL":    "Pure SQL extensions and functions",
		"Java":   "Extensions running on JVM",
		"Data":   "Data-only extensions",
	}
}

func getLicenseInfo(license string) LicenseInfo {
	licenses := map[string]LicenseInfo{
		"PostgreSQL":   {"https://opensource.org/licenses/postgresql", "Very liberal license based on the BSD license, allowing almost unlimited freedom.", 1},
		"Apache-2.0":   {"https://opensource.org/licenses/Apache-2.0", "Permissive license with patent protection and attribution requirements.", 2},
		"MIT":          {"https://opensource.org/licenses/MIT", "A permissive license that allows commercial use, modification, and private use.", 3},
		"BSD 3-Clause": {"https://opensource.org/license/bsd-3-clause", "Permissive license with attribution and endorsement restriction clauses.", 4},
		"BSD 2-Clause": {"https://opensource.org/license/bsd-2-clause", "Permissive license requiring attribution but allowing commercial use.", 5},
		"GPL-2.0":      {"https://opensource.org/licenses/GPL-2.0", "Strong copyleft license requiring derivative works to be open source.", 6},
		"GPL-3.0":      {"https://opensource.org/licenses/GPL-3.0", "Strong copyleft license with additional patent and hardware restrictions.", 7},
		"AGPL-3.0":     {"https://opensource.org/licenses/AGPL-3.0", "Network copyleft license extending GPL to cover network-distributed software.", 8},
		"ISC":          {"https://opensource.org/licenses/ISC", "A permissive license similar to MIT, allowing commercial use and modification.", 9},
		"Artistic":     {"https://opensource.org/license/artistic-2-0", "Copyleft license allowing modification with certain distribution requirements.", 10},
		"Timescale":    {"https://www.timescale.com/legal/licenses", "Proprietary license with restrictions on commercial use and distribution.", 11},
		"BSD 0-Clause": {"https://opensource.org/license/0bsd", "Public domain equivalent license with no restrictions on use.", 12},
		"LGPL-3.0":     {"https://opensource.org/licenses/LGPL-3.0", "Weak copyleft license with additional patent and hardware provisions.", 13},
		"MPL-2.0":      {"https://opensource.org/licenses/MPL-2.0", "Weak copyleft license allowing proprietary combinations with file-level copyleft.", 14},
		"LGPL-2.1":     {"https://opensource.org/licenses/LGPL-2.1", "Weak copyleft license allowing proprietary applications to link dynamically.", 15},
	}

	if info, ok := licenses[license]; ok {
		return info
	}
	return LicenseInfo{"#", "Unknown license", 999}
}

func getLicenseOrder(license string) int {
	return getLicenseInfo(license).Order
}
