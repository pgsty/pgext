package cli

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

// PigstyConfigGenerator handles the generation of Pigsty configuration files
type PigstyConfigGenerator struct {
	db      *sql.DB
	osName  string
	osCode  string // el9, d12, u24, etc.
	arch    string // x86_64, aarch64
	verbose bool
}

// PigstyExtInfo holds extension availability information for Pigsty config generation
type PigstyExtInfo struct {
	Name      string
	Pkg       string
	Category  string
	Available bool
	Hidden    bool
	Version   string
	Comment   string
	PGVersions []int // Available PostgreSQL versions
}

// PigstyCategoryExts holds extensions grouped by category for Pigsty config
type PigstyCategoryExts struct {
	Category   string
	Extensions []PigstyExtInfo
}

// PackageAlias represents a package alias mapping
type PackageAlias struct {
	Name     string
	Packages []string // List of actual packages
	Comment  string
}

// GeneratePigstyConfig generates Pigsty configuration for a specific OS
func GeneratePigstyConfig(osName string, outputDir string, dryRun bool, verbose bool) error {
	// Parse OS name (e.g., "el9.x86_64" -> osCode="el9", arch="x86_64")
	parts := strings.Split(osName, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid OS name format: %s (expected format: osCode.arch, e.g., el9.x86_64)", osName)
	}

	// Get standard SQL DB from pgxpool
	sqlDB, err := GetStdDB()
	if err != nil {
		return fmt.Errorf("failed to get SQL DB: %w", err)
	}

	// Use V2 generator
	generator := NewPigstyConfigGenerator(sqlDB, osName, verbose)

	// Generate configuration content
	content, err := generator.GenerateContent()
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	// Determine output file path
	outputFile := filepath.Join(outputDir, osName+".yml")

	if dryRun {
		if verbose {
			logrus.Infof("Dry run mode - would write to: %s", outputFile)
			fmt.Println(content)
		}
		return nil
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if verbose {
		logrus.Infof("Written configuration to: %s", outputFile)
	}

	return nil
}

// GenerateContent generates the YAML configuration content
func (g *PigstyConfigGenerator) GenerateContent() (string, error) {
	// Load extension availability data
	extensions, err := g.loadExtensions()
	if err != nil {
		return "", fmt.Errorf("failed to load extensions: %w", err)
	}

	// Group extensions by category
	categoryMap := g.groupByCategory(extensions)

	// Generate package mappings for each PG version
	pgVersionMaps := make(map[int][]PackageAlias)
	for _, pgVer := range []int{18, 17, 16, 15, 14, 13} {
		pgVersionMaps[pgVer] = g.generatePGPackageMappings(pgVer, categoryMap)
	}

	// Prepare template data
	data := struct {
		OSName       string
		OSCode       string
		Arch         string
		IsRPM        bool
		IsDEB        bool
		PGVersions   []int
		PGMappings   map[int][]PackageAlias
		Categories   []PigstyCategoryExts
	}{
		OSName:     g.osName,
		OSCode:     g.osCode,
		Arch:       g.arch,
		IsRPM:      isRPMDistro(g.osCode),
		IsDEB:      isDEBDistro(g.osCode),
		PGVersions: []int{18, 17, 16, 15, 14, 13},
		PGMappings: pgVersionMaps,
		Categories: g.orderCategories(categoryMap),
	}

	// Select appropriate template
	var tmplContent string
	if isRPMDistro(g.osCode) {
		tmplContent = rpmConfigTemplate
	} else {
		tmplContent = debConfigTemplate
	}

	// Parse and execute template
	tmpl, err := template.New("config").Funcs(template.FuncMap{
		"join": strings.Join,
		"formatPkgLine": func(alias PackageAlias) string {
			return g.formatPackageLine(alias)
		},
	}).Parse(tmplContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// loadExtensions loads extension availability from database
func (g *PigstyConfigGenerator) loadExtensions() ([]PigstyExtInfo, error) {
	query := `
		SELECT
			e.name,
			e.pkg,
			e.category,
			p.state,
			p.version,
			p.hide,
			p.pg,
			e.comment
		FROM pgext.extension e
		LEFT JOIN pgext.pkg p ON e.name = p.ext AND p.os = $1
		WHERE e.lead = true
		ORDER BY e.category, e.name, p.pg
	`

	rows, err := g.db.Query(query, g.osName)
	if err != nil {
		return nil, fmt.Errorf("failed to query extensions: %w", err)
	}
	defer rows.Close()

	// Build extension map to aggregate PG versions
	extMap := make(map[string]*PigstyExtInfo)

	for rows.Next() {
		var (
			name     string
			pkg      string
			category string
			state    sql.NullString
			version  sql.NullString
			hide     sql.NullBool
			pgVer    sql.NullInt32
			comment  sql.NullString
		)

		err := rows.Scan(&name, &pkg, &category, &state, &version, &hide, &pgVer, &comment)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Get or create extension info
		if _, exists := extMap[name]; !exists {
			extMap[name] = &PigstyExtInfo{
				Name:       name,
				Pkg:        pkg,
				Category:   category,
				Available:  false,
				Hidden:     false,
				PGVersions: []int{},
			}
			if comment.Valid {
				extMap[name].Comment = comment.String
			}
		}

		// Update availability based on state
		if state.Valid && state.String == "AVAIL" {
			extMap[name].Available = true
			if pgVer.Valid {
				extMap[name].PGVersions = append(extMap[name].PGVersions, int(pgVer.Int32))
			}
			if version.Valid && extMap[name].Version == "" {
				extMap[name].Version = version.String
			}
		}

		// Update hidden status
		if hide.Valid && hide.Bool {
			extMap[name].Hidden = true
		}
	}

	// Convert map to slice
	var extensions []PigstyExtInfo
	for _, ext := range extMap {
		extensions = append(extensions, *ext)
	}

	// Sort by category and name
	sort.Slice(extensions, func(i, j int) bool {
		if extensions[i].Category != extensions[j].Category {
			return extensions[i].Category < extensions[j].Category
		}
		return extensions[i].Name < extensions[j].Name
	})

	return extensions, nil
}

// groupByCategory groups extensions by their category
func (g *PigstyConfigGenerator) groupByCategory(extensions []PigstyExtInfo) map[string][]PigstyExtInfo {
	categoryMap := make(map[string][]PigstyExtInfo)
	for _, ext := range extensions {
		categoryMap[ext.Category] = append(categoryMap[ext.Category], ext)
	}
	return categoryMap
}

// orderCategories returns categories in the standard order
func (g *PigstyConfigGenerator) orderCategories(categoryMap map[string][]PigstyExtInfo) []PigstyCategoryExts {
	// Standard category order
	categoryOrder := []string{
		"TIME", "GIS", "RAG", "FTS", "OLAP", "FEAT",
		"LANG", "TYPE", "UTIL", "FUNC", "ADMIN", "STAT",
		"SEC", "FDW", "SIM", "ETL",
	}

	var result []PigstyCategoryExts
	for _, cat := range categoryOrder {
		if exts, ok := categoryMap[cat]; ok {
			result = append(result, PigstyCategoryExts{
				Category:   cat,
				Extensions: exts,
			})
		}
	}
	return result
}

// generatePGPackageMappings generates package mappings for a specific PG version
func (g *PigstyConfigGenerator) generatePGPackageMappings(pgVer int, categoryMap map[string][]PigstyExtInfo) []PackageAlias {
	var mappings []PackageAlias

	// Generate category-based package groups
	categoryOrder := []string{
		"time", "gis", "rag", "fts", "olap", "feat",
		"lang", "type", "util", "func", "admin", "stat",
		"sec", "fdw", "sim", "etl",
	}

	for _, catName := range categoryOrder {
		catUpper := strings.ToUpper(catName)
		if exts, ok := categoryMap[catUpper]; ok {
			var packages []string

			for _, ext := range exts {
				// Check if extension is available for this PG version
				isAvailableForPG := false
				for _, v := range ext.PGVersions {
					if v == pgVer {
						isAvailableForPG = true
						break
					}
				}

				if isAvailableForPG {
					pkgName := g.getPackageName(ext, pgVer)
					// Add asterisk for optional package
					if !ext.Hidden && !g.shouldHideExtension(ext.Name, pgVer) {
						packages = append(packages, pkgName+"*")
					}
				}
			}

			// Also add hidden/commented packages
			for _, ext := range exts {
				isAvailableForPG := false
				for _, v := range ext.PGVersions {
					if v == pgVer {
						isAvailableForPG = true
						break
					}
				}

				if isAvailableForPG && (ext.Hidden || g.shouldHideExtension(ext.Name, pgVer)) {
					pkgName := g.getPackageName(ext, pgVer)
					packages = append(packages, "#"+pkgName)
				}
			}

			if len(packages) > 0 {
				alias := PackageAlias{
					Name:     fmt.Sprintf("pg%d-%s", pgVer, catName),
					Packages: packages,
				}
				mappings = append(mappings, alias)
			}
		}
	}

	return mappings
}

// getPackageName returns the package name for an extension
func (g *PigstyConfigGenerator) getPackageName(ext PigstyExtInfo, pgVer int) string {
	// Handle special cases first
	switch ext.Name {
	case "pgaudit":
		return g.getPGAuditPackageName(pgVer)
	case "hunspell_cs_cz", "hunspell_de_de", "hunspell_en_us", "hunspell_fr",
	     "hunspell_ne_np", "hunspell_nl_nl", "hunspell_nn_no", "hunspell_ru_ru", "hunspell_ru_ru_aot":
		// All hunspell variants get their specific package name
		if isRPMDistro(g.osCode) {
			return fmt.Sprintf("%s_%d", ext.Name, pgVer)
		} else {
			return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(ext.Name, "_", "-"))
		}
	case "babelfishpg_common":
		// Merge babelfish into wiltondb
		if g.osCode == "el7" || g.osCode == "el8" || g.osCode == "el9" ||
		   g.osCode == "u20" || g.osCode == "u22" || g.osCode == "u24" {
			return "wiltondb"
		}
	}

	// Query-based package naming from database
	var pkgName string
	query := `
		SELECT name
		FROM pgext.pkg
		WHERE os = $1 AND pg = $2 AND ext = $3
		LIMIT 1
	`
	err := g.db.QueryRow(query, g.osName, pgVer, ext.Name).Scan(&pkgName)
	if err == nil && pkgName != "" {
		return pkgName
	}

	// Fallback to standard package naming
	if isRPMDistro(g.osCode) {
		// For RPM, typically: extension_name_$v
		if ext.Pkg != "" && ext.Pkg != ext.Name {
			// Use the pkg field from extension table
			return strings.ReplaceAll(ext.Pkg, "$v", fmt.Sprintf("%d", pgVer))
		}
		return fmt.Sprintf("%s_%d", ext.Name, pgVer)
	} else {
		// For DEB, typically: postgresql-$v-extension-name
		if ext.Pkg != "" && ext.Pkg != ext.Name {
			return strings.ReplaceAll(ext.Pkg, "$v", fmt.Sprintf("%d", pgVer))
		}
		return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(ext.Name, "_", "-"))
	}
}

// getPGAuditPackageName returns the correct pgaudit package name based on PG version
func (g *PigstyConfigGenerator) getPGAuditPackageName(pgVer int) string {
	switch pgVer {
	case 18, 17, 16:
		return fmt.Sprintf("pgaudit_%d", pgVer)
	case 15:
		return "pgaudit17_15"
	case 14:
		return "pgaudit16_14"
	case 13:
		return "pgaudit15_13"
	default:
		return fmt.Sprintf("pgaudit_%d", pgVer)
	}
}

// shouldHideExtension determines if an extension should be hidden
func (g *PigstyConfigGenerator) shouldHideExtension(extName string, pgVer int) bool {
	// List of extensions to hide
	hideList := []string{
		"hydra", "duckdb_fdw", "pg_timeseries", "pgpool", "plr",
		"pgagent", "dbt2", "pgtap", "faker", "repmgr", "slony",
		"oracle_fdw", "pg_strom", "db2_fdw", "orioledb",
	}

	for _, hidden := range hideList {
		if extName == hidden {
			return true
		}
	}

	// Special cases based on OS and architecture
	if extName == "rdkit" && g.osCode != "u24" {
		return true
	}
	if extName == "jdbc_fdw" && (g.osCode == "el8" || g.osCode == "el9") && g.arch == "aarch64" {
		return true
	}
	if extName == "pllua" && (g.osCode == "el8" || g.osCode == "el9") && g.arch == "aarch64" && pgVer < 16 {
		return true
	}

	return false
}

// formatPackageLine formats a package line for the YAML output
func (g *PigstyConfigGenerator) formatPackageLine(alias PackageAlias) string {
	// Packages already have asterisks or # prefix added
	return fmt.Sprintf(`  %-24s "%s"`, alias.Name+":", strings.Join(alias.Packages, " "))
}

// Helper functions
func isRPMDistro(osCode string) bool {
	return strings.HasPrefix(osCode, "el") // el7, el8, el9, el10
}

func isDEBDistro(osCode string) bool {
	return strings.HasPrefix(osCode, "d") || strings.HasPrefix(osCode, "u") // d11, d12, d13, u20, u22, u24
}