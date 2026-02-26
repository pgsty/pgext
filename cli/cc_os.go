/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC OS Generator - generates OS-specific availability pages for pigsty.cc (Chinese only)
*/
package cli

import (
	"context"
	"fmt"
	"os"
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
// Uses flat file structure: os/xxx.md instead of os/xxx/_index.md
func (g *CCOSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
	// Validate OS exists and is active
	osInfo, err := g.getOSInfo(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to get OS info for %s: %w", osName, err)
	}

	// Load all packages for this OS
	packages, err := g.loadOSPackages(ctx, osName)
	if err != nil {
		return fmt.Errorf("failed to load packages for %s: %w", osName, err)
	}

	// Create OS directory (os/)
	osDir := filepath.Join(g.OutputDir, "os")
	if err := os.MkdirAll(osDir, 0755); err != nil {
		return fmt.Errorf("failed to create OS directory: %w", err)
	}

	// Generate Chinese content only
	content := g.generateOSContent(osInfo, packages)
	outputPath := filepath.Join(osDir, osName+".md")
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
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

	// Organize packages by pkg name
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

		// Get or create package info
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

	// Convert map to sorted slice
	var packages []*OSPackageInfo
	for _, pkg := range pkgMap {
		packages = append(packages, pkg)
	}

	// Sort packages by extension ID
	sortOSPackages(packages)

	return packages, nil
}

// generateOSContent generates the markdown content for an OS page
func (g *CCOSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	// Frontmatter
	b.WriteString(g.generateOSFrontmatter(osInfo))

	// Overview
	b.WriteString(g.generateOSOverview(osInfo, packages))

	// Availability Matrix
	b.WriteString(g.generateOSAvailabilityMatrix(packages, osInfo.OS))

	return b.String()
}

// generateOSFrontmatter generates the YAML frontmatter
func (g *CCOSGenerator) generateOSFrontmatter(osInfo *OSInfo) string {
	weight := 100

	// Adjust weight based on OS
	switch osInfo.OS {
	case "el7.x86_64":
		weight = 710
	case "el8.x86_64":
		weight = 720
	case "el8.aarch64":
		weight = 721
	case "el9.x86_64":
		weight = 730
	case "el9.aarch64":
		weight = 731
	case "el10.x86_64":
		weight = 740
	case "el10.aarch64":
		weight = 741
	case "d11.x86_64":
		weight = 810
	case "d11.aarch64":
		weight = 811
	case "d12.x86_64":
		weight = 820
	case "d12.aarch64":
		weight = 821
	case "d13.x86_64":
		weight = 830
	case "d13.aarch64":
		weight = 831
	case "u20.x86_64":
		weight = 910
	case "u20.aarch64":
		weight = 911
	case "u22.x86_64":
		weight = 920
	case "u22.aarch64":
		weight = 921
	case "u24.x86_64":
		weight = 930
	case "u24.aarch64":
		weight = 931
	}

	return fmt.Sprintf(`---
title: "%s"
linkTitle: "%s"
description: "%s 系统的 PostgreSQL 扩展可用性"
weight: %d
---

`, osInfo.OS, osInfo.OS, osInfo.Desc, weight)
}

// generateOSOverview generates the overview section
func (g *CCOSGenerator) generateOSOverview(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("本页面展示了 **%s** (`%s`) 系统的 PostgreSQL 扩展可用性。\n\n", osInfo.Desc, osInfo.OS))
	b.WriteString(fmt.Sprintf("共有 **%d** 个扩展包可用。\n\n", len(packages)))

	return b.String()
}

// generateOSAvailabilityMatrix generates the availability matrix table
func (g *CCOSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo, osName string) string {
	var b strings.Builder

	// Table header
	b.WriteString("| **扩展** / **PG版本** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")

	// Separator
	b.WriteString("|:----------------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":------:|")
	}
	b.WriteString("\n")

	// Generate rows for each package
	for _, ospkg := range packages {
		// Extension link
		extLink := CCExtLinkWithLabel(ospkg.Lead, ospkg.Pkg)
		b.WriteString(fmt.Sprintf("| %s |", extLink))

		// Generate cells for each PG version
		for _, pg := range g.Cache.PGVersions {
			cell := g.generateOSMatrixCell(ospkg, pg, osName)
			b.WriteString(fmt.Sprintf(" %s |", cell))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// generateOSMatrixCell generates a single cell in the availability matrix
func (g *CCOSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int, osName string) string {
	pkg, exists := ospkg.PGData[pg]
	if !exists {
		// Not available - check why
		ext, hasExt := g.Cache.PkgMap[ospkg.Pkg]
		if hasExt {
			// Check if the extension supports this PG version
			pgSupported := false
			pgStr := fmt.Sprintf("%d", pg)
			for _, ver := range ext.PgVer {
				if ver == pgStr {
					pgSupported = true
					break
				}
			}

			if !pgSupported {
				return "✗"
			}

			// Check if the platform provides this extension
			osPrefix := strings.Split(osName, ".")[0]
			isRPMBased := strings.HasPrefix(osPrefix, "el")
			isDEBBased := strings.HasPrefix(osPrefix, "d") || strings.HasPrefix(osPrefix, "u")

			platformUnsupported := false
			if isRPMBased && (!ext.RpmRepo.Valid || ext.RpmRepo.String == "") {
				platformUnsupported = true
			} else if isDEBBased && (!ext.DebRepo.Valid || ext.DebRepo.String == "") {
				platformUnsupported = true
			}

			if platformUnsupported {
				return "⚠"
			}
		}

		return "✗"
	}

	// Get package state
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

	// Generate display based on state
	switch state {
	case "AVAIL":
		if version != "" && org != "" {
			return fmt.Sprintf("%s %s", org, version)
		} else if version != "" {
			return version
		}
		return "✓"
	case "MISS":
		return "✗"
	case "HIDE":
		return "-"
	case "THROW":
		return "THROW"
	case "BREAK":
		return "BREAK"
	case "TBD":
		return "TBD"
	default:
		return "?"
	}
}

// GenerateAllOSPages generates OS pages for all active OS distributions
func (g *CCOSGenerator) GenerateAllOSPages(ctx context.Context) error {
	// Get all active OS
	osList, err := GetActiveOS(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active OS list: %w", err)
	}

	if len(osList) == 0 {
		return fmt.Errorf("no active OS distributions found")
	}

	logrus.Infof("Generating OS pages for %d distributions...", len(osList))

	// Generate pages in parallel
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

	// Collect results
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
