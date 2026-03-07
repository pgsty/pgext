/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC Page Generator - generates extension detail pages for pigsty.cc (Chinese only)
Uses CSS classes and Hugo/Docsy shortcodes (tabpane) for styling
*/
package cli

import (
	"context"
	"database/sql"
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
}

// NewCCPageGenerator creates a new CC page generator
func NewCCPageGenerator(cache *ExtensionCache, outputDir string) *CCPageGenerator {
	return &CCPageGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateExtensionPage generates a single extension page (Chinese only)
func (g *CCPageGenerator) GenerateExtensionPage(ctx context.Context, ext *Extension) error {
	extDir := filepath.Join(g.OutputDir, "e")
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extension directory: %w", err)
	}

	content := g.generateExtensionContent(ctx, ext)

	// Append stub-zh content if exists
	stubPath := filepath.Join("stub-zh", ext.Name+".md")
	if stubContent, err := os.ReadFile(stubPath); err == nil && len(stubContent) > 0 {
		content += "\n" + string(stubContent)
	}

	outputPath := filepath.Join(extDir, ext.Name+".md")
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
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
	b.WriteString("\n> Bin: 二进制；Lib：库文件；Load：显式加载；Create：需要DDL创建；Trust：普通用户可创建；Reloc：安装至其他模式\n\n")
	b.WriteString(g.generatePackages(ext))

	if !ext.Contrib {
		b.WriteString(g.generateAvailability(packages))
		b.WriteString(g.generateDownloads(binaries))
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
	desc := ext.Name
	if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
		desc = ext.ZhDesc.String
	} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
		desc = ext.EnDesc.String
	}
	desc = strings.ReplaceAll(desc, `"`, `\"`)

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
		source := ext.Source.String
		// Extract filename from source URL/path
		fileName := source
		if idx := strings.LastIndex(source, "/"); idx >= 0 {
			fileName = source[idx+1:]
		}
		b.WriteString(fmt.Sprintf("  <a class=\"ext-card ext-card--source\" href=\"%s\">\n", source))
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

	// Extension badge with URL
	extBadge := fmt.Sprintf("**`%s`**", ext.Name)
	if ext.URL.Valid && ext.URL.String != "" {
		extBadge = fmt.Sprintf("[**`%s`**](%s)", ext.Name, ext.URL.String)
	}

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
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", extBadge, version, category, license, lang))
	b.WriteString("{.ext-table}\n\n")

	return b.String()
}

// generateExtensionTable generates the extension attributes table showing all extensions in the package
func (g *CCPageGenerator) generateExtensionTable(ext *Extension, allExts []*Extension) string {
	var b strings.Builder

	b.WriteString("|  ID   | **扩展名** | **Bin** | **Lib** | **Load** | **Create** | **Trust** | **Reloc** | **模式** |\n")
	b.WriteString("|:-----:|:-------------------------------------------------------------------------|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:--------------------------------------------:|:----------|\n")

	for _, e := range allExts {
		extLink := CCExtBoldLink(e.Name, "")
		if ext.URL.Valid && ext.URL.String != "" {
			extLink = CCExtBoldLink(e.Name, ext.URL.String)
		}

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
	if len(ext.Requires) == 0 && len(ext.RequireBy) == 0 && len(ext.SeeAlso) == 0 && len(siblings) == 0 {
		return ""
	}

	var b strings.Builder

	b.WriteString("| **关联扩展** | |\n")
	b.WriteString("|:--------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|\n")

	if len(ext.Requires) > 0 {
		links := make([]string, len(ext.Requires))
		for i, req := range ext.Requires {
			links[i] = CCExtLink(req)
		}
		b.WriteString(fmt.Sprintf("| 上游依赖 | %s |\n", strings.Join(links, " ")))
	}

	if len(ext.RequireBy) > 0 {
		links := make([]string, len(ext.RequireBy))
		for i, req := range ext.RequireBy {
			links[i] = CCExtLink(req)
		}
		b.WriteString(fmt.Sprintf("| 下游依赖 | %s |\n", strings.Join(links, " ")))
	}

	if len(ext.SeeAlso) > 0 {
		links := make([]string, len(ext.SeeAlso))
		for i, see := range ext.SeeAlso {
			links[i] = CCExtLink(see)
		}
		b.WriteString(fmt.Sprintf("| 相关扩展 | %s |\n", strings.Join(links, " ")))
	}

	if len(siblings) > 0 {
		links := make([]string, len(siblings))
		for i, sib := range siblings {
			links[i] = CCExtLink(sib.Name)
		}
		b.WriteString(fmt.Sprintf("| 兄弟扩展 | %s |\n", strings.Join(links, " ")))
	}

	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

// generatePackages generates the packages table
func (g *CCPageGenerator) generatePackages(ext *Extension) string {
	if ext.Contrib {
		return g.generateContribPackages(ext)
	}

	var b strings.Builder
	b.WriteString("\n## 软件包\n\n")
	b.WriteString("| 类型 | 仓库 | 版本 | PG兼容 | 包名模式 | 依赖 |\n")
	b.WriteString("|:----:|:----:|:----:|:------:|:--------:|:----:|\n")

	// EXT row
	if ext.Repo.Valid || ext.Version.Valid || len(ext.PgVer) > 0 {
		b.WriteString(g.buildPackageRow("EXT", ext.Repo.String, ext.Version.String, ext.Pkg, ext.PgVer, ext.Requires))
	}

	// RPM row
	if ext.RpmRepo.Valid || ext.RpmVer.Valid || len(ext.RpmPg) > 0 {
		pkg := ext.Pkg
		if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
			pkg = ext.RpmPkg.String
		}
		b.WriteString(g.buildPackageRow("RPM", ext.RpmRepo.String, ext.RpmVer.String, pkg, ext.RpmPg, ext.RpmDeps))
	}

	// DEB row
	if ext.DebRepo.Valid || ext.DebVer.Valid || len(ext.DebPg) > 0 {
		pkg := ext.Pkg
		if ext.DebPkg.Valid && ext.DebPkg.String != "" {
			pkg = ext.DebPkg.String
		}
		b.WriteString(g.buildPackageRow("DEB", ext.DebRepo.String, ext.DebVer.String, pkg, ext.DebPg, ext.DebDeps))
	}

	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

// buildPackageRow builds a single package row
func (g *CCPageGenerator) buildPackageRow(label, repo, version, pattern string, pgVers, deps []string) string {
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

	pgBadges := g.buildPGCompatBadges(pgVers)

	depsStr := "-"
	if len(deps) > 0 {
		depList := make([]string, len(deps))
		for i, dep := range deps {
			depList[i] = fmt.Sprintf("`%s`", dep)
		}
		depsStr = strings.Join(depList, ", ")
	}

	return fmt.Sprintf("| **%s** | %s | %s | %s | %s | %s |\n",
		label, repoBadge, verStr, pgBadges, patternStr, depsStr)
}

// buildPGCompatBadges builds PG version compatibility badges with CSS classes
func (g *CCPageGenerator) buildPGCompatBadges(supported []string) string {
	if len(supported) == 0 {
		return "-"
	}

	supportedMap := make(map[int]bool)
	for _, v := range supported {
		v = strings.TrimSuffix(strings.TrimSpace(v), "+")
		var ver int
		if _, err := fmt.Sscanf(v, "%d", &ver); err == nil {
			supportedMap[ver] = true
		}
	}

	var badges []string
	for _, pg := range g.Cache.PGVersions {
		badges = append(badges, CCPGVerBadge(fmt.Sprintf("%d", pg), supportedMap[pg]))
	}

	return strings.Join(badges, " ")
}

// generateContribPackages generates package info for contrib extensions
func (g *CCPageGenerator) generateContribPackages(ext *Extension) string {
	var b strings.Builder
	b.WriteString("\n## 软件包\n\n")

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

	version := "-"
	if ext.Version.Valid {
		version = ext.Version.String
	}

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

// generateAvailability generates the availability matrix
func (g *CCPageGenerator) generateAvailability(packages []*PkgInfo) string {
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

	b.WriteString("## 可用性\n\n")

	// Header
	b.WriteString("| **Linux / PGSQL** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:-----------------------------------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":-----------------------------------------------------:|")
	}
	b.WriteString("\n")

	// Data rows
	for _, os := range g.Cache.OSVersions {
		if !os.Active {
			continue
		}

		b.WriteString(fmt.Sprintf("| %s |", CCOSLink(os.OS)))

		for _, pg := range g.Cache.PGVersions {
			key := fmt.Sprintf("%d-%s", pg, os.OS)

			if pkg, ok := pkgMap[key]; ok {
				cell := g.formatAvailCell(pkg)
				b.WriteString(fmt.Sprintf(" %s |", cell))
			} else {
				b.WriteString(fmt.Sprintf(" %s |", CCMissBadge()))
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("{.ext-table .ext-table--matrix}\n\n")
	return b.String()
}

// formatAvailCell formats an availability cell
func (g *CCPageGenerator) formatAvailCell(pkg *PkgInfo) string {
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}

	switch state {
	case "MISS":
		return CCMissBadge()
	case "HIDE":
		return `<span class="ext-badge ext-badge--hide">-</span>`
	case "THROW":
		return `<span class="ext-badge ext-badge--throw">!</span>`
	case "BREAK":
		return `<span class="ext-badge ext-badge--break">!</span>`
	case "AVAIL":
		version := ""
		if pkg.Version.Valid {
			version = pkg.Version.String
		}
		if version != "" {
			return CCAvailBadge(version)
		}
		return CCAvailBadge("✓")
	default:
		return `<span class="ext-badge ext-badge--hide">-</span>`
	}
}

// generateDownloads generates the downloads section with Hugo tabpane shortcodes
func (g *CCPageGenerator) generateDownloads(binaries []*Binary) string {
	if len(binaries) == 0 {
		return ""
	}

	activePG := make(map[int]bool, len(g.Cache.PGVersions))
	for _, pg := range g.Cache.PGVersions {
		activePG[pg] = true
	}

	pgGroups := make(map[int][]*Binary)
	for _, bin := range binaries {
		if !activePG[bin.PG] {
			continue
		}
		pgGroups[bin.PG] = append(pgGroups[bin.PG], bin)
	}

	if len(pgGroups) == 0 {
		return ""
	}

	var pgVersions []int
	for pg := range pgGroups {
		pgVersions = append(pgVersions, pg)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pgVersions)))

	var b strings.Builder
	b.WriteString("## 下载\n\n")
	b.WriteString("<div class=\"ext-tabs\">\n")
	b.WriteString("{{< tabpane text=true persist=\"disabled\" >}}\n\n")

	for _, pg := range pgVersions {
		bins := pgGroups[pg]
		if len(bins) == 0 {
			continue
		}

		b.WriteString(fmt.Sprintf("{{%% tab header=\"PG%d\" %%}}\n", pg))
		b.WriteString("| 包名 | 版本 | 系统 | 来源 | 大小 |\n")
		b.WriteString("|:-----|:----:|:----:|:----:|-----:|\n")

		for _, bin := range bins {
			org := "-"
			if bin.Org.Valid {
				org = bin.Org.String
			}

			b.WriteString(fmt.Sprintf("| [`%s`](%s) | `%s` | [%s](/ext/os/%s) | %s | %s |\n",
				bin.Name, bin.Href, bin.Version, bin.OS, bin.OS, org, FormatSize(bin.Size)))
		}

		b.WriteString("{.ext-table .ext-table--download}\n")
		b.WriteString("{{% /tab %}}\n\n")
	}

	b.WriteString("{{< /tabpane >}}\n")
	b.WriteString("</div>\n\n")

	return b.String()
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

// NullBool type alias for sql.NullBool
type NullBool = sql.NullBool
