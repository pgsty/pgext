/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC Attr Generator - generates extension attribute pages for pigsty.cc (Chinese only)
Sub-pages: load (动态加载), ddl (无头扩展), deps (依赖关系), multi (多扩展包), fork (分支扩展)
*/
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// CCAttrGenerator generates extension attribute pages for pigsty.cc
type CCAttrGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewCCAttrGenerator creates a new CC Attr generator
func NewCCAttrGenerator(cache *ExtensionCache, outputDir string) *CCAttrGenerator {
	return &CCAttrGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateAllAttrPages generates all attribute sub-pages and the index page
func (g *CCAttrGenerator) GenerateAllAttrPages() error {
	attrDir := filepath.Join(g.OutputDir, "attr")

	// Generate _index.md (only if not existing)
	if err := g.generateIndexPage(filepath.Join(attrDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate attr index: %w", err)
	}

	type pageGen struct {
		name string
		fn   func(string) error
	}
	pages := []pageGen{
		{"load", g.GenerateLoadPage},
		{"ddl", g.GenerateDDLPage},
		{"deps", g.GenerateDepsPage},
		{"multi", g.GenerateMultiPage},
		{"fork", g.GenerateForkPage},
	}

	for _, p := range pages {
		if err := p.fn(filepath.Join(attrDir, p.name+".md")); err != nil {
			return fmt.Errorf("failed to generate %s page: %w", p.name, err)
		}
		logrus.Infof("Generated attr page: %s", p.name)
	}

	logrus.Infof("Successfully generated all attr pages")
	return nil
}

// generateIndexPage generates the attr/_index.md page
func (g *CCAttrGenerator) generateIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	content := `---
title: "扩展属性"
linkTitle: "扩展属性"
description: "按属性筛选的扩展列表"
weight: 300
icon: fa-solid fa-tags
---
`
	return WriteMarkdownFile(outputPath, content)
}


// GenerateLoadPage generates the load.md page (extensions needing shared_preload_libraries)
func (g *CCAttrGenerator) GenerateLoadPage(outputPath string) error {
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if ext.NeedLoad {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(`---
title: "动态加载"
linkTitle: "动态加载"
description: "需要动态加载的 PostgreSQL 扩展"
weight: 10
---

`)
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展需要在 [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) 中动态加载，才能正常使用。\n\n", len(exts)))
	b.WriteString("也就是说，您需要修改 PostgreSQL 配置文件 `postgresql.conf` 中的 [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) 参数，将扩展的库名添加进去，然后重启数据库才能生效。\n\n")

	b.WriteString("| **扩展名** | **动态库名** | **描述** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range exts {
		desc := SanitizeText(ext.GetZhDesc())
		libName := fmt.Sprintf("`%s`", ext.GetLibName())
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s |\n",
			ext.Name, ext.Name, libName, desc))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateDDLPage generates the ddl.md page (extensions not needing CREATE EXTENSION)
func (g *CCAttrGenerator) GenerateDDLPage(outputPath string) error {
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if !ext.NeedDDL {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(`---
title: "无头扩展"
linkTitle: "无头扩展"
description: "不需要 CREATE EXTENSION 的 PostgreSQL 扩展"
weight: 20
slug: ddl
---

`)
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展不需要执行 `CREATE EXTENSION` 即可使用。\n\n", len(exts)))
	b.WriteString("这些扩展通常以共享库（Hook）或独立工具的形式存在，安装后直接通过配置参数启用或自动生效，无需在数据库中显式创建扩展对象。\n\n")

	b.WriteString("| **扩展名** | **扩展包名** | **版本** | **属性** | **描述** |\n")
	b.WriteString("|:-----------|:-------------|:--------:|:--------:|:---------|\n")
	for _, ext := range exts {
		version := ext.GetVersion()
		desc := SanitizeText(ext.GetZhDesc())
		pkgLink := ext.GetPkgURLLink()
		attr := fmt.Sprintf("`%s`", ext.GetAttributeBadge())
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s |\n",
			ext.Name, ext.Name, pkgLink, version, attr, desc))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateDepsPage generates the deps.md page (extensions with dependencies)
func (g *CCAttrGenerator) GenerateDepsPage(outputPath string) error {
	allExts := g.Cache.ReadyExtensions()

	// Upstream: extensions that require other extensions
	var upstream []*Extension
	for _, ext := range allExts {
		if len(ext.Requires) > 0 {
			upstream = append(upstream, ext)
		}
	}

	// Downstream: extensions that are required by other extensions
	var downstream []*Extension
	for _, ext := range allExts {
		if len(ext.RequireBy) > 0 {
			downstream = append(downstream, ext)
		}
	}

	var b strings.Builder
	b.WriteString(`---
title: "依赖关系"
linkTitle: "依赖关系"
description: "具有扩展依赖关系的 PostgreSQL 扩展"
weight: 30
---

`)
	b.WriteString(fmt.Sprintf("共有 **%d** 个扩展依赖其他扩展，**%d** 个扩展被其他扩展所依赖。\n\n", len(upstream), len(downstream)))

	// Upstream table
	b.WriteString("## 上游依赖\n\n")
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展需要先安装其他扩展才能使用：\n\n", len(upstream)))
	b.WriteString("| **扩展名** | **上游依赖** | **描述** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range upstream {
		b.WriteString(g.depsRow(ext, ext.Requires))
	}
	b.WriteString("{.ext-table}\n\n")

	// Downstream table
	b.WriteString("## 下游依赖\n\n")
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展被其他扩展所依赖：\n\n", len(downstream)))
	b.WriteString("| **扩展名** | **下游依赖** | **描述** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range downstream {
		b.WriteString(g.depsRow(ext, ext.RequireBy))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// depsRow generates a table row with dependency links
func (g *CCAttrGenerator) depsRow(ext *Extension, deps []string) string {
	desc := SanitizeText(ext.GetZhDesc())

	// Format dependency links
	depLinks := make([]string, 0, len(deps))
	for _, dep := range deps {
		depLinks = append(depLinks, CCExtLink(dep))
	}
	depStr := strings.Join(depLinks, " ")

	return fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s |\n",
		ext.Name, ext.Name, depStr, desc)
}

// GenerateMultiPage generates the multi.md page (packages containing multiple extensions)
func (g *CCAttrGenerator) GenerateMultiPage(outputPath string) error {
	// Group extensions by package, only consider packages with multiple extensions
	pkgExts := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		pkgExts[ext.Pkg] = append(pkgExts[ext.Pkg], ext)
	}

	// Collect packages with multiple extensions, sorted by lead extension ID
	type multiPkg struct {
		Pkg  string
		Lead *Extension
		Exts []*Extension
	}
	var pkgs []multiPkg
	for pkg, exts := range pkgExts {
		if len(exts) < 2 {
			continue
		}
		// Find lead extension
		var lead *Extension
		for _, ext := range exts {
			if ext.Lead {
				lead = ext
				break
			}
		}
		if lead == nil {
			lead = exts[0]
		}
		pkgs = append(pkgs, multiPkg{Pkg: pkg, Lead: lead, Exts: exts})
	}
	sort.Slice(pkgs, func(i, j int) bool {
		return pkgs[i].Lead.ID < pkgs[j].Lead.ID
	})

	// Count total extensions
	totalExts := 0
	for _, p := range pkgs {
		totalExts += len(p.Exts)
	}

	var b strings.Builder
	b.WriteString(`---
title: "多扩展包"
linkTitle: "多扩展包"
description: "包含多个扩展的 PostgreSQL 扩展包"
weight: 40
---

`)
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展包中包含多个扩展，共计 **%d** 个扩展。\n\n", len(pkgs), totalExts))
	b.WriteString("在安装这些包时，您将同时获得包中的所有扩展。主扩展用粗体标出。\n\n")

	for _, p := range pkgs {
		b.WriteString(fmt.Sprintf("### %s\n\n", p.Pkg))

		pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", p.Pkg, p.Lead.Name)
		b.WriteString(fmt.Sprintf("%s 扩展包共有 **%d** 个扩展：\n\n", pkgLink, len(p.Exts)))

		b.WriteString("| **ID** | **扩展名** | **版本** | **属性** | **模式** | **描述** |\n")
		b.WriteString("|:------:|:-----------|:--------:|:--------:|:---------|:---------|\n")
		for _, ext := range p.Exts {
			version := ext.GetVersion()
			desc := SanitizeText(ext.GetZhDesc())
			attr := fmt.Sprintf("`%s`", ext.GetAttributeBadge())

			schema := "-"
			if len(ext.Schemas) > 0 {
				schema = fmt.Sprintf("`%s`", ext.Schemas[0])
			}

			name := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)
			if ext.Lead {
				name = fmt.Sprintf("[**`%s`**](/ext/e/%s)", ext.Name, ext.Name)
			}
			b.WriteString(fmt.Sprintf("| %d | %s | `%s` | %s | %s | %s |\n",
				ext.ID, name, version, attr, schema, desc))
		}
		b.WriteString("{.ext-table}\n\n")
	}

	return WriteMarkdownFile(outputPath, b.String())
}

// kernelInfo holds display name and documentation link for a PostgreSQL kernel fork
type kernelInfo struct {
	Name string
	Link string
}

// knownKernels maps kernel key (from extra.kernel) to display info
var knownKernels = map[string]kernelInfo{
	"babelfish":  {Name: "Babelfish", Link: "/docs/pgsql/kernel/babelfish"},
	"oriole":     {Name: "OrioleDB", Link: "/docs/pgsql/kernel/orioledb"},
	"percona":    {Name: "Percona", Link: "/docs/pgsql/kernel/percona"},
	"ivory":      {Name: "IvorySQL", Link: "/docs/pgsql/kernel/ivorysql"},
	"openhalodb": {Name: "openHalo", Link: "/docs/pgsql/kernel/openhalo"},
	"pgedge":     {Name: "pgEdge", Link: "/docs/pgsql/kernel/pgedge"},
	"polardb":    {Name: "PolarDB PG", Link: "/docs/pgsql/kernel/polardb"},
	"citus":      {Name: "Citus", Link: "/docs/pgsql/kernel/citus"},
	"agensgraph": {Name: "AgensGraph", Link: "/docs/pgsql/kernel/agensgraph"},
	"neon":       {Name: "Neon", Link: "/docs/pgsql/kernel/neon"},
}

// GenerateForkPage generates the fork.md page (extensions forked from kernel)
func (g *CCAttrGenerator) GenerateForkPage(outputPath string) error {
	// Group extensions by kernel
	kernelExts := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		if ext.Extra == nil {
			continue
		}
		kernel, ok := ext.Extra["kernel"]
		if !ok {
			continue
		}
		kernelStr, ok := kernel.(string)
		if !ok || kernelStr == "" {
			continue
		}
		kernelExts[kernelStr] = append(kernelExts[kernelStr], ext)
	}

	if len(kernelExts) == 0 {
		return WriteMarkdownFile(outputPath, "---\ntitle: \"分支扩展\"\nweight: 50\n---\n\n暂无分支扩展。\n")
	}

	// Sort kernel names for stable output
	kernelNames := make([]string, 0, len(kernelExts))
	for k := range kernelExts {
		kernelNames = append(kernelNames, k)
	}
	sort.Strings(kernelNames)

	// Count total
	totalExts := 0
	for _, exts := range kernelExts {
		totalExts += len(exts)
	}

	var b strings.Builder
	b.WriteString(`---
title: "分支扩展"
linkTitle: "分支扩展"
description: "基于 PostgreSQL 内核分支的扩展"
weight: 50
---

`)
	b.WriteString(fmt.Sprintf("以下 **%d** 个扩展基于 **%d** 种不同的 PostgreSQL 内核分支。\n\n", totalExts, len(kernelNames)))
	b.WriteString("这些扩展需要使用特定的 PostgreSQL 内核分支版本，而非原版 PostgreSQL 内核。\n\n")

	for _, kernel := range kernelNames {
		exts := kernelExts[kernel]
		info, known := knownKernels[kernel]
		if known {
			b.WriteString(fmt.Sprintf("## %s\n\n", info.Name))
			b.WriteString(fmt.Sprintf("以下扩展基于 [**%s**](%s) 内核分支：\n\n", info.Name, info.Link))
		} else {
			b.WriteString(fmt.Sprintf("## %s\n\n", kernel))
			b.WriteString(fmt.Sprintf("以下扩展基于 **%s** 内核分支：\n\n", kernel))
		}
		b.WriteString(CCExtensionTable(exts))
	}

	return WriteMarkdownFile(outputPath, b.String())
}
