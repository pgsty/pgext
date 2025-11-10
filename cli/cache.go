/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"fmt"
	"strings"
)

// parseArrayValue converts database array values to string slices
// Compatible with both pgx and lib/pq array representations
func parseArrayValue(v interface{}) []string {
	if v == nil {
		return []string{}
	}

	// Handle different array types from pgx
	switch arr := v.(type) {
	case []string:
		return arr
	case []interface{}:
		result := make([]string, 0, len(arr))
		for _, item := range arr {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	case string:
		// Handle PostgreSQL array literal format {val1,val2,val3}
		if arr == "" || arr == "{}" {
			return []string{}
		}
		if strings.HasPrefix(arr, "{") && strings.HasSuffix(arr, "}") {
			arr = arr[1 : len(arr)-1]
			if arr == "" {
				return []string{}
			}
			// Simple split for non-quoted values
			// TODO: handle quoted values with commas inside
			return strings.Split(arr, ",")
		}
		return []string{arr}
	default:
		return []string{}
	}
}

// LoadExtensionCache loads all extension data into memory cache
func LoadExtensionCache(ctx context.Context) (*ExtensionCache, error) {
	cache := &ExtensionCache{
		ExtMap:     make(map[string]*Extension),
		PkgMap:     make(map[string]*Extension),
		ExtDeps:    make(map[string][]*Extension),
		CateMap:    make(map[string]*Category),
		CateExtMap: make(map[string][]*Extension),
	}

	// Load PostgreSQL versions
	if err := loadPGVersions(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load PG versions: %w", err)
	}

	// Load OS versions
	if err := loadOSVersions(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load OS versions: %w", err)
	}

	// Load categories
	if err := loadCategories(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load categories: %w", err)
	}

	// Load extensions
	if err := loadExtensions(ctx, cache); err != nil {
		return nil, fmt.Errorf("failed to load extensions: %w", err)
	}

	// Build dependency map
	buildDependencyMap(cache)

	return cache, nil
}

func loadPGVersions(ctx context.Context, cache *ExtensionCache) error {
	query := `SELECT pg FROM pgext.active_pg ORDER BY pg DESC`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var pg int
		if err := rows.Scan(&pg); err != nil {
			return err
		}
		cache.PGVersions = append(cache.PGVersions, pg)
	}

	return rows.Err()
}

func loadOSVersions(ctx context.Context, cache *ExtensionCache) error {
	query := `
		SELECT os, os_full, active, os_major, os_arch
		FROM pgext.os
		WHERE active
		ORDER BY os_major, os_arch DESC
	`

	rows, err := QueryContext(ctx, query)
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
	query := `SELECT id, name, en_desc, zh_desc FROM pgext.category ORDER BY id`

	rows, err := QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cat := &Category{}
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.EnDesc, &cat.ZhDesc); err != nil {
			return err
		}
		cache.Categories = append(cache.Categories, cat)
		cache.CateMap[strings.ToUpper(cat.Name)] = cat
	}

	return rows.Err()
}

func loadExtensions(ctx context.Context, cache *ExtensionCache) error {
	query := `
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

		// Temporary variables for arrays - use interface{} for pgx compatibility
		var tags, schemas, pgVer, requires, requireBy, seeAlso interface{}
		var rpmPg, rpmDeps, debPg, debDeps interface{}
		var extra interface{}

		err := rows.Scan(
			&ext.ID, &ext.Name, &ext.Pkg, &ext.LeadExt, &ext.Category,
			&ext.State, &ext.URL, &ext.License, &tags, &ext.Version,
			&ext.Repo, &ext.Lang, &ext.Contrib, &ext.Lead, &ext.HasBin,
			&ext.HasLib, &ext.NeedDDL, &ext.NeedLoad, &ext.Trusted,
			&ext.Relocatable, &schemas, &pgVer, &requires, &requireBy,
			&seeAlso, &ext.RpmVer, &ext.RpmRepo, &ext.RpmPkg, &rpmPg,
			&rpmDeps, &ext.DebVer, &ext.DebRepo, &ext.DebPkg, &debDeps,
			&debPg, &ext.Source, &ext.EnDesc, &ext.ZhDesc, &ext.Comment, &extra,
		)
		if err != nil {
			return err
		}

		// Convert arrays - handle pgx array types
		ext.Tags = parseArrayValue(tags)
		ext.Schemas = parseArrayValue(schemas)
		ext.PgVer = parseArrayValue(pgVer)
		ext.Requires = parseArrayValue(requires)
		ext.RequireBy = parseArrayValue(requireBy)
		ext.SeeAlso = parseArrayValue(seeAlso)
		ext.RpmPg = parseArrayValue(rpmPg)
		ext.RpmDeps = parseArrayValue(rpmDeps)
		ext.DebPg = parseArrayValue(debPg)
		ext.DebDeps = parseArrayValue(debDeps)

		// Parse extra JSONB field
		if extra != nil {
			if jsonMap, ok := extra.(map[string]interface{}); ok {
				ext.Extra = JsonMap(jsonMap)
			}
		}

		// Add to cache
		cache.Extensions = append(cache.Extensions, ext)
		cache.ExtMap[ext.Name] = ext

		// Add to package map if lead
		if ext.Lead {
			cache.PkgMap[ext.Pkg] = ext
		}

		// Add to category map
		if ext.Category.Valid {
			catKey := strings.ToUpper(ext.Category.String)
			cache.CateExtMap[catKey] = append(cache.CateExtMap[catKey], ext)
		}
	}

	// Second pass: fill in package map for non-lead extensions
	for _, ext := range cache.Extensions {
		if _, exists := cache.PkgMap[ext.Pkg]; !exists {
			cache.PkgMap[ext.Pkg] = ext
		}
	}

	return rows.Err()
}

func buildDependencyMap(cache *ExtensionCache) {
	// Build reverse dependency map
	for _, ext := range cache.Extensions {
		for _, reqName := range ext.Requires {
			if reqExt, exists := cache.ExtMap[reqName]; exists {
				cache.ExtDeps[reqName] = append(cache.ExtDeps[reqName], reqExt)
			}
		}
	}
}

// GetSiblingExtensions returns other extensions that share the same package
func (cache *ExtensionCache) GetSiblingExtensions(pkgName, currentExtName string) []*Extension {
	var siblings []*Extension

	for _, ext := range cache.Extensions {
		if ext.Pkg == pkgName && ext.Name != currentExtName {
			siblings = append(siblings, ext)
		}
	}

	return siblings
}

// LoadPackages loads package availability data for a specific package
func LoadPackages(ctx context.Context, pkgName string) ([]*PkgInfo, error) {
	query := `
		SELECT pg, os, name, pkg, ext, state, org, version, count
		FROM pgext.pkg
		WHERE pkg = $1
		ORDER BY pg DESC, os
	`

	rows, err := QueryContext(ctx, query, pkgName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*PkgInfo
	for rows.Next() {
		pkg := &PkgInfo{}
		err := rows.Scan(&pkg.PG, &pkg.OS, &pkg.Name, &pkg.Pkg,
			&pkg.Ext, &pkg.State, &pkg.Org, &pkg.Version, &pkg.Count)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	return packages, rows.Err()
}

// LoadBinaries loads binary package details for a specific extension with proper ordering
func LoadBinaries(ctx context.Context, extName string) ([]*Binary, error) {
	// Query binaries with proper ordering: pg DESC, os_major, version DESC, org DESC
	query := `
		SELECT b.pg,b.os,b.name,b.version,r.org,b.size,b.file,
		       format('%s/%s', r.default_url, b.href) AS url
		FROM pgext.pkg p
			 JOIN pgext.bin b USING (pg, os, name)
			 JOIN pgext.repository r ON b.repo = r.id
			 JOIN pgext.os o ON p.os = o.os
		WHERE p.ext = $1
		ORDER BY b.pg DESC, o.os_major, o.os_arch DESC, b.version::pgext.VERSION DESC, r.org DESC
	`

	rows, err := QueryContext(ctx, query, extName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var binaries []*Binary
	for rows.Next() {
		bin := &Binary{}
		var url string
		err := rows.Scan(&bin.PG, &bin.OS, &bin.Name, &bin.Version,
			&bin.Org, &bin.Size, &bin.File, &url)
		if err != nil {
			return nil, err
		}
		// Set the URL in the Href field for compatibility
		bin.Href = url
		binaries = append(binaries, bin)
	}

	return binaries, rows.Err()
}
