package cli

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"pgext/cli/repo"
	"sort"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

func el9ARMPatroniPackages(extra ...string) string {
	packages := []string{
		"patroni.noarch",
		"patroni-etcd.noarch",
	}
	packages = append(packages, extra...)
	return strings.Join(packages, " ")
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
	defer sqlDB.Close()

	// Create generator
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

// PigstyConfigGenerator handles the generation of Pigsty configuration files
type PigstyConfigGenerator struct {
	db        *sql.DB
	osName    string
	osCode    string // el9, d12, u24, etc.
	arch      string // x86_64, aarch64
	verbose   bool
	constants *ConfigConstants
}

func (g *PigstyConfigGenerator) isEL9ARM() bool {
	return g.osCode == "el9" && g.arch == "aarch64"
}

func (g *PigstyConfigGenerator) tunedProfileDir() string {
	switch g.osCode {
	case "el10", "d13", "u26":
		return "/etc/tuned/profiles"
	default:
		return "/etc/tuned"
	}
}

// NewPigstyConfigGenerator creates a new generator
func NewPigstyConfigGenerator(db *sql.DB, osName string, verbose bool) *PigstyConfigGenerator {
	parts := strings.Split(osName, ".")
	osCode := parts[0]
	arch := ""
	if len(parts) > 1 {
		arch = parts[1]
	}

	return &PigstyConfigGenerator{
		db:        db,
		osName:    osName,
		osCode:    osCode,
		arch:      arch,
		verbose:   verbose,
		constants: GetConfigConstants(),
	}
}

// ExtensionData holds processed extension information
type ExtensionData struct {
	ID              int // Extension ID for sorting
	Name            string
	Alias           string
	Package         string
	Category        string
	CategoryID      int            // Category ID for sorting
	Available       map[int]bool   // PG version -> availability
	VersionPackages map[int]string // PG version -> package name from pgext.pkg
	VersionHidden   map[int]bool   // PG version -> hidden status
	Hidden          bool           // deprecated, kept for compatibility
	ForkBound       bool           // true when tags contains fork; render only via manual fork aliases
	Comment         string
	RPMPkg          string // RPM package template from extension.rpm_pkg
	DEBPkg          string // DEB package template from extension.deb_pkg
	RPMPG           []int  // Declared RPM PG versions from extension.rpm_pg
	DEBPG           []int  // Declared DEB PG versions from extension.deb_pg
}

// CategoryPkgList holds visible and hidden packages for a category
type CategoryPkgList struct {
	Visible string // Packages in quotes
	Hidden  string // Hidden packages as comments
}

// ExtensionMapping holds extension to package mapping
type ExtensionMapping struct {
	Name     string
	Package  string
	PGVers   []int
	Comment  string
	Category string
}

// GenerateContent generates the complete configuration content
func (g *PigstyConfigGenerator) GenerateContent() (string, error) {
	// Load extension data from database
	extensions, err := g.loadExtensionData()
	if err != nil {
		return "", fmt.Errorf("failed to load extension data: %w", err)
	}

	// Generate content based on distribution type
	if g.isRPM() {
		return g.generateRPMConfig(extensions)
	}
	return g.generateDEBConfig(extensions)
}

// loadExtensionData loads and processes extension information from database
func (g *PigstyConfigGenerator) loadExtensionData() (map[string]*ExtensionData, error) {
	query := `
		WITH ext_avail AS (
			SELECT
				e.id as ext_id,
				e.name,
				COALESCE(e.pkg, e.name) as alias,
				e.category,
				COALESCE(c.id, 999) as cat_id,
				e.comment,
				e.rpm_pkg,
				e.deb_pkg,
				e.rpm_pg,
				e.deb_pg,
				COALESCE(e.tags @> ARRAY['fork']::text[], false) AS fork_bound,
				p.pg,
				p.name as pkg_name,
				p.state,
				p.hide
			FROM pgext.extension e
			LEFT JOIN pgext.category c ON c.name = e.category
			LEFT JOIN pgext.pkg p ON p.ext = e.name AND p.os = $1
			WHERE e.lead = true AND e.contrib IS NOT TRUE
		)
		SELECT ext_id, name, alias, category, cat_id, comment, rpm_pkg, deb_pkg, rpm_pg, deb_pg, fork_bound, pg, pkg_name, state, hide
		FROM ext_avail
		ORDER BY cat_id, ext_id, pg
	`

	rows, err := g.db.Query(query, g.osName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	extensions := make(map[string]*ExtensionData)

	for rows.Next() {
		var (
			extID     int
			name      string
			alias     string
			category  string
			catID     int
			comment   sql.NullString
			rpmPkg    sql.NullString
			debPkg    sql.NullString
			rpmPG     interface{}
			debPG     interface{}
			forkBound bool
			pg        sql.NullInt32
			pkgName   sql.NullString
			state     sql.NullString
			hide      sql.NullBool
		)

		err := rows.Scan(&extID, &name, &alias, &category, &catID, &comment, &rpmPkg, &debPkg, &rpmPG, &debPG, &forkBound, &pg, &pkgName, &state, &hide)
		if err != nil {
			return nil, err
		}

		// Get or create extension entry
		if _, exists := extensions[name]; !exists {
			extensions[name] = &ExtensionData{
				ID:              extID,
				Name:            name,
				Alias:           alias,
				Category:        category,
				CategoryID:      catID,
				Available:       make(map[int]bool),
				VersionPackages: make(map[int]string),
				VersionHidden:   make(map[int]bool),
				ForkBound:       forkBound,
			}
			if comment.Valid {
				extensions[name].Comment = comment.String
			}
			if rpmPkg.Valid {
				extensions[name].RPMPkg = rpmPkg.String
			}
			if debPkg.Valid {
				extensions[name].DEBPkg = debPkg.String
			}
			extensions[name].RPMPG = ParsePGVersions(parseArrayValue(rpmPG))
			extensions[name].DEBPG = ParsePGVersions(parseArrayValue(debPG))
		}

		// Update availability and package name for specific PG version
		if pg.Valid && state.Valid && state.String == "AVAIL" {
			pgVersion := int(pg.Int32)
			extensions[name].Available[pgVersion] = true

			// Store the version-specific package name if available
			if pkgName.Valid && pkgName.String != "" {
				extensions[name].VersionPackages[pgVersion] = pkgName.String
				// Also store a default package name (use the latest one we see)
				extensions[name].Package = pkgName.String
			}

			// Store version-specific hidden status
			if hide.Valid {
				extensions[name].VersionHidden[pgVersion] = hide.Bool
			}
		}
	}

	return extensions, nil
}

// isRPM returns true if this is an RPM-based distribution
func (g *PigstyConfigGenerator) isRPM() bool {
	return strings.HasPrefix(g.osCode, "el")
}

// generateRPMConfig generates configuration for RPM-based systems
func (g *PigstyConfigGenerator) generateRPMConfig(extensions map[string]*ExtensionData) (string, error) {
	tmpl := template.Must(template.New("rpm").Funcs(g.getFuncMap()).Parse(rpmTemplate))

	data := g.prepareTemplateData(extensions)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// generateDEBConfig generates configuration for DEB-based systems
func (g *PigstyConfigGenerator) generateDEBConfig(extensions map[string]*ExtensionData) (string, error) {
	tmpl := template.Must(template.New("deb").Funcs(g.getFuncMap()).Parse(debTemplate))

	data := g.prepareTemplateData(extensions)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// prepareTemplateData prepares all data for template rendering
func (g *PigstyConfigGenerator) prepareTemplateData(extensions map[string]*ExtensionData) map[string]interface{} {
	catExts := g.generateCategoryPackages(extensions)

	// Get ordered category list
	categories, err := g.loadCategories()
	if err != nil {
		// Fallback to default order if database query fails
		categories = []string{"time", "gis", "rag", "fts", "olap", "feat", "lang", "type", "util", "func", "admin", "stat", "sec", "fdw", "sim", "etl"}
	}

	return map[string]interface{}{
		"OSName":          g.osName,
		"OSCode":          g.osCode,
		"Arch":            g.arch,
		"TunedProfileDir": g.tunedProfileDir(),
		"Constants":       g.constants,
		"Extensions":      extensions,
		"PGVersions":      []int{18, 17, 16, 15, 14},
		"CategoryExts":    catExts,
		"Categories":      categories,
		"ExtMappings":     g.generateExtensionMappings(extensions),
	}
}

// loadCategories loads categories from database in ID order
func (g *PigstyConfigGenerator) loadCategories() ([]string, error) {
	query := `SELECT LOWER(name) FROM pgext.category ORDER BY id`

	rows, err := g.db.Query(query)
	if err != nil {
		if g.verbose {
			fmt.Printf("Failed to load categories from database: %v\n", err)
		}
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var cat string
		if err := rows.Scan(&cat); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	if g.verbose && len(categories) > 0 {
		fmt.Printf("Loaded %d categories from database in ID order: %v\n", len(categories), categories)
	}

	return categories, nil
}

// generateCategoryPackages generates category-based package listings for each PG version
func (g *PigstyConfigGenerator) generateCategoryPackages(extensions map[string]*ExtensionData) map[int]map[string]*CategoryPkgList {
	result := make(map[int]map[string]*CategoryPkgList)

	// Load categories from database in proper ID order
	categories, err := g.loadCategories()
	if err != nil {
		// Fallback to default order if database query fails
		categories = []string{"time", "gis", "rag", "fts", "olap", "feat", "lang", "type", "util", "func", "admin", "stat", "sec", "fdw", "sim", "etl"}
	}
	pgVersions := []int{18, 17, 16, 15, 14}

	// Convert extensions to sorted slice
	sortedExtensions := sortExtensionsByID(extensions)

	for _, pgVer := range pgVersions {
		result[pgVer] = make(map[string]*CategoryPkgList)

		for _, cat := range categories {
			var packages []string
			var hiddenPackages []string

			// Collect all packages for this category and version
			type pkgInfo struct {
				name   string
				id     int
				hidden bool
			}
			var pkgList []pkgInfo

			for _, ext := range sortedExtensions {
				if ext.ForkBound {
					continue
				}
				if strings.ToLower(ext.Category) != cat {
					continue
				}

				// Check if available for this PG version
				if !ext.Available[pgVer] {
					continue
				}

				// Generate package name based on distribution type
				pkgName := g.getPackageNameForCategory(ext, pgVer)

				// Use version-specific hide flag from database configuration
				isHidden := false
				if hidden, exists := ext.VersionHidden[pgVer]; exists {
					isHidden = hidden
				}

				// Don't add asterisk for wildcard matching - the package name already includes it if needed

				pkgList = append(pkgList, pkgInfo{
					name:   pkgName,
					id:     ext.ID,
					hidden: isHidden,
				})
			}

			// Sort the package list by ID
			sort.Slice(pkgList, func(i, j int) bool {
				return pkgList[i].id < pkgList[j].id
			})

			// Separate into visible and hidden packages
			for _, pkg := range pkgList {
				if pkg.hidden {
					hiddenPackages = append(hiddenPackages, pkg.name)
				} else {
					packages = append(packages, pkg.name)
				}
			}

			// Build the final package strings
			if len(packages) > 0 || len(hiddenPackages) > 0 {
				catPkgs := &CategoryPkgList{}

				// Visible packages go in quotes
				if len(packages) > 0 {
					catPkgs.Visible = strings.Join(packages, " ")
				}

				// Hidden packages become YAML comments
				if len(hiddenPackages) > 0 {
					catPkgs.Hidden = strings.Join(hiddenPackages, " ")
				}

				result[pgVer][cat] = catPkgs
			}
		}
	}

	return result
}

// getExtensionMappingVersions returns PG versions for extension alias comments.
// Prefer real AVAIL data from pkg table, and fallback to declared rpm_pg/deb_pg when AVAIL is empty.
func (g *PigstyConfigGenerator) getExtensionMappingVersions(ext *ExtensionData) []int {
	var versions []int
	for pg := range ext.Available {
		if ext.Available[pg] {
			versions = append(versions, pg)
		}
	}
	if len(versions) > 0 {
		return versions
	}

	if g.isRPM() && len(ext.RPMPG) > 0 {
		return append([]int(nil), ext.RPMPG...)
	}
	if !g.isRPM() && len(ext.DEBPG) > 0 {
		return append([]int(nil), ext.DEBPG...)
	}
	return versions
}

// generateExtensionMappings generates extension to package name mappings
func (g *PigstyConfigGenerator) generateExtensionMappings(extensions map[string]*ExtensionData) []ExtensionMapping {
	var mappings []ExtensionMapping

	// Sort extensions by category ID and extension ID
	sortedExts := sortExtensionsByID(extensions)
	seenAliases := make(map[string]struct{})

	// Generate mappings
	for _, ext := range sortedExts {
		// Skip contrib extensions
		if ext.Category == "CONTRIB" {
			continue
		}
		if ext.ForkBound {
			continue
		}
		if _, exists := seenAliases[ext.Alias]; exists {
			continue
		}
		seenAliases[ext.Alias] = struct{}{}

		// Determine package pattern
		var pkgPattern string
		if g.isRPM() {
			pkgPattern = g.getRPMPackagePattern(ext)
		} else {
			pkgPattern = g.getDEBPackagePattern(ext)
		}

		// Determine available versions
		availVersions := g.getExtensionMappingVersions(ext)
		sort.Sort(sort.Reverse(sort.IntSlice(availVersions)))

		mappings = append(mappings, ExtensionMapping{
			Name:     ext.Alias,
			Package:  pkgPattern,
			PGVers:   availVersions,
			Comment:  g.getExtensionComment(ext),
			Category: ext.Category,
		})
	}

	return mappings
}

// getPackageNameForCategory returns the package name for use in category lists
func (g *PigstyConfigGenerator) getPackageNameForCategory(ext *ExtensionData, pgVer int) string {
	// Use rpm_pkg or deb_pkg templates from extension
	if g.isRPM() {
		if ext.RPMPkg != "" {
			// Replace $v with the PG version
			return strings.ReplaceAll(ext.RPMPkg, "$v", fmt.Sprintf("%d", pgVer))
		}
		// Fallback: check version-specific package name from pgext.pkg
		if pkgName, exists := ext.VersionPackages[pgVer]; exists && pkgName != "" {
			return pkgName
		}
		// Final fallback if rpm_pkg not set
		if ext.Alias != "" {
			return fmt.Sprintf("%s_%d", ext.Alias, pgVer)
		}
		return fmt.Sprintf("%s_%d", ext.Name, pgVer)
	} else {
		// DEB format
		if ext.DEBPkg != "" {
			// Replace $v with the PG version
			return strings.ReplaceAll(ext.DEBPkg, "$v", fmt.Sprintf("%d", pgVer))
		}
		// Fallback: check version-specific package name from pgext.pkg
		if pkgName, exists := ext.VersionPackages[pgVer]; exists && pkgName != "" {
			return pkgName
		}
		// Final fallback if deb_pkg not set
		if ext.Alias != "" {
			return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(ext.Alias, "_", "-"))
		}
		return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(ext.Name, "_", "-"))
	}
}

// getRPMPackagePattern returns the RPM package pattern for an extension
func (g *PigstyConfigGenerator) getRPMPackagePattern(ext *ExtensionData) string {
	// Use RPMPkg if available, keeping $v placeholder
	if ext.RPMPkg != "" {
		return ext.RPMPkg
	}
	return ext.Alias + "_$v"
}

// getDEBPackagePattern returns the DEB package pattern for an extension
func (g *PigstyConfigGenerator) getDEBPackagePattern(ext *ExtensionData) string {
	// Use DEBPkg if available, keeping $v placeholder
	if ext.DEBPkg != "" {
		return ext.DEBPkg
	}
	return "postgresql-$v-" + strings.ReplaceAll(ext.Alias, "_", "-")
}

// getExtensionComment returns the comment for an extension
func (g *PigstyConfigGenerator) getExtensionComment(ext *ExtensionData) string {
	comments := map[string]string{
		"hydra":      "exclude from 16 category alias lists (special case)",
		"pgpool":     "exclude due to not used",
		"plr":        "exclude due to heavy deps",
		"pgagent":    "exclude due to not used",
		"dbt2":       "exclude due to not used",
		"pgtap":      "exclude due to broken perl deps",
		"faker":      "exclude due to broken perl deps",
		"repmgr":     "exclude due to not used",
		"slony":      "exclude due to not used",
		"oracle_fdw": "exclude due to heavy oracle deps",
		"pg_strom":   "exclude due to heavy cuda deps",
		"db2_fdw":    "exclude due to heavy db2 deps",
		"orioledb":   "work with orioledb postgres fork",
		"pg_tde":     "work with percona postgres fork",
	}

	if comment, ok := comments[ext.Alias]; ok {
		return comment
	}

	if ext.Comment != "" {
		return ext.Comment
	}

	return ""
}

// getFuncMap returns template functions
func (g *PigstyConfigGenerator) getFuncMap() template.FuncMap {
	return template.FuncMap{
		"join": strings.Join,
		"replaceAll": func(old, new, s string) string {
			// Note: when used with pipe, the piped value comes LAST
			// e.g., {{ $str | replaceAll "old" "new" }} calls replaceAll("old", "new", $str)
			return strings.ReplaceAll(s, old, new)
		},
		"hasPrefix": strings.HasPrefix,
		"trimPrefix": func(prefix, s string) string {
			// Note: template calls with prefix first, but strings.TrimPrefix expects string first
			return strings.TrimPrefix(s, prefix)
		},
		"toUpper": strings.ToUpper,
		"sub": func(a, b int) int {
			return a - b
		},
		"formatPGVersions": func(versions []int) string {
			var result []string
			allVersions := []int{18, 17, 16, 15, 14}
			for _, v := range allVersions {
				found := false
				for _, av := range versions {
					if av == v {
						found = true
						break
					}
				}
				if found {
					result = append(result, fmt.Sprintf("%2d", v))
				} else {
					result = append(result, "  ")
				}
			}
			return strings.Join(result, " ")
		},
		"getBootstrap": func() string {
			// Try OS-specific first
			if pkg, ok := g.constants.DistroAdhocPkg[g.osCode]; ok {
				return pkg
			}
			// Fallback to generic
			if g.isRPM() {
				return g.constants.DistroAdhocPkg["rpm"]
			}
			return g.constants.DistroAdhocPkg["deb"]
		},
		"getJavaRuntime": func() string {
			// el10 and d13 use Java 21, others use Java 17
			if g.isRPM() {
				if g.osCode == "el10" {
					return "java-21-openjdk-src java-21-openjdk-headless"
				}
				return "java-17-openjdk-src java-17-openjdk-headless"
			}
			// DEB systems
			if g.osCode == "d13" {
				return "openjdk-21-jdk"
			}
			return "openjdk-17-jdk"
		},
		"getNodePackage1": func() string {
			return g.constants.RPMCommonPkg[3]
		},
		"getDebNodePackage1": func() string {
			return g.constants.DEBCommonPkg[3]
		},
		"getNodePackage2": func() string {
			if g.isRPM() {
				if g.osCode == "el7" {
					return g.constants.RPMCommonPkg[4] + " iproute"
				}
				return g.constants.RPMCommonPkg[4]
			}
			return g.constants.DEBCommonPkg[4]
		},
		"getNodePackage3": func() string {
			if !g.isRPM() {
				return g.constants.DEBCommonPkg[5]
			}
			// el10 doesn't have flamegraph package
			basePkgs := g.constants.RPMCommonPkg[5] + " netcat ftp perf"
			if g.osCode == "el10" {
				return basePkgs
			}
			return basePkgs + " flamegraph"
		},
		"getPgsqlUtility": func() string {
			if g.isEL9ARM() {
				return el9ARMPatroniPackages("pgbouncer", "pgbackrest", "pgbadger", "pg-timetable", "pgFormatter", "pg_filedump", "pgxnclient", "timescaledb-tools", "timescaledb-event-streamer")
			}
			return g.constants.RPMCommonPkg[6]
		},
		"getUtilPkg": func(key, rpm string) string {
			if g.isEL9ARM() {
				if key == "pgsql-common" {
					return el9ARMPatroniPackages("pgbouncer", "pgbackrest", "pg-exporter", "pgbackrest-exporter", "vip-manager")
				}
				if key == "patroni" {
					return el9ARMPatroniPackages()
				}
			}
			return rpm
		},
	}
}

// sortExtensionsByID converts extension map to slice sorted by ID
func sortExtensionsByID(extensions map[string]*ExtensionData) []*ExtensionData {
	// Convert map to slice
	extSlice := make([]*ExtensionData, 0, len(extensions))
	for _, ext := range extensions {
		extSlice = append(extSlice, ext)
	}

	// Sort by category ID first, then by extension ID
	sort.Slice(extSlice, func(i, j int) bool {
		if extSlice[i].CategoryID != extSlice[j].CategoryID {
			return extSlice[i].CategoryID < extSlice[j].CategoryID
		}
		return extSlice[i].ID < extSlice[j].ID
	})

	return extSlice
}

// ConfigConstants holds all the constant package lists from Python script
type ConfigConstants struct {
	// Package groups
	RPMCommonPkg []string
	DEBCommonPkg []string
	// Legacy exporter package names mapped to their normalized package names.
	InfraAliasMap []PackageMapping

	// Kernel packages
	PGSQLKernelMap []PackageMapping

	// Utility packages
	PGSQLUtilMap []PackageMapping

	// Exotic/Pro modules
	PGSQLExoticMap []PackageMapping

	// OS-specific bootstrap packages
	DistroAdhocPkg map[string]string
}

// PackageMapping holds package name mappings between RPM and DEB
type PackageMapping struct {
	Key string
	RPM string
	DEB string
}

// GetConfigConstants returns all the constant definitions
func GetConfigConstants() *ConfigConstants {
	constants := &ConfigConstants{
		RPMCommonPkg: []string{
			// 0: infra-package
			"nginx dnsmasq etcd haproxy vip-manager node-exporter keepalived-exporter pg-exporter pgbackrest-exporter redis-exporter redis valkey silo mcli openbao kes sow pig",
			// 1: infra-addons
			"grafana grafana-plugins grafana-victoriametrics-ds grafana-victorialogs-ds victoria-metrics victoria-logs victoria-traces vlogscli vmutils vector alertmanager",
			// 2: extra-modules
			"blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx docker-ce docker-compose-plugin duckdb restic juicefs vray grafana-infinity-ds",
			// 3: node-package1
			"bash python3 sudo acl ca-certificates openssl curl wget lz4 zstd unzip bzip2 gzip tar tzdata chrony openssh-server util-linux rsync psmisc logrotate",
			// 4: node-package2
			"pv jq git make patch lsof less ncdu htop iotop socat net-tools telnet ipvsadm tuned numactl nvme-cli sysstat keepalived etcd haproxy vector pig uv",
			// 5: node-package3
			"zlib readline xz glibc-langpack-en cronie openssh-clients node-exporter bind-utils iproute iputils nmap-ncat procps-ng vim-minimal yum audit grubby chkconfig",
			// 6: pgsql-utility
			"patroni patroni-etcd pgbouncer pgbackrest pgbadger pg-timetable pgFormatter pg_filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
		},

		DEBCommonPkg: []string{
			// 0: infra-package
			"nginx dnsmasq etcd haproxy vip-manager node-exporter keepalived-exporter pg-exporter pgbackrest-exporter redis-exporter redis valkey silo mcli openbao kes sow pig",
			// 1: infra-addons
			"grafana grafana-plugins grafana-victoriametrics-ds grafana-victorialogs-ds victoria-metrics victoria-logs victoria-traces vlogscli vmutils vector alertmanager",
			// 2: extra-modules
			"blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx docker-ce docker-compose-plugin duckdb restic juicefs vray grafana-infinity-ds",
			// 3: node-package1
			"bash python3 sudo acl ca-certificates openssl curl wget lz4 zstd unzip bzip2 gzip tar tzdata chrony openssh-server util-linux rsync psmisc logrotate",
			// 4: node-package2
			"pv jq git make patch lsof less ncdu htop iotop socat net-tools telnet ipvsadm tuned numactl nvme-cli sysstat keepalived etcd haproxy vector pig uv",
			// 5: node-package3
			"zlib1g libreadline-dev xz-utils locales cron openssh-client node-exporter bind9-dnsutils iproute2 iputils-ping netcat-openbsd procps vim-tiny",
			// 6: pgsql-utility
			"patroni python3-etcd pgbouncer pgbackrest pgbadger pg-timetable pgformatter postgresql-filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
		},

		InfraAliasMap: []PackageMapping{
			{"blackbox_exporter", "blackbox-exporter", "blackbox-exporter"},
			{"kafka_exporter", "kafka-exporter", "kafka-exporter"},
			{"keepalived_exporter", "keepalived-exporter", "keepalived-exporter"},
			{"mongodb_exporter", "mongodb-exporter", "mongodb-exporter"},
			{"mysqld_exporter", "mysqld-exporter", "mysqld-exporter"},
			{"nginx_exporter", "nginx-exporter", "nginx-exporter"},
			{"node_exporter", "node-exporter", "node-exporter"},
			{"redis_exporter", "redis-exporter", "redis-exporter"},
			{"zfs_exporter", "zfs-exporter", "zfs-exporter"},
		},

		PGSQLKernelMap: []PackageMapping{
			{"pgsql", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v"},
			{"pgsql-mini", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib", "postgresql-$v postgresql-client-$v"},
			{"pgsql-core", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v"},
			{"pgsql-full", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit postgresql$v-test postgresql$v-devel", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-server-dev-$v"},
			{"pgsql-main", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl pg_repack_$v wal2json_$v pgvector_$v", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"},
			{"pgsql-client", "postgresql$v", "postgresql-client-$v"},
			{"pgsql-server", "postgresql$v-server postgresql$v-libs postgresql$v-contrib", "postgresql-$v"},
			{"pgsql-devel", "postgresql$v-devel", "postgresql-server-dev-$v"},
			{"pgsql-basic", "pg_repack_$v wal2json_$v pgvector_$v", "postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"},
		},

		PGSQLUtilMap: []PackageMapping{
			{"pgsql-common", "patroni patroni-etcd pgbouncer pgbackrest pg-exporter pgbackrest-exporter vip-manager", "patroni python3-etcd pgbouncer pgbackrest pg-exporter pgbackrest-exporter vip-manager"},
			{"patroni", "patroni patroni-etcd", "patroni python3-etcd"},
			{"pgbouncer", "pgbouncer", "pgbouncer"},
			{"pgbackrest", "pgbackrest", "pgbackrest"},
			{"pg_exporter", "pg-exporter", "pg-exporter"},
			{"pgbackrest_exporter", "pgbackrest-exporter", "pgbackrest-exporter"},
			{"vip-manager", "vip-manager", "vip-manager"},
			{"pg-hardstorage", "pg-hardstorage", "pg-hardstorage"},
			{"pg_hardstorage", "pg-hardstorage", "pg-hardstorage"},
			{"pgbadger", "pgbadger", "pgbadger"},
			{"pg_activity", "pg_activity", "pg-activity"},
			{"pg_filedump", "pg_filedump", "postgresql-filedump"},
			{"pgxnclient", "pgxnclient", "pgxnclient"},
			{"pgformatter", "pgformatter", "pgformatter"},
			{"pgcopydb", "pgcopydb", "pgcopydb"},
			{"pgloader", "pgloader", "pgloader"},
			{"pg_timetable", "pg-timetable", "pg-timetable"},
			{"timescaledb-utils", "timescaledb-tools timescaledb-event-streamer", "timescaledb-tools timescaledb-event-streamer"},
		},

		PGSQLExoticMap: []PackageMapping{

			{"agensgraph", "agensgraph-$v", "agensgraph-$v"},
			{"polardb", "polardb-$v", "polardb-$v"},
			{"openhalo", "openhalodb-$v", "openhalodb-$v"},
			{"ivorysql", "ivorysql-$v ivorysql-$v-contrib", "ivorysql-$v ivorysql-$v-contrib"},
			{"babelfish", "babelfish-$v", "babelfish-$v"},
			{"orioledb", "orioledb-$v", "orioledb-$v"},
			{"pgedge", "pgedge-$v", "pgedge-$v"},
			{"supabase", "pg_tle_$v,pgvector_$v,pg_cron_$v,pgsodium_$v,pg_graphql_$v,pg_jsonschema_$v,wrappers_$v,vault_$v,pgjwt_$v,pgsql_http_$v,pg_net_$v,supautils_$v,index_advisor_$v,safeupdate_$v,pg_plan_filter_$v", "postgresql-$v-pg-tle,postgresql-$v-pg-graphql,postgresql-$v-pg-jsonschema,postgresql-$v-wrappers,postgresql-$v-pgvector,postgresql-$v-cron,postgresql-$v-pgsodium,postgresql-$v-vault,postgresql-$v-pgjwt,postgresql-$v-http,postgresql-$v-pg-net,postgresql-$v-supautils,postgresql-$v-index-advisor,postgresql-$v-pg-safeupdate,postgresql-$v-pg-plan-filter"},
			{"cloudberry", "cloudberry cloudberry-backup cloudberry-pxf", "cloudberry cloudberry-backup cloudberry-pxf"},
			{"pgtde", "pgtde-18,pgtde-18-contrib", "pgtde-18 pgtde-18-contrib"},
		},

		DistroAdhocPkg: map[string]string{
			"rpm":  "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils sow sshpass",
			"deb":  "ansible python3 python3-jmespath sow sshpass",
			"el7":  "ansible python3 python36-requests python36-idna yum-utils sow sshpass",
			"el8":  "ansible python3 python3-requests python3.12-jmespath python3-cryptography dnf-utils sow sshpass",
			"el9":  "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils sow sshpass",
			"el10": "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils sow sshpass crypto-policies-scripts",
			"d11":  "ansible python3 python3-requests python3-jmespath sow sshpass tnftp linux-perf",
			"d12":  "ansible python3 python3-requests python3-jmespath sow sshpass tnftp linux-perf",
			"d13":  "ansible python3 python3-requests python3-jmespath sow sshpass tnftp linux-perf",
			"u22":  "ansible python3 python3-requests python3-jmespath sow sshpass ftp linux-tools-generic",
			"u24":  "ansible python3 python3-requests python3-jmespath sow sshpass ftp linux-tools-generic",
			"u26":  "ansible python3 python3-requests python3-jmespath sow sshpass ftp linux-tools-generic",
		},
	}

	return constants
}

// ============================================================================
// Configuration Templates
// ============================================================================

// rpmTemplate is the template for RPM-based distributions
var rpmTemplate = `---
# {{ .OSName }} distribution ad hoc configuration file
# License    :   Apache-2.0 @ https://pigsty.io/docs/about/license
# Copyright  :   2018-2026  Ruohang Feng / Vonng (rh@vonng.com)

syslog_path: /var/log/messages
node_tuned_profile_dir: {{ .TunedProfileDir }}

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, node-package3, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd,unzip,bzip2,gzip,tar,tzdata,chrony,openssh-server,util-linux,rsync,psmisc,logrotate
  - pv,jq,git,make,patch,lsof,less,ncdu,htop,iotop,socat,net-tools,telnet,ipvsadm,tuned,numactl,nvme-cli,sysstat,keepalived,etcd,haproxy,vector,pig,uv
  - zlib,readline,xz,glibc-langpack-en,cronie,openssh-clients,node-exporter,bind-utils,iproute,iputils,nmap-ncat,procps-ng,vim-minimal,yum,audit,grubby,chkconfig

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,grafana-victorialogs-ds,grafana-victoriametrics-ds,victoria-metrics,victoria-logs,victoria-traces,vmutils,vlogscli,alertmanager
  - node-exporter,blackbox-exporter,nginx-exporter,pg-exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli,restic,certbot,python3-certbot-nginx

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/pgsql-$v'
  citus:  '/usr/pgsql-$v'
  mssql:  '/usr/babelfish-$v/'
  agens:  '/usr/agens-$v'
  ivory:  '/usr/ivory-$v'
  mysql:  '/usr/halo-$v'
  pgtde:  '/usr/pgtde-$v'
  polar:  '/usr/polar-$v'
  oriole: '/usr/oriole-$v'
  pgedge: '/usr/pgedge-$v'
  gpsql:  '/usr/cloudberry'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
` + repo.RPMCatalog + `
# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Constants.RPMCommonPkg 0 }}"
  infra-addons:            "{{ index .Constants.RPMCommonPkg 1 }}"
  extra-modules:           "{{ index .Constants.RPMCommonPkg 2 }}"
  node-package1:           "{{ getNodePackage1 }}"
  node-package2:           "{{ getNodePackage2 }}"
  node-package3:           "{{ getNodePackage3 }}"
  pgsql-utility:           "{{ getPgsqlUtility }}"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl"{{ range .Constants.PGSQLKernelMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#{{ range .Constants.PGSQLExoticMap }}{{ if and (or (eq .Key "greenplum") (eq .Key "gpsql")) (eq $.Arch "aarch64") }}{{ else if .RPM }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}{{ end }}
  java-runtime:            "{{ getJavaRuntime }}"
  kafka-stack:             "kafka kafka-exporter jmx-exporter"
  mysql:                   "mysql-community-server mysql-community-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter"
  kube-runtime:            "containerd.io"
  sealos:                  "sealos"
  kubernetes:              "kubeadm kubelet kubectl"
  docker:                  "docker-ce docker-compose-plugin"
  infra-extra:             "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds rclone mysqld-exporter mongodb-exporter kafka-exporter"
  victoria:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds"
  vmetrics:                "victoria-metrics victoria-metrics-cluster vmutils"
  vlogs:                   "victoria-logs vlogscli vlagent"
  vtraces:                 "victoria-traces"
  tigerbeetle:             "tigerbeetle"
  clickhouse:              "clickhouse-server clickhouse-client clickhouse-common-static"

  # legacy underscore exporter package names{{ range .Constants.InfraAliasMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}

  #--------------------------------#
  # PGSQL: tools & utils
  #--------------------------------#{{ range .Constants.PGSQLUtilMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ getUtilPkg .Key .RPM }}"{{ end }}

  #--------------------------------#
  # PGSQL: extensions (default pg18)
  #--------------------------------#
{{ range $cat := .Categories }}{{ $pkgs := index $.CategoryExts 18 $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pgsql-%s:" $cat) }} "{{ replaceAll "_18" "_$v" $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}

{{ range $pgVer := .PGVersions }}
  #--------------------------------#
  # PG{{ $pgVer }}: packages
  #--------------------------------#{{ range $.Constants.PGSQLKernelMap }}{{ if eq .Key "pgsql" }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .RPM }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .RPM }}"{{ end }}{{ end }}
{{ range $cat := $.Categories }}{{ $pkgs := index $.CategoryExts $pgVer $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer $cat) }} "{{ $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}
{{ end }}
{{ $lastCat := "" }}{{ range .ExtMappings }}{{ if ne .Category $lastCat }}{{ $lastCat = .Category }}

  #--------------------------------#
  # {{ toUpper .Category }}: extension
  #--------------------------------#
{{ end }}  {{ printf "%-24s %-42s" (printf "%s:" .Name) (printf "\"%s\"" .Package) }} # {{ formatPGVersions .PGVers }}
{{ end }}

...
`

// debTemplate is the template for DEB-based distributions
var debTemplate = `---
# {{ .OSName }} distribution ad hoc configuration file
# License    :   Apache-2.0 @ https://pigsty.io/docs/about/license
# Copyright  :   2018-2026  Ruohang Feng / Vonng (rh@vonng.com)

syslog_path: /var/log/syslog
node_tuned_profile_dir: {{ .TunedProfileDir }}

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, node-package3, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd,unzip,bzip2,gzip,tar,tzdata,chrony,openssh-server,util-linux,rsync,psmisc,logrotate
  - pv,jq,git,make,patch,lsof,less,ncdu,htop,iotop,socat,net-tools,telnet,ipvsadm,tuned,numactl,nvme-cli,sysstat,keepalived,etcd,haproxy,vector,pig,uv
  - zlib1g,libreadline-dev,xz-utils,locales,cron,openssh-client,node-exporter,bind9-dnsutils,iproute2,iputils-ping,netcat-openbsd,procps,vim-tiny

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,grafana-victorialogs-ds,grafana-victoriametrics-ds,victoria-metrics,victoria-logs,victoria-traces,vmutils,vlogscli,alertmanager
  - node-exporter,blackbox-exporter,nginx-exporter,pg-exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli,restic,certbot,python3-certbot-nginx

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/lib/postgresql/$v'
  citus:  '/usr/lib/postgresql/$v'
  mssql:  '/usr/babelfish-$v/'
  agens:  '/usr/agens-$v'
  ivory:  '/usr/ivory-$v'
  mysql:  '/usr/halo-$v'
  pgtde:  '/usr/pgtde-$v'
  polar:  '/usr/polar-$v'
  oriole: '/usr/oriole-$v'
  pgedge: '/usr/pgedge-$v'
  gpsql:  '/usr/cloudberry'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
` + repo.DEBCatalog + `
# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Constants.DEBCommonPkg 0 }}"
  infra-addons:            "{{ index .Constants.DEBCommonPkg 1 }}"
  extra-modules:           "{{ index .Constants.DEBCommonPkg 2 }}"
  node-package1:           "{{ getDebNodePackage1 }}"
  node-package2:           "{{ getNodePackage2 }}"
  node-package3:           "{{ getNodePackage3 }}"
  pgsql-utility:           "{{ index .Constants.DEBCommonPkg 6 }}"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql-$v"{{ range .Constants.PGSQLKernelMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#{{ range .Constants.PGSQLExoticMap }}{{ if and (or (eq .Key "greenplum") (eq .Key "gpsql")) (eq $.Arch "aarch64") }}{{ else if .DEB }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}{{ end }}
  java-runtime:            "{{ getJavaRuntime }}"
  kafka-stack:             "kafka kafka-exporter jmx-exporter"
  mysql:                   "mysql-server mysql-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter"
  kube-runtime:            "containerd.io"
  sealos:                  "sealos"
  kubernetes:              "kubeadm kubelet kubectl"
  docker:                  "docker-ce docker-compose-plugin"
  infra-extra:             "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds rclone mysqld-exporter mongodb-exporter kafka-exporter"
  victoria:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds"
  vmetrics:                "victoria-metrics victoria-metrics-cluster vmutils"
  vlogs:                   "victoria-logs vlogscli vlagent"
  vtraces:                 "victoria-traces"
  tigerbeetle:             "tigerbeetle"
  clickhouse:              "clickhouse-server clickhouse-client clickhouse-common-static"

  # legacy underscore exporter package names{{ range .Constants.InfraAliasMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}

  #--------------------------------#
  # PGSQL: tools & utils
  #--------------------------------#{{ range .Constants.PGSQLUtilMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}

  #--------------------------------#
  # PGSQL: extensions (default pg18)
  #--------------------------------#
{{ range $cat := .Categories }}{{ $pkgs := index $.CategoryExts 18 $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pgsql-%s:" $cat) }} "{{ replaceAll "-18" "-$v" $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}

{{ range $pgVer := .PGVersions }}
  #--------------------------------#
  # PG{{ $pgVer }}: packages
  #--------------------------------#{{ range $.Constants.PGSQLKernelMap }}{{ if eq .Key "pgsql" }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB }}"{{ end }}{{ end }}
{{ range $cat := $.Categories }}{{ $pkgs := index $.CategoryExts $pgVer $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer $cat) }} "{{ $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}
{{ end }}
{{ $lastCat := "" }}{{ range .ExtMappings }}{{ if ne .Category $lastCat }}{{ $lastCat = .Category }}

  #--------------------------------#
  # {{ toUpper .Category }}: extension
  #--------------------------------#
{{ end }}  {{ printf "%-24s %-42s" (printf "%s:" .Name) (printf "\"%s\"" .Package) }} # {{ formatPGVersions .PGVers }}
{{ end }}

...
`
