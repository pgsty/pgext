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

// isEL9ARM returns true if this is el9.aarch64 (has special patroni issues)
func (g *PigstyConfigGenerator) isEL9ARM() bool {
	return g.osCode == "el9" && g.arch == "aarch64"
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
	Comment         string
	RPMPkg          string // RPM package template from extension.rpm_pkg
	DEBPkg          string // DEB package template from extension.deb_pkg
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
				p.pg,
				p.name as pkg_name,
				p.state,
				p.hide
			FROM pgext.extension e
			LEFT JOIN pgext.category c ON c.name = e.category
			LEFT JOIN pgext.pkg p ON p.ext = e.name AND p.os = $1
			WHERE e.lead = true AND e.contrib IS NOT TRUE
		)
		SELECT ext_id, name, alias, category, cat_id, comment, rpm_pkg, deb_pkg, pg, pkg_name, state, hide
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
			extID    int
			name     string
			alias    string
			category string
			catID    int
			comment  sql.NullString
			rpmPkg   sql.NullString
			debPkg   sql.NullString
			pg       sql.NullInt32
			pkgName  sql.NullString
			state    sql.NullString
			hide     sql.NullBool
		)

		err := rows.Scan(&extID, &name, &alias, &category, &catID, &comment, &rpmPkg, &debPkg, &pg, &pkgName, &state, &hide)
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
		"OSName":       g.osName,
		"OSCode":       g.osCode,
		"Arch":         g.arch,
		"Constants":    g.constants,
		"Extensions":   extensions,
		"PGVersions":   []int{18, 17, 16, 15, 14, 13},
		"CategoryExts": catExts,
		"Categories":   categories,
		"ExtMappings":  g.generateExtensionMappings(extensions),
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
	pgVersions := []int{18, 17, 16, 15, 14, 13}

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

// generateExtensionMappings generates extension to package name mappings
func (g *PigstyConfigGenerator) generateExtensionMappings(extensions map[string]*ExtensionData) []ExtensionMapping {
	var mappings []ExtensionMapping

	// Sort extensions by category ID and extension ID
	sortedExts := sortExtensionsByID(extensions)

	// Collect all hunspell extensions for aggregation
	hunspellExts := []string{}
	hunspellVersions := make(map[int]bool)
	for _, ext := range sortedExts {
		if strings.HasPrefix(ext.Name, "hunspell_") && ext.Name != "hunspell_pt_pt" { // Exclude broken hunspell_pt_pt
			hunspellExts = append(hunspellExts, ext.Alias)
			for pg := range ext.Available {
				if ext.Available[pg] {
					hunspellVersions[pg] = true
				}
			}
		}
	}

	// Generate mappings
	for _, ext := range sortedExts {
		// Skip contrib extensions
		if ext.Category == "CONTRIB" {
			continue
		}

		// Special handling: insert hunspell aggregation alias before hunspell_cs_cz
		if ext.Name == "hunspell_cs_cz" && len(hunspellExts) > 0 {
			// Create hunspell aggregate package pattern
			var hunspellPkgPattern string
			if g.isRPM() {
				// For RPM: hunspell_cs_cz_$v hunspell_de_de_$v hunspell_en_us_$v ...
				var pkgs []string
				for _, hext := range hunspellExts {
					pkgs = append(pkgs, hext+"_$v")
				}
				hunspellPkgPattern = strings.Join(pkgs, " ")
			} else {
				// For DEB: postgresql-$v-hunspell-cs-cz,postgresql-$v-hunspell-de-de,...
				var pkgs []string
				for _, hext := range hunspellExts {
					pkgs = append(pkgs, "postgresql-$v-"+strings.ReplaceAll(hext, "_", "-"))
				}
				hunspellPkgPattern = strings.Join(pkgs, ",")
			}

			// Convert map to slice for versions
			var hunspellVerList []int
			for v := range hunspellVersions {
				hunspellVerList = append(hunspellVerList, v)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(hunspellVerList)))

			// Add hunspell aggregate alias
			mappings = append(mappings, ExtensionMapping{
				Name:     "hunspell",
				Package:  hunspellPkgPattern,
				PGVers:   hunspellVerList,
				Comment:  "aggregate all hunspell extensions",
				Category: ext.Category,
			})
		}

		// Determine package pattern
		var pkgPattern string
		if g.isRPM() {
			pkgPattern = g.getRPMPackagePattern(ext)
		} else {
			pkgPattern = g.getDEBPackagePattern(ext)
		}

		// Determine available versions
		var availVersions []int
		for pg := range ext.Available {
			if ext.Available[pg] {
				availVersions = append(availVersions, pg)
			}
		}
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
	// Special case handling
	switch ext.Name {
	case "pgaudit":
		// pgaudit has special versioning for older PG versions
		if g.isRPM() {
			return g.getPGAuditPackageName(pgVer)
		}
		// For Debian/Ubuntu, use standard postgresql-$v-pgaudit pattern
		return fmt.Sprintf("postgresql-%d-pgaudit", pgVer)
	case "babelfishpg_common":
		// Merge babelfish into wiltondb
		if g.osCode == "el7" || g.osCode == "el8" || g.osCode == "el9" ||
			g.osCode == "u20" || g.osCode == "u22" || g.osCode == "u24" {
			return "wiltondb"
		}
	}

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

// getPGAuditPackageName returns the correct pgaudit package name for RPM systems
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
			allVersions := []int{18, 17, 16, 15, 14, 13}
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
			// el10 doesn't have flamegraph package
			basePkgs := "lz4 unzip bzip2 zlib yum pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl grubby sysstat iotop htop rsync tcpdump perf"
			if g.osCode == "el10" {
				return basePkgs + " chkconfig"
			}
			return basePkgs + " flamegraph chkconfig"
		},
		"getDebNodePackage1": func() string {
			// u24 temporarily disables tcpdump
			if g.osCode == "u24" {
				return "lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl sysstat iotop htop rsync acl chrony" // tcpdump
			}
			return g.constants.DEBCommonPkg[3]
		},
		"getNodePackage2": func() string {
			// Handle OS-specific package differences
			if g.isRPM() {
				// RPM systems use bind-utils
				return g.constants.RPMCommonPkg[4]
			}
			// DEB systems - check for Debian 13 special case
			if g.osCode == "d13" {
				// Debian 13 renamed dnsutils to bind9-dnsutils
				return strings.ReplaceAll(g.constants.DEBCommonPkg[4], "dnsutils", "bind9-dnsutils")
			}
			// Other DEB systems use standard dnsutils
			return g.constants.DEBCommonPkg[4]
		},
		"getDNSPackage": func() string {
			// Return the appropriate DNS utilities package name
			if g.isRPM() {
				return "bind-utils"
			}
			if g.osCode == "d13" {
				// Debian 13 renamed dnsutils to bind9-dnsutils
				return "bind9-dnsutils"
			}
			// Other DEB systems use dnsutils
			return "dnsutils"
		},
		"getDockerRepo": func() string {
			// For el10, use RHEL repositories instead of CentOS
			if g.osCode == "el10" {
				return "- { name: docker-ce      ,description: 'Docker CE'          ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/rhel/$releasever/$basearch/stable'      ,china: 'https://mirrors.aliyun.com/docker-ce/linux/rhel/$releasever/$basearch/stable'   ,europe: 'https://mirrors.xtom.de/docker-ce/linux/rhel/$releasever/$basearch/stable' } ,meta: { skip_if_unavailable: 1 }}"
			}
			// For other EL versions, use CentOS repositories
			return "- { name: docker-ce      ,description: 'Docker CE'          ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/centos/$releasever/$basearch/stable'    ,china: 'https://mirrors.aliyun.com/docker-ce/linux/centos/$releasever/$basearch/stable' ,europe: 'https://mirrors.xtom.de/docker-ce/linux/centos/$releasever/$basearch/stable' } ,meta: { skip_if_unavailable: 1 }}"
		},
		"getPgsqlUtility": func() string {
			// For el9.aarch64, lock patroni version due to PGDG upstream issues
			if g.isEL9ARM() {
				return "patroni-4.1.0 patroni-etcd-4.1.0 pgbouncer pgbackrest pgbadger pg_timetable pgFormatter pg_filedump pgxnclient timescaledb-tools timescaledb-event-streamer"
			}
			return g.constants.RPMCommonPkg[5]
		},
		"getUtilPkg": func(key, rpm string) string {
			// For el9.aarch64, lock patroni and pgsql-common versions
			if g.isEL9ARM() {
				if key == "pgsql-common" {
					return "patroni-4.1.0 patroni-etcd-4.1.0 pgbouncer pgbackrest pg_exporter pgbackrest_exporter vip-manager"
				}
				if key == "patroni" {
					return "patroni-4.1.0 patroni-etcd-4.1.0"
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
			"nginx dnsmasq etcd haproxy vip-manager node_exporter keepalived_exporter pg_exporter pgbackrest_exporter redis_exporter redis minio mcli pig",
			// 1: infra-addons
			"grafana grafana-plugins grafana-victoriametrics-ds grafana-victorialogs-ds victoria-metrics victoria-logs victoria-traces vlogscli vmutils vector alertmanager uv",
			// 2: extra-modules
			"blackbox_exporter nginx_exporter pev2 certbot python3-certbot-nginx docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds",
			// 3: node-package1
			"lz4 unzip bzip2 zlib yum pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl grubby sysstat iotop htop rsync tcpdump perf flamegraph chkconfig",
			// 4: node-package2
			"netcat socat ftp lrzsz net-tools ipvsadm bind-utils telnet audit ca-certificates readline vim-minimal keepalived chrony openssl openssh-server openssh-clients",
			// 5: pgsql-utility
			"patroni patroni-etcd pgbouncer pgbackrest pgbadger pg_timetable pgFormatter pg_filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
		},

		DEBCommonPkg: []string{
			// 0: infra-package
			"nginx dnsmasq etcd haproxy vip-manager node-exporter keepalived-exporter pg-exporter pgbackrest-exporter redis-exporter redis minio mcli pig",
			// 1: infra-addons
			"grafana grafana-plugins grafana-victoriametrics-ds grafana-victorialogs-ds victoria-metrics victoria-logs victoria-traces vlogscli vmutils vector alertmanager uv",
			// 2: extra-modules
			"blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds",
			// 3: node-package1
			"lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl sysstat iotop htop rsync tcpdump acl chrony",
			// 4: node-package2
			"netcat-openbsd socat lrzsz net-tools ipvsadm dnsutils telnet ca-certificates libreadline-dev vim-tiny keepalived openssl openssh-server openssh-client",
			// 5: pgsql-utility
			"patroni python3-etcd pgbouncer pgbackrest pgbadger pg-timetable pgformatter postgresql-filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
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
			{"pgsql-common", "patroni patroni-etcd pgbouncer pgbackrest pg_exporter pgbackrest_exporter vip-manager", "patroni python3-etcd pgbouncer pgbackrest pg-exporter pgbackrest-exporter vip-manager"},
			{"patroni", "patroni patroni-etcd", "patroni python3-etcd"},
			{"pgbouncer", "pgbouncer", "pgbouncer"},
			{"pgbackrest", "pgbackrest", "pgbackrest"},
			{"pg_exporter", "pg_exporter", "pg-exporter"},
			{"pgbackrest_exporter", "pgbackrest_exporter", "pgbackrest-exporter"},
			{"vip-manager", "vip-manager", "vip-manager"},
			{"pgbadger", "pgbadger", "pgbadger"},
			{"pg_activity", "pg_activity", "pg-activity"},
			{"pg_filedump", "pg_filedump", "postgresql-filedump"},
			{"pgxnclient", "pgxnclient", "pgxnclient"},
			{"pgformatter", "pgformatter", "pgformatter"},
			{"pgcopydb", "pgcopydb", "pgcopydb"},
			{"pgloader", "pgloader", "pgloader"},
			{"pg_timetable", "pg_timetable", "pg-timetable"},
			{"timescaledb-utils", "timescaledb-tools timescaledb-event-streamer", "timescaledb-tools timescaledb-event-streamer"},
		},

		PGSQLExoticMap: []PackageMapping{
			{"wiltondb", "wiltondb", "wiltondb"},
			{"polardb", "PolarDB", "polardb-for-postgresql"},
			{"openhalodb", "openhalodb", "openhalodb"},
			{"ivorysql", "ivorysql5", "ivorysql-5"},
			{"oriole", "orioledb_17 oriolepg_17", "oriolepg-17 oriolepg-17-orioledb"},
			{"supabase", "pg_tle_$v,pgvector_$v,pg_cron_$v,pgsodium_$v,pg_graphql_$v,pg_jsonschema_$v,wrappers_$v,vault_$v,pgjwt_$v,pgsql_http_$v,pg_net_$v,supautils_$v,index_advisor_$v,safeupdate_$v,pg_plan_filter_$v", "postgresql-$v-pg-tle,postgresql-$v-pg-graphql,postgresql-$v-pg-jsonschema,postgresql-$v-wrappers,postgresql-$v-pgvector,postgresql-$v-cron,postgresql-$v-pgsodium,postgresql-$v-vault,postgresql-$v-pgjwt,postgresql-$v-http,postgresql-$v-pg-net,postgresql-$v-supautils,postgresql-$v-index-advisor,postgresql-$v-pg-safeupdate,postgresql-$v-pg-plan-filter"},
			{"greenplum", "open-source-greenplum-db-7", ""},
			{"cloudberry", "cloudberry-db cloudberry-hll cloudberry-pgvector", ""},
			{"percona-core", "percona-postgresql18,percona-postgresql18-server,percona-postgresql18-contrib,percona-postgresql18-plperl,percona-postgresql18-plpython3,percona-postgresql18-pltcl,percona-pg_tde18", "percona-postgresql-18 percona-postgresql-client-18 percona-postgresql-plperl-18 percona-postgresql-plpython3-18 percona-postgresql-pltcl-18 percona-pg-tde18"},
			{"percona-main", "percona-postgresql18,percona-postgresql18-server,percona-postgresql18-contrib,percona-postgresql18-plperl,percona-postgresql18-plpython3,percona-postgresql18-pltcl,percona-pg_tde18,percona-postgis35_18,percona-postgis35_18-client,percona-postgis35_18-utils,percona-pgvector_18,percona-wal2json18,percona-pg_repack18,percona-pgaudit18,percona-pgaudit18_set_user,percona-pg_stat_monitor18,percona-pg_gather", "percona-postgresql-18 percona-postgresql-client-18 percona-postgresql-plperl-18 percona-postgresql-plpython3-18 percona-postgresql-pltcl-18 percona-pg-tde18 percona-postgresql-18-postgis-3 percona-postgresql-18-pgvector percona-postgresql-18-wal2json percona-postgresql-18-repack percona-postgresql-18-pgaudit percona-pgaudit18-set-user percona-pg-stat-monitor18 percona-pg-gather"},
		},

		DistroAdhocPkg: map[string]string{
			"rpm":  "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"deb":  "ansible python3 python3-jmespath dpkg-dev sshpass",
			"el7":  "ansible python3 python36-requests python36-idna yum-utils createrepo_c sshpass",
			"el8":  "ansible python3 python3-requests python3.12-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"el9":  "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"el10": "ansible python3 python3-requests python3-jmespath python3-cryptography dnf-utils createrepo_c sshpass crypto-policies-scripts",
			"d11":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"d12":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"d13":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"u20":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
			"u22":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
			"u24":  "ansible python3 python3-requests python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
		},
	}

	return constants
}

// ============================================================================
// Configuration Templates
// ============================================================================

// rpmTemplate is the template for RPM-based distributions
const rpmTemplate = `---
# {{ .OSName }} distribution ad hoc configuration file
# License    :   Apache-2.0 @ https://pigsty.io/docs/about/license
# Copyright  :   2018-2026  Ruohang Feng / Vonng (rh@vonng.com)

# where to register systemd files
systemd_dir: /usr/lib/systemd/system
syslog_path: /var/log/messages

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,uuid,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync,tcpdump
  - python3,socat,lrzsz,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,pig
  - zlib,yum,audit,bind-utils,readline,vim-minimal,node_exporter,grubby,openssh-server,openssh-clients,chkconfig,vector

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,grafana-victorialogs-ds,grafana-victoriametrics-ds,victoria-metrics,victoria-logs,victoria-traces,vmutils,vlogscli,alertmanager,uv
  - node_exporter,blackbox_exporter,nginx_exporter,pg_exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli,restic,certbot,python3-certbot-nginx

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/pgsql-$v'
  citus:  '/usr/pgsql-$v'
  mssql:  '/usr/'
  ivory:  '/usr/ivory-5'
  mysql:  '/usr/halo-14'
  gpsql:  '/usr/gpsql'
  polar:  '/u01/polardb_pg'
  oracle: '/u01/polardb_pg'
  oriole: '/usr/oriole-$v'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
repo_upstream_default:
  - { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://${admin_ip}/pigsty' } ,meta: { skip_if_unavailable: 1 ,priority: 1 }} # used by intranet nodes
  - { name: pigsty-infra   ,description: 'Pigsty INFRA'       ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/infra/$basearch' ,china: 'https://repo.pigsty.cc/yum/infra/$basearch' } ,meta: { priority: 12 }}
  - { name: pigsty-pgsql   ,description: 'Pigsty PGSQL'       ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/pgsql/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/pgsql/el$releasever.$basearch' } ,meta: { priority: 11 }}
  - { name: nginx          ,description: 'Nginx Repo'         ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://nginx.org/packages/rhel/$releasever/$basearch/' }}
  - { name: docker-ce      ,description: 'Docker CE'          ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/centos/$releasever/$basearch/stable'    ,china: 'https://mirrors.aliyun.com/docker-ce/linux/centos/$releasever/$basearch/stable' ,europe: 'https://mirrors.xtom.de/docker-ce/linux/centos/$releasever/$basearch/stable' } ,meta: { skip_if_unavailable: 1 }}
  - { name: baseos         ,description: 'EL 8+ BaseOS'       ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/BaseOS/$basearch/os/'                        ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/BaseOS/$basearch/os/'                        ,europe: 'https://mirrors.xtom.de/rocky/$releasever/BaseOS/$basearch/os/'     }}
  - { name: appstream      ,description: 'EL 8+ AppStream'    ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/AppStream/$basearch/os/'                     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/AppStream/$basearch/os/'                     ,europe: 'https://mirrors.xtom.de/rocky/$releasever/AppStream/$basearch/os/'  }}
  - { name: extras         ,description: 'EL 8+ Extras'       ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/extras/$basearch/os/'                        ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/extras/$basearch/os/'                        ,europe: 'https://mirrors.xtom.de/rocky/$releasever/extras/$basearch/os/'     }}
  - { name: powertools     ,description: 'EL 8 PowerTools'    ,module: node    ,releases: [8     ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/PowerTools/$basearch/os/'                    ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/PowerTools/$basearch/os/'                    ,europe: 'https://mirrors.xtom.de/rocky/$releasever/PowerTools/$basearch/os/' }}
  - { name: crb            ,description: 'EL 9 CRB'           ,module: node    ,releases: [  9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/CRB/$basearch/os/'                           ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/CRB/$basearch/os/'                           ,europe: 'https://mirrors.xtom.de/rocky/$releasever/CRB/$basearch/os/'        }}
  - { name: epel           ,description: 'EL 8+ EPEL'         ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever/Everything/$basearch/'               ,china: 'https://mirrors.aliyun.com/epel/$releasever/Everything/$basearch/'                             ,europe: 'https://mirrors.xtom.de/epel/$releasever/Everything/$basearch/'     }}
  - { name: pgdg-common    ,description: 'PostgreSQL Common'  ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/redhat/rhel-$releasever-$basearch'      ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/common/redhat/rhel-$releasever-$basearch'      ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg13         ,description: 'PostgreSQL 13'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/13/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/13/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch'          ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch'          ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg-beta      ,description: 'PostgreSQL Testing' ,module: beta    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/testing/19/redhat/rhel-$releasever-$basearch'  ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch'  ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch'  }}
  - { name: pgdg-extras    ,description: 'PostgreSQL Extra'   ,module: extra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/extras/redhat/rhel-$releasever-$basearch'      ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/extras/redhat/rhel-$releasever-$basearch'      ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/extras/redhat/rhel-$releasever-$basearch'      }}
  - { name: pgdg13-nonfree ,description: 'PostgreSQL 13+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: pgdg14-nonfree ,description: 'PostgreSQL 14+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: pgdg15-nonfree ,description: 'PostgreSQL 15+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: pgdg16-nonfree ,description: 'PostgreSQL 16+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: pgdg17-nonfree ,description: 'PostgreSQL 17+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: pgdg18-nonfree ,description: 'PostgreSQL 18+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-$releasever-$basearch' ,china: 'https://mirrors.aliyun.com/postgresql/repos/yum/non-free/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/18/redhat/rhel-$releasever-$basearch' } ,meta: { skip_if_unavailable: 1 }}
  - { name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/timescale/timescaledb/el/$releasever/$basearch'  }}
  - { name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/percona/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/percona/el$releasever.$basearch' ,origin: 'http://repo.percona.com/ppg-18.1/yum/release/$releasever/RPMS/$basearch'  }}
  - { name: wiltondb       ,description: 'WiltonDB'           ,module: mssql   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/mssql/el$releasever.$basearch', china: 'https://repo.pigsty.cc/yum/mssql/el$releasever.$basearch' , origin: 'https://download.copr.fedorainfracloud.org/results/wiltondb/wiltondb/epel-$releasever-$basearch/' }}
  - { name: groonga        ,description: 'Groonga'            ,module: groonga ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.groonga.org/almalinux/$releasever/$basearch/' }}
  - { name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mysql.com/yum/mysql-8.4-community/el/$releasever/$basearch/' }}
  - { name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' ,china: 'https://mirrors.aliyun.com/mongodb/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' }}
  - { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpmfind.net/linux/remi/enterprise/$releasever/redis72/$basearch/' }}
  - { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpm.grafana.com', china: 'https://mirrors.aliyun.com/grafana/yum/' }}
  - { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/rpm/', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/rpm/' }}
  - { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/el/$releasever/$basearch' }}
  - { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/el/$releasever/$basearch' }}
  - { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/rpm/stable/', china: 'https://mirrors.aliyun.com/clickhouse/rpm/stable/' }}

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
  node-package2:           "{{ index .Constants.RPMCommonPkg 4 }}"
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
  kafka:                   "kafka kafka_exporter"
  kube-runtime:            "containerd.io"
  sealos:                  "sealos"
  kubernetes:              "kubeadm kubelet kubectl"
  docker:                  "docker-ce docker-compose-plugin"
  infra-extra:             "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds rclone mysqld_exporter mongodb_exporter kafka_exporter"
  victoria:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent victoria-traces grafana-victorialogs-ds"
  vmetrics:                "victoria-metrics victoria-metrics-cluster vmutils"
  vlogs:                   "victoria-logs vlogscli vlagent"
  vtraces:                 "victoria-traces"
  tigerbeetle:             "tigerbeetle"
  clickhouse:              "clickhouse-server clickhouse-client clickhouse-common-static"

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

...`

// debTemplate is the template for DEB-based distributions
const debTemplate = `---
# {{ .OSName }} distribution ad hoc configuration file
# License    :   Apache-2.0 @ https://pigsty.io/docs/about/license
# Copyright  :   2018-2026  Ruohang Feng / Vonng (rh@vonng.com)

# where to register systemd files
systemd_dir: /lib/systemd/system
syslog_path: /var/log/syslog

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,uuid,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync{{ if ne .OSCode "u24" }},tcpdump{{ else }} #tcpdump{{ end }}
  - python3,socat,lrzsz,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,pig
  - zlib1g,acl,{{ getDNSPackage }},libreadline-dev,vim-tiny,node-exporter,openssh-server,openssh-client,vector

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,grafana-victorialogs-ds,grafana-victoriametrics-ds,victoria-metrics,victoria-logs,victoria-traces,vmutils,vlogscli,alertmanager,uv
  - node-exporter,blackbox-exporter,nginx-exporter,pg-exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli,restic,certbot,python3-certbot-nginx

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/lib/postgresql/$v'
  citus:  '/usr/lib/postgresql/$v'
  mssql:  '/usr/lib/postgresql/$v'
  ivory:  '/usr/ivory-5'
  mysql:  '/usr/halo-14'
  polar:  '/u01/polardb_pg'
  oracle: '/u01/polardb_pg'
  oriole: '/usr/oriole-$v'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
repo_upstream_default:
  - { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://${admin_ip}/pigsty ./' }}
  - { name: pigsty-pgsql   ,description: 'Pigsty PgSQL'       ,module: pgsql   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main', china: 'https://repo.pigsty.cc/apt/pgsql/${distro_codename} ${distro_codename} main' }}
  - { name: pigsty-infra   ,description: 'Pigsty Infra'       ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/infra/ generic main' ,china: 'https://repo.pigsty.cc/apt/infra/ generic main' }}
  - { name: nginx          ,description: 'Nginx'              ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://nginx.org/packages/${distro_name} ${distro_codename} nginx' }}
  - { name: docker-ce      ,description: 'Docker'             ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/${distro_name} ${distro_codename} stable'                               ,china: 'https://mirrors.aliyun.com/docker-ce/linux/${distro_name} ${distro_codename} stable' }}
  - { name: base           ,description: 'Debian Basic'       ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename} main non-free-firmware'                                  ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main restricted universe multiverse' }}
  - { name: updates        ,description: 'Debian Updates'     ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename}-updates main non-free-firmware'                          ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main restricted universe multiverse' }}
  - { name: security       ,description: 'Debian Security'    ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://security.debian.org/debian-security ${distro_codename}-security main non-free-firmware'            ,china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main non-free-firmware' }}
  - { name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}           main universe multiverse restricted' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}           main restricted universe multiverse' }}
  - { name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-backports main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-updates   main restricted universe multiverse' }}
  - { name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-security  main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-backports main restricted universe multiverse' }}
  - { name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-updates   main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-security  main restricted universe multiverse' }}
  - { name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}             main universe multiverse restricted' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}           main restricted universe multiverse' }}
  - { name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-backports   main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-updates   main restricted universe multiverse' }}
  - { name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-security    main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-backports main restricted universe multiverse' }}
  - { name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-updates     main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-security  main restricted universe multiverse' }}
  - { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main' ,china: 'https://mirrors.aliyun.com/postgresql/repos/apt/ ${distro_codename}-pgdg main' }}
  - { name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta    ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg-testing main 19' ,china: 'https://mirrors.aliyun.com/postgresql/repos/apt/ ${distro_codename}-pgdg-testing main 19' }}
  - { name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/timescale/timescaledb/${distro_name}/ ${distro_codename} main' }}
  - { name: citus          ,description: 'Citus'              ,module: extra   ,releases: [11,12,   20,22   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/citusdata/community/${distro_name}/ ${distro_codename} main' } }
  - { name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/percona ${distro_codename} main' ,china: 'https://repo.pigsty.cc/apt/percona ${distro_codename} main' ,origin: 'http://repo.percona.com/ppg-18.1/apt ${distro_codename} main' }}
  - { name: wiltondb       ,description: 'WiltonDB'           ,module: mssql   ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/mssql/ ${distro_codename} main'  ,china: 'https://repo.pigsty.cc/apt/mssql/ ${distro_codename} main'  ,origin: 'https://ppa.launchpadcontent.net/wiltondb/wiltondb/ubuntu/ ${distro_codename} main'  }}
  - { name: groonga        ,description: 'Groonga Debian'     ,module: groonga ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.groonga.org/debian/ ${distro_codename} main' }}
  - { name: groonga        ,description: 'Groonga Ubuntu'     ,module: groonga ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://ppa.launchpadcontent.net/groonga/ppa/ubuntu/ ${distro_codename} main' }}
  - { name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mysql.com/apt/${distro_name} ${distro_codename} mysql-8.0 mysql-tools', china: 'https://mirrors.tuna.tsinghua.edu.cn/mysql/apt/${distro_name} ${distro_codename} mysql-8.0 mysql-tools' }}
  - { name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mongodb.org/apt/${distro_name} ${distro_codename}/mongodb-org/8.0 multiverse', china: 'https://mirrors.aliyun.com/mongodb/apt/${distro_name} ${distro_codename}/mongodb-org/8.0 multiverse' }}
  - { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.redis.io/deb ${distro_codename} main' }}
  - { name: llvm           ,description: 'LLVM'               ,module: llvm    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.llvm.org/${distro_codename}/ llvm-toolchain-${distro_codename} main' ,china: 'https://mirrors.tuna.tsinghua.edu.cn/llvm-apt/${distro_codename}/ llvm-toolchain-${distro_codename} main' }}  
  - { name: haproxyd       ,description: 'Haproxy Debian'     ,module: haproxy ,releases: [11,12            ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://haproxy.debian.net/ ${distro_codename}-backports-3.1 main' }}
  - { name: haproxyu       ,description: 'Haproxy Ubuntu'     ,module: haproxy ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://ppa.launchpadcontent.net/vbernat/haproxy-3.1/ubuntu/ ${distro_codename} main' }}
  - { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://apt.grafana.com stable main' ,china: 'https://mirrors.aliyun.com/grafana/apt/ stable main' }}
  - { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/deb/ /', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/deb/ /' }}
  - { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/${distro_name}/ ${distro_codename} main' }}
  - { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/${distro_name}/ ${distro_codename} main' }}
  - { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/deb/ stable main', china: 'https://mirrors.aliyun.com/clickhouse/deb/ stable main' }}


# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Constants.DEBCommonPkg 0 }}"
  infra-addons:            "{{ index .Constants.DEBCommonPkg 1 }}"
  extra-modules:           "{{ index .Constants.DEBCommonPkg 2 }}"
  node-package1:           "{{ getDebNodePackage1 }}"{{ if eq .OSCode "u24" }} #tcpdump{{ end }}
  node-package2:           "{{ getNodePackage2 }}"
  pgsql-utility:           "{{ index .Constants.DEBCommonPkg 5 }}"

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
  kafka:                   "kafka kafka-exporter"
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

...`
