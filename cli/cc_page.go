/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

OINK Page Generator - generates extension detail pages for pigsty.cc (ZH) & pigsty.io (EN).
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// CCPageGenerator generates extension detail pages for pigsty.cc (DocsZH) or pigsty.io (DocsEN).
type CCPageGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	StubDir   string
	Locale    DocsLocale
}

func NewCCPageGenerator(cache *ExtensionCache, outputDir, stubDir string) *CCPageGenerator {
	return &CCPageGenerator{Cache: cache, OutputDir: outputDir, StubDir: stubDir, Locale: DocsZH}
}

func NewIOPageGenerator(cache *ExtensionCache, outputDir, stubDir string) *CCPageGenerator {
	return &CCPageGenerator{Cache: cache, OutputDir: outputDir, StubDir: stubDir, Locale: DocsEN}
}

func (g *CCPageGenerator) locale() DocsLocale {
	if g.Locale == "" {
		return DocsZH
	}
	return g.Locale
}

// GenerateExtensionPage generates a single extension detail page.
func (g *CCPageGenerator) GenerateExtensionPage(ctx context.Context, ext *Extension) error {
	stubContent, err := ReadValidatedStub(filepath.Join(g.StubDir, ext.Name+".md"))
	if err != nil {
		return err
	}
	content, err := g.generateExtensionContent(ctx, ext)
	if err != nil {
		return err
	}
	if stubContent != "" {
		content += "\n" + stubContent
	}
	return WriteMarkdownFile(filepath.Join(g.OutputDir, "e", ext.Name+".md"), content)
}

// PruneStaleExtensionPages removes generated detail pages no longer present in the packaged catalog.
func (g *CCPageGenerator) PruneStaleExtensionPages(extensions []*Extension) ([]string, error) {
	if len(extensions) == 0 {
		return nil, fmt.Errorf("refusing to prune extension pages with an empty catalog")
	}

	extDir := filepath.Join(g.OutputDir, "e")
	entries, err := os.ReadDir(extDir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read extension page directory %s: %w", extDir, err)
	}

	expected := make(map[string]struct{}, len(extensions))
	for _, ext := range extensions {
		if ext != nil && ext.Name != "" {
			expected[ext.Name+".md"] = struct{}{}
		}
	}

	var removed []string
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".md") || strings.HasPrefix(name, "_") {
			continue
		}
		if _, ok := expected[name]; ok {
			continue
		}

		path := filepath.Join(extDir, name)
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read candidate stale extension page %s: %w", path, err)
		}
		stem := strings.TrimSuffix(name, ".md")
		if !strings.Contains(string(content), fmt.Sprintf("title: %q", stem)) ||
			!strings.Contains(string(content), fmt.Sprintf("linkTitle: %q", stem)) {
			continue
		}
		if err := os.Remove(path); err != nil {
			return nil, fmt.Errorf("remove stale extension page %s: %w", path, err)
		}
		removed = append(removed, stem)
	}
	sort.Strings(removed)
	return removed, nil
}

func (g *CCPageGenerator) generateExtensionContent(ctx context.Context, ext *Extension) (string, error) {
	packages, err := LoadPackages(ctx, ext.Pkg)
	if err != nil {
		return "", fmt.Errorf("load packages for %s: %w", ext.Name, err)
	}
	binaries, err := LoadBinaries(ctx, ext.Name)
	if err != nil {
		return "", fmt.Errorf("load binaries for %s: %w", ext.Name, err)
	}
	siblings := g.Cache.GetSiblingExtensions(ext.Pkg, ext.Name)
	allExts := append([]*Extension{ext}, siblings...)
	sort.Slice(allExts, func(i, j int) bool { return allExts[i].ID < allExts[j].ID })

	var b strings.Builder
	b.WriteString(g.generateFrontmatter(ext))
	b.WriteString(g.generateCards(ext))
	b.WriteString("\n---------\n\n")
	b.WriteString(g.generateOverview(ext))
	b.WriteString(g.generateExtensionTable(allExts))
	b.WriteString(g.generateRelationships(ext))
	if ext.Comment.Valid && ext.Comment.String != "" {
		b.WriteString(fmt.Sprintf("\n> %s\n\n", ext.Comment.String))
	}
	b.WriteString(g.generatePackages(ext))
	if !ext.Contrib {
		b.WriteString(g.generateAvailability(ext, packages, binaries))
		b.WriteString(g.generateBuild(ext))
	}
	b.WriteString(g.generateInstall(ext))
	return b.String(), nil
}

func (g *CCPageGenerator) generateFrontmatter(ext *Extension) string {
	desc := strings.ReplaceAll(g.locale().extDesc(ext), `"`, `\"`)
	return fmt.Sprintf("---\ntitle: \"%s\"\nlinkTitle: \"%s\"\ndescription: \"%s\"\nweight: %d\n---\n\n",
		ext.Name, ext.Name, desc, ext.ID)
}

func (g *CCPageGenerator) generateCards(ext *Extension) string {
	hasRepo := ext.URL.Valid && ext.URL.String != ""
	hasSource := ext.Source.Valid && ext.Source.String != ""
	if !hasRepo && !hasSource {
		return ""
	}
	l := g.locale()
	var b strings.Builder
	b.WriteString("<div class=\"ext-cards\">\n")
	if hasRepo {
		url := ext.URL.String
		title := url
		if strings.Contains(url, "github.com") {
			parts := strings.Split(strings.TrimRight(url, "/"), "/")
			if len(parts) >= 2 {
				title = parts[len(parts)-2] + "/" + parts[len(parts)-1]
			}
		}
		b.WriteString(fmt.Sprintf("  <a class=\"ext-card ext-card--repo\" href=\"%s\">\n    <div class=\"ext-card__kicker\">%s</div>\n    <div class=\"ext-card__title\">%s</div>\n    <div class=\"ext-card__desc\">%s</div>\n  </a>\n",
			url, l.t("仓库", "Repository"), title, url))
	}
	if hasSource {
		sourceURL := ext.GetSourceURL(l.region())
		fileName := ext.GetSourceFilename()
		b.WriteString(fmt.Sprintf("  <a class=\"ext-card ext-card--source\" href=\"%s\">\n    <div class=\"ext-card__kicker\">%s</div>\n    <div class=\"ext-card__title\">%s</div>\n    <div class=\"ext-card__desc\">%s</div>\n  </a>\n",
			sourceURL, l.t("源码", "Source"), fileName, fileName))
	}
	b.WriteString("</div>\n\n")
	return b.String()
}

func (g *CCPageGenerator) generateOverview(ext *Extension) string {
	l := g.locale()
	pkgLink := fmt.Sprintf("[**`%s`**](/ext/e/%s)", ext.Pkg, ext.GetPkgPageName())
	version, category, license, lang := "-", "-", "-", "-"
	if ext.Version.Valid {
		version = fmt.Sprintf("`%s`", ext.Version.String)
	}
	if ext.Category.Valid {
		category = CCCategoryBadge(ext.Category.String)
	}
	if ext.License.Valid {
		license = CCLicenseBadge(ext.License.String)
	}
	if ext.Lang.Valid {
		lang = CCLanguageBadge(ext.Lang.String)
	}

	var b strings.Builder
	b.WriteString(l.t("## 概览\n\n", "## Overview\n\n"))
	b.WriteString(l.t("| **扩展包名** | **版本** | **分类** | **许可证** | **语言** |\n", "| **Package** | **Version** | **Category** | **License** | **Language** |\n"))
	b.WriteString("|:---------------------------------------------------:|:-------:|:--------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------:|:--------------------------------------------------------------------:|\n")
	b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n{.ext-table}\n\n", pkgLink, version, category, license, lang))
	return b.String()
}

func (g *CCPageGenerator) generateExtensionTable(allExts []*Extension) string {
	l := g.locale()
	var b strings.Builder
	b.WriteString(l.t(
		"|  ID   | **扩展名** | **Bin** | **Lib** | **Load** | **Create** | **Trust** | **Reloc** | **模式** |\n",
		"|  ID   | **Extension** | **Bin** | **Lib** | **Load** | **Create** | **Trust** | **Reloc** | **Schema** |\n",
	))
	b.WriteString("|:-----:|:-------------------------------------------------------------------------|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:---------------------------------------------:|:--------------------------------------------:|:--------------------------------------------:|:----------|\n")
	for _, e := range allExts {
		schema := "-"
		if len(e.Schemas) > 0 {
			schema = fmt.Sprintf("`%s`", e.Schemas[0])
		}
		b.WriteString(fmt.Sprintf("| %d  | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			e.ID,
			CCExtBoldLink(e.Name, fmt.Sprintf("/ext/e/%s", e.Name)),
			DocsFlagBadge(l, e.HasBin),
			DocsFlagBadge(l, e.HasLib),
			DocsFlagBadge(l, e.NeedLoad),
			DocsFlagBadge(l, e.NeedDDL),
			DocsOptFlagBadge(l, e.Trusted),
			DocsOptFlagBadge(l, e.Relocatable),
			schema,
		))
	}
	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

func (g *CCPageGenerator) generateRelationships(ext *Extension) string {
	if len(ext.Requires) == 0 && len(ext.RequireBy) == 0 && len(ext.SeeAlso) == 0 {
		return ""
	}
	l := g.locale()
	headerParts := make([]string, 0, len(ext.Requires)+len(ext.SeeAlso))
	for _, req := range ext.Requires {
		headerParts = append(headerParts, CCExtRef(g.Cache, req))
	}
	for _, see := range ext.SeeAlso {
		headerParts = append(headerParts, CCExtRef(g.Cache, see))
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("| **%s** | %s |\n", l.t("相关扩展", "Related"), strings.Join(headerParts, " ")))
	b.WriteString("|:--------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|\n")
	if len(ext.RequireBy) > 0 {
		links := make([]string, len(ext.RequireBy))
		for i, req := range ext.RequireBy {
			links[i] = CCExtRef(g.Cache, req)
		}
		b.WriteString(fmt.Sprintf("| **%s** | %s |\n", l.t("下游依赖", "Depended By"), strings.Join(links, " ")))
	}
	b.WriteString("{.ext-table .ext-table--rel}\n\n")
	return b.String()
}

func (g *CCPageGenerator) generatePackages(ext *Extension) string {
	if ext.Contrib {
		return g.generateContribPackages(ext)
	}
	l := g.locale()
	var b strings.Builder
	b.WriteString(l.t("\n## 版本\n\n", "\n## Version\n\n"))
	b.WriteString(l.t("| 类型 | 仓库 | 版本 | PG 大版本 | 包名 | 依赖 |\n", "| Type | Repo | Version | PG Ver | Package | Deps |\n"))
	b.WriteString("|:----:|:----:|:----:|:------:|:--------:|:----:|\n")

	if ext.Repo.Valid || ext.Version.Valid || len(ext.PgVer) > 0 {
		b.WriteString(g.buildPackageRow("EXT", ext.Category.String, ext.Repo.String, ext.Version.String, ext.Pkg, ext.PgVer, ext.Requires))
	}
	if ext.RpmRepo.Valid || ext.RpmVer.Valid || len(ext.RpmPg) > 0 {
		pkg := ext.Pkg
		if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
			pkg = ext.RpmPkg.String
		}
		b.WriteString(g.buildPackageRow("RPM", ext.Category.String, ext.RpmRepo.String, ext.RpmVer.String, pkg, ext.RpmPg, ext.RpmDeps))
	}
	if ext.DebRepo.Valid || ext.DebVer.Valid || len(ext.DebPg) > 0 {
		pkg := ext.Pkg
		if ext.DebPkg.Valid && ext.DebPkg.String != "" {
			pkg = ext.DebPkg.String
		}
		b.WriteString(g.buildPackageRow("DEB", ext.Category.String, ext.DebRepo.String, ext.DebVer.String, pkg, ext.DebPg, ext.DebDeps))
	}
	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

func (g *CCPageGenerator) buildPackageRow(label, category, repo, version, pattern string, pgVers, deps []string) string {
	repoBadge, verStr, patternStr, depsStr := "-", "-", "-", "-"
	if repo != "" {
		repoBadge = CCRepoBadge(repo)
	}
	if version != "" {
		verStr = fmt.Sprintf("`%s`", version)
	}
	if pattern != "" {
		patternStr = fmt.Sprintf("`%s`", pattern)
	}
	if len(deps) > 0 {
		depList := make([]string, len(deps))
		for i, dep := range deps {
			depList[i] = fmt.Sprintf("`%s`", dep)
		}
		depsStr = strings.Join(depList, ", ")
	}
	catAnchor := ""
	if category != "" {
		catAnchor = "#" + strings.ToLower(category)
	}
	labelLink := fmt.Sprintf("**%s**", label)
	switch label {
	case "EXT":
		labelLink = fmt.Sprintf("[**EXT**](/ext/list%s)", catAnchor)
	case "RPM":
		labelLink = fmt.Sprintf("[**RPM**](/ext/rpm%s)", catAnchor)
	case "DEB":
		labelLink = fmt.Sprintf("[**DEB**](/ext/deb%s)", catAnchor)
	}
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
		labelLink, repoBadge, verStr, CCPGVersShortcode(pgVers), patternStr, depsStr)
}

func (g *CCPageGenerator) generateContribPackages(ext *Extension) string {
	var b strings.Builder
	b.WriteString(g.locale().t("\n## 版本\n\n|", "\n## Version\n\n|"))
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n|")

	availablePgs := make(map[int]bool)
	for _, pg := range ParsePGVersions(ext.PgVer) {
		availablePgs[pg] = true
	}
	version := ext.GetVersion()
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

func (g *CCPageGenerator) generateAvailability(ext *Extension, packages []*PkgInfo, binaries []*Binary) string {
	if len(packages) == 0 {
		return ""
	}
	pkgMap := make(map[string]*PkgInfo, len(packages))
	for _, pkg := range packages {
		key := fmt.Sprintf("%d-%s", pkg.PG, pkg.OS)
		if _, exists := pkgMap[key]; !exists {
			pkgMap[key] = pkg
		}
	}

	var b strings.Builder
	b.WriteString("{{< pgext_matrix >}}\n| **OS / PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:--:|")
	for range g.Cache.PGVersions {
		b.WriteString(":--:|")
	}
	b.WriteString("\n")

	for _, osv := range g.Cache.OSVersions {
		if !osv.Active {
			continue
		}
		b.WriteString(fmt.Sprintf("| %s |", osv.OS))
		for _, pg := range g.Cache.PGVersions {
			if pkg, ok := pkgMap[fmt.Sprintf("%d-%s", pg, osv.OS)]; ok {
				b.WriteString(fmt.Sprintf(" %s |", g.formatAvailCell(pkg, ext, osv.OS)))
			} else {
				b.WriteString(" N/A - - 0 |")
			}
		}
		b.WriteString("\n")
	}

	if len(binaries) > 0 {
		activePG := make(map[int]bool, len(g.Cache.PGVersions))
		for _, pg := range g.Cache.PGVersions {
			activePG[pg] = true
		}
		region := g.locale().region()
		for _, bin := range binaries {
			if !activePG[bin.PG] {
				continue
			}
			org := "-"
			if bin.Org.Valid && bin.Org.String != "" {
				org = bin.Org.String
			}
			size := strings.ReplaceAll(FormatSize(bin.Size), " ", "")
			b.WriteString(fmt.Sprintf("@ %s %d %s %s %s %s %s %s\n",
				bin.OS, bin.PG, bin.Name, bin.File, org, bin.Version, size, bin.GetDownloadURL(region)))
		}
	}
	b.WriteString("{{< /pgext_matrix >}}\n\n")
	return b.String()
}

func (g *CCPageGenerator) formatAvailCell(pkg *PkgInfo, ext *Extension, osName string) string {
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}
	org := InferRepo(ext, osName)
	if pkg.Org.Valid && pkg.Org.String != "" {
		org = strings.ToUpper(pkg.Org.String)
	}
	version := "-"
	if pkg.Version.Valid && pkg.Version.String != "" {
		version = pkg.Version.String
	}
	count := int64(0)
	if pkg.Count.Valid {
		count = pkg.Count.Int64
	}
	return fmt.Sprintf("%s %s %s %d", state, org, version, count)
}

func (g *CCPageGenerator) generateBuild(ext *Extension) string {
	if ext.Contrib {
		return ""
	}
	buildRPM, buildDEB := extBuildFlags(ext)
	if !buildRPM && !buildDEB {
		return ""
	}
	l := g.locale()
	var buildHint string
	switch {
	case buildRPM && buildDEB:
		buildHint = l.t("RPM / DEB 包", "RPM / DEB packages")
	case buildRPM:
		buildHint = l.t("RPM 包", "RPM packages")
	default:
		buildHint = l.t("DEB 包", "DEB packages")
	}

	var b strings.Builder
	b.WriteString(l.t("## 构建\n\n", "## Build\n\n"))
	if l.isZH() {
		b.WriteString(fmt.Sprintf("您可以使用 `pig build` 命令构建 `%s` 扩展的 %s：\n\n", ext.Pkg, buildHint))
	} else {
		b.WriteString(fmt.Sprintf("You can build the %s for `%s` using `pig build`:\n\n", buildHint, ext.Pkg))
	}
	b.WriteString("```bash\n")
	b.WriteString(fmt.Sprintf(l.t("pig build pkg %s         # 构建 %s\n", "pig build pkg %s         # build %s\n"), ext.Pkg, buildHint))
	b.WriteString("```\n\n")
	return b.String()
}

func (g *CCPageGenerator) generateInstall(ext *Extension) string {
	l := g.locale()
	var b strings.Builder
	b.WriteString(l.t("\n## 安装\n\n", "\n## Install\n\n"))

	if ext.Contrib {
		b.WriteString(l.t(
			"> **提示**：这是 PostgreSQL 内核自带的 contrib 扩展\n\n",
			"> **Note**: This is a built-in contrib extension of PostgreSQL\n\n",
		))
		if ext.NeedDDL {
			b.WriteString("```sql\nCREATE EXTENSION " + ext.Name + ";\n```\n")
		}
		return b.String()
	}

	if ext.UsesPGDGOnly() {
		b.WriteString(fmt.Sprintf(l.t(
			"您可以直接安装 `%s` 扩展包的预置二进制包，首先确保 [**PGDG**](/docs/repo/pgdg) 仓库已经添加并启用：\n\n",
			"You can install `%s` directly. First, make sure the [**PGDG**](/docs/repo/pgdg) repository is added and enabled:\n\n",
		), ext.Pkg))
		b.WriteString(l.t(
			"```bash\npig repo add pgdg -u          # 添加 PGDG 仓库并更新缓存\n```\n\n",
			"```bash\npig repo add pgdg -u          # Add PGDG repo and update cache\n```\n\n",
		))
	} else {
		b.WriteString(fmt.Sprintf(l.t(
			"您可以直接安装 `%s` 扩展包的预置二进制包，首先确保 [**PGDG**](/docs/repo/pgdg) 和 [**PIGSTY**](/docs/repo/pgsql) 仓库已经添加并启用：\n\n",
			"You can install `%s` directly. First, make sure the [**PGDG**](/docs/repo/pgdg) and [**PIGSTY**](/docs/repo/pgsql) repositories are added and enabled:\n\n",
		), ext.Pkg))
		b.WriteString(l.t(
			"```bash\npig repo add pgsql -u          # 添加仓库并更新缓存\n```\n\n",
			"```bash\npig repo add pgsql -u          # Add repo and update cache\n```\n\n",
		))
	}

	pgVersions := activePGVersionsFor(g.Cache, ext)
	rpmPkg, debPkg := ext.Pkg, ext.Pkg
	if ext.RpmPkg.Valid && ext.RpmPkg.String != "" {
		rpmPkg = ext.RpmPkg.String
	}
	if ext.DebPkg.Valid && ext.DebPkg.String != "" {
		debPkg = ext.DebPkg.String
	}

	b.WriteString(l.t(
		"使用 [**pig**](https://pig.pgsty.com/zh) 或者是 `apt/yum/dnf` 安装扩展：\n\n",
		"Install the extension using [**pig**](https://pig.pgsty.com) or `apt/yum/dnf`:\n\n",
	))
	b.WriteString(fmt.Sprintf("```bash {tab=%q group=\"extension-install\" value=\"install\"}\npig install %s\n```\n\n",
		l.t("安装", "Install"), ext.Pkg))

	b.WriteString("```bash {tab=\"pig\" value=\"pig\"}\n")
	for _, pg := range pgVersions {
		b.WriteString(fmt.Sprintf("pig ext install -y %s -v %d  # PG %d\n", ext.Pkg, pg, pg))
	}
	b.WriteString("```\n\n")

	if ext.RpmRepo.Valid || ext.RpmVer.Valid || len(ext.RpmPg) > 0 {
		b.WriteString("```bash {tab=\"dnf\" value=\"dnf\"}\n")
		for _, pg := range pgVersions {
			pkgName := strings.ReplaceAll(rpmPkg, "$v", fmt.Sprintf("%d", pg))
			b.WriteString(fmt.Sprintf("dnf install -y %s       # PG %d\n", pkgName, pg))
		}
		b.WriteString("```\n\n")
	}

	if ext.DebRepo.Valid || ext.DebVer.Valid || len(ext.DebPg) > 0 {
		b.WriteString("```bash {tab=\"apt\" value=\"apt\"}\n")
		for _, pg := range pgVersions {
			pkgName := strings.ReplaceAll(debPkg, "$v", fmt.Sprintf("%d", pg))
			b.WriteString(fmt.Sprintf("apt install -y %s   # PG %d\n", pkgName, pg))
		}
		b.WriteString("```\n\n")
	}

	if ext.NeedLoad {
		b.WriteString(l.t("\n**预加载配置**：\n\n", "\n**Preload**:\n\n"))
		libConfig := strings.Join(preloadLibraries(g.Cache, ext), ", ")
		b.WriteString(fmt.Sprintf("```ini\nshared_preload_libraries = '%s';\n```\n\n", libConfig))
	}

	if ext.NeedDDL {
		b.WriteString(l.t("\n**创建扩展**：\n\n", "\n**Create Extension**:\n\n"))
		b.WriteString("```sql\n")
		if len(ext.Requires) > 0 {
			b.WriteString(fmt.Sprintf(l.t(
				"CREATE EXTENSION %s CASCADE;  -- 依赖: %s\n",
				"CREATE EXTENSION %s CASCADE;  -- requires: %s\n",
			), ext.Name, strings.Join(ext.Requires, ", ")))
		} else {
			b.WriteString(fmt.Sprintf("CREATE EXTENSION %s;\n", ext.Name))
		}
		b.WriteString("```\n")
	} else if !ext.NeedLoad {
		b.WriteString(l.t("> 此扩展不需要执行 `CREATE EXTENSION` 语句\n", "> This extension does not require `CREATE EXTENSION`\n"))
	}
	return b.String()
}
