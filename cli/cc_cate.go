/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

OINK Category Generator - generates per-category extension index pages for pigsty.cc (ZH) & pigsty.io (EN).
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

// OSCodeEntry represents a unique os_code (e.g., el8, el9, d12) with its two arch variants.
type OSCodeEntry struct {
	Code   string
	X86    string
	ARM    string
	HasX86 bool
	HasARM bool
}

// CCCateGenerator generates per-category extension index pages for pigsty.cc (DocsZH) or pigsty.io (DocsEN).
type CCCateGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	OSCodes   []OSCodeEntry
	Locale    DocsLocale
}

func newDocsCateGenerator(cache *ExtensionCache, outputDir string, locale DocsLocale) *CCCateGenerator {
	var osCodes []OSCodeEntry
	if cache != nil {
		osCodes = buildOSCodeEntries(cache.OSVersions)
	}
	return &CCCateGenerator{Cache: cache, OutputDir: outputDir, OSCodes: osCodes, Locale: locale}
}

func NewCCCateGenerator(cache *ExtensionCache, outputDir string) *CCCateGenerator {
	return newDocsCateGenerator(cache, outputDir, DocsZH)
}

func NewIOCateGenerator(cache *ExtensionCache, outputDir string) *CCCateGenerator {
	return newDocsCateGenerator(cache, outputDir, DocsEN)
}

func (g *CCCateGenerator) locale() DocsLocale {
	if g.Locale == "" {
		return DocsZH
	}
	return g.Locale
}

// CategoryWeight returns the frontmatter weight for a category (800 + category ID).
func CategoryWeight(catID int) int {
	return 800 + catID
}

// GenerateCategoryPage generates a single category page.
func (g *CCCateGenerator) GenerateCategoryPage(ctx context.Context, cat *Category) error {
	catKey := strings.ToUpper(cat.Name)
	allExts := g.Cache.CateExtMap[catKey]
	if len(allExts) == 0 {
		return nil
	}

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

	pkgSet := make(map[string]bool)
	for _, ext := range exts {
		hasKernel := ext.Extra != nil && ext.Extra["kernel"] != nil
		if !ext.Contrib && !hasKernel {
			pkgSet[ext.Pkg] = true
		}
	}
	pkgNames := make([]string, 0, len(pkgSet))
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

func (g *CCCateGenerator) loadAvailability(ctx context.Context, pkgNames []string) (map[string]map[string]*PkgInfo, error) {
	if len(pkgNames) == 0 {
		return make(map[string]map[string]*PkgInfo), nil
	}
	const query = `
		SELECT pg, os, pkg, state, version
		FROM pgext.pkg
		WHERE pkg = ANY($1)
		ORDER BY pkg, os, pg DESC
	`
	rows, err := QueryContext(ctx, query, pkgNames)
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

func (g *CCCateGenerator) generateContent(cat *Category, exts []*Extension, pkgAvail map[string]map[string]*PkgInfo) string {
	var b strings.Builder
	b.WriteString(g.generateFrontmatter(cat))
	b.WriteString(g.generateOverview(exts))
	b.WriteString(DocsExtensionTable(g.locale(), exts))
	for _, ext := range exts {
		avail := pkgAvail[ext.Pkg]
		if avail == nil {
			avail = make(map[string]*PkgInfo)
		}
		b.WriteString(g.generateExtCard(ext, avail))
	}
	return b.String()
}

func (g *CCCateGenerator) generateFrontmatter(cat *Category) string {
	l := g.locale()
	catUpper := strings.ToUpper(cat.Name)
	icon := ""
	if cat.Icon1.Valid && cat.Icon1.String != "" {
		icon = fmt.Sprintf("icon: %s\n", cat.Icon1.String)
	}
	return fmt.Sprintf("---\ntitle: \"%s - %s\"\nlinkTitle: \"%s\"\ndescription: \"%s\"\nweight: %d\n%s---\n\n",
		catUpper, l.catName(catUpper), catUpper, l.catDesc(cat), CategoryWeight(cat.ID), icon)
}

func (g *CCCateGenerator) generateOverview(exts []*Extension) string {
	pkgSet := make(map[string]bool)
	for _, ext := range exts {
		if ext.Lead {
			pkgSet[ext.Pkg] = true
		}
	}
	return fmt.Sprintf(g.locale().t(
		"## 扩展列表\n\n共有 **%d** 个扩展，位于 **%d** 个扩展包中。\n\n",
		"## Extension List\n\nThere are **%d** extensions in **%d** packages.\n\n",
	), len(exts), len(pkgSet))
}

func (g *CCCateGenerator) generateExtCard(ext *Extension, avail map[string]*PkgInfo) string {
	l := g.locale()
	var b strings.Builder

	b.WriteString(fmt.Sprintf("\n---------\n\n## %s {#%s}\n\n", ext.Name, ext.Name))
	b.WriteString(fmt.Sprintf("[**`%s`**](/ext/e/%s) - `%s` : %s\n\n", ext.Pkg, ext.Name, ext.GetVersion(), l.extDesc(ext)))

	leftRows := g.buildLeftRows(ext)
	hasKernel := ext.Extra != nil && ext.Extra["kernel"] != nil
	usePgVer := ext.Contrib || hasKernel
	var pgVerSet map[int]bool
	if usePgVer {
		pgVerSet = make(map[int]bool)
		for _, pg := range ParsePGVersions(ext.PgVer) {
			pgVerSet[pg] = true
		}
	}

	b.WriteString(l.t("| **条目** | **属性** | **OS** | **x86_64** | **aarch64** |\n", "| **Item** | **Value** | **OS** | **x86_64** | **aarch64** |\n"))
	b.WriteString("|:---:|:---|:---:|:---:|:---:|\n")

	totalRows := max(len(leftRows), len(g.OSCodes))
	for i := 0; i < totalRows; i++ {
		if i < len(leftRows) {
			b.WriteString(fmt.Sprintf("| **%s** | %s |", leftRows[i].Label, leftRows[i].Value))
		} else {
			b.WriteString("| | |")
		}
		if i < len(g.OSCodes) {
			osc := g.OSCodes[i]
			b.WriteString(fmt.Sprintf(" **%s** |", osc.Code))
			g.writeAvailCells(&b, osc, avail, usePgVer, pgVerSet)
		} else {
			b.WriteString(" | | |")
		}
		b.WriteString("\n")
	}

	b.WriteString("{.ext-table .ext-table--cate}\n\n")
	return b.String()
}

type leftRow struct {
	Label string
	Value string
}

func (g *CCCateGenerator) buildLeftRows(ext *Extension) []leftRow {
	l := g.locale()
	rpmPkg, debPkg, lang, repo, license := "-", "-", "-", "-", "-"
	if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
		rpmPkg = fmt.Sprintf("`%s`", ext.RpmPkg.String)
	}
	if ext.DebPkg.Valid && ext.DebPkg.String != "" {
		debPkg = fmt.Sprintf("`%s`", ext.DebPkg.String)
	}
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}
	if ext.Repo.Valid && ext.Repo.String != "" {
		repo = CCRepoBadge(ext.Repo.String)
	}
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}
	return []leftRow{
		{l.t("扩展名", "Extension"), fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)},
		{l.t("扩展包", "Package"), ext.GetPkgURLLink()},
		{"RPM", rpmPkg},
		{"DEB", debPkg},
		{l.t("语言", "Language"), lang},
		{l.t("仓库", "Repo"), repo},
		{l.t("协议", "License"), license},
	}
}

func (g *CCCateGenerator) collectAvailPGVers(avail map[string]*PkgInfo, osName string) []string {
	if osName == "" {
		return nil
	}
	var vers []string
	for _, pg := range g.Cache.PGVersions {
		key := fmt.Sprintf("%d-%s", pg, osName)
		if pkg, exists := avail[key]; exists && pkg.State.Valid && pkg.State.String == "AVAIL" {
			vers = append(vers, fmt.Sprintf("%d", pg))
		}
	}
	return vers
}

func (g *CCCateGenerator) writeAvailCells(b *strings.Builder, osc OSCodeEntry, avail map[string]*PkgInfo, usePgVer bool, pgVerSet map[int]bool) {
	if usePgVer {
		var vers []string
		for _, pg := range g.Cache.PGVersions {
			if pgVerSet[pg] {
				vers = append(vers, fmt.Sprintf("%d", pg))
			}
		}
		badge := CCPGVersShortcode(vers)
		b.WriteString(fmt.Sprintf(" %s | %s |", badge, badge))
		return
	}
	for _, osName := range []string{osc.X86, osc.ARM} {
		b.WriteString(fmt.Sprintf(" %s |", CCPGVersShortcode(g.collectAvailPGVers(avail, osName))))
	}
}

func (g *CCCateGenerator) GenerateCategoryIndexPage() error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 500\nicon: fa-solid fa-cubes\nsidebar_expanded: true\n---\n\n",
		l.t("扩展分类", "Categories"),
		l.t("扩展分类", "Categories"),
		l.t("将 PostgreSQL 扩展按照功能分为 16 个大类别，方便用户按功能浏览扩展包索引。", "PostgreSQL extensions organized into 16 functional categories for easy browsing."),
	))
	b.WriteString(l.t("| **分类** | **扩展数** | **包数** | **描述** |\n", "| **Category** | **Extensions** | **Packages** | **Description** |\n"))
	b.WriteString("|:--------:|:-------:|:------:|:-------|\n")

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		extCount, pkgCount := 0, 0
		for _, ext := range g.Cache.CateExtMap[catKey] {
			if !ext.IsReady() {
				continue
			}
			extCount++
			if ext.Lead {
				pkgCount++
			}
		}
		b.WriteString(fmt.Sprintf("| %s | %d | %d | %s |\n", CCCategoryBadge(cat.Name), extCount, pkgCount, l.catDesc(cat)))
	}
	b.WriteString("{.ext-table}\n\n")
	return WriteMarkdownFile(filepath.Join(g.OutputDir, "cate", "_index.md"), b.String())
}

func (g *CCCateGenerator) GenerateAllCategoryPages(ctx context.Context) error {
	if err := g.GenerateCategoryIndexPage(); err != nil {
		return fmt.Errorf("failed to generate category index: %w", err)
	}
	var failed []string
	for _, cat := range g.Cache.Categories {
		if err := g.GenerateCategoryPage(ctx, cat); err != nil {
			logrus.Errorf("Failed to generate category page for %s: %v", cat.Name, err)
			failed = append(failed, cat.Name)
		}
	}
	logrus.Infof("Generated %d/%d category pages", len(g.Cache.Categories)-len(failed), len(g.Cache.Categories))
	if len(failed) > 0 {
		return fmt.Errorf("failed to generate %d category pages: %s", len(failed), strings.Join(failed, ", "))
	}
	return nil
}
