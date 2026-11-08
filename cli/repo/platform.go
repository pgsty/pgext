package repo

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Platform uses the same OS and architecture names as confgen.
type Platform struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Family   string `json:"family"`
	Version  string `json:"version"`
	Codename string `json:"codename,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
}

var codenames = map[string]string{
	"el8": "", "el9": "", "el10": "",
	"d11": "bullseye", "d12": "bookworm", "d13": "trixie",
	"u22": "jammy", "u24": "noble", "u26": "resolute",
}

func normalizeArch(arch string) string {
	switch arch {
	case "amd64", "x86_64":
		return "x86_64"
	case "arm64", "aarch64":
		return "aarch64"
	default:
		return arch
	}
}

func ParsePlatform(value, version string) (Platform, error) {
	code, arch, ok := strings.Cut(value, ".")
	codename, known := codenames[code]
	arch = normalizeArch(arch)
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
