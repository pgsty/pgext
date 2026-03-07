/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC Category Generator - generates per-category extension index pages for pigsty.cc (Chinese only)
Each category page lists all extensions with unified tables showing metadata and OS availability.
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// OSCodeEntry represents a unique os_code (e.g., el8, el9, d12) with its two arch variants
type OSCodeEntry struct {
	Code   string // e.g., "el8"
	X86    string // full OS name for x86_64, e.g., "el8.x86_64"
	ARM    string // full OS name for aarch64, e.g., "el8.aarch64"
	HasX86 bool
	HasARM bool
}

// CCCateGenerator generates per-category extension index pages for pigsty.cc
type CCCateGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	OSCodes   []OSCodeEntry
}

// NewCCCateGenerator creates a new CC category generator
func NewCCCateGenerator(cache *ExtensionCache, outputDir string) *CCCateGenerator {
	g := &CCCateGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
	g.buildOSCodes()
	return g
}

// buildOSCodes builds the ordered OS code list from cache OSVersions
func (g *CCCateGenerator) buildOSCodes() {
	codeMap := make(map[string]*OSCodeEntry)
	var codeOrder []string

	for _, osv := range g.Cache.OSVersions {
		if !osv.Active {
			continue
		}
		code := osv.OS
		if idx := strings.Index(osv.OS, "."); idx > 0 {
			code = osv.OS[:idx]
		}
		entry, exists := codeMap[code]
		if !exists {
			entry = &OSCodeEntry{Code: code}
			codeMap[code] = entry
			codeOrder = append(codeOrder, code)
		}
		if osv.OSArch == "x86_64" {
			entry.X86 = osv.OS
			entry.HasX86 = true
		} else if osv.OSArch == "aarch64" {
			entry.ARM = osv.OS
			entry.HasARM = true
		}
	}

	g.OSCodes = make([]OSCodeEntry, 0, len(codeOrder))
	for _, code := range codeOrder {
		g.OSCodes = append(g.OSCodes, *codeMap[code])
	}
}

// CategoryWeight returns the frontmatter weight for a category (800 + category ID)
func CategoryWeight(catID int) int {
	return 800 + catID
}

// CategoryZhName returns the Chinese display name for a category
func CategoryZhName(catName string) string {
	names := map[string]string{
		"TIME": "时序", "GIS": "地理", "RAG": "向量", "FTS": "全文检索",
		"OLAP": "数据分析", "FEAT": "特性", "LANG": "编程语言", "TYPE": "数据类型",
		"UTIL": "工具", "FUNC": "数学函数", "ADMIN": "管理", "STAT": "监控统计",
		"SEC": "安全", "FDW": "外部数据", "SIM": "兼容模拟", "ETL": "数据集成",
	}
	if n, ok := names[strings.ToUpper(catName)]; ok {
		return n
	}
	return catName
}

// GenerateCategoryPage generates a single category page
func (g *CCCateGenerator) GenerateCategoryPage(ctx context.Context, cat *Category) error {
	catKey := strings.ToUpper(cat.Name)
	allExts := g.Cache.CateExtMap[catKey]
	if len(allExts) == 0 {
		return nil
	}

	// Filter out not-ready extensions and sort by ID
	var exts []*Extension
	for _, ext := range allExts {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			exts = append(exts, ext)
		}
	}
	sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })

	if len(exts) == 0 {
		return nil
	}

	// Collect unique package names for availability loading (skip contrib and kernel-fork)
	pkgSet := make(map[string]bool)
	for _, ext := range exts {
		hasKernel := ext.Extra != nil && ext.Extra["kernel"] != nil
		if !ext.Contrib && !hasKernel {
			pkgSet[ext.Pkg] = true
		}
	}
	var pkgNames []string
	for name := range pkgSet {
		pkgNames = append(pkgNames, name)
	}

	pkgAvail, err := g.loadAvailability(ctx, pkgNames)
	if err != nil {
		logrus.Warnf("Failed to load availability for category %s: %v", catKey, err)
		pkgAvail = make(map[string]map[string]*PkgInfo)
	}

	content := g.generateContent(cat, exts, pkgAvail)

	cateDir := filepath.Join(g.OutputDir, "cate")
	if err := os.MkdirAll(cateDir, 0755); err != nil {
		return fmt.Errorf("failed to create category directory: %w", err)
	}
	outputPath := filepath.Join(cateDir, strings.ToLower(cat.Name)+".md")
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Count unique packages for log
	pkgCount := 0
	seen := make(map[string]bool)
	for _, ext := range exts {
		if ext.Lead && !seen[ext.Pkg] {
			pkgCount++
			seen[ext.Pkg] = true
		}
	}
	logrus.Infof("Generated category page for %s with %d packages", catKey, pkgCount)
	return nil
}

// loadAvailability loads pkg availability for given package names
func (g *CCCateGenerator) loadAvailability(ctx context.Context, pkgNames []string) (map[string]map[string]*PkgInfo, error) {
	if len(pkgNames) == 0 {
		return make(map[string]map[string]*PkgInfo), nil
	}

	placeholders := make([]string, len(pkgNames))
	args := make([]interface{}, len(pkgNames))
	for i, name := range pkgNames {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = name
	}

	query := fmt.Sprintf(`
		SELECT pg, os, pkg, state, version
		FROM pgext.pkg
		WHERE pkg IN (%s)
		ORDER BY pkg, os, pg DESC
	`, strings.Join(placeholders, ","))

	rows, err := QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]map[string]*PkgInfo)
	for rows.Next() {
		var pkg PkgInfo
		if err := rows.Scan(&pkg.PG, &pkg.OS, &pkg.Pkg, &pkg.State, &pkg.Version); err != nil {
			return nil, err
		}
		key := fmt.Sprintf("%d-%s", pkg.PG, pkg.OS)
		if result[pkg.Pkg] == nil {
			result[pkg.Pkg] = make(map[string]*PkgInfo)
		}
		if _, exists := result[pkg.Pkg][key]; !exists {
			p := pkg
			result[pkg.Pkg][key] = &p
		}
	}
	return result, rows.Err()
}

// generateContent generates the full markdown content for a category page
func (g *CCCateGenerator) generateContent(cat *Category, exts []*Extension, pkgAvail map[string]map[string]*PkgInfo) string {
	var b strings.Builder

	b.WriteString(g.generateFrontmatter(cat))
	b.WriteString(g.generateOverview(exts))
	b.WriteString(g.generateExtensionTable(exts))

	for _, ext := range exts {
		avail := pkgAvail[ext.Pkg]
		if avail == nil {
			avail = make(map[string]*PkgInfo)
		}
		b.WriteString(g.generateExtCard(ext, avail))
	}

	return b.String()
}

// generateFrontmatter generates the YAML frontmatter with icon
func (g *CCCateGenerator) generateFrontmatter(cat *Category) string {
	catUpper := strings.ToUpper(cat.Name)
	zhName := CategoryZhName(catUpper)
	desc := cat.Name
	if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
		desc = cat.ZhDesc.String
	}

	icon := ""
	if cat.Icon1.Valid && cat.Icon1.String != "" {
		icon = fmt.Sprintf("icon: %s\n", cat.Icon1.String)
	}

	return fmt.Sprintf(`---
title: "%s - %s"
linkTitle: "%s"
description: "%s"
weight: %d
%s---

`, catUpper, zhName, catUpper, desc, CategoryWeight(cat.ID), icon)
}

// generateOverview generates the overview heading with counts
func (g *CCCateGenerator) generateOverview(exts []*Extension) string {
	pkgSet := make(map[string]bool)
	for _, ext := range exts {
		if ext.Lead {
			pkgSet[ext.Pkg] = true
		}
	}
	return fmt.Sprintf("## 扩展列表\n\n共有 **%d** 个扩展，位于 **%d** 个扩展包中。\n\n", len(exts), len(pkgSet))
}

// generateExtensionTable generates the overview extension table at the top
func (g *CCCateGenerator) generateExtensionTable(exts []*Extension) string {
	return CCExtensionTable(exts)
}

// generateExtCard generates a unified table section for a single extension
func (g *CCCateGenerator) generateExtCard(ext *Extension, avail map[string]*PkgInfo) string {
	var b strings.Builder

	// Description
	desc := ext.Name
	if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
		desc = ext.ZhDesc.String
	} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
		desc = ext.EnDesc.String
	}

	// H2 heading: extension name
	b.WriteString(fmt.Sprintf("\n---------\n\n## %s {#%s}\n\n", ext.Name, ext.Name))

	// Intro line: [**`pkg`**](url) - `version` : description
	pkgLink := fmt.Sprintf("**`%s`**", ext.Pkg)
	if ext.URL.Valid && ext.URL.String != "" {
		pkgLink = fmt.Sprintf("[**`%s`**](%s)", ext.Pkg, ext.URL.String)
	}
	version := "-"
	if ext.Version.Valid {
		version = ext.Version.String
	}
	b.WriteString(fmt.Sprintf("%s - `%s` : %s\n\n", pkgLink, version, desc))

	// Build the 7 left-side metadata rows
	leftRows := g.buildLeftRows(ext)

	pgVersions := g.Cache.PGVersions
	pgCount := len(pgVersions)

	// For contrib or kernel-fork extensions, build synthetic availability from PgVer
	hasKernel := ext.Extra != nil && ext.Extra["kernel"] != nil
	usePgVer := ext.Contrib || hasKernel
	var pgVerSet map[int]bool
	if usePgVer {
		pgVerSet = make(map[int]bool)
		for _, v := range ext.PgVer {
			v = strings.TrimSuffix(v, "+")
			var pg int
			if _, err := fmt.Sscanf(v, "%d", &pg); err == nil {
				pgVerSet[pg] = true
			}
		}
	}

	// Header row with centered x86_64 / aarch64 spanning across PG columns
	// Use HTML colspan-like centering via merged empty headers
	b.WriteString("| **条目** | **属性** | **OS** |")
	// x86_64 header: first cell has text, rest empty
	for i := 0; i < pgCount; i++ {
		if i == pgCount/2 {
			b.WriteString(" **x86_64** |")
		} else {
			b.WriteString(" |")
		}
	}
	// aarch64 header: first cell has text, rest empty
	for i := 0; i < pgCount; i++ {
		if i == pgCount/2 {
			b.WriteString(" **aarch64** |")
		} else {
			b.WriteString(" |")
		}
	}
	b.WriteString("\n")

	// Separator
	b.WriteString("|:---:|:---|:---:|")
	for i := 0; i < pgCount*2; i++ {
		b.WriteString(":---:|")
	}
	b.WriteString("\n")

	// 7 data rows: left metadata + right OS availability
	numRows := len(leftRows)
	if len(g.OSCodes) < numRows {
		numRows = len(g.OSCodes)
	}

	for i := 0; i < numRows; i++ {
		left := leftRows[i]
		osc := g.OSCodes[i]

		b.WriteString(fmt.Sprintf("| **%s** | %s | **%s** |", left.Label, left.Value, osc.Code))

		if usePgVer {
			// Contrib/kernel-fork: use PgVer for all OS uniformly
			for _, pg := range pgVersions {
				if pgVerSet[pg] {
					b.WriteString(fmt.Sprintf(` <span class="ext-pgver ext-pgver--ok">%d</span> |`, pg))
				} else {
					b.WriteString(" |")
				}
			}
			for _, pg := range pgVersions {
				if pgVerSet[pg] {
					b.WriteString(fmt.Sprintf(` <span class="ext-pgver ext-pgver--ok">%d</span> |`, pg))
				} else {
					b.WriteString(" |")
				}
			}
		} else {
			// Normal: use actual package availability
			for _, pg := range pgVersions {
				badge := g.pgAvailBadge(avail, pg, osc.X86)
				b.WriteString(fmt.Sprintf(" %s |", badge))
			}
			for _, pg := range pgVersions {
				badge := g.pgAvailBadge(avail, pg, osc.ARM)
				b.WriteString(fmt.Sprintf(" %s |", badge))
			}
		}
		b.WriteString("\n")
	}

	// Extra OS rows beyond the 7 left-side rows
	for i := numRows; i < len(g.OSCodes); i++ {
		osc := g.OSCodes[i]
		b.WriteString(fmt.Sprintf("| | | **%s** |", osc.Code))
		if usePgVer {
			for _, pg := range pgVersions {
				if pgVerSet[pg] {
					b.WriteString(fmt.Sprintf(` <span class="ext-pgver ext-pgver--ok">%d</span> |`, pg))
				} else {
					b.WriteString(" |")
				}
			}
			for _, pg := range pgVersions {
				if pgVerSet[pg] {
					b.WriteString(fmt.Sprintf(` <span class="ext-pgver ext-pgver--ok">%d</span> |`, pg))
				} else {
					b.WriteString(" |")
				}
			}
		} else {
			for _, pg := range pgVersions {
				badge := g.pgAvailBadge(avail, pg, osc.X86)
				b.WriteString(fmt.Sprintf(" %s |", badge))
			}
			for _, pg := range pgVersions {
				badge := g.pgAvailBadge(avail, pg, osc.ARM)
				b.WriteString(fmt.Sprintf(" %s |", badge))
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("{.ext-table .ext-table--cate}\n\n")
	return b.String()
}

// leftRow represents a label-value pair for the left side of the unified table
type leftRow struct {
	Label string
	Value string
}

// buildLeftRows builds the 7 left-side metadata rows for an extension
func (g *CCCateGenerator) buildLeftRows(ext *Extension) []leftRow {
	// Row 1: 扩展名
	extLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)

	// Row 2: 扩展包
	pkgLink := fmt.Sprintf("`%s`", ext.Pkg)
	if ext.URL.Valid && ext.URL.String != "" {
		pkgLink = fmt.Sprintf("[`%s`](%s)", ext.Pkg, ext.URL.String)
	}

	// Row 3: RPM 包名
	rpmPkg := "-"
	if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
		rpmPkg = fmt.Sprintf("`%s`", ext.RpmPkg.String)
	}

	// Row 4: DEB 包名
	debPkg := "-"
	if ext.DebPkg.Valid && ext.DebPkg.String != "" {
		debPkg = fmt.Sprintf("`%s`", ext.DebPkg.String)
	}

	// Row 5: 编程语言
	lang := "-"
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}

	// Row 6: 来源仓库
	repo := "-"
	if ext.Repo.Valid && ext.Repo.String != "" {
		repo = CCRepoBadge(ext.Repo.String)
	}

	// Row 7: 开源协议
	license := "-"
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}

	return []leftRow{
		{"扩展名", extLink},
		{"扩展包", pkgLink},
		{"RPM", rpmPkg},
		{"DEB", debPkg},
		{"语言", lang},
		{"仓库", repo},
		{"协议", license},
	}
}

// pgAvailBadge generates a PG version availability badge for the category page
func (g *CCCateGenerator) pgAvailBadge(avail map[string]*PkgInfo, pg int, osName string) string {
	if osName == "" {
		return ""
	}
	key := fmt.Sprintf("%d-%s", pg, osName)
	pkg, exists := avail[key]
	if !exists {
		return ""
	}
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}
	if state == "AVAIL" {
		return fmt.Sprintf(`<span class="ext-pgver ext-pgver--ok">%d</span>`, pg)
	}
	return ""
}

// GenerateCategoryIndexPage generates the _index.md for the cate/ directory
func (g *CCCateGenerator) GenerateCategoryIndexPage() error {
	cateDir := filepath.Join(g.OutputDir, "cate")
	if err := os.MkdirAll(cateDir, 0755); err != nil {
		return fmt.Errorf("failed to create category directory: %w", err)
	}

	var b strings.Builder
	b.WriteString(`---
title: "扩展分类"
linkTitle: "扩展分类"
description: "将 PostgreSQL 扩展按照功能分为 16 个大类别，方便用户按功能浏览扩展包索引。"
weight: 810
sidebar_expanded: true
---

`)

	b.WriteString("| **分类** | **扩展数** | **包数** | **描述** |\n")
	b.WriteString("|:--------:|:-------:|:------:|:-------|\n")

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		allExts := g.Cache.CateExtMap[catKey]

		extCount := 0
		pkgCount := 0
		for _, ext := range allExts {
			if ext.State.Valid && ext.State.String == "not-ready" {
				continue
			}
			extCount++
			if ext.Lead {
				pkgCount++
			}
		}

		desc := cat.Name
		if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
			desc = cat.ZhDesc.String
		}

		lower := strings.ToLower(cat.Name)
		badge := fmt.Sprintf(`<a class="ext-badge ext-badge--cate %s" href="/ext/cate/%s">%s</a>`,
			lower, lower, catKey)

		b.WriteString(fmt.Sprintf("| %s | %d | %d | %s |\n",
			badge, extCount, pkgCount, desc))
	}

	b.WriteString("{.ext-table}\n\n")

	outputPath := filepath.Join(cateDir, "_index.md")
	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateAllCategoryPages generates all 16 category pages into cate/ subdirectory
func (g *CCCateGenerator) GenerateAllCategoryPages(ctx context.Context) error {
	if err := g.GenerateCategoryIndexPage(); err != nil {
		return fmt.Errorf("failed to generate category index: %w", err)
	}

	for _, cat := range g.Cache.Categories {
		if err := g.GenerateCategoryPage(ctx, cat); err != nil {
			logrus.Errorf("Failed to generate category page for %s: %v", cat.Name, err)
		}
	}

	logrus.Infof("Generated %d category pages", len(g.Cache.Categories))
	return nil
}
