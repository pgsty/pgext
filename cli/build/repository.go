package build

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"
)

//go:embed assets/pigsty.asc
var pigstyKey []byte

type RepoFile struct {
	Path    string `json:"path"`
	Content string `json:"content,omitempty"`
	Exists  bool   `json:"exists"`
	Remove  bool   `json:"remove,omitempty"`
}
type RepoResult struct {
	Action   string     `json:"action"`
	Platform Platform   `json:"platform"`
	Plan     bool       `json:"plan"`
	Success  bool       `json:"success"`
	Files    []RepoFile `json:"files"`
	Commands []Command  `json:"commands,omitempty"`
	Backup   string     `json:"backup,omitempty"`
	Error    string     `json:"error,omitempty"`
}

func (e *Engine) Repo(ctx context.Context, action string, modules []string, update bool) (*RepoResult, error) {
	r := &RepoResult{Action: "env repo " + action, Platform: e.Platform, Plan: e.Options.Plan, Files: []RepoFile{}}
	err := e.repo(ctx, r, action, modules, update)
	r.Success = err == nil
	if err != nil {
		r.Error = err.Error()
	}
	return r, err
}

func (e *Engine) repo(ctx context.Context, r *RepoResult, action string, modules []string, update bool) error {
	if e.Platform.Family == "" {
		return fmt.Errorf("repository setup requires a supported Linux builder")
	}
	dir, suffix := "/etc/apt/sources.list.d", ".sources"
	cache := Command{Args: []string{"apt-get", "update", "-o", "APT::Update::Error-Mode=any"}, Root: true}
	if e.Platform.Family == "rpm" {
		dir, suffix = "/etc/yum.repos.d", ".repo"
		cache.Args = []string{"dnf", "makecache"}
	}
	if action == "list" || action == "status" {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, ent := range entries {
			if strings.HasSuffix(ent.Name(), suffix) || strings.HasSuffix(ent.Name(), ".list") {
				r.Files = append(r.Files, RepoFile{Path: filepath.Join(dir, ent.Name()), Exists: true})
			}
		}
		return nil
	}
	if action == "update" {
		r.Commands = []Command{cache}
		if !r.Plan {
			return e.Exec.Run(ctx, cache, e.Err)
		}
		return nil
	}
	if action != "add" && action != "rm" {
		return fmt.Errorf("unknown repository action %s", action)
	}
	if len(modules) == 0 {
		if action == "rm" {
			return fmt.Errorf("specify pgdg or pigsty to remove")
		}
		modules = []string{"all"}
	}
	var expanded []string
	for _, m := range modules {
		for _, s := range strings.Split(m, ",") {
			switch s {
			case "all":
				expanded = append(expanded, "node", "pgdg", "pigsty")
			case "node", "pgdg", "pigsty":
				expanded = append(expanded, s)
			default:
				return fmt.Errorf("unknown repository module %q (all, node, pgdg, pigsty)", s)
			}
		}
	}
	modules = unique(expanded)
	if action == "rm" && slices.Contains(modules, "node") {
		return fmt.Errorf("node repositories are OS-owned; only pgext's pgdg/pigsty files can be removed")
	}
	for _, m := range modules {
		if m == "node" {
			continue
		}
		content, err := e.repoContent(m)
		if err != nil {
			return err
		}
		p := filepath.Join(dir, "pgext-"+m+suffix)
		_, err = os.Stat(p)
		r.Files = append(r.Files, RepoFile{Path: p, Content: content, Exists: err == nil})
	}
	if action == "add" && (slices.Contains(modules, "pgdg") || slices.Contains(modules, "pigsty")) {
		legacySuffix := ".list"
		if e.Platform.Family == "rpm" {
			legacySuffix = ".repo"
		}
		for _, name := range []string{"infra", "pgsql"} {
			path := filepath.Join(dir, name+legacySuffix)
			b, err := os.ReadFile(path)
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return err
			}
			if !recognizedLegacyRepos(string(b), e.Platform.Family) {
				return fmt.Errorf("%s contains custom repository definitions; migrate it explicitly before adding PGEXT repositories", path)
			}
			if !e.Options.ReplaceRepos || !slices.Contains(modules, "pgdg") || !slices.Contains(modules, "pigsty") {
				return fmt.Errorf("legacy PIG repository file %s; use pgext env repo add pgdg pigsty --replace to back up and migrate both modules", path)
			}
			r.Files = append(r.Files, RepoFile{Path: path, Exists: true, Remove: true})
		}
	}
	if action == "add" && slices.Contains(modules, "node") && e.Platform.Family == "rpm" {
		crb := "crb"
		if e.Platform.OS == "el8" {
			crb = "powertools"
		}
		r.Commands = append(r.Commands, Command{Args: []string{"dnf", "install", "-y", "dnf-plugins-core", "epel-release"}, Root: true}, Command{Args: []string{"dnf", "config-manager", "--set-enabled", crb}, Root: true})
		if e.Platform.OS != "el10" {
			r.Commands = append(r.Commands, Command{Args: []string{"dnf", "module", "disable", "-y", "postgresql"}, Root: true})
		}
	}
	if update {
		r.Commands = append(r.Commands, cache)
	}
	if r.Plan {
		return nil
	}
	if os.Geteuid() != 0 {
		return fmt.Errorf("env repo add/rm writes system repository files; run as root")
	}
	if action == "add" {
		if slices.Contains(modules, "pigsty") {
			p := "/etc/apt/keyrings/pgext-pigsty.asc"
			if e.Platform.Family == "rpm" {
				p = "/etc/pki/rpm-gpg/pgext-pigsty.asc"
			}
			if err := AtomicWrite(p, pigstyKey, 0644); err != nil {
				return err
			}
		}
		if slices.Contains(modules, "pgdg") {
			base, name, p := "https://apt.postgresql.org/pub/repos/apt", "ACCC4CF8.asc", "/etc/apt/keyrings/pgext-pgdg.asc"
			if e.Platform.Family == "rpm" {
				base = "https://download.postgresql.org/pub/repos/yum/keys"
				name = "PGDG-RPM-GPG-KEY-RHEL"
				if e.Platform.Arch == "arm64" {
					name = "PGDG-RPM-GPG-KEY-AARCH64-RHEL"
				}
				p = "/etc/pki/rpm-gpg/pgext-pgdg.asc"
			}
			if _, err := Download(ctx, []string{base}, name, p, "", false, e.Err); err != nil {
				return err
			}
		}
	}
	backup := filepath.Join(dir, "pgext-backup-"+time.Now().UTC().Format("20060102T150405.000000000"))
	r.Backup = backup
	if err := os.MkdirAll(backup, 0700); err != nil {
		return err
	}
	saved := map[string][]byte{}
	for _, f := range r.Files {
		b, err := os.ReadFile(f.Path)
		if err == nil {
			saved[f.Path] = b
			if err := AtomicWrite(filepath.Join(backup, filepath.Base(f.Path)), b, 0600); err != nil {
				return err
			}
		} else if !os.IsNotExist(err) {
			return err
		}
	}
	rollback := func() {
		for _, f := range r.Files {
			if b, ok := saved[f.Path]; ok {
				_ = AtomicWrite(f.Path, b, 0644)
			} else {
				_ = os.Remove(f.Path)
			}
		}
	}
	for _, f := range r.Files {
		var err error
		if action == "rm" || f.Remove {
			err = os.Remove(f.Path)
			if os.IsNotExist(err) {
				err = nil
			}
		} else {
			err = AtomicWrite(f.Path, []byte(f.Content), 0644)
		}
		if err != nil {
			rollback()
			return err
		}
	}
	for _, c := range r.Commands {
		if err := e.Exec.Run(ctx, c, e.Err); err != nil {
			rollback()
			return fmt.Errorf("repository operation failed; restored managed definitions (backup %s): %w", backup, err)
		}
	}
	return nil
}

// Only the old PIG-generated infra/pgsql files qualify for bounded migration.
// A mixed file containing any custom repository must be handled explicitly.
func recognizedLegacyRepos(content, family string) bool {
	seen := false
	section, hasURL := "", false
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if family == "deb" {
			if !regexp.MustCompile(`^deb(?:-src)?\s+(?:\[[^\]]+\]\s+)?https?://(?:repo\.pigsty\.(?:io|cc)/apt/(?:infra|pgsql)(?:/|\s)|apt\.postgresql\.org/pub/repos/apt(?:/|\s)|mirrors\.(?:tuna\.tsinghua\.edu\.cn|ustc\.edu\.cn)/postgresql/repos/apt(?:/|\s)|mirrors\.cloud\.tencent\.com/postgresql/repos/apt(?:/|\s))`).MatchString(line) {
				return false
			}
			seen = true
		} else if strings.HasPrefix(line, "[") {
			if section != "" && !hasURL {
				return false
			}
			if !regexp.MustCompile(`^\[(?:pigsty-(?:infra|pgsql)|pgdg(?:-common|1[4-8]))\]$`).MatchString(line) {
				return false
			}
			seen = true
			section, hasURL = strings.Trim(line, "[]"), false
		} else {
			key, value, ok := strings.Cut(line, "=")
			if !seen || !ok {
				return false
			}
			key, value = strings.TrimSpace(key), strings.TrimSpace(value)
			switch key {
			case "baseurl":
				urls := strings.Fields(value)
				if len(urls) == 0 {
					return false
				}
				for _, u := range urls {
					if !legacyRPMURL(section, u) {
						return false
					}
				}
				hasURL = true
			case "name", "enabled", "gpgcheck", "repo_gpgcheck", "module_hotfixes", "priority":
				// Standard PIG-rendered metadata; no custom URL or package filters.
			default:
				return false
			}
		}
	}
	return seen && (family == "deb" || hasURL)
}

func legacyRPMURL(section, value string) bool {
	if strings.HasPrefix(section, "pigsty-") {
		module := strings.TrimPrefix(section, "pigsty-")
		return regexp.MustCompile(`^https?://repo\.pigsty\.(?:io|cc)/yum/` + module + `(?:/|$)`).MatchString(value)
	}
	return regexp.MustCompile(`^https?://(?:download\.postgresql\.org/pub/repos/yum|repo\.pigsty\.(?:cc|io)/yum/pgdg|mirrors\.(?:cloud\.tencent\.com|tuna\.tsinghua\.edu\.cn|ustc\.edu\.cn|xtom\.de)/postgresql/repos/yum)/(?:common|1[4-8])/redhat/rhel-[^\s]+$`).MatchString(value)
}

func (e *Engine) repoContent(module string) (string, error) {
	p := e.Platform
	base := "https://repo.pigsty.io"
	if e.Options.Mirror {
		base = "https://repo.pigsty.cc"
	}
	if p.Family == "deb" {
		if module == "pgdg" {
			return fmt.Sprintf("# Managed by pgext env repo\nTypes: deb\nURIs: https://apt.postgresql.org/pub/repos/apt\nSuites: %s-pgdg\nComponents: main\nSigned-By: /etc/apt/keyrings/pgext-pgdg.asc\n", p.Codename), nil
		}
		return fmt.Sprintf("# Managed by pgext env repo\nTypes: deb\nURIs: %s/apt/infra\nSuites: generic\nComponents: main\nSigned-By: /etc/apt/keyrings/pgext-pigsty.asc\n\nTypes: deb\nURIs: %s/apt/pgsql/%s\nSuites: %s\nComponents: main\nSigned-By: /etc/apt/keyrings/pgext-pigsty.asc\n", base, base, p.Codename, p.Codename), nil
	}
	var b strings.Builder
	write := func(name, url, key string, priority int) {
		fmt.Fprintf(&b, "[pgext-%s]\nname=PGEXT %s\nbaseurl=%s\nenabled=1\ngpgcheck=1\ngpgkey=file:///etc/pki/rpm-gpg/pgext-%s.asc\nmodule_hotfixes=1\npriority=%d\n\n", name, name, url, key, priority)
	}
	if module == "pigsty" {
		write("pigsty-infra", base+"/yum/infra/"+p.NativeArch(), "pigsty", 12)
		write("pigsty-pgsql", base+"/yum/pgsql/"+p.OS+"."+p.NativeArch(), "pigsty", 11)
	} else {
		release := p.Version
		if p.OS == "el8" {
			release = "8"
		}
		if (p.OS == "el9" || p.OS == "el10") && !strings.Contains(release, ".") {
			return "", fmt.Errorf("PGDG requires the full EL minor version; run repository setup/preview inside the target builder")
		}
		base = "https://download.postgresql.org/pub/repos/yum"
		if e.Options.Mirror {
			base = "https://repo.pigsty.cc/yum/pgdg"
		}
		write("pgdg-common", base+"/common/redhat/rhel-"+release+"-"+p.NativeArch(), "pgdg", 10)
		for _, pg := range []int{14, 15, 16, 17, 18} {
			write(fmt.Sprintf("pgdg%d", pg), fmt.Sprintf("%s/%d/redhat/rhel-%s-%s", base, pg, release, p.NativeArch()), "pgdg", 10)
		}
	}
	return b.String(), nil
}
