/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// LoadExtensionCache loads all extension data into memory cache.
func LoadExtensionCache(ctx context.Context) (*ExtensionCache, error) {
	cache := &ExtensionCache{
		ExtMap:     make(map[string]*Extension),
		PkgMap:     make(map[string]*Extension),
		CateMap:    make(map[string]*Category),
		CateExtMap: make(map[string][]*Extension),
	}

	var err error
	if cache.PGVersions, err = ActivePGVersions(ctx); err != nil {
		return nil, fmt.Errorf("failed to load PG versions: %w", err)
	}
	if err := loadOSVersions(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load OS versions: %w", err)
	}
	if err := loadCategories(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load categories: %w", err)
	}
	if err := loadExtensions(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load extensions: %w", err)
	}
	return cache, nil
}

func loadOSVersions(ctx context.Context, cache *ExtensionCache) error {
	rows, err := QueryContext(ctx, `
		SELECT os, os_full, active, os_major, os_arch
		FROM pgext.os
		WHERE active
		ORDER BY os_major, os_arch DESC
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var os OSVersion
		if err := rows.Scan(&os.OS, &os.OSFull, &os.Active, &os.OSMajor, &os.OSArch); err != nil {
			return err
		}
		cache.OSVersions = append(cache.OSVersions, os)
	}
	return rows.Err()
}

func loadCategories(ctx context.Context, cache *ExtensionCache) error {
	rows, err := QueryContext(ctx, `SELECT id, name, icon1, en_desc, zh_desc FROM pgext.category ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cat := &Category{}
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Icon1, &cat.EnDesc, &cat.ZhDesc); err != nil {
			return err
		}
		cache.Categories = append(cache.Categories, cat)
		cache.CateMap[strings.ToUpper(cat.Name)] = cat
	}
	return rows.Err()
}

func loadExtensions(ctx context.Context, cache *ExtensionCache) error {
	const query = `
		SELECT id, name, pkg, lead_ext, category, state, url, license,
		       tags, version, repo, lang, contrib, lead, has_bin, has_lib,
		       need_ddl, need_load, trusted, relocatable, schemas, pg_ver,
		       requires, require_by, see_also, rpm_ver, rpm_repo, rpm_pkg,
		       rpm_pg, rpm_deps, deb_ver, deb_repo, deb_pkg, deb_deps,
		       deb_pg, source, en_desc, zh_desc, comment, extra
		FROM pgext.extension
		ORDER BY id
	`
	rows, err := QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		ext := &Extension{}
		var extraJSON []byte

		err := rows.Scan(
			&ext.ID, &ext.Name, &ext.Pkg, &ext.LeadExt, &ext.Category,
			&ext.State, &ext.URL, &ext.License, &ext.Tags, &ext.Version,
			&ext.Repo, &ext.Lang, &ext.Contrib, &ext.Lead, &ext.HasBin,
			&ext.HasLib, &ext.NeedDDL, &ext.NeedLoad, &ext.Trusted,
			&ext.Relocatable, &ext.Schemas, &ext.PgVer, &ext.Requires, &ext.RequireBy,
			&ext.SeeAlso, &ext.RpmVer, &ext.RpmRepo, &ext.RpmPkg, &ext.RpmPg,
			&ext.RpmDeps, &ext.DebVer, &ext.DebRepo, &ext.DebPkg, &ext.DebDeps,
			&ext.DebPg, &ext.Source, &ext.EnDesc, &ext.ZhDesc, &ext.Comment, &extraJSON,
		)
		if err != nil {
			return err
		}
		if len(extraJSON) > 0 {
			var m map[string]interface{}
			if err := json.Unmarshal(extraJSON, &m); err == nil {
				ext.Extra = JsonMap(m)
			}
		}

		cache.Extensions = append(cache.Extensions, ext)
		cache.ExtMap[ext.Name] = ext
		if ext.Lead {
			cache.PkgMap[ext.Pkg] = ext
		}
		if ext.Category.Valid {
			catKey := strings.ToUpper(ext.Category.String)
			cache.CateExtMap[catKey] = append(cache.CateExtMap[catKey], ext)
		}
	}

	// Fill in package map for non-lead extensions whose package had no lead row.
	for _, ext := range cache.Extensions {
		if _, exists := cache.PkgMap[ext.Pkg]; !exists {
			cache.PkgMap[ext.Pkg] = ext
		}
	}
	return rows.Err()
}

// ReadyExtensions returns all extensions that are not in "not-ready" state.
func (cache *ExtensionCache) ReadyExtensions() []*Extension {
	var exts []*Extension
	for _, ext := range cache.Extensions {
		if ext.IsReady() {
			exts = append(exts, ext)
		}
	}
	return exts
}

// GetSiblingExtensions returns other extensions that share the same package.
func (cache *ExtensionCache) GetSiblingExtensions(pkgName, currentExtName string) []*Extension {
	var siblings []*Extension
	for _, ext := range cache.Extensions {
		if ext.Pkg == pkgName && ext.Name != currentExtName {
			siblings = append(siblings, ext)
		}
	}
	return siblings
}

// LoadPackages loads package availability data for a specific package.
func LoadPackages(ctx context.Context, pkgName string) ([]*PkgInfo, error) {
	rows, err := QueryContext(ctx, `
		SELECT pg, os, name, pkg, ext, state, org, version, count
		FROM pgext.pkg
		WHERE pkg = $1
		ORDER BY pg DESC, os
	`, pkgName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*PkgInfo
	for rows.Next() {
		pkg := &PkgInfo{}
		if err := rows.Scan(&pkg.PG, &pkg.OS, &pkg.Name, &pkg.Pkg,
			&pkg.Ext, &pkg.State, &pkg.Org, &pkg.Version, &pkg.Count); err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}
	return packages, rows.Err()
}

// LoadBinaries loads binary package details for a specific extension with deterministic ordering.
// URL prefixes are resolved at render time so callers can choose default vs mirror base URLs.
func LoadBinaries(ctx context.Context, extName string) ([]*Binary, error) {
	rows, err := QueryContext(ctx, `
		SELECT b.pg, b.os, b.name, b.version, r.org, b.size, b.file,
		       b.href, r.default_url, r.mirror_url
		FROM pgext.pkg p
		     JOIN pgext.bin b USING (pg, os, name)
		     JOIN pgext.repository r ON b.repo = r.id
		     JOIN pgext.os o ON p.os = o.os
		WHERE p.ext = $1
		ORDER BY b.pg DESC, o.os_major, o.os_arch DESC,
		         b.ver::pgext.VERSION USING OPERATOR (pgext.>),
		         r.org DESC, b.file DESC, b.repo DESC, b.href DESC
	`, extName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var binaries []*Binary
	for rows.Next() {
		bin := &Binary{}
		if err := rows.Scan(&bin.PG, &bin.OS, &bin.Name, &bin.Version,
			&bin.Org, &bin.Size, &bin.File, &bin.Href, &bin.DefaultURL, &bin.MirrorURL); err != nil {
			return nil, err
		}
		binaries = append(binaries, bin)
	}
	return binaries, rows.Err()
}
