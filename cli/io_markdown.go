/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

IO (pigsty.io) Markdown generators - English-specific badge/table helpers
Reuses CC badge helpers for language-neutral badges (category, language, license, repo, pgver, etc.)
Only overrides helpers that contain Chinese text.
*/
package cli

import (
	"database/sql"
	"fmt"
	"strings"
)

// IOFlagBadge generates a yes/no flag badge with CSS class (English)
func IOFlagBadge(val bool) string {
	if val {
		return `<span class="ext-flag ext-flag--yes">Yes</span>`
	}
	return `<span class="ext-flag ext-flag--no">No</span>`
}

// IOOptFlagBadge generates a yes/no/unknown flag badge (English)
func IOOptFlagBadge(val sql.NullBool) string {
	if !val.Valid {
		return `<span class="ext-flag ext-flag--no">No</span>`
	}
	return IOFlagBadge(val.Bool)
}

// IOExtensionTable generates a markdown extension table with English headers
func IOExtensionTable(exts []*Extension) string {
	var b strings.Builder
	b.WriteString("| **Extension** | **Package** | **Version** | **License** | **Language** | **Description** |\n")
	b.WriteString("|:---------|:-------|:--------:|:----------:|:--------:|:---------|\n")
	for _, ext := range exts {
		version := ext.GetVersion()
		license := "-"
		if ext.License.Valid {
			license = CCLicenseBadge(ext.License.String)
		}
		lang := "-"
		if ext.Lang.Valid {
			lang = CCLanguageBadge(ext.Lang.String)
		}
		desc := SanitizeText(ext.GetEnDesc())
		pkgLink := ext.GetPkgURLLink()
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
			ext.Name, ext.Name, pkgLink, version, license, lang, desc))
	}
	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

// CategoryEnName returns the English display name for a category
func CategoryEnName(catName string) string {
	names := map[string]string{
		"TIME": "Timeseries", "GIS": "Geospatial", "RAG": "Vector", "FTS": "Full-Text Search",
		"OLAP": "Analytics", "FEAT": "Feature", "LANG": "Programming Language", "TYPE": "Data Type",
		"UTIL": "Utility", "FUNC": "Math & Function", "ADMIN": "Administration", "STAT": "Monitoring & Stats",
		"SEC": "Security", "FDW": "Foreign Data", "SIM": "Simulation & Compatibility", "ETL": "Data Integration",
	}
	if n, ok := names[strings.ToUpper(catName)]; ok {
		return n
	}
	return catName
}
