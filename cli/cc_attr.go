/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

OINK Attr Generator - generates extension attribute pages for pigsty.cc (ZH) & pigsty.io (EN).
Sub-pages: load, ddl, deps, multi, fork.
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

// CCAttrGenerator generates extension attribute pages for pigsty.cc (DocsZH) or pigsty.io (DocsEN).
type CCAttrGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	Locale    DocsLocale
}

func NewCCAttrGenerator(cache *ExtensionCache, outputDir string) *CCAttrGenerator {
	return &CCAttrGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsZH}
}

func NewIOAttrGenerator(cache *ExtensionCache, outputDir string) *CCAttrGenerator {
	return &CCAttrGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsEN}
}

func (g *CCAttrGenerator) locale() DocsLocale {
	if g.Locale == "" {
		return DocsZH
	}
	return g.Locale
}

// GenerateAllAttrPages generates all attribute sub-pages and the index page.
func (g *CCAttrGenerator) GenerateAllAttrPages() error {
	attrDir := filepath.Join(g.OutputDir, "attr")
	if err := g.generateIndexPage(filepath.Join(attrDir, "_index.md")); err != nil {
		return fmt.Errorf("failed to generate attr index: %w", err)
	}

	pages := []struct {
		name string
		fn   func(string) error
	}{
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

func (g *CCAttrGenerator) generateIndexPage(outputPath string) error {
	if _, err := os.Stat(outputPath); err == nil {
		logrus.Infof("Keeping existing %s", outputPath)
		return nil
	}
	l := g.locale()
	content := fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 300\nicon: fa-solid fa-tags\n---\n",
		l.t("扩展属性", "Attributes"),
		l.t("扩展属性", "Attributes"),
		l.t("按属性筛选的扩展列表", "Extensions filtered by attributes"),
	)
	return WriteMarkdownFile(outputPath, content)
}

func (g *CCAttrGenerator) GenerateLoadPage(outputPath string) error {
	l := g.locale()
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if ext.NeedLoad {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 10\n---\n\n",
		l.t("动态加载", "Preloading"),
		l.t("动态加载", "Preloading"),
		l.t("需要动态加载的 PostgreSQL 扩展", "PostgreSQL extensions that require dynamic loading"),
	))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展需要在 [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) 中动态加载，才能正常使用。\n\n",
		"The following **%d** extensions require loading in [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) to function properly.\n\n",
	), len(exts)))
	b.WriteString(l.t(
		"也就是说，您需要修改 PostgreSQL 配置文件 `postgresql.conf` 中的 [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) 参数，将扩展的库名添加进去，然后重启数据库才能生效。\n\n",
		"You need to modify the [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES) parameter in `postgresql.conf`, add the extension library, and restart the database.\n\n",
	))
	b.WriteString(l.t("| **扩展名** | **动态库名** | **描述** |\n", "| **Extension** | **Library** | **Description** |\n"))
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range exts {
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | `%s` | %s |\n",
			ext.Name, ext.Name, ext.GetLibName(), SanitizeText(l.extDesc(ext))))
	}
	b.WriteString("{.ext-table}\n\n")
	return WriteMarkdownFile(outputPath, b.String())
}

func (g *CCAttrGenerator) GenerateDDLPage(outputPath string) error {
	l := g.locale()
	var exts []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if !ext.NeedDDL {
			exts = append(exts, ext)
		}
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 20\nslug: ddl\n---\n\n",
		l.t("无头扩展", "Headless"),
		l.t("无头扩展", "Headless"),
		l.t("不需要 CREATE EXTENSION 的 PostgreSQL 扩展", "PostgreSQL extensions that do not require CREATE EXTENSION"),
	))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展不需要执行 `CREATE EXTENSION` 即可使用。\n\n",
		"The following **%d** extensions can be used without running `CREATE EXTENSION`.\n\n",
	), len(exts)))
	b.WriteString(l.t(
		"这些扩展通常以共享库（Hook）或独立工具的形式存在，安装后直接通过配置参数启用或自动生效，无需在数据库中显式创建扩展对象。\n\n",
		"These extensions typically exist as shared libraries (hooks) or standalone tools that take effect through configuration parameters.\n\n",
	))
	b.WriteString(l.t("| **扩展名** | **扩展包名** | **版本** | **属性** | **描述** |\n", "| **Extension** | **Package** | **Version** | **Attr** | **Description** |\n"))
	b.WriteString("|:-----------|:-------------|:--------:|:--------:|:---------|\n")
	for _, ext := range exts {
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | `%s` | %s |\n",
			ext.Name, ext.Name, ext.GetPkgURLLink(), ext.GetVersion(), ext.GetAttributeBadge(), SanitizeText(l.extDesc(ext))))
	}
	b.WriteString("{.ext-table}\n\n")
	return WriteMarkdownFile(outputPath, b.String())
}

func (g *CCAttrGenerator) GenerateDepsPage(outputPath string) error {
	l := g.locale()
	var upstream, downstream []*Extension
	for _, ext := range g.Cache.ReadyExtensions() {
		if len(ext.Requires) > 0 {
			upstream = append(upstream, ext)
		}
		if len(ext.RequireBy) > 0 {
			downstream = append(downstream, ext)
		}
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 30\n---\n\n",
		l.t("依赖关系", "Dependencies"),
		l.t("依赖关系", "Dependencies"),
		l.t("具有扩展依赖关系的 PostgreSQL 扩展", "PostgreSQL extensions with dependency relationships"),
	))
	b.WriteString(fmt.Sprintf(l.t(
		"共有 **%d** 个扩展依赖其他扩展，**%d** 个扩展被其他扩展所依赖。\n\n",
		"**%d** extensions depend on other extensions, **%d** extensions are depended upon by others.\n\n",
	), len(upstream), len(downstream)))

	b.WriteString(l.t("## 上游依赖\n\n", "## Upstream Dependencies\n\n"))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展需要先安装其他扩展才能使用：\n\n",
		"The following **%d** extensions require other extensions to be installed first:\n\n",
	), len(upstream)))
	b.WriteString(l.t("| **扩展名** | **上游依赖** | **描述** |\n", "| **Extension** | **Requires** | **Description** |\n"))
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range upstream {
		b.WriteString(g.depsRow(ext, ext.Requires))
	}
	b.WriteString("{.ext-table}\n\n")

	b.WriteString(l.t("## 下游依赖\n\n", "## Downstream Dependencies\n\n"))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展被其他扩展所依赖：\n\n",
		"The following **%d** extensions are depended upon by other extensions:\n\n",
	), len(downstream)))
	b.WriteString(l.t("| **扩展名** | **下游依赖** | **描述** |\n", "| **Extension** | **Required By** | **Description** |\n"))
	b.WriteString("|:-----------|:-------------|:---------|\n")
	for _, ext := range downstream {
		b.WriteString(g.depsRow(ext, ext.RequireBy))
	}
	b.WriteString("{.ext-table}\n\n")
	return WriteMarkdownFile(outputPath, b.String())
}

func (g *CCAttrGenerator) depsRow(ext *Extension, deps []string) string {
	depLinks := make([]string, 0, len(deps))
	for _, dep := range deps {
		depLinks = append(depLinks, CCExtRef(g.Cache, dep))
	}
	return fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | %s |\n",
		ext.Name, ext.Name, strings.Join(depLinks, " "), SanitizeText(g.locale().extDesc(ext)))
}

func (g *CCAttrGenerator) GenerateMultiPage(outputPath string) error {
	l := g.locale()
	pkgExts := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		pkgExts[ext.Pkg] = append(pkgExts[ext.Pkg], ext)
	}

	type multiPkg struct {
		Pkg  string
		Lead *Extension
		Exts []*Extension
	}
	var pkgs []multiPkg
	totalExts := 0
	for pkg, exts := range pkgExts {
		if len(exts) < 2 {
			continue
		}
		lead := exts[0]
		for _, ext := range exts {
			if ext.Lead {
				lead = ext
				break
			}
		}
		pkgs = append(pkgs, multiPkg{Pkg: pkg, Lead: lead, Exts: exts})
		totalExts += len(exts)
	}
	sort.Slice(pkgs, func(i, j int) bool { return pkgs[i].Lead.ID < pkgs[j].Lead.ID })

	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 40\n---\n\n",
		l.t("多扩展包", "Multi-Ext PKG"),
		l.t("多扩展包", "Multi-Ext PKG"),
		l.t("包含多个扩展的 PostgreSQL 扩展包", "PostgreSQL packages containing multiple extensions"),
	))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展包中包含多个扩展，共计 **%d** 个扩展。\n\n",
		"The following **%d** packages contain multiple extensions, totaling **%d** extensions.\n\n",
	), len(pkgs), totalExts))
	b.WriteString(l.t(
		"在安装这些包时，您将同时获得包中的所有扩展。主扩展用粗体标出。\n\n",
		"When installing these packages, you will get all extensions in the package. The lead extension is shown in bold.\n\n",
	))

	for _, p := range pkgs {
		b.WriteString(fmt.Sprintf("## %s\n\n", p.Pkg))
		pkgLink := fmt.Sprintf("[`%s`](/ext/e/%s)", p.Pkg, p.Lead.Name)
		b.WriteString(fmt.Sprintf(l.t("%s 扩展包共有 **%d** 个扩展：\n\n", "Package %s contains **%d** extensions:\n\n"), pkgLink, len(p.Exts)))
		b.WriteString(l.t("| **ID** | **扩展名** | **版本** | **属性** | **模式** | **描述** |\n", "| **ID** | **Extension** | **Version** | **Attr** | **Schema** | **Description** |\n"))
		b.WriteString("|:------:|:-----------|:--------:|:--------:|:---------|:---------|\n")
		for _, ext := range p.Exts {
			schema := "-"
			if len(ext.Schemas) > 0 {
				schema = fmt.Sprintf("`%s`", ext.Schemas[0])
			}
			name := fmt.Sprintf("[`%s`](/ext/e/%s)", ext.Name, ext.Name)
			if ext.Lead {
				name = fmt.Sprintf("[**`%s`**](/ext/e/%s)", ext.Name, ext.Name)
			}
			b.WriteString(fmt.Sprintf("| %d | %s | `%s` | `%s` | %s | %s |\n",
				ext.ID, name, ext.GetVersion(), ext.GetAttributeBadge(), schema, SanitizeText(l.extDesc(ext))))
		}
		b.WriteString("{.ext-table}\n\n")
	}
	return WriteMarkdownFile(outputPath, b.String())
}

type kernelInfo struct {
	Name string
	Link string
}

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

func (g *CCAttrGenerator) GenerateForkPage(outputPath string) error {
	l := g.locale()
	kernelExts := make(map[string][]*Extension)
	for _, ext := range g.Cache.ReadyExtensions() {
		if ext.Extra == nil {
			continue
		}
		if kernelStr, ok := ext.Extra["kernel"].(string); ok && kernelStr != "" {
			kernelExts[kernelStr] = append(kernelExts[kernelStr], ext)
		}
	}

	if len(kernelExts) == 0 {
		return WriteMarkdownFile(outputPath, fmt.Sprintf("---\ntitle: %q\nweight: 50\n---\n\n%s\n",
			l.t("分支扩展", "Kernel Forks"),
			l.t("暂无分支扩展。", "No kernel fork extensions found."),
		))
	}

	kernelNames := make([]string, 0, len(kernelExts))
	totalExts := 0
	for k, exts := range kernelExts {
		kernelNames = append(kernelNames, k)
		totalExts += len(exts)
	}
	sort.Strings(kernelNames)

	var b strings.Builder
	b.WriteString(fmt.Sprintf("---\ntitle: %q\nlinkTitle: %q\ndescription: %q\nweight: 50\n---\n\n",
		l.t("分支扩展", "Kernel Forks"),
		l.t("分支扩展", "Kernel Forks"),
		l.t("基于 PostgreSQL 内核分支的扩展", "Extensions based on PostgreSQL kernel forks"),
	))
	b.WriteString(fmt.Sprintf(l.t(
		"以下 **%d** 个扩展基于 **%d** 种不同的 PostgreSQL 内核分支。\n\n",
		"The following **%d** extensions are based on **%d** different PostgreSQL kernel forks.\n\n",
	), totalExts, len(kernelNames)))
	b.WriteString(l.t(
		"这些扩展需要使用特定的 PostgreSQL 内核分支版本，而非原版 PostgreSQL 内核。\n\n",
		"These extensions require a specific PostgreSQL kernel fork, not the vanilla PostgreSQL kernel.\n\n",
	))

	for _, kernel := range kernelNames {
		exts := kernelExts[kernel]
		if info, known := knownKernels[kernel]; known {
			b.WriteString(fmt.Sprintf("## %s\n\n", info.Name))
			b.WriteString(fmt.Sprintf(l.t(
				"以下扩展基于 [**%s**](%s) 内核分支：\n\n",
				"The following extensions are based on the [**%s**](%s) kernel fork:\n\n",
			), info.Name, info.Link))
		} else {
			b.WriteString(fmt.Sprintf("## %s\n\n", kernel))
			b.WriteString(fmt.Sprintf(l.t(
				"以下扩展基于 **%s** 内核分支：\n\n",
				"The following extensions are based on the **%s** kernel fork:\n\n",
			), kernel))
		}
		b.WriteString(DocsExtensionTable(l, exts))
	}
	return WriteMarkdownFile(outputPath, b.String())
}
