/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

Shared OINK documentation generators and presentation helpers for pigsty.cc (ZH)
and pigsty.io (EN), plus common catalog-generation helpers.
*/
package cli

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

// CCBaseURL is the base URL path for extension pages on pigsty.cc and pigsty.io.
const CCBaseURL = "/ext"

// DocsLocale selects the language and mirror region for OINK documentation output.
type DocsLocale string

const (
	DocsZH DocsLocale = "zh"
	DocsEN DocsLocale = "en"
)

func (l DocsLocale) isZH() bool { return l == DocsZH }

func (l DocsLocale) t(zh, en string) string {
	if l == DocsZH {
		return zh
	}
	return en
}

func (l DocsLocale) region() Region {
	if l == DocsZH {
		return RegionChina
	}
	return RegionDefault
}

func (l DocsLocale) extDesc(ext *Extension) string {
	if l == DocsZH {
		return ext.GetZhDesc()
	}
	return ext.GetEnDesc()
}

func (l DocsLocale) catDesc(cat *Category) string {
	if l == DocsZH {
		if cat.ZhDesc.Valid && cat.ZhDesc.String != "" {
			return cat.ZhDesc.String
		}
	} else {
		if cat.EnDesc.Valid && cat.EnDesc.String != "" {
			return cat.EnDesc.String
		}
	}
	return cat.Name
}

var categoryDisplayNames = map[string][2]string{
	"TIME":  {"时序", "Timeseries"},
	"GIS":   {"地理", "Geospatial"},
	"RAG":   {"向量", "Vector"},
	"FTS":   {"全文检索", "Full-Text Search"},
	"OLAP":  {"数据分析", "Analytics"},
	"FEAT":  {"特性", "Feature"},
	"LANG":  {"编程语言", "Programming Language"},
	"TYPE":  {"数据类型", "Data Type"},
	"UTIL":  {"工具", "Utility"},
	"FUNC":  {"数学函数", "Math & Function"},
	"ADMIN": {"管理", "Administration"},
	"STAT":  {"监控统计", "Monitoring & Stats"},
	"SEC":   {"安全", "Security"},
	"FDW":   {"外部数据", "Foreign Data"},
	"SIM":   {"兼容模拟", "Simulation & Compatibility"},
	"ETL":   {"数据集成", "Data Integration"},
}

func (l DocsLocale) catName(catName string) string {
	if pair, ok := categoryDisplayNames[strings.ToUpper(catName)]; ok {
		return l.t(pair[0], pair[1])
	}
	return catName
}

// CategoryZhName returns the Chinese display name for a category.
func CategoryZhName(catName string) string { return DocsZH.catName(catName) }

// CategoryEnName returns the English display name for a category.
func CategoryEnName(catName string) string { return DocsEN.catName(catName) }

// CCCategoryBadge generates a category badge with CSS class.
func CCCategoryBadge(category string) string {
	if category == "" {
		return "-"
	}
	lower := strings.ToLower(category)
	return fmt.Sprintf(`<a class="ext-badge ext-badge--cate %s" href="%s/cate/%s">%s</a>`,
		lower, CCBaseURL, lower, strings.ToUpper(category))
}

// CCLanguageBadge generates a language badge with CSS class.
func CCLanguageBadge(lang string) string {
	if lang == "" {
		return "-"
	}
	anchor := strings.ToLower(lang)
	if anchor == "c++" {
		anchor = "cpp"
	}
	return fmt.Sprintf(`<a class="ext-badge ext-badge--lang %s" href="%s/language#%s">%s</a>`,
		anchor, CCBaseURL, anchor, lang)
}

// CCLicenseBadge generates a license badge with CSS class.
func CCLicenseBadge(license string) string {
	if license == "" {
		return "-"
	}
	cls := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(license, ".", ""), "-", ""))
	anchor := strings.ReplaceAll(cls, " ", "")
	return fmt.Sprintf(`<a class="ext-badge ext-badge--license %s" href="%s/license#%s">%s</a>`,
		cls, CCBaseURL, anchor, license)
}

// CCRepoBadge generates a repo badge (PGDG or PIGSTY) with CSS class.
func CCRepoBadge(repo string) string {
	if repo == "" {
		return "-"
	}
	repoLower := strings.ToLower(repo)
	return fmt.Sprintf(`<a class="ext-badge ext-badge--repo %s" href="%s/repo#%s">%s</a>`,
		repoLower, CCBaseURL, repoLower, strings.ToUpper(repo))
}

// CCPGVersShortcode generates a {{< pgvers "..." >}} shortcode call from supported version strings.
func CCPGVersShortcode(supported []string) string {
	if len(supported) == 0 {
		return "-"
	}
	cleaned := make([]string, 0, len(supported))
	for _, v := range supported {
		v = strings.TrimSuffix(strings.TrimSpace(v), "+")
		if v != "" {
			cleaned = append(cleaned, v)
		}
	}
	if len(cleaned) == 0 {
		return "-"
	}
	return fmt.Sprintf(`{{< pgvers "%s" >}}`, strings.Join(cleaned, ","))
}

// CCPGVersAllShortcode generates a {{< pgvers "..." >}} shortcode with all active PG versions.
func CCPGVersAllShortcode(pgVersions []int) string {
	vers := make([]string, len(pgVersions))
	for i, pg := range pgVersions {
		vers[i] = fmt.Sprintf("%d", pg)
	}
	return fmt.Sprintf(`{{< pgvers "%s" >}}`, strings.Join(vers, ","))
}

// DocsFlagBadge generates a localized yes/no flag badge with CSS class.
func DocsFlagBadge(locale DocsLocale, val bool) string {
	if val {
		return fmt.Sprintf(`<span class="ext-flag ext-flag--yes">%s</span>`, locale.t("是", "Yes"))
	}
	return fmt.Sprintf(`<span class="ext-flag ext-flag--no">%s</span>`, locale.t("否", "No"))
}

// DocsOptFlagBadge generates a localized yes/no flag badge for nullable booleans.
func DocsOptFlagBadge(locale DocsLocale, val sql.NullBool) string {
	return DocsFlagBadge(locale, val.Valid && val.Bool)
}

// CCFlagBadge / IOFlagBadge / CCOptFlagBadge / IOOptFlagBadge preserve exported helpers.
func CCFlagBadge(val bool) string            { return DocsFlagBadge(DocsZH, val) }
func IOFlagBadge(val bool) string            { return DocsFlagBadge(DocsEN, val) }
func CCOptFlagBadge(val sql.NullBool) string { return DocsOptFlagBadge(DocsZH, val) }
func IOOptFlagBadge(val sql.NullBool) string { return DocsOptFlagBadge(DocsEN, val) }

// CCAvailBadge generates an availability badge with CSS class.
func CCAvailBadge(version string) string {
	return fmt.Sprintf(`<span class="ext-badge ext-badge--avail">%s</span>`, version)
}

// CCMissBadge generates a missing badge.
func CCMissBadge() string {
	return `<span class="ext-badge ext-badge--miss">✗</span>`
}

// CCPkgStateBadge generates an availability badge based on package state and version.
func CCPkgStateBadge(state, version string) string {
	switch state {
	case "AVAIL":
		if version != "" {
			return CCAvailBadge(version)
		}
		return CCAvailBadge("✓")
	case "MISS":
		return CCMissBadge()
	default:
		return `<span class="ext-badge ext-badge--na">N/A</span>`
	}
}

// CCExtLink generates a plain markdown extension link.
func CCExtLink(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("[`%s`](%s/e/%s)", name, CCBaseURL, name)
}

// CCExtRef links an extension only when the packaged catalog will emit its page.
func CCExtRef(cache *ExtensionCache, name string) string {
	if name == "" {
		return ""
	}
	if cache != nil {
		if ext, ok := cache.ExtMap[name]; ok && ext.IsReady() {
			return CCExtLink(name)
		}
	}
	return fmt.Sprintf("`%s`", name)
}

// CCExtBoldLink generates a bold code extension link.
func CCExtBoldLink(name, url string) string {
	if url != "" {
		return fmt.Sprintf("[**`%s`**](%s)", name, url)
	}
	return fmt.Sprintf("**`%s`**", name)
}

// DocsExtensionTable generates a localized markdown extension table.
func DocsExtensionTable(locale DocsLocale, exts []*Extension) string {
	var b strings.Builder
	b.WriteString(locale.t(
		"| **扩展** | **包** | **版本** | **许可证** | **语言** | **描述** |\n",
		"| **Extension** | **Package** | **Version** | **License** | **Language** | **Description** |\n",
	))
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
		desc := SanitizeText(locale.extDesc(ext))
		pkgLink := ext.GetPkgURLLink()
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
			ext.Name, ext.Name, pkgLink, version, license, lang, desc))
	}
	b.WriteString("{.ext-table}\n\n")
	return b.String()
}

func CCExtensionTable(exts []*Extension) string { return DocsExtensionTable(DocsZH, exts) }
func IOExtensionTable(exts []*Extension) string { return DocsExtensionTable(DocsEN, exts) }

// ============================================================================
// Shared presentation & query helpers across OINK (cc/io) and catalog (gen)
// ============================================================================

// preloadLibraries resolves ordered shared_preload_libraries (preload dependencies first, then self).
func preloadLibraries(cache *ExtensionCache, ext *Extension) []string {
	var libs []string
	seen := make(map[string]bool)
	add := func(raw string) {
		for _, lib := range strings.Split(raw, ",") {
			lib = strings.TrimSpace(lib)
			if lib != "" && !seen[lib] {
				seen[lib] = true
				libs = append(libs, lib)
			}
		}
	}
	if cache != nil {
		for _, depName := range ext.Requires {
			if depExt, ok := cache.ExtMap[depName]; ok && depExt.NeedLoad {
				add(depExt.GetLibName())
			}
		}
	}
	add(ext.GetLibName())
	return libs
}

// activePGVersionsFor returns ext's supported PG versions filtered to active catalog PG versions, descending.
func activePGVersionsFor(cache *ExtensionCache, ext *Extension) []int {
	if cache == nil {
		return nil
	}
	active := make(map[int]bool, len(cache.PGVersions))
	for _, pg := range cache.PGVersions {
		active[pg] = true
	}
	var out []int
	for _, pg := range ParsePGVersions(ext.PgVer) {
		if active[pg] {
			out = append(out, pg)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	return out
}

// extBuildFlags reports extra.rpm and extra.deb boolean build indicators.
func extBuildFlags(ext *Extension) (rpm, deb bool) {
	if ext.Extra == nil {
		return false, false
	}
	if v, ok := ext.Extra["rpm"].(bool); ok {
		rpm = v
	}
	if v, ok := ext.Extra["deb"].(bool); ok {
		deb = v
	}
	return rpm, deb
}

// buildOSCodeEntries groups active OSVersions into ordered OSCodeEntry pairs (x86_64 + aarch64).
func buildOSCodeEntries(osVersions []OSVersion) []OSCodeEntry {
	codeMap := make(map[string]*OSCodeEntry)
	var order []string
	for _, osv := range osVersions {
		if !osv.Active {
			continue
		}
		code, _, _ := strings.Cut(osv.OS, ".")
		entry, exists := codeMap[code]
		if !exists {
			entry = &OSCodeEntry{Code: code}
			codeMap[code] = entry
			order = append(order, code)
		}
		switch osv.OSArch {
		case "x86_64":
			entry.X86, entry.HasX86 = osv.OS, true
		case "aarch64":
			entry.ARM, entry.HasARM = osv.OS, true
		}
	}
	out := make([]OSCodeEntry, 0, len(order))
	for _, code := range order {
		out = append(out, *codeMap[code])
	}
	return out
}

// loadOSInfo retrieves active OS metadata from pgext.os.
func loadOSInfo(ctx context.Context, osName string) (*OSInfo, error) {
	const query = `SELECT os, os_full, os_major, os_arch, active FROM pgext.os WHERE os = $1 AND active = true`
	var info OSInfo
	err := QueryRowContext(ctx, query, osName).Scan(&info.OS, &info.Desc, &info.Major, &info.Arch, &info.Active)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// loadOSPackages loads lead package availability rows for one OS, sorted by extension ID.
func loadOSPackages(ctx context.Context, osName string) ([]*OSPackageInfo, error) {
	const query = `
		WITH lead_extensions AS (
			SELECT DISTINCT e.id, e.pkg, e.name AS lead
			FROM pgext.extension e
			WHERE e.lead = true
		)
		SELECT p.pg, p.pkg, le.lead, le.id, p.name, p.state, p.org, p.version, p.count
		FROM pgext.pkg p
		JOIN lead_extensions le ON p.pkg = le.pkg
		WHERE p.os = $1
		ORDER BY le.id, p.pg DESC
	`
	rows, err := QueryContext(ctx, query, osName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pkgMap := make(map[string]*OSPackageInfo)
	for rows.Next() {
		var pg, extID int
		var pkgName, lead string
		info := &PkgInfo{}
		if err := rows.Scan(&pg, &pkgName, &lead, &extID, &info.Name, &info.State, &info.Org, &info.Version, &info.Count); err != nil {
			return nil, err
		}
		info.PG, info.OS, info.Pkg, info.Ext = pg, osName, pkgName, lead
		entry, ok := pkgMap[pkgName]
		if !ok {
			entry = &OSPackageInfo{ID: extID, Pkg: pkgName, Lead: lead, PGData: make(map[int]*PkgInfo)}
			pkgMap[pkgName] = entry
		}
		entry.PGData[pg] = info
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	packages := make([]*OSPackageInfo, 0, len(pkgMap))
	for _, pkg := range pkgMap {
		packages = append(packages, pkg)
	}
	sort.Slice(packages, func(i, j int) bool { return packages[i].ID < packages[j].ID })
	return packages, nil
}

// runParallelOSPages runs genFn across all active OS targets with bounded parallelism.
func runParallelOSPages(ctx context.Context, genFn func(context.Context, string) error) error {
	osList, err := GetActiveOS(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active OS list: %w", err)
	}
	if len(osList) == 0 {
		return fmt.Errorf("no active OS distributions found")
	}

	logrus.Infof("Generating OS pages for %d distributions...", len(osList))
	sem := make(chan struct{}, 8)
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		failedOS []string
	)
	for _, osName := range osList {
		wg.Add(1)
		go func(os string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			if err := genFn(ctx, os); err != nil {
				logrus.Errorf("Failed to generate OS page for %s: %v", os, err)
				mu.Lock()
				failedOS = append(failedOS, os)
				mu.Unlock()
			}
		}(osName)
	}
	wg.Wait()

	logrus.Infof("Successfully generated %d/%d OS pages", len(osList)-len(failedOS), len(osList))
	if len(failedOS) > 0 {
		sort.Strings(failedOS)
		logrus.Warnf("Failed to generate pages for: %v", failedOS)
		return fmt.Errorf("failed to generate %d/%d OS pages: %s", len(failedOS), len(osList), strings.Join(failedOS, ", "))
	}
	return nil
}
