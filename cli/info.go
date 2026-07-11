/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// BinPackage represents a binary package record for display.
type BinPackage struct {
	Name    string
	OS      string
	Org     string
	Size    string
	Version string
	URL     string
}

// ShowBin displays binary packages from pgext.bin table.
func ShowBin(name string, pgVer int, osFilter string, region string) error {
	ctx := context.Background()
	if name == "" {
		return fmt.Errorf("package or extension name is required")
	}

	// Choose URL field based on region
	urlField := "r.default_url"
	if region == "china" || region == "mirror" {
		urlField = "COALESCE(r.mirror_url, r.default_url)"
	}

	// Build query
	query := fmt.Sprintf(`
		SELECT b.name, p.os, r.org, b.size, b.version, format('%%s/%%s', %s, b.href) AS url
		FROM pgext.pkg p
		JOIN pgext.bin b USING (pg, os, name)
		JOIN pgext.repository r ON b.repo = r.id
		WHERE (p.pkg = $1 OR p.ext = $1)
	`, urlField)

	args := []interface{}{name}
	argCount := 2

	if pgVer > 0 {
		query += fmt.Sprintf(" AND p.pg = $%d", argCount)
		args = append(args, pgVer)
		argCount++
	}
	if osFilter != "" {
		query += fmt.Sprintf(" AND p.os LIKE $%d", argCount)
		args = append(args, "%"+osFilter+"%")
		argCount++
	}

	query += " ORDER BY p.pg DESC, p.os, b.ver::pgext.version DESC"

	// Execute query
	rows, err := QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Collect results
	var packages []BinPackage
	for rows.Next() {
		var pkg BinPackage
		var size int64
		err := rows.Scan(&pkg.Name, &pkg.OS, &pkg.Org, &size, &pkg.Version, &pkg.URL)
		if err != nil {
			logrus.Debugf("scan error: %v", err)
			continue
		}
		pkg.Size = formatSize(size)
		packages = append(packages, pkg)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration failed: %w", err)
	}

	// Display results
	displayBinTable(packages)

	return nil
}

// formatSize formats byte size to human-readable format with 2 significant digits.
func formatSize(bytes int64) string {
	if bytes < 0 {
		return "0"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	units := []string{"KiB", "MiB", "GiB", "TiB"}
	size := float64(bytes) / unit
	unitIdx := 0

	for size >= unit && unitIdx < len(units)-1 {
		size /= unit
		unitIdx++
	}

	// Format: 100+ = integer, 1-99.9 = one decimal
	if size >= 100 {
		return fmt.Sprintf("%.0f %s", size, units[unitIdx])
	}
	return fmt.Sprintf("%.1f %s", size, units[unitIdx])
}

// displayBinTable displays binary packages in a table format.
func displayBinTable(packages []BinPackage) {
	if len(packages) == 0 {
		fmt.Println()
		fmt.Println("No packages found.")
		fmt.Println()
		return
	}

	// Calculate column widths
	widths := map[string]int{
		"name":    4,
		"os":      2,
		"org":     3,
		"size":    4,
		"version": 7,
	}

	for _, pkg := range packages {
		if len(pkg.Name) > widths["name"] {
			widths["name"] = len(pkg.Name)
		}
		if len(pkg.OS) > widths["os"] {
			widths["os"] = len(pkg.OS)
		}
		if len(pkg.Org) > widths["org"] {
			widths["org"] = len(pkg.Org)
		}
		if len(pkg.Size) > widths["size"] {
			widths["size"] = len(pkg.Size)
		}
		if len(pkg.Version) > widths["version"] {
			widths["version"] = len(pkg.Version)
		}
	}

	// Print header
	fmt.Println()
	fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n",
		widths["name"], "NAME",
		widths["os"], "OS",
		widths["org"], "ORG",
		widths["size"], "SIZE",
		widths["version"], "VER", "URL")

	// Print separator
	totalWidth := widths["name"] + widths["os"] + widths["org"] + widths["size"] + widths["version"] + 12
	fmt.Println(strings.Repeat("─", totalWidth))

	// Print rows
	for _, pkg := range packages {
		fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n",
			widths["name"], pkg.Name,
			widths["os"], pkg.OS,
			widths["org"], pkg.Org,
			widths["size"], pkg.Size,
			widths["version"], pkg.Version,
			pkg.URL)
	}

	// Print footer
	fmt.Println(strings.Repeat("─", totalWidth))
	fmt.Printf("Total: %d packages\n\n", len(packages))
}

const extensionInfoQuery = `
	SELECT id, name, pkg, lead_ext, category, state, url, license,
	       tags, version, repo, lang, contrib, lead, has_bin, has_lib,
	       need_ddl, need_load, trusted, relocatable, schemas, pg_ver,
	       requires, require_by, see_also, rpm_ver, rpm_repo, rpm_pkg,
	       rpm_pg, rpm_deps, deb_ver, deb_repo, deb_pkg, deb_deps,
	       deb_pg, source, extra, en_desc, zh_desc, comment, mtime::text
	FROM pgext.extension
	WHERE name = $1
`

// ExtensionInfo is the stable JSON representation returned by pgext ext --json.
// Its field names deliberately follow the pgext.extension catalog vocabulary.
type ExtensionInfo struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Pkg         string   `json:"pkg"`
	LeadExt     string   `json:"lead_ext"`
	Category    string   `json:"category"`
	State       string   `json:"state"`
	URL         string   `json:"url"`
	License     string   `json:"license"`
	Tags        []string `json:"tags"`
	Version     string   `json:"version"`
	Repo        string   `json:"repo"`
	Lang        string   `json:"lang"`
	Contrib     *bool    `json:"contrib"`
	Lead        *bool    `json:"lead"`
	HasBin      *bool    `json:"has_bin"`
	HasLib      *bool    `json:"has_lib"`
	NeedDDL     *bool    `json:"need_ddl"`
	NeedLoad    *bool    `json:"need_load"`
	Trusted     *bool    `json:"trusted"`
	Relocatable *bool    `json:"relocatable"`
	Schemas     []string `json:"schemas"`
	PGVer       []string `json:"pg_ver"`
	Requires    []string `json:"requires"`
	RequireBy   []string `json:"require_by"`
	SeeAlso     []string `json:"see_also"`
	RPMVer      string   `json:"rpm_ver"`
	RPMRepo     string   `json:"rpm_repo"`
	RPMPkg      string   `json:"rpm_pkg"`
	RPMPG       []string `json:"rpm_pg"`
	RPMDeps     []string `json:"rpm_deps"`
	DEBVer      string   `json:"deb_ver"`
	DEBRepo     string   `json:"deb_repo"`
	DEBPkg      string   `json:"deb_pkg"`
	DEBDeps     []string `json:"deb_deps"`
	DEBPG       []string `json:"deb_pg"`
	Source      string   `json:"source"`
	Extra       JsonMap  `json:"extra"`
	EnDesc      string   `json:"en_desc"`
	ZhDesc      string   `json:"zh_desc"`
	Comment     string   `json:"comment"`
	MTime       string   `json:"mtime"`
}

// ShowExtOptions controls extension output without changing the legacy ShowExt API.
type ShowExtOptions struct {
	JSON   bool
	Writer io.Writer
}

// ShowExt displays extension information using the legacy text format.
func ShowExt(extName string) error {
	return ShowExtWithOptions(context.Background(), extName, ShowExtOptions{Writer: os.Stdout})
}

// ShowExtWithOptions displays extension information from pgext.extension.
func ShowExtWithOptions(ctx context.Context, extName string, opts ShowExtOptions) error {
	info, err := LoadExtensionInfo(ctx, extName)
	if err != nil {
		return err
	}

	out := opts.Writer
	if out == nil {
		out = os.Stdout
	}
	if opts.JSON {
		if err := writeExtensionInfoJSON(out, info); err != nil {
			return fmt.Errorf("encode extension %q as JSON: %w", extName, err)
		}
		return nil
	}

	if err := writeExtensionInfo(out, info); err != nil {
		return fmt.Errorf("write extension %q: %w", extName, err)
	}
	return nil
}

// LoadExtensionInfo loads one record from the standard packaged extension catalog.
func LoadExtensionInfo(ctx context.Context, extName string) (*ExtensionInfo, error) {
	ext := &Extension{}
	var tags, schemas, pgVer, requires, requireBy, seeAlso interface{}
	var rpmPG, rpmDeps, debPG, debDeps interface{}
	var contrib, lead, hasBin, hasLib, needDDL, needLoad sql.NullBool
	var extraJSON []byte
	var mtime sql.NullString

	err := QueryRowContext(ctx, extensionInfoQuery, extName).Scan(
		&ext.ID, &ext.Name, &ext.Pkg, &ext.LeadExt, &ext.Category,
		&ext.State, &ext.URL, &ext.License, &tags, &ext.Version,
		&ext.Repo, &ext.Lang, &contrib, &lead, &hasBin,
		&hasLib, &needDDL, &needLoad, &ext.Trusted,
		&ext.Relocatable, &schemas, &pgVer, &requires, &requireBy,
		&seeAlso, &ext.RpmVer, &ext.RpmRepo, &ext.RpmPkg, &rpmPG,
		&rpmDeps, &ext.DebVer, &ext.DebRepo, &ext.DebPkg, &debDeps,
		&debPG, &ext.Source, &extraJSON, &ext.EnDesc, &ext.ZhDesc,
		&ext.Comment, &mtime,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("extension %q not found", extName)
	}
	if err != nil {
		return nil, fmt.Errorf("query extension %q: %w", extName, err)
	}

	ext.Tags = nullableInfoArray(tags)
	ext.Schemas = nullableInfoArray(schemas)
	ext.PgVer = nullableInfoArray(pgVer)
	ext.Requires = nullableInfoArray(requires)
	ext.RequireBy = nullableInfoArray(requireBy)
	ext.SeeAlso = nullableInfoArray(seeAlso)
	ext.RpmPg = nullableInfoArray(rpmPG)
	ext.RpmDeps = nullableInfoArray(rpmDeps)
	ext.DebPg = nullableInfoArray(debPG)
	ext.DebDeps = nullableInfoArray(debDeps)
	if len(extraJSON) > 0 {
		if err := json.Unmarshal(extraJSON, &ext.Extra); err != nil {
			return nil, fmt.Errorf("decode extension %q metadata: %w", extName, err)
		}
	}

	info := &ExtensionInfo{
		ID: ext.ID, Name: ext.Name, Pkg: ext.Pkg,
		LeadExt: ext.LeadExt.String, Category: ext.Category.String,
		State: ext.State.String, URL: ext.URL.String, License: ext.License.String,
		Tags: ext.Tags, Version: ext.Version.String, Repo: ext.Repo.String,
		Lang: ext.Lang.String, Contrib: nullBoolPointer(contrib), Lead: nullBoolPointer(lead),
		HasBin: nullBoolPointer(hasBin), HasLib: nullBoolPointer(hasLib), NeedDDL: nullBoolPointer(needDDL),
		NeedLoad: nullBoolPointer(needLoad), Schemas: ext.Schemas, PGVer: ext.PgVer,
		Requires: ext.Requires, RequireBy: ext.RequireBy, SeeAlso: ext.SeeAlso,
		RPMVer: ext.RpmVer.String, RPMRepo: ext.RpmRepo.String,
		RPMPkg: ext.RpmPkg.String, RPMPG: ext.RpmPg, RPMDeps: ext.RpmDeps,
		DEBVer: ext.DebVer.String, DEBRepo: ext.DebRepo.String,
		DEBPkg: ext.DebPkg.String, DEBDeps: ext.DebDeps, DEBPG: ext.DebPg,
		Source: ext.Source.String, Extra: ext.Extra, EnDesc: ext.EnDesc.String,
		ZhDesc: ext.ZhDesc.String, Comment: ext.Comment.String, MTime: mtime.String,
	}
	if ext.Trusted.Valid {
		value := ext.Trusted.Bool
		info.Trusted = &value
	}
	if ext.Relocatable.Valid {
		value := ext.Relocatable.Bool
		info.Relocatable = &value
	}
	return info, nil
}

func nullBoolPointer(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}
	result := value.Bool
	return &result
}

func nullableInfoArray(value interface{}) []string {
	if value == nil {
		return nil
	}
	return parseArrayValue(value)
}

func writeExtensionInfoJSON(out io.Writer, info *ExtensionInfo) error {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(info)
}

func writeExtensionInfo(out io.Writer, info *ExtensionInfo) error {
	var text strings.Builder
	fmt.Fprintln(&text, "\n╔═══════════════════════════════════════════╗")
	fmt.Fprintln(&text, "║        Extension Information              ║")
	fmt.Fprintln(&text, "╚═══════════════════════════════════════════╝")
	fmt.Fprintln(&text)
	fmt.Fprintf(&text, "Name:        %s\n", info.Name)
	fmt.Fprintf(&text, "Package:     %s\n", info.Pkg)
	fmt.Fprintf(&text, "Category:    %s\n", info.Category)
	fmt.Fprintf(&text, "Repository:  %s\n", info.Repo)
	fmt.Fprintf(&text, "Language:    %s\n", info.Lang)
	fmt.Fprintf(&text, "License:     %s\n", info.License)
	fmt.Fprintf(&text, "Version:     %s\n", info.Version)
	description := info.EnDesc
	if description == "" {
		description = info.ZhDesc
	}
	fmt.Fprintf(&text, "\nDescription:\n%s\n", wrapText(description, 60))
	if len(info.Requires) > 0 {
		fmt.Fprintf(&text, "\nRequires:    %s\n", strings.Join(info.Requires, ", "))
	}
	if info.URL != "" {
		fmt.Fprintf(&text, "Website:     %s\n", info.URL)
	}
	if info.Source != "" {
		fmt.Fprintf(&text, "Source:      %s\n", info.Source)
	}
	fmt.Fprintln(&text)
	_, err := io.WriteString(out, text.String())
	return err
}

// wrapText wraps text to specified width.
func wrapText(text string, width int) string {
	if len(text) <= width {
		return "  " + text
	}

	var result strings.Builder
	words := strings.Fields(text)
	line := "  "

	for _, word := range words {
		if len(line)+len(word)+1 > width {
			result.WriteString(line)
			result.WriteString("\n")
			line = "  " + word
		} else {
			if line != "  " {
				line += " "
			}
			line += word
		}
	}

	if line != "  " {
		result.WriteString(line)
	}

	return result.String()
}
