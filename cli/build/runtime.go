package build

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"syscall"
	"time"
)

type Platform struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Family   string `json:"family"`
	Version  string `json:"version"`
	Codename string `json:"codename,omitempty"`
}

func DetectPlatform() (Platform, error) {
	p := Platform{OS: runtime.GOOS, Arch: runtime.GOARCH}
	if runtime.GOOS != "linux" {
		return p, nil
	}
	b, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return p, err
	}
	return ParsePlatform(string(b), runtime.GOARCH)
}

func ParsePlatform(release, arch string) (Platform, error) {
	m := map[string]string{}
	for _, l := range strings.Split(release, "\n") {
		if k, v, ok := strings.Cut(l, "="); ok {
			m[k] = strings.Trim(v, `"'`)
		}
	}
	p := Platform{Arch: arch, Version: m["VERSION_ID"], Codename: m["VERSION_CODENAME"]}
	major := strings.Split(p.Version, ".")[0]
	switch m["ID"] {
	case "rocky", "rhel", "almalinux", "centos", "ol":
		p.OS, p.Family = "el"+major, "rpm"
	case "debian":
		p.OS, p.Family = "d"+major, "deb"
	case "ubuntu":
		p.OS, p.Family = "u"+major, "deb"
	default:
		return p, fmt.Errorf("unsupported Linux distribution %q", m["ID"])
	}
	if _, ok := supportedOS[p.OS]; !ok {
		return p, fmt.Errorf("unsupported build platform %s", p.OS)
	}
	if arch != "amd64" && arch != "arm64" {
		return p, fmt.Errorf("unsupported architecture %s", arch)
	}
	return p, nil
}

var supportedOS = map[string]string{"el8": "", "el9": "", "el10": "", "d12": "bookworm", "d13": "trixie", "u22": "jammy", "u24": "noble", "u26": "resolute"}

func PlanPlatform(value string) (Platform, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 2 {
		return Platform{}, fmt.Errorf("platform must be OS.ARCH, e.g. el9.arm64")
	}
	code, ok := supportedOS[parts[0]]
	if !ok {
		return Platform{}, fmt.Errorf("unsupported platform %s", value)
	}
	arch := parts[1]
	if arch == "x86_64" {
		arch = "amd64"
	}
	if arch == "aarch64" {
		arch = "arm64"
	}
	if arch != "amd64" && arch != "arm64" {
		return Platform{}, fmt.Errorf("unsupported architecture %s", arch)
	}
	family := "deb"
	if strings.HasPrefix(parts[0], "el") {
		family = "rpm"
	}
	return Platform{OS: parts[0], Family: family, Arch: arch, Codename: code, Version: strings.TrimLeft(parts[0], "eldu")}, nil
}

func (p Platform) PGConfig(pg int) string {
	if p.OS == "darwin" {
		for _, path := range []string{fmt.Sprintf("/opt/homebrew/opt/postgresql@%d/bin/pg_config", pg), fmt.Sprintf("/usr/local/opt/postgresql@%d/bin/pg_config", pg), fmt.Sprintf("/Applications/Postgres.app/Contents/Versions/%d/bin/pg_config", pg)} {
			if fileExists(path) {
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
func (p Platform) NativeArch() string {
	if p.Family == "rpm" {
		if p.Arch == "arm64" {
			return "aarch64"
		}
		return "x86_64"
	}
	return p.Arch
}

type Command struct {
	Args []string `json:"args"`
	Dir  string   `json:"dir,omitempty"`
	Env  []string `json:"env,omitempty"`
	Root bool     `json:"root,omitempty"`
}

// Executor is injected by tests; all native output goes to diagnostics/logs.
type Executor interface {
	Run(context.Context, Command, io.Writer) error
	Capture(context.Context, Command) (string, error)
}
type NativeExecutor struct{}

func nativeCommand(ctx context.Context, c Command) *exec.Cmd {
	args := c.Args
	if c.Root && os.Geteuid() != 0 {
		prefix := []string{"sudo", "--"}
		if len(c.Env) > 0 {
			prefix = append(append(prefix, "env"), c.Env...)
		}
		args = append(prefix, args...)
	}
	env := MergeEnv(os.Environ(), c.Env)
	program, lookupErr := executableInEnv(args[0], env)
	cmd := exec.CommandContext(ctx, program, args[1:]...)
	if lookupErr != nil {
		cmd.Err = lookupErr
	}
	cmd.Dir, cmd.Env = c.Dir, env
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		if cmd.Process == nil {
			return nil
		}
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
	cmd.WaitDelay = 5 * time.Second
	return cmd
}

// exec.Command resolves a basename before Cmd.Env is applied. Resolve it
// against the actual build PATH so Rust provenance and execution agree.
func executableInEnv(name string, env []string) (string, error) {
	if strings.ContainsRune(name, filepath.Separator) {
		return name, nil
	}
	var path string
	for _, v := range env {
		if p, ok := strings.CutPrefix(v, "PATH="); ok {
			path = p
		}
	}
	for _, dir := range filepath.SplitList(path) {
		p := filepath.Join(dir, name)
		if i, err := os.Stat(p); err == nil && !i.IsDir() && i.Mode()&0111 != 0 {
			if !filepath.IsAbs(p) {
				return name, fmt.Errorf("%s: %w", name, exec.ErrDot)
			}
			return p, nil
		}
	}
	return name, fmt.Errorf("%s: %w", name, exec.ErrNotFound)
}
func (NativeExecutor) Run(ctx context.Context, c Command, w io.Writer) error {
	cmd := nativeCommand(ctx, c)
	cmd.Stdout, cmd.Stderr = w, w
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", filepath.Base(c.Args[0]), err)
	}
	return nil
}
func (NativeExecutor) Capture(ctx context.Context, c Command) (string, error) {
	cmd := nativeCommand(ctx, c)
	var b, diagnostics strings.Builder
	cmd.Stdout, cmd.Stderr = &b, &diagnostics
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %w: %s", filepath.Base(c.Args[0]), err, strings.TrimSpace(diagnostics.String()+b.String()))
	}
	return b.String(), nil
}
func MergeEnv(base, overrides []string) []string {
	m := map[string]string{}
	for _, e := range append(append([]string{}, base...), overrides...) {
		if k, v, ok := strings.Cut(e, "="); ok {
			m[k] = v
		}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]string, 0, len(m))
	for _, k := range keys {
		out = append(out, k+"="+m[k])
	}
	return out
}
