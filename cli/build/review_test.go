package build

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestEnvironmentDoesNotDependOnCatalog(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	dir, _ := ConfigDir()
	put(t, filepath.Join(dir, "extension.csv"), "broken catalog")
	e, err := NewEnvironment(Options{Plan: true, Platform: "u24.arm64"}, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	f := &fakeExecutor{capture: func(Command) (string, error) {
		t.Fatal("environment plan executed a native command")
		return "", nil
	}}
	e.Exec = f
	for _, action := range []string{"status", "tool", "spec", "rust", "pgrx"} {
		if _, err := e.Environment(context.Background(), action, EnvOptions{}); err != nil {
			t.Fatalf("%s: %v", action, err)
		}
	}
	if len(f.commands) != 0 {
		t.Fatal("plan executed commands")
	}
	if _, err := New(Options{}, io.Discard); err == nil {
		t.Fatal("build accepted invalid catalog")
	}
}

func TestRecipeSourcesIncludeAuxiliaryInputs(t *testing.T) {
	e := testEngine(t)
	for _, tc := range []struct {
		pkg, recipe string
		want        []string
	}{
		{"omnigres", "TARBALL := omnigres-20260212.tar.gz\nWHEEL := wheel-0.45.1-py3-none-any.whl\nSOURCE_WHEEL := ../SOURCES/$(WHEEL)\n", []string{"omnigres-20260212.tar.gz", "wheel-0.45.1-py3-none-any.whl"}},
		{"rdkit", "# old rdkit_202503.6.orig.tar.xz\nTARBALL := rdkit-Release_2026_09_2.tar.gz\nSOURCE_ENUM := ../SOURCES/better-enums-0.11.3-enum.h\nDSC_FILE := rdkit_202303.3-3.pgdg120+1.dsc\n", []string{"rdkit-Release_2026_09_2.tar.gz", "rdkit_202303.3-3.pgdg120+1.dsc", "better-enums-0.11.3-enum.h"}},
		{"cloudberry-pxf", "TARBALL := apache-cloudberry-pxf-2.2.0-incubating-src.tar.gz\n\tcp ../SOURCES/gradle-wrapper-8.14.4.jar ./build/gradle-wrapper.jar\n\tcp ../SOURCES/gradle-wrapper-8.14.4.jar.sha256 ./build/gradle-8.14.4-wrapper.jar.sha256\n", []string{"apache-cloudberry-pxf-2.2.0-incubating-src.tar.gz", "gradle-wrapper-8.14.4.jar", "gradle-wrapper-8.14.4.jar.sha256"}},
		{"valkey", "VERSION := 9.1.3\nSOURCE := ../SOURCES/valkey-$(VERSION).tar.gz\nSOURCE_URL := https://github.com/valkey-io/valkey/archive/refs/tags/$(VERSION).tar.gz\n", []string{"valkey-9.1.3.tar.gz"}},
	} {
		t.Run(tc.pkg, func(t *testing.T) {
			got := e.sourceNames(&Job{Package: tc.pkg}, []byte(tc.recipe))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestLiteralAPTExtractionPreservesCompoundShell(t *testing.T) {
	for _, line := range []string{
		"\tapt-get update && ./prepare",
		"\tapt-get update; \\",
		"\tapt-get install -y 'compiler (>= 1)'",
		"\tapt-get install -y $(BUILD_DEPS)",
		"\tapt-get install -y gcc > install.log",
		"\tdnf install -y gcc",
	} {
		if got := withoutRecipeInstalls(line); got != line {
			t.Fatalf("changed compound or non-APT command: %q -> %q", line, got)
		}
		if got := recipeInstallCommands(line); len(got) != 0 {
			t.Fatalf("extracted nonliteral command: %v", got)
		}
	}
	if got := recipeInstallCommands("\tapt install -y make # literal dependency\n"); !reflect.DeepEqual(got, [][]string{{"apt-get", "install", "-y", "make"}}) {
		t.Fatal(got)
	}
}

func TestRecipeTreeDriftIncludesPackagingAndPatches(t *testing.T) {
	for _, input := range []string{"debian/control.in", "patches/fix.patch", "helper.sh"} {
		t.Run(input, func(t *testing.T) {
			e := testEngine(t)
			root := filepath.Join(e.Options.SpecDir, "pgvector")
			put(t, filepath.Join(root, "Makefile"), "PG_VERSIONS := 18\n")
			put(t, filepath.Join(root, input), "original")
			plan, err := e.Prepare("make", []string{"vector"})
			if err != nil {
				t.Fatal(err)
			}
			put(t, filepath.Join(root, input), "modified")
			r, err := e.Run(context.Background(), plan)
			if err == nil || !strings.Contains(err.Error(), "recipe changed") || r.Success {
				t.Fatal(r, err)
			}
			if !fileExists(r.Manifest) || len(e.Exec.(*fakeExecutor).commands) != 0 {
				t.Fatal("drift did not stop before native execution and save result")
			}
		})
	}
}

func TestRecipeCopyFingerprintAndPermissions(t *testing.T) {
	src, dst := t.TempDir(), filepath.Join(t.TempDir(), "staged")
	helper := filepath.Join(src, "build")
	put(t, helper, "#!/bin/sh\n")
	if err := os.Chmod(helper, 0775); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(src, "x", "build", "old"), "old workspace")
	put(t, filepath.Join(src, "SOURCES", "old.tar.gz"), "old cache")
	if err := copyRecipe(src, dst); err != nil {
		t.Fatal(err)
	}
	want, err := recipeTreeHash(src)
	if err != nil {
		t.Fatal(err)
	}
	got, err := recipeTreeHash(dst)
	if err != nil || got != want {
		t.Fatalf("staged inputs differ: %s %s %v", want, got, err)
	}
}

func TestFailedBuildRetainsStagedSourceProvenance(t *testing.T) {
	for _, family := range []string{"rpm", "deb"} {
		t.Run(family, func(t *testing.T) {
			e := testEngine(t)
			e.Platform.Family = family
			put(t, filepath.Join(e.Options.Root, "src", "source.tar.gz"), "original archive")
			unit, out := t.TempDir(), t.TempDir()
			j := &Job{Package: "example", Sources: []string{"source.tar.gz"}}
			if family == "rpm" {
				j.Recipe = filepath.Join(e.Options.SpecDir, "SPECS", "example.spec")
			} else {
				j.Recipe = filepath.Join(e.Options.SpecDir, "example", "Makefile")
			}
			put(t, j.Recipe, "recipe")
			result := &Result{}
			e.Exec = &fakeExecutor{capture: func(Command) (string, error) {
				return "Source0: https://example.org/release.tar.gz#/source.tar.gz\n", nil
			}, run: func(Command) error {
				if len(result.Sources) != 1 || result.Sources[0].SHA256 != Hash([]byte("original archive")) {
					t.Fatal("source provenance missing before native execution")
				}
				put(t, result.Sources[0].Path, "recipe modified source")
				return errors.New("compiler failed")
			}}
			var err error
			if family == "rpm" {
				_, err = e.makeRPM(context.Background(), j, 0, unit, out, result, io.Discard)
			} else {
				_, err = e.makeDEB(context.Background(), j, 0, unit, out, result, io.Discard)
			}
			if err == nil || !strings.Contains(err.Error(), "compiler failed") || result.Sources[0].SHA256 != Hash([]byte("original archive")) {
				t.Fatal(result, err)
			}
		})
	}
}

func TestLegacyRPMCustomURLsAreNotReplaceable(t *testing.T) {
	for _, line := range []string{
		"baseurl=https://private.example/postgresql",
		"baseurl=https://repo.pigsty.cc.evil.example/yum/pgdg/18/redhat/rhel-9.8-aarch64",
		"mirrorlist=https://private.example/list",
		"baseurl=https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-9.8-aarch64\nexclude=postgresql18*",
	} {
		if recognizedLegacyRepos("[pgdg18]\n"+line+"\n", "rpm") {
			t.Fatalf("accepted custom definition: %s", line)
		}
	}
}

func TestRPMSourceURLRename(t *testing.T) {
	for _, raw := range []string{"source.tar.gz", "https://example.org/source.tar.gz", "https://example.org/release.tar.gz#/source.tar.gz"} {
		if got, err := rpmSourceName(raw); err != nil || got != "source.tar.gz" {
			t.Fatal(got, err)
		}
	}
	e := testEngine(t)
	e.Platform.Family = "rpm"
	put(t, filepath.Join(e.Options.Root, "src", "source.tar.gz"), "archive")
	e.Exec = &fakeExecutor{capture: func(Command) (string, error) {
		return "Source0: https://example.org/release.tar.gz#/source.tar.gz\n", nil
	}}
	j := &Job{Package: "example", Recipe: "example.spec", PG: []int{18}}
	if err := e.get(context.Background(), j, &Result{}, io.Discard); err != nil || !slices.Equal(j.Sources, []string{"source.tar.gz"}) {
		t.Fatal(j, err)
	}
}
