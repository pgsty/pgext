package build

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//go:embed assets/verify-payload.sh
var payloadVerifier []byte

func (e *Engine) verify(ctx context.Context, j *Job, pg int, recipe, out string, w io.Writer) ([]Artifact, error) {
	var paths []string
	err := filepath.WalkDir(out, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (strings.HasSuffix(p, ".deb") || strings.HasSuffix(p, ".ddeb") || strings.HasSuffix(p, ".rpm")) {
			paths = append(paths, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("build produced no packages")
	}
	expected, err := e.expectedPackages(ctx, pg, recipe)
	if err != nil {
		return nil, err
	}
	seen := map[string]Artifact{}
	var result []Artifact
	for _, p := range paths {
		f, err := fileInfo(p)
		if err != nil {
			return nil, err
		}
		if f.Size == 0 {
			return nil, fmt.Errorf("empty package %s", p)
		}
		var args []string
		if e.Platform.Family == "rpm" {
			args = []string{"rpm", "-qp", "--qf", "%{NAME}\t%{EPOCHNUM}:%{VERSION}-%{RELEASE}\t%{ARCH}", p}
		} else {
			args = []string{"dpkg-deb", "-W", "--showformat=${Package}\t${Version}\t${Architecture}", p}
		}
		metadata, err := e.Exec.Capture(ctx, Command{Args: args})
		if err != nil {
			return nil, err
		}
		parts := strings.Split(strings.TrimSpace(metadata), "\t")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid package metadata: %s", p)
		}
		a := Artifact{File: f, Package: parts[0], Version: parts[1], Arch: parts[2]}
		if e.Platform.Family == "deb" && pg > 0 {
			if m := regexp.MustCompile(`^postgresql-([0-9]+)-`).FindStringSubmatch(a.Package); len(m) > 0 && m[1] != strconv.Itoa(pg) {
				return nil, fmt.Errorf("recipe ignored PG%d selection and produced %s", pg, a.Package)
			}
		}
		if a.Arch != e.Platform.NativeArch() && a.Arch != "noarch" && a.Arch != "all" {
			return nil, fmt.Errorf("unexpected architecture %s: %s", a.Arch, p)
		}
		if _, ok := seen[a.Package]; ok {
			return nil, fmt.Errorf("duplicate binary package %s", a.Package)
		}
		version, ok := expected[a.Package]
		if !ok && strings.HasSuffix(a.Package, "-dbgsym") {
			version, ok = expected[strings.TrimSuffix(a.Package, "-dbgsym")]
		}
		if !ok {
			return nil, fmt.Errorf("unexpected binary package %s", a.Package)
		}
		if version != a.Version {
			return nil, fmt.Errorf("%s version %s differs from recipe %s", a.Package, a.Version, version)
		}
		seen[a.Package] = a
		result = append(result, a)
	}
	for name := range expected {
		if _, ok := seen[name]; !ok {
			return nil, fmt.Errorf("missing expected package %s", name)
		}
	}
	if j.Extension != nil && pg > 0 {
		template := j.Extension.DEB
		if e.Platform.Family == "rpm" {
			template = j.Extension.RPM
		}
		primary := strings.ReplaceAll(template, "$v", strconv.Itoa(pg))
		if primary != "" {
			if _, ok := seen[primary]; !ok {
				return nil, fmt.Errorf("catalog main package %s not produced", primary)
			}
		}
	}
	script := filepath.Join(filepath.Dir(out), "verify-payload.sh")
	if err := AtomicWrite(script, payloadVerifier, 0700); err != nil {
		return nil, err
	}
	args := []string{"bash", script, e.Platform.Family, out, strconv.FormatBool(e.Options.NoDebug)}
	if err := e.Exec.Run(ctx, Command{Args: args}, w); err != nil {
		return nil, err
	}
	// Check the existing RPM LLVM policy only when this recipe declares it.
	if e.Platform.Family == "rpm" && pg > 0 && j.Extension != nil {
		b, err := os.ReadFile(recipe)
		if err != nil {
			return nil, err
		}
		for _, a := range result {
			if strings.Contains(a.Package, "-llvmjit") {
				return nil, fmt.Errorf("extension llvmjit subpackage is forbidden: %s", a.Package)
			}
			if strings.HasSuffix(a.Package, "-debuginfo") || strings.HasSuffix(a.Package, "-debugsource") {
				continue
			}
			deps, err := e.Exec.Capture(ctx, Command{Args: []string{"rpm", "-qp", "--requires", a.Path}})
			if err != nil {
				return nil, err
			}
			if regexp.MustCompile(`(?im)(^|[/\s])(llvm|clang|libLLVM|libclang)|postgresql\d+-llvmjit`).MatchString(deps) {
				return nil, fmt.Errorf("forbidden LLVM runtime dependency: %s", a.Package)
			}
			primary := strings.ReplaceAll(j.Extension.RPM, "$v", strconv.Itoa(pg))
			if a.Package == primary && strings.Contains(string(b), "%{!?llvm:%global llvm 1}") {
				payload, err := e.Exec.Capture(ctx, Command{Args: []string{"rpm", "-qpl", a.Path}})
				if err != nil {
					return nil, err
				}
				if !strings.Contains(payload, "/bitcode/") || !strings.Contains(payload, ".index.bc") {
					return nil, fmt.Errorf("main package lacks LLVM bitcode/index: %s", primary)
				}
			}
		}
	}
	return result, nil
}

func (e *Engine) expectedPackages(ctx context.Context, pg int, recipe string) (map[string]string, error) {
	expected := map[string]string{}
	if e.Platform.Family == "rpm" {
		args := []string{"rpmspec", "-q", "--qf", "%{NAME}\t%{EPOCHNUM}:%{VERSION}-%{RELEASE}\n", "--define", "_topdir " + filepath.Dir(filepath.Dir(recipe))}
		if pg > 0 {
			args = append(args, "--define", fmt.Sprintf("pgmajorversion %d", pg))
		}
		if e.Options.NoDebug {
			args = append(args, "--define", "debug_package %{nil}")
		}
		s, err := e.Exec.Capture(ctx, Command{Args: append(args, recipe)})
		if err != nil {
			return nil, err
		}
		for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
			parts := strings.Split(line, "\t")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid rpmspec package query")
			}
			expected[parts[0]] = parts[1]
		}
	} else {
		var controls []string
		// Native recipes place generated packaging in build/debian,
		// build-18/build/debian, build/columnar/debian, or an isolated SRC_DIR.
		unit := filepath.Dir(filepath.Dir(recipe))
		err := filepath.WalkDir(unit, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if d.Name() == "target" || d.Name() == ".git" || d.Name() == "SOURCES" {
					return filepath.SkipDir
				}
				return nil
			}
			if d.Name() == "control" && filepath.Base(filepath.Dir(path)) == "debian" && path != filepath.Join(recipe, "debian", "control") {
				if _, err := os.Stat(filepath.Join(filepath.Dir(path), "changelog")); err == nil {
					controls = append(controls, path)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		if len(controls) == 0 {
			if p := filepath.Join(recipe, "debian", "control"); fileExists(p) {
				controls = append(controls, p)
			}
		}
		if len(controls) == 0 {
			return nil, fmt.Errorf("no generated Debian control in %s", recipe)
		}
		if len(controls) > 1 {
			return nil, fmt.Errorf("ambiguous generated Debian controls: %v", controls)
		}
		control := controls[0]
		version, err := e.Exec.Capture(ctx, Command{Args: []string{"dpkg-parsechangelog", "-l" + filepath.Join(filepath.Dir(control), "changelog"), "-SVersion"}})
		if err != nil {
			return nil, err
		}
		b, err := os.ReadFile(control)
		if err != nil {
			return nil, err
		}
		for _, m := range regexp.MustCompile(`(?m)^Package:\s*(\S+)`).FindAllSubmatch(b, -1) {
			name := string(m[1])
			if strings.Contains(name, "PGVERSION") {
				return nil, fmt.Errorf("unexpanded package name %s", name)
			}
			expected[name] = strings.TrimSpace(version)
		}
	}
	if len(expected) == 0 {
		return nil, fmt.Errorf("recipe declares no binary packages")
	}
	return expected, nil
}
