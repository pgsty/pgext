package cli

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"pgext/cli/repo"
	"pgext/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func el9ARMPatroniPackages(extra ...string) string {
	packages := append([]string{"patroni.noarch", "patroni-etcd.noarch"}, extra...)
	return strings.Join(packages, " ")
}

// GeneratePigstyConfig generates Pigsty configuration for a specific OS.
func GeneratePigstyConfig(osName string, outputDir string, dryRun bool, verbose bool) error {
	if !db.SupportsOSPlatform(osName) {
		return fmt.Errorf("unsupported configuration target %q", osName)
	}
	if strings.Count(osName, ".") != 1 {
		return fmt.Errorf("invalid OS name format: %s (expected format: osCode.arch, e.g., el9.x86_64)", osName)
	}
	if DB == nil {
		return fmt.Errorf("database connection pool not initialized")
	}

	generator := NewPigstyConfigGenerator(DB, osName, verbose)
	content, err := generator.GenerateContent()
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	outputFile := filepath.Join(outputDir, osName+".yml")
	if dryRun {
		if verbose {
			logrus.Infof("Dry run mode - would write to: %s", outputFile)
			fmt.Println(content)
		}
		return nil
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	if verbose {
		logrus.Infof("Written configuration to: %s", outputFile)
	}
	return nil
}

// PigstyConfigGenerator handles the generation of Pigsty configuration files.
type PigstyConfigGenerator struct {
	db        *pgxpool.Pool
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

// NewPigstyConfigGenerator creates a new generator.
func NewPigstyConfigGenerator(pool *pgxpool.Pool, osName string, verbose bool) *PigstyConfigGenerator {
	osCode, arch, _ := strings.Cut(osName, ".")
	return &PigstyConfigGenerator{
		db:        pool,
		osName:    osName,
		osCode:    osCode,
		arch:      arch,
		verbose:   verbose,
		constants: GetConfigConstants(),
	}
}

// ExtensionData holds processed extension information.
type ExtensionData struct {
	ID              int // Extension ID for sorting
	Name            string
	Alias           string
	Category        string
	CategoryID      int            // Category ID for sorting
	Available       map[int]bool   // PG version -> availability
	VersionPackages map[int]string // PG version -> package name from pgext.pkg
	VersionHidden   map[int]bool   // PG version -> hidden status
	ForkBound       bool           // true when tags contains fork; render only via manual fork aliases
	Comment         string
	RPMPkg          string // RPM package template from extension.rpm_pkg
	DEBPkg          string // DEB package template from extension.deb_pkg
	RPMPG           []int  // Declared RPM PG versions from extension.rpm_pg
	DEBPG           []int  // Declared DEB PG versions from extension.deb_pg
}

// CategoryPkgList holds visible and hidden packages for a category.
type CategoryPkgList struct {
	Visible string // Packages in quotes
	Hidden  string // Hidden packages as comments
}

// ExtensionMapping holds extension to package mapping.
type ExtensionMapping struct {
	Name     string
	Package  string
	PGVers   []int
	Comment  string
	Category string
}

// GenerateContent generates the complete configuration content.
func (g *PigstyConfigGenerator) GenerateContent() (string, error) {
	ctx := context.Background()
	extensions, err := g.loadExtensionData(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to load extension data: %w", err)
	}
	return g.renderConfig(g.prepareTemplateData(ctx, extensions))
}

func (g *PigstyConfigGenerator) loadExtensionData(ctx context.Context) (map[string]*ExtensionData, error) {
	const query = `
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

	rows, err := g.db.Query(ctx, query, g.osName)
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
			rpmPG     []string
			debPG     []string
			forkBound bool
			pg        sql.NullInt32
			pkgName   sql.NullString
			state     sql.NullString
			hide      sql.NullBool
		)
		if err := rows.Scan(&extID, &name, &alias, &category, &catID, &comment, &rpmPkg, &debPkg, &rpmPG, &debPG, &forkBound, &pg, &pkgName, &state, &hide); err != nil {
			return nil, err
		}

		ext, exists := extensions[name]
		if !exists {
			ext = &ExtensionData{
				ID:              extID,
				Name:            name,
				Alias:           alias,
				Category:        category,
				CategoryID:      catID,
				Available:       make(map[int]bool),
				VersionPackages: make(map[int]string),
				VersionHidden:   make(map[int]bool),
				ForkBound:       forkBound,
				Comment:         comment.String,
				RPMPkg:          rpmPkg.String,
				DEBPkg:          debPkg.String,
				RPMPG:           ParsePGVersions(rpmPG),
				DEBPG:           ParsePGVersions(debPG),
			}
			extensions[name] = ext
		}

		if pg.Valid && state.Valid && state.String == "AVAIL" {
			pgVersion := int(pg.Int32)
			ext.Available[pgVersion] = true
			if pkgName.Valid && pkgName.String != "" {
				ext.VersionPackages[pgVersion] = pkgName.String
			}
			if hide.Valid {
				ext.VersionHidden[pgVersion] = hide.Bool
			}
		}
	}
	return extensions, rows.Err()
}

func (g *PigstyConfigGenerator) isRPM() bool {
	return strings.HasPrefix(g.osCode, "el")
}

func (g *PigstyConfigGenerator) renderConfig(data map[string]interface{}) (string, error) {
	tmpl, err := template.New("pigsty-config").Funcs(g.getFuncMap()).Parse(pigstyConfigTemplate)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type familyConfig struct {
	SyslogPath    string
	NodePkg3Line  string
	PGHome        string
	RepoCatalog   string
	CommonPkg     []string
	PostgreSQLPkg string
	MySQLPkg      string
	VerSep        string
}

func (g *PigstyConfigGenerator) familyConfig() familyConfig {
	if g.isRPM() {
		return familyConfig{
			SyslogPath:    "/var/log/messages",
			NodePkg3Line:  "zlib,readline,xz,glibc-langpack-en,cronie,openssh-clients,node-exporter,bind-utils,iproute,iputils,nmap-ncat,procps-ng,vim-minimal,yum,audit,grubby,chkconfig",
			PGHome:        "/usr/pgsql-$v",
			RepoCatalog:   repo.RPMCatalog,
			CommonPkg:     g.constants.RPMCommonPkg,
			PostgreSQLPkg: "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl",
			MySQLPkg:      "mysql-community-server mysql-community-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter",
			VerSep:        "_",
		}
	}
	return familyConfig{
		SyslogPath:    "/var/log/syslog",
		NodePkg3Line:  "zlib1g,libreadline-dev,xz-utils,locales,cron,openssh-client,node-exporter,bind9-dnsutils,iproute2,iputils-ping,netcat-openbsd,procps,vim-tiny",
		PGHome:        "/usr/lib/postgresql/$v",
		RepoCatalog:   repo.DEBCatalog,
		CommonPkg:     g.constants.DEBCommonPkg,
		PostgreSQLPkg: "postgresql-$v",
		MySQLPkg:      "mysql-server mysql-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter",
		VerSep:        "-",
	}
}

func (g *PigstyConfigGenerator) prepareTemplateData(ctx context.Context, extensions map[string]*ExtensionData) map[string]interface{} {
	categories, err := g.loadCategories(ctx)
	if err != nil {
		categories = []string{"time", "gis", "rag", "fts", "olap", "feat", "lang", "type", "util", "func", "admin", "stat", "sec", "fdw", "sim", "etl"}
	}
	return g.buildTemplateData(extensions, categories, g.generateCategoryPackages(extensions, categories), g.generateExtensionMappings(extensions))
}

func (g *PigstyConfigGenerator) buildTemplateData(
	extensions map[string]*ExtensionData,
	categories []string,
	catExts map[int]map[string]*CategoryPkgList,
	extMappings []ExtensionMapping,
) map[string]interface{} {
	return map[string]interface{}{
		"OSName":          g.osName,
		"OSCode":          g.osCode,
		"Arch":            g.arch,
		"Family":          g.familyConfig(),
		"TunedProfileDir": g.tunedProfileDir(),
		"Constants":       g.constants,
		"Extensions":      extensions,
		"PGVersions":      db.SupportedPGVersionsForOS(g.osCode),
		"DefaultPG":       db.DefaultPGVersion,
		"CategoryExts":    catExts,
		"Categories":      categories,
		"ExtMappings":     extMappings,
	}
}

func (g *PigstyConfigGenerator) loadCategories(ctx context.Context) ([]string, error) {
	if g.db == nil {
		return nil, fmt.Errorf("database pool not initialized")
	}
	rows, err := g.db.Query(ctx, `SELECT LOWER(name) FROM pgext.category ORDER BY id`)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if g.verbose && len(categories) > 0 {
		fmt.Printf("Loaded %d categories from database in ID order: %v\n", len(categories), categories)
	}
	return categories, nil
}

func (g *PigstyConfigGenerator) generateCategoryPackages(extensions map[string]*ExtensionData, categories []string) map[int]map[string]*CategoryPkgList {
	result := make(map[int]map[string]*CategoryPkgList)
	pgVersions := db.SupportedPGVersionsForOS(g.osCode)
	sortedExtensions := sortExtensionsByID(extensions)

	for _, pgVer := range pgVersions {
		result[pgVer] = make(map[string]*CategoryPkgList)
		for _, cat := range categories {
			var packages, hiddenPackages []string
			for _, ext := range sortedExtensions {
				if ext.ForkBound || strings.ToLower(ext.Category) != cat || !ext.Available[pgVer] {
					continue
				}
				pkgName := g.getPackageNameForCategory(ext, pgVer)
				if ext.VersionHidden[pgVer] {
					hiddenPackages = append(hiddenPackages, pkgName)
				} else {
					packages = append(packages, pkgName)
				}
			}
			if len(packages) > 0 || len(hiddenPackages) > 0 {
				result[pgVer][cat] = &CategoryPkgList{
					Visible: strings.Join(packages, " "),
					Hidden:  strings.Join(hiddenPackages, " "),
				}
			}
		}
	}
	return result
}

func (g *PigstyConfigGenerator) getExtensionMappingVersions(ext *ExtensionData) []int {
	var versions []int
	for pg, avail := range ext.Available {
		if avail {
			versions = append(versions, pg)
		}
	}
	if len(versions) > 0 {
		return versions
	}
	declared := ext.DEBPG
	if g.isRPM() {
		declared = ext.RPMPG
	}
	for _, pg := range declared {
		if db.SupportsOSPG(g.osCode, pg) {
			versions = append(versions, pg)
		}
	}
	return versions
}

func (g *PigstyConfigGenerator) generateExtensionMappings(extensions map[string]*ExtensionData) []ExtensionMapping {
	var mappings []ExtensionMapping
	seenAliases := make(map[string]struct{})

	for _, ext := range sortExtensionsByID(extensions) {
		if ext.Category == "CONTRIB" || ext.ForkBound {
			continue
		}
		if _, exists := seenAliases[ext.Alias]; exists {
			continue
		}
		seenAliases[ext.Alias] = struct{}{}

		pkgPattern := g.getDEBPackagePattern(ext)
		if g.isRPM() {
			pkgPattern = g.getRPMPackagePattern(ext)
		}
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

func (g *PigstyConfigGenerator) getPackageNameForCategory(ext *ExtensionData, pgVer int) string {
	verStr := fmt.Sprintf("%d", pgVer)
	if g.isRPM() {
		if ext.RPMPkg != "" {
			return strings.ReplaceAll(ext.RPMPkg, "$v", verStr)
		}
		if pkgName := ext.VersionPackages[pgVer]; pkgName != "" {
			return pkgName
		}
		name := ext.Alias
		if name == "" {
			name = ext.Name
		}
		return fmt.Sprintf("%s_%d", name, pgVer)
	}
	if ext.DEBPkg != "" {
		return strings.ReplaceAll(ext.DEBPkg, "$v", verStr)
	}
	if pkgName := ext.VersionPackages[pgVer]; pkgName != "" {
		return pkgName
	}
	name := ext.Alias
	if name == "" {
		name = ext.Name
	}
	return fmt.Sprintf("postgresql-%d-%s", pgVer, strings.ReplaceAll(name, "_", "-"))
}

func (g *PigstyConfigGenerator) getRPMPackagePattern(ext *ExtensionData) string {
	if ext.RPMPkg != "" {
		return ext.RPMPkg
	}
	return ext.Alias + "_$v"
}

func (g *PigstyConfigGenerator) getDEBPackagePattern(ext *ExtensionData) string {
	if ext.DEBPkg != "" {
		return ext.DEBPkg
	}
	return "postgresql-$v-" + strings.ReplaceAll(ext.Alias, "_", "-")
}

var extensionAliasComments = map[string]string{
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

func (g *PigstyConfigGenerator) getExtensionComment(ext *ExtensionData) string {
	if comment, ok := extensionAliasComments[ext.Alias]; ok {
		return comment
	}
	return ext.Comment
}

func (g *PigstyConfigGenerator) getFuncMap() template.FuncMap {
	return template.FuncMap{
		"replaceAll": func(old, new, s string) string { return strings.ReplaceAll(s, old, new) },
		"trimPrefix": func(prefix, s string) string { return strings.TrimPrefix(s, prefix) },
		"toUpper":    strings.ToUpper,
		"mapPkg": func(m PackageMapping) string {
			if g.isRPM() {
				return m.RPM
			}
			return m.DEB
		},
		"supportsKernelAlias": func(key string) bool {
			switch key {
			case "greenplum", "gpsql":
				return g.arch != "aarch64"
			case "openhalo", "openhalodb", "openhalo-19", "openhalo19":
				return db.SupportsOSPG(g.osCode, 19)
			}
			return true
		},
		"formatPGVersions": func(versions []int) string {
			avail := make(map[int]bool, len(versions))
			for _, av := range versions {
				avail[av] = true
			}
			allVersions := db.SupportedPGVersionsForOS(g.osCode)
			result := make([]string, len(allVersions))
			for i, v := range allVersions {
				if avail[v] {
					result[i] = fmt.Sprintf("%2d", v)
				} else {
					result[i] = "  "
				}
			}
			return strings.Join(result, " ")
		},
		"getBootstrap": func() string {
			if pkg, ok := g.constants.DistroAdhocPkg[g.osCode]; ok {
				return pkg
			}
			if g.isRPM() {
				return g.constants.DistroAdhocPkg["rpm"]
			}
			return g.constants.DistroAdhocPkg["deb"]
		},
		"getJavaRuntime": func() string {
			if g.isRPM() {
				if g.osCode == "el10" {
					return "java-21-openjdk-src java-21-openjdk-headless"
				}
				return "java-17-openjdk-src java-17-openjdk-headless"
			}
			if g.osCode == "d13" {
				return "openjdk-21-jdk"
			}
			return "openjdk-17-jdk"
		},
		"getNodePackage1": func() string {
			if g.isRPM() {
				return g.constants.RPMCommonPkg[3]
			}
			return g.constants.DEBCommonPkg[3]
		},
		"getDebNodePackage1": func() string { return g.constants.DEBCommonPkg[3] },
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
			basePkgs := g.constants.RPMCommonPkg[5] + " netcat ftp perf"
			if g.osCode == "el10" {
				return basePkgs
			}
			return basePkgs + " flamegraph"
		},
		"getPgsqlUtility": func() string {
			if !g.isRPM() {
				return g.constants.DEBCommonPkg[6]
			}
			if g.isEL9ARM() {
				return el9ARMPatroniPackages("pgbouncer", "pgbackrest", "pgbadger", "pg-timetable", "pgFormatter", "pg_filedump", "pgxnclient", "timescaledb-tools", "timescaledb-event-streamer")
			}
			return g.constants.RPMCommonPkg[6]
		},
		"getUtilPkg": func(key, val string) string {
			if g.isEL9ARM() {
				if key == "pgsql-common" {
					return el9ARMPatroniPackages("pgbouncer", "pgbackrest", "pg-exporter", "pgbackrest-exporter", "vip-manager")
				}
				if key == "patroni" {
					return el9ARMPatroniPackages()
				}
			}
			return val
		},
	}
}

func sortExtensionsByID(extensions map[string]*ExtensionData) []*ExtensionData {
	extSlice := make([]*ExtensionData, 0, len(extensions))
	for _, ext := range extensions {
		extSlice = append(extSlice, ext)
	}
	sort.Slice(extSlice, func(i, j int) bool {
		if extSlice[i].CategoryID != extSlice[j].CategoryID {
			return extSlice[i].CategoryID < extSlice[j].CategoryID
		}
		return extSlice[i].ID < extSlice[j].ID
	})
	return extSlice
}

// ConfigConstants holds package group definitions for Pigsty node_id vars.
type ConfigConstants struct {
	RPMCommonPkg   []string
	DEBCommonPkg   []string
	InfraAliasMap  []PackageMapping
	PGSQLKernelMap []PackageMapping
	PGSQLUtilMap   []PackageMapping
	PGSQLExoticMap []PackageMapping
	DistroAdhocPkg map[string]string
}

// PackageMapping holds package name mappings between RPM and DEB.
type PackageMapping struct {
	Key string
	RPM string
	DEB string
}

// GetConfigConstants returns all constant package definitions.
func GetConfigConstants() *ConfigConstants {
	const (
		commonInfra  = "nginx dnsmasq etcd haproxy vip-manager node-exporter keepalived-exporter pg-exporter pgbackrest-exporter redis-exporter redis valkey silo mcli openbao kes sow pig"
		commonAddons = "grafana grafana-plugins grafana-victoriametrics-ds grafana-victorialogs-ds victoria-metrics victoria-logs victoria-traces vlogscli vmutils vector alertmanager"
		commonExtra  = "blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx docker-ce docker-compose-plugin duckdb restic juicefs vray grafana-infinity-ds"
		commonNode1  = "bash python3 sudo acl ca-certificates openssl curl wget lz4 zstd unzip bzip2 gzip tar tzdata chrony openssh-server util-linux rsync psmisc logrotate"
		commonNode2  = "pv jq git make patch lsof less ncdu htop iotop socat net-tools telnet ipvsadm tuned numactl nvme-cli sysstat keepalived etcd haproxy vector pig uv"
	)
	return &ConfigConstants{
		RPMCommonPkg: []string{
			commonInfra,
			commonAddons,
			commonExtra,
			commonNode1,
			commonNode2,
			"zlib readline xz glibc-langpack-en cronie openssh-clients node-exporter bind-utils iproute iputils nmap-ncat procps-ng vim-minimal yum audit grubby chkconfig",
			"patroni patroni-etcd pgbouncer pgbackrest pgbadger pg-timetable pgFormatter pg_filedump pgxnclient timescaledb-tools timescaledb-event-streamer",
		},
		DEBCommonPkg: []string{
			commonInfra,
			commonAddons,
			commonExtra,
			commonNode1,
			commonNode2,
			"zlib1g libreadline-dev xz-utils locales cron openssh-client node-exporter bind9-dnsutils iproute2 iputils-ping netcat-openbsd procps vim-tiny",
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
			{"openhalo", "openhalodb-19", "openhalodb-19"},
			{"openhalodb", "openhalodb-19", "openhalodb-19"},
			{"openhalo-19", "openhalodb-19", "openhalodb-19"},
			{"openhalo19", "openhalodb-19", "openhalodb-19"},
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
}

// ============================================================================
// Unified Configuration Template (RPM & DEB)
// ============================================================================

var pigstyConfigTemplate = `---
# {{ .OSName }} distribution ad hoc configuration file
# License    :   Apache-2.0 @ https://pigsty.io/docs/about/license
# Copyright  :   2018-2026  Ruohang Feng / Vonng (rh@vonng.com)

syslog_path: {{ .Family.SyslogPath }}
node_tuned_profile_dir: {{ .TunedProfileDir }}

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, node-package3, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd,unzip,bzip2,gzip,tar,tzdata,chrony,openssh-server,util-linux,rsync,psmisc,logrotate
  - pv,jq,git,make,patch,lsof,less,ncdu,htop,iotop,socat,net-tools,telnet,ipvsadm,tuned,numactl,nvme-cli,sysstat,keepalived,etcd,haproxy,vector,pig,uv
  - {{ .Family.NodePkg3Line }}

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,grafana-victorialogs-ds,grafana-victoriametrics-ds,victoria-metrics,victoria-logs,victoria-traces,vmutils,vlogscli,alertmanager
  - node-exporter,blackbox-exporter,nginx-exporter,pg-exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli,restic,certbot,python3-certbot-nginx

# postgres home dir in various mode
pg_home_map:
  pgsql:  '{{ .Family.PGHome }}'
  citus:  '{{ .Family.PGHome }}'
  mssql:  '/usr/babelfish-$v/'
  agens:  '/usr/agens-$v'
  ivory:  '/usr/ivory-$v'{{ if supportsKernelAlias "openhalo" }}
  mysql:  '/usr/halo-19'{{ end }}
  pgtde:  '/usr/pgtde-$v'
  polar:  '/usr/polar-$v'
  oriole: '/usr/oriole-$v'
  pgedge: '/usr/pgedge-$v'
  gpsql:  '/usr/cloudberry'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
{{ .Family.RepoCatalog }}# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Family.CommonPkg 0 }}"
  infra-addons:            "{{ index .Family.CommonPkg 1 }}"
  extra-modules:           "{{ index .Family.CommonPkg 2 }}"
  node-package1:           "{{ getNodePackage1 }}"
  node-package2:           "{{ getNodePackage2 }}"
  node-package3:           "{{ getNodePackage3 }}"
  pgsql-utility:           "{{ getPgsqlUtility }}"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "{{ .Family.PostgreSQLPkg }}"{{ range .Constants.PGSQLKernelMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ mapPkg . }}"{{ end }}

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#{{ range .Constants.PGSQLExoticMap }}{{ $v := mapPkg . }}{{ if and $v (supportsKernelAlias .Key) }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ $v }}"{{ end }}{{ end }}
  java-runtime:            "{{ getJavaRuntime }}"
  kafka-stack:             "kafka kafka-exporter jmx-exporter"
  mysql:                   "{{ .Family.MySQLPkg }}"
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
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ mapPkg . }}"{{ end }}

  #--------------------------------#
  # PGSQL: tools & utils
  #--------------------------------#{{ range .Constants.PGSQLUtilMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ getUtilPkg .Key (mapPkg .) }}"{{ end }}

  #--------------------------------#
  # PGSQL: extensions (default pg{{ .DefaultPG }})
  #--------------------------------#
{{ range $cat := .Categories }}{{ $pkgs := index $.CategoryExts $.DefaultPG $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pgsql-%s:" $cat) }} "{{ replaceAll (printf "%s%d" $.Family.VerSep $.DefaultPG) (printf "%s$v" $.Family.VerSep) $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}

{{ range $pgVer := .PGVersions }}
  #--------------------------------#
  # PG{{ $pgVer }}: packages
  #--------------------------------#{{ range $.Constants.PGSQLKernelMap }}{{ if eq .Key "pgsql" }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) (mapPkg .) }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) (mapPkg .) }}"{{ end }}{{ end }}
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
