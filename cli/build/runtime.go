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

	"pgext/db"
)

type Platform = db.Platform

func DetectPlatform() (Platform, error) {
	p, err := db.DetectPlatform()
	if err != nil {
		return p, err
	}
	if p.OS == "" {
		return Platform{OS: runtime.GOOS, Arch: db.NormalizeArch(runtime.GOARCH)}, nil
	}
	return p, nil
}

func ParsePlatform(release, arch string) (Platform, error) {
	return db.ParseOSRelease(release, arch)
}

func PlanPlatform(value string) (Platform, error) {
	return db.ParsePlatform(value, "")
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
