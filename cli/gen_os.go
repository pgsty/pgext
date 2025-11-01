/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// OSGenerator generates Hugo content for OS-specific extension availability
type OSGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	DB        *pgxpool.Pool
}

// NewOSGenerator creates a new OS generator
func NewOSGenerator(cache *ExtensionCache, outputDir string) *OSGenerator {
	return &OSGenerator{
		Cache:     cache,
		OutputDir: outputDir,
		DB:        DB, // Use the global DB connection
	}
}

// OSPackageInfo represents package availability for a specific OS
type OSPackageInfo struct {
	ID     int              // Extension ID for sorting
	Pkg    string           // Extension package name
	Lead   string           // Lead extension name
	PGData map[int]*PkgInfo // Package info by PG version
}

// GenerateOSPage generates a single OS-specific availability page
func (g *OSGenerator) GenerateOSPage(ctx context.Context, osName string) error {
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

	// Generate content
	content := g.generateOSContent(osInfo, packages)

	// Write to file
	outputPath := filepath.Join(g.OutputDir, fmt.Sprintf("%s.md", osName))
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	logrus.Infof("Generated OS page for %s with %d packages", osName, len(packages))
	return nil
}

// getOSInfo retrieves OS information from database
func (g *OSGenerator) getOSInfo(ctx context.Context, osName string) (*OSInfo, error) {
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
func (g *OSGenerator) loadOSPackages(ctx context.Context, osName string) ([]*OSPackageInfo, error) {
	// Query to get all packages for this OS with lead extensions and their IDs
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

	// Sort packages by lead extension name
	sortOSPackages(packages)

	return packages, nil
}

// sortOSPackages sorts packages by extension ID
func sortOSPackages(packages []*OSPackageInfo) {
	// Sort by extension ID
	for i := 0; i < len(packages)-1; i++ {
		for j := i + 1; j < len(packages); j++ {
			if packages[i].ID > packages[j].ID {
				packages[i], packages[j] = packages[j], packages[i]
			}
		}
	}
}

// generateOSContent generates the markdown content for an OS page
func (g *OSGenerator) generateOSContent(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	// Frontmatter
	b.WriteString(g.generateOSFrontmatter(osInfo))

	// Overview
	b.WriteString(g.generateOSOverview(osInfo, packages))

	// Availability Matrix
	b.WriteString(g.generateOSAvailabilityMatrix(packages))

	return b.String()
}

// generateOSFrontmatter generates the YAML frontmatter
func (g *OSGenerator) generateOSFrontmatter(osInfo *OSInfo) string {
	weight := 100 // Default weight, can be adjusted based on OS priority

	// Adjust weight based on OS (newer/popular OSes get lower weight = higher priority)
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
description: "PostgreSQL Extension Availability for %s"
weight: %d
width: full
---

# %s

`, osInfo.OS, osInfo.OS, osInfo.Desc, weight, osInfo.Desc)
}

// generateOSOverview generates the overview section
func (g *OSGenerator) generateOSOverview(osInfo *OSInfo, packages []*OSPackageInfo) string {
	var b strings.Builder

	b.WriteString("## Overview\n\n")
	b.WriteString(fmt.Sprintf("This page shows the availability of PostgreSQL extensions for **%s**.\n\n", osInfo.Desc))
	b.WriteString(fmt.Sprintf("- **OS Code**: `%s`\n", osInfo.OS))
	b.WriteString(fmt.Sprintf("- **Architecture**: `%s`\n", osInfo.Arch))
	b.WriteString(fmt.Sprintf("- **Total Packages**: %d\n\n", len(packages)))

	return b.String()
}

// generateOSAvailabilityMatrix generates the availability matrix table
func (g *OSGenerator) generateOSAvailabilityMatrix(packages []*OSPackageInfo) string {
	var b strings.Builder

	b.WriteString("## Extension Availability Matrix\n\n")

	// Table header
	b.WriteString("| **Extension** / **PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf("       **PG%d**        |", pg))
	}
	b.WriteString("\n")

	// Table separator
	b.WriteString("|:----------------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":--------------------:|")
	}
	b.WriteString("\n")

	// Generate rows for each package
	for _, ospkg := range packages {
		// Extension link
		extLink := fmt.Sprintf("[%s](/e/%s)", ospkg.Lead, ospkg.Lead)
		b.WriteString(fmt.Sprintf("| %s |", extLink))

		// Generate cells for each PG version
		for _, pg := range g.Cache.PGVersions {
			cell := g.generateOSMatrixCell(ospkg, pg)
			b.WriteString(fmt.Sprintf(" %s |", cell))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// generateOSMatrixCell generates a single cell in the availability matrix
func (g *OSGenerator) generateOSMatrixCell(ospkg *OSPackageInfo, pg int) string {
	pkg, exists := ospkg.PGData[pg]
	if !exists {
		// Not available - MISS state
		return Badge("✗", "red", "", "", "")
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

	// Generate badge based on state
	switch state {
	case "AVAIL":
		if version != "" && org != "" {
			// Show org and version
			text := fmt.Sprintf("%s %s", org, version)
			color := getOSBadgeColor(state, org)
			return Badge(text, color, "", "", "")
		} else if version != "" {
			return Badge(version, "blue", "", "", "")
		}
		return Badge("✓", "green", "", "", "")
	case "MISS":
		return Badge("✗", "red", "", "", "")
	case "HIDE":
		if version != "" && org != "" {
			text := fmt.Sprintf("%s %s", org, version)
			return Badge(text, "gray", "", "", "")
		}
		return Badge("HIDE", "gray", "", "", "")
	case "THROW":
		return Badge("THROW", "purple", "", "", "")
	case "BREAK":
		return Badge("BREAK", "orange", "", "", "")
	case "TBD":
		return Badge("TBD", "yellow", "", "", "")
	default:
		return Badge("?", "gray", "", "", "")
	}
}

// getOSBadgeColor determines the badge color based on state and org
func getOSBadgeColor(state, org string) string {
	// First check state
	switch state {
	case "MISS":
		return "red"
	case "HIDE":
		return "" // No color (gray)
	case "THROW":
		return "purple"
	case "BREAK":
		return "orange"
	case "AVAIL":
		// For AVAIL, check org
		switch strings.ToUpper(org) {
		case "PGDG":
			return "blue"
		case "PIGSTY":
			return "green"
		default:
			return "yellow"
		}
	default:
		return ""
	}
}

// OSInfo represents OS information
type OSInfo struct {
	OS     string
	Desc   string
	Major  *int // os_major can be NULL
	Arch   string
	Active bool
}
