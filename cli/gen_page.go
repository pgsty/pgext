/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// ExtensionGenerator generates Hugo content for PostgreSQL extensions
type ExtensionGenerator struct {
	Cache     *ExtensionCache
	OutputDir string
	DB        *pgxpool.Pool
}

// NewExtensionGenerator creates a new extension generator
func NewExtensionGenerator(cache *ExtensionCache, outputDir string) *ExtensionGenerator {
	return &ExtensionGenerator{
		Cache:     cache,
		OutputDir: outputDir,
		DB:        DB, // Use the global DB connection
	}
}

// GenerateExtensionPage generates a single extension detail page
func (g *ExtensionGenerator) GenerateExtensionPage(ctx context.Context, ext *Extension) error {
	// Load package and binary data
	packages, err := LoadPackages(ctx, ext.Pkg)
	if err != nil {
		return fmt.Errorf("failed to load packages: %w", err)
	}

	binaries, err := LoadBinaries(ctx, ext.Name)
	if err != nil {
		return fmt.Errorf("failed to load binaries: %w", err)
	}

	siblings := g.Cache.GetSiblingExtensions(ext.Pkg, ext.Name)

	// Generate content
	content := g.generateExtensionContent(ext, packages, binaries, siblings)

	// Append stub content if exists
	stubPath := filepath.Join(filepath.Dir(g.OutputDir), "..", "stub", fmt.Sprintf("%s.md", ext.Name))
	if stubContent, err := os.ReadFile(stubPath); err == nil && len(stubContent) > 0 {
		content += "\n" + string(stubContent)
		logrus.Debugf("Appended stub content for %s", ext.Name)
	}

	// Write to file
	outputPath := filepath.Join(g.OutputDir, fmt.Sprintf("%s.md", ext.Name))
	return os.WriteFile(outputPath, []byte(content), 0644)
}

func (g *ExtensionGenerator) generateExtensionContent(ext *Extension, packages []*PkgInfo, binaries []*Binary, siblings []*Extension) string {
	var b strings.Builder

	// Frontmatter
	b.WriteString(g.generateFrontmatter(ext))

	// Overview section
	b.WriteString(g.generateOverview(ext))

	// Attributes section
	b.WriteString(g.generateAttributes(ext))

	// Relationships section
	if rel := g.generateRelationships(ext, siblings); rel != "" {
		b.WriteString(rel)
	}

	// Comment note
	if note := g.generateCommentNote(ext); note != "" {
		b.WriteString(note)
	}

	// Packages section
	b.WriteString(g.generatePackagesTable(ext, packages))

	// For non-contrib extensions, add detailed availability info
	if !ext.Contrib {
		b.WriteString(g.generateAvailabilityMatrix(packages, binaries))
		b.WriteString(g.generatePackageDetailsTabs(ext.Name, packages, binaries))
	} else {
		b.WriteString(g.generateContribTip(ext))
	}

	// Source section
	if source := g.generateSourceSection(ext); source != "" {
		b.WriteString(source)
	}

	// Install section
	b.WriteString(g.generateInstallSection(ext))

	return b.String()
}

func (g *ExtensionGenerator) generateFrontmatter(ext *Extension) string {
	categoryTitle := "Unknown"
	if ext.Category.Valid {
		categoryTitle = strings.Title(ext.Category.String)
	}

	desc := ext.EnDesc.String
	if desc == "" {
		desc = "PostgreSQL Extension"
	}

	// Build the subtitle line with format: [**<pkg>**](**url**) : <zh_desc>
	subtitle := ""
	if ext.URL.String != "" {
		subtitle = fmt.Sprintf("[**%s**](%s) : %s", ext.Pkg, ext.URL.String, ext.EnDesc.String)
	} else {
		subtitle = fmt.Sprintf("**%s** : %s", ext.Pkg, ext.EnDesc.String)
	}

	return fmt.Sprintf(`---
title: "%s"
linkTitle: "%s"
description: "%s"
weight: %d
categories: ["%s"]
width: full
---

%s

`, ext.Name, ext.Name, desc, ext.ID, categoryTitle, subtitle)
}

func (g *ExtensionGenerator) generateOverview(ext *Extension) string {
	category := ""
	if ext.Category.Valid {
		category = CategoryShortcode(ext.Category.String)
	}

	license := ""
	if ext.License.Valid {
		license = LicenseShortcode(ext.License.String)
	}

	lang := ""
	if ext.Lang.Valid {
		lang = LanguageShortcode(ext.Lang.String)
	}

	url := ""
	if ext.URL.Valid {
		url = ext.URL.String
	}
	extBadge := Badge(ext.Name, "", "", url, "")

	pkgShortcode := ExtShortcode(ext.Name, ext.Pkg)

	version := "-"
	if ext.Version.Valid {
		version = ext.Version.String
	}

	var b strings.Builder
	b.WriteString("\n## Overview\n\n")
	b.WriteString("|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |\n")
	b.WriteString("|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|\n")
	b.WriteString(fmt.Sprintf("| **%d** | %s | %s | `%s` | %s | %s | %s |\n\n",
		ext.ID, extBadge, pkgShortcode, version, category, license, lang))

	return b.String()
}

func (g *ExtensionGenerator) generateAttributes(ext *Extension) string {
	attrs := ext.GetAttributeBadge()

	hasBin := Badge("No", "blue", "", "", "")
	if ext.HasBin {
		hasBin = Badge("Yes", "green", "", "", "")
	}

	hasLib := Badge("No", "blue", "", "", "")
	if ext.HasLib {
		hasLib = Badge("Yes", "green", "", "", "")
	}

	needLoad := Badge("No", "blue", "", "", "")
	if ext.NeedLoad {
		needLoad = Badge("Yes", "orange", "", "", "")
	}

	needDDL := Badge("No", "orange", "", "", "")
	if ext.NeedDDL {
		needDDL = Badge("Yes", "green", "", "", "")
	}

	relocatable := Badge("no", "orange", "", "", "")
	if ext.Relocatable.Valid && ext.Relocatable.Bool {
		relocatable = Badge("yes", "green", "", "", "")
	}

	trusted := Badge("no", "orange", "", "", "")
	if ext.Trusted.Valid && ext.Trusted.Bool {
		trusted = Badge("yes", "green", "", "", "")
	}

	// Skip first two characters (c and l flags)
	attrBadge := Badge("--"+attrs[2:], "blue", "", "", "")

	return fmt.Sprintf(`
|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| %s | %s | %s | %s | %s | %s | %s |

`, attrBadge, hasBin, hasLib, needLoad, needDDL, relocatable, trusted)
}

func (g *ExtensionGenerator) generateRelationships(ext *Extension, siblings []*Extension) string {
	if len(ext.Schemas) == 0 && len(ext.Requires) == 0 && len(ext.RequireBy) == 0 && len(ext.SeeAlso) == 0 && len(siblings) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("\n| **Relationships** |   |\n")
	b.WriteString("|:-----------------:|:----|\n")

	if len(ext.Schemas) > 0 {
		schemas := make([]string, len(ext.Schemas))
		for i, s := range ext.Schemas {
			schemas[i] = fmt.Sprintf("`%s`", s)
		}
		b.WriteString(fmt.Sprintf("|    **Schemas**    | %s |\n", strings.Join(schemas, " ")))
	}

	if len(ext.Requires) > 0 {
		requires := make([]string, len(ext.Requires))
		for i, e := range ext.Requires {
			requires[i] = ExtShortcode(e, "")
		}
		b.WriteString(fmt.Sprintf("|   **Requires**    | %s |\n", strings.Join(requires, " ")))
	}

	if len(ext.RequireBy) > 0 {
		requireBy := make([]string, len(ext.RequireBy))
		for i, e := range ext.RequireBy {
			requireBy[i] = ExtShortcode(e, "")
		}
		b.WriteString(fmt.Sprintf("|    **Need By**    | %s |\n", strings.Join(requireBy, " ")))
	}

	if len(ext.SeeAlso) > 0 {
		seeAlso := make([]string, len(ext.SeeAlso))
		for i, e := range ext.SeeAlso {
			seeAlso[i] = ExtShortcode(e, "")
		}
		b.WriteString(fmt.Sprintf("|   **See Also**    | %s |\n", strings.Join(seeAlso, " ")))
	}

	if len(siblings) > 0 {
		siblingLinks := make([]string, len(siblings))
		for i, sib := range siblings {
			siblingLinks[i] = ExtShortcode(sib.Name, "")
		}
		b.WriteString(fmt.Sprintf("|    **Siblings**   | %s |\n", strings.Join(siblingLinks, " ")))
	}

	b.WriteString("\n")
	return b.String()
}

func (g *ExtensionGenerator) generateContribTip(ext *Extension) string {
	if !ext.Contrib {
		return ""
	}
	return "> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel\n\n"
}

func (g *ExtensionGenerator) generateCommentNote(ext *Extension) string {
	if !ext.Comment.Valid || ext.Comment.String == "" {
		return ""
	}
	return fmt.Sprintf("> [!Note] %s\n\n", ext.Comment.String)
}

func (g *ExtensionGenerator) generatePackagesTable(ext *Extension, packages []*PkgInfo) string {
	// Special handling for contrib extensions
	if ext.Contrib {
		return g.generateContribPackagesTable(ext)
	}

	var rows []string

	// SRC row - add source package row first
	if ext.Repo.Valid || ext.Version.Valid || len(ext.PgVer) > 0 {
		srcPgVersions := ParsePGVersions(ext.PgVer)
		srcRow := g.buildPackageSummaryRow("EXT", ext.Repo.String, ext.Version.String,
			ext.Pkg, srcPgVersions, ext.Requires, ext, true) // true indicates this is a EXT row
		if srcRow != "" {
			rows = append(rows, srcRow)
		}
	}

	// RPM row
	if ext.RpmRepo.Valid || ext.RpmVer.Valid || ext.RpmPkg.Valid || len(ext.RpmPg) > 0 {
		rpmPgVersions := ParsePGVersions(ext.RpmPg)
		rpmRow := g.buildPackageSummaryRow("RPM", ext.RpmRepo.String, ext.RpmVer.String,
			ext.RpmPkg.String, rpmPgVersions, ext.RpmDeps, ext, false)
		if rpmRow != "" {
			rows = append(rows, rpmRow)
		}
	}

	// DEB row
	if ext.DebRepo.Valid || ext.DebVer.Valid || ext.DebPkg.Valid || len(ext.DebPg) > 0 {
		debPgVersions := ParsePGVersions(ext.DebPg)
		debRow := g.buildPackageSummaryRow("DEB", ext.DebRepo.String, ext.DebVer.String,
			ext.DebPkg.String, debPgVersions, ext.DebDeps, ext, false)
		if debRow != "" {
			rows = append(rows, debRow)
		}
	}

	if len(rows) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("\n## Packages\n\n")
	b.WriteString("| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |\n")
	b.WriteString("|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|\n")

	for _, row := range rows {
		b.WriteString(row)
		b.WriteString("\n")
	}
	b.WriteString("\n")

	return b.String()
}

func (g *ExtensionGenerator) buildPackageSummaryRow(label, repo, version, pattern string,
	supportedPg []int, deps []string, ext *Extension, isSrcRow bool) string {

	if repo == "" && version == "" && pattern == "" && len(deps) == 0 && len(supportedPg) == 0 {
		return ""
	}

	// Generate repo badge with correct link
	repoLabel := repo
	if repoLabel == "" {
		repoLabel = "N/A"
	}

	// Determine the repo link based on repo name
	repoLink := "/repo/pgsql" // Default to PIGSTY
	if strings.ToUpper(repo) == "PGDG" {
		repoLink = "/repo/pgdg"
	}
	repoBadge := Badge(repoLabel, "", "", repoLink, "")

	versionDisplay := "-"
	if version != "" {
		versionDisplay = fmt.Sprintf("`%s`", version)
	}

	patternDisplay := "-"
	if pattern != "" {
		patternDisplay = fmt.Sprintf("`%s`", pattern)
	}

	// Build PG badges using bg shortcode
	supported := make(map[int]bool)
	for _, pg := range supportedPg {
		supported[pg] = true
	}

	var pgBadges []string
	for _, pg := range g.Cache.PGVersions {
		text := fmt.Sprintf("%d", pg)
		color := "red"

		// For SRC row, don't include alt text; for RPM/DEB rows, include package name pattern
		alt := ""
		if !isSrcRow && pattern != "" {
			alt = strings.ReplaceAll(pattern, "$v", fmt.Sprintf("%d", pg))
		}

		if supported[pg] {
			color = "green"
		}

		pgBadges = append(pgBadges, BgShortcode(text, alt, color))
	}
	pgCell := strings.Join(pgBadges, " ")

	depsStr := "-"
	if len(deps) > 0 {
		depsList := make([]string, len(deps))
		for i, dep := range deps {
			depsList[i] = fmt.Sprintf("`%s`", dep)
		}
		depsStr = strings.Join(depsList, ", ")
	}

	return fmt.Sprintf("| **%s** | %s | %s | %s | %s | %s |",
		label, repoBadge, versionDisplay, pgCell, patternDisplay, depsStr)
}

func (g *ExtensionGenerator) generateContribPackagesTable(ext *Extension) string {
	// Get available PG versions from pg_ver field
	availablePgs := make(map[int]bool)
	for _, pgStr := range ext.PgVer {
		if pgVersions := ParsePGVersions([]string{pgStr}); len(pgVersions) > 0 {
			for _, pg := range pgVersions {
				availablePgs[pg] = true
			}
		}
	}

	var b strings.Builder
	b.WriteString("\n## Packages\n\n")

	// Build table header
	b.WriteString("|")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf(" **PG%d** |", pg))
	}
	b.WriteString("\n")

	// Separator
	b.WriteString("|")
	for range g.Cache.PGVersions {
		b.WriteString(":--------:|")
	}
	b.WriteString("\n")

	// Availability row
	b.WriteString("|")
	version := "-"
	if ext.Version.Valid {
		version = ext.Version.String
	}

	for _, pg := range g.Cache.PGVersions {
		if availablePgs[pg] {
			alt := fmt.Sprintf("PostgreSQL %d: version %s", pg, version)
			b.WriteString(fmt.Sprintf(" %s |", BgShortcode(version, alt, "green")))
		} else {
			alt := fmt.Sprintf("PostgreSQL %d: not available", pg)
			b.WriteString(fmt.Sprintf(" %s |", BgShortcode("N/A", alt, "red")))
		}
	}
	b.WriteString("\n\n")

	return b.String()
}

// getBadgeColor determines the badge color based on state and org
func getBadgeColor(state, org string) string {
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

func (g *ExtensionGenerator) generateAvailabilityMatrix(packages []*PkgInfo, binaries []*Binary) string {
	// Build package lookup map from pgext.pkg data
	pkgMap := make(map[string]*PkgInfo)
	for _, pkg := range packages {
		key := fmt.Sprintf("%d-%s", pkg.PG, pkg.OS)
		pkgMap[key] = pkg
	}

	var b strings.Builder
	b.WriteString("\n| **Linux** / **PG** |")
	for _, pg := range g.Cache.PGVersions {
		b.WriteString(fmt.Sprintf("                  **PG%d**                   |", pg))
	}
	b.WriteString("\n")

	b.WriteString("|:------------------:|")
	for range g.Cache.PGVersions {
		b.WriteString(":-------------------------------------------:|")
	}
	b.WriteString("\n")

	// Generate rows for each OS (already sorted by os_major, os_arch DESC in cache)
	for _, os := range g.Cache.OSVersions {
		b.WriteString(fmt.Sprintf("| %s |", OsShortcode(os.OS)))

		for _, pg := range g.Cache.PGVersions {
			key := fmt.Sprintf("%d-%s", pg, os.OS)

			if pkg, ok := pkgMap[key]; ok {
				// Get package info from pgext.pkg
				name := ""
				if pkg.Name.Valid {
					name = pkg.Name.String
				} else {
					name = pkg.Pkg // fallback to package name
				}

				version := ""
				if pkg.Version.Valid {
					version = pkg.Version.String
				}

				org := ""
				if pkg.Org.Valid {
					org = strings.ToUpper(pkg.Org.String)
				}

				state := "MISS" // default state
				if pkg.State.Valid {
					state = pkg.State.String
				}

				count := "0"
				if pkg.Count.Valid {
					count = fmt.Sprintf("%d", pkg.Count.Int64)
				}

				// Build badge text
				text := "MISS"
				if state != "MISS" && org != "" && version != "" {
					text = fmt.Sprintf("%s %s", org, version)
				}

				// Build alt text: "package_name : STATE count"
				alt := fmt.Sprintf("%s : %s %s", name, state, count)

				// Get color based on state and org
				// If text shows "MISS", always use red color
				var color string
				if text == "MISS" {
					color = "red"
				} else {
					color = getBadgeColor(state, org)
				}

				// Generate bg shortcode
				cellContent := BgShortcode(text, alt, color)

				// Adjust spacing based on state
				if state == "AVAIL" {
					b.WriteString(fmt.Sprintf(" %s |", cellContent))
				} else if state == "HIDE" {
					b.WriteString(fmt.Sprintf("  %s   |", cellContent))
				} else {
					b.WriteString(fmt.Sprintf("      %s      |", cellContent))
				}
			} else {
				// Not available - MISS state
				// Try to infer package name from other available packages
				pkgName := "N/A"

				// Look for any package with the same extension but different PG version
				for _, p := range packages {
					if p.PG == pg {
						// Found a package for this PG version (shouldn't happen in this else block)
						pkgName = p.Name.String
						if pkgName == "" {
							pkgName = p.Pkg
						}
						break
					}
				}

				// If still not found, try to construct from first available package
				if pkgName == "N/A" && len(packages) > 0 {
					basePkg := packages[0]
					if basePkg.Name.Valid && basePkg.Name.String != "" {
						// Replace version number in package name
						pkgName = strings.ReplaceAll(basePkg.Name.String, fmt.Sprintf("_%d", basePkg.PG), fmt.Sprintf("_%d", pg))
						pkgName = strings.ReplaceAll(pkgName, fmt.Sprintf("-%d-", basePkg.PG), fmt.Sprintf("-%d-", pg))
					} else {
						pkgName = strings.ReplaceAll(basePkg.Pkg, fmt.Sprintf("_%d", basePkg.PG), fmt.Sprintf("_%d", pg))
					}
				}

				text := "MISS"
				alt := fmt.Sprintf("%s : MISS 0", pkgName)
				cellContent := BgShortcode(text, alt, "red")
				b.WriteString(fmt.Sprintf("    %s     |", cellContent))
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	return b.String()
}

func (g *ExtensionGenerator) generatePackageDetailsTabs(extName string, packages []*PkgInfo, binaries []*Binary) string {
	// Group binaries by PG version (binaries are already properly sorted from SQL)
	pgGroups := make(map[int][]*Binary)
	for _, binary := range binaries {
		pgGroups[binary.PG] = append(pgGroups[binary.PG], binary)
	}

	if len(pgGroups) == 0 {
		return ""
	}

	// Sort PG versions in descending order
	var pgVersions []int
	for pg := range pgGroups {
		pgVersions = append(pgVersions, pg)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pgVersions)))

	// Generate tab items
	tabItems := make([]string, len(pgVersions))
	for i, pg := range pgVersions {
		tabItems[i] = fmt.Sprintf("PG%d", pg)
	}

	var b strings.Builder
	b.WriteString("\n")

	// Start tabs
	tabContent := ""
	for _, pg := range pgVersions {
		pkgList := pgGroups[pg]

		tab := "\n{{< tab >}}\n\n"
		if len(pkgList) > 0 {
			tab += "| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |\n"
			tab += "|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|\n"

			// Packages are already sorted by os_major ASC, version DESC, file DESC from SQL
			for _, pkg := range pkgList {
				sizeStr := FormatSize(pkg.Size)
				org := pkg.Org.String
				if org == "" {
					org = "N/A"
				}

				// URL is already formatted in the SQL query
				url := pkg.Href
				fileName := filepath.Base(url)

				tab += fmt.Sprintf("| `%s` | `%s` | [%s](/os/%s) | %s | %s | [%s](%s) |\n",
					pkg.Name, pkg.Version, pkg.OS, pkg.OS, org, sizeStr, fileName, url)
			}
		} else {
			tab += "*No packages available for this PostgreSQL version.*\n"
		}
		tab += "\n{{< /tab >}}"
		tabContent += tab
	}

	b.WriteString(TabsShortcode(tabItems, tabContent))
	b.WriteString("\n")

	return b.String()
}

func (g *ExtensionGenerator) generateSourceSection(ext *Extension) string {
	// Skip for contrib extensions
	if ext.Contrib {
		return ""
	}

	var cards []string

	if ext.URL.Valid && ext.URL.String != "" {
		url := ext.URL.String
		icon := "link"
		subtitle := strings.TrimPrefix(url, "https://")
		subtitle = strings.TrimPrefix(subtitle, "http://")

		if strings.Contains(url, "github.com") {
			icon = "github"
		}

		cards = append(cards, CardShortcode("Repository", subtitle, url, icon))
	}

	if ext.Source.Valid && ext.Source.String != "" {
		cards = append(cards, CardShortcode("Source Tarball", ext.Source.String, "/list", "clipboard-list"))
	}

	if len(cards) == 0 {
		return ""
	}

	section := "\n## Source\n\n"
	section += CardsShortcode(3, strings.Join(cards, "\n"))
	section += "\n\n"

	// Check extra.rpm and extra.deb to determine if build command should be shown
	var buildRPM, buildDEB bool
	if ext.Extra != nil {
		if rpm, ok := ext.Extra["rpm"]; ok {
			if rpmBool, ok := rpm.(bool); ok && rpmBool {
				buildRPM = true
			}
		}
		if deb, ok := ext.Extra["deb"]; ok {
			if debBool, ok := deb.(bool); ok && debBool {
				buildDEB = true
			}
		}
	}

	// Only show build command if either rpm or deb is true in extra
	if ext.Source.Valid && ext.Source.String != "" && (buildRPM || buildDEB) {
		var buildHint string
		if buildRPM && buildDEB {
			buildHint = "build rpm/deb"
		} else if buildRPM {
			buildHint = "build rpm"
		} else {
			buildHint = "build deb"
		}
		commands := fmt.Sprintf("pig build pkg %s;\t\t# %s", ext.Pkg, buildHint)
		section += "\n" + TripleQuoteBash(commands) + "\n\n"
	}

	return section
}

func (g *ExtensionGenerator) generateInstallSection(ext *Extension) string {

	// Build the "Install" section with better formatting
	var b strings.Builder
	b.WriteString("\n## Install\n\n")

	loadCreate := g.generateLoadCreate(ext)
	// Special handling for contrib extensions
	if ext.Contrib {
		return b.String() + loadCreate
	}

	// Parse PG versions
	pgVersions := ParsePGVersions(ext.PgVer)
	sort.Sort(sort.Reverse(sort.IntSlice(pgVersions)))

	// Generate install commands
	var installCmds []string
	for _, pg := range pgVersions {
		installCmds = append(installCmds, fmt.Sprintf("pig install %s -v %d;   # install for PG %d", ext.Name, pg, pg))
	}

	// Build the Repo Section based on extension repo
	if ext.Repo.Valid && ext.Repo.String == "PGDG" {
		b.WriteString("Make sure [**PGDG**](/repo/pgdg) repo available:\n\n")
		b.WriteString(TripleQuoteBash("pig repo add pgdg -u    # add pgdg repo and update cache"))
	} else {
		b.WriteString("Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:\n\n")
		b.WriteString(TripleQuoteBash("pig repo add pgsql -u   # add both repo and update cache"))
	}
	b.WriteString("\n\n")

	// Build install commands
	b.WriteString("[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):\n\n")
	installScript := fmt.Sprintf("pig install %s;\t\t# install via package name, for the active PG version\n", ext.Pkg)
	if ext.Name != ext.Pkg {
		installScript = installScript + fmt.Sprintf("pig install %s;\t\t# install by extension name, for the current active PG version\n", ext.Name)
	}
	if len(installCmds) > 0 {
		installScript += "\n" + strings.Join(installCmds, "\n") + "\n"
	}
	b.WriteString(TripleQuoteBash(installScript))
	b.WriteString("\n\n")

	// load + create
	b.WriteString(loadCreate)
	return b.String()
}

func (g *ExtensionGenerator) generateLoadCreate(ext *Extension) string {
	var b strings.Builder
	// Add shared_preload_libraries section if needed
	if ext.NeedLoad {
		b.WriteString("\n[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):\n\n")

		// Build the list of libraries to load
		libSet := make(map[string]bool)
		var libsToLoad []string

		// Check if extension has dependencies that also need loading
		if len(ext.Requires) > 0 {
			for _, depName := range ext.Requires {
				if depExt, exists := g.Cache.ExtMap[depName]; exists && depExt.NeedLoad {
					// Get library name(s) for the dependency
					depLibName := depExt.GetLibName()
					// Handle case where lib name might contain multiple libraries separated by comma
					for _, lib := range strings.Split(depLibName, ",") {
						lib = strings.TrimSpace(lib)
						if lib != "" && !libSet[lib] {
							libSet[lib] = true
							libsToLoad = append(libsToLoad, lib)
						}
					}
				}
			}
		}

		// Add current extension's library name(s)
		extLibName := ext.GetLibName()
		for _, lib := range strings.Split(extLibName, ",") {
			lib = strings.TrimSpace(lib)
			if lib != "" && !libSet[lib] {
				libSet[lib] = true
				libsToLoad = append(libsToLoad, lib)
			}
		}

		// Generate the shared_preload_libraries config
		libConfig := strings.Join(libsToLoad, ", ")
		b.WriteString(TripleQuoteINI(fmt.Sprintf("shared_preload_libraries = '%s';", libConfig)))
		b.WriteString("\n\n")
	}

	// Add CREATE section
	if ext.NeedDDL {
		b.WriteString("\n[**Create**](https://ext.pgsty.com/usage/create) this extension with:\n\n")
		if len(ext.Requires) > 0 {
			b.WriteString(TripleQuoteSQL(fmt.Sprintf("CREATE EXTENSION %s CASCADE; -- requires %s", ext.Name, strings.Join(ext.Requires, ", "))))
		} else {
			b.WriteString(TripleQuoteSQL(fmt.Sprintf("CREATE EXTENSION %s;", ext.Name)))
		}

	} else {
		b.WriteString("\nThis extension does not need `CREATE EXTENSION` DDL command\n\n")
	}
	b.WriteString("\n")

	return b.String()
}
