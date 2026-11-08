// Package repo manages Linux package sources using confgen's Pigsty catalog.
package repo

import (
	_ "embed"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"go.yaml.in/yaml/v3"
)

// These blocks are embedded verbatim in confgen as well as parsed by repo.
//
//go:embed catalogs/rpm.yml
var RPMCatalog string

//go:embed catalogs/deb.yml
var DEBCatalog string

type Repository struct {
	Name        string            `yaml:"name" json:"name"`
	Description string            `yaml:"description" json:"description"`
	Module      string            `yaml:"module" json:"module"`
	Releases    []int             `yaml:"releases" json:"releases"`
	Arch        []string          `yaml:"arch" json:"arch"`
	BaseURL     map[string]string `yaml:"baseurl" json:"baseurl"`
	Meta        map[string]string `yaml:"meta" json:"meta,omitempty"`
	Minor       bool              `yaml:"minor" json:"minor,omitempty"`
	Family      string            `yaml:"-" json:"family"`
}

var identifier = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)

func ParseCatalog(data []byte, family string) ([]Repository, error) {
	if family != "rpm" && family != "deb" {
		return nil, fmt.Errorf("a custom catalog requires a Linux --platform")
	}
	var doc struct {
		Upstream *[]Repository `yaml:"repo_upstream"`
		Default  *[]Repository `yaml:"repo_upstream_default"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parse repository catalog: %w", err)
	}
	entries := doc.Upstream
	if entries == nil {
		entries = doc.Default
	}
	if entries == nil {
		return nil, fmt.Errorf("catalog has neither repo_upstream nor repo_upstream_default")
	}
	for i := range *entries {
		r := &(*entries)[i]
		r.Family = family
		if !identifier.MatchString(r.Name) || !identifier.MatchString(r.Module) {
			return nil, fmt.Errorf("invalid repository name/module %q/%q", r.Name, r.Module)
		}
		if len(r.Releases) == 0 || len(r.Arch) == 0 || strings.TrimSpace(r.BaseURL["default"]) == "" {
			return nil, fmt.Errorf("repository %s requires releases, arch and a default baseurl", r.Name)
		}
		for j, arch := range r.Arch {
			r.Arch[j] = normalizeArch(arch)
			if r.Arch[j] != "x86_64" && r.Arch[j] != "aarch64" {
				return nil, fmt.Errorf("repository %s: unsupported architecture %q", r.Name, arch)
			}
		}
		for _, value := range append(slices.Collect(maps.Values(r.BaseURL)), r.Description) {
			if strings.ContainsAny(value, "\r\n\x00") {
				return nil, fmt.Errorf("repository %s contains a multiline value", r.Name)
			}
		}
		for k, v := range r.Meta {
			if !identifier.MatchString(k) || strings.ContainsAny(v, "\r\n\x00") {
				return nil, fmt.Errorf("repository %s has invalid metadata %q", r.Name, k)
			}
		}
		for _, other := range (*entries)[:i] {
			if other.Name != r.Name {
				continue
			}
			for _, release := range r.Releases {
				for _, arch := range r.Arch {
					if slices.Contains(other.Releases, release) && slices.Contains(other.Arch, arch) {
						return nil, fmt.Errorf("overlapping definitions for repository %s", r.Name)
					}
				}
			}
		}
	}
	return *entries, nil
}

func DefaultCatalog(family string) ([]Repository, error) {
	var result []Repository
	for _, f := range []string{"rpm", "deb"} {
		if family != "" && family != f {
			continue
		}
		data := RPMCatalog
		if f == "deb" {
			data = DEBCatalog
		}
		entries, err := ParseCatalog([]byte(data), f)
		if err != nil {
			return nil, err
		}
		result = append(result, entries...)
	}
	return result, nil
}

func (r Repository) Available(p Platform) bool {
	return p.OS == "" || r.Family == p.Family && slices.Contains(r.Releases, p.Major()) && slices.Contains(r.Arch, p.Arch)
}

// Select validates every selector, filters before deduplication, and keeps catalog order.
func Select(catalog []Repository, selectors []string, p Platform) ([]Repository, error) {
	var tokens []string
	for _, arg := range selectors {
		for _, token := range strings.Split(arg, ",") {
			token = strings.TrimSpace(token)
			if token == "" {
				return nil, fmt.Errorf("empty repository selector")
			}
			tokens = append(tokens, token)
		}
	}
	match := func(r Repository, token string) bool {
		switch token {
		case "all":
			return slices.Contains([]string{"node", "pgsql", "infra"}, r.Module)
		case "pgdg":
			return r.Module == "pgsql" && strings.HasPrefix(r.Name, "pgdg")
		case "pigsty":
			return slices.Contains([]string{"pgsql", "infra"}, r.Module) && strings.HasPrefix(r.Name, "pigsty-")
		default:
			return token == r.Name || token == r.Module
		}
	}
	selected := map[int]bool{}
	for _, token := range tokens {
		known, available := false, false
		for i, r := range catalog {
			if match(r, token) {
				known = true
				if r.Available(p) {
					available, selected[i] = true, true
				}
			}
		}
		if !known {
			return nil, fmt.Errorf("unknown repository or module %q; see pgext repo list", token)
		}
		if !available {
			return nil, fmt.Errorf("repository or module %q is unavailable on %s", token, p)
		}
	}
	result := []Repository{}
	for i, r := range catalog {
		if (len(tokens) == 0 || selected[i]) && r.Available(p) {
			result = append(result, r)
		}
	}
	return result, nil
}
