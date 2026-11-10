package build

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

type fakeExecutor struct {
	commands []Command
	run      func(Command) error
	capture  func(Command) (string, error)
}

func (f *fakeExecutor) Run(_ context.Context, c Command, _ io.Writer) error {
	f.commands = append(f.commands, c)
	if f.run != nil {
		return f.run(c)
	}
	return nil
}
func (f *fakeExecutor) Capture(_ context.Context, c Command) (string, error) {
	if f.capture != nil {
		return f.capture(c)
	}
	return "", errors.New("not installed")
}
func put(t *testing.T, p, s string) {
	t.Helper()
	if err := AtomicWrite(p, []byte(s), 0644); err != nil {
		t.Fatal(err)
	}
}
func testEngine(t *testing.T) *Engine {
	t.Helper()
	c, err := ParseCatalog([]byte("name,pkg,lead_ext,source,pg_ver,rpm_pkg,deb_pkg\nvector,pgvector,vector,pgvector-1.0.tar.gz,\"{18,17}\",pgvector_$v,postgresql-$v-pgvector\n"))
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	return &Engine{Options: Options{Root: filepath.Join(dir, "ext"), SpecDir: filepath.Join(dir, "debbuild"), PG: "18"}, Platform: Platform{OS: "u24", Arch: "arm64", Family: "deb", Codename: "noble"}, Catalog: c, Exec: &fakeExecutor{}, Err: io.Discard}
}

func TestEmbeddedCatalogAndMatrixIndependence(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	c, err := LoadCatalog("")
	if err != nil {
		t.Fatal(err)
	}
	if c.Resolve("vector").Package != "pgvector" || c.Resolve("citus_columnar").Package != "citus" {
		t.Fatal("family resolution")
	}
	for _, extra := range []string{`{}`, `{"matrix":["u24a:18:N:f:0::"]}`} {
		row := "name,pkg,source,extra\nx,x,x-1.tar.gz,\"" + strings.ReplaceAll(extra, `"`, `""`) + "\"\n"
		c, err := ParseCatalog([]byte(row))
		if err != nil || c.Resolve("x") == nil {
			t.Fatalf("matrix must not gate source builds: %v", err)
		}
	}
	for _, b := range []string{"name,pkg,source\nx,x,x\nx,y,y\n", "name,source\nx,x\n", "name,pkg,source\n../x,x,x\n"} {
		if _, err := ParseCatalog([]byte(b)); err == nil {
			t.Fatalf("accepted invalid catalog %q", b)
		}
	}
	dir, _ := ConfigDir()
	put(t, filepath.Join(dir, "extension.csv"), "broken")
	if _, err := LoadCatalog(""); err == nil {
		t.Fatal("invalid local catalog silently fell back")
	}
	if _, err := LoadCatalog(filepath.Join(t.TempDir(), "absent")); err == nil {
		t.Fatal("explicit missing catalog fell back")
	}
}

func TestStageBoundariesAndFamilyDedup(t *testing.T) {
	e := testEngine(t)
	put(t, filepath.Join(e.Options.SpecDir, "pgvector", "Makefile"), "all: setup build move\nPG_VERSIONS := 17 18\nTARBALL := pgvector-1.0.tar.gz\n")
	for action, want := range map[string][]string{"get": {"get"}, "dep": {"dep"}, "make": {"make"}, "pkg": {"get", "dep", "make"}} {
		p, err := e.Prepare(action, []string{"vector", "pgvector"})
		if err != nil {
			t.Fatal(err)
		}
		if len(p.Jobs) != 1 || !reflect.DeepEqual(p.Jobs[0].Stages, want) || !reflect.DeepEqual(p.Jobs[0].PG, []int{18}) {
			t.Fatalf("bad plan: %+v", p)
		}
	}
	e.Options.PG = "16"
	if _, err := e.Prepare("make", []string{"vector"}); err == nil {
		t.Fatal("unsupported target accepted")
	}
	for _, pg := range []string{"19", "13", "18,", "17;touch /tmp/oops"} {
		if _, err := ParsePG(pg); err == nil {
			t.Fatalf("accepted %s", pg)
		}
	}
}

func TestDebianSupportIsNotLimitedByMakeDefault(t *testing.T) {
	for _, tc := range []struct {
		name, versions, makeDefault, requested string
		wantError                              bool
	}{
		{"explicit supported major", "14\n15\n16\n17\n18\n", "18", "14", false},
		{"unsupported major", "17\n18\n", "14 15 16 17 18", "16", true},
		{"missing file fallback", "", "17 18", "17", false},
		{"fallback still constrains support", "", "17 18", "16", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := testEngine(t)
			e.Options.PG = tc.requested
			root := filepath.Join(e.Options.SpecDir, "pgvector")
			put(t, filepath.Join(root, "Makefile"), "PG_VERSIONS ?= "+tc.makeDefault+"\nTARBALL := pgvector-1.0.tar.gz\n")
			if tc.versions != "" {
				put(t, filepath.Join(root, "debian", "pgversions"), tc.versions)
			}
			p, err := e.Prepare("get", []string{"vector"})
			if (err != nil) != tc.wantError {
				t.Fatalf("Prepare() error = %v, wantError %v", err, tc.wantError)
			}
			if err == nil && (len(p.Jobs) != 1 || fmt.Sprint(p.Jobs[0].PG) != "["+tc.requested+"]") {
				t.Fatalf("requested major was lost: %+v", p.Jobs)
			}
		})
	}
}

func TestPipelineStopsOnDependencyFailure(t *testing.T) {
	e := testEngine(t)
	put(t, filepath.Join(e.Options.SpecDir, "pgvector", "Makefile"), "PG_VERSIONS := 18\nTARBALL := pgvector-1.0.tar.gz\n")
	put(t, filepath.Join(e.Options.SpecDir, "pgvector", "debian", "control.in"), "Source: pgvector\nBuild-Depends: make (>= 1), gcc | clang\n\nPackage: postgresql-PGVERSION-pgvector\n")
	put(t, filepath.Join(e.Options.Root, "src", "pgvector-1.0.tar.gz"), "already downloaded")
	f := &fakeExecutor{run: func(c Command) error {
		if c.Args[0] != "apt-get" {
			t.Fatalf("unexpected stage: %v", c.Args)
		}
		if c.Args[1] == "update" {
			return nil
		}
		return errors.New("dependency unavailable")
	}}
	e.Exec = f
	p, err := e.Prepare("pkg", []string{"vector"})
	if err != nil {
		t.Fatal(err)
	}
	r, err := e.Run(context.Background(), p)
	if err == nil || r.Success || len(r.Stages) != 2 || r.Stages[1].Stage != "dep" || len(f.commands) != 2 {
		t.Fatalf("failed to stop: %+v %v", r, err)
	}
	if b, err := os.ReadFile(r.Manifest); err != nil || !strings.Contains(string(b), "dependency unavailable") {
		t.Fatalf("failed manifest: %v", err)
	}
}

func TestDebianTargetActuallyReachesRecipe(t *testing.T) {
	for _, special := range []bool{false, true} {
		t.Run(map[bool]string{true: "numbered", false: "pg_buildext"}[special], func(t *testing.T) {
			e := testEngine(t)
			mk := "all: setup build move\nPG_VERSIONS := 17 18\n"
			if special {
				mk += "build18 build17: setup\n"
			}
			p := filepath.Join(e.Options.SpecDir, "pgvector", "Makefile")
			put(t, p, mk)
			put(t, filepath.Join(filepath.Dir(p), "debian", "pgversions"), "17\n18\n")
			out := filepath.Join(t.TempDir(), "output")
			os.MkdirAll(out, 0755)
			f := &fakeExecutor{run: func(c Command) error {
				joined := strings.Join(c.Args, " ")
				if !strings.Contains(joined, "PG_VERSIONS=17") || !strings.Contains(joined, "PG_MAJOR=17") {
					t.Fatalf("PG not passed: %s", joined)
				}
				if special && !strings.Contains(joined, "build17") {
					t.Fatalf("numbered target not selected: %s", joined)
				}
				b, err := os.ReadFile(filepath.Join(c.Dir, "debian", "pgversions"))
				if err != nil || string(b) != "17\n" {
					t.Fatalf("unfiltered pgversions: %s, %v", b, err)
				}
				if !strings.Contains(strings.Join(c.Env, " "), "PG_SUPPORTED_VERSIONS=17") {
					t.Fatal("pg_buildext not constrained")
				}
				return nil
			}}
			e.Exec = f
			_, err := e.makeDEB(context.Background(), &Job{Package: "pgvector", Recipe: p}, 17, t.TempDir(), out, &Result{}, io.Discard)
			if err != nil {
				t.Fatal(err)
			}
			b, _ := os.ReadFile(filepath.Join(filepath.Dir(p), "debian", "pgversions"))
			if string(b) != "17\n18\n" {
				t.Fatal("authoritative recipe changed")
			}
		})
	}
}

func TestDownloadVerifiesCacheAndReplacesAtomically(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "new bytes") }))
	defer srv.Close()
	p := filepath.Join(t.TempDir(), "source.tar.gz")
	put(t, p, "old bytes")
	if _, err := Download(context.Background(), []string{srv.URL}, "source.tar.gz", p, Hash([]byte("new bytes")), false, io.Discard); err == nil {
		t.Fatal("unverified cache accepted")
	}
	if _, err := Download(context.Background(), []string{srv.URL}, "source.tar.gz", p, Hash([]byte("wrong")), true, io.Discard); err == nil {
		t.Fatal("wrong download accepted")
	}
	b, _ := os.ReadFile(p)
	if string(b) != "old bytes" {
		t.Fatal("failed fetch damaged old source")
	}
	f, err := Download(context.Background(), []string{srv.URL}, "source.tar.gz", p, Hash([]byte("new bytes")), true, io.Discard)
	if err != nil || f.SHA256 != Hash([]byte("new bytes")) {
		t.Fatal(f, err)
	}
}

func TestArchiveCannotEscapeOrInjectLinks(t *testing.T) {
	for _, name := range []string{"../outside", "rpmbuild/../../outside", "rpmbuild/link"} {
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		tw := tar.NewWriter(gz)
		kind := byte(tar.TypeReg)
		if strings.HasSuffix(name, "link") {
			kind = tar.TypeSymlink
		}
		tw.WriteHeader(&tar.Header{Name: name, Mode: 0644, Typeflag: kind, Linkname: "/etc"})
		tw.Close()
		gz.Close()
		p := filepath.Join(t.TempDir(), "recipes.tar.gz")
		os.WriteFile(p, b.Bytes(), 0600)
		if err := extractRecipes(p, t.TempDir(), "rpmbuild"); err == nil {
			t.Fatalf("unsafe archive accepted %s", name)
		}
	}
}

func TestAllSupportedRepoAndToolPlans(t *testing.T) {
	for code := range supportedOS {
		for _, arch := range []string{"amd64", "arm64"} {
			p, err := PlanPlatform(code + "." + arch)
			if err != nil {
				t.Fatal(err)
			}
			if code == "el9" {
				p.Version = "9.8"
			}
			if code == "el10" {
				p.Version = "10.2"
			}
			e := testEngine(t)
			e.Platform = p
			e.Options.Plan = true
			for _, m := range []string{"pgdg", "pigsty"} {
				content, err := e.repoContent(m)
				if err != nil {
					t.Fatal(err)
				}
				if strings.Contains(content, "trusted=yes") || strings.Contains(content, "gpgcheck=0") {
					t.Fatal("signature verification disabled")
				}
			}
			cmd, err := e.toolCommand("full")
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(strings.Join(cmd.Args, " "), "postgresql19") {
				t.Fatal("implicit beta")
			}
			if _, err := e.Repo(context.Background(), "add", []string{"unknown"}, true); err == nil {
				t.Fatal("unknown module accepted")
			}
		}
	}
}

func TestNativeCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	start := time.Now()
	err := (NativeExecutor{}).Run(ctx, Command{Args: []string{"sh", "-c", "sleep 30 & wait"}}, io.Discard)
	if err == nil || time.Since(start) > 3*time.Second {
		t.Fatalf("process tree did not cancel: %v", err)
	}
}

func TestRecipeSourceVariablesAndGroupedTargets(t *testing.T) {
	content := "VERSION := 2.31.0\nTARBALL := timescaledb-$(VERSION).tar.gz\nSOURCE_TARBALL ?= ../SOURCES/$(TARBALL)\nbuild18 build17 build16:\n\ttar -xf $(SOURCE_TARBALL)\n"
	if got := recipeSources(content); !reflect.DeepEqual(got, []string{"timescaledb-2.31.0.tar.gz"}) {
		t.Fatal(got)
	}
	if got := recipePG("deb", t.TempDir(), "x", []byte(content)); !reflect.DeepEqual(got, []int{16, 17, 18}) {
		t.Fatal(got)
	}
	conditional := "TARBALL_OLD := citus-13.5.0.tar.gz\nTARBALL_NEW := citus-14.3.0.tar.gz\nifeq ($(PG_MAJOR),15)\nTARBALL := $(TARBALL_OLD)\nelse\nTARBALL := $(TARBALL_NEW)\nendif\n"
	if got := recipeSources(conditional); !reflect.DeepEqual(got, []string{"citus-13.5.0.tar.gz", "citus-14.3.0.tar.gz"}) {
		t.Fatal(got)
	}
}

func TestRPMGetIncludesAllRecipeMajorSources(t *testing.T) {
	e := testEngine(t)
	e.Options.PG = ""
	e.Platform = Platform{OS: "el9", Family: "rpm", Arch: "arm64"}
	put(t, filepath.Join(e.Options.SpecDir, "SPECS", "pgvector.spec"), "Name: pgvector_%{pgmajorversion}\n")
	put(t, filepath.Join(e.Options.SpecDir, "Makefile"), "pgvector:\n\t./build pgvector 17 18\n")
	p, err := e.Prepare("get", []string{"pgvector"})
	if err != nil || !reflect.DeepEqual(p.Jobs[0].PG, []int{17, 18}) {
		t.Fatal(p, err)
	}
	put(t, filepath.Join(e.Options.Root, "src", "old.tar.gz"), "old")
	put(t, filepath.Join(e.Options.Root, "src", "new.tar.gz"), "new")
	e.Exec = &fakeExecutor{capture: func(c Command) (string, error) {
		if strings.Contains(strings.Join(c.Args, " "), "pgmajorversion 17") {
			return "Source0: old.tar.gz\n", nil
		}
		return "Source0: new.tar.gz\n", nil
	}}
	r, err := e.Run(context.Background(), p)
	if err != nil || len(r.Sources) != 2 {
		t.Fatal(r, err)
	}
}

func TestMakeDoesNotInstallRecipeDependencies(t *testing.T) {
	e := testEngine(t)
	p := filepath.Join(e.Options.SpecDir, "pgvector", "Makefile")
	content := "all: pre setup build\npre:\n\tapt-get update -qq\n\tapt-get install -y cmake libicu-dev\n\t./compile-bundled-library\nsetup:\n\ttar -xf ../SOURCES/vector.tar.gz\nbuild18:\n\t./build-native\n"
	put(t, p, content)
	put(t, filepath.Join(e.Options.Root, "src", "vector.tar.gz"), "input archive")
	f := &fakeExecutor{run: func(c Command) error {
		if c.Args[0] != "make" || !strings.Contains(strings.Join(c.Args, " "), "setup build18") {
			t.Fatalf("unexpected build operation: %v", c.Args)
		}
		b, _ := os.ReadFile(filepath.Join(c.Dir, "Makefile"))
		if strings.Contains(string(b), "apt-get") || !strings.Contains(string(b), "./compile-bundled-library") {
			t.Fatalf("bad staged recipe: %s", b)
		}
		return nil
	}}
	e.Exec = f
	_, err := e.makeDEB(context.Background(), &Job{Package: "pgvector", Recipe: p, Sources: []string{"vector.tar.gz"}}, 18, t.TempDir(), t.TempDir(), &Result{}, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(p); string(b) != content {
		t.Fatal("source recipe changed")
	}
	if got := recipeInstallCommands(content); !reflect.DeepEqual(got, [][]string{{"apt-get", "install", "-y", "cmake", "libicu-dev"}}) {
		t.Fatal(got)
	}
}

func TestVerificationRejectsIncompleteOrWrongPackages(t *testing.T) {
	for _, tc := range []struct {
		name, control, packages, problem string
	}{
		{"missing-companion", "main companion", "main:arm64", "missing expected package"},
		{"wrong-architecture", "main", "main:amd64", "unexpected architecture"},
		{"unrequested-major", "postgresql-18-x postgresql-19-x", "postgresql-18-x:arm64 postgresql-19-x:arm64", "ignored PG18"},
		{"duplicate", "main", "main:arm64 main:arm64", "duplicate binary package"},
		{"unexpected", "main", "foreign:arm64", "unexpected binary package"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := testEngine(t)
			recipe := filepath.Join(t.TempDir(), "debbuild", "test")
			var control strings.Builder
			for _, name := range strings.Fields(tc.control) {
				fmt.Fprintf(&control, "Package: %s\nArchitecture: any\n\n", name)
			}
			put(t, filepath.Join(recipe, "build", "debian", "control"), control.String())
			put(t, filepath.Join(recipe, "build", "debian", "changelog"), "test changelog")
			out := t.TempDir()
			metadata := map[string]string{}
			for i, value := range strings.Fields(tc.packages) {
				name, arch, _ := strings.Cut(value, ":")
				path := filepath.Join(out, fmt.Sprintf("%02d.deb", i))
				put(t, path, "test package bytes")
				metadata[path] = name + "\t1.0-1PGSTY\t" + arch
			}
			e.Exec = &fakeExecutor{capture: func(c Command) (string, error) {
				if c.Args[0] == "dpkg-parsechangelog" {
					return "1.0-1PGSTY", nil
				}
				return metadata[c.Args[len(c.Args)-1]], nil
			}, run: func(c Command) error {
				t.Fatal("invalid package set reached payload verification")
				return nil
			}}
			_, err := e.verify(context.Background(), &Job{}, 18, recipe, out, io.Discard)
			if err == nil || !strings.Contains(err.Error(), tc.problem) {
				t.Fatalf("wanted %s, got %v", tc.problem, err)
			}
		})
	}
}

func TestRecipeDriftStillSavesFailure(t *testing.T) {
	e := testEngine(t)
	p := filepath.Join(e.Options.SpecDir, "pgvector", "Makefile")
	put(t, p, "PG_VERSIONS := 18\n")
	plan, err := e.Prepare("make", []string{"vector"})
	if err != nil {
		t.Fatal(err)
	}
	put(t, p, "PG_VERSIONS := 17\n")
	r, err := e.Run(context.Background(), plan)
	if err == nil || r.Error == "" || r.Success {
		t.Fatal(r, err)
	}
	if b, err := os.ReadFile(r.Manifest); err != nil || !strings.Contains(string(b), "recipe changed after planning") {
		t.Fatal(string(b), err)
	}
}

func TestLegacyRepoRecognition(t *testing.T) {
	for _, tc := range []struct {
		family, content string
		want            bool
	}{
		{"deb", "# pigsty-infra\ndeb [signed-by=/etc/apt/keyrings/pigsty.asc trusted=yes] https://repo.pigsty.cc/apt/infra/ generic main\ndeb [trusted=yes] https://apt.postgresql.org/pub/repos/apt/ noble-pgdg main\n", true},
		{"deb", "deb https://example.com/custom stable main", false},
		{"deb", "deb https://repo.pigsty.cc/apt/infra generic main\ndeb https://example.com/custom stable main", false},
		{"rpm", "[pigsty-pgsql]\nbaseurl=https://repo.pigsty.cc/yum/pgsql/el9.$basearch\n[pgdg18]\nbaseurl=https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-9.8-$basearch\nenabled=1", true},
		{"rpm", "[pgdg18]\nenabled=1\n[custom]\nenabled=1", false},
	} {
		if got := recognizedLegacyRepos(tc.content, tc.family); got != tc.want {
			t.Fatalf("%s: %t", tc.content, got)
		}
	}
}

func TestRecipeCopyPreservesHelpersAndOmitsArtifacts(t *testing.T) {
	src, dst := t.TempDir(), filepath.Join(t.TempDir(), "recipes")
	put(t, filepath.Join(src, "build"), "#!/bin/sh\necho build helper\n")
	put(t, filepath.Join(src, "extension", "debian", "rules"), "recipe")
	put(t, filepath.Join(src, "extension", "build-18", "binary"), "old build")
	put(t, filepath.Join(src, "extension", "old.deb"), "old package")
	put(t, filepath.Join(src, "SOURCES", "source.tar.gz"), "archive")
	if err := copyRecipe(src, dst); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"build", "extension/debian/rules"} {
		if !fileExists(filepath.Join(dst, path)) {
			t.Fatalf("missing recipe input: %s", path)
		}
	}
	for _, path := range []string{"extension/build-18", "extension/old.deb", "SOURCES"} {
		if _, err := os.Stat(filepath.Join(dst, path)); !os.IsNotExist(err) {
			t.Fatalf("copied old artifacts: %s", path)
		}
	}
}

func TestNativeCaptureKeepsMetadataSeparate(t *testing.T) {
	s, err := (NativeExecutor{}).Capture(context.Background(), Command{Args: []string{"sh", "-c", "echo warning >&2; printf 'package\\tversion\\tarch'"}})
	if err != nil || s != "package\tversion\tarch" {
		t.Fatal(s, err)
	}
}

func TestNativeToolUsesBuildPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "pgext-test-tool")
	put(t, path, "#!/bin/sh\necho selected-toolchain\n")
	if err := os.Chmod(path, 0755); err != nil {
		t.Fatal(err)
	}
	s, err := (NativeExecutor{}).Capture(context.Background(), Command{Args: []string{"pgext-test-tool", "--version"}, Env: []string{"PATH=" + dir}})
	if err != nil || strings.TrimSpace(s) != "selected-toolchain" {
		t.Fatal(s, err)
	}
}
