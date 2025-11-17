package cli

import (
	"bytes"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

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

// ExtensionMapping holds extension to package mapping
type ExtensionMapping struct {
	Name     string
	Package  string
	PGVers   []int
	Comment  string
	Category string
}

// GetConfigConstants returns all the constant definitions
func GetConfigConstants() *ConfigConstants {
	constants := &ConfigConstants{
		RPMCommonPkg: []string{
			// 0: infra-package
			"nginx dnsmasq etcd haproxy vip-manager node_exporter keepalived_exporter pg_exporter pgbackrest_exporter redis_exporter redis minio mcli pig",
			// 1: infra-addons
			"grafana grafana-plugins loki logcli vector promtail prometheus alertmanager pushgateway blackbox_exporter nginx_exporter pev2 certbot python3-certbot-nginx",
			// 2: extra-modules
			"docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds",
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
			"grafana grafana-plugins loki logcli vector promtail prometheus alertmanager pushgateway blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx",
			// 2: extra-modules
			"docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds",
			// 3: node-package1
			"lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl sysstat iotop htop rsync tcpdump acl chrony",
			// 4: node-package2
			"netcat-openbsd socat lrzsz net-tools ipvsadm dnsutils telnet ca-certificates libreadline-dev vim-tiny keepalived openssl openssh-server openssh-client",
			// 5: pgsql-utility
			"patroni pgbouncer pgbackrest pgbadger pg-timetable pgformatter postgresql-filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
		},

		PGSQLKernelMap: []PackageMapping{
			{"pgsql", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v"},
			{"pgsql-mini", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib", "postgresql-$v postgresql-client-$v"},
			{"pgsql-core", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v"},
			{"pgsql-full", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit postgresql$v-test postgresql$v-devel", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-server-dev-$v"},
			{"pgsql-main", "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit pg_repack_$v* wal2json_$v* pgvector_$v*", "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"},
			{"pgsql-client", "postgresql$v", "postgresql-client-$v"},
			{"pgsql-server", "postgresql$v-server postgresql$v-libs postgresql$v-contrib", "postgresql-$v"},
			{"pgsql-devel", "postgresql$v-devel", "postgresql-server-dev-$v"},
			{"pgsql-basic", "pg_repack_$v* wal2json_$v* pgvector_$v*", "postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"},
		},

		PGSQLUtilMap: []PackageMapping{
			{"pgsql-common", "patroni patroni-etcd pgbouncer pgbackrest pg_exporter pgbackrest_exporter vip-manager", "patroni pgbouncer pgbackrest pg-exporter pgbackrest-exporter vip-manager"},
			{"patroni", "patroni patroni-etcd", "patroni"},
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
			{"ivorysql", "ivorysql4", "ivorysql-4"},
			{"oriole", "orioledb_17 oriolepg_17", "oriolepg-17 oriolepg-17-orioledb"},
			{"supabase", "pg_tle_$v*,pgvector_$v*,pg_cron_$v*,pgsodium_$v*,pg_graphql_$v,pg_jsonschema_$v,wrappers_$v,vault_$v,pgjwt_$v*,pgsql_http_$v*,pg_net_$v*,supautils_$v*,index_advisor_$v,safeupdate_$v*,pg_plan_filter_$v*", "postgresql-$v-pg-tle,postgresql-$v-pg-graphql,postgresql-$v-pg-jsonschema,postgresql-$v-wrappers,postgresql-$v-pgvector,postgresql-$v-cron,postgresql-$v-pgsodium,postgresql-$v-vault,postgresql-$v-pgjwt,postgresql-$v-http,postgresql-$v-pg-net,postgresql-$v-supautils,postgresql-$v-index-advisor,postgresql-$v-pg-safeupdate,postgresql-$v-pg-plan-filter"},
			{"greenplum", "open-source-greenplum-db-7", ""},
			{"cloudberry", "cloudberry-db cloudberry-hll cloudberry-pgvector", ""},
			{"percona-core", "percona-postgresql17,percona-postgresql17-server,percona-postgresql17-contrib,percona-postgresql17-plperl,percona-postgresql17-plpython3,percona-postgresql17-pltcl", "percona-postgresql-17 percona-postgresql-client-17 percona-postgresql-plperl-17 percona-postgresql-plpython3-17 percona-postgresql-pltcl-17"},
			{"percona-main", "percona-postgresql17,percona-postgresql17-server,percona-postgresql17-contrib,percona-postgresql17-plperl,percona-postgresql17-plpython3,percona-postgresql17-pltcl,percona-postgis33_17,percona-postgis33_17-client,percona-postgis33_17-utils,percona-pgvector_17,percona-wal2json17,percona-pg_repack17,percona-pgaudit17,percona-pgaudit17_set_user,percona-pg_stat_monitor17,percona-pg_gather", "percona-postgresql-17 percona-postgresql-client-17 percona-postgresql-plperl-17 percona-postgresql-plpython3-17 percona-postgresql-pltcl-17 percona-postgresql-17-postgis-3 percona-postgresql-17-pgvector percona-postgresql-17-wal2json percona-postgresql-17-repack percona-postgresql-17-pgaudit percona-pgaudit17-set-user percona-pg-stat-monitor17 percona-pg-gather"},
		},

		DistroAdhocPkg: map[string]string{
			"rpm": "ansible python3 python3-pip python3-virtualenv python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"deb": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass",
			"el7": "ansible python3 python3-pip python36-virtualenv python36-requests python36-idna yum-utils createrepo_c sshpass",
			"el8": "ansible python3 python3-pip python3-virtualenv python3-requests python3.12-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"el9": "ansible python3 python3-pip python3-virtualenv python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"el10": "ansible-core python3 python3-pip python3-virtualenv python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass",
			"d11": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"d12": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"d13": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass tnftp linux-perf",
			"u20": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
			"u22": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
			"u24": "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass ftp linux-tools-generic",
		},
	}

	return constants
}

// ExtensionData holds processed extension information
type ExtensionData struct {
	ID             int            // Extension ID for sorting
	Name           string
	Alias          string
	Package        string
	Category       string
	CategoryID     int            // Category ID for sorting
	Available      map[int]bool   // PG version -> availability
	VersionPackages map[int]string // PG version -> package name from pgext.pkg
	Hidden         bool
	Comment        string
	RPMPkg         string         // RPM package template from extension.rpm_pkg
	DEBPkg         string         // DEB package template from extension.deb_pkg
}

// NewPigstyConfigGenerator creates a new generator with V2 implementation
func NewPigstyConfigGenerator(db *sql.DB, osName string, verbose bool) *PigstyConfigGeneratorV2 {
	parts := strings.Split(osName, ".")
	osCode := parts[0]
	arch := ""
	if len(parts) > 1 {
		arch = parts[1]
	}

	return &PigstyConfigGeneratorV2{
		db:        db,
		osName:    osName,
		osCode:    osCode,
		arch:      arch,
		verbose:   verbose,
		constants: GetConfigConstants(),
	}
}

// PigstyConfigGeneratorV2 is the improved config generator
type PigstyConfigGeneratorV2 struct {
	db        *sql.DB
	osName    string
	osCode    string
	arch      string
	verbose   bool
	constants *ConfigConstants
}

// GenerateContent generates the complete configuration content
func (g *PigstyConfigGeneratorV2) GenerateContent() (string, error) {
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
func (g *PigstyConfigGeneratorV2) loadExtensionData() (map[string]*ExtensionData, error) {
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
		}

		// Update hidden status
		if hide.Valid && hide.Bool {
			extensions[name].Hidden = true
		}
	}

	return extensions, nil
}

// isRPM returns true if this is an RPM-based distribution
func (g *PigstyConfigGeneratorV2) isRPM() bool {
	return strings.HasPrefix(g.osCode, "el")
}

// generateRPMConfig generates configuration for RPM-based systems
func (g *PigstyConfigGeneratorV2) generateRPMConfig(extensions map[string]*ExtensionData) (string, error) {
	tmpl := template.Must(template.New("rpm").Funcs(g.getFuncMap()).Parse(rpmTemplateV2))

	data := g.prepareTemplateData(extensions)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// generateDEBConfig generates configuration for DEB-based systems
func (g *PigstyConfigGeneratorV2) generateDEBConfig(extensions map[string]*ExtensionData) (string, error) {
	tmpl := template.Must(template.New("deb").Funcs(g.getFuncMap()).Parse(debTemplateV2))

	data := g.prepareTemplateData(extensions)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// prepareTemplateData prepares all data for template rendering
func (g *PigstyConfigGeneratorV2) prepareTemplateData(extensions map[string]*ExtensionData) map[string]interface{} {
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
		"Categories":   categories, // Add ordered category list
		"ExtMappings":  g.generateExtensionMappings(extensions),
	}
}

// loadCategories loads categories from database in ID order
func (g *PigstyConfigGeneratorV2) loadCategories() ([]string, error) {
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

// CategoryPkgList holds visible and hidden packages for a category
type CategoryPkgList struct {
	Visible string // Packages in quotes
	Hidden  string // Hidden packages as comments
}

// generateCategoryPackages generates category-based package listings for each PG version
func (g *PigstyConfigGeneratorV2) generateCategoryPackages(extensions map[string]*ExtensionData) map[int]map[string]*CategoryPkgList {
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

			// Collect all packages for this category and version (now in ID order)
			// We use a temporary structure to maintain ordering
			type pkgInfo struct {
				name string
				id   int
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

				// Determine if should be hidden
				isHidden := ext.Hidden || g.shouldHideInCategory(ext.Name, pgVer)

				// Add asterisk for wildcard matching if RPM and not hidden
				if !isHidden && g.isRPM() && !strings.HasSuffix(pkgName, "*") {
					pkgName = pkgName + "*"
				}

				pkgList = append(pkgList, pkgInfo{
					name:   pkgName,
					id:     ext.ID,
					hidden: isHidden,
				})
			}

			// Sort the package list by ID (it should already be sorted, but ensure consistency)
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
func (g *PigstyConfigGeneratorV2) generateExtensionMappings(extensions map[string]*ExtensionData) []ExtensionMapping {
	var mappings []ExtensionMapping

	// Sort extensions by category ID and extension ID
	sortedExts := sortExtensionsByID(extensions)

	// Generate mappings
	for _, ext := range sortedExts {
		// Skip contrib extensions
		if ext.Category == "CONTRIB" {
			continue
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

// getExtensionPackageName returns the actual package name for an extension
func (g *PigstyConfigGeneratorV2) getExtensionPackageName(ext *ExtensionData, pgVer int) string {
	// Special cases
	if ext.Name == "pgaudit" {
		return g.getPGAuditPackageName(pgVer)
	}

	// Query for actual package name from database
	if ext.Package != "" {
		return strings.ReplaceAll(ext.Package, "$v", fmt.Sprintf("%d", pgVer))
	}

	// Generate default package name
	if g.isRPM() {
		return fmt.Sprintf("%s_%d", ext.Alias, pgVer)
	}
	return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(ext.Alias, "_", "-"))
}

// getRPMPackagePattern returns the RPM package pattern for an extension
func (g *PigstyConfigGeneratorV2) getRPMPackagePattern(ext *ExtensionData) string {
	// Special handling for specific extensions
	switch ext.Alias {
	case "pgaudit":
		return "pgaudit_$v*"
	case "hunspell_cs_cz":
		return "hunspell_cs_cz_$v hunspell_de_de_$v hunspell_en_us_$v hunspell_fr_$v hunspell_ne_np_$v hunspell_nl_nl_$v hunspell_nn_no_$v hunspell_ru_ru_$v hunspell_ru_ru_aot_$v"
	default:
		// Use RPMPkg if available, keeping $v placeholder
		if ext.RPMPkg != "" {
			return ext.RPMPkg
		}
		return ext.Alias + "_$v"
	}
}

// getDEBPackagePattern returns the DEB package pattern for an extension
func (g *PigstyConfigGeneratorV2) getDEBPackagePattern(ext *ExtensionData) string {
	// Special handling for specific extensions
	switch ext.Alias {
	case "hunspell_cs_cz":
		return "postgresql-$v-hunspell-cs-cz,postgresql-$v-hunspell-de-de,postgresql-$v-hunspell-en-us,postgresql-$v-hunspell-fr,postgresql-$v-hunspell-ne-np,postgresql-$v-hunspell-nl-nl,postgresql-$v-hunspell-nn-no,postgresql-$v-hunspell-ru-ru,postgresql-$v-hunspell-ru-ru-aot"
	case "pgaudit":
		// Fix: use proper pgaudit naming for DEB systems
		return "postgresql-$v-pgaudit"
	default:
		// Use DEBPkg if available, keeping $v placeholder
		if ext.DEBPkg != "" {
			return ext.DEBPkg
		}
		return "postgresql-$v-" + strings.ReplaceAll(ext.Alias, "_", "-")
	}
}

// shouldHide determines if an extension should be hidden
func (g *PigstyConfigGeneratorV2) shouldHide(extName string, pgVer int) bool {
	hideList := []string{
		"hydra", "duckdb_fdw", "pg_timeseries", "pgpool", "plr",
		"pgagent", "dbt2", "pgtap", "faker", "repmgr", "slony",
		"oracle_fdw", "pg_strom", "db2_fdw", "orioledb",
	}

	for _, h := range hideList {
		if extName == h {
			return true
		}
	}

	// OS/arch specific hiding
	if extName == "rdkit" && g.osCode != "u24" {
		return true
	}

	return false
}

// shouldHideInCategory determines if an extension should be hidden in category packages
func (g *PigstyConfigGeneratorV2) shouldHideInCategory(extName string, pgVer int) bool {
	// Same logic as shouldHide but can have additional category-specific rules
	return g.shouldHide(extName, pgVer)
}

// getPackageNameForCategory returns the package name for use in category lists
func (g *PigstyConfigGeneratorV2) getPackageNameForCategory(ext *ExtensionData, pgVer int) string {
	// Special case handling
	switch ext.Name {
	case "pgaudit":
		// Fix: use proper pgaudit naming for Debian/Ubuntu systems
		if !g.isRPM() {
			// For Debian/Ubuntu, use postgresql-$v-pgaudit pattern
			return fmt.Sprintf("postgresql-%d-pgaudit", pgVer)
		}
		return g.getPGAuditPackageName(pgVer)
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

// getPGAuditPackageName returns the correct pgaudit package name
func (g *PigstyConfigGeneratorV2) getPGAuditPackageName(pgVer int) string {
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
func (g *PigstyConfigGeneratorV2) getExtensionComment(ext *ExtensionData) string {
	comments := map[string]string{
		"pgpool": "exclude due to not used",
		"plr": "exclude due to heavy deps",
		"pgagent": "exclude due to not used",
		"dbt2": "exclude due to not used",
		"pgtap": "exclude due to broken perl deps",
		"faker": "exclude due to broken perl deps",
		"repmgr": "exclude due to not used",
		"slony": "exclude due to not used",
		"oracle_fdw": "exclude due to heavy oracle deps",
		"pg_strom": "exclude due to heavy cuda deps",
		"db2_fdw": "exclude due to heavy db2 deps",
		"orioledb": "work with orioledb postgres fork",
		"pg_tde": "work with percona postgres fork",
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
func (g *PigstyConfigGeneratorV2) getFuncMap() template.FuncMap {
	return template.FuncMap{
		"join":        strings.Join,
		"replaceAll": func(old, new, s string) string {
			// Note: when used with pipe, the piped value comes LAST
			// e.g., {{ $str | replaceAll "old" "new" }} calls replaceAll("old", "new", $str)
			return strings.ReplaceAll(s, old, new)
		},
		"hasPrefix":   strings.HasPrefix,
		"trimPrefix": func(prefix, s string) string {
			// Note: template calls with prefix first, but strings.TrimPrefix expects string first
			return strings.TrimPrefix(s, prefix)
		},
		"toUpper":     strings.ToUpper,
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
	}
}