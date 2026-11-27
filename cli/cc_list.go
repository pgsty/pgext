/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

OINK List Generator - generates list and overview pages for pigsty.cc (ZH) & pigsty.io (EN).
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

// CCListGenerator generates list pages for pigsty.cc (DocsZH) or pigsty.io (DocsEN).
type CCListGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	Locale    DocsLocale
}

func NewCCListGenerator(cache *ExtensionCache, outputDir string) *CCListGenerator {
	return &CCListGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsZH}
}

func NewIOListGenerator(cache *ExtensionCache, outputDir string) *CCListGenerator {
	return &CCListGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsEN}
}

func (g *CCListGenerator) locale() DocsLocale {
	if g.Locale == "" {
		return DocsZH
	}
	return g.Locale
}

// GenerateExtensionIndexPage generates e/_index.md only if it does not already exist.
func (g *CCListGenerator) GenerateExtensionIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	l := g.locale()
	content := fmt.Sprintf("---\ntitle: %q\ndescription: %q\nweight: 600\nicon: fas fa-puzzle-piece\n---\n",
		l.t("扩展详情", "Extensions"),
		l.t("每个 PostgreSQL 扩展的详细信息，可用性情况，下载链接，使用说明", "Detailed information for each PostgreSQL extension"),
	)
	return WriteMarkdownFile(outputPath, content)
}

// GenerateExtensionList generates the full extension list page (list.md).
func (g *CCListGenerator) GenerateExtensionList(outputPath string) error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\ndescription: %q\nweight: 120\nicon: fas fa-puzzle-piece\n---\n\n",
		l.t("扩展列表", "EXT List"),
		l.t("在 Pigsty 扩展目录中收录的 PostgreSQL 扩展列表", "PostgreSQL extensions available in the Pigsty Extension Catalog"),
	))

	allExts := g.Cache.ReadyExtensions()
	pgVersions := g.Cache.PGVersions

	var extAll, extPGDG, extPigsty, extContrib int
	extPGCount := make(map[int]int)
	var pkgAll, pkgPGDG, pkgPigsty int
	pkgPGCount := make(map[int]int)

	for _, ext := range allExts {
		repo := ""
		if ext.Repo.Valid {
			repo = ext.Repo.String
		}
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
			for _, pg := range ParsePGVersions(ext.PgVer) {
				extPGCount[pg]++
			}
		}

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
			for _, pg := range ParsePGVersions(ext.PgVer) {
				pkgPGCount[pg]++
			}
		}
	}

	b.WriteString(fmt.Sprintf(l.t(
		"在 Pigsty 扩展目录中共有 **%d** 个 PostgreSQL 扩展可用，共计 **%d** 个扩展包。\n\n",
		"There are **%d** PostgreSQL extensions available in the Pigsty Extension Catalog, in **%d** packages.\n\n",
	), extAll, pkgAll))

	b.WriteString(l.t("| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |", "| **Type** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |"))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--------:|:------:|:--------:|:----------:|:-----------:|:---:|")
	for range pgVersions {
		b.WriteString(":--------:|")
	}
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("| [**EXT**](/ext/list) | %d | %d | %d | %d | |", extAll, extPGDG, extPigsty, extContrib))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", extPGCount[pg]))
	}
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("| [**PKG**](/ext/repo) | %d | %d | %d | %d | |", pkgAll, pkgPGDG, pkgPigsty, 0))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", pkgPGCount[pg]))
	}
	b.WriteString("\n{.ext-table}\n\n")

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		var exts []*Extension
		for _, ext := range g.Cache.CateExtMap[catKey] {
			if !ext.Contrib && ext.IsReady() {
				exts = append(exts, ext)
			}
		}
		if len(exts) == 0 {
			continue
		}
		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })

		b.WriteString(fmt.Sprintf("\n--------\n\n## %s\n\n", catKey))
		b.WriteString(l.t(
			"| **扩展** | **包名** | **版本** | **属性** | **仓库** | **PG大版本** | **依赖** |\n",
			"| **Extension** | **Package** | **Version** | **Attr** | **Repo** | **PG Ver** | **Deps** |\n",
		))
		b.WriteString("|:---------|:---------|:--------:|:--------:|:--------:|:-----------:|:---------|\n")

		for _, ext := range exts {
			extLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)
			pkgName := ext.GetPkgURLLink()
			verStr := "-"
			if ext.Version.Valid && ext.Version.String != "" {
				verStr = fmt.Sprintf("`%s`", ext.Version.String)
			}
			attrStr := fmt.Sprintf(`<span style="white-space:nowrap"><code>%s</code></span>`, ext.GetAttributeBadge())
			repoBadge := "-"
			if ext.Repo.Valid && ext.Repo.String != "" {
				repoBadge = CCRepoBadge(ext.Repo.String)
			}
			depsStr := "-"
			if len(ext.Requires) > 0 {
				depList := make([]string, len(ext.Requires))
				for i, d := range ext.Requires {
					depList[i] = fmt.Sprintf("[`%s`](/ext/e/%s)", d, d)
				}
				depsStr = strings.Join(depList, ", ")
			}
			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				extLink, pkgName, verStr, attrStr, repoBadge, CCPGVersShortcode(ext.PgVer), depsStr))
		}
		b.WriteString("{.ext-table}\n\n")
	}
	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateLanguageList generates the language list page.
func (g *CCListGenerator) GenerateLanguageList(outputPath string) error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 220\nicon: fas fa-code\n---\n\n",
		l.t("编程语言", "Language"),
		l.t("编程语言", "Language"),
		l.t("按编程语言分类的扩展列表", "Extensions by programming language"),
	))

	langMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		lang := "Unknown"
		if ext.Lang.Valid && ext.Lang.String != "" {
			lang = ext.Lang.String
		}
		langMap[lang] = append(langMap[lang], ext)
	}

	langs := make([]string, 0, len(langMap))
	for lang := range langMap {
		langs = append(langs, lang)
	}
	sort.Slice(langs, func(i, j int) bool {
		if len(langMap[langs[i]]) != len(langMap[langs[j]]) {
			return len(langMap[langs[i]]) > len(langMap[langs[j]])
		}
		return langs[i] < langs[j]
	})

	b.WriteString(l.t("| **语言** |", "| **Language** |"))
	for _, lang := range langs {
		b.WriteString(fmt.Sprintf(" %s |", CCLanguageBadge(lang)))
	}
	b.WriteString("\n|:------:|")
	for range langs {
		b.WriteString(":------:|")
	}
	b.WriteString(l.t("\n| **扩展数** |", "\n| **Count** |"))
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
		b.WriteString(fmt.Sprintf("\n--------\n\n## %s {#%s}\n\n", lang, anchor))
		b.WriteString(fmt.Sprintf(l.t("%s 语言编写的扩展（%d 个）\n\n", "Extensions written in %s (%d)\n\n"), lang, len(exts)))

		isRust := lang == "Rust"
		if isRust {
			b.WriteString(l.t(
				"| **扩展** | **包** | **版本** | **PGRX** | **分类** | **许可证** | **描述** |\n",
				"| **Extension** | **Package** | **Version** | **PGRX** | **Type** | **License** | **Description** |\n",
			))
			b.WriteString("|:---------|:-------|:--------:|:--------:|:--------:|:----------:|:---------|\n")
		} else {
			b.WriteString(l.t(
				"| **扩展** | **包** | **版本** | **分类** | **许可证** | **描述** |\n",
				"| **Extension** | **Package** | **Version** | **Type** | **License** | **Description** |\n",
			))
			b.WriteString("|:---------|:-------|:--------:|:--------:|:----------:|:---------|\n")
		}

		for _, ext := range exts {
			catBadge, license := "-", "-"
			if ext.Category.Valid {
				catBadge = CCCategoryBadge(ext.Category.String)
			}
			if ext.License.Valid && ext.License.String != "" {
				license = CCLicenseBadge(ext.License.String)
			}
			desc := SanitizeText(l.extDesc(ext))
			pkgLink := ext.GetPkgURLLink()
			version := ext.GetVersion()
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
	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateLicenseList generates the license list page.
func (g *CCListGenerator) GenerateLicenseList(outputPath string) error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 230\nicon: fas fa-scale-balanced\n---\n\n",
		l.t("开源协议", "License"),
		l.t("开源协议", "License"),
		l.t("按开源许可证分类的扩展列表", "Extensions by open-source license"),
	))

	licenseMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		license := "Unknown"
		if ext.License.Valid && ext.License.String != "" {
			license = ext.License.String
		}
		licenseMap[license] = append(licenseMap[license], ext)
	}

	permissive := []string{"MIT", "ISC", "PostgreSQL", "0BSD", "BSD-2-Clause", "BSD-3-Clause", "Apache-2.0"}
	other := []string{"Artistic-1.0", "Artistic-2.0", "MPL-2.0", "GPL-2.0", "GPL-3.0", "LGPL-2.1", "LGPL-3.0", "AGPL-3.0", "Timescale"}
	known := make(map[string]bool, len(permissive)+len(other))
	for _, lic := range permissive {
		known[lic] = true
	}
	for _, lic := range other {
		known[lic] = true
	}
	var extra []string
	for lic := range licenseMap {
		if !known[lic] {
			extra = append(extra, lic)
		}
	}
	sort.Strings(extra)
	other = append(other, extra...)

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
	other = filterPresent(other)
	allLicenses := append(append([]string{}, permissive...), other...)

	maxCols := max(len(permissive), len(other))
	writeLicenseRow := func(label string, lics []string, asCount bool) {
		b.WriteString(fmt.Sprintf("| **%s** |", label))
		for _, lic := range lics {
			if asCount {
				b.WriteString(fmt.Sprintf(" %d |", len(licenseMap[lic])))
			} else {
				b.WriteString(fmt.Sprintf(" %s |", CCLicenseBadge(lic)))
			}
		}
		for i := len(lics); i < maxCols; i++ {
			b.WriteString(" |")
		}
	}

	licLabel := l.t("许可证", "License")
	writeLicenseRow(licLabel, permissive, false)
	b.WriteString("\n|:------:|")
	for i := 0; i < maxCols; i++ {
		b.WriteString(":------:|")
	}
	b.WriteString("\n")
	writeLicenseRow(l.t("宽松", "Permissive"), permissive, true)
	b.WriteString("\n")
	writeLicenseRow(licLabel, other, false)
	b.WriteString("\n")
	writeLicenseRow(l.t("其他许可证", "Other licenses"), other, true)
	b.WriteString("\n{.ext-table}\n\n")

	for _, lic := range allLicenses {
		exts := licenseMap[lic]
		if len(exts) == 0 {
			continue
		}
		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })
		anchor := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(lic, ".", ""), "-", ""), " ", ""))

		b.WriteString(fmt.Sprintf("\n--------\n\n## %s {#%s}\n\n", lic, anchor))
		b.WriteString(fmt.Sprintf(l.t("使用 %s 许可证的扩展（%d 个）\n\n", "Extensions using %s license (%d)\n\n"), lic, len(exts)))
		b.WriteString(l.t(
			"| **扩展** | **包** | **版本** | **分类** | **语言** | **描述** |\n",
			"| **Extension** | **Package** | **Version** | **Type** | **Language** | **Description** |\n",
		))
		b.WriteString("|:---------|:-------|:--------:|:--------:|:--------:|:---------|\n")

		for _, ext := range exts {
			catBadge, lang := "-", "-"
			if ext.Category.Valid {
				catBadge = CCCategoryBadge(ext.Category.String)
			}
			if ext.Lang.Valid && ext.Lang.String != "" {
				lang = CCLanguageBadge(ext.Lang.String)
			}
			b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
				ext.Name, ext.Name, ext.GetPkgURLLink(), ext.GetVersion(), catBadge, lang, SanitizeText(l.extDesc(ext))))
		}
		b.WriteString("{.ext-table}\n\n")
	}
	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateRepoList generates the repository list page.
func (g *CCListGenerator) GenerateRepoList(outputPath string) error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 210\nicon: fas fa-warehouse\n---\n\n",
		l.t("归属仓库", "Repository"),
		l.t("归属仓库", "Repository"),
		l.t("按归属仓库分类的扩展列表", "Extensions by repository"),
	))

	var pgdgExts, pigstyExts, mixedExts, contribExts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if !ext.Repo.Valid {
			continue
		}
		switch ext.Repo.String {
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
	for _, exts := range [][]*Extension{pgdgExts, pigstyExts, mixedExts, contribExts} {
		sort.Slice(exts, func(i, j int) bool { return exts[i].ID < exts[j].ID })
	}

	total := len(pgdgExts) + len(pigstyExts) + len(mixedExts) + len(contribExts)
	b.WriteString(fmt.Sprintf(l.t(
		"| **仓库** | %s | %s | %s | %s | **总计** |\n",
		"| **Repo** | %s | %s | %s | %s | **Total** |\n",
	), CCRepoBadge("PGDG"), CCRepoBadge("PIGSTY"), CCRepoBadge("MIXED"), CCRepoBadge("CONTRIB")))
	b.WriteString("|:------:|:------:|:------:|:------:|:------:|:------:|\n")
	b.WriteString(fmt.Sprintf(l.t(
		"| **扩展数量** | %d | %d | %d | %d | **%d** |\n{.ext-table}\n\n",
		"| **Extension Count** | %d | %d | %d | %d | **%d** |\n{.ext-table}\n\n",
	), len(pgdgExts), len(pigstyExts), len(mixedExts), len(contribExts), total))

	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf(l.t(
		"## PGDG 仓库 {#pgdg}\n\nPostgreSQL 全球开发组提供的二进制制成品仓库，包括 %d 个扩展。\n\n",
		"## PGDG Repository {#pgdg}\n\nBinary packages from the PostgreSQL Global Development Group, including %d extensions.\n\n",
	), len(pgdgExts)))
	writeRepoExtTable(&b, l, pgdgExts)

	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf(l.t(
		"## PIGSTY 仓库 {#pigsty}\n\nPigsty 扩展仓库在 PGDG 仓库的基础上，额外提供了 %d 个扩展。\n\n",
		"## PIGSTY Repository {#pigsty}\n\nPigsty extension repo provides %d additional extensions on top of PGDG.\n\n",
	), len(pigstyExts)))
	writeRepoExtTable(&b, l, pigstyExts)

	if len(mixedExts) > 0 {
		b.WriteString("\n--------\n\n")
		b.WriteString(fmt.Sprintf(l.t(
			"## 两者皆有 {#mixed}\n\n有一些扩展在 PIGSTY 和 PGDG 仓库中同时提供（%d 个）。\n\n",
			"## Both {#mixed}\n\nSome extensions are available in both PIGSTY and PGDG repositories (%d).\n\n",
		), len(mixedExts)))
		writeRepoExtTable(&b, l, mixedExts)
	}

	b.WriteString("\n--------\n\n")
	b.WriteString(fmt.Sprintf(l.t(
		"## CONTRIB {#contrib}\n\nPostgreSQL 内置的 Contrib 扩展模块，总计 %d 个扩展，PG 随内核 contrib 包交付。\n\n",
		"## CONTRIB {#contrib}\n\nPostgreSQL built-in contrib extension modules, %d extensions total, shipped with the PG kernel contrib package.\n\n",
	), len(contribExts)))
	writeContribExtTable(&b, l, contribExts)

	return WriteMarkdownFile(outputPath, b.String())
}

func writeRepoExtTable(b *strings.Builder, l DocsLocale, exts []*Extension) {
	b.WriteString(l.t(
		"| **扩展** | **RPM仓库** | **RPM包** | **RPM版本** | **DEB仓库** | **DEB包** | **DEB版本** |\n",
		"| **Extension** | **RPM Repo** | **RPM Pkg** | **RPM Ver** | **DEB Repo** | **DEB Pkg** | **DEB Ver** |\n",
	))
	b.WriteString("|:---------|:----------:|:---------|:----------:|:----------:|:---------|:----------:|\n")
	codeOrDash := func(s sql.NullString) string {
		if s.Valid && s.String != "" {
			return fmt.Sprintf("`%s`", s.String)
		}
		return "-"
	}
	repoOrDash := func(s sql.NullString) string {
		if s.Valid && s.String != "" {
			return CCRepoBadge(s.String)
		}
		return "-"
	}
	for _, ext := range exts {
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s | %s | %s | %s | %s |\n",
			ext.Name, ext.Name,
			repoOrDash(ext.RpmRepo), codeOrDash(ext.RpmPkg), codeOrDash(ext.RpmVer),
			repoOrDash(ext.DebRepo), codeOrDash(ext.DebPkg), codeOrDash(ext.DebVer)))
	}
	b.WriteString("{.ext-table}\n\n")
}

func writeContribExtTable(b *strings.Builder, l DocsLocale, exts []*Extension) {
	b.WriteString(l.t(
		"| **扩展** | **分类** | **版本** | **PG版本** | **说明** |\n",
		"| **Extension** | **Type** | **Version** | **PG Ver** | **Description** |\n",
	))
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
			vers := make([]string, len(ext.PgVer))
			for i, v := range ext.PgVer {
				vers[i] = strings.TrimSuffix(v, "+")
			}
			pgVer = strings.Join(vers, ", ")
		}
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s | %s | %s |\n",
			ext.Name, ext.Name, catBadge, version, pgVer, SanitizeText(l.extDesc(ext))))
	}
	b.WriteString("{.ext-table}\n\n")
}

type platformListConfig struct {
	Name         string
	OSDesc       string
	Title        string
	Desc         string
	Icon         string
	Weight       int
	OtherName    string
	GetRepo      func(*Extension) sql.NullString
	GetPkg       func(*Extension) sql.NullString
	GetVer       func(*Extension) sql.NullString
	GetPG        func(*Extension) []string
	GetDeps      func(*Extension) []string
	GetOtherRepo func(*Extension) sql.NullString
}

func (g *CCListGenerator) GenerateRPMList(outputPath string) error {
	l := g.locale()
	return g.generatePlatformList(outputPath, &platformListConfig{
		Name:         "RPM",
		OSDesc:       "EL",
		Title:        l.t("RPM 列表", "RPM List"),
		Desc:         l.t("在 EL 系统上可用的 PostgreSQL 扩展 RPM 二进制包", "PostgreSQL extension RPM binary packages available on EL systems"),
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

func (g *CCListGenerator) GenerateDEBList(outputPath string) error {
	l := g.locale()
	return g.generatePlatformList(outputPath, &platformListConfig{
		Name:         "DEB",
		OSDesc:       "Debian/Ubuntu",
		Title:        l.t("DEB 列表", "DEB List"),
		Desc:         l.t("在 Debian/Ubuntu 系统上可用的 PostgreSQL 扩展 DEB 二进制包", "PostgreSQL extension DEB binary packages available on Debian/Ubuntu systems"),
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

func (g *CCListGenerator) generatePlatformList(outputPath string, cfg *platformListConfig) error {
	l := g.locale()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: \"%s\"\ndescription: \"%s\"\nweight: %d\nicon: %s\n---\n\n",
		cfg.Title, cfg.Desc, cfg.Weight, cfg.Icon))

	allExts := g.Cache.ReadyExtensions()
	pgVersions := g.Cache.PGVersions

	var extAll, extPGDG, extPigsty, extContrib int
	extPGCount := make(map[int]int)
	var pkgAll, pkgPGDG, pkgPigsty int
	pkgPGCount := make(map[int]int)
	var exclusive, otherMissing int

	for _, ext := range allExts {
		repo := cfg.GetRepo(ext)
		hasRepo := repo.Valid && repo.String != ""
		otherRepo := cfg.GetOtherRepo(ext)
		hasOther := otherRepo.Valid && otherRepo.String != ""

		if hasRepo || ext.Contrib {
			extAll++
			if !ext.Contrib {
				switch repo.String {
				case "PGDG":
					extPGDG++
				case "PIGSTY":
					extPigsty++
				}
			} else {
				extContrib++
			}
			pgs := cfg.GetPG(ext)
			if ext.Contrib && len(pgs) == 0 {
				pgs = ext.PgVer
			}
			for _, pg := range ParsePGVersions(pgs) {
				extPGCount[pg]++
			}
		}

		if !ext.Lead {
			continue
		}
		if hasRepo {
			pkgAll++
			switch repo.String {
			case "PGDG":
				pkgPGDG++
			case "PIGSTY":
				pkgPigsty++
			}
			for _, pg := range ParsePGVersions(cfg.GetPG(ext)) {
				pkgPGCount[pg]++
			}
		}
		if !ext.Contrib {
			if hasRepo && !hasOther {
				exclusive++
			}
			if hasOther && !hasRepo {
				otherMissing++
			}
		}
	}

	b.WriteString(fmt.Sprintf(l.t(
		"在 **%d** 个已打包扩展总目录中，%s 系统当前覆盖 **%d** 个 PostgreSQL 扩展，共计 **%d** 个扩展包。\n\n",
		"Across the **%d** packaged extensions catalog, %s systems currently cover **%d** PostgreSQL extensions in **%d** packages.\n\n",
	), len(allExts), cfg.OSDesc, extAll, pkgAll))
	b.WriteString(fmt.Sprintf(l.t(
		"其中 **%d** 个扩展包是 %s 独有，**%d** 个 %s 扩展包缺少对应的 %s 包。\n\n",
		"**%d** packages are %s exclusive, **%d** %s packages are missing corresponding %s packages.\n\n",
	), exclusive, cfg.OSDesc, otherMissing, cfg.OtherName, cfg.Name))

	extLink := fmt.Sprintf("/ext/%s", strings.ToLower(cfg.Name))
	b.WriteString(l.t("| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |", "| **Type** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |"))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--------:|:------:|:--------:|:----------:|:-----------:|:---:|")
	for range pgVersions {
		b.WriteString(":--------:|")
	}
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("| [**EXT**](%s) | %d | %d | %d | %d | |", extLink, extAll, extPGDG, extPigsty, extContrib))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", extPGCount[pg]))
	}
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("| [**PKG**](/ext/repo) | %d | %d | %d | %d | |", pkgAll, pkgPGDG, pkgPigsty, 0))
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf(" %d |", pkgPGCount[pg]))
	}
	b.WriteString("\n{.ext-table}\n\n")

	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		var pkgs []*Extension
		for _, ext := range g.Cache.CateExtMap[catKey] {
			if !ext.Lead || ext.Contrib || !ext.IsReady() {
				continue
			}
			if repo := cfg.GetRepo(ext); repo.Valid && repo.String != "" {
				pkgs = append(pkgs, ext)
			}
		}
		if len(pkgs) == 0 {
			continue
		}
		sort.Slice(pkgs, func(i, j int) bool { return pkgs[i].ID < pkgs[j].ID })

		b.WriteString(fmt.Sprintf("\n--------\n\n## %s\n\n", catKey))
		b.WriteString(l.t(
			"| **扩展包** | **版本** | **仓库** | **包名** | **PG大版本** | **依赖** |\n",
			"| **Package** | **Version** | **Repo** | **Pattern** | **PG Ver** | **Deps** |\n",
		))
		b.WriteString("|:----------|:--------:|:--------:|:---------|:-----------:|:---------|\n")

		for _, ext := range pkgs {
			pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Pkg, ext.Name)
			verStr, repoBadge, pkgName, depsStr := "-", "-", "-", "-"
			if ver := cfg.GetVer(ext); ver.Valid && ver.String != "" {
				verStr = fmt.Sprintf("`%s`", ver.String)
			}
			if repo := cfg.GetRepo(ext); repo.Valid && repo.String != "" {
				repoBadge = CCRepoBadge(repo.String)
			}
			if pkg := cfg.GetPkg(ext); pkg.Valid && pkg.String != "" {
				pkgName = fmt.Sprintf("`%s`", pkg.String)
			}
			pgs := cfg.GetPG(ext)
			if ext.Contrib && len(pgs) == 0 {
				pgs = ext.PgVer
			}
			if deps := cfg.GetDeps(ext); len(deps) > 0 {
				depList := make([]string, len(deps))
				for i, d := range deps {
					depList[i] = fmt.Sprintf("`%s`", d)
				}
				depsStr = strings.Join(depList, ", ")
			}
			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
				pkgLink, verStr, repoBadge, pkgName, CCPGVersShortcode(pgs), depsStr))
		}
		b.WriteString("{.ext-table}\n\n")
	}
	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateOSIndexPage generates os/_index.md only if it does not already exist.
func (g *CCListGenerator) GenerateOSIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	l := g.locale()
	content := fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 400\nicon: fa-brands fa-linux\n---\n",
		l.t("操作系统", "Platform"),
		l.t("操作系统", "Platform"),
		l.t("各操作系统的扩展可用性", "Extension availability by operating system"),
	)
	return WriteMarkdownFile(outputPath, content)
}

// GenerateOverviewPage generates the ext/_index.md overview page with dynamic data.
func (g *CCListGenerator) GenerateOverviewPage(outputPath string) error {
	l := g.locale()
	var b strings.Builder

	type rowStats struct {
		All, PGDG, PIGSTY, CONTRIB int
		PGCounts                   map[int]int
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

	for _, ext := range g.Cache.ReadyExtensions() {
		repo := ""
		if ext.Repo.Valid {
			repo = ext.Repo.String
		}
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

		addSide := func(row *rowStats, sideRepo sql.NullString, sidePG []string) {
			hasSide := sideRepo.Valid && sideRepo.String != ""
			if !hasSide && !ext.Contrib {
				return
			}
			row.All++
			if hasSide {
				switch sideRepo.String {
				case "PGDG":
					row.PGDG++
				case "PIGSTY":
					row.PIGSTY++
				}
			}
			if ext.Contrib {
				row.CONTRIB++
			}
			for _, pg := range g.Cache.PGVersions {
				if (hasSide && pgContains(sidePG, pg)) || (ext.Contrib && pgContains(ext.PgVer, pg)) {
					row.PGCounts[pg]++
				}
			}
		}
		addSide(&elRow, ext.RpmRepo, ext.RpmPg)
		addSide(&debRow, ext.DebRepo, ext.DebPg)
	}

	totalExts := allRow.All

	type cateInfo struct{ Name, Pkgs string }
	var cateInfos []cateInfo
	for _, cat := range g.Cache.Categories {
		cateName := strings.ToUpper(cat.Name)
		var pkgLinks []string
		for _, ext := range g.Cache.CateExtMap[cateName] {
			if !ext.IsReady() || !ext.Lead || ext.Contrib {
				continue
			}
			pkgLinks = append(pkgLinks, fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Pkg, ext.Name))
		}
		cateInfos = append(cateInfos, cateInfo{Name: cateName, Pkgs: strings.Join(pkgLinks, "  ")})
	}

	activeOSCount := 0
	for _, osv := range g.Cache.OSVersions {
		if osv.Active {
			activeOSCount++
		}
	}

	writeStatsRow := func(label string, r rowStats) {
		b.WriteString(fmt.Sprintf("| **%s** | %d | %d | %d | %d | |", label, r.All, r.PGDG, r.PIGSTY, r.CONTRIB))
		for _, pg := range g.Cache.PGVersions {
			b.WriteString(fmt.Sprintf(" %d |", r.PGCounts[pg]))
		}
		b.WriteString("\n")
	}

	b.WriteString(fmt.Sprintf("---\ntitle: %s\nlinkTitle: %s\nicon: fas fa-puzzle-piece\ndescription: >\n  %s\nweight: 100\n%s\n---\n\n",
		l.t("PG 扩展目录", "PG Extension Catalog"),
		l.t("PG 扩展目录", "PG Extension Catalog"),
		fmt.Sprintf(l.t(
			"扩展是 PostgreSQL 的灵魂所在，Pigsty 制作、打包、收录并整合了 %d 个 PG 生态扩展，供用户开箱即用。",
			"Extensions are the soul of PostgreSQL. Pigsty builds, packages, collects, and integrates %d PG ecosystem extensions for out-of-the-box use.",
		), totalExts),
		extOverviewOutputsFrontMatter,
	))

	if l.isZH() {
		b.WriteString(fmt.Sprintf(`Pigsty 提供了以下三样基础设施，帮助用户更好的利用 PostgreSQL 扩展生态系统的协同超能力：

- [**已打包扩展目录**](/ext/list)：查阅 [**%d 个已打包扩展**](/ext/list) 的详细信息，使用方法，元数据，下载链接与文档
- [**扩展仓库**](/docs/repo/pgsql)：获取预先打包的 RPM/DEB 二进制包，在 [**%d 个 Linux 系统**](/ext/os) 上可用
- [**包管理器**](https://pig.pgsty.com/zh)：使用 [**`+"`"+`pig`+"`"+`**](https://pig.pgsty.com/zh) 命令行工具，屏蔽复杂度与操作系统与架构差异

`+"```bash"+`
curl -fsSL https://repo.pigsty.cc/pig | bash  # 安装 pig 命令行工具
pig repo add pigsty pgdg -u                   # 在您的 Linux 发行版上配置软件仓库
pig install pg18                              # 从 PGDG 官方仓库安装 PostgreSQL 18 内核包
pig install pg_duckdb -v 18                   # 例：针对 PG 18 安装 pg_duckdb
`+"```"+`

一切都皆可用 PostgreSQL 解决！请参阅我们的博客文章：[**PostgreSQL 正在吞噬数据库世界！**](/blog/pg/pg-eat-db-world)

![Pigsty PostgreSQL 扩展生态系统](/img/pigsty/ecosystem.png)

--------

## 核心特点

- **数量**：无与伦比的扩展数量：**%d 个已打包扩展**，为 PG 扩展生态之最
- **质量**：原生 Linux RPM/DEB 包，完全兼容 PGDG 打包规范
- **易用**：提供包管理器 [**`+"`"+`pig`+"`"+`**](https://pig.pgsty.com/zh)，屏蔽操作系统与架构差异，开箱即用
- **兼容**：扩展完全兼容 PGDG 打包规范，可与 PGDG 仓库无缝混用
- **分发**：由 Cloudflare CDN 进行全球仓库分发，提供国内 CDN 加速
- **开源**：完全开源，提供便利的构建工具，免费对公众提供服务的软件基础设施


`, totalExts, activeOSCount, totalExts))
	} else {
		b.WriteString(fmt.Sprintf(`Pigsty provides the following three pieces of infrastructure to help users harness the collaborative superpowers of the PostgreSQL extension ecosystem:

- [**Packaged Extension Catalog**](/ext/list): Browse detailed information, usage, metadata, download links and docs for [**%d packaged extensions**](/ext/list)
- [**Extension Repository**](/docs/repo/pgsql): Get pre-built RPM/DEB binary packages, available on [**%d Linux systems**](/ext/os)
- [**Package Manager**](https://pig.pgsty.com): Use the [`+"`"+`pig`+"`"+`](https://pig.pgsty.com) CLI tool to abstract away OS and architecture differences

`+"```bash"+`
curl -fsSL https://repo.pigsty.io/pig | bash   # Install pig CLI tool
pig repo add pigsty pgdg -u                    # Configure repos on your Linux distro
pig install pg18                               # Install PostgreSQL 18 kernel from PGDG
pig install pg_duckdb -v 18                    # e.g., install pg_duckdb for PG 18
`+"```"+`

Everything can be solved with PostgreSQL! Check out our blog post: [**PostgreSQL is eating the database world!**](/blog/pg/pg-eat-db-world)

![Pigsty PostgreSQL extension ecosystem](/img/pigsty/ecosystem.png)

--------

## Key Features

- **Quantity**: Unparalleled extension count: **%d packaged extensions**, the most in the PG extension ecosystem
- **Quality**: Native Linux RPM/DEB packages, fully compatible with PGDG packaging standards
- **Ease of Use**: Provides the [**`+"`"+`pig`+"`"+`**](https://pig.pgsty.com) package manager, abstracting away OS and architecture differences for out-of-the-box use
- **Compatibility**: Extensions are fully compatible with PGDG packaging standards, seamlessly usable with PGDG repositories
- **Distribution**: Global repository distribution via Cloudflare CDN
- **Open Source**: Fully open source, with convenient build tools and free public software infrastructure


`, totalExts, activeOSCount, totalExts))
	}

	b.WriteString(l.t("--------\n\n## 扩展统计\n\n| **分类** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |", "--------\n\n## Extension Statistics\n\n| **Type** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | |"))
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:------:|:------:|:------:|:------:|:------:|:---:|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n")
	writeStatsRow("ALL", allRow)
	writeStatsRow("EL", elRow)
	writeStatsRow("Debian", debRow)
	b.WriteString("{.ext-table}\n\n")
	b.WriteString(fmt.Sprintf(l.t(
		"> **%d** 是已打包扩展总目录口径；EL、Debian/Ubuntu 及各 PostgreSQL 大版本行展示对应平台的实际覆盖数，因此可能少于 %d。\n>\n> 详见：[扩展列表](/ext/list)，[RPM 列表](/ext/rpm)，[DEB 列表](/ext/deb)，[归属仓库](/ext/repo)\n\n",
		"> **%d** is the packaged catalog total; EL, Debian/Ubuntu, and per-major PostgreSQL rows show actual platform coverage and may be smaller than %d.\n>\n> See: [Extension List](/ext/list), [RPM List](/ext/rpm), [DEB List](/ext/deb), [Repository](/ext/repo)\n\n",
	), totalExts, totalExts))

	b.WriteString(l.t("--------\n\n## 扩展分类\n\n\n| **分类** | **扩展包** |\n", "--------\n\n## Extension Categories\n\n\n| **Type** | **Package** |\n"))
	b.WriteString("|:--|:--|\n")
	for _, ci := range cateInfos {
		b.WriteString(fmt.Sprintf("| %s | %s |\n", CCCategoryBadge(ci.Name), ci.Pkgs))
	}
	b.WriteString("{.ext-table}\n\n")

	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n{.full-width}\n\n",
		CCLicenseBadge("MIT"), CCLicenseBadge("ISC"), CCLicenseBadge("PostgreSQL"),
		CCLicenseBadge("0BSD"), CCLicenseBadge("BSD-2-Clause"), CCLicenseBadge("BSD-3-Clause"),
		CCLicenseBadge("Artistic-1.0")+" "+CCLicenseBadge("Artistic-2.0")))
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n|:------:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n{.full-width}\n\n",
		CCLicenseBadge("Apache-2.0"), CCLicenseBadge("MPL-2.0"),
		CCLicenseBadge("GPL-2.0"), CCLicenseBadge("GPL-3.0"),
		CCLicenseBadge("LGPL-2.1"), CCLicenseBadge("LGPL-3.0"),
		CCLicenseBadge("AGPL-3.0"), CCLicenseBadge("Timescale")))
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n|:------:|:------:|:------:|:------:|:------:|:------:|:------:|\n{.full-width}\n\n\n\n",
		CCLanguageBadge("C"), CCLanguageBadge("C++"), CCLanguageBadge("Rust"),
		CCLanguageBadge("Java"), CCLanguageBadge("Python"), CCLanguageBadge("SQL"),
		CCLanguageBadge("Data")))

	b.WriteString(l.t(
		"## 兼容系统\n\n| **系统** | **版本** | **x86_64** | **aarch64** | **PG版本** |\n",
		"## Compatible Systems\n\n| **System** | **Version** | **x86_64** | **aarch64** | **PG Ver** |\n",
	))
	b.WriteString("|:------:|:------:|:----------:|:-----------:|:--------:|\n")

	pgBadgeStr := CCPGVersAllShortcode(g.Cache.PGVersions)
	osMeta := make(map[string]OSVersion)
	for _, osv := range g.Cache.OSVersions {
		if osv.Active {
			code, _, _ := strings.Cut(osv.OS, ".")
			if _, ok := osMeta[code]; !ok {
				osMeta[code] = osv
			}
		}
	}
	for _, osc := range buildOSCodeEntries(g.Cache.OSVersions) {
		osv := osMeta[osc.Code]
		sysName := osv.OSFull
		for _, suffix := range []string{" x86_64", " aarch64", " x86", " ARM"} {
			sysName = strings.TrimSuffix(sysName, suffix)
		}
		version := osv.OSMajor
		if strings.HasPrefix(osc.Code, "u") && len(osv.OSMajor) == 2 {
			version = osv.OSMajor + ".04"
		}
		x86Link, armLink := "-", "-"
		if osc.X86 != "" {
			x86Link = fmt.Sprintf("[`%s`](/ext/os/%s)", osc.X86, osc.X86)
		}
		if osc.ARM != "" {
			armLink = fmt.Sprintf("[`%s`](/ext/os/%s)", osc.ARM, osc.ARM)
		}
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s | %s |\n", sysName, version, x86Link, armLink, pgBadgeStr))
	}
	b.WriteString("{.ext-table}\n\n")
	b.WriteString(l.t(
		"> 详见：[操作系统](/ext/os)，[RPM 列表](/ext/rpm)，[DEB 列表](/ext/deb)\n\n\n\n",
		"> See: [Operating System](/ext/os), [RPM List](/ext/rpm), [DEB List](/ext/deb)\n\n\n\n",
	))

	if l.isZH() {
		b.WriteString(`--------

## 合作伙伴

Pigsty 向用户提供无可比拟的 PostgreSQL 扩展交付体验，已有多家 PostgreSQL Vendor 使用它进行交付。

- [Pigsty](https://github.com/pgsty/pigsty) — 开箱即用的开源企业级 PostgreSQL RDS 发行版
- [Omnigres](https://docs.omnigres.org/quick_start/) — PostgreSQL as a Platform，在数据库中进行应用开发
- [AutoBase](https://autobase.tech/docs/extensions/install) — 基于 Ansible 的 PG 集群自动化部署，开源 DBaaS
- [TensorChord](https://github.com/tensorchord) — 云原生 AI 基础设施，开发了多个知名 PG 扩展
- [文武IT](https://w3.ww-it.cn/) — 企业级 PostgreSQL 服务商
{.cards}


--------

## 关于项目

本项目由 [**PGSTY**](https://github.com/pgsty) / [**VONNG**](https://vonng.com) 开发维护，并使用 [**Apache 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。

| **仓库** | **说明** |
|:------|:------|
| [pgsty](https://github.com/pgsty) | PGSTY 组织主页 |
| [pgsty/pgext](https://github.com/pgsty/pgext) | 扩展元数据、用法源稿与管理工具 |
| [pgsty/pigsty](https://github.com/pgsty/pigsty) | PostgreSQL 数据库发行版 |
| [pgsty/pig](https://github.com/pgsty/pig) | PostgreSQL 包管理器 |
| [pgsty/rpm](https://github.com/pgsty/rpm) | RPM 构建源代码 |
| [pgsty/deb](https://github.com/pgsty/deb) | DEB 构建源代码 |
| [pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | 基础设施包仓库 |
`)
	} else {
		b.WriteString(`--------

## Partners

Pigsty provides an unparalleled PostgreSQL extension delivery experience. Several PostgreSQL vendors are using it for delivery.

- [Pigsty](https://github.com/pgsty/pigsty) — Out-of-the-box, open-source, enterprise-grade PostgreSQL RDS distribution
- [Omnigres](https://docs.omnigres.org/quick_start/) — PostgreSQL as a Platform, application development inside the database
- [AutoBase](https://autobase.tech/docs/extensions/install) — Ansible-based PG cluster automation deployment, open-source DBaaS
- [TensorChord](https://github.com/tensorchord) — Cloud-native AI infrastructure, developers of several well-known PG extensions
- [WENWU IT](https://w3.ww-it.cn/) — Enterprise PostgreSQL service provider
{.cards}


--------

## About

This project is developed and maintained by [**PGSTY**](https://github.com/pgsty) / [**VONNG**](https://vonng.com), and is open-sourced under the [**Apache 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) license.

| **Repository** | **Description** |
|:------|:------|
| [pgsty](https://github.com/pgsty) | PGSTY organization homepage |
| [pgsty/pgext](https://github.com/pgsty/pgext) | Extension metadata, usage sources, and management tools |
| [pgsty/pigsty](https://github.com/pgsty/pigsty) | PostgreSQL database distribution |
| [pgsty/pig](https://github.com/pgsty/pig) | PostgreSQL package manager |
| [pgsty/rpm](https://github.com/pgsty/rpm) | RPM build source code |
| [pgsty/deb](https://github.com/pgsty/deb) | DEB build source code |
| [pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | Infrastructure package repository |
`)
	}
	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateAllLists generates all list pages under OutputDir.
func (g *CCListGenerator) GenerateAllLists() error {
	extDir := filepath.Join(g.OutputDir, "e")
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extension directory: %w", err)
	}
	osDir := filepath.Join(g.OutputDir, "os")
	if err := os.MkdirAll(osDir, 0755); err != nil {
		return fmt.Errorf("failed to create os directory: %w", err)
	}

	steps := []struct {
		name string
		fn   func() error
	}{
		{"overview page", func() error { return g.GenerateOverviewPage(filepath.Join(g.OutputDir, "_index.md")) }},
		{"extension list", func() error { return g.GenerateExtensionList(filepath.Join(g.OutputDir, "list.md")) }},
		{"RPM list", func() error { return g.GenerateRPMList(filepath.Join(g.OutputDir, "rpm.md")) }},
		{"DEB list", func() error { return g.GenerateDEBList(filepath.Join(g.OutputDir, "deb.md")) }},
		{"language list", func() error { return g.GenerateLanguageList(filepath.Join(g.OutputDir, "language.md")) }},
		{"license list", func() error { return g.GenerateLicenseList(filepath.Join(g.OutputDir, "license.md")) }},
		{"repo list", func() error { return g.GenerateRepoList(filepath.Join(g.OutputDir, "repo.md")) }},
		{"dimension divider", func() error { return g.GenerateDivider(filepath.Join(g.OutputDir, "_div_dimension.md"), 200) }},
		{"category divider", func() error { return g.GenerateDivider(filepath.Join(g.OutputDir, "_div_category.md"), 499) }},
		{"details divider", func() error { return g.GenerateDivider(filepath.Join(g.OutputDir, "_div_details.md"), 599) }},
		{"extension index", func() error { return g.GenerateExtensionIndexPage(filepath.Join(extDir, "_index.md")) }},
		{"OS index", func() error { return g.GenerateOSIndexPage(filepath.Join(osDir, "_index.md")) }},
	}
	for _, s := range steps {
		if err := s.fn(); err != nil {
			return fmt.Errorf("failed to generate %s: %w", s.name, err)
		}
	}
	return nil
}

// GenerateDivider generates a sidebar divider page.
func (g *CCListGenerator) GenerateDivider(outputPath string, weight int) error {
	content := fmt.Sprintf("---\ntitle: \"\"\nweight: %d\nsidebar_divider: true\ntoc_hide: false\n---\n", weight)
	return WriteMarkdownFile(outputPath, content)
}
