package repo

import (
	"embed"
	"strings"
)

// The Pigsty key is retained from the builder. PGDG keys come from PostgreSQL's
// media/keys and pub/repos/yum/keys directories.
//
//go:embed keys/*.asc
var signingKeys embed.FS

func repoProvider(name string) string {
	switch {
	case strings.HasPrefix(name, "pgdg"):
		return "pgdg"
	case name == "pigsty-infra" || name == "pigsty-pgsql":
		return "pigsty"
	default:
		return ""
	}
}

func signingKeyPath(provider, family string) string {
	dir := "etc/apt/keyrings/"
	if family == "rpm" {
		dir = "etc/pki/rpm-gpg/"
	}
	return dir + "pgext-" + provider + ".asc"
}

func bundledSigningKey(r Repository, p Platform) (path, asset string) {
	provider := repoProvider(r.Name)
	if provider == "" {
		return "", ""
	}
	if p.Family == "rpm" {
		if r.Meta["gpgkey"] != "" || r.Meta["gpgcheck"] == "0" {
			return "", ""
		}
	} else if r.Meta["signed-by"] != "" || strings.EqualFold(r.Meta["trusted"], "yes") {
		return "", ""
	}

	key := provider
	if provider == "pgdg" {
		key = "pgdg-apt"
		if p.Family == "rpm" {
			key = "pgdg-rpm"
			if p.RPMArch() == "aarch64" {
				key = "pgdg-rpm-aarch64"
			}
		}
	}
	return signingKeyPath(provider, p.Family), "keys/" + key + ".asc"
}
