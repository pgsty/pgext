package build

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

const DefaultPgrxVersion = "0.19.3"

type EnvOptions struct {
	Options
	From, Mode, PgrxVersion string
}

type EnvResult struct {
	Action   string    `json:"action"`
	Platform Platform  `json:"platform"`
	Success  bool      `json:"success"`
	Plan     bool      `json:"plan"`
	Commands []Command `json:"commands,omitempty"`
	Paths    []string  `json:"paths,omitempty"`
	Error    string    `json:"error,omitempty"`
}

func (e *Engine) Environment(ctx context.Context, action string, opts EnvOptions) (*EnvResult, error) {
	r := &EnvResult{Action: "env " + action, Platform: e.Platform, Plan: e.Options.Plan}
	var err error
	switch action {
	case "status":
		r.Paths = []string{e.Options.Root, e.Options.SpecDir}
	case "tool":
		var c Command
		c, err = e.toolCommand(opts.Mode)
		if err == nil {
			r.Commands = append(r.Commands, c)
			if !r.Plan {
				err = e.Exec.Run(ctx, c, e.Err)
			}
		}
	case "spec":
		r.Paths = []string{e.Options.SpecDir, filepath.Join(e.Options.Root, "src"), filepath.Join(e.Options.Root, "pkg")}
		if !r.Plan {
			err = e.setupSpec(ctx, opts.From)
		}
	case "rust":
		r.Paths = []string{"~/.cargo/bin/cargo", "~/.cargo/config.toml"}
		if !r.Plan {
			err = e.setupRust(ctx)
		}
	case "pgrx":
		r.Commands, err = e.pgrxCommands(ctx, opts.PgrxVersion)
		if err == nil && !r.Plan {
			for _, c := range r.Commands {
				if err = e.Exec.Run(ctx, c, e.Err); err != nil {
					break
				}
			}
		}
	default:
		err = fmt.Errorf("unknown environment action %s", action)
	}
	r.Success = err == nil
	if err != nil {
		r.Error = err.Error()
	}
	return r, err
}

func (e *Engine) toolCommand(mode string) (Command, error) {
	if mode != "" && mode != "full" && mode != "mini" {
		return Command{}, fmt.Errorf("tool mode must be full or mini")
	}
	common := []string{"make", "cmake", "ninja-build", "git", "rsync", "tar", "gzip", "unzip", "jq", "file", "binutils", "patch"}
	var args []string
	if e.Platform.Family == "rpm" {
		args = []string{"dnf", "install", "-y"}
		common = append(common, "/usr/bin/curl", "xz", "gcc", "gcc-c++", "pkgconf-pkg-config", "rpm-build", "rpmdevtools", "dnf-plugins-core", "cpio", "lld", "openssl-devel", "readline-devel", "zlib-devel", "libxml2-devel", "lz4-devel", "libzstd-devel", "krb5-devel")
		if mode != "mini" {
			common = append(common, "pgdg-srpm-macros")
		}
	} else if e.Platform.Family == "deb" {
		args = []string{"apt-get", "install", "-y", "--no-install-recommends"}
		common = append(common, "curl", "xz-utils", "build-essential", "pkg-config", "debhelper", "devscripts", "fakeroot", "dpkg-dev", "postgresql-common", "lsb-release", "lld", "libssl-dev", "libreadline-dev", "zlib1g-dev", "libxml2-dev", "libxslt1-dev", "liblz4-dev", "libzstd-dev", "libkrb5-dev")
	} else {
		return Command{}, fmt.Errorf("env tool requires a Linux builder")
	}
	if mode != "mini" {
		pgs, err := ParsePG(e.Options.PG)
		if err != nil {
			return Command{}, err
		}
		if len(pgs) == 0 {
			pgs = []int{14, 15, 16, 17, 18}
		}
		for _, pg := range pgs {
			if e.Platform.Family == "rpm" {
				common = append(common, fmt.Sprintf("postgresql%d-devel", pg), fmt.Sprintf("postgresql%d-server", pg))
			} else {
				common = append(common, fmt.Sprintf("postgresql-server-dev-%d", pg), fmt.Sprintf("postgresql-%d", pg))
			}
		}
	}
	return Command{Args: append(args, common...), Root: true, Env: []string{"DEBIAN_FRONTEND=noninteractive"}}, nil
}

func (e *Engine) setupSpec(ctx context.Context, from string) error {
	if e.Platform.Family == "" {
		return fmt.Errorf("env spec requires a Linux builder")
	}
	for _, d := range []string{"src", "pkg", "log", "tmp"} {
		if err := os.MkdirAll(filepath.Join(e.Options.Root, d), 0755); err != nil {
			return err
		}
	}
	stage, err := os.MkdirTemp(filepath.Join(e.Options.Root, "tmp"), "spec-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(stage)
	root := e.Platform.Family + "build"
	if e.Platform.Family == "deb" {
		root = "debbuild"
	}
	if from == "" {
		bases := []string{"https://repo.pigsty.io/ext/spec", "https://repo.pigsty.cc/ext/spec"}
		if e.Options.Mirror {
			bases[0], bases[1] = bases[1], bases[0]
		}
		f, err := Download(ctx, bases, root+".tar.gz", filepath.Join(e.Options.Root, root+".tar.gz"), "", e.Options.Force, e.Err)
		if err != nil {
			return err
		}
		from = f.Path
	}
	i, err := os.Stat(from)
	if err != nil {
		return err
	}
	if i.IsDir() {
		if err := copyRecipe(from, filepath.Join(stage, root)); err != nil {
			return err
		}
	} else {
		if err := extractRecipes(from, stage, root); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(e.Options.SpecDir, 0755); err != nil {
		return err
	}
	// Updating recipes does not delete local files, sources, or artifacts.
	if err := e.Exec.Run(ctx, Command{Args: []string{"rsync", "-a", "--exclude=SOURCES", "--exclude=RPMS", "--exclude=DEBS", "--exclude=BUILD", "--exclude=BUILDROOT", "--exclude=SRPMS", filepath.Join(stage, root) + "/", e.Options.SpecDir + "/"}}, e.Err); err != nil {
		return err
	}
	for link, target := range map[string]string{"SOURCES": "src", map[bool]string{true: "RPMS", false: "DEBS"}[e.Platform.Family == "rpm"]: "pkg"} {
		if err := ensureLink(filepath.Join(e.Options.SpecDir, link), filepath.Join(e.Options.Root, target)); err != nil {
			return err
		}
	}
	return nil
}

func ensureLink(path, target string) error {
	if i, err := os.Lstat(path); err == nil {
		if i.Mode()&os.ModeSymlink != 0 {
			actual, err := filepath.EvalSymlinks(path)
			wanted, e := filepath.EvalSymlinks(target)
			if err == nil && e == nil && actual == wanted {
				return nil
			}
			return fmt.Errorf("existing link %s points elsewhere", path)
		}
		if i.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(entries) == 0 {
				if err := os.Remove(path); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("preserving nonempty %s; migrate it to %s before setup", path, target)
			}
		} else {
			return fmt.Errorf("preserving existing file %s", path)
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return os.Symlink(target, path)
}

func extractRecipes(path, dest, root string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name := filepath.Clean(strings.TrimPrefix(h.Name, "./"))
		if name == "." {
			continue
		}
		if name != root && !strings.HasPrefix(name, root+"/") {
			return fmt.Errorf("recipe archive entry outside %s: %s", root, name)
		}
		parts := strings.Split(name, "/")
		skip := false
		for _, part := range parts {
			if slices.Contains([]string{"SOURCES", "RPMS", "DEBS", "BUILD", "BUILDROOT", "SRPMS", ".git"}, part) {
				skip = true
			}
		}
		if skip {
			continue
		}
		to := filepath.Join(dest, name)
		switch h.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(to, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(to), 0755); err != nil {
				return err
			}
			out, err := os.OpenFile(to, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.FileMode(h.Mode)&0777)
			if err != nil {
				return err
			}
			_, err = io.Copy(out, tr)
			ce := out.Close()
			if err != nil {
				return err
			}
			if ce != nil {
				return ce
			}
		default:
			return fmt.Errorf("unsupported recipe archive entry %s", h.Name)
		}
	}
	return nil
}

func (e *Engine) setupRust(ctx context.Context) error {
	h, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cargo := filepath.Join(h, ".cargo")
	if e.Options.Mirror {
		b := []byte("[source.crates-io]\nreplace-with = \"rsproxy-sparse\"\n\n[source.rsproxy-sparse]\nregistry = \"sparse+https://rsproxy.cn/index/\"\n\n[net]\ngit-fetch-with-cli = true\n")
		p := filepath.Join(cargo, "config.toml")
		if old, err := os.ReadFile(p); err == nil && string(old) != string(b) {
			if !e.Options.Force {
				return fmt.Errorf("existing Cargo configuration differs; --force backs it up before replacement")
			}
			if err := AtomicWrite(p+".pgext-backup", old, 0600); err != nil {
				return err
			}
		} else if err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := AtomicWrite(p, b, 0644); err != nil {
			return err
		}
	}
	if _, err := os.Stat(filepath.Join(cargo, "bin", "cargo")); err == nil && !e.Options.Force {
		return nil
	}
	f, err := os.CreateTemp("", "pgext-rustup-*.sh")
	if err != nil {
		return err
	}
	path := f.Name()
	f.Close()
	defer os.Remove(path)
	url := "https://sh.rustup.rs"
	env := []string{}
	if e.Options.Mirror {
		url = "https://rsproxy.cn/rustup-init.sh"
		env = []string{"RUSTUP_DIST_SERVER=https://rsproxy.cn", "RUSTUP_UPDATE_ROOT=https://rsproxy.cn/rustup"}
	}
	if err := e.Exec.Run(ctx, Command{Args: []string{"curl", "--fail", "--location", "--retry", "3", url, "--output", path}}, e.Err); err != nil {
		return err
	}
	return e.Exec.Run(ctx, Command{Args: []string{"sh", path, "-y"}, Env: env}, e.Err)
}

func (e *Engine) pgrxCommands(ctx context.Context, version string) ([]Command, error) {
	if version == "" {
		version = DefaultPgrxVersion
	}
	if !regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(version) {
		return nil, fmt.Errorf("invalid pgrx version")
	}
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cargo := filepath.Join(h, ".cargo", "bin", "cargo")
	pgs, err := ParsePG(e.Options.PG)
	if err != nil {
		return nil, err
	}
	explicit := len(pgs) > 0
	if !explicit {
		pgs = []int{14, 15, 16, 17, 18}
	}
	args := []string{cargo, "pgrx", "init"}
	for _, pg := range pgs {
		path := e.Platform.PGConfig(pg)
		if _, err := os.Stat(path); err == nil || e.Options.Plan {
			args = append(args, fmt.Sprintf("--pg%d=%s", pg, path))
		} else if explicit {
			return nil, fmt.Errorf("missing %s", path)
		}
	}
	if len(args) == 3 {
		return nil, fmt.Errorf("no PostgreSQL installations; run pgext env tool")
	}
	commands := []Command{}
	var current string
	if !e.Options.Plan {
		current, err = e.Exec.Capture(ctx, Command{Args: []string{cargo, "pgrx", "--version"}, Env: e.buildEnv(0)})
	}
	if e.Options.Plan || err != nil || strings.TrimSpace(current) != "cargo-pgrx "+version {
		commands = append(commands, Command{Args: []string{cargo, "install", "--locked", "cargo-pgrx@" + version}, Env: e.buildEnv(0)})
	}
	return append(commands, Command{Args: args, Env: e.buildEnv(0)}), nil
}
