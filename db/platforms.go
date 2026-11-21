package db

import "strings"

// Build platforms are independent of historical repository/extension metadata.
// Keep this policy aligned with the active rows in os.csv (checked by tests).
var buildPlatforms = []struct{ code, codename string }{
	{"el8", ""}, {"el9", ""}, {"el10", ""},
	{"d12", "bookworm"}, {"d13", "trixie"},
	{"u22", "jammy"}, {"u24", "noble"}, {"u26", "resolute"},
}

// SupportsOSPG reports whether an active OS code supports a PostgreSQL major.
// EL8 remains in the active platform matrix for PG15–18; PG19+ has no EL8 PGDG/Pigsty packages.
func SupportsOSPG(osCode string, pg int) bool {
	if !SupportsPGVersion(pg) {
		return false
	}
	if osCode == "el8" && pg >= 19 {
		return false
	}
	return true
}

// SupportedPGVersionsForOS returns active PostgreSQL majors for an OS code in descending order.
func SupportedPGVersionsForOS(osCode string) []int {
	all := SupportedPGVersions()
	out := make([]int, 0, len(all))
	for _, pg := range all {
		if SupportsOSPG(osCode, pg) {
			out = append(out, pg)
		}
	}
	return out
}

// BuildOSCodename recognizes an active build OS and its APT suite, if any.
func BuildOSCodename(code string) (string, bool) {
	for _, platform := range buildPlatforms {
		if platform.code == code {
			return platform.codename, true
		}
	}
	return "", false
}

// SupportedOSPlatforms returns canonical build targets in display order.
func SupportedOSPlatforms() []string {
	platforms := make([]string, 0, len(buildPlatforms)*2)
	for _, platform := range buildPlatforms {
		for _, arch := range []string{"x86_64", "aarch64"} {
			platforms = append(platforms, platform.code+"."+arch)
		}
	}
	return platforms
}

// SupportsOSPlatform accepts canonical OS.arch names for current build targets.
func SupportsOSPlatform(platform string) bool {
	code, arch, ok := strings.Cut(platform, ".")
	_, supported := BuildOSCodename(code)
	return ok && supported && (arch == "x86_64" || arch == "aarch64")
}
