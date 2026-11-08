package repo

import (
	"fmt"
	"maps"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

const managedHeader = "# Managed by pgext repo\n"

type RenderOptions struct {
	Region  string
	AdminIP string
}

func Render(r Repository, p Platform, opts RenderOptions) (string, error) {
	if p.OS == "" || !r.Available(p) {
		return "", fmt.Errorf("repository %s cannot be rendered for %s; specify --platform", r.Name, p)
	}
	base := r.BaseURL[opts.Region]
	if strings.TrimSpace(base) == "" {
		base = r.BaseURL["default"]
	}
	version := strconv.Itoa(p.Major())
	if strings.HasPrefix(r.Name, "epel") && p.Major() >= 10 {
		version += "z"
	}
	if p.Family == "rpm" && strings.HasPrefix(r.Name, "pgdg") {
		parts := strings.Split(p.Version, ".")
		if p.Major() >= 10 {
			if len(parts) < 2 {
				return "", fmt.Errorf("%s needs the full EL minor version; use --os-version (e.g. 10.1)", r.Name)
			}
			version = p.Version
		} else if p.Major() == 9 && len(parts) > 1 {
			minor, _ := strconv.Atoi(parts[1])
			if minor >= 6 {
				version = strings.Join(parts[:2], ".")
			}
		}
	}
	if r.Minor {
		version = p.Version
	}
	meta := maps.Clone(r.Meta)
	if meta == nil {
		meta = map[string]string{}
	}
	arch := p.Arch
	if p.Family == "deb" {
		arch = "amd64"
		if p.Arch == "aarch64" {
			arch = "arm64"
		}
	}
	replacer := strings.NewReplacer(
		"${admin_ip}", opts.AdminIP,
		"${distro_codename}", p.Codename, "${distro_name}", p.Vendor,
		"${releasever}", version, "$releasever", version,
		"${basearch}", arch, "$basearch", arch,
	)
	expand := func(value string) (string, error) {
		if strings.Contains(value, "${admin_ip}") {
			u, err := url.Parse("http://" + opts.AdminIP)
			if err != nil || u.Host == "" || u.User != nil || u.Path != "" || u.RawQuery != "" || u.Fragment != "" || strings.ContainsAny(opts.AdminIP, " \t\r\n") {
				return "", fmt.Errorf("%s requires a valid --admin-ip host[:port]", r.Name)
			}
		}
		value = replacer.Replace(value)
		if strings.ContainsAny(value, "$\r\n\x00") {
			return "", fmt.Errorf("%s contains an unresolved variable or multiline value: %q", r.Name, value)
		}
		return value, nil
	}
	var err error
	base, err = expand(base)
	if err != nil {
		return "", err
	}
	for key, value := range meta {
		meta[key], err = expand(value)
		if err != nil {
			return "", err
		}
	}
	var b strings.Builder
	b.WriteString(managedHeader)
	fmt.Fprintf(&b, "# %s: %s\n", r.Name, r.Description)
	if p.Family == "rpm" {
		if _, ok := meta["enabled"]; !ok {
			meta["enabled"] = "1"
		}
		if _, ok := meta["gpgcheck"]; !ok {
			meta["gpgcheck"] = "0"
		}
		for _, endpoint := range strings.Fields(base) {
			if err := validateEndpoint(endpoint); err != nil {
				return "", fmt.Errorf("%s: %w", r.Name, err)
			}
		}
		fmt.Fprintf(&b, "[%s]\nname=%s\nbaseurl=%s\n", r.Name, r.Description, base)
		for _, key := range slices.Sorted(maps.Keys(meta)) {
			fmt.Fprintf(&b, "%s=%s\n", key, meta[key])
		}
	} else {
		parts := strings.Fields(base)
		if len(parts) < 2 || len(parts) == 2 && !strings.HasSuffix(parts[1], "/") {
			return "", fmt.Errorf("%s needs an APT URL, suite and components (or a flat suite ending in /)", r.Name)
		}
		if err := validateEndpoint(parts[0]); err != nil {
			return "", fmt.Errorf("%s: %w", r.Name, err)
		}
		if _, ok := meta["arch"]; !ok {
			meta["arch"] = arch
		}
		// Preserve Pigsty's trust policy unless the catalog declares a key/policy.
		if _, ok := meta["trusted"]; !ok && meta["signed-by"] == "" {
			meta["trusted"] = "yes"
		}
		var options []string
		for _, key := range slices.Sorted(maps.Keys(meta)) {
			value := meta[key]
			if value == "" || strings.ContainsAny(value, " \t[]#") {
				return "", fmt.Errorf("%s has invalid APT option %s=%q", r.Name, key, value)
			}
			options = append(options, key+"="+value)
		}
		fmt.Fprintf(&b, "deb [%s] %s\n", strings.Join(options, " "), strings.Join(parts, " "))
	}
	return b.String(), nil
}

func validateEndpoint(value string) error {
	u, err := url.Parse(value)
	if err != nil || u == nil || u.Scheme != "file" && u.Scheme != "http" && u.Scheme != "https" || u.Scheme != "file" && u.Host == "" || u.Fragment != "" {
		return fmt.Errorf("invalid repository URL %q", value)
	}
	return nil
}

func repoPath(name, family string) string {
	if family == "rpm" {
		return "etc/yum.repos.d/pgext-" + name + ".repo"
	}
	return "etc/apt/sources.list.d/pgext-" + name + ".list"
}
