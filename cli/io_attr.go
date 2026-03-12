/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

IO Attr Generator - generates extension attribute pages for pigsty.io (English only)
Sub-pages: load (Preload Required), ddl (Headless Extensions), deps (Dependencies), multi (Multi-Extension Packages), fork (Kernel Fork Extensions)
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

// IOAttrGenerator generates extension attribute pages for pigsty.io
type IOAttrGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewIOAttrGenerator creates a new IO Attr generator
func NewIOAttrGenerator(cache *ExtensionCache, outputDir string) *IOAttrGenerator {
	return &IOAttrGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateAllAttrPages generates all attribute sub-pages and the index page
func (g *IOAttrGenerator) GenerateAllAttrPages() error {
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
func (g *IOAttrGenerator) generateIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	content := `---
title: "Attributes"
linkTitle: "Attributes"
description: "Extensions filtered by attributes"
weight: 300
icon: fa-solid fa-tags
---
`
	return WriteMarkdownFile(outputPath, content)
}

// GenerateLoadPage generates the load.md page (extensions needing shared_preload_libraries)
func (g *IOAttrGenerator) GenerateLoadPage(outputPath string) error {
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if ext.NeedLoad {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(`---
title: "Preloading"
linkTitle: "Preloading"
description: "PostgreSQL extensions that require dynamic loading"
weight: 10
---

`)
	b.WriteString(fmt.Sprintf("The following **%d** extensions require loading in [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) to function properly.\n\n", len(exts)))
	b.WriteString("You need to modify the [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) parameter in `postgresql.conf`, add the extension library, and restart the database.\n\n")

	b.WriteString("| **Extension** | **Library** | **Description** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range exts {
		desc := SanitizeText(ext.GetEnDesc())
		libName := fmt.Sprintf("`%s`", ext.GetLibName())
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s |\n",
			ext.Name, ext.Name, libName, desc))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateDDLPage generates the ddl.md page (extensions not needing CREATE EXTENSION)
func (g *IOAttrGenerator) GenerateDDLPage(outputPath string) error {
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if !ext.NeedDDL {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(`---
title: "Headless"
linkTitle: "Headless"
description: "PostgreSQL extensions that do not require CREATE EXTENSION"
weight: 20
slug: ddl
---

`)
	b.WriteString(fmt.Sprintf("The following **%d** extensions can be used without running `CREATE EXTENSION`.\n\n", len(exts)))
	b.WriteString("These extensions typically exist as shared libraries (hooks) or standalone tools that take effect through configuration parameters.\n\n")

	b.WriteString("| **Extension** | **Package** | **Version** | **Attr** | **Description** |\n")
	b.WriteString("|:-----------|:-------------|:--------:|:--------:|:---------|\n")
	for _, ext := range exts {
		version := ext.GetVersion()
		desc := SanitizeText(ext.GetEnDesc())
		pkgLink := ext.GetPkgURLLink()
		attr := fmt.Sprintf("`%s`", ext.GetAttributeBadge())
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s |\n",
			ext.Name, ext.Name, pkgLink, version, attr, desc))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// GenerateDepsPage generates the deps.md page (extensions with dependencies)
func (g *IOAttrGenerator) GenerateDepsPage(outputPath string) error {
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
title: "Dependencies"
linkTitle: "Dependencies"
description: "PostgreSQL extensions with dependency relationships"
weight: 30
---

`)
	b.WriteString(fmt.Sprintf("**%d** extensions depend on other extensions, **%d** extensions are depended upon by others.\n\n", len(upstream), len(downstream)))

	// Upstream table
	b.WriteString("## Upstream Dependencies\n\n")
	b.WriteString(fmt.Sprintf("The following **%d** extensions require other extensions to be installed first:\n\n", len(upstream)))
	b.WriteString("| **Extension** | **Requires** | **Description** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range upstream {
		b.WriteString(g.depsRow(ext, ext.Requires))
	}
	b.WriteString("{.ext-table}\n\n")

	// Downstream table
	b.WriteString("## Downstream Dependencies\n\n")
	b.WriteString(fmt.Sprintf("The following **%d** extensions are depended upon by other extensions:\n\n", len(downstream)))
	b.WriteString("| **Extension** | **Required By** | **Description** |\n")
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range downstream {
		b.WriteString(g.depsRow(ext, ext.RequireBy))
	}
	b.WriteString("{.ext-table}\n\n")

	return WriteMarkdownFile(outputPath, b.String())
}

// depsRow generates a table row with dependency links
func (g *IOAttrGenerator) depsRow(ext *Extension, deps []string) string {
	desc := SanitizeText(ext.GetEnDesc())

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
func (g *IOAttrGenerator) GenerateMultiPage(outputPath string) error {
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
title: "Multi-Ext PKG"
linkTitle: "Multi-Ext PKG"
description: "PostgreSQL packages containing multiple extensions"
weight: 40
---

`)
	b.WriteString(fmt.Sprintf("The following **%d** packages contain multiple extensions, totaling **%d** extensions.\n\n", len(pkgs), totalExts))
	b.WriteString("When installing these packages, you will get all extensions in the package. The lead extension is shown in bold.\n\n")

	for _, p := range pkgs {
		b.WriteString(fmt.Sprintf("### %s\n\n", p.Pkg))

		pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", p.Pkg, p.Lead.Name)
		b.WriteString(fmt.Sprintf("Package %s contains **%d** extensions:\n\n", pkgLink, len(p.Exts)))

		b.WriteString("| **ID** | **Extension** | **Version** | **Attr** | **Schema** | **Description** |\n")
		b.WriteString("|:------:|:-----------|:--------:|:--------:|:---------|:---------|\n")
		for _, ext := range p.Exts {
			version := ext.GetVersion()
			desc := SanitizeText(ext.GetEnDesc())
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

// GenerateForkPage generates the fork.md page (extensions forked from kernel)
func (g *IOAttrGenerator) GenerateForkPage(outputPath string) error {
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
		return WriteMarkdownFile(outputPath, "---\ntitle: \"Kernel Forks\"\nweight: 50\n---\n\nNo kernel fork extensions found.\n")
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
title: "Kernel Forks"
linkTitle: "Kernel Forks"
description: "Extensions based on PostgreSQL kernel forks"
weight: 50
---

`)
	b.WriteString(fmt.Sprintf("The following **%d** extensions are based on **%d** different PostgreSQL kernel forks.\n\n", totalExts, len(kernelNames)))
	b.WriteString("These extensions require a specific PostgreSQL kernel fork, not the vanilla PostgreSQL kernel.\n\n")

	for _, kernel := range kernelNames {
		exts := kernelExts[kernel]
		info, known := knownKernels[kernel]
		if known {
			b.WriteString(fmt.Sprintf("## %s\n\n", info.Name))
			b.WriteString(fmt.Sprintf("The following extensions are based on the [**%s**](%s) kernel fork:\n\n", info.Name, info.Link))
		} else {
			b.WriteString(fmt.Sprintf("## %s\n\n", kernel))
			b.WriteString(fmt.Sprintf("The following extensions are based on the **%s** kernel fork:\n\n", kernel))
		}
		b.WriteString(IOExtensionTable(exts))
	}

	return WriteMarkdownFile(outputPath, b.String())
}
