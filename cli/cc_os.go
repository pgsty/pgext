/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC OS Generator - generates OS-specific availability pages for pigsty.cc (Chinese only)
Uses CSS classes and {.ext-table} for consistent styling
*/
package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

// CCOSGenerator generates OS availability pages for pigsty.cc
type CCOSGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewCCOSGenerator creates a new CC OS generator
func NewCCOSGenerator(cache *ExtensionCache, outputDir string) *CCOSGenerator {
	return &CCOSGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateOSPage generates an OS-specific availability page (Chinese only)
func (g *CCOSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
	osInfo, err := g.getOSInfo(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to get OS info for %s: %w", osName, err)
	}

	packages, err := g.loadOSPackages(ctx, osName)
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

// getOSInfo retrieves OS information from database
func (g *CCOSGenerator) getOSInfo(ctx context.Context, osName string) (*OSInfo, error) {
	query := `SELECT os, os_full, os_major, os_arch, active FROM pgext.os WHERE os = $1 AND active = true`

	var info OSInfo
	err := QueryRowContext(ctx, query, osName).Scan(
		&info.OS, &info.Desc, &info.Major, &info.Arch, &info.Active)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// loadOSPackages loads all packages with lead extensions for a specific OS
func (g *CCOSGenerator) loadOSPackages(ctx context.Context, osName string) ([]*OSPackageInfo, error) {
	query := `
		WITH lead_extensions AS (
			SELECT DISTINCT e.id, e.pkg, e.name AS lead
			FROM pgext.extension e
			WHERE e.lead = true
		)
		SELECT p.pg, p.pkg, le.lead, le.id, p.name, p.state, p.org, p.version, p.count
		FROM pgext.pkg p
		JOIN lead_extensions le ON p.pkg = le.pkg
		WHERE p.os = $1
		ORDER BY le.id, p.pg DESC
	`

	rows, err := QueryContext(ctx, query, osName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pkgMap := make(map[string]*OSPackageInfo)

	for rows.Next() {
		var pg, extID int
		var pkgName, lead string
		info := &PkgInfo{}

		err := rows.Scan(&pg, &pkgName, &lead, &extID, &info.Name,
			&info.State, &info.Org, &info.Version, &info.Count)
		if err != nil {
			return nil, err
		}

		info.PG = pg
		info.OS = osName
		info.Pkg = pkgName
		info.Ext = lead

		if _, exists := pkgMap[pkgName]; !exists {
			pkgMap[pkgName] = &OSPackageInfo{
				ID:     extID,
				Pkg:    pkgName,
				Lead:   lead,
				PGData: make(map[int]*PkgInfo),
			}
		}

		pkgMap[pkgName].PGData[pg] = info
	}

	var packages []*OSPackageInfo
	for _, pkg := range pkgMap {
		packages = append(packages, pkg)
	}

	sortOSPackages(packages)

	return packages, nil
}

// generateOSContent generates the markdown content for an OS page
func (g *CCOSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	b.WriteString(g.generateOSFrontmatter(osInfo))
	b.WriteString(g.generateOSOverview(osInfo, packages))
	b.WriteString(g.generateOSAvailabilityMatrix(packages, osInfo.OS))

	return b.String()
}

// generateOSFrontmatter generates the YAML frontmatter
func (g *CCOSGenerator) generateOSFrontmatter(osInfo *OSInfo) string {
	weight := computeOSWeight(osInfo.OS)

	// Icon: brand icon based on OS family, flipped for aarch64
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

	return fmt.Sprintf(`---
title: "Linux %s"
description: "%s 系统的 PostgreSQL 扩展可用性"
weight: %d
icon: %s
---

`, osInfo.OS, osInfo.Desc, weight, iconBrand)
}

// generateOSOverview generates the overview section
func (g *CCOSGenerator) generateOSOverview(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	// Count available packages (at least one AVAIL PG version)
	availPkgs := make(map[string]bool)
	for _, pkg := range packages {
		for _, pgData := range pkg.PGData {
			if pgData.State.Valid && pgData.State.String == "AVAIL" {
				availPkgs[pkg.Pkg] = true
				break
			}
		}
	}

	// Count non-contrib extensions belonging to available packages
	extCount := 0
	for _, ext := range g.Cache.Extensions {
		if ext.IsReady() && !ext.Contrib && availPkgs[ext.Pkg] {
			extCount++
		}
	}

	// Determine alias YAML file URL
	aliasFile := fmt.Sprintf("%s.yml", osInfo.OS)
	aliasURL := fmt.Sprintf("https://github.com/pgsty/pigsty/blob/main/roles/node_id/vars/%s", aliasFile)

	b.WriteString(fmt.Sprintf("当前系统共有 **%d** 个非 contrib 扩展可用，分布于 **%d** 个扩展包中。", extCount, len(availPkgs)))
	b.WriteString(fmt.Sprintf("完整可用的 [**包别名**](/docs/pgsql/config/alias/) 请参考 [**`%s`**](%s)。\n\n", aliasFile, aliasURL))

	return b.String()
}

// generateOSAvailabilityMatrix generates the availability matrix using pgext_os_matrix shortcode
func (g *CCOSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo, osName string) string {
	var b strings.Builder

	b.WriteString("{{< pgext_os_matrix >}}\n")

	// Header row
	b.WriteString("| **PKG / PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")

	// Separator
	b.WriteString("|:---|")
	for range g.Cache.PGVersions {
		b.WriteString(":--:|")
	}
	b.WriteString("\n")

	// Data rows
	for _, ospkg := range packages {
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) |", ospkg.Pkg, ospkg.Lead))

		for _, pg := range g.Cache.PGVersions {
			cell := g.generateOSMatrixCell(ospkg, pg, osName)
			b.WriteString(fmt.Sprintf(" %s |", cell))
		}
		b.WriteString("\n")
	}

	b.WriteString("{{< /pgext_os_matrix >}}\n\n")

	return b.String()
}

// generateOSMatrixCell generates a cell as "STATE REPO VERSION COUNT" for pgext_matrix
func (g *CCOSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int, osName string) string {
	// Infer the expected repo from extension metadata (rpm_repo / deb_repo)
	fallbackRepo := "-"
	if ext, ok := g.Cache.PkgMap[ospkg.Pkg]; ok {
		fallbackRepo = InferRepo(ext, osName)
	}

	pkg, exists := ospkg.PGData[pg]
	if !exists {
		return fmt.Sprintf("MISS %s - 0", fallbackRepo)
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

	// MISS takes highest priority when no actual package
	if version == "-" && state != "AVAIL" && state != "HIDE" {
		state = "MISS"
	}

	count := int64(1)
	if pkg.Count.Valid {
		count = pkg.Count.Int64
	}

	return fmt.Sprintf("%s %s %s %d", state, org, version, count)
}


// GenerateAllOSPages generates OS pages for all active OS distributions
func (g *CCOSGenerator) GenerateAllOSPages(ctx context.Context) error {
	osList, err := GetActiveOS(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active OS list: %w", err)
	}

	if len(osList) == 0 {
		return fmt.Errorf("no active OS distributions found")
	}

	logrus.Infof("Generating OS pages for %d distributions...", len(osList))

	type result struct {
		os      string
		success bool
		err     error
	}

	results := make(chan result, len(osList))
	const maxParallel = 8
	semaphore := make(chan struct{}, maxParallel)
	var wg sync.WaitGroup

	for _, osName := range osList {
		wg.Add(1)
		go func(os string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			err := g.GenerateOSPage(ctx, os)
			if err != nil {
				logrus.Errorf("Failed to generate OS page for %s: %v", os, err)
				results <- result{os: os, success: false, err: err}
			} else {
				results <- result{os: os, success: true, err: nil}
			}
		}(osName)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	successCount := 0
	failedOS := []string{}
	for res := range results {
		if res.success {
			successCount++
		} else {
			failedOS = append(failedOS, res.os)
		}
	}

	logrus.Infof("Successfully generated %d/%d OS pages", successCount, len(osList))

	if len(failedOS) > 0 {
		logrus.Warnf("Failed to generate pages for: %v", failedOS)
	}

	if successCount == 0 {
		return fmt.Errorf("failed to generate any OS pages")
	}
	return nil
}

// computeOSWeight computes Hugo frontmatter weight for an OS page
// Pattern: EL base=700 offset=(major-6)*10, Debian base=800 offset=(major-10)*10, Ubuntu base=900 offset=(major-18)/2*10
// aarch64 adds +1 to the weight
func computeOSWeight(osName string) int {
	parts := strings.SplitN(osName, ".", 2)
	code := parts[0]

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

	if len(parts) > 1 && parts[1] == "aarch64" {
		base++
	}
	return base
}
