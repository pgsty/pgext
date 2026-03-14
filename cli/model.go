/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Extension represents a PostgreSQL extension from pgext.extension table
type Extension struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Pkg         string         `db:"pkg"`
	LeadExt     sql.NullString `db:"lead_ext"`
	Category    sql.NullString `db:"category"`
	State       sql.NullString `db:"state"`
	URL         sql.NullString `db:"url"`
	License     sql.NullString `db:"license"`
	Tags        []string       `db:"tags"`
	Version     sql.NullString `db:"version"`
	Repo        sql.NullString `db:"repo"`
	Lang        sql.NullString `db:"lang"`
	Contrib     bool           `db:"contrib"`
	Lead        bool           `db:"lead"`
	HasBin      bool           `db:"has_bin"`
	HasLib      bool           `db:"has_lib"`
	NeedDDL     bool           `db:"need_ddl"`
	NeedLoad    bool           `db:"need_load"`
	Trusted     sql.NullBool   `db:"trusted"`
	Relocatable sql.NullBool   `db:"relocatable"`
	Schemas     []string       `db:"schemas"`
	PgVer       []string       `db:"pg_ver"`
	Requires    []string       `db:"requires"`
	RequireBy   []string       `db:"require_by"`
	SeeAlso     []string       `db:"see_also"`
	RpmVer      sql.NullString `db:"rpm_ver"`
	RpmRepo     sql.NullString `db:"rpm_repo"`
	RpmPkg      sql.NullString `db:"rpm_pkg"`
	RpmPg       []string       `db:"rpm_pg"`
	RpmDeps     []string       `db:"rpm_deps"`
	DebVer      sql.NullString `db:"deb_ver"`
	DebRepo     sql.NullString `db:"deb_repo"`
	DebPkg      sql.NullString `db:"deb_pkg"`
	DebDeps     []string       `db:"deb_deps"`
	DebPg       []string       `db:"deb_pg"`
	Source      sql.NullString `db:"source"`
	EnDesc      sql.NullString `db:"en_desc"`
	ZhDesc      sql.NullString `db:"zh_desc"`
	Comment     sql.NullString `db:"comment"`
	Extra       JsonMap        `db:"extra"`
}

// JsonMap is a type alias for map[string]interface{} to handle JSONB fields
type JsonMap map[string]interface{}

// IsReady returns true if the extension is not in "not-ready" state
func (e *Extension) IsReady() bool {
	return !e.State.Valid || e.State.String != "not-ready"
}

// GetZhDesc returns the Chinese description, falling back to English description, then extension name
func (e *Extension) GetZhDesc() string {
	if e.ZhDesc.Valid && e.ZhDesc.String != "" {
		return e.ZhDesc.String
	}
	if e.EnDesc.Valid && e.EnDesc.String != "" {
		return e.EnDesc.String
	}
	return e.Name
}

// GetEnDesc returns the English description, falling back to extension name
func (e *Extension) GetEnDesc() string {
	if e.EnDesc.Valid && e.EnDesc.String != "" {
		return e.EnDesc.String
	}
	return e.Name
}

// GetSourceFilename returns the source tarball filename.
func (e *Extension) GetSourceFilename() string {
	if !e.Source.Valid {
		return ""
	}

	source := strings.TrimSpace(e.Source.String)
	if source == "" {
		return ""
	}
	if idx := strings.LastIndex(source, "/"); idx >= 0 {
		return source[idx+1:]
	}
	return source
}

// GetSourceURL returns a site-specific source tarball URL.
func (e *Extension) GetSourceURL(region Region) string {
	if !e.Source.Valid {
		return ""
	}

	source := strings.TrimSpace(e.Source.String)
	if source == "" {
		return ""
	}
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return source
	}

	base := DefaultSourceURL
	if region == RegionChina {
		base = ChinaSourceURL
	}
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(source, "/")
}

// GetLibName returns the library name for shared_preload_libraries
// It checks the Extra field for a "lib" key, otherwise returns the extension Name
func (ext *Extension) GetLibName() string {
	if ext.Extra != nil {
		if lib, ok := ext.Extra["lib"]; ok {
			if libStr, ok := lib.(string); ok && libStr != "" {
				return libStr
			}
		}
	}
	return ext.Name
}

// Category represents a PostgreSQL extension category
type Category struct {
	ID     int            `db:"id"`
	Name   string         `db:"name"`
	Icon1  sql.NullString `db:"icon1"`
	EnDesc sql.NullString `db:"en_desc"`
	ZhDesc sql.NullString `db:"zh_desc"`
}

// PkgInfo represents availability information from pgext.pkg
type PkgInfo struct {
	PG      int            `db:"pg"`
	OS      string         `db:"os"`
	Name    sql.NullString `db:"name"`
	Pkg     string         `db:"pkg"`
	Ext     string         `db:"ext"`
	State   sql.NullString `db:"state"`
	Org     sql.NullString `db:"org"`
	Version sql.NullString `db:"version"`
	Count   sql.NullInt64  `db:"count"`
}

// Binary represents binary package details from pgext.bin
type Binary struct {
	PG         int            `db:"pg"`
	OS         string         `db:"os"`
	Name       string         `db:"name"`
	Repo       string         `db:"repo"`
	Org        sql.NullString `db:"org"`
	Type       string         `db:"type"`
	Version    string         `db:"version"`
	File       string         `db:"file"`
	SHA256     sql.NullString `db:"sha256"`
	Href       string         `db:"href"`
	DefaultURL string         `db:"default_url"`
	MirrorURL  sql.NullString `db:"mirror_url"`
	Size       int64          `db:"size"`
	SizeFull   int64          `db:"size_full"`
}

// PGVersion represents a PostgreSQL major version configuration
type PGVersion struct {
	PG     int  `db:"pg"`
	Active bool `db:"active"`
}

// OSVersion represents an OS configuration
type OSVersion struct {
	OS      string `db:"os"`
	OSFull  string `db:"os_full"`
	Active  bool   `db:"active"`
	OSMajor string `db:"os_major"`
	OSArch  string `db:"os_arch"`
}

// ExtensionCache is a cache for all extensions and their relationships
type ExtensionCache struct {
	// All extensions as a slice
	Extensions []*Extension

	// Map from extension name to extension object
	ExtMap map[string]*Extension

	// Map from package name to lead extension
	PkgMap map[string]*Extension

	// Map from extension name to extensions that depend on it
	ExtDeps map[string][]*Extension

	// All categories
	Categories []*Category

	// Map from category name to category object
	CateMap map[string]*Category

	// Map from category name to extensions in that category
	CateExtMap map[string][]*Extension

	// Active PostgreSQL versions
	PGVersions []int

	// Active OS versions
	OSVersions []OSVersion
}

// ParsePGVersions converts a list of version strings to sorted integer PG major versions
func ParsePGVersions(values []string) []int {
	pgVersions := make([]int, 0, len(values))
	seen := make(map[int]bool)

	for _, val := range values {
		stripped := strings.TrimSpace(val)
		if strings.HasSuffix(stripped, "+") {
			stripped = stripped[:len(stripped)-1]
		}

		var version int
		if _, err := fmt.Sscanf(stripped, "%d", &version); err == nil {
			if !seen[version] {
				pgVersions = append(pgVersions, version)
				seen[version] = true
			}
		}
	}

	return pgVersions
}

// GetAttributeBadge generates attribute badge string like 'cbsLdtr'
func (e *Extension) GetAttributeBadge() string {
	attrs := make([]byte, 0, 8)

	if e.Contrib {
		attrs = append(attrs, 'c')
	} else {
		attrs = append(attrs, '-')
	}

	if e.HasBin {
		attrs = append(attrs, 'b')
	} else {
		attrs = append(attrs, '-')
	}

	if e.HasLib {
		attrs = append(attrs, 's')
	} else {
		attrs = append(attrs, '-')
	}

	if e.NeedLoad {
		attrs = append(attrs, 'L')
	} else {
		attrs = append(attrs, '-')
	}

	if e.NeedDDL {
		attrs = append(attrs, 'd')
	} else {
		attrs = append(attrs, '-')
	}

	if e.Trusted.Valid && e.Trusted.Bool {
		attrs = append(attrs, 't')
	} else {
		attrs = append(attrs, '-')
	}

	if e.Relocatable.Valid && e.Relocatable.Bool {
		attrs = append(attrs, 'r')
	} else {
		attrs = append(attrs, '-')
	}

	return string(attrs)
}

// GetVersion returns the version string, or "-" if not set
func (e *Extension) GetVersion() string {
	if e.Version.Valid {
		return e.Version.String
	}
	return "-"
}

// GetPkgURLLink returns a markdown link to the package URL, or plain backtick-wrapped name
func (e *Extension) GetPkgURLLink() string {
	if e.URL.Valid && e.URL.String != "" {
		return fmt.Sprintf("[`%s`](%s)", e.Pkg, e.URL.String)
	}
	return fmt.Sprintf("`%s`", e.Pkg)
}

// GetPkgPageName returns the extension page slug that should represent this package.
// Lead pages should always point to themselves; non-lead extensions can point to
// the package's lead extension page when provided.
func (e *Extension) GetPkgPageName() string {
	if e.Lead {
		return e.Name
	}
	if e.LeadExt.Valid && e.LeadExt.String != "" {
		return e.LeadExt.String
	}
	return e.Name
}

// InferRepo infers the expected repository (PGDG/PIGSTY) for an extension on a given OS
func InferRepo(ext *Extension, osName string) string {
	if strings.HasPrefix(osName, "el") {
		if ext.RpmRepo.Valid && ext.RpmRepo.String != "" {
			return strings.ToUpper(ext.RpmRepo.String)
		}
	} else {
		if ext.DebRepo.Valid && ext.DebRepo.String != "" {
			return strings.ToUpper(ext.DebRepo.String)
		}
	}
	if ext.Repo.Valid && ext.Repo.String != "" {
		r := strings.ToUpper(ext.Repo.String)
		if r == "MIXED" {
			return "PIGSTY"
		}
		return r
	}
	return "-"
}

// SanitizeText normalizes text for safe markdown table output
func SanitizeText(value string) string {
	if value == "" {
		return ""
	}

	// Replace pipe characters for markdown tables
	sanitized := strings.ReplaceAll(value, "|", "\\|")
	// Replace newlines with spaces
	sanitized = strings.ReplaceAll(sanitized, "\n", " ")
	// Normalize whitespace
	fields := strings.Fields(sanitized)
	return strings.Join(fields, " ")
}

// WriteMarkdownFile writes content to a markdown file, creating parent directories as needed
func WriteMarkdownFile(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(path), err)
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// FormatSize formats byte size to human-readable format
func FormatSize(bytes int64) string {
	if bytes < 0 {
		return "0 B"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	suffixes := []string{"KiB", "MiB", "GiB", "TiB"}
	if exp >= len(suffixes) {
		exp = len(suffixes) - 1
	}

	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), suffixes[exp])
}
