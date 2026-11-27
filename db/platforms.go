package db

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Platform is the canonical OS and architecture descriptor shared across
// repository management, extension building, and configuration generation.
type Platform struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Family   string `json:"family"`
	Version  string `json:"version"`
	Codename string `json:"codename,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
}

// Build platforms are independent of historical repository/extension metadata.
// Keep this policy aligned with the active rows in os.csv (checked by tests).
var buildPlatforms = []struct{ code, codename string }{
	{"el8", ""}, {"el9", ""}, {"el10", ""},
	{"d12", "bookworm"}, {"d13", "trixie"},
	{"u22", "jammy"}, {"u24", "noble"}, {"u26", "resolute"},
}

// NormalizeArch maps Go/Debian/RPM architecture names to canonical x86_64/aarch64.
func NormalizeArch(arch string) string {
	switch arch {
	case "amd64", "x86_64":
		return "x86_64"
	case "arm64", "aarch64":
		return "aarch64"
	default:
		return arch
	}
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

// SupportedOSCodes returns the active OS codes in display order.
func SupportedOSCodes() []string {
	codes := make([]string, 0, len(buildPlatforms))
	for _, p := range buildPlatforms {
		codes = append(codes, p.code)
	}
	return codes
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

// ParsePlatform parses an OS.ARCH selector (e.g. el9.x86_64 or u24.arm64) and optional version.
func ParsePlatform(value, version string) (Platform, error) {
	code, arch, ok := strings.Cut(value, ".")
	codename, known := BuildOSCodename(code)
	arch = NormalizeArch(arch)
	if !ok || !known || (arch != "x86_64" && arch != "aarch64") {
		return Platform{}, fmt.Errorf("unsupported platform %q; use OS.ARCH, e.g. el9.x86_64 or u24.arm64", value)
	}
	p := Platform{OS: code, Arch: arch, Codename: codename, Family: "deb", Vendor: "debian"}
	major := strings.TrimLeft(code, "eldu")
	if strings.HasPrefix(code, "el") {
		p.Family, p.Vendor = "rpm", "rocky"
	} else if strings.HasPrefix(code, "u") {
		p.Vendor = "ubuntu"
	}
	if version == "" {
		version = major
	}
	for _, part := range strings.Split(version, ".") {
		if part == "" || strings.IndexFunc(part, func(r rune) bool { return r < '0' || r > '9' }) >= 0 {
			return Platform{}, fmt.Errorf("invalid OS version %q", version)
		}
	}
	if strings.Split(version, ".")[0] != major {
		return Platform{}, fmt.Errorf("OS version %s does not match %s", version, code)
	}
	p.Version = version
	return p, nil
}

// DetectPlatform reads /etc/os-release on Linux; returns zero Platform on non-Linux hosts.
func DetectPlatform() (Platform, error) {
	if runtime.GOOS != "linux" {
		return Platform{}, nil
	}
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return Platform{}, err
	}
	return ParseOSRelease(string(data), runtime.GOARCH)
}

// ParseOSRelease parses /etc/os-release content and a Go/Linux architecture string.
func ParseOSRelease(data, arch string) (Platform, error) {
	values := map[string]string{}
	for _, line := range strings.Split(data, "\n") {
		if k, v, ok := strings.Cut(line, "="); ok {
			values[k] = strings.Trim(v, `"'`)
		}
	}
	major := strings.Split(values["VERSION_ID"], ".")[0]
	var code string
	switch values["ID"] {
	case "rocky", "rhel", "almalinux", "centos", "ol":
		code = "el" + major
	case "debian":
		code = "d" + major
	case "ubuntu":
		code = "u" + major
	default:
		return Platform{}, fmt.Errorf("unsupported Linux distribution %q", values["ID"])
	}
	p, err := ParsePlatform(code+"."+arch, values["VERSION_ID"])
	if err == nil {
		p.Vendor = values["ID"]
		if values["VERSION_CODENAME"] != "" && p.Codename == "" {
			p.Codename = values["VERSION_CODENAME"]
		}
	}
	return p, err
}

func (p Platform) Major() int {
	n, _ := strconv.Atoi(strings.TrimLeft(p.OS, "eldu"))
	return n
}

func (p Platform) String() string {
	if p.OS == "" {
		return "all platforms"
	}
	return p.OS + "." + p.Arch
}

func (p Platform) RPMArch() string {
	return NormalizeArch(p.Arch)
}

func (p Platform) DEBArch() string {
	switch NormalizeArch(p.Arch) {
	case "aarch64":
		return "arm64"
	case "x86_64":
		return "amd64"
	default:
		return p.Arch
	}
}

// NativeArch returns the package manager's architecture spelling (x86_64/aarch64 for RPM, amd64/arm64 for DEB).
func (p Platform) NativeArch() string {
	if p.Family == "rpm" {
		return p.RPMArch()
	}
	if p.Family == "deb" {
		return p.DEBArch()
	}
	return p.Arch
}

// PGConfig returns the pg_config path for a PostgreSQL major version on this platform.
func (p Platform) PGConfig(pg int) string {
	if p.OS == "darwin" {
		for _, path := range []string{
			fmt.Sprintf("/opt/homebrew/opt/postgresql@%d/bin/pg_config", pg),
			fmt.Sprintf("/usr/local/opt/postgresql@%d/bin/pg_config", pg),
			fmt.Sprintf("/Applications/Postgres.app/Contents/Versions/%d/bin/pg_config", pg),
		} {
			if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
				return path
			}
		}
		return fmt.Sprintf("/opt/homebrew/opt/postgresql@%d/bin/pg_config", pg)
	}
	if p.Family == "rpm" {
		return fmt.Sprintf("/usr/pgsql-%d/bin/pg_config", pg)
	}
	return fmt.Sprintf("/usr/lib/postgresql/%d/bin/pg_config", pg)
}
