/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"fmt"
	"os"
	"strings"
)

// GenerateIndexPage generates extension index pages (_index.md and _index.zh.md)
func (g *ExtensionGenerator) GenerateIndexPage(locale, outputPath string) error {
	// Filter out extensions that are not ready
	var extensions []*Extension
	for _, ext := range g.Cache.Extensions {
		if !ext.State.Valid || ext.State.String != "not-ready" {
			extensions = append(extensions, ext)
		}
	}

	// Get lead packages only
	var packages []*Extension
	for _, ext := range extensions {
		if ext.Lead {
			packages = append(packages, ext)
		}
	}

	content := g.generateIndexContent(extensions, packages, locale)
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) generateIndexContent(extensions, packages []*Extension, locale string) string {
	extCount := len(extensions)
	pkgCount := len(packages)
	isZh := locale == "zh"

	var b strings.Builder

	// Frontmatter
	if isZh {
		b.WriteString(fmt.Sprintf(`---
title: "扩展"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

## 包

共有 %d 个可用的 PostgreSQL 扩展软件包：

`, pkgCount))
	} else {
		b.WriteString(fmt.Sprintf(`---
title: "Extensions"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

## Packages

There are %d available PostgreSQL packages:

`, pkgCount))
	}

	// Generate packages table
	b.WriteString(g.generatePackagesIndexTable(packages, isZh))

	// Extensions section
	if isZh {
		b.WriteString(fmt.Sprintf("\n## 扩展\n\n共有 %d 个可用的 PostgreSQL 扩展：\n\n", extCount))
	} else {
		b.WriteString(fmt.Sprintf("\n## Extensions\n\nThere are %d available PostgreSQL extensions:\n\n", extCount))
	}

	// Generate extensions table
	b.WriteString(g.generateExtensionsIndexTable(extensions, isZh))

	// Attribute legend
	if isZh {
		b.WriteString("\n**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位")
	} else {
		b.WriteString("\n**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable")
	}

	return b.String()
}

func (g *ExtensionGenerator) generatePackagesIndexTable(packages []*Extension, isZh bool) string {
	var b strings.Builder

	// Table header
	if isZh {
		b.WriteString("| 包 | 版本 | 仓库 | 分类 | RPM | DEB |\n")
		b.WriteString("|:---|:-----|:-----|:-----|:-----|:-----|\n")
	} else {
		b.WriteString("| Package | Version | Repo | Category | RPM | DEB |\n")
		b.WriteString("|:--------|:--------|:-----|:---------|:-----|:-----|\n")
	}

	// Table rows
	for _, pkg := range packages {
		pkgLink := ExtShortcode(pkg.Name, pkg.Pkg)

		version := "-"
		if pkg.Version.Valid {
			version = fmt.Sprintf("`%s`", pkg.Version.String)
		}

		// URL badge
		urlBadge := Badge("N/A", "gray", "", "", "")
		if pkg.URL.Valid && pkg.URL.String != "" {
			linkText := "Link"
			urlBadge = Badge(linkText, "", "", pkg.URL.String, "")
		}

		// Category badge
		cateBadge := "-"
		if pkg.Category.Valid {
			cateBadge = CategoryShortcode(pkg.Category.String)
		}

		// RPM and DEB packages
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

func (g *ExtensionGenerator) generateExtensionsIndexTable(extensions []*Extension, isZh bool) string {
	var b strings.Builder

	// Table header
	if isZh {
		b.WriteString("| 扩展 | PG 版本列表          | 属性 | 描述 |\n")
		b.WriteString("|:-----|:-------------------|:----:|:-----|\n")
	} else {
		b.WriteString("| Extension | PG Versions | Attribute | Description |\n")
		b.WriteString("|:----------|:------------|:---------:|:--------------|\n")
	}

	// Table rows
	for _, ext := range extensions {
		extLink := ExtShortcode(ext.Name, ext.Pkg)

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

		// Truncate description to 100 chars
		if len(desc) > 100 {
			desc = desc[:100]
		}

		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
			extLink, pgBadges, attrBadge, desc))
	}

	return b.String()
}
