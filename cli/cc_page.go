/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC Page Generator - generates extension detail pages for pigsty.cc (Chinese only)
Uses CSS classes and Hugo/Docsy shortcodes (tabpane) for styling
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// CCPageGenerator generates extension pages for pigsty.cc
type CCPageGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	StubDir   string
}

// NewCCPageGenerator creates a new CC page generator
func NewCCPageGenerator(cache *ExtensionCache, outputDir, stubDir string) *CCPageGenerator {
	return &CCPageGenerator{
		Cache:     cache,
		OutputDir: outputDir,
		StubDir:   stubDir,
	}
}

// GenerateExtensionPage generates a single extension page (Chinese only)
func (g *CCPageGenerator) GenerateExtensionPage(ctx context.Context, ext *Extension) error {
	content := g.generateExtensionContent(ctx, ext)

	// Append stub content if exists
	stubPath := filepath.Join(g.StubDir, ext.Name+".md")
	if stubContent, err := os.ReadFile(stubPath); err == nil && len(stubContent) > 0 {
		content += "\n" + string(stubContent)
	}

	return WriteMarkdownFile(filepath.Join(g.OutputDir, "e", ext.Name+".md"), content)
}

// generateExtensionContent generates the markdown content
func (g *CCPageGenerator) generateExtensionContent(ctx context.Context, ext *Extension) string {
	var b strings.Builder

	packages, _ := LoadPackages(ctx, ext.Pkg)
	binaries, _ := LoadBinaries(ctx, ext.Name)
	siblings := g.Cache.GetSiblingExtensions(ext.Pkg, ext.Name)
	allExts := g.getAllPackageExtensions(ext, siblings)

	b.WriteString(g.generateFrontmatter(ext))
	b.WriteString(g.generateCards(ext))
	b.WriteString("\n---------\n\n")
	b.WriteString(g.generateOverview(ext))
	b.WriteString(g.generateExtensionTable(ext, allExts))
	b.WriteString(g.generateRelationships(ext, siblings))
	if ext.Comment.Valid && ext.Comment.String != "" {
		b.WriteString(fmt.Sprintf("\n> %s\n\n", ext.Comment.String))
	}
	b.WriteString(g.generatePackages(ext))

	if !ext.Contrib {
		b.WriteString(g.generateAvailability(ext, packages, binaries))
		b.WriteString(g.generateBuild(ext))
	}

	b.WriteString(g.generateInstall(ext))

	return b.String()
}

// getAllPackageExtensions returns all extensions for the package (self + siblings), sorted by ID
func (g *CCPageGenerator) getAllPackageExtensions(ext *Extension, siblings []*Extension) []*Extension {
	all := make([]*Extension, 0, 1+len(siblings))
	all = append(all, ext)
	all = append(all, siblings...)
	sort.Slice(all, func(i, j int) bool { return all[i].ID < all[j].ID })
	return all
}

// generateFrontmatter generates YAML frontmatter
func (g *CCPageGenerator) generateFrontmatter(ext *Extension) string {
	desc := strings.ReplaceAll(ext.GetZhDesc(), `"`, `\"`)

	return fmt.Sprintf(`---
title: "%s"
linkTitle: "%s"
description: "%s"
weight: %d
---

`, ext.Name, ext.Name, desc, ext.ID)
}

// generateCards generates the top card section with repo and source links
func (g *CCPageGenerator) generateCards(ext *Extension) string {
	hasRepo := ext.URL.Valid && ext.URL.String != ""
	hasSource := ext.Source.Valid && ext.Source.String != ""

	if !hasRepo && !hasSource {
		return ""
	}

	var b strings.Builder
	b.WriteString("<div class=\"ext-cards\">\n")

	if hasRepo {
		url := ext.URL.String
		// Extract org/repo from GitHub URL
		title := url
		if strings.Contains(url, "github.com") {
			parts := strings.Split(strings.TrimRight(url, "/"), "/")
			if len(parts) >= 2 {
				title = parts[len(parts)-2] + "/" + parts[len(parts)-1]
			}
		}
		b.WriteString(fmt.Sprintf("  <a class=\"ext-card ext-card--repo\" href=\"%s\">\n", url))
		b.WriteString("    <div class=\"ext-card__kicker\">仓库</div>\n")
		b.WriteString(fmt.Sprintf("    <div class=\"ext-card__title\">%s</div>\n", title))
		b.WriteString(fmt.Sprintf("    <div class=\"ext-card__desc\">%s</div>\n", url))
		b.WriteString("  </a>\n")
	}

	if hasSource {
		sourceURL := ext.GetSourceURL(RegionChina)
		fileName := ext.GetSourceFilename()
		b.WriteString(fmt.Sprintf("  <a class=\"ext-card ext-card--source\" href=\"%s\">\n", sourceURL))
		b.WriteString("    <div class=\"ext-card__kicker\">源码</div>\n")
		b.WriteString(fmt.Sprintf("    <div class=\"ext-card__title\">%s</div>\n", fileName))
		b.WriteString(fmt.Sprintf("    <div class=\"ext-card__desc\">%s</div>\n", fileName))
		b.WriteString("  </a>\n")
	}

	b.WriteString("</div>\n\n")
	return b.String()
}

// generateOverview generates the overview table
func (g *CCPageGenerator) generateOverview(ext *Extension) string {
	var b strings.Builder
	b.WriteString("## 概览\n\n")

	// Package badge: lead pages point to themselves, sibling pages point to the lead page.
	pkgLink := fmt.Sprintf("[**`%s`**](/ext/e/%s)", ext.Pkg, ext.GetPkgPageName())

	version := "-"
	if ext.Version.Valid {
		version = fmt.Sprintf("`%s`", ext.Version.String)
	}

	category := "-"
	if ext.Category.Valid {
		category = CCCategoryBadge(ext.Category.String)
	}

	license := "-"
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}

	lang := "-"
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}

	b.WriteString("| **扩展包名** | **版本** | **分类** | **许可证** | **语言** |\n")
	b.WriteString("|:---------------------------------------------------:|:-------:|:--------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------:|:--------------------------------------------------------------------:|\n")
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", pkgLink, version, category, license, lang))
	b.WriteString("{.ext-table}\n\n")

	return b.String()
}

// generateExtensionTable generates the extension attributes table showing all extensions in the package
func (g *CCPageGenerator) generateExtensionTable(ext *Extension, allExts []*Extension) string {
	var b strings.Builder

	b.WriteString("|  ID   | **扩展名** | **Bin** | **Lib** | **Load** | **Create** | **Trust** | **Reloc** | **模式** |\n")
	b.WriteString("|:-----:|:-------------------------------------------------------------------------|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:--------------------------------------------:|:----------|\n")

	for _, e := range allExts {
		extLink := CCExtBoldLink(e.Name, fmt.Sprintf("/ext/e/%s", e.Name))

		schema := "-"
		if len(e.Schemas) > 0 {
			schema = fmt.Sprintf("`%s`", e.Schemas[0])
		}

		b.WriteString(fmt.Sprintf("| %d  | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			e.ID,
			extLink,
			CCFlagBadge(e.HasBin),
			CCFlagBadge(e.HasLib),
			CCFlagBadge(e.NeedLoad),
			CCFlagBadge(e.NeedDDL),
			CCOptFlagBadge(e.Trusted),
			CCOptFlagBadge(e.Relocatable),
			schema,
		))
	}

	b.WriteString("{.ext-table}\n\n")

	return b.String()
}

// generateRelationships generates the relationships section
func (g *CCPageGenerator) generateRelationships(ext *Extension, siblings []*Extension) string {
	if len(ext.Requires) == 0 && len(ext.RequireBy) == 0 && len(ext.SeeAlso) == 0 {
		return ""
	}

	var b strings.Builder

	// Build "相关扩展" header row content from requires + see_also
	headerParts := make([]string, 0)
	if len(ext.Requires) > 0 {
		for _, req := range ext.Requires {
			headerParts = append(headerParts, CCExtLink(req))
		}
	}
	if len(ext.SeeAlso) > 0 {
		for _, see := range ext.SeeAlso {
			headerParts = append(headerParts, CCExtLink(see))
		}
	}

	// If no related extensions at all (only require_by), use empty header
	headerContent := ""
	if len(headerParts) > 0 {
		headerContent = strings.Join(headerParts, " ")
	}

	b.WriteString(fmt.Sprintf("| **相关扩展** | %s |\n", headerContent))
	b.WriteString("|:--------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|\n")

	if len(ext.RequireBy) > 0 {
		links := make([]string, len(ext.RequireBy))
		for i, req := range ext.RequireBy {
			links[i] = CCExtLink(req)
		}
		b.WriteString(fmt.Sprintf("| **下游依赖** | %s |\n", strings.Join(links, " ")))
	}

	b.WriteString("{.ext-table .ext-table--rel}\n\n")
	return b.String()
}

// generatePackages generates the packages table
func (g *CCPageGenerator) generatePackages(ext *Extension) string {
	if ext.Contrib {
		return g.generateContribPackages(ext)
	}

	var b strings.Builder
	b.WriteString("\n## 版本\n\n")
	b.WriteString("| 类型 | 仓库 | 版本 | PG 大版本 | 包名 | 依赖 |\n")
	b.WriteString("|:----:|:----:|:----:|:------:|:--------:|:----:|\n")

	// EXT row
	if ext.Repo.Valid || ext.Version.Valid || len(ext.PgVer) > 0 {
		b.WriteString(g.buildPackageRow("EXT", ext.Category.String, ext.Repo.String, ext.Version.String, ext.Pkg, ext.PgVer, ext.Requires))
	}

	// RPM row
	if ext.RpmRepo.Valid || ext.RpmVer.Valid || len(ext.RpmPg) > 0 {
		pkg := ext.Pkg
		if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
			pkg = ext.RpmPkg.String
		}
		b.WriteString(g.buildPackageRow("RPM", ext.Category.String, ext.RpmRepo.String, ext.RpmVer.String, pkg, ext.RpmPg, ext.RpmDeps))
	}

	// DEB row
	if ext.DebRepo.Valid || ext.DebVer.Valid || len(ext.DebPg) > 0 {
		pkg := ext.Pkg
		if ext.DebPkg.Valid && ext.DebPkg.String != "" {
			pkg = ext.DebPkg.String
		}
		b.WriteString(g.buildPackageRow("DEB", ext.Category.String, ext.DebRepo.String, ext.DebVer.String, pkg, ext.DebPg, ext.DebDeps))
	}

	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

// buildPackageRow builds a single package row
func (g *CCPageGenerator) buildPackageRow(label, category, repo, version, pattern string, pgVers, deps []string) string {
	repoBadge := "-"
	if repo != "" {
		repoBadge = CCRepoBadge(repo)
	}

	verStr := "-"
	if version != "" {
		verStr = fmt.Sprintf("`%s`", version)
	}

	patternStr := "-"
	if pattern != "" {
		patternStr = fmt.Sprintf("`%s`", pattern)
	}

	pgBadges := CCPGVersShortcode(pgVers)

	depsStr := "-"
	if len(deps) > 0 {
		depList := make([]string, len(deps))
		for i, dep := range deps {
			depList[i] = fmt.Sprintf("`%s`", dep)
		}
		depsStr = strings.Join(depList, ", ")
	}

	// Map label to URL with category anchor
	catAnchor := ""
	if category != "" {
		catAnchor = "#" + strings.ToLower(category)
	}
	labelLink := fmt.Sprintf("**%s**", label)
	switch label {
	case "EXT":
		labelLink = fmt.Sprintf("[**EXT**](/ext/list%s)", catAnchor)
	case "RPM":
		labelLink = fmt.Sprintf("[**RPM**](/ext/rpm%s)", catAnchor)
	case "DEB":
		labelLink = fmt.Sprintf("[**DEB**](/ext/deb%s)", catAnchor)
	}

	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
		labelLink, repoBadge, verStr, pgBadges, patternStr, depsStr)
}

// generateContribPackages generates package info for contrib extensions
func (g *CCPageGenerator) generateContribPackages(ext *Extension) string {
	var b strings.Builder
	b.WriteString("\n## 版本\n\n")

	// Build header
	b.WriteString("|")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n|")

	availablePgs := make(map[int]bool)
	for _, pgStr := range ext.PgVer {
		if pgVersions := ParsePGVersions([]string{pgStr}); len(pgVersions) > 0 {
			for _, pg := range pgVersions {
				availablePgs[pg] = true
			}
		}
	}

	version := ext.GetVersion()

	for _, pg := range g.Cache.PGVersions {
		if availablePgs[pg] {
			b.WriteString(fmt.Sprintf(" %s |", CCAvailBadge(version)))
		} else {
			b.WriteString(fmt.Sprintf(" %s |", CCMissBadge()))
		}
	}
	b.WriteString("\n{.ext-table}\n\n")

	return b.String()
}

// generateAvailability generates the availability matrix using the pgext_matrix shortcode.
// Each cell is plain text: "STATE REPO VERSION PKG_NAME"
// The shortcode parses cells and renders two-line badges with CSS coloring.
func (g *CCPageGenerator) generateAvailability(ext *Extension, packages []*PkgInfo, binaries []*Binary) string {
	if len(packages) == 0 {
		return ""
	}

	var b strings.Builder

	pkgMap := make(map[string]*PkgInfo)
	for _, pkg := range packages {
		key := fmt.Sprintf("%d-%s", pkg.PG, pkg.OS)
		if _, exists := pkgMap[key]; !exists {
			pkgMap[key] = pkg
		}
	}

	b.WriteString("{{< pgext_matrix >}}\n")

	// Header
	b.WriteString("| **OS / PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--:|")
	for range g.Cache.PGVersions {
		b.WriteString(":--:|")
	}
	b.WriteString("\n")

	// Data rows
	for _, osv := range g.Cache.OSVersions {
		if !osv.Active {
			continue
		}

		b.WriteString(fmt.Sprintf("| %s |", osv.OS))

		for _, pg := range g.Cache.PGVersions {
			key := fmt.Sprintf("%d-%s", pg, osv.OS)

			if pkg, ok := pkgMap[key]; ok {
				cell := g.formatAvailCell(pkg, ext, osv.OS)
				b.WriteString(fmt.Sprintf(" %s |", cell))
			} else {
				// Infer repo for missing cells, count is 0
				repo := InferRepo(ext, osv.OS)
				b.WriteString(fmt.Sprintf(" MISS %s - 0 |", repo))
			}
		}
		b.WriteString("\n")
	}

	// Download data lines (@ OS PG PKG FILENAME REPO VERSION SIZE URL)
	if len(binaries) > 0 {
		activePG := make(map[int]bool, len(g.Cache.PGVersions))
		for _, pg := range g.Cache.PGVersions {
			activePG[pg] = true
		}
		for _, bin := range binaries {
			if !activePG[bin.PG] {
				continue
			}
			org := "-"
			if bin.Org.Valid && bin.Org.String != "" {
				org = bin.Org.String
			}
			size := strings.ReplaceAll(FormatSize(bin.Size), " ", "")
			url := bin.GetDownloadURL(RegionChina)
			b.WriteString(fmt.Sprintf("@ %s %d %s %s %s %s %s %s\n",
				bin.OS, bin.PG, bin.Name, bin.File, org, bin.Version, size, url))
		}
	}

	b.WriteString("{{< /pgext_matrix >}}\n\n")
	return b.String()
}

// formatAvailCell formats a cell as plain text: "STATE REPO VERSION COUNT"
func (g *CCPageGenerator) formatAvailCell(pkg *PkgInfo, ext *Extension, osName string) string {
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}

	org := InferRepo(ext, osName)
	if pkg.Org.Valid && pkg.Org.String != "" {
		org = strings.ToUpper(pkg.Org.String)
	}

	version := "-"
	if pkg.Version.Valid && pkg.Version.String != "" {
		version = pkg.Version.String
	}

	// MISS takes highest priority: if no actual package exists, override to MISS
	if version == "-" && state != "AVAIL" && state != "HIDE" {
		state = "MISS"
	}

	count := int64(1)
	if pkg.Count.Valid {
		count = pkg.Count.Int64
	}

	return fmt.Sprintf("%s %s %s %d", state, org, version, count)
}

// generateBuild generates the build section
func (g *CCPageGenerator) generateBuild(ext *Extension) string {
	if ext.Contrib {
		return ""
	}

	// Check extra.rpm and extra.deb to determine if build command should be shown
	var buildRPM, buildDEB bool
	if ext.Extra != nil {
		if rpm, ok := ext.Extra["rpm"]; ok {
			if rpmBool, ok := rpm.(bool); ok && rpmBool {
				buildRPM = true
			}
		}
		if deb, ok := ext.Extra["deb"]; ok {
			if debBool, ok := deb.(bool); ok && debBool {
				buildDEB = true
			}
		}
	}

	if !buildRPM && !buildDEB {
		return ""
	}

	var buildHint string
	if buildRPM && buildDEB {
		buildHint = "RPM / DEB 包"
	} else if buildRPM {
		buildHint = "RPM 包"
	} else {
		buildHint = "DEB 包"
	}

	var b strings.Builder
	b.WriteString("## 构建\n\n")
	b.WriteString(fmt.Sprintf("您可以使用 `pig build` 命令构建 `%s` 扩展的 %s：\n\n", ext.Pkg, buildHint))
	b.WriteString("```bash\n")
	b.WriteString(fmt.Sprintf("pig build pkg %s         # 构建 %s\n", ext.Pkg, buildHint))
	b.WriteString("```\n\n")
	return b.String()
}

// generateInstall generates the install section
func (g *CCPageGenerator) generateInstall(ext *Extension) string {
	var b strings.Builder
	b.WriteString("\n## 安装\n\n")

	// For contrib, just show CREATE EXTENSION
	if ext.Contrib {
		b.WriteString("> **提示**：这是 PostgreSQL 内核自带的 contrib 扩展\n\n")
		if ext.NeedDDL {
			b.WriteString("```sql\nCREATE EXTENSION " + ext.Name + ";\n```\n")
		}
		return b.String()
	}

	// Determine repo setup text
	if ext.Repo.Valid && ext.Repo.String == "PGDG" {
		b.WriteString("您可以直接安装 `" + ext.Pkg + "` 扩展包的预置二进制包，首先确保 [**PGDG**](/docs/repo/pgdg) 仓库已经添加并启用：\n\n")
		b.WriteString("```bash\npig repo add pgdg -u          # 添加 PGDG 仓库并更新缓存\n```\n\n")
	} else {
		b.WriteString("您可以直接安装 `" + ext.Pkg + "` 扩展包的预置二进制包，首先确保 [**PGDG**](/docs/repo/pgdg) 和 [**PIGSTY**](/docs/repo/pgsql) 仓库已经添加并启用：\n\n")
		b.WriteString("```bash\npig repo add pgsql -u          # 添加仓库并更新缓存\n```\n\n")
	}

	// Determine available PG versions
	activePG := make(map[int]bool, len(g.Cache.PGVersions))
	for _, pg := range g.Cache.PGVersions {
		activePG[pg] = true
	}

	var pgVersions []int
	for _, pg := range ParsePGVersions(ext.PgVer) {
		if activePG[pg] {
			pgVersions = append(pgVersions, pg)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pgVersions)))

	// Get RPM and DEB package patterns
	rpmPkg := ext.Pkg
	if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
		rpmPkg = ext.RpmPkg.String
	}
	debPkg := ext.Pkg
	if ext.DebPkg.Valid && ext.DebPkg.String != "" {
		debPkg = ext.DebPkg.String
	}

	b.WriteString("使用 [**pig**](/docs/pig) 或者是 `apt/yum/dnf` 安装扩展：\n\n")
	b.WriteString("{{< tabpane text=true persist=header >}}\n")

	// Tab: 安装 (simple install)
	b.WriteString("{{% tab header=\"安装\" %}}\n")
	b.WriteString("```bash\n")
	b.WriteString(fmt.Sprintf("pig install %s;          # 当前活跃 PG 版本安装\n", ext.Pkg))
	b.WriteString("```\n")
	b.WriteString("{{% /tab %}}\n")

	// Tab: pig (per-version)
	b.WriteString("{{% tab header=\"pig\" %}}\n")
	b.WriteString("```bash\n")
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf("pig ext install -y %s -v %d  # PG %d\n", ext.Pkg, pg, pg))
	}
	b.WriteString("```\n")
	b.WriteString("{{% /tab %}}\n")

	// Tab: dnf (RPM)
	if ext.RpmRepo.Valid || ext.RpmVer.Valid || len(ext.RpmPg) > 0 {
		b.WriteString("{{% tab header=\"dnf\" %}}\n")
		b.WriteString("```bash\n")
		for _, pg := range pgVersions {
			pkgName := strings.ReplaceAll(rpmPkg, "$v", fmt.Sprintf("%d", pg))
			b.WriteString(fmt.Sprintf("dnf install -y %s       # PG %d\n", pkgName, pg))
		}
		b.WriteString("```\n")
		b.WriteString("{{% /tab %}}\n")
	}

	// Tab: apt (DEB)
	if ext.DebRepo.Valid || ext.DebVer.Valid || len(ext.DebPg) > 0 {
		b.WriteString("{{% tab header=\"apt\" %}}\n")
		b.WriteString("```bash\n")
		for _, pg := range pgVersions {
			pkgName := strings.ReplaceAll(debPkg, "$v", fmt.Sprintf("%d", pg))
			b.WriteString(fmt.Sprintf("apt install -y %s   # PG %d\n", pkgName, pg))
		}
		b.WriteString("```\n")
		b.WriteString("{{% /tab %}}\n")
	}

	b.WriteString("{{< /tabpane >}}\n\n")

	// Shared preload libraries
	if ext.NeedLoad {
		b.WriteString("\n**预加载配置**：\n\n")

		var libsToLoad []string
		libSet := make(map[string]bool)

		for _, depName := range ext.Requires {
			if depExt, exists := g.Cache.ExtMap[depName]; exists && depExt.NeedLoad {
				depLibName := depExt.GetLibName()
				for _, lib := range strings.Split(depLibName, ",") {
					lib = strings.TrimSpace(lib)
					if lib != "" && !libSet[lib] {
						libSet[lib] = true
						libsToLoad = append(libsToLoad, lib)
					}
				}
			}
		}

		extLibName := ext.GetLibName()
		for _, lib := range strings.Split(extLibName, ",") {
			lib = strings.TrimSpace(lib)
			if lib != "" && !libSet[lib] {
				libSet[lib] = true
				libsToLoad = append(libsToLoad, lib)
			}
		}

		libConfig := strings.Join(libsToLoad, ", ")
		b.WriteString("```bash\n")
		b.WriteString(fmt.Sprintf("shared_preload_libraries = '%s';\n", libConfig))
		b.WriteString("```\n\n")
	}

	// CREATE EXTENSION
	if ext.NeedDDL {
		b.WriteString("\n**创建扩展**：\n\n")
		b.WriteString("```sql\n")
		if len(ext.Requires) > 0 {
			b.WriteString(fmt.Sprintf("CREATE EXTENSION %s CASCADE;  -- 依赖: %s\n", ext.Name, strings.Join(ext.Requires, ", ")))
		} else {
			b.WriteString(fmt.Sprintf("CREATE EXTENSION %s;\n", ext.Name))
		}
		b.WriteString("```\n")
	} else if !ext.NeedLoad {
		b.WriteString("> 此扩展不需要执行 `CREATE EXTENSION` 语句\n")
	}

	return b.String()
}
