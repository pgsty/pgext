/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC List Generator - generates list pages for pigsty.cc (Chinese only)
*/
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
func (g *CCListGenerator) GenerateExtensionIndexPage(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "PostgreSQL 扩展"
linkTitle: "扩展"
description: "Pigsty 收录的所有 PostgreSQL 扩展"
weight: 10
---

`)

	b.WriteString("Pigsty 收录了 **")
	b.WriteString(fmt.Sprintf("%d", len(g.Cache.Extensions)))
	b.WriteString("** 个 PostgreSQL 扩展，涵盖了几乎所有常用的扩展。\n\n")

	// Quick stats
	leadCount := 0
	for _, ext := range g.Cache.Extensions {
		if ext.Lead {
			leadCount++
		}
	}
	b.WriteString(fmt.Sprintf("其中包含 **%d** 个独立扩展包。\n\n", leadCount))

	// Links to list pages
	b.WriteString("## 扩展列表\n\n")
	b.WriteString("- [**按扩展名浏览**](/ext/list/ext/)：按扩展名字母顺序排列的完整扩展列表\n")
	b.WriteString("- [**按包名浏览**](/ext/list/pkg/)：按安装包名排列的扩展列表\n")
	b.WriteString("- [**按分类浏览**](/ext/list/cate/)：按功能分类组织的扩展列表\n")
	b.WriteString("- [**按语言浏览**](/ext/list/lang/)：按编程语言分类的扩展列表\n")
	b.WriteString("- [**按许可证浏览**](/ext/list/license/)：按开源许可证分类的扩展列表\n\n")

	// Category quick links
	b.WriteString("## 分类概览\n\n")
	b.WriteString("| **分类** | **描述** | **扩展数** |\n")
	b.WriteString("|:--------:|:---------|:----------:|\n")

	for _, cat := range g.Cache.Categories {
		count := len(g.Cache.CateExtMap[strings.ToUpper(cat.Name)])
		desc := cat.Name
		if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
			desc = cat.ZhDesc.String
		}
		b.WriteString(fmt.Sprintf("| %s | %s | %d |\n",
			CCCategoryLink(cat.Name), desc, count))
	}

	b.WriteString("\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateExtensionList generates the full extension list page
func (g *CCListGenerator) GenerateExtensionList(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "扩展列表"
linkTitle: "按扩展名"
description: "按扩展名排序的完整扩展列表"
weight: 10
---

`)

	b.WriteString("## 扩展列表\n\n")
	b.WriteString("按扩展名字母顺序排列的完整扩展列表。\n\n")

	// Group by first letter
	letterMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.State.Valid && ext.State.String == "not-ready" {
			continue
		}
		firstLetter := strings.ToUpper(string(ext.Name[0]))
		if firstLetter < "A" || firstLetter > "Z" {
			firstLetter = "#"
		}
		letterMap[firstLetter] = append(letterMap[firstLetter], ext)
	}

	// Sort letters
	var letters []string
	for letter := range letterMap {
		letters = append(letters, letter)
	}
	sort.Strings(letters)

	// Generate table
	b.WriteString("| **扩展名** | **包名** | **分类** | **描述** |\n")
	b.WriteString("|:-----------|:---------|:--------:|:---------|\n")

	for _, letter := range letters {
		exts := letterMap[letter]
		// Sort by name
		sort.Slice(exts, func(i, j int) bool {
			return exts[i].Name < exts[j].Name
		})

		for _, ext := range exts {
			desc := SanitizeText(ext.Name)
			if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
				desc = SanitizeText(ext.EnDesc.String)
			}

			catLink := "-"
			if ext.Category.Valid {
				catLink = CCCategoryLink(ext.Category.String)
			}

			b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %s |\n",
				CCExtLink(ext.Name, ""),
				ext.Pkg,
				catLink,
				desc))
		}
	}

	b.WriteString("\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GeneratePackageList generates the package list page
func (g *CCListGenerator) GeneratePackageList(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "包列表"
linkTitle: "按包名"
description: "按安装包名排序的扩展列表"
weight: 20
---

`)

	b.WriteString("## 包列表\n\n")
	b.WriteString("按安装包名排序的扩展列表，每个包可能包含一个或多个扩展。\n\n")

	// Get unique packages (lead extensions only)
	var packages []*Extension
	for _, ext := range g.Cache.Extensions {
		if ext.Lead && (!ext.State.Valid || ext.State.String != "not-ready") {
			packages = append(packages, ext)
		}
	}

	// Sort by package name
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Pkg < packages[j].Pkg
	})

	// Generate table
	b.WriteString("| **包名** | **主扩展** | **版本** | **分类** | **RPM** | **DEB** |\n")
	b.WriteString("|:---------|:-----------|:--------:|:--------:|:-------:|:-------:|\n")

	for _, ext := range packages {
		version := "-"
		if ext.Version.Valid {
			version = ext.Version.String
		}

		catLink := "-"
		if ext.Category.Valid {
			catLink = CCCategoryLink(ext.Category.String)
		}

		rpmRepo := "✗"
		if ext.RpmRepo.Valid && ext.RpmRepo.String != "" {
			rpmRepo = strings.ToUpper(ext.RpmRepo.String)
		}

		debRepo := "✗"
		if ext.DebRepo.Valid && ext.DebRepo.String != "" {
			debRepo = strings.ToUpper(ext.DebRepo.String)
		}

		b.WriteString(fmt.Sprintf("| `%s` | %s | `%s` | %s | %s | %s |\n",
			ext.Pkg,
			CCExtLink(ext.Name, ""),
			version,
			catLink,
			rpmRepo,
			debRepo))
	}

	b.WriteString("\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateCategoryList generates the category list page
func (g *CCListGenerator) GenerateCategoryList(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "分类列表"
linkTitle: "按分类"
description: "按功能分类组织的扩展列表"
weight: 30
---

`)

	b.WriteString("## 分类列表\n\n")
	b.WriteString("按功能分类组织的 PostgreSQL 扩展列表。\n\n")

	// Generate section for each category
	for _, cat := range g.Cache.Categories {
		catKey := strings.ToUpper(cat.Name)
		exts := g.Cache.CateExtMap[catKey]

		if len(exts) == 0 {
			continue
		}

		// Sort by extension ID
		sort.Slice(exts, func(i, j int) bool {
			return exts[i].ID < exts[j].ID
		})

		desc := cat.Name
		if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
			desc = cat.ZhDesc.String
		}

		b.WriteString(fmt.Sprintf("### %s {#%s}\n\n", catKey, strings.ToLower(cat.Name)))
		b.WriteString(fmt.Sprintf("%s（%d 个扩展）\n\n", desc, len(exts)))

		b.WriteString("| **扩展名** | **包名** | **版本** | **描述** |\n")
		b.WriteString("|:-----------|:---------|:--------:|:---------|\n")

		for _, ext := range exts {
			if ext.State.Valid && ext.State.String == "not-ready" {
				continue
			}

			version := "-"
			if ext.Version.Valid {
				version = ext.Version.String
			}

			desc := SanitizeText(ext.Name)
			if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
				desc = SanitizeText(ext.EnDesc.String)
			}

			b.WriteString(fmt.Sprintf("| %s | `%s` | `%s` | %s |\n",
				CCExtLink(ext.Name, ""),
				ext.Pkg,
				version,
				desc))
		}

		b.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateLanguageList generates the language list page
func (g *CCListGenerator) GenerateLanguageList(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "语言列表"
linkTitle: "按语言"
description: "按编程语言分类的扩展列表"
weight: 40
---

`)

	b.WriteString("## 语言列表\n\n")
	b.WriteString("按编程语言分类的 PostgreSQL 扩展列表。\n\n")

	// Group by language
	langMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.State.Valid && ext.State.String == "not-ready" {
			continue
		}
		lang := "Unknown"
		if ext.Lang.Valid && ext.Lang.String != "" {
			lang = ext.Lang.String
		}
		langMap[lang] = append(langMap[lang], ext)
	}

	// Sort languages by count
	var langs []string
	for lang := range langMap {
		langs = append(langs, lang)
	}
	sort.Slice(langs, func(i, j int) bool {
		return len(langMap[langs[i]]) > len(langMap[langs[j]])
	})

	// Overview table
	b.WriteString("| **语言** | **扩展数** |\n")
	b.WriteString("|:---------|:----------:|\n")
	for _, lang := range langs {
		anchor := strings.ToLower(lang)
		anchor = strings.ReplaceAll(anchor, "+", "")
		b.WriteString(fmt.Sprintf("| [**%s**](#%s) | %d |\n", lang, anchor, len(langMap[lang])))
	}
	b.WriteString("\n")

	// Generate section for each language
	for _, lang := range langs {
		exts := langMap[lang]
		if len(exts) == 0 {
			continue
		}

		// Sort by extension name
		sort.Slice(exts, func(i, j int) bool {
			return exts[i].Name < exts[j].Name
		})

		anchor := strings.ToLower(lang)
		anchor = strings.ReplaceAll(anchor, "+", "")

		b.WriteString(fmt.Sprintf("### %s {#%s}\n\n", lang, anchor))
		b.WriteString(fmt.Sprintf("%s 语言编写的扩展（%d 个）\n\n", lang, len(exts)))

		b.WriteString("| **扩展名** | **包名** | **分类** | **描述** |\n")
		b.WriteString("|:-----------|:---------|:--------:|:---------|\n")

		for _, ext := range exts {
			catLink := "-"
			if ext.Category.Valid {
				catLink = CCCategoryLink(ext.Category.String)
			}

			desc := SanitizeText(ext.Name)
			if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
				desc = SanitizeText(ext.EnDesc.String)
			}

			b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %s |\n",
				CCExtLink(ext.Name, ""),
				ext.Pkg,
				catLink,
				desc))
		}

		b.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateLicenseList generates the license list page
func (g *CCListGenerator) GenerateLicenseList(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "许可证列表"
linkTitle: "按许可证"
description: "按开源许可证分类的扩展列表"
weight: 50
---

`)

	b.WriteString("## 许可证列表\n\n")
	b.WriteString("按开源许可证分类的 PostgreSQL 扩展列表。\n\n")

	// Group by license
	licenseMap := make(map[string][]*Extension)
	for _, ext := range g.Cache.Extensions {
		if ext.State.Valid && ext.State.String == "not-ready" {
			continue
		}
		license := "Unknown"
		if ext.License.Valid && ext.License.String != "" {
			license = ext.License.String
		}
		licenseMap[license] = append(licenseMap[license], ext)
	}

	// Sort licenses by count
	var licenses []string
	for lic := range licenseMap {
		licenses = append(licenses, lic)
	}
	sort.Slice(licenses, func(i, j int) bool {
		return len(licenseMap[licenses[i]]) > len(licenseMap[licenses[j]])
	})

	// Overview table
	b.WriteString("| **许可证** | **扩展数** |\n")
	b.WriteString("|:-----------|:----------:|\n")
	for _, lic := range licenses {
		anchor := strings.ToLower(lic)
		anchor = strings.ReplaceAll(anchor, " ", "-")
		b.WriteString(fmt.Sprintf("| [**%s**](#%s) | %d |\n", lic, anchor, len(licenseMap[lic])))
	}
	b.WriteString("\n")

	// Generate section for each license
	for _, lic := range licenses {
		exts := licenseMap[lic]
		if len(exts) == 0 {
			continue
		}

		// Sort by extension name
		sort.Slice(exts, func(i, j int) bool {
			return exts[i].Name < exts[j].Name
		})

		anchor := strings.ToLower(lic)
		anchor = strings.ReplaceAll(anchor, " ", "-")

		b.WriteString(fmt.Sprintf("### %s {#%s}\n\n", lic, anchor))
		b.WriteString(fmt.Sprintf("使用 %s 许可证的扩展（%d 个）\n\n", lic, len(exts)))

		b.WriteString("| **扩展名** | **包名** | **分类** | **描述** |\n")
		b.WriteString("|:-----------|:---------|:--------:|:---------|\n")

		for _, ext := range exts {
			catLink := "-"
			if ext.Category.Valid {
				catLink = CCCategoryLink(ext.Category.String)
			}

			desc := SanitizeText(ext.Name)
			if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
				desc = SanitizeText(ext.ZhDesc.String)
			} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
				desc = SanitizeText(ext.EnDesc.String)
			}

			b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %s |\n",
				CCExtLink(ext.Name, ""),
				ext.Pkg,
				catLink,
				desc))
		}

		b.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateCatalogPage generates the main list catalog page
func (g *CCListGenerator) GenerateCatalogPage(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "扩展目录"
linkTitle: "目录"
description: "PostgreSQL 扩展分类目录"
weight: 5
---

`)

	b.WriteString("## 扩展目录\n\n")
	b.WriteString("Pigsty 收录的 PostgreSQL 扩展分类目录。\n\n")

	// Stats
	leadCount := 0
	for _, ext := range g.Cache.Extensions {
		if ext.Lead {
			leadCount++
		}
	}

	b.WriteString(fmt.Sprintf("- **扩展总数**：%d\n", len(g.Cache.Extensions)))
	b.WriteString(fmt.Sprintf("- **扩展包数**：%d\n", leadCount))
	b.WriteString(fmt.Sprintf("- **分类数量**：%d\n", len(g.Cache.Categories)))
	b.WriteString(fmt.Sprintf("- **活跃 PG 版本**：%d（", len(g.Cache.PGVersions)))
	for i, pg := range g.Cache.PGVersions {
		if i > 0 {
			b.WriteString("、")
		}
		b.WriteString(fmt.Sprintf("%d", pg))
	}
	b.WriteString("）\n")
	b.WriteString(fmt.Sprintf("- **支持操作系统**：%d\n\n", len(g.Cache.OSVersions)))

	// Navigation
	b.WriteString("## 浏览方式\n\n")
	b.WriteString("- [**按扩展名浏览**](/ext/list/ext/)：按扩展名字母顺序排列\n")
	b.WriteString("- [**按包名浏览**](/ext/list/pkg/)：按安装包名排列\n")
	b.WriteString("- [**按分类浏览**](/ext/list/cate/)：按功能分类组织\n")
	b.WriteString("- [**按语言浏览**](/ext/list/lang/)：按编程语言分类\n")
	b.WriteString("- [**按许可证浏览**](/ext/list/license/)：按开源许可证分类\n\n")

	// Category summary
	b.WriteString("## 分类汇总\n\n")
	b.WriteString("| **分类** | **描述** | **扩展数** |\n")
	b.WriteString("|:--------:|:---------|:----------:|\n")

	for _, cat := range g.Cache.Categories {
		count := len(g.Cache.CateExtMap[strings.ToUpper(cat.Name)])
		desc := cat.Name
		if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
			desc = cat.ZhDesc.String
		}
		b.WriteString(fmt.Sprintf("| %s | %s | %d |\n",
			CCCategoryLink(cat.Name), desc, count))
	}

	b.WriteString("\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateOSIndexPage generates the OS index page
func (g *CCListGenerator) GenerateOSIndexPage(outputPath string) error {
	var b strings.Builder

	// Frontmatter
	b.WriteString(`---
title: "操作系统"
linkTitle: "系统"
description: "各操作系统的扩展可用性"
weight: 60
---

`)

	b.WriteString("## 操作系统支持\n\n")
	b.WriteString("Pigsty 支持以下操作系统平台的 PostgreSQL 扩展：\n\n")

	// Group by OS family
	rhelOS := make([]OSVersion, 0)
	debOS := make([]OSVersion, 0)
	ubuntuOS := make([]OSVersion, 0)

	for _, os := range g.Cache.OSVersions {
		if !os.Active {
			continue
		}
		switch {
		case strings.HasPrefix(os.OS, "el"):
			rhelOS = append(rhelOS, os)
		case strings.HasPrefix(os.OS, "d"):
			debOS = append(debOS, os)
		case strings.HasPrefix(os.OS, "u"):
			ubuntuOS = append(ubuntuOS, os)
		}
	}

	// RHEL family
	if len(rhelOS) > 0 {
		b.WriteString("### RHEL / Rocky / AlmaLinux\n\n")
		b.WriteString("| **系统** | **描述** | **架构** |\n")
		b.WriteString("|:---------|:---------|:--------:|\n")
		for _, os := range rhelOS {
			b.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				CCOSLink(os.OS), os.OSFull, os.OSArch))
		}
		b.WriteString("\n")
	}

	// Debian family
	if len(debOS) > 0 {
		b.WriteString("### Debian\n\n")
		b.WriteString("| **系统** | **描述** | **架构** |\n")
		b.WriteString("|:---------|:---------|:--------:|\n")
		for _, os := range debOS {
			b.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				CCOSLink(os.OS), os.OSFull, os.OSArch))
		}
		b.WriteString("\n")
	}

	// Ubuntu family
	if len(ubuntuOS) > 0 {
		b.WriteString("### Ubuntu\n\n")
		b.WriteString("| **系统** | **描述** | **架构** |\n")
		b.WriteString("|:---------|:---------|:--------:|\n")
		for _, os := range ubuntuOS {
			b.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				CCOSLink(os.OS), os.OSFull, os.OSArch))
		}
		b.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

// GenerateAllLists generates all list pages
func (g *CCListGenerator) GenerateAllLists() error {
	// Create list directory
	listDir := filepath.Join(g.OutputDir, "list")
	if err := os.MkdirAll(listDir, 0755); err != nil {
		return fmt.Errorf("failed to create list directory: %w", err)
	}

	// Create e directory
	extDir := filepath.Join(g.OutputDir, "e")
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extension directory: %w", err)
	}

	// Create os directory
	osDir := filepath.Join(g.OutputDir, "os")
	if err := os.MkdirAll(osDir, 0755); err != nil {
		return fmt.Errorf("failed to create os directory: %w", err)
	}

	// Generate catalog page
	if err := g.GenerateCatalogPage(filepath.Join(listDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate catalog page: %w", err)
	}

	// Generate extension list
	if err := g.GenerateExtensionList(filepath.Join(listDir, "ext.md")); err != nil {
		return fmt.Errorf("failed to generate extension list: %w", err)
	}

	// Generate package list
	if err := g.GeneratePackageList(filepath.Join(listDir, "pkg.md")); err != nil {
		return fmt.Errorf("failed to generate package list: %w", err)
	}

	// Generate category list
	if err := g.GenerateCategoryList(filepath.Join(listDir, "cate.md")); err != nil {
		return fmt.Errorf("failed to generate category list: %w", err)
	}

	// Generate language list
	if err := g.GenerateLanguageList(filepath.Join(listDir, "lang.md")); err != nil {
		return fmt.Errorf("failed to generate language list: %w", err)
	}

	// Generate license list
	if err := g.GenerateLicenseList(filepath.Join(listDir, "license.md")); err != nil {
		return fmt.Errorf("failed to generate license list: %w", err)
	}

	// Generate extension index page
	if err := g.GenerateExtensionIndexPage(filepath.Join(extDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate extension index: %w", err)
	}

	// Generate OS index page
	if err := g.GenerateOSIndexPage(filepath.Join(osDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate OS index: %w", err)
	}

	return nil
}
