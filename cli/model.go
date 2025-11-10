/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"database/sql"
	"fmt"
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