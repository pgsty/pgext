/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// OSGenerator generates Hugo content for OS-specific extension availability.
type OSGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewOSGenerator creates a new OS generator.
func NewOSGenerator(cache *ExtensionCache, outputDir string) *OSGenerator {
	return &OSGenerator{Cache: cache, OutputDir: outputDir}
}

// OSPackageInfo represents package availability for a specific OS.
type OSPackageInfo struct {
	ID     int
	Pkg    string
	Lead   string
	PGData map[int]*PkgInfo
}

// OSInfo represents OS information.
type OSInfo struct {
	OS     string
	Desc   string
	Major  *int
	Arch   string
	Active bool
}

// GenerateOSPage generates OS-specific availability pages in both English and Chinese.
func (g *OSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
	osInfo, err := loadOSInfo(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to get OS info for %s: %w", osName, err)
	}
	packages, err := loadOSPackages(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to load packages for %s: %w", osName, err)
	}

	if err := WriteMarkdownFile(filepath.Join(g.OutputDir, osName+".md"), g.generateOSContent(osInfo, packages, "en")); err != nil {
		return fmt.Errorf("failed to write English file: %w", err)
	}
	if err := WriteMarkdownFile(filepath.Join(g.OutputDir, osName+".zh.md"), g.generateOSContent(osInfo, packages, "zh")); err != nil {
		return fmt.Errorf("failed to write Chinese file: %w", err)
	}

	logrus.Infof("Generated OS pages for %s with %d packages (en/zh)", osName, len(packages))
	return nil
}

// GenerateAllOSPages generates OS pages for all active OS distributions in parallel.
func (g *OSGenerator) GenerateAllOSPages(ctx context.Context) error {
	return runParallelOSPages(ctx, g.GenerateOSPage)
}

func (g *OSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo, locale string) string {
	var b strings.Builder
	b.WriteString(g.generateOSFrontmatter(osInfo, locale))
	b.WriteString(g.generateOSOverview(osInfo, locale))
	b.WriteString(g.generateOSAvailabilityMatrix(packages, locale))
	return b.String()
}

func (g *OSGenerator) generateOSFrontmatter(osInfo *OSInfo, locale string) string {
	weight := computeOSWeight(osInfo.OS)
	desc := fmt.Sprintf("PostgreSQL Extension Availability for %s", osInfo.Desc)
	if locale == "zh" {
		desc = fmt.Sprintf("%s 的 PostgreSQL 扩展可用性", osInfo.Desc)
	}
	return fmt.Sprintf("---\ntitle: \"OS: %s\"\nlinkTitle: \"%s\"\ndescription: \"%s\"\nweight: %d\npage_width: full\n---\n\n",
		osInfo.OS, osInfo.OS, desc, weight)
}

func (g *OSGenerator) generateOSOverview(osInfo *OSInfo, locale string) string {
	if locale == "zh" {
		return fmt.Sprintf("本页面展示了 **%s** (`%s`) 系统的 PostgreSQL 扩展可用性。\n\n", osInfo.Desc, osInfo.OS)
	}
	return fmt.Sprintf("This page shows the availability of PostgreSQL extensions for **%s** (`%s`) Platform.\n\n", osInfo.Desc, osInfo.OS)
}

func (g *OSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo, locale string) string {
	var b strings.Builder
	if locale == "zh" {
		b.WriteString("| **扩展** / **PG版本** |")
	} else {
		b.WriteString("| **Extension** / **PG** |")
	}
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf("       **PG%d**        |", pg))
	}
	b.WriteString("\n|:----------------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":--------------------:|")
	}
	b.WriteString("\n")

	for _, ospkg := range packages {
		b.WriteString(fmt.Sprintf("| {{< ext %q %q >}} |", ospkg.Lead, ospkg.Pkg))
		for _, pg := range g.Cache.PGVersions {
			b.WriteString(fmt.Sprintf(" %s |", g.generateOSMatrixCell(ospkg, pg)))
		}
		b.WriteString("\n")
	}
	b.WriteString("{.matrix}\n")
	return b.String()
}

func (g *OSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int) string {
	pkg, exists := ospkg.PGData[pg]
	if !exists {
		return Badge("N/A", "gray", "", "", "")
	}
	state := "MISS"
	if pkg.State.Valid {
		state = pkg.State.String
	}
	version := ""
	if pkg.Version.Valid {
		version = pkg.Version.String
	}
	org := ""
	if pkg.Org.Valid {
		org = strings.ToUpper(pkg.Org.String)
	}
	switch state {
	case "AVAIL":
		if version != "" && org != "" {
			return Badge(fmt.Sprintf("%s %s", org, version), getBadgeColor(state, org), "", "", "")
		}
		if version != "" {
			return Badge(version, "blue", "", "", "")
		}
		return Badge("✓", "green", "", "", "")
	case "MISS":
		return Badge("✗", "red", "", "", "")
	default:
		return Badge("N/A", "gray", "", "", "")
	}
}

// getOSBadgeColor aliases getBadgeColor for callers/tests.
func getOSBadgeColor(state, org string) string { return getBadgeColor(state, org) }
