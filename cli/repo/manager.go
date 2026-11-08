package repo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

type Options struct {
	RenderOptions
	DryRun  bool
	Replace bool
	Update  bool
	All     bool
}

type Result struct {
	Action       string          `json:"action"`
	Platform     Platform        `json:"platform"`
	Catalog      string          `json:"catalog,omitempty"`
	Root         string          `json:"root"`
	DryRun       bool            `json:"dry_run"`
	Success      bool            `json:"success"`
	Repositories []Repository    `json:"repositories,omitempty"`
	Installed    []InstalledFile `json:"installed,omitempty"`
	Changes      []Change        `json:"changes"`
	Commands     [][]string      `json:"commands,omitempty"`
	Backup       string          `json:"backup,omitempty"`
	RolledBack   bool            `json:"rolled_back,omitempty"`
	Error        string          `json:"error,omitempty"`
}

type Manager struct {
	Root     string
	Platform Platform
	Catalog  []Repository
	// RunCommand allows deterministic package-manager failure tests.
	RunCommand func(context.Context, []string) error
}

func NewManager(root string, p Platform, catalog []Repository, diagnostics io.Writer) (*Manager, error) {
	if root == "" {
		root = "/"
	}
	root, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	// Resolve only the explicitly selected root. Paths beneath it may not be symlinks.
	resolved, err := filepath.EvalSymlinks(root)
	if err != nil {
		return nil, fmt.Errorf("repository root must exist: %w", err)
	}
	info, err := os.Stat(resolved)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("repository root must be a directory")
	}
	m := &Manager{Root: resolved, Platform: p, Catalog: catalog}
	m.RunCommand = func(ctx context.Context, args []string) error {
		c := exec.CommandContext(ctx, args[0], args[1:]...)
		c.Stdout, c.Stderr = diagnostics, diagnostics
		if err := c.Run(); err != nil {
			return fmt.Errorf("%s: %w", strings.Join(args, " "), err)
		}
		return nil
	}
	return m, nil
}

func (m *Manager) Execute(ctx context.Context, action string, args []string, opts Options) (*Result, error) {
	r := &Result{Action: action, Root: m.Root, Platform: m.Platform, DryRun: opts.DryRun, Changes: []Change{}}
	err := m.execute(ctx, r, args, opts)
	r.Success = err == nil
	if err != nil {
		r.Error = err.Error()
	}
	return r, err
}

func (m *Manager) execute(ctx context.Context, r *Result, args []string, opts Options) error {
	if m.Platform.OS == "" {
		return fmt.Errorf("repository operations require Linux; use --platform for preview or --root for staging")
	}
	if opts.Replace && r.Action != "set" && r.Action != "add" {
		return fmt.Errorf("--replace is only valid with add or set")
	}
	if opts.All && r.Action != "rm" {
		return fmt.Errorf("--all is only valid with rm")
	}
	if (opts.Update || r.Action == "update") && m.Root != "/" {
		return fmt.Errorf("cache updates are unavailable with --root; use --update=false")
	}
	files, err := m.scan()
	if err != nil {
		return err
	}
	r.Installed = installed(m.Root, files)
	if r.Action == "status" {
		return nil
	}
	after := maps.Clone(files)
	switch r.Action {
	case "add", "set":
		if len(args) == 0 {
			args = []string{"all"}
		}
		selected, err := Select(m.Catalog, args, m.Platform)
		if err != nil {
			return err
		}
		if len(selected) == 0 {
			return fmt.Errorf("no repositories selected")
		}
		r.Repositories = selected
		for path, image := range files {
			if opts.Replace || r.Action == "set" && owned(path, image) {
				delete(after, path)
			}
		}
		for _, repository := range selected {
			path := repoPath(repository.Name, m.Platform.Family)
			if old := after[path]; old != nil && !owned(path, old) {
				return fmt.Errorf("%s is not managed by pgext repo; use set --replace to back up and replace existing sources", path)
			}
			content, err := Render(repository, m.Platform, opts.RenderOptions)
			if err != nil {
				return err
			}
			after[path] = &fileImage{Content: content, Mode: 0644}
		}
		// Check selected definitions against every retained source, including deb822.
		for _, repository := range selected {
			path := repoPath(repository.Name, m.Platform.Family)
			keys := sourceKeys(path, after[path].Content, m.Platform.Family)
			for other, image := range after {
				if other == path {
					continue
				}
				for _, key := range sourceKeys(other, image.Content, m.Platform.Family) {
					if slices.Contains(keys, key) {
						return fmt.Errorf("%s conflicts with %s (%s); use set --replace to back up and replace existing sources", path, other, key)
					}
				}
			}
		}
	case "rm":
		if opts.All && len(args) != 0 {
			return fmt.Errorf("use either rm --all or named selectors")
		}
		if !opts.All && len(args) == 0 {
			return fmt.Errorf("specify repository names/modules, or --all for all pgext-managed files")
		}
		selected, err := Select(m.Catalog, args, m.Platform)
		if err != nil {
			return err
		}
		paths := map[string]bool{}
		for _, repository := range selected {
			paths[repoPath(repository.Name, m.Platform.Family)] = true
		}
		for path, image := range files {
			if owned(path, image) && (opts.All || paths[path]) {
				delete(after, path)
			}
		}
	case "restore":
		if len(args) != 1 {
			return fmt.Errorf("restore requires one backup ID")
		}
		r.Changes, err = m.restoreChanges(args[0])
		if err != nil {
			return err
		}
	case "update":
	default:
		return fmt.Errorf("unknown repository action %q", r.Action)
	}
	if r.Action != "restore" {
		paths := maps.Clone(files)
		maps.Copy(paths, after)
		for _, path := range slices.Sorted(maps.Keys(paths)) {
			if !sameImage(files[path], after[path]) {
				r.Changes = append(r.Changes, Change{Path: path, Before: files[path], After: after[path]})
			}
		}
	}
	if opts.Update || r.Action == "update" {
		command := []string{"apt-get", "update", "-o", "APT::Update::Error-Mode=any"}
		if m.Platform.Family == "rpm" {
			command = []string{"dnf", "makecache"}
		}
		r.Commands = [][]string{command}
	}
	if opts.DryRun || len(r.Changes) == 0 && len(r.Commands) == 0 {
		return nil
	}
	if m.Root == "/" {
		host, err := DetectPlatform()
		if err != nil {
			return err
		}
		if runtime.GOOS != "linux" || !compatibleHostPlatform(host, m.Platform) {
			return fmt.Errorf("system repository changes require the detected Linux platform; use --dry-run or --root")
		}
		if os.Geteuid() != 0 {
			return fmt.Errorf("system repository changes require root; run with sudo or use --dry-run")
		}
	}
	return m.apply(ctx, r, files)
}

// compatibleHostPlatform guards system writes without making an explicit
// generic target such as el9.x86_64 fail on an RHEL-compatible host whose
// /etc/os-release says almalinux or rhel, or whose VERSION_ID has a minor.
func compatibleHostPlatform(host, requested Platform) bool {
	if host.OS != requested.OS || host.Arch != requested.Arch || host.Family != requested.Family {
		return false
	}
	if requested.Codename != "" && host.Codename != requested.Codename {
		return false
	}
	hostMajor := strings.Split(host.Version, ".")[0]
	requestedMajor := strings.Split(requested.Version, ".")[0]
	return requested.Version == "" || requested.Version == host.Version || requested.Version == requestedMajor && hostMajor == requestedMajor
}

func (m *Manager) apply(ctx context.Context, r *Result, files map[string]*fileImage) error {
	unlock, err := m.lock()
	if err != nil {
		return err
	}
	defer unlock()
	current, err := m.scan()
	if err != nil {
		return err
	}
	if !maps.EqualFunc(files, current, sameImage) {
		return fmt.Errorf("repository files changed during planning; retry the operation")
	}
	for _, c := range r.Changes {
		if !sameImage(current[c.Path], c.Before) {
			return fmt.Errorf("%s changed since backup validation", c.Path)
		}
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(r.Changes) > 0 {
		r.Backup, err = m.saveBackup(r.Changes)
		if err != nil {
			return err
		}
	}
	var applied []Change
	fail := func(cause error) error {
		if len(applied) == 0 {
			return cause
		}
		err := m.rollback(applied)
		r.RolledBack = err == nil
		if err != nil {
			return errors.Join(cause, fmt.Errorf("rollback incomplete (backup %s): %w", r.Backup, err))
		}
		return fmt.Errorf("repository changes rolled back (backup %s): %w", r.Backup, cause)
	}
	for _, c := range r.Changes {
		if err := ctx.Err(); err != nil {
			return fail(err)
		}
		current, err := m.read(c.Path)
		if err != nil {
			return fail(err)
		}
		if !sameImage(current, c.Before) {
			return fail(fmt.Errorf("%s changed concurrently", c.Path))
		}
		if err := m.write(c.Path, c.After); err != nil {
			return fail(err)
		}
		applied = append(applied, c)
	}
	for _, command := range r.Commands {
		if err := m.RunCommand(ctx, command); err != nil {
			return fail(err)
		}
	}
	return nil
}

// sourceKeys identifies RPM section IDs and active APT endpoint/suite pairs.
// Comparing http and https equally catches conflicting Signed-By/Trusted policy.
func sourceKeys(path, content, family string) []string {
	var keys []string
	add := func(endpoint, suite string) {
		u, err := url.Parse(endpoint)
		if err == nil {
			keys = append(keys, strings.ToLower(u.Host)+strings.TrimRight(u.Path, "/")+" "+suite)
		}
	}
	if family == "rpm" {
		for _, line := range strings.Split(content, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "[") {
				if name, _, ok := strings.Cut(line[1:], "]"); ok {
					keys = append(keys, name)
				}
			}
		}
	} else if strings.HasSuffix(path, ".sources") {
		fields := map[string]string{}
		flush := func() {
			if strings.ToLower(fields["enabled"]) != "no" && slices.Contains(strings.Fields(fields["types"]), "deb") {
				for _, endpoint := range strings.Fields(fields["uris"]) {
					for _, suite := range strings.Fields(fields["suites"]) {
						add(endpoint, suite)
					}
				}
			}
			fields = map[string]string{}
		}
		last := ""
		for _, line := range strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n") {
			if strings.TrimSpace(line) == "" {
				flush()
				last = ""
			} else if strings.HasPrefix(strings.TrimSpace(line), "#") {
				continue
			} else if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
				fields[last] += " " + strings.TrimSpace(line)
			} else if k, v, ok := strings.Cut(line, ":"); ok {
				last = strings.ToLower(strings.TrimSpace(k))
				fields[last] = strings.TrimSpace(v)
			}
		}
		flush()
	} else {
		for _, line := range strings.Split(content, "\n") {
			parts := strings.Fields(line)
			if len(parts) == 0 || parts[0] != "deb" {
				continue
			}
			parts = parts[1:]
			if len(parts) > 0 && strings.HasPrefix(parts[0], "[") {
				for len(parts) > 0 {
					last := strings.HasSuffix(parts[0], "]")
					parts = parts[1:]
					if last {
						break
					}
				}
			}
			if len(parts) >= 2 {
				add(parts[0], parts[1])
			}
		}
	}
	return keys
}
