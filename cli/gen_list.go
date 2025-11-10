/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ListGenerator generates list pages for categories, languages, and licenses
type ListGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewListGenerator creates a new list generator
func NewListGenerator(cache *ExtensionCache, outputDir string) *ListGenerator {
	return &ListGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateCategoryList generates category-based extension list pages
func (g *ListGenerator) GenerateCategoryList(locale, outputPath string) error {
	isZh := locale == "zh"

	// Count unique packages
	pkgMap := make(map[string]bool)
	for _, ext := range g.Cache.Extensions {
		pkgMap[ext.Pkg] = true
	}
	pkgCount := len(pkgMap)

	var b strings.Builder

	// Frontmatter
	if isZh {
		b.WriteString(fmt.Sprintf(`---
title: "按分类"
weight: 100
---

PostgreSQL 扩展（%d ext / %d pkg）归属 %d 个分类。

`, len(g.Cache.Extensions), pkgCount, len(g.Cache.Categories)))
	} else {
		b.WriteString(fmt.Sprintf(`---
title: "By Category"
weight: 100
---

PostgreSQL Extensions (%d ext in %d pkg) categorized into %d categories.

`, len(g.Cache.Extensions), pkgCount, len(g.Cache.Categories)))
	}

	b.WriteString(`

| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------| 
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |


`)

	// Generate sections for each category
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

	// Add category description
	desc := ""
	if isZh && category.ZhDesc.Valid {
		desc = category.ZhDesc.String
	} else if category.EnDesc.Valid {
		desc = category.EnDesc.String
	}

	if desc != "" {
		// Remove redundant category name prefix
		prefix := fmt.Sprintf("%s:", category.Name)
		if strings.HasPrefix(desc, prefix) {
			desc = strings.TrimSpace(desc[len(prefix):])
		}
		b.WriteString(SanitizeText(desc))
		b.WriteString("\n\n")
	}

	// Table header
	if isZh {
		b.WriteString("| ID | 扩展/包 | 版本 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension / Package | Version | Description |\n")
		b.WriteString("|:---:|:---|:---|:---|\n")
	}

	// Table rows
	if len(extensions) > 0 {
		// Sort by ID
		sort.Slice(extensions, func(i, j int) bool {
			return extensions[i].ID < extensions[j].ID
		})

		for _, ext := range extensions {
			desc := ""
			if isZh && ext.ZhDesc.Valid {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
			if desc == "" {
				desc = "-"
			}

			version := "-"
			if ext.Version.Valid {
				version = ext.Version.String
			}

			extPkgCell := AliasShortcode(ext.Name, ext.Pkg)
			b.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n",
				ext.ID, extPkgCell, version, desc))
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

// GenerateLanguageList generates language-based extension list pages
func (g *ListGenerator) GenerateLanguageList(locale, outputPath string) error {
	isZh := locale == "zh"

	// Group extensions by language
	langMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.Lang.Valid && ext.Lang.String != "" {
			lang := ext.Lang.String
			langMap[lang] = append(langMap[lang], ext)
		}
	}

	// Sort languages by extension count
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

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "按语言"
description: "按实现语言组织的 PostgreSQL 扩展"
excludeSearch: true
weight: 200
---

`)
	} else {
		b.WriteString(`---
title: "By Language"
description: "PostgreSQL extensions organized by implementation language"
excludeSearch: true
weight: 200
---

`)
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
		b.WriteString(fmt.Sprintf("| %s | %d | %s |\n",
			LanguageShortcode(lc.lang), lc.count, desc))
	}
	b.WriteString("\n\n")

	// Individual language sections
	for _, lc := range langs {
		b.WriteString(g.generateLanguageSection(lc.lang, langMap[lc.lang], isZh))
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func (g *ListGenerator) generateLanguageSection(lang string, extensions []*Extension, isZh bool) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n\n", lang))

	// Language badge with count
	countText := fmt.Sprintf("%d Extensions", len(extensions))
	if isZh {
		countText = fmt.Sprintf("%d 个扩展", len(extensions))
	}
	b.WriteString(fmt.Sprintf("%s %s\n\n", LanguageShortcode(lang), Badge(countText, "gray", "", "", "cube")))

	// Description
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

	// Table
	if isZh {
		b.WriteString("| ID | 扩展 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension | Description |\n")
		b.WriteString("|:---:|:---|:---|\n")
	}

	// Sort by ID
	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	for _, ext := range extensions {
		extCell := AliasShortcode(ext.Name, ext.Pkg)

		desc := ""
		if isZh {
			if ext.ZhDesc.Valid {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
			if desc == "" {
				desc = "暂无描述"
			}
		} else {
			if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
			if desc == "" {
				desc = "No description"
			}
		}

		b.WriteString(fmt.Sprintf("| %d | %s | %s |\n", ext.ID, extCell, desc))
	}
	b.WriteString("\n")

	return b.String()
}

// GenerateLicenseList generates license-based extension list pages
func (g *ListGenerator) GenerateLicenseList(locale, outputPath string) error {
	isZh := locale == "zh"

	// Group extensions by license
	licenseMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.License.Valid && ext.License.String != "" {
			license := ext.License.String
			licenseMap[license] = append(licenseMap[license], ext)
		}
	}

	// Sort licenses by count and display order

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

	// Frontmatter
	if isZh {
		b.WriteString(`---
title: "按许可证"
description: "按开源许可证组织的 PostgreSQL 扩展"
weight: 300
---

按照所使用开源许可证，对 PostgreSQL 扩展进行分类。

`)
	} else {
		b.WriteString(`---
title: "By License"
description: "PostgreSQL extensions organized by open source license"
weight: 300
---

PostgreSQL extension categorized by license.

`)
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

	// Individual license sections
	for _, lic := range licenses {
		b.WriteString(g.generateLicenseSection(lic.name, lic.exts, isZh))
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func (g *ListGenerator) generateLicenseSection(license string, extensions []*Extension, isZh bool) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("## %s\n\n", license))

	// License badge with count
	countText := fmt.Sprintf("%d Extensions", len(extensions))
	if isZh {
		countText = fmt.Sprintf("%d 个扩展", len(extensions))
	}
	info := getLicenseInfo(license)
	refText := "License Text"
	if isZh {
		refText = "许可证文本"
	}
	b.WriteString(fmt.Sprintf(`

| %s | %s  |
|:----|:---|
| %s | %s |

`, LicenseShortcode(license), Badge(countText, "gray", "", "", "cube"), Badge(refText, "gray", "", info.URL, "scale"), info.Description))

	// Table
	if isZh {
		b.WriteString("| ID | 扩展 | 描述 |\n")
		b.WriteString("|:---:|:---|:---|\n")
	} else {
		b.WriteString("| ID | Extension | Description |\n")
		b.WriteString("|:---:|:---|:---|\n")
	}

	// Sort by ID
	sort.Slice(extensions, func(i, j int) bool {
		return extensions[i].ID < extensions[j].ID
	})

	for _, ext := range extensions {
		desc := ""
		if isZh {
			if ext.ZhDesc.Valid {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
			if desc == "" {
				desc = "暂无描述"
			}
		} else {
			if ext.EnDesc.Valid {
				desc = SanitizeText(ext.EnDesc.String)
			}
		}

		b.WriteString(fmt.Sprintf("| %d | %s | %s |\n",
			ext.ID, AliasShortcode(ext.Name, ext.Pkg), desc))
	}

	b.WriteString("\n")
	return b.String()
}

type licenseItem struct {
	name  string
	exts  []*Extension
	order int
}

// Helper functions
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

type LicenseInfo struct {
	URL         string
	Description string
	Order       int
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
	info := getLicenseInfo(license)
	return info.Order
}
