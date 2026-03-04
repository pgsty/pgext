/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC Page Generator - generates extension detail pages for pigsty.cc (Chinese only)
Complete information with beautiful Tailwind-inspired styling
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
	outputPath := filepath.Join(extDir, ext.Name+".md")
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// generateExtensionContent generates the markdown content
func (g *CCPageGenerator) generateExtensionContent(ctx context.Context, ext *Extension) string {
	var b strings.Builder

	// Load data
	packages, _ := LoadPackages(ctx, ext.Pkg)
	binaries, _ := LoadBinaries(ctx, ext.Name)
	siblings := g.Cache.GetSiblingExtensions(ext.Pkg, ext.Name)

	b.WriteString(g.generateFrontmatter(ext))
	b.WriteString(g.generateHeader(ext))
	b.WriteString(g.generateOverview(ext))
	b.WriteString(g.generateAttributes(ext))
	b.WriteString(g.generateRelationships(ext, siblings))
	b.WriteString(g.generateComment(ext))
	b.WriteString(g.generatePackages(ext))

	if !ext.Contrib {
		b.WriteString(g.generateAvailability(packages))
		b.WriteString(g.generateDownloads(binaries))
	} else {
		b.WriteString(g.generateContribTip())
	}

	b.WriteString(g.generateSource(ext))
	b.WriteString(g.generateInstall(ext))

	return b.String()
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

// generateHeader generates the title with package link and description
func (g *CCPageGenerator) generateHeader(ext *Extension) string {
	var b strings.Builder

	// Description
	desc := ext.EnDesc.String
	if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
		desc = ext.ZhDesc.String
	}

	// Package link with URL if available
	if ext.URL.Valid && ext.URL.String != "" {
		b.WriteString(fmt.Sprintf("[**`%s`**](%s): %s\n\n", ext.Pkg, ext.URL.String, desc))
	} else {
		b.WriteString(fmt.Sprintf("**`%s`**: %s\n\n", ext.Pkg, desc))
	}

	return b.String()
}

// generateOverview generates the overview table
func (g *CCPageGenerator) generateOverview(ext *Extension) string {
	var b strings.Builder

	b.WriteString("## 概览\n\n")

	// Extension badge with URL
	extBadge := fmt.Sprintf("`%s`", ext.Name)
	if ext.URL.Valid && ext.URL.String != "" {
		extBadge = fmt.Sprintf("[`%s`](%s)", ext.Name, ext.URL.String)
	}

	// Package link
	pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Pkg, ext.Name)

	// Version
	version := "-"
	if ext.Version.Valid {
		version = ext.Version.String
	}

	// Category badge
	category := "-"
	if ext.Category.Valid {
		category = CCCategoryBadge(ext.Category.String)
	}

	// License badge
	license := "-"
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}

	// Language badge
	lang := "-"
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}

	b.WriteString("| **编号** | **扩展名** | **包名** | **版本** | **分类** | **许可证** | **语言** |\n")
	b.WriteString("|:--------:|:----------:|:--------:|:--------:|:--------:|:----------:|:--------:|\n")
	b.WriteString(fmt.Sprintf("| %d | %s | %s | `%s` | %s | %s | %s |\n\n",
		ext.ID, extBadge, pkgLink, version, category, license, lang))

	return b.String()
}

// generateAttributes generates the attributes table
func (g *CCPageGenerator) generateAttributes(ext *Extension) string {
	var b strings.Builder

	// Attribute badge string
	attrs := ext.GetAttributeBadge()
	attrBadge := fmt.Sprintf("`%s`", attrs)

	b.WriteString("| **属性** | **二进制** | **动态库** | **预加载** | **DDL** | **可信** | **可重定位** |\n")
	b.WriteString("|:--------:|:----------:|:----------:|:----------:|:-------:|:--------:|:------------:|\n")
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n\n",
		attrBadge,
		g.boolBadge(ext.HasBin, "blue", "green"),
		g.boolBadge(ext.HasLib, "blue", "green"),
		g.boolBadge(ext.NeedLoad, "blue", "amber"),
		g.boolBadge(ext.NeedDDL, "amber", "green"),
		g.optBoolBadge(ext.Trusted),
		g.optBoolBadge(ext.Relocatable),
	))

	return b.String()
}

// boolBadge generates a colored badge for boolean value
func (g *CCPageGenerator) boolBadge(val bool, falseColor, trueColor string) string {
	if val {
		return CCColorBadge("是", trueColor)
	}
	return CCColorBadge("否", falseColor)
}

// optBoolBadge generates a colored badge for optional boolean
func (g *CCPageGenerator) optBoolBadge(val NullBool) string {
	if !val.Valid {
		return CCColorBadge("-", "gray")
	}
	if val.Bool {
		return CCColorBadge("是", "green")
	}
	return CCColorBadge("否", "amber")
}

// generateRelationships generates the relationships section
func (g *CCPageGenerator) generateRelationships(ext *Extension, siblings []*Extension) string {
	if len(ext.Schemas) == 0 && len(ext.Requires) == 0 && len(ext.RequireBy) == 0 && len(ext.SeeAlso) == 0 && len(siblings) == 0 {
		return ""
	}

	var b strings.Builder

	b.WriteString("| **依赖关系** | |\n")
	b.WriteString("|:------------:|:---|\n")

	if len(ext.Schemas) > 0 {
		schemas := make([]string, len(ext.Schemas))
		for i, s := range ext.Schemas {
			schemas[i] = fmt.Sprintf("`%s`", s)
		}
		b.WriteString(fmt.Sprintf("| **模式** | %s |\n", strings.Join(schemas, " ")))
	}

	if len(ext.Requires) > 0 {
		links := make([]string, len(ext.Requires))
		for i, req := range ext.Requires {
			links[i] = CCExtBadge(req)
		}
		b.WriteString(fmt.Sprintf("| **依赖** | %s |\n", strings.Join(links, " ")))
	}

	if len(ext.RequireBy) > 0 {
		links := make([]string, len(ext.RequireBy))
		for i, req := range ext.RequireBy {
			links[i] = CCExtBadge(req)
		}
		b.WriteString(fmt.Sprintf("| **被依赖** | %s |\n", strings.Join(links, " ")))
	}

	if len(ext.SeeAlso) > 0 {
		links := make([]string, len(ext.SeeAlso))
		for i, see := range ext.SeeAlso {
			links[i] = CCExtBadge(see)
		}
		b.WriteString(fmt.Sprintf("| **相关** | %s |\n", strings.Join(links, " ")))
	}

	if len(siblings) > 0 {
		links := make([]string, len(siblings))
		for i, sib := range siblings {
			links[i] = CCExtBadge(sib.Name)
		}
		b.WriteString(fmt.Sprintf("| **同包** | %s |\n", strings.Join(links, " ")))
	}

	b.WriteString("\n")
	return b.String()
}

// generateComment generates the comment note if present
func (g *CCPageGenerator) generateComment(ext *Extension) string {
	if !ext.Comment.Valid || ext.Comment.String == "" {
		return ""
	}
	return fmt.Sprintf("> **备注**：%s\n\n", ext.Comment.String)
}

// generatePackages generates the packages table
func (g *CCPageGenerator) generatePackages(ext *Extension) string {
	var b strings.Builder

	// For contrib extensions
	if ext.Contrib {
		return g.generateContribPackages(ext)
	}

	b.WriteString("## 安装包\n\n")
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

	b.WriteString("\n")
	return b.String()
}

// buildPackageRow builds a single package row
func (g *CCPageGenerator) buildPackageRow(label, repo, version, pattern string, pgVers, deps []string) string {
	// Repo badge
	repoBadge := "-"
	if repo != "" {
		repoBadge = CCRepoBadge(repo)
	}

	// Version
	verStr := "-"
	if version != "" {
		verStr = fmt.Sprintf("`%s`", version)
	}

	// Pattern
	patternStr := "-"
	if pattern != "" {
		patternStr = fmt.Sprintf("`%s`", pattern)
	}

	// PG badges
	pgBadges := g.buildPGCompatBadges(pgVers)

	// Dependencies
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

// buildPGCompatBadges builds PG version compatibility badges
func (g *CCPageGenerator) buildPGCompatBadges(supported []string) string {
	if len(supported) == 0 {
		return "-"
	}

	// Parse supported versions
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
		if supportedMap[pg] {
			badges = append(badges, CCPGBadgeGreen(fmt.Sprintf("%d", pg)))
		} else {
			badges = append(badges, CCPGBadgeRed(fmt.Sprintf("%d", pg)))
		}
	}

	return strings.Join(badges, " ")
}

// generateContribPackages generates package info for contrib extensions
func (g *CCPageGenerator) generateContribPackages(ext *Extension) string {
	var b strings.Builder

	b.WriteString("## 安装包\n\n")

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

	// Parse available versions
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
			b.WriteString(fmt.Sprintf(" %s |", CCColorBadge(version, "green")))
		} else {
			b.WriteString(fmt.Sprintf(" %s |", CCColorBadge("N/A", "red")))
		}
	}
	b.WriteString("\n\n")

	return b.String()
}

// generateAvailability generates the availability matrix
func (g *CCPageGenerator) generateAvailability(packages []*PkgInfo) string {
	if len(packages) == 0 {
		return ""
	}

	var b strings.Builder

	// Build package lookup map
	pkgMap := make(map[string]*PkgInfo)
	for _, pkg := range packages {
		key := fmt.Sprintf("%d-%s", pkg.PG, pkg.OS)
		if _, exists := pkgMap[key]; !exists {
			pkgMap[key] = pkg
		}
	}

	b.WriteString("## 可用性\n\n")

	// Header
	b.WriteString("| **发行版 / PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:---------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n")

	// Data rows
	for _, os := range g.Cache.OSVersions {
		if !os.Active {
			continue
		}

		b.WriteString(fmt.Sprintf("| [**%s**](/ext/os/%s) |", os.OS, os.OS))

		for _, pg := range g.Cache.PGVersions {
			key := fmt.Sprintf("%d-%s", pg, os.OS)

			if pkg, ok := pkgMap[key]; ok {
				cell := g.formatAvailCell(pkg)
				b.WriteString(fmt.Sprintf(" %s |", cell))
			} else {
				b.WriteString(fmt.Sprintf(" %s |", CCColorBadge("✗", "red")))
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
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
		return CCColorBadge("✗", "red")
	case "HIDE":
		return CCColorBadge("-", "gray")
	case "THROW":
		return CCColorBadge("!", "purple")
	case "BREAK":
		return CCColorBadge("!", "amber")
	case "AVAIL":
		org := ""
		if pkg.Org.Valid {
			org = strings.ToUpper(pkg.Org.String)
		}
		version := ""
		if pkg.Version.Valid {
			version = pkg.Version.String
		}

		if version != "" {
			color := "green"
			if org == "PGDG" {
				color = "blue"
			}
			return CCColorBadge(version, color)
		}
		return CCColorBadge("✓", "green")
	default:
		return CCColorBadge("-", "gray")
	}
}

// generateDownloads generates the downloads section with expandable tabs
func (g *CCPageGenerator) generateDownloads(binaries []*Binary) string {
	if len(binaries) == 0 {
		return ""
	}

	var b strings.Builder

	// Group by PG version
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

	// Sort PG versions descending
	var pgVersions []int
	for pg := range pgGroups {
		pgVersions = append(pgVersions, pg)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pgVersions)))

	b.WriteString("## 下载\n\n")

	for _, pg := range pgVersions {
		bins := pgGroups[pg]
		if len(bins) == 0 {
			continue
		}

		// Collapsible section
		b.WriteString(fmt.Sprintf("<details>\n<summary><b>PG%d</b> (%d 个包)</summary>\n\n", pg, len(bins)))

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

		b.WriteString("\n</details>\n\n")
	}

	return b.String()
}

// generateContribTip generates the contrib tip
func (g *CCPageGenerator) generateContribTip() string {
	return "> **提示**：这是 PostgreSQL 内核自带的 contrib 扩展\n\n"
}

// generateSource generates the source section
func (g *CCPageGenerator) generateSource(ext *Extension) string {
	if ext.Contrib {
		return ""
	}

	if !ext.URL.Valid && !ext.Source.Valid {
		return ""
	}

	var b strings.Builder

	b.WriteString("## 源码\n\n")

	if ext.URL.Valid && ext.URL.String != "" {
		url := ext.URL.String
		icon := "🔗"
		if strings.Contains(url, "github.com") {
			icon = "🐙"
		} else if strings.Contains(url, "gitlab.com") {
			icon = "🦊"
		}
		b.WriteString(fmt.Sprintf("- %s **仓库**: [%s](%s)\n", icon, url, url))
	}

	if ext.Source.Valid && ext.Source.String != "" {
		b.WriteString(fmt.Sprintf("- 📦 **源码包**: `%s`\n", ext.Source.String))
	}

	b.WriteString("\n")

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

	// Only show build command if either rpm or deb is true in extra
	if ext.Source.Valid && ext.Source.String != "" && (buildRPM || buildDEB) {
		var buildHint string
		if buildRPM && buildDEB {
			buildHint = "构建 RPM/DEB"
		} else if buildRPM {
			buildHint = "构建 RPM"
		} else {
			buildHint = "构建 DEB"
		}

		b.WriteString("```bash\n")
		b.WriteString(fmt.Sprintf("pig build pkg %s;  # %s\n", ext.Pkg, buildHint))
		b.WriteString("```\n\n")
	}

	return b.String()
}

// generateInstall generates the install section
func (g *CCPageGenerator) generateInstall(ext *Extension) string {
	var b strings.Builder

	b.WriteString("## 安装\n\n")

	// For contrib, just show CREATE EXTENSION
	if ext.Contrib {
		if ext.NeedDDL {
			b.WriteString("```sql\nCREATE EXTENSION " + ext.Name + ";\n```\n\n")
		}
		return b.String()
	}

	// Repo setup
	if ext.Repo.Valid && ext.Repo.String == "PGDG" {
		b.WriteString("确保 [**PGDG**](/ext/repo/pgdg) 仓库已启用：\n\n")
		b.WriteString("```bash\npig repo add pgdg -u    # 添加 PGDG 仓库并更新缓存\n```\n\n")
	} else {
		b.WriteString("确保 [**PGDG**](/ext/repo/pgdg) 和 [**PIGSTY**](/ext/repo/pgsql) 仓库已启用：\n\n")
		b.WriteString("```bash\npig repo add pgsql -u   # 添加仓库并更新缓存\n```\n\n")
	}

	// Install commands
	b.WriteString("使用 [**pig**](/ext/pig) 安装扩展：\n\n")
	b.WriteString("```bash\n")
	b.WriteString(fmt.Sprintf("pig ext install %s;          # 当前活跃 PG 版本安装\n", ext.Pkg))
	if ext.Name != ext.Pkg {
		b.WriteString(fmt.Sprintf("pig ext install %s;           # 用扩展名安装\n", ext.Name))
	}

	// Version-specific install
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
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf("pig ext install %s -v %d;    # 为 PG %d 安装\n", ext.Name, pg, pg))
		break // Only show one example
	}
	b.WriteString("```\n\n")

	// Shared preload libraries
	if ext.NeedLoad {
		b.WriteString("**预加载配置**：\n\n")

		// Collect all libs to load
		var libsToLoad []string
		libSet := make(map[string]bool)

		// Check dependencies that need loading
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

		// Add current extension
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
		b.WriteString("**创建扩展**：\n\n")
		b.WriteString("```sql\n")
		if len(ext.Requires) > 0 {
			b.WriteString(fmt.Sprintf("CREATE EXTENSION %s CASCADE;  -- 依赖: %s\n", ext.Name, strings.Join(ext.Requires, ", ")))
		} else {
			b.WriteString(fmt.Sprintf("CREATE EXTENSION %s;\n", ext.Name))
		}
		b.WriteString("```\n\n")
	} else {
		b.WriteString("> 此扩展不需要执行 `CREATE EXTENSION` 语句\n\n")
	}

	return b.String()
}

// NullBool type alias for sql.NullBool
type NullBool = sql.NullBool
