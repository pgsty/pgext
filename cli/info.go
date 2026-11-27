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

	urlField := "r.default_url"
	if region == "china" || region == "mirror" {
		urlField = "COALESCE(r.mirror_url, r.default_url)"
	}

	query := fmt.Sprintf(`
		SELECT b.name, p.os, r.org, b.size, b.version, format('%%s/%%s', %s, b.href) AS url
		FROM pgext.pkg p
		JOIN pgext.bin b USING (pg, os, name)
		JOIN pgext.repository r ON b.repo = r.id
		WHERE (p.pkg = $1 OR p.ext = $1)
	`, urlField)

	args := []interface{}{name}
	if pgVer > 0 {
		query += fmt.Sprintf(" AND p.pg = $%d", len(args)+1)
		args = append(args, pgVer)
	}
	if osFilter != "" {
		query += fmt.Sprintf(" AND p.os LIKE $%d", len(args)+1)
		args = append(args, "%"+osFilter+"%")
	}
	query += " ORDER BY p.pg DESC, p.os, b.ver::pgext.version DESC"

	rows, err := QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var packages []BinPackage
	for rows.Next() {
		var pkg BinPackage
		var size int64
		if err := rows.Scan(&pkg.Name, &pkg.OS, &pkg.Org, &size, &pkg.Version, &pkg.URL); err != nil {
			logrus.Warnf("scan bin row: %v", err)
			continue
		}
		pkg.Size = FormatSize(size)
		packages = append(packages, pkg)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration failed: %w", err)
	}

	displayBinTable(packages)
	return nil
}

// displayBinTable displays binary packages in a table format.
func displayBinTable(packages []BinPackage) {
	if len(packages) == 0 {
		fmt.Print("\nNo packages found.\n\n")
		return
	}

	wName, wOS, wOrg, wSize, wVer := 4, 2, 3, 4, 7
	for _, pkg := range packages {
		wName = max(wName, len(pkg.Name))
		wOS = max(wOS, len(pkg.OS))
		wOrg = max(wOrg, len(pkg.Org))
		wSize = max(wSize, len(pkg.Size))
		wVer = max(wVer, len(pkg.Version))
	}

	totalWidth := wName + wOS + wOrg + wSize + wVer + 12
	sep := strings.Repeat("─", totalWidth)
	fmt.Println()
	fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n", wName, "NAME", wOS, "OS", wOrg, "ORG", wSize, "SIZE", wVer, "VER", "URL")
	fmt.Println(sep)
	for _, pkg := range packages {
		fmt.Printf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n", wName, pkg.Name, wOS, pkg.OS, wOrg, pkg.Org, wSize, pkg.Size, wVer, pkg.Version, pkg.URL)
	}
	fmt.Println(sep)
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

// ShowExtOptions controls extension output formatting and destination.
type ShowExtOptions struct {
	JSON   bool
	Writer io.Writer
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
	info := &ExtensionInfo{}
	var (
		leadExt, category, state, url, license, version, repo, lang sql.NullString
		rpmVer, rpmRepo, rpmPkg, debVer, debRepo, debPkg            sql.NullString
		source, enDesc, zhDesc, comment, mtime                      sql.NullString
		extraJSON                                                   []byte
	)

	err := QueryRowContext(ctx, extensionInfoQuery, extName).Scan(
		&info.ID, &info.Name, &info.Pkg, &leadExt, &category,
		&state, &url, &license, &info.Tags, &version,
		&repo, &lang, &info.Contrib, &info.Lead, &info.HasBin,
		&info.HasLib, &info.NeedDDL, &info.NeedLoad, &info.Trusted,
		&info.Relocatable, &info.Schemas, &info.PGVer, &info.Requires, &info.RequireBy,
		&info.SeeAlso, &rpmVer, &rpmRepo, &rpmPkg, &info.RPMPG,
		&info.RPMDeps, &debVer, &debRepo, &debPkg, &info.DEBDeps,
		&info.DEBPG, &source, &extraJSON, &enDesc, &zhDesc,
		&comment, &mtime,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("extension %q not found", extName)
	}
	if err != nil {
		return nil, fmt.Errorf("query extension %q: %w", extName, err)
	}

	info.LeadExt, info.Category, info.State = leadExt.String, category.String, state.String
	info.URL, info.License, info.Version = url.String, license.String, version.String
	info.Repo, info.Lang = repo.String, lang.String
	info.RPMVer, info.RPMRepo, info.RPMPkg = rpmVer.String, rpmRepo.String, rpmPkg.String
	info.DEBVer, info.DEBRepo, info.DEBPkg = debVer.String, debRepo.String, debPkg.String
	info.Source, info.EnDesc, info.ZhDesc = source.String, enDesc.String, zhDesc.String
	info.Comment, info.MTime = comment.String, mtime.String

	if len(extraJSON) > 0 {
		if err := json.Unmarshal(extraJSON, &info.Extra); err != nil {
			return nil, fmt.Errorf("decode extension %q metadata: %w", extName, err)
		}
	}
	return info, nil
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

func wrapText(text string, width int) string {
	if len(text) <= width {
		return "  " + text
	}
	var result strings.Builder
	line := "  "
	for _, word := range strings.Fields(text) {
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
