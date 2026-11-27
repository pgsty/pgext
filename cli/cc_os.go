/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

OINK OS Generator - generates OS-specific availability pages for pigsty.cc (ZH) & pigsty.io (EN).
*/
package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// CCOSGenerator generates OS availability pages for pigsty.cc (DocsZH) or pigsty.io (DocsEN).
type CCOSGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	Locale    DocsLocale
}

func NewCCOSGenerator(cache *ExtensionCache, outputDir string) *CCOSGenerator {
	return &CCOSGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsZH}
}

func NewIOOSGenerator(cache *ExtensionCache, outputDir string) *CCOSGenerator {
	return &CCOSGenerator{Cache: cache, OutputDir: outputDir, Locale: DocsEN}
}

func (g *CCOSGenerator) locale() DocsLocale {
	if g.Locale == "" {
		return DocsZH
	}
	return g.Locale
}

// GenerateOSPage generates an OS-specific availability page.
func (g *CCOSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
	osInfo, err := loadOSInfo(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to get OS info for %s: %w", osName, err)
	}
	packages, err := loadOSPackages(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to load packages for %s: %w", osName, err)
	}

	content := g.generateOSContent(osInfo, packages)
	outputPath := filepath.Join(g.OutputDir, "os", osName+".md")
	if err := WriteMarkdownFile(outputPath, content); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	logrus.Infof("Generated OS page for %s with %d packages", osName, len(packages))
	return nil
}

func (g *CCOSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder
	b.WriteString(g.generateOSFrontmatter(osInfo))
	b.WriteString(g.generateOSOverview(osInfo, packages))
	b.WriteString(g.generateOSAvailabilityMatrix(packages, osInfo.OS))
	return b.String()
}

func (g *CCOSGenerator) generateOSFrontmatter(osInfo *OSInfo) string {
	l := g.locale()
	weight := computeOSWeight(osInfo.OS)

	var iconBrand string
	switch {
	case strings.HasPrefix(osInfo.OS, "el"):
		iconBrand = "fa-brands fa-redhat"
	case strings.HasPrefix(osInfo.OS, "d"):
		iconBrand = "fa-brands fa-debian"
	case strings.HasPrefix(osInfo.OS, "u"):
		iconBrand = "fa-brands fa-ubuntu"
	}
	if strings.Contains(osInfo.OS, "aarch64") {
		iconBrand += " fa-flip-vertical"
	}

	desc := fmt.Sprintf(l.t("%s 系统的 PostgreSQL 扩展可用性", "PostgreSQL extension availability on %s"), osInfo.Desc)
	return fmt.Sprintf("---\ntitle: \"Linux %s\"\ndescription: \"%s\"\nweight: %d\nicon: %s\n---\n\n",
		osInfo.OS, desc, weight, iconBrand)
}

func (g *CCOSGenerator) generateOSOverview(osInfo *OSInfo, packages []*OSPackageInfo) string {
	l := g.locale()
	availPkgs := make(map[string]bool)
	for _, pkg := range packages {
		for _, pgData := range pkg.PGData {
			if pgData.State.Valid && pgData.State.String == "AVAIL" {
				availPkgs[pkg.Pkg] = true
				break
			}
		}
	}

	extCount := 0
	for _, ext := range g.Cache.Extensions {
		if ext.IsReady() && !ext.Contrib && availPkgs[ext.Pkg] {
			extCount++
		}
	}

	aliasFile := fmt.Sprintf("%s.yml", osInfo.OS)
	aliasURL := fmt.Sprintf("https://github.com/pgsty/pigsty/blob/main/roles/node_id/vars/%s", aliasFile)

	var b strings.Builder
	b.WriteString(fmt.Sprintf(l.t(
		"当前系统共有 **%d** 个非 contrib 扩展可用，分布于 **%d** 个扩展包中。",
		"There are **%d** non-contrib extensions available in **%d** packages on this system.",
	), extCount, len(availPkgs)))
	b.WriteString(fmt.Sprintf(l.t(
		"完整可用的 [**包别名**](/docs/pgsql/config/alias/) 请参考 [**`%s`**](%s)。\n\n",
		" For the complete [**package aliases**](/docs/pgsql/config/alias/), see [**`%s`**](%s).\n\n",
	), aliasFile, aliasURL))
	return b.String()
}

func (g *CCOSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo, osName string) string {
	var b strings.Builder
	b.WriteString("{{< pgext_os_matrix >}}\n| **PKG / PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n|:---|")
	for range g.Cache.PGVersions {
		b.WriteString(":--:|")
	}
	b.WriteString("\n")

	for _, ospkg := range packages {
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) |", ospkg.Pkg, ospkg.Lead))
		for _, pg := range g.Cache.PGVersions {
			b.WriteString(fmt.Sprintf(" %s |", g.generateOSMatrixCell(ospkg, pg, osName)))
		}
		b.WriteString("\n")
	}
	b.WriteString("{{< /pgext_os_matrix >}}\n\n")
	return b.String()
}

func (g *CCOSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int, osName string) string {
	fallbackRepo := "-"
	if ext, ok := g.Cache.PkgMap[ospkg.Pkg]; ok {
		fallbackRepo = InferRepo(ext, osName)
	}
	pkg, exists := ospkg.PGData[pg]
	if !exists {
		return "N/A - - 0"
	}
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}
	org := fallbackRepo
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

// GenerateAllOSPages generates OS pages for all active OS distributions.
func (g *CCOSGenerator) GenerateAllOSPages(ctx context.Context) error {
	return runParallelOSPages(ctx, g.GenerateOSPage)
}

// computeOSWeight computes Hugo frontmatter weight for an OS page.
// Pattern: EL base=700 offset=(major-6)*10, Debian base=800 offset=(major-10)*10, Ubuntu base=900 offset=(major-18)/2*10.
// aarch64 adds +1 to the weight.
func computeOSWeight(osName string) int {
	code, arch, _ := strings.Cut(osName, ".")
	var base, major int
	switch {
	case strings.HasPrefix(code, "el"):
		fmt.Sscanf(code, "el%d", &major)
		base = 700 + (major-6)*10
	case strings.HasPrefix(code, "d"):
		fmt.Sscanf(code, "d%d", &major)
		base = 800 + (major-10)*10
	case strings.HasPrefix(code, "u"):
		fmt.Sscanf(code, "u%d", &major)
		base = 900 + (major-18)/2*10
	default:
		base = 100
	}
	if arch == "aarch64" {
		base++
	}
	return base
}
