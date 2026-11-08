package build

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func recipePG(family, root, name string, recipe []byte) []int {
	var raw string
	if family == "deb" {
		if b, err := os.ReadFile(filepath.Join(root, name, "debian", "pgversions")); err == nil {
			raw = string(b)
		}
		if m := regexp.MustCompile(`(?m)^PG_VERSIONS\s*[:?+]?=\s*([0-9 ]+)`).FindSubmatch(recipe); len(m) > 0 {
			raw = string(m[1])
		}
		if raw == "" {
			for pg := 14; pg <= 18; pg++ {
				if makeHasTarget(string(recipe), fmt.Sprintf("build%d", pg)) {
					raw += " " + strconv.Itoa(pg)
				}
			}
		}
	} else if b, err := os.ReadFile(filepath.Join(root, "Makefile")); err == nil {
		pat := regexp.MustCompile(`(?m)^\s*\./build\s+` + regexp.QuoteMeta(name) + `\s+(?:debug\s+)?([0-9 ]+)`)
		if m := pat.FindSubmatch(b); len(m) > 0 {
			raw = string(m[1])
		}
	}
	var out []int
	for _, s := range strings.Fields(raw) {
		n, err := strconv.Atoi(s)
		if err == nil && n >= 14 && n <= 18 && !slices.Contains(out, n) {
			out = append(out, n)
		}
	}
	return out
}

func (e *Engine) dep(ctx context.Context, j *Job, w io.Writer) error {
	if e.Platform.Family == "rpm" {
		pgs := j.PG
		if len(pgs) == 0 {
			pgs = []int{0}
		}
		for _, pg := range pgs {
			args := []string{"dnf", "builddep", "-y", "--define", "_topdir " + e.Options.SpecDir}
			if e.Platform.OS == "el8" {
				args = append(args, "--nobest")
			}
			if pg > 0 {
				args = append(args, "--define", fmt.Sprintf("pgmajorversion %d", pg))
			}
			args = append(args, j.Recipe)
			if err := e.Exec.Run(ctx, Command{Args: args, Root: true}, w); err != nil {
				return err
			}
		}
		return nil
	}
	dir := filepath.Join(filepath.Dir(j.Recipe), "debian")
	b, err := os.ReadFile(filepath.Join(dir, "control.in"))
	if os.IsNotExist(err) {
		b, err = os.ReadFile(filepath.Join(dir, "control"))
	}
	if err != nil {
		return err
	}
	deps := buildDepends(string(b))
	if deps == "" {
		return fmt.Errorf("no Build-Depends in %s", dir)
	}
	if !e.aptUpdated {
		if err := e.Exec.Run(ctx, Command{Args: []string{"apt-get", "update", "-o", "APT::Update::Error-Mode=any"}, Root: true}, w); err != nil {
			return err
		}
		e.aptUpdated = true
	}
	pgs := j.PG
	if len(pgs) == 0 {
		pgs = []int{0}
	}
	for _, pg := range pgs {
		d := strings.ReplaceAll(deps, "PGVERSION", strconv.Itoa(pg))
		d = strings.ReplaceAll(d, "$v", strconv.Itoa(pg))
		// Let APT resolve version constraints, alternatives and architecture
		// restrictions rather than stripping them into a lossy package list.
		if err := e.Exec.Run(ctx, Command{Args: []string{"apt-get", "satisfy", "-y", "--no-install-recommends", d}, Root: true, Env: []string{"DEBIAN_FRONTEND=noninteractive"}}, w); err != nil {
			return err
		}
	}
	if raw, err := os.ReadFile(j.Recipe); err == nil {
		for _, args := range recipeInstallCommands(string(raw)) {
			if err := e.Exec.Run(ctx, Command{Args: args, Root: true}, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildDepends(control string) string {
	fields := map[string]string{}
	current := ""
	for _, line := range strings.Split(control, "\n") {
		if strings.TrimSpace(line) == "" {
			break
		}
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			fields[current] += " " + strings.TrimSpace(line)
			continue
		}
		k, v, ok := strings.Cut(line, ":")
		if ok {
			current = k
			fields[k] = strings.TrimSpace(v)
		}
	}
	var out []string
	for _, k := range []string{"Build-Depends", "Build-Depends-Arch", "Build-Depends-Indep"} {
		if fields[k] != "" {
			out = append(out, fields[k])
		}
	}
	return strings.Join(out, ", ")
}

func (e *Engine) buildEnv(pg int) []string {
	h, _ := os.UserHomeDir()
	path := filepath.Join(h, ".cargo", "bin") + ":" + os.Getenv("PATH")
	if pg > 0 {
		path = filepath.Dir(e.Platform.PGConfig(pg)) + ":" + path
	}
	env := []string{"PATH=" + path, "CARGO_PROFILE_RELEASE_DEBUG=2", "CARGO_PROFILE_RELEASE_STRIP=none"}
	if e.Platform.OS == "el10" {
		env = append(env, "QA_RPATHS=3")
	}
	debOptions := strings.Fields(os.Getenv("DEB_BUILD_OPTIONS"))
	if e.Options.Jobs > 0 {
		env = append(env, fmt.Sprintf("CARGO_BUILD_JOBS=%d", e.Options.Jobs))
		debOptions = slices.DeleteFunc(debOptions, func(s string) bool { return strings.HasPrefix(s, "parallel=") })
		debOptions = append(debOptions, fmt.Sprintf("parallel=%d", e.Options.Jobs))
	}
	if e.Options.NoDebug {
		env = append(env, "CARGO_PROFILE_RELEASE_DEBUG=0", "CARGO_PROFILE_RELEASE_STRIP=debuginfo")
		debOptions = append(debOptions, "noautodbgsym")
	}
	if len(debOptions) > 0 {
		env = append(env, "DEB_BUILD_OPTIONS="+strings.Join(unique(debOptions), " "))
	}
	return env
}

func (e *Engine) make(ctx context.Context, j *Job, r *Result, w io.Writer) error {
	for _, pg := range j.PG {
		if _, err := os.Stat(e.Platform.PGConfig(pg)); err != nil {
			return fmt.Errorf("PG%d development tools missing: %w", pg, err)
		}
	}
	for _, tool := range []string{"make", "gcc", "rustc", "cargo", "rpmbuild", "dpkg-buildpackage"} {
		if out, err := e.Exec.Capture(ctx, Command{Args: []string{tool, "--version"}, Env: e.buildEnv(0)}); err == nil {
			r.ToolVersions[tool] = strings.Split(strings.TrimSpace(out), "\n")[0]
		}
	}
	if err := os.MkdirAll(filepath.Join(e.Options.Root, "tmp"), 0755); err != nil {
		return err
	}
	work, err := os.MkdirTemp(filepath.Join(e.Options.Root, "tmp"), j.Package+"-")
	if err != nil {
		return err
	}
	pgs := j.PG
	if len(pgs) == 0 {
		pgs = []int{0}
	}
	for _, pg := range pgs {
		unit := filepath.Join(work, fmt.Sprintf("pg%d", pg))
		out := filepath.Join(e.Options.Root, "pkg", filepath.Base(work), fmt.Sprintf("pg%d", pg))
		if err := os.MkdirAll(out, 0755); err != nil {
			return err
		}
		var recipe string
		if e.Platform.Family == "rpm" {
			recipe, err = e.makeRPM(ctx, j, pg, unit, out, r, w)
		} else {
			recipe, err = e.makeDEB(ctx, j, pg, unit, out, r, w)
		}
		if err != nil {
			return fmt.Errorf("%s PG%d: %w", j.Package, pg, err)
		}
		artifacts, err := e.verify(ctx, j, pg, recipe, out, w)
		if err != nil {
			return fmt.Errorf("verify %s PG%d: %w", j.Package, pg, err)
		}
		r.Artifacts = append(r.Artifacts, artifacts...)
	}
	return nil
}

func (e *Engine) makeRPM(ctx context.Context, j *Job, pg int, unit, out string, r *Result, w io.Writer) (string, error) {
	root := filepath.Join(unit, "rpmbuild")
	spec := filepath.Join(root, "SPECS", j.Package+".spec")
	if err := copyRecipe(filepath.Join(e.Options.SpecDir, "SPECS"), filepath.Join(root, "SPECS")); err != nil {
		return spec, err
	}
	if err := checkRecipeCopy(j, filepath.Join(root, "SPECS")); err != nil {
		return spec, err
	}
	for _, d := range []string{"BUILD", "BUILDROOT", "SRPMS", "SOURCES"} {
		if err := os.MkdirAll(filepath.Join(root, d), 0755); err != nil {
			return spec, err
		}
	}
	if err := os.Symlink(out, filepath.Join(root, "RPMS")); err != nil {
		return spec, err
	}
	args := []string{"--define", "_topdir " + root}
	if pg > 0 {
		args = append(args, "--define", fmt.Sprintf("pgmajorversion %d", pg))
	}
	if e.Options.Jobs > 0 {
		args = append(args, "--define", fmt.Sprintf("_smp_mflags -j%d", e.Options.Jobs))
	}
	if e.Options.NoDebug {
		args = append(args, "--define", "debug_package %{nil}")
	}
	expanded, err := e.Exec.Capture(ctx, Command{Args: append(append([]string{"rpmspec", "-P"}, args...), spec)})
	if err != nil {
		return spec, err
	}
	for _, line := range strings.Split(expanded, "\n") {
		if !regexp.MustCompile(`^(Source|Patch)[0-9]*\s*:`).MatchString(line) {
			continue
		}
		kind, raw, _ := strings.Cut(line, ":")
		name, err := rpmSourceName(raw)
		if err != nil {
			return spec, err
		}
		from := filepath.Join(e.Options.Root, "src", name)
		if strings.HasPrefix(kind, "Patch") {
			from = filepath.Join(root, "SPECS", "patches", name)
		}
		staged := filepath.Join(root, "SOURCES", name)
		if err := copyFile(from, staged, 0644); err != nil && !os.IsExist(err) {
			return spec, err
		}
		if !strings.HasPrefix(kind, "Patch") {
			if err := e.checkSource(staged); err != nil {
				return spec, err
			}
		}
	}
	if err := recordSources(unit, r); err != nil {
		return spec, err
	}
	cmd := Command{Args: append(append([]string{"rpmbuild", "-ba"}, args...), spec), Dir: root, Env: e.buildEnv(pg)}
	return spec, e.Exec.Run(ctx, cmd, w)
}

func (e *Engine) makeDEB(ctx context.Context, j *Job, pg int, unit, out string, r *Result, w io.Writer) (string, error) {
	root := filepath.Join(unit, "debbuild")
	recipe := filepath.Join(root, j.Package)
	if err := copyRecipe(filepath.Dir(j.Recipe), recipe); err != nil {
		return recipe, err
	}
	if err := checkRecipeCopy(j, recipe); err != nil {
		return recipe, err
	}
	if err := os.MkdirAll(filepath.Join(root, "SOURCES"), 0755); err != nil {
		return recipe, err
	}
	for _, name := range j.Sources {
		from := filepath.Join(e.Options.Root, "src", name)
		if _, err := os.Stat(from); os.IsNotExist(err) {
			continue
		}
		staged := filepath.Join(root, "SOURCES", name)
		if err := copyFile(from, staged, 0644); err != nil {
			return recipe, err
		}
		if err := e.checkSource(staged); err != nil {
			return recipe, err
		}
	}
	// Shared kernel version files are recipe inputs, not source caches.
	for _, dir := range []string{"pgedge"} {
		if dir == j.Package {
			continue
		}
		from := filepath.Join(e.Options.SpecDir, dir, "versions.mk")
		if _, err := os.Stat(from); err == nil {
			if err := copyFile(from, filepath.Join(root, dir, "versions.mk"), 0644); err != nil {
				return recipe, err
			}
		}
	}
	b, err := os.ReadFile(filepath.Join(recipe, "Makefile"))
	if err != nil {
		return recipe, err
	}
	text := strings.ReplaceAll(string(b), "~/ext/pkg", `"$(PKG_OUTPUT_DIR)"`)
	text = withoutRecipeInstalls(text)
	if err := AtomicWrite(filepath.Join(recipe, "Makefile"), []byte(text), 0644); err != nil {
		return recipe, err
	}
	env := e.buildEnv(pg)
	args := []string{"make", "--no-print-directory", "-j1", "PKG_OUTPUT_DIR=" + out}
	if regexp.MustCompile(`(?m)^SRC_DIR\s*[:?]?=`).MatchString(text) {
		args = append(args, "SRC_DIR="+filepath.Join(unit, "src"))
	}
	if pg > 0 {
		v := strconv.Itoa(pg)
		args = append(args, "PG_VERSIONS="+v, "PG_MAJOR="+v)
		env = append(env, "PG_VERSIONS="+v, "PG_SUPPORTED_VERSIONS="+v, "PG_CONFIG="+e.Platform.PGConfig(pg))
		pv := filepath.Join(recipe, "debian", "pgversions")
		if _, err := os.Stat(pv); err == nil {
			if err := AtomicWrite(pv, []byte(v+"\n"), 0644); err != nil {
				return recipe, err
			}
		}
		if makeHasTarget(text, "build"+v) {
			if makeHasTarget(text, "pre") {
				args = append(args, "pre")
			}
			if makeHasTarget(text, "setup") {
				args = append(args, "setup")
			}
			args = append(args, "build"+v)
		}
	}
	if err := recordSources(unit, r); err != nil {
		return recipe, err
	}
	if err := e.Exec.Run(ctx, Command{Args: args, Dir: recipe, Env: env}, w); err != nil {
		return recipe, err
	}
	// Isolated inputs contain no old packages; collect every package, including
	// dbgsym and companions, regardless of a legacy recipe's move target.
	err = filepath.WalkDir(unit, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "target" || d.Name() == ".git" || d.Name() == "debian" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".deb") || strings.HasSuffix(path, ".ddeb") {
			to := filepath.Join(out, filepath.Base(path))
			if _, err := os.Stat(to); os.IsNotExist(err) {
				return copyFile(path, to, 0644)
			}
		}
		return nil
	})
	return recipe, err
}

// Some historical recipes install pre-build OS packages in their pre target.
// Move only literal package-manager commands to dep; retain the remaining pre
// build work (for example compiling PLV8's bundled V8).
var installLine = regexp.MustCompile(`^\t(?:sudo\s+)?(?:apt-get|apt)\s+(install|update)\b(.*)$`)

func literalAPTCommand(line string) []string {
	m := installLine.FindStringSubmatch(line)
	if len(m) == 0 {
		return nil
	}
	args := strings.Fields(strings.Split(m[2], "#")[0])
	for _, arg := range args {
		if strings.ContainsAny(arg, "$;|&`\\\"'<>*?(){}[]~") {
			return nil
		}
	}
	return append([]string{"apt-get", m[1]}, args...)
}

func recipeInstallCommands(content string) [][]string {
	var out [][]string
	for _, line := range strings.Split(content, "\n") {
		if args := literalAPTCommand(line); len(args) > 0 && args[1] == "install" {
			out = append(out, args)
		}
	}
	return out
}
func withoutRecipeInstalls(content string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if len(literalAPTCommand(line)) > 0 {
			lines[i] = "\t@true # OS packages are prepared by pgext dep"
		}
	}
	return strings.Join(lines, "\n")
}

func (e *Engine) checkSource(path string) error {
	f, err := fileInfo(path)
	if err != nil {
		return err
	}
	if f.Size == 0 {
		return fmt.Errorf("empty source: %s", path)
	}
	if e.Options.Checksums == "" {
		return nil
	}
	sums, err := loadChecksums(e.Options.Checksums)
	if err != nil {
		return err
	}
	want := sums[filepath.Base(path)]
	if want == "" {
		return fmt.Errorf("checksum manifest lacks %s", filepath.Base(path))
	}
	if f.SHA256 != want {
		return fmt.Errorf("source checksum mismatch: %s", path)
	}
	return nil
}
func recordSources(unit string, r *Result) error {
	root := filepath.Join(unit, "rpmbuild", "SOURCES")
	entries, err := os.ReadDir(root)
	if os.IsNotExist(err) {
		root = filepath.Join(unit, "debbuild", "SOURCES")
		entries, err = os.ReadDir(root)
	}
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		f, err := fileInfo(filepath.Join(root, entry.Name()))
		if err != nil {
			return err
		}
		r.Sources = append(r.Sources, f)
	}
	return nil
}

func makeHasTarget(content, target string) bool {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "\t") {
			continue
		}
		head, _, ok := strings.Cut(line, ":")
		if ok && slices.Contains(strings.Fields(head), target) {
			return true
		}
	}
	return false
}

func checkRecipeCopy(j *Job, root string) error {
	if j.RecipeTreeSHA256 == "" {
		return nil
	}
	digest, err := recipeTreeHash(root)
	if err != nil {
		return err
	}
	if digest != j.RecipeTreeSHA256 {
		return fmt.Errorf("recipe changed while staging: %s", j.Recipe)
	}
	return nil
}
