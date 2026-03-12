/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

IO Category Generator - generates per-category extension index pages for pigsty.io (English only)
Each category page lists all extensions with unified tables showing metadata and OS availability.
*/
package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// IOCateGenerator generates per-category extension index pages for pigsty.io
type IOCateGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	OSCodes   []OSCodeEntry
}

// NewIOCateGenerator creates a new IO category generator
func NewIOCateGenerator(cache *ExtensionCache, outputDir string) *IOCateGenerator {
	g := &IOCateGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
	g.buildOSCodes()
	return g
}

// buildOSCodes builds the ordered OS code list from cache OSVersions
func (g *IOCateGenerator) buildOSCodes() {
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

// GenerateCategoryPage generates a single category page
func (g *IOCateGenerator) GenerateCategoryPage(ctx context.Context, cat *Category) error {
	catKey := strings.ToUpper(cat.Name)
	allExts := g.Cache.CateExtMap[catKey]
	if len(allExts) == 0 {
		return nil
	}

	// Filter out not-ready extensions and sort by ID
	var exts []*Extension
	for _, ext := range allExts {
		if ext.IsReady() {
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
	outputPath := filepath.Join(g.OutputDir, "cate", strings.ToLower(cat.Name)+".md")
	if err := WriteMarkdownFile(outputPath, content); err != nil {
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
func (g *IOCateGenerator) loadAvailability(ctx context.Context, pkgNames []string) (map[string]map[string]*PkgInfo, error) {
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
func (g *IOCateGenerator) generateContent(cat *Category, exts []*Extension, pkgAvail map[string]map[string]*PkgInfo) string {
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
func (g *IOCateGenerator) generateFrontmatter(cat *Category) string {
	catUpper := strings.ToUpper(cat.Name)
	enName := CategoryEnName(catUpper)
	desc := cat.Name
	if cat.EnDesc.Valid && cat.EnDesc.String != "" {
		desc = cat.EnDesc.String
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

`, catUpper, enName, catUpper, desc, CategoryWeight(cat.ID), icon)
}

// generateOverview generates the overview heading with counts
func (g *IOCateGenerator) generateOverview(exts []*Extension) string {
	pkgSet := make(map[string]bool)
	for _, ext := range exts {
		if ext.Lead {
			pkgSet[ext.Pkg] = true
		}
	}
	return fmt.Sprintf("## Extension List\n\nThere are **%d** extensions in **%d** packages.\n\n", len(exts), len(pkgSet))
}

// generateExtensionTable generates the overview extension table at the top
func (g *IOCateGenerator) generateExtensionTable(exts []*Extension) string {
	return IOExtensionTable(exts)
}

// generateExtCard generates a unified table section for a single extension
func (g *IOCateGenerator) generateExtCard(ext *Extension, avail map[string]*PkgInfo) string {
	var b strings.Builder

	desc := ext.GetEnDesc()

	// H2 heading: extension name
	b.WriteString(fmt.Sprintf("\n---------\n\n## %s {#%s}\n\n", ext.Name, ext.Name))

	// Intro line: [**`pkg`**](/ext/e/name) - `version` : description
	pkgLink := fmt.Sprintf("[**`%s`**](/ext/e/%s)", ext.Pkg, ext.Name)
	version := ext.GetVersion()
	b.WriteString(fmt.Sprintf("%s - `%s` : %s\n\n", pkgLink, version, desc))

	// Build the 7 left-side metadata rows
	leftRows := g.buildLeftRows(ext)

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

	// Header row with x86_64 / aarch64 columns
	b.WriteString("| **Item** | **Value** | **OS** | **x86_64** | **aarch64** |\n")
	b.WriteString("|:---:|:---|:---:|:---:|:---:|\n")

	// 7 data rows: left metadata + right OS availability
	numRows := len(leftRows)
	if len(g.OSCodes) < numRows {
		numRows = len(g.OSCodes)
	}

	for i := 0; i < numRows; i++ {
		left := leftRows[i]
		osc := g.OSCodes[i]
		b.WriteString(fmt.Sprintf("| **%s** | %s | **%s** |", left.Label, left.Value, osc.Code))
		g.writeAvailCells(&b, osc, avail, usePgVer, pgVerSet)
		b.WriteString("\n")
	}

	// Extra OS rows beyond the 7 left-side rows
	for i := numRows; i < len(g.OSCodes); i++ {
		osc := g.OSCodes[i]
		b.WriteString(fmt.Sprintf("| | | **%s** |", osc.Code))
		g.writeAvailCells(&b, osc, avail, usePgVer, pgVerSet)
		b.WriteString("\n")
	}

	b.WriteString("{.ext-table .ext-table--cate}\n\n")
	return b.String()
}

// buildLeftRows builds the 7 left-side metadata rows for an extension
func (g *IOCateGenerator) buildLeftRows(ext *Extension) []leftRow {
	// Row 1: Extension
	extLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)

	// Row 2: Package
	pkgLink := ext.GetPkgURLLink()

	// Row 3: RPM package name
	rpmPkg := "-"
	if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
		rpmPkg = fmt.Sprintf("`%s`", ext.RpmPkg.String)
	}

	// Row 4: DEB package name
	debPkg := "-"
	if ext.DebPkg.Valid && ext.DebPkg.String != "" {
		debPkg = fmt.Sprintf("`%s`", ext.DebPkg.String)
	}

	// Row 5: Language
	lang := "-"
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}

	// Row 6: Repo
	repo := "-"
	if ext.Repo.Valid && ext.Repo.String != "" {
		repo = CCRepoBadge(ext.Repo.String)
	}

	// Row 7: License
	license := "-"
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}

	return []leftRow{
		{"Extension", extLink},
		{"Package", pkgLink},
		{"RPM", rpmPkg},
		{"DEB", debPkg},
		{"Language", lang},
		{"Repo", repo},
		{"License", license},
	}
}

// collectAvailPGVers returns the list of available PG versions for a given OS name
func (g *IOCateGenerator) collectAvailPGVers(avail map[string]*PkgInfo, osName string) []string {
	if osName == "" {
		return nil
	}
	var vers []string
	for _, pg := range g.Cache.PGVersions {
		key := fmt.Sprintf("%d-%s", pg, osName)
		if pkg, exists := avail[key]; exists {
			state := "MISS"
			if pkg.State.Valid {
				state = pkg.State.String
			}
			if state == "AVAIL" {
				vers = append(vers, fmt.Sprintf("%d", pg))
			}
		}
	}
	return vers
}

// writeAvailCells writes two cells (x86_64, aarch64) with pgvers shortcodes
func (g *IOCateGenerator) writeAvailCells(b *strings.Builder, osc OSCodeEntry, avail map[string]*PkgInfo, usePgVer bool, pgVerSet map[int]bool) {
	if usePgVer {
		// Contrib/kernel-fork: use PgVer for both arches uniformly
		var vers []string
		for _, pg := range g.Cache.PGVersions {
			if pgVerSet[pg] {
				vers = append(vers, fmt.Sprintf("%d", pg))
			}
		}
		badge := CCPGVersShortcode(vers)
		b.WriteString(fmt.Sprintf(" %s | %s |", badge, badge))
	} else {
		// Normal: use actual package availability per arch
		for _, osName := range []string{osc.X86, osc.ARM} {
			vers := g.collectAvailPGVers(avail, osName)
			badge := CCPGVersShortcode(vers)
			b.WriteString(fmt.Sprintf(" %s |", badge))
		}
	}
}

// GenerateCategoryIndexPage generates the _index.md for the cate/ directory
func (g *IOCateGenerator) GenerateCategoryIndexPage() error {
	var b strings.Builder
	b.WriteString(`---
title: "Categories"
linkTitle: "Categories"
description: "PostgreSQL extensions organized into 16 functional categories for easy browsing."
weight: 500
icon: fa-solid fa-cubes
sidebar_expanded: true
---

`)

	b.WriteString("| **Category** | **Extensions** | **Packages** | **Description** |\n")
	b.WriteString("|:--------:|:-------:|:------:|:-------|\n")

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		allExts := g.Cache.CateExtMap[catKey]

		extCount := 0
		pkgCount := 0
		for _, ext := range allExts {
			if !ext.IsReady() {
				continue
			}
			extCount++
			if ext.Lead {
				pkgCount++
			}
		}

		desc := cat.Name
		if cat.EnDesc.Valid && cat.EnDesc.String != "" {
			desc = cat.EnDesc.String
		}

		lower := strings.ToLower(cat.Name)
		badge := fmt.Sprintf(`<a class="ext-badge ext-badge--cate %s" href="/ext/cate/%s">%s</a>`,
			lower, lower, catKey)

		b.WriteString(fmt.Sprintf("| %s | %d | %d | %s |\n",
			badge, extCount, pkgCount, desc))
	}

	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(filepath.Join(g.OutputDir, "cate", "_index.md"), b.String())
}

// GenerateAllCategoryPages generates all 16 category pages into cate/ subdirectory
func (g *IOCateGenerator) GenerateAllCategoryPages(ctx context.Context) error {
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
