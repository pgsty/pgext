/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC List Generator - generates list pages for pigsty.cc (Chinese only)
Uses CSS classes and {.ext-table} for consistent styling
*/
package cli

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// CCListGenerator generates list pages for pigsty.cc
type CCListGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewCCListGenerator creates a new CC list generator
func NewCCListGenerator(cache *ExtensionCache, outputDir string) *CCListGenerator {
	return &CCListGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateExtensionIndexPage generates the extension index page (e/_index.md)
// If the file already exists, it is kept as-is (user may have custom content).
// If generating fresh, only the YAML frontmatter is written.
func (g *CCListGenerator) GenerateExtensionIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	content := `---
title: "扩展详情"
description: "每个 PostgreSQL 扩展的详细信息，可用性情况，下载链接，使用说明"
weight: 600
icon: fas fa-puzzle-piece
---
`
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// GenerateExtensionList generates the full extension list page
func (g *CCListGenerator) GenerateExtensionList(outputPath string) error {
	var b strings.Builder

	b.WriteString(`---
title: "扩展列表"
description: "在 Pigsty 扩展目录中收录的的 PostgreSQL 扩展列表"
weight: 120
icon: fas fa-puzzle-piece
---

`)

	// Collect all non-not-ready extensions
	var allExts []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		allExts = append(allExts, ext)
	}

	pgVersions := g.Cache.PGVersions

	// EXT stats (all extensions)
	var extAll, extPGDG, extPigsty, extContrib, extMiss int
	extPGCount := make(map[int]int)
	// PKG stats (lead non-contrib only)
	var pkgAll, pkgPGDG, pkgPigsty, pkgMiss int
	pkgPGCount := make(map[int]int)

	for _, ext := range allExts {
		repo := ""
		if ext.Repo.Valid {
			repo = ext.Repo.String
		}

		// EXT stats
		if repo != "" {
			extAll++
			switch repo {
			case "PGDG":
				extPGDG++
			case "PIGSTY":
				extPigsty++
			case "CONTRIB":
				extContrib++
			case "MIXED":
				extPGDG++
				extPigsty++
			}
			for _, v := range ext.PgVer {
				v = strings.TrimSuffix(v, "+")
				var pg int
				if _, err := fmt.Sscanf(v, "%d", &pg); err == nil {
					extPGCount[pg]++
				}
			}
		} else {
			extMiss++
		}

		// PKG stats (lead non-contrib only)
		if !ext.Lead || ext.Contrib {
			continue
		}
		if repo != "" && repo != "CONTRIB" {
			pkgAll++
			switch repo {
			case "PGDG":
				pkgPGDG++
			case "PIGSTY":
				pkgPigsty++
			case "MIXED":
				pkgPGDG++
				pkgPigsty++
			}
			for _, v := range ext.PgVer {
				v = strings.TrimSuffix(v, "+")
				var pg int
				if _, err := fmt.Sscanf(v, "%d", &pg); err == nil {
					pkgPGCount[pg]++
				}
			}
		} else if repo == "" {
			pkgMiss++
		}
	}

	// Summary text
	b.WriteString(fmt.Sprintf("在 Pigsty 扩展目录中共有 **%d** 个 PostgreSQL 扩展可用，共计 **%d** 个扩展包。\n\n",
		extAll, pkgAll))

	// Stats table - all fields centered
	b.WriteString("| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |")
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--------:|:------:|:--------:|:----------:|:-----------:|:---:|")
	for range pgVersions {
		b.WriteString(":--------:|")
	}
	b.WriteString("\n")

	// EXT row
	b.WriteString(fmt.Sprintf("| [**EXT**](/ext/list) | %d | %d | %d | %d | |",
		extAll, extPGDG, extPigsty, extContrib))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", extPGCount[pg]))
	}
	b.WriteString("\n")

	// PKG row
	b.WriteString(fmt.Sprintf("| [**PKG**](/ext/repo) | %d | %d | %d | %d | |",
		pkgAll, pkgPGDG, pkgPigsty, 0))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", pkgPGCount[pg]))
	}
	b.WriteString("\n")
	b.WriteString("{.ext-table}\n\n")

	// Per-category sections (EXT-based: all non-contrib extensions)
	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		catExts := g.Cache.CateExtMap[catKey]

		// Filter: all non-contrib extensions, sorted by ID
		var exts []*Extension
		for _, ext := range catExts {
			if ext.Contrib {
				continue
			}
			if !ext.IsReady() {
				continue
			}
			exts = append(exts, ext)
		}
		if len(exts) == 0 {
			continue
		}
		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })

		b.WriteString("\n--------\n\n")
		b.WriteString(fmt.Sprintf("## %s\n\n", catKey))

		b.WriteString("| **扩展** | **包名** | **版本** | **属性** | **仓库** | **PG大版本** | **依赖** |\n")
		b.WriteString("|:---------|:---------|:--------:|:--------:|:--------:|:-----------:|:---------|\n")

		for _, ext := range exts {
			// Column 1: 扩展名 - linked to ext page
			extLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)

			// Column 2: 包名 - linked to upstream URL
			pkgName := fmt.Sprintf("`%s`", ext.Pkg)
			if ext.URL.Valid && ext.URL.String != "" {
				pkgName = fmt.Sprintf("[`%s`](%s)", ext.Pkg, ext.URL.String)
			}

			// Column 3: 版本
			verStr := "-"
			if ext.Version.Valid && ext.Version.String != "" {
				verStr = fmt.Sprintf("`%s`", ext.Version.String)
			}

			// Column 4: 属性 (attribute flags)
			attrStr := fmt.Sprintf(`<span style="white-space:nowrap"><code>%s</code></span>`, ext.GetAttributeBadge())

			// Column 5: 仓库
			repoBadge := "-"
			if ext.Repo.Valid && ext.Repo.String != "" {
				repoBadge = CCRepoBadge(ext.Repo.String)
			}

			// Column 6: PG大版本 - pgvers shortcode
			pgStr := CCPGVersShortcode(ext.PgVer)

			// Column 7: 依赖 (requires - linked to ext pages)
			depsStr := "-"
			if len(ext.Requires) > 0 {
				depList := make([]string, len(ext.Requires))
				for i, d := range ext.Requires {
					depList[i] = fmt.Sprintf("[`%s`](/ext/e/%s)", d, d)
				}
				depsStr = strings.Join(depList, ", ")
			}

			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				extLink, pkgName, verStr, attrStr, repoBadge, pgStr, depsStr))
		}

		b.WriteString("{.ext-table}\n\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GeneratePackageList generates the package list page
func (g *CCListGenerator) GeneratePackageList(outputPath string) error {
	var b strings.Builder

	b.WriteString(`---
title: "扩展包"
linkTitle: "扩展包"
description: "按安装包名排序的扩展列表"
weight: 150
icon: fas fa-box
---

`)

	var packages []*Extension
	for _, ext := range g.Cache.Extensions {
		if ext.Lead && ext.IsReady() {
			packages = append(packages, ext)
		}
	}
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Pkg < packages[j].Pkg
	})

	b.WriteString("| **包名** | **主扩展** | **版本** | **分类** | **RPM** | **DEB** |\n")
	b.WriteString("|:---------|:-----------|:--------:|:--------:|:-------:|:-------:|\n")

	for _, ext := range packages {
		version := ext.GetVersion()

		catBadge := "-"
		if ext.Category.Valid {
			catBadge = CCCategoryBadge(ext.Category.String)
		}

		rpmBadge := ccRepoBadgeOrMiss(ext.RpmRepo)
		debBadge := ccRepoBadgeOrMiss(ext.DebRepo)

		b.WriteString(fmt.Sprintf("| `%s` | [%s](/ext/e/%s) | `%s` | %s | %s | %s |\n",
			ext.Pkg, ext.Name, ext.Name, version, catBadge, rpmBadge, debBadge))
	}

	b.WriteString("{.ext-table}\n\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateLanguageList generates the language list page
func (g *CCListGenerator) GenerateLanguageList(outputPath string) error {
	var b strings.Builder

	b.WriteString(`---
title: "编程语言"
linkTitle: "编程语言"
description: "按编程语言分类的扩展列表"
weight: 220
icon: fas fa-code
---

`)

	langMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		lang := "Unknown"
		if ext.Lang.Valid && ext.Lang.String != "" {
			lang = ext.Lang.String
		}
		langMap[lang] = append(langMap[lang], ext)
	}

	var langs []string
	for lang := range langMap {
		langs = append(langs, lang)
	}
	sort.Slice(langs, func(i, j int) bool {
		return len(langMap[langs[i]]) > len(langMap[langs[j]])
	})

	// Horizontal summary table
	b.WriteString("| **语言** |")
	for _, lang := range langs {
		b.WriteString(fmt.Sprintf(" %s |", CCLanguageBadge(lang)))
	}
	b.WriteString("\n|:------:|")
	for range langs {
		b.WriteString(":------:|")
	}
	b.WriteString("\n| **扩展数** |")
	for _, lang := range langs {
		b.WriteString(fmt.Sprintf(" %d |", len(langMap[lang])))
	}
	b.WriteString("\n{.ext-table}\n\n")

	for _, lang := range langs {
		exts := langMap[lang]
		if len(exts) == 0 {
			continue
		}

		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })

		anchor := strings.ToLower(lang)
		if anchor == "c++" {
			anchor = "cpp"
		}

		b.WriteString(fmt.Sprintf("\n--------\n\n"))
		b.WriteString(fmt.Sprintf("## %s {#%s}\n\n", lang, anchor))
		b.WriteString(fmt.Sprintf("%s 语言编写的扩展（%d 个）\n\n", lang, len(exts)))

		isRust := lang == "Rust"
		if isRust {
			b.WriteString("| **扩展** | **包** | **版本** | **PGRX** | **分类** | **许可证** | **描述** |\n")
			b.WriteString("|:---------|:-------|:--------:|:--------:|:--------:|:----------:|:---------|\n")
		} else {
			b.WriteString("| **扩展** | **包** | **版本** | **分类** | **许可证** | **描述** |\n")
			b.WriteString("|:---------|:-------|:--------:|:--------:|:----------:|:---------|\n")
		}

		for _, ext := range exts {
			catBadge := "-"
			if ext.Category.Valid {
				catBadge = CCCategoryBadge(ext.Category.String)
			}
			license := "-"
			if ext.License.Valid && ext.License.String != "" {
				license = CCLicenseBadge(ext.License.String)
			}
			version := "-"
			if ext.Version.Valid && ext.Version.String != "" {
				version = ext.Version.String
			}
			desc := SanitizeText(ext.GetZhDesc())
			pkgLink := fmt.Sprintf("`%s`", ext.Pkg)
			if ext.URL.Valid && ext.URL.String != "" {
				pkgLink = fmt.Sprintf("[`%s`](%s)", ext.Pkg, ext.URL.String)
			}

			if isRust {
				pgrx := "-"
				if ext.Extra != nil {
					if v, ok := ext.Extra["pgrx"]; ok {
						pgrx = fmt.Sprintf("`%v`", v)
					}
				}
				b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s | %s |\n",
					ext.Name, ext.Name, pkgLink, version, pgrx, catBadge, license, desc))
			} else {
				b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
					ext.Name, ext.Name, pkgLink, version, catBadge, license, desc))
			}
		}

		b.WriteString("{.ext-table}\n\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateLicenseList generates the license list page
func (g *CCListGenerator) GenerateLicenseList(outputPath string) error {
	var b strings.Builder

	b.WriteString(`---
title: "开源协议"
linkTitle: "开源协议"
description: "按开源许可证分类的扩展列表"
weight: 230
icon: fas fa-scale-balanced
---

`)

	licenseMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		license := "Unknown"
		if ext.License.Valid && ext.License.String != "" {
			license = ext.License.String
		}
		licenseMap[license] = append(licenseMap[license], ext)
	}

	// Fixed license ordering: permissive then copyleft
	permissive := []string{"MIT", "ISC", "PostgreSQL", "BSD 0-Clause", "BSD 2-Clause", "BSD 3-Clause", "Artistic", "Apache-2.0", "MPL-2.0"}
	copyleft := []string{"GPL-2.0", "GPL-3.0", "LGPL-2.1", "LGPL-3.0", "AGPL-3.0", "Timescale"}

	// Collect any licenses not in the predefined lists
	known := make(map[string]bool)
	for _, lic := range permissive {
		known[lic] = true
	}
	for _, lic := range copyleft {
		known[lic] = true
	}
	var extra []string
	for lic := range licenseMap {
		if !known[lic] {
			extra = append(extra, lic)
		}
	}
	sort.Strings(extra)
	copyleft = append(copyleft, extra...)

	// Filter to only licenses that actually have extensions
	filterPresent := func(lics []string) []string {
		var out []string
		for _, lic := range lics {
			if len(licenseMap[lic]) > 0 {
				out = append(out, lic)
			}
		}
		return out
	}
	permissive = filterPresent(permissive)
	copyleft = filterPresent(copyleft)

	// Combined ordered list for detail sections
	allLicenses := append(append([]string{}, permissive...), copyleft...)

	// Single summary table: row 1 = permissive, row 2 = copyleft
	maxCols := len(permissive)
	if len(copyleft) > maxCols {
		maxCols = len(copyleft)
	}

	// Header row
	b.WriteString("| **许可证** |")
	for _, lic := range permissive {
		b.WriteString(fmt.Sprintf(" %s |", CCLicenseBadge(lic)))
	}
	for i := len(permissive); i < maxCols; i++ {
		b.WriteString(" |")
	}
	b.WriteString("\n|:------:|")
	for i := 0; i < maxCols; i++ {
		b.WriteString(":------:|")
	}
	// Permissive count row
	b.WriteString("\n| **宽松** |")
	for _, lic := range permissive {
		b.WriteString(fmt.Sprintf(" %d |", len(licenseMap[lic])))
	}
	for i := len(permissive); i < maxCols; i++ {
		b.WriteString(" |")
	}
	// Copyleft license row
	b.WriteString("\n| **许可证** |")
	for _, lic := range copyleft {
		b.WriteString(fmt.Sprintf(" %s |", CCLicenseBadge(lic)))
	}
	for i := len(copyleft); i < maxCols; i++ {
		b.WriteString(" |")
	}
	// Copyleft count row
	b.WriteString("\n| **严格** |")
	for _, lic := range copyleft {
		b.WriteString(fmt.Sprintf(" %d |", len(licenseMap[lic])))
	}
	for i := len(copyleft); i < maxCols; i++ {
		b.WriteString(" |")
	}
	b.WriteString("\n{.ext-table}\n\n")

	// Per-license detail sections
	for _, lic := range allLicenses {
		exts := licenseMap[lic]
		if len(exts) == 0 {
			continue
		}

		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })

		anchor := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(lic, ".", ""), "-", ""), " ", ""))

		b.WriteString("\n--------\n\n")
		b.WriteString(fmt.Sprintf("## %s {#%s}\n\n", lic, anchor))
		b.WriteString(fmt.Sprintf("使用 %s 许可证的扩展（%d 个）\n\n", lic, len(exts)))

		b.WriteString("| **扩展** | **包** | **版本** | **分类** | **语言** | **描述** |\n")
		b.WriteString("|:---------|:-------|:--------:|:--------:|:--------:|:---------|\n")

		for _, ext := range exts {
			catBadge := "-"
			if ext.Category.Valid {
				catBadge = CCCategoryBadge(ext.Category.String)
			}
			lang := "-"
			if ext.Lang.Valid && ext.Lang.String != "" {
				lang = CCLanguageBadge(ext.Lang.String)
			}
			version := "-"
			if ext.Version.Valid && ext.Version.String != "" {
				version = ext.Version.String
			}
			desc := SanitizeText(ext.GetZhDesc())
			pkgLink := fmt.Sprintf("`%s`", ext.Pkg)
			if ext.URL.Valid && ext.URL.String != "" {
				pkgLink = fmt.Sprintf("[`%s`](%s)", ext.Pkg, ext.URL.String)
			}

			b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
				ext.Name, ext.Name, pkgLink, version, catBadge, lang, desc))
		}

		b.WriteString("{.ext-table}\n\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateRepoList generates the repository list page
func (g *CCListGenerator) GenerateRepoList(outputPath string) error {
	var b strings.Builder

	b.WriteString(`---
title: "归属仓库"
linkTitle: "归属仓库"
description: "按归属仓库分类的扩展列表"
weight: 210
icon: fas fa-warehouse
---

`)

	// Collect all non-not-ready extensions, group by Repo field
	var pgdgExts, pigstyExts, mixedExts, contribExts []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		repo := ""
		if ext.Repo.Valid {
			repo = ext.Repo.String
		}
		switch repo {
		case "PGDG":
			pgdgExts = append(pgdgExts, ext)
		case "PIGSTY":
			pigstyExts = append(pigstyExts, ext)
		case "MIXED":
			mixedExts = append(mixedExts, ext)
		case "CONTRIB":
			contribExts = append(contribExts, ext)
		}
	}

	// Sort all groups by ID
	for _, exts := range [][]*Extension{pgdgExts, pigstyExts, mixedExts, contribExts} {
		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })
	}

	// Horizontal summary table
	total := len(pgdgExts) + len(pigstyExts) + len(mixedExts) + len(contribExts)
	b.WriteString(fmt.Sprintf("| **仓库** | %s | %s | %s | %s | **总计** |\n",
		CCRepoBadge("PGDG"), CCRepoBadge("PIGSTY"), CCRepoBadge("MIXED"), CCRepoBadge("CONTRIB")))
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:------:|\n")
	b.WriteString(fmt.Sprintf("| **扩展数量** | %d | %d | %d | %d | **%d** |\n",
		len(pgdgExts), len(pigstyExts), len(mixedExts), len(contribExts), total))
	b.WriteString("{.ext-table}\n\n")

	// PGDG section
	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf("## PGDG 仓库 {#pgdg}\n\nPostgreSQL 全球开发组提供的二进制制成品仓库，包括 %d 个扩展。\n\n", len(pgdgExts)))
	writeRepoExtTable(&b, pgdgExts)

	// PIGSTY section
	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf("## PIGSTY 仓库 {#pigsty}\n\nPigsty 扩展仓库在 PGDG 仓库的基础上，额外提供了 %d 个扩展。\n\n", len(pigstyExts)))
	writeRepoExtTable(&b, pigstyExts)

	// MIXED section
	if len(mixedExts) > 0 {
		b.WriteString("\n--------\n\n")
		b.WriteString(fmt.Sprintf("## 两者皆有 {#mixed}\n\n有一些扩展在 PIGSTY 和 PGDG 仓库中同时提供（%d 个）。\n\n", len(mixedExts)))
		writeRepoExtTable(&b, mixedExts)
	}

	// CONTRIB section
	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf("## CONTRIB {#contrib}\n\nPostgreSQL 内置的 Contrib 扩展模块，总计 %d 个扩展，PG 随内核 contrib 包交付。\n\n", len(contribExts)))
	writeContribExtTable(&b, contribExts)

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// writeRepoExtTable writes a 7-column table for PGDG/PIGSTY/MIXED sections
func writeRepoExtTable(b *strings.Builder, exts []*Extension) {
	b.WriteString("| **扩展** | **RPM仓库** | **RPM包** | **RPM版本** | **DEB仓库** | **DEB包** | **DEB版本** |\n")
	b.WriteString("|:---------|:----------:|:---------|:----------:|:----------:|:---------|:----------:|\n")

	for _, ext := range exts {
		rpmRepo := "-"
		if ext.RpmRepo.Valid && ext.RpmRepo.String != "" {
			rpmRepo = CCRepoBadge(ext.RpmRepo.String)
		}
		rpmPkg := "-"
		if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
			rpmPkg = fmt.Sprintf("`%s`", ext.RpmPkg.String)
		}
		rpmVer := "-"
		if ext.RpmVer.Valid && ext.RpmVer.String != "" {
			rpmVer = fmt.Sprintf("`%s`", ext.RpmVer.String)
		}
		debRepo := "-"
		if ext.DebRepo.Valid && ext.DebRepo.String != "" {
			debRepo = CCRepoBadge(ext.DebRepo.String)
		}
		debPkg := "-"
		if ext.DebPkg.Valid && ext.DebPkg.String != "" {
			debPkg = fmt.Sprintf("`%s`", ext.DebPkg.String)
		}
		debVer := "-"
		if ext.DebVer.Valid && ext.DebVer.String != "" {
			debVer = fmt.Sprintf("`%s`", ext.DebVer.String)
		}

		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s | %s | %s | %s | %s |\n",
			ext.Name, ext.Name, rpmRepo, rpmPkg, rpmVer, debRepo, debPkg, debVer))
	}

	b.WriteString("{.ext-table}\n\n")
}

// writeContribExtTable writes the CONTRIB section table
func writeContribExtTable(b *strings.Builder, exts []*Extension) {
	b.WriteString("| **扩展** | **分类** | **版本** | **PG版本** | **说明** |\n")
	b.WriteString("|:---------|:------:|:--------:|:---------|:---------|\n")

	for _, ext := range exts {
		catBadge := "-"
		if ext.Category.Valid {
			catBadge = CCCategoryBadge(ext.Category.String)
		}
		version := ext.GetVersion()
		if version != "-" {
			version = fmt.Sprintf("`%s`", version)
		}
		pgVer := "-"
		if len(ext.PgVer) > 0 {
			var vers []string
			for _, v := range ext.PgVer {
				vers = append(vers, strings.TrimSuffix(v, "+"))
			}
			pgVer = strings.Join(vers, ", ")
		}
		desc := SanitizeText(ext.GetZhDesc())

		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s | %s | %s |\n",
			ext.Name, ext.Name, catBadge, version, pgVer, desc))
	}

	b.WriteString("{.ext-table}\n\n")
}

// platformListConfig configures a platform-specific package list page (RPM or DEB)
type platformListConfig struct {
	Name      string // "RPM" or "DEB"
	OSDesc    string // "EL" or "Debian/Ubuntu"
	Title     string
	Desc      string
	Icon      string
	Weight    int
	OtherName string // "DEB" or "RPM"
	// Field accessors
	GetRepo func(*Extension) sql.NullString
	GetPkg  func(*Extension) sql.NullString
	GetVer  func(*Extension) sql.NullString
	GetPG   func(*Extension) []string
	GetDeps func(*Extension) []string
	// Cross-platform accessor
	GetOtherRepo func(*Extension) sql.NullString
}

// GenerateRPMList generates the RPM package list page
func (g *CCListGenerator) GenerateRPMList(outputPath string) error {
	return g.generatePlatformList(outputPath, &platformListConfig{
		Name:         "RPM",
		OSDesc:       "EL",
		Title:        "RPM 列表",
		Desc:         "在 EL 系统上可用的 PostgreSQL 扩展 RPM 二进制包",
		Icon:         "fa-brands fa-redhat",
		Weight:       130,
		OtherName:    "DEB",
		GetRepo:      func(e *Extension) sql.NullString { return e.RpmRepo },
		GetPkg:       func(e *Extension) sql.NullString { return e.RpmPkg },
		GetVer:       func(e *Extension) sql.NullString { return e.RpmVer },
		GetPG:        func(e *Extension) []string { return e.RpmPg },
		GetDeps:      func(e *Extension) []string { return e.RpmDeps },
		GetOtherRepo: func(e *Extension) sql.NullString { return e.DebRepo },
	})
}

// GenerateDEBList generates the DEB package list page
func (g *CCListGenerator) GenerateDEBList(outputPath string) error {
	return g.generatePlatformList(outputPath, &platformListConfig{
		Name:         "DEB",
		OSDesc:       "Debian/Ubuntu",
		Title:        "DEB 列表",
		Desc:         "在 Debian/Ubuntu 系统上可用的 PostgreSQL 扩展 DEB 二进制包",
		Icon:         "fa-brands fa-debian",
		Weight:       140,
		OtherName:    "RPM",
		GetRepo:      func(e *Extension) sql.NullString { return e.DebRepo },
		GetPkg:       func(e *Extension) sql.NullString { return e.DebPkg },
		GetVer:       func(e *Extension) sql.NullString { return e.DebVer },
		GetPG:        func(e *Extension) []string { return e.DebPg },
		GetDeps:      func(e *Extension) []string { return e.DebDeps },
		GetOtherRepo: func(e *Extension) sql.NullString { return e.RpmRepo },
	})
}

// generatePlatformList generates a platform-specific package list page
func (g *CCListGenerator) generatePlatformList(outputPath string, cfg *platformListConfig) error {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("---\ntitle: \"%s\"\ndescription: \"%s\"\nweight: %d\nicon: %s\n---\n\n",
		cfg.Title, cfg.Desc, cfg.Weight, cfg.Icon))

	// Collect all non-not-ready extensions
	var allExts []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		allExts = append(allExts, ext)
	}

	pgVersions := g.Cache.PGVersions

	// Build PG version set for quick lookup
	pgVerSet := make(map[string]bool)

	// EXT stats
	var extAll, extPGDG, extPigsty, extContrib, extMiss int
	extPGCount := make(map[int]int)
	// PKG stats (lead only)
	var pkgAll, pkgPGDG, pkgPigsty, pkgMiss int
	pkgPGCount := make(map[int]int)
	// Cross-platform stats (lead only, non-contrib)
	var exclusive, otherMissing int

	for _, ext := range allExts {
		repo := cfg.GetRepo(ext)
		hasRepo := repo.Valid && repo.String != ""
		otherRepo := cfg.GetOtherRepo(ext)
		hasOther := otherRepo.Valid && otherRepo.String != ""

		// EXT stats
		if hasRepo || ext.Contrib {
			extAll++
			if !ext.Contrib {
				if repo.String == "PGDG" {
					extPGDG++
				} else if repo.String == "PIGSTY" {
					extPigsty++
				}
			}
			if ext.Contrib {
				extContrib++
			}
			// PG versions for EXT
			pgs := cfg.GetPG(ext)
			if ext.Contrib && len(pgs) == 0 {
				pgs = ext.PgVer
			}
			for _, v := range pgs {
				v = strings.TrimSuffix(v, "+")
				var pg int
				if _, err := fmt.Sscanf(v, "%d", &pg); err == nil {
					extPGCount[pg]++
				}
			}
		} else {
			extMiss++
		}

		// PKG stats (lead only)
		if !ext.Lead {
			continue
		}
		if hasRepo {
			pkgAll++
			if repo.String == "PGDG" {
				pkgPGDG++
			} else if repo.String == "PIGSTY" {
				pkgPigsty++
			}
			for _, v := range cfg.GetPG(ext) {
				v = strings.TrimSuffix(v, "+")
				pgVerSet[v] = true
				var pg int
				if _, err := fmt.Sscanf(v, "%d", &pg); err == nil {
					pkgPGCount[pg]++
				}
			}
		}
		if !ext.Contrib {
			if hasRepo && !hasOther {
				exclusive++
			}
			if hasOther && !hasRepo {
				otherMissing++
			}
			if !hasRepo {
				pkgMiss++
			}
		}
	}

	// Summary text - two lines separated by blank line
	b.WriteString(fmt.Sprintf("在 %s 系统上共有 **%d** 个 PostgreSQL 扩展可用，共计 **%d** 个扩展包。\n\n",
		cfg.OSDesc, extAll, pkgAll))
	b.WriteString(fmt.Sprintf("其中 **%d** 个扩展包是 %s 独有，**%d** 个 %s 扩展包缺少对应的 %s 包。\n\n",
		exclusive, cfg.OSDesc, otherMissing, cfg.OtherName, cfg.Name))

	// Stats table - all fields centered
	extLink := fmt.Sprintf("/ext/%s", strings.ToLower(cfg.Name))
	b.WriteString("| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |")
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--------:|:------:|:--------:|:----------:|:-----------:|:---:|")
	for range pgVersions {
		b.WriteString(":--------:|")
	}
	b.WriteString("\n")

	// EXT row
	b.WriteString(fmt.Sprintf("| [**EXT**](%s) | %d | %d | %d | %d | |",
		extLink, extAll, extPGDG, extPigsty, extContrib))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", extPGCount[pg]))
	}
	b.WriteString("\n")

	// PKG row
	b.WriteString(fmt.Sprintf("| [**PKG**](/ext/repo) | %d | %d | %d | %d | |",
		pkgAll, pkgPGDG, pkgPigsty, 0))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", pkgPGCount[pg]))
	}
	b.WriteString("\n")
	b.WriteString("{.ext-table}\n\n")

	// Per-category sections
	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		catExts := g.Cache.CateExtMap[catKey]

		// Filter: lead extensions with repo set, exclude contrib, sorted by ID
		var pkgs []*Extension
		for _, ext := range catExts {
			if !ext.Lead || ext.Contrib {
				continue
			}
			if !ext.IsReady() {
				continue
			}
			repo := cfg.GetRepo(ext)
			if repo.Valid && repo.String != "" {
				pkgs = append(pkgs, ext)
			}
		}
		if len(pkgs) == 0 {
			continue
		}
		sort.Slice(pkgs, func(i, j int) bool { return pkgs[i].ID < pkgs[j].ID })

		b.WriteString("\n--------\n\n")
		b.WriteString(fmt.Sprintf("## %s\n\n", catKey))

		b.WriteString("| **扩展包** | **版本** | **仓库** | **包名** | **PG大版本** | **依赖** |\n")
		b.WriteString("|:----------|:--------:|:--------:|:---------|:-----------:|:---------|\n")

		for _, ext := range pkgs {
			// Column 1: 扩展包 - pkg name linked to upstream URL
			pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Pkg, ext.Name)

			// Column 2: 版本
			ver := cfg.GetVer(ext)
			verStr := "-"
			if ver.Valid && ver.String != "" {
				verStr = fmt.Sprintf("`%s`", ver.String)
			}

			// Column 3: 仓库
			repo := cfg.GetRepo(ext)
			repoBadge := "-"
			if repo.Valid && repo.String != "" {
				repoBadge = CCRepoBadge(repo.String)
			}

			// Column 4: 包名模式
			pkg := cfg.GetPkg(ext)
			pkgName := "-"
			if pkg.Valid && pkg.String != "" {
				pkgName = fmt.Sprintf("`%s`", pkg.String)
			}

			// Column 5: PG大版本 - pgvers shortcode
			pgs := cfg.GetPG(ext)
			if ext.Contrib && len(pgs) == 0 {
				pgs = ext.PgVer
			}
			pgStr := CCPGVersShortcode(pgs)

			// Column 6: 依赖
			deps := cfg.GetDeps(ext)
			depsStr := "-"
			if len(deps) > 0 {
				depList := make([]string, len(deps))
				for i, d := range deps {
					depList[i] = fmt.Sprintf("`%s`", d)
				}
				depsStr = strings.Join(depList, ", ")
			}

			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
				pkgLink, verStr, repoBadge, pkgName, pgStr, depsStr))
		}

		b.WriteString("{.ext-table}\n\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateOSIndexPage generates the OS index page (os/_index.md)
// If the file already exists, it is kept as-is (user may have custom content).
// If generating fresh, only the YAML frontmatter is written.
func (g *CCListGenerator) GenerateOSIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	content := `---
title: "操作系统"
linkTitle: "操作系统"
description: "各操作系统的扩展可用性"
weight: 400
icon: fa-brands fa-linux
---
`
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// GenerateOverviewPage generates the ext/_index.md overview page with dynamic data
func (g *CCListGenerator) GenerateOverviewPage(outputPath string) error {
	var b strings.Builder

	// ── Compute statistics ────────────────────────────────────────────
	type rowStats struct {
		All, PGDG, PIGSTY, CONTRIB, MISS int
		PGCounts                          map[int]int
	}
	allRow := rowStats{PGCounts: make(map[int]int)}
	elRow := rowStats{PGCounts: make(map[int]int)}
	debRow := rowStats{PGCounts: make(map[int]int)}

	pgContains := func(slice []string, ver int) bool {
		s := fmt.Sprintf("%d", ver)
		for _, v := range slice {
			if v == s {
				return true
			}
		}
		return false
	}

	for _, ext := range g.Cache.Extensions {
		if !ext.IsReady() {
			continue
		}
		repo := ""
		if ext.Repo.Valid {
			repo = ext.Repo.String
		}

		// ALL row
		allRow.All++
		switch repo {
		case "PGDG":
			allRow.PGDG++
		case "PIGSTY":
			allRow.PIGSTY++
		case "MIXED":
			allRow.PGDG++
			allRow.PIGSTY++
		case "CONTRIB":
			allRow.CONTRIB++
		}
		for _, pg := range g.Cache.PGVersions {
			if pgContains(ext.PgVer, pg) {
				allRow.PGCounts[pg]++
			}
		}

		// EL row: has RPM or is contrib
		hasRPM := ext.RpmRepo.Valid && ext.RpmRepo.String != ""
		if hasRPM || ext.Contrib {
			elRow.All++
			if hasRPM {
				switch ext.RpmRepo.String {
				case "PGDG":
					elRow.PGDG++
				case "PIGSTY":
					elRow.PIGSTY++
				}
			}
			if ext.Contrib {
				elRow.CONTRIB++
			}
			for _, pg := range g.Cache.PGVersions {
				if hasRPM && pgContains(ext.RpmPg, pg) {
					elRow.PGCounts[pg]++
				} else if ext.Contrib && pgContains(ext.PgVer, pg) {
					elRow.PGCounts[pg]++
				}
			}
		}

		// Debian row: has DEB or is contrib
		hasDEB := ext.DebRepo.Valid && ext.DebRepo.String != ""
		if hasDEB || ext.Contrib {
			debRow.All++
			if hasDEB {
				switch ext.DebRepo.String {
				case "PGDG":
					debRow.PGDG++
				case "PIGSTY":
					debRow.PIGSTY++
				}
			}
			if ext.Contrib {
				debRow.CONTRIB++
			}
			for _, pg := range g.Cache.PGVersions {
				if hasDEB && pgContains(ext.DebPg, pg) {
					debRow.PGCounts[pg]++
				} else if ext.Contrib && pgContains(ext.PgVer, pg) {
					debRow.PGCounts[pg]++
				}
			}
		}
	}
	elRow.MISS = allRow.All - elRow.All
	debRow.MISS = allRow.All - debRow.All

	totalExts := allRow.All

	// ── Per-category package lists (PKG logic: lead, non-contrib, ordered by ID) ──
	type cateInfo struct {
		Name string
		Pkgs string
	}
	var cateInfos []cateInfo
	for _, cat := range g.Cache.Categories {
		cateName := strings.ToUpper(cat.Name)
		exts := g.Cache.CateExtMap[cateName]
		var pkgLinks []string
		for _, ext := range exts {
			if !ext.Lead || ext.Contrib {
				continue
			}
			pkgLinks = append(pkgLinks, fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Pkg, ext.Name))
		}
		cateInfos = append(cateInfos, cateInfo{Name: cateName, Pkgs: strings.Join(pkgLinks, "  ")})
	}

	// ── Count active OS variants ──────────────────────────────────────
	activeOSCount := 0
	for _, osv := range g.Cache.OSVersions {
		if osv.Active {
			activeOSCount++
		}
	}

	// ── PG version badge helper ───────────────────────────────────────
	pgVerBadges := func() string {
		return CCPGVersAllShortcode(g.Cache.PGVersions)
	}

	// ── Stats row helper ──────────────────────────────────────────────
	writeStatsRow := func(label string, r rowStats) {
		b.WriteString(fmt.Sprintf("| **%s** | %d | %d | %d | %d | |", label, r.All, r.PGDG, r.PIGSTY, r.CONTRIB))
		for _, pg := range g.Cache.PGVersions {
			b.WriteString(fmt.Sprintf(" %d |", r.PGCounts[pg]))
		}
		b.WriteString("\n")
	}

	// ══════════════════════════════════════════════════════════════════
	// YAML Front Matter
	// ══════════════════════════════════════════════════════════════════
	b.WriteString(fmt.Sprintf(`---
title: Pigsty 扩展目录
linkTitle: Pigsty 扩展目录
description: >
  扩展是 PostgreSQL 的灵魂所在，Pigsty 制作、收录、整合了 %d 个 PG 生态扩展，供用户开箱即用。
weight: 20
cascade:
  type: docs
---

`, totalExts))

	// ══════════════════════════════════════════════════════════════════
	// Introduction
	// ══════════════════════════════════════════════════════════════════
	b.WriteString(fmt.Sprintf(`Pigsty 提供了以下三样基础设施，帮助用户更好的利用 PostgreSQL 扩展生态系统的协同超能力：

- [**扩展目录**](/ext/list)：查阅 [**%d**](/ext/list) 个扩展插件的详细信息，使用方法，元数据，下载链接与文档
- [**扩展仓库**](/docs/repo/pgsql)：获取预先打包的 RPM/DEB 二进制包，在 [**%d 个 Linux 系统**](/ext/os) 上可用
- [**包管理器**](/docs/pig)：使用 [**`+"`"+`pig`+"`"+`**](/docs/pig) 命令行工具，屏蔽复杂度与操作系统与架构差异

`, totalExts, activeOSCount))
	b.WriteString("```bash\n")
	b.WriteString("curl -fsSL https://repo.pigsty.cc/pig | bash  # 安装 pig 命令行工具\n")
	b.WriteString("pig repo add pigsty pgdg -u                   # 在您的 Linux 发行版上配置软件仓库\n")
	b.WriteString("pig install pg18                              # 从 PGDG 官方仓库安装 PostgreSQL 18 内核包\n")
	b.WriteString("pig install pg_duckdb -v 18                   # 例：针对 PG 18 安装 pg_duckdb\n")
	b.WriteString("```\n\n")
	b.WriteString("一切都皆可用 PostgreSQL 解决！请参阅我们的博客文章：[**PostgreSQL 正在吞噬数据库世界！**](/blog/pg/pg-eat-db-world)\n\n")
	b.WriteString("![](/img/pigsty/ecosystem.png)\n\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 核心特点
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("--------\n\n## 核心特点\n\n")
	b.WriteString(fmt.Sprintf("- **数量**：无与伦比的扩展数量：**%d** 个可用扩展，为 PG 扩展生态之最\n", totalExts))
	b.WriteString("- **质量**：原生 Linux RPM/DEB 包，完全兼容 PGDG 打包规范\n")
	b.WriteString("- **易用**：提供包管理器 [**`pig`**](/docs/pig)，屏蔽操作系统与架构差异，开箱即用\n")
	b.WriteString("- **兼容**：扩展完全兼容 PGDG 打包规范，可与 PGDG 仓库无缝混用\n")
	b.WriteString("- **分发**：由 Cloudflare CDN 进行全球仓库分发，提供国内 CDN 加速\n")
	b.WriteString("- **开源**：完全开源，提供便利的构建工具，免费对公众提供服务的软件基础设施\n\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 扩展统计
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("--------\n\n## 扩展统计\n\n")
	b.WriteString("| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:---:|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n")
	writeStatsRow("ALL", allRow)
	writeStatsRow("EL", elRow)
	writeStatsRow("Debian", debRow)
	b.WriteString("{.ext-table}\n\n")
	b.WriteString("> 详见：[扩展列表](/ext/list)，[RPM 列表](/ext/rpm)，[DEB 列表](/ext/deb)，[归属仓库](/ext/repo)\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 扩展分类
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("--------\n\n## 扩展分类\n\n\n")
	b.WriteString("| **分类** | **扩展包** |\n")
	b.WriteString("|:--|:--|\n")
	for _, ci := range cateInfos {
		b.WriteString(fmt.Sprintf("| %s | %s |\n", CCCategoryBadge(ci.Name), ci.Pkgs))
	}
	b.WriteString("{.ext-table}\n\n")

	// License badges
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
		CCLicenseBadge("MIT"), CCLicenseBadge("ISC"), CCLicenseBadge("PostgreSQL"),
		CCLicenseBadge("BSD 0-Clause"), CCLicenseBadge("BSD 2-Clause"), CCLicenseBadge("BSD 3-Clause"),
		CCLicenseBadge("Artistic")))
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n")
	b.WriteString("{.full-width}\n\n")
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n",
		CCLicenseBadge("Apache-2.0"), CCLicenseBadge("MPL-2.0"),
		CCLicenseBadge("GPL-2.0"), CCLicenseBadge("GPL-3.0"),
		CCLicenseBadge("LGPL-2.1"), CCLicenseBadge("LGPL-3.0"),
		CCLicenseBadge("AGPL-3.0"), CCLicenseBadge("Timescale")))
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n")
	b.WriteString("{.full-width}\n\n")

	// Language badges
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
		CCLanguageBadge("C"), CCLanguageBadge("C++"), CCLanguageBadge("Rust"),
		CCLanguageBadge("Java"), CCLanguageBadge("Python"), CCLanguageBadge("SQL"),
		CCLanguageBadge("Data")))
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n")
	b.WriteString("{.full-width}\n\n\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 兼容系统
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("## 兼容系统\n\n")
	b.WriteString("| **系统** | **版本** | **x86_64** | **aarch64** | **PG版本** |\n")
	b.WriteString("|:------:|:------:|:----------:|:-----------:|:--------:|\n")

	// Group OS versions by prefix (el8, el9, d12, u22, ...) derived from OS field
	type osPair struct {
		Prefix, SysName, Version, X86, ARM string
	}
	osPairMap := make(map[string]*osPair)
	var osPrefixOrder []string
	for _, osv := range g.Cache.OSVersions {
		if !osv.Active {
			continue
		}
		// Derive prefix from OS field: "el8.x86_64" → "el8"
		prefix := strings.Split(osv.OS, ".")[0]
		if _, exists := osPairMap[prefix]; !exists {
			// Derive system name from OSFull by stripping arch suffix
			sysName := osv.OSFull
			for _, suffix := range []string{" x86_64", " aarch64", " x86", " ARM"} {
				sysName = strings.TrimSuffix(sysName, suffix)
			}
			// Derive version: for Ubuntu use "XX.04" format, otherwise use OSMajor
			version := osv.OSMajor
			if strings.HasPrefix(prefix, "u") && len(osv.OSMajor) == 2 {
				version = osv.OSMajor + ".04"
			}
			osPairMap[prefix] = &osPair{Prefix: prefix, SysName: sysName, Version: version}
			osPrefixOrder = append(osPrefixOrder, prefix)
		}
		pair := osPairMap[prefix]
		if strings.HasSuffix(osv.OS, ".x86_64") || strings.Contains(osv.OSArch, "x86_64") {
			pair.X86 = osv.OS
		} else {
			pair.ARM = osv.OS
		}
	}
	for _, prefix := range osPrefixOrder {
		pair := osPairMap[prefix]
		x86Link := "-"
		if pair.X86 != "" {
			x86Link = fmt.Sprintf("[`%s`](/ext/os/%s)", pair.X86, pair.X86)
		}
		armLink := "-"
		if pair.ARM != "" {
			armLink = fmt.Sprintf("[`%s`](/ext/os/%s)", pair.ARM, pair.ARM)
		}
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s | %s |\n",
			pair.SysName, pair.Version, x86Link, armLink, pgVerBadges()))
	}
	b.WriteString("{.ext-table}\n\n")
	b.WriteString("> 详见：[操作系统](/ext/os)，[RPM 列表](/ext/rpm)，[DEB 列表](/ext/deb)\n\n\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 合作伙伴
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("--------\n\n## 合作伙伴\n\n")
	b.WriteString("Pigsty 向用户提供无可比拟的 PostgreSQL 扩展交付体验，已有多家 PostgreSQL Vendor 使用它进行交付。\n\n")
	b.WriteString("{{< cardpane >}}\n")
	b.WriteString("{{< card header=\"[**Pigsty**](https://github.com/pgsty/pigsty)\" >}}\n")
	b.WriteString("开箱即用的开源企业级 PostgreSQL RDS 发行版\n")
	b.WriteString("{{< /card >}}\n")
	b.WriteString("{{< card header=\"[**Omnigres**](https://docs.omnigres.org/quick_start/)\" >}}\n")
	b.WriteString("PostgreSQL as a Platform，在数据库中进行应用开发\n")
	b.WriteString("{{< /card >}}\n")
	b.WriteString("{{< card header=\"[**AutoBase**](https://autobase.tech/docs/extensions/install)\" >}}\n")
	b.WriteString("基于 Ansible 的 PG 集群自动化部署，开源 DBaaS\n")
	b.WriteString("{{< /card >}}\n")
	b.WriteString("{{< /cardpane >}}\n\n")
	b.WriteString("{{< cardpane >}}\n")
	b.WriteString("{{< card header=\"[**TensorChord**](https://github.com/tensorchord)\" >}}\n")
	b.WriteString("云原生 AI 基础设施，开发了多个知名 PG 扩展\n")
	b.WriteString("{{< /card >}}\n")
	b.WriteString("{{< card header=\"[**文武IT**](https://w3.ww-it.cn/)\" >}}\n")
	b.WriteString("企业级 PostgreSQL 服务商\n")
	b.WriteString("{{< /card >}}\n")
	b.WriteString("{{< /cardpane >}}\n\n\n")

	// ══════════════════════════════════════════════════════════════════
	// 关于项目
	// ══════════════════════════════════════════════════════════════════
	b.WriteString("--------\n\n## 关于项目\n\n")
	b.WriteString("本项目由 [**PGSTY**](https://github.com/pgsty) / [**VONNG**](https://vonng.com) 开发维护，并使用 [**Apache 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。\n\n")
	b.WriteString("| **仓库** | **说明** |\n")
	b.WriteString("|:------|:------|\n")
	b.WriteString("| [pgsty](https://github.com/pgsty) | PGSTY 组织主页 |\n")
	b.WriteString("| [pgsty/pgext](https://github.com/pgsty/pgext) | 本网站，扩展元数据，以及管理工具 |\n")
	b.WriteString("| [pgsty/pigsty](https://github.com/pgsty/pigsty) | PostgreSQL 数据库发行版 |\n")
	b.WriteString("| [pgsty/pig](https://github.com/pgsty/pig) | PostgreSQL 包管理器 |\n")
	b.WriteString("| [pgsty/rpm](https://github.com/pgsty/rpm) | RPM 构建源代码 |\n")
	b.WriteString("| [pgsty/deb](https://github.com/pgsty/deb) | DEB 构建源代码 |\n")
	b.WriteString("| [pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | 基础设施包仓库 |\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateAllLists generates all list pages (output to top-level ext/ directory)
func (g *CCListGenerator) GenerateAllLists() error {
	extDir := filepath.Join(g.OutputDir, "e")
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extension directory: %w", err)
	}

	osDir := filepath.Join(g.OutputDir, "os")
	if err := os.MkdirAll(osDir, 0755); err != nil {
		return fmt.Errorf("failed to create os directory: %w", err)
	}

	// Overview page (ext/_index.md)
	if err := g.GenerateOverviewPage(filepath.Join(g.OutputDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate overview page: %w", err)
	}

	// List pages go directly to ext/ top level
	if err := g.GenerateExtensionList(filepath.Join(g.OutputDir, "list.md")); err != nil {
		return fmt.Errorf("failed to generate extension list: %w", err)
	}

	if err := g.GenerateRPMList(filepath.Join(g.OutputDir, "rpm.md")); err != nil {
		return fmt.Errorf("failed to generate RPM list: %w", err)
	}

	if err := g.GenerateDEBList(filepath.Join(g.OutputDir, "deb.md")); err != nil {
		return fmt.Errorf("failed to generate DEB list: %w", err)
	}

	// Category list is now handled by cate/_index.md from the category generator

	if err := g.GenerateLanguageList(filepath.Join(g.OutputDir, "language.md")); err != nil {
		return fmt.Errorf("failed to generate language list: %w", err)
	}

	if err := g.GenerateLicenseList(filepath.Join(g.OutputDir, "license.md")); err != nil {
		return fmt.Errorf("failed to generate license list: %w", err)
	}

	if err := g.GenerateRepoList(filepath.Join(g.OutputDir, "repo.md")); err != nil {
		return fmt.Errorf("failed to generate repo list: %w", err)
	}

	// Divider: before dimension pages (归属仓库, 编程语言, 开源协议)
	if err := g.GenerateDivider(filepath.Join(g.OutputDir, "_div_dimension.md"), 200); err != nil {
		return fmt.Errorf("failed to generate dimension divider: %w", err)
	}

	// Divider: before category pages
	if err := g.GenerateDivider(filepath.Join(g.OutputDir, "_div_category.md"), 499); err != nil {
		return fmt.Errorf("failed to generate category divider: %w", err)
	}

	// Divider: before extension detail pages
	if err := g.GenerateDivider(filepath.Join(g.OutputDir, "_div_details.md"), 599); err != nil {
		return fmt.Errorf("failed to generate details divider: %w", err)
	}

	// Extension index page
	if err := g.GenerateExtensionIndexPage(filepath.Join(extDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate extension index: %w", err)
	}

	// OS index page
	if err := g.GenerateOSIndexPage(filepath.Join(osDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate OS index: %w", err)
	}

	return nil
}

// GenerateDivider generates a sidebar divider page
func (g *CCListGenerator) GenerateDivider(outputPath string, weight int) error {
	content := fmt.Sprintf(`---
title: ""
weight: %d
sidebar_divider: true
toc_hide: false
---
`, weight)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// ccRepoBadgeOrMiss returns a repo badge or a miss badge
func ccRepoBadgeOrMiss(repo sql.NullString) string {
	if repo.Valid && repo.String != "" {
		return CCRepoBadge(repo.String)
	}
	return `<span class="ext-badge ext-badge--miss">✗</span>`
}
