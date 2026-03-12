/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

IO OS Generator - generates OS-specific availability pages for pigsty.io (English only)
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

// IOOSGenerator generates OS availability pages for pigsty.io
type IOOSGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
}

// NewIOOSGenerator creates a new IO OS generator
func NewIOOSGenerator(cache *ExtensionCache, outputDir string) *IOOSGenerator {
	return &IOOSGenerator{
		Cache:     cache,
		OutputDir: outputDir,
	}
}

// GenerateOSPage generates an OS-specific availability page (English only)
func (g *IOOSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
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
func (g *IOOSGenerator) getOSInfo(ctx context.Context, osName string) (*OSInfo, error) {
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
func (g *IOOSGenerator) loadOSPackages(ctx context.Context, osName string) ([]*OSPackageInfo, error) {
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
func (g *IOOSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	b.WriteString(g.generateOSFrontmatter(osInfo))
	b.WriteString(g.generateOSOverview(osInfo, packages))
	b.WriteString(g.generateOSAvailabilityMatrix(packages, osInfo.OS))

	return b.String()
}

// generateOSFrontmatter generates the YAML frontmatter
func (g *IOOSGenerator) generateOSFrontmatter(osInfo *OSInfo) string {
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
description: "PostgreSQL extension availability on %s"
weight: %d
icon: %s
---

`, osInfo.OS, osInfo.Desc, weight, iconBrand)
}

// generateOSOverview generates the overview section
func (g *IOOSGenerator) generateOSOverview(osInfo *OSInfo, packages []*OSPackageInfo) string {
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

	b.WriteString(fmt.Sprintf("There are **%d** non-contrib extensions available in **%d** packages on this system.", extCount, len(availPkgs)))
	b.WriteString(fmt.Sprintf(" For the complete [**package aliases**](/docs/pgsql/config/alias/), see [**`%s`**](%s).\n\n", aliasFile, aliasURL))

	return b.String()
}

// generateOSAvailabilityMatrix generates the availability matrix using pgext_os_matrix shortcode
func (g *IOOSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo, osName string) string {
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
func (g *IOOSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int, osName string) string {
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
func (g *IOOSGenerator) GenerateAllOSPages(ctx context.Context) error {
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
