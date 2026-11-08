package build

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestCatalogRecipeAliases(t *testing.T) {
	for _, tc := range []struct{ family, name, pkg, recipe string }{
		{"deb", "h3", "pg_h3", "h3"},
		{"rpm", "h3", "pg_h3", "pg_h3"},
		{"deb", "onesparse", "one_sparse", "onesparse"},
		{"rpm", "onesparse", "one_sparse", "onesparse"},
		{"deb", "address_standardizer", "address_standardizer", "address-standardizer"},
		{"deb", "postgis_tiger_geocoder", "postgis_tiger_geocoder", "postgis-tiger-geocoder"},
	} {
		t.Run(tc.family+"/"+tc.name, func(t *testing.T) {
			e := testEngine(t)
			e.Platform.Family = tc.family
			var err error
			e.Catalog, err = ParseCatalog([]byte("name,pkg,source\n" + tc.name + "," + tc.pkg + ",source.tar.gz\n"))
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(e.Options.SpecDir, tc.recipe, "Makefile")
			content := "PG_VERSIONS := 18\n"
			if tc.family == "rpm" {
				path = filepath.Join(e.Options.SpecDir, "SPECS", tc.recipe+".spec")
				content = "%{!?pgmajorversion:%global pgmajorversion 18}\n"
			}
			put(t, path, content)
			p, err := e.Prepare("make", []string{tc.name, tc.pkg, tc.recipe})
			if err != nil {
				t.Fatal(err)
			}
			if len(p.Jobs) != 1 || p.Jobs[0].Recipe != path || p.Jobs[0].Extension == nil {
				t.Fatalf("recipe aliases lost catalog identity or deduplication: %+v", p.Jobs)
			}
		})
	}
}

func TestDebianStagedInputs(t *testing.T) {
	if _, err := exec.LookPath("make"); err != nil {
		t.Skip("make unavailable")
	}
	e := testEngine(t)
	recipe := filepath.Join(e.Options.SpecDir, "documentdb", "Makefile")
	stale := t.TempDir()
	for _, name := range []string{"old.tar.gz", "new.tar.gz", "intel.tar.gz"} {
		put(t, filepath.Join(stale, name), "stale cache")
	}
	t.Setenv("SOURCE_TARBALL", filepath.Join(stale, "new.tar.gz"))
	put(t, recipe, `PG_VERSIONS := 17 18
ifeq ($(PG_MAJOR),17)
TARBALL := old.tar.gz
else
TARBALL := new.tar.gz
endif
INTEL_TARBALL := intel.tar.gz
SOURCE_TARBALL ?= `+stale+`/$(TARBALL)
INTEL_SOURCE_TARBALL ?= `+stale+`/$(INTEL_TARBALL)
all:
	@cat "$(SOURCE_TARBALL)" "$(INTEL_SOURCE_TARBALL)"
	@cat "$$LIBBSON_SOURCE" "$$PCRE2_SOURCE" "$$UNCRUSTIFY_SOURCE" "$$CITUS_TOOLS_SOURCE"
`)
	p, err := e.Prepare("get", []string{"documentdb"})
	if err != nil {
		t.Fatal(err)
	}
	j := p.Jobs[0]
	wantSources := []string{"old.tar.gz", "new.tar.gz", "intel.tar.gz", "mongo-c-driver-1.28.0.tar.gz", "pcre2-10.40.tar.gz", "uncrustify-uncrustify-0.68.1.tar.gz", "citus-tools-e36e4ea4258989bf527744334f6c633bb67e0686.tar.gz"}
	if !slices.Equal(j.Sources, wantSources) {
		t.Fatalf("incomplete DocumentDB source plan: %v", j.Sources)
	}
	sums := map[string]string{}
	for _, name := range j.Sources {
		put(t, filepath.Join(e.Options.Root, "src", name), name+"\n")
		sums[name] = Hash([]byte(name + "\n"))
	}
	b, _ := json.Marshal(sums)
	e.Options.Checksums = filepath.Join(t.TempDir(), "checksums.json")
	put(t, e.Options.Checksums, string(b))
	e.Exec = NativeExecutor{}
	for _, pg := range []int{17, 18} {
		var output bytes.Buffer
		r := &Result{}
		_, err := e.makeDEB(context.Background(), j, pg, t.TempDir(), t.TempDir(), r, &output)
		if err != nil {
			t.Fatal(err, output.String())
		}
		want := append([]string{map[int]string{17: "old.tar.gz", 18: "new.tar.gz"}[pg], "intel.tar.gz"}, wantSources[3:]...)
		if output.String() != strings.Join(want, "\n")+"\n" || len(r.Sources) != len(j.Sources) {
			t.Fatalf("PG%d did not consume its checksummed inputs: %s", pg, output.String())
		}
	}
}

func TestDebianMissingOrUnverifiedSourceStopsBeforeMake(t *testing.T) {
	for _, tc := range []struct{ name, content, sums, problem string }{
		{"missing", "", "{}", "stage source"},
		{"unlisted", "archive", "{}", "checksum manifest lacks"},
		{"mismatch", "archive", `{"source.tar.gz":"` + Hash([]byte("other")) + `"}`, "checksum mismatch"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := testEngine(t)
			j := &Job{Package: "example", Recipe: filepath.Join(e.Options.SpecDir, "example", "Makefile"), Sources: []string{"source.tar.gz"}}
			put(t, j.Recipe, "all:\n\t@true\n")
			if tc.content != "" {
				put(t, filepath.Join(e.Options.Root, "src", "source.tar.gz"), tc.content)
			}
			e.Options.Checksums = filepath.Join(t.TempDir(), "checksums.json")
			put(t, e.Options.Checksums, tc.sums)
			_, err := e.makeDEB(context.Background(), j, 0, t.TempDir(), t.TempDir(), &Result{}, io.Discard)
			if err == nil || !strings.Contains(err.Error(), tc.problem) || len(e.Exec.(*fakeExecutor).commands) != 0 {
				t.Fatalf("unsafe input reached make: %v", err)
			}
		})
	}
}

func TestPrivateCMakeDependencyAndStaging(t *testing.T) {
	for _, tc := range []struct{ os, pkg, constraint, failure string }{
		{"u22", "graphblas", "3.23", ""},
		{"u22", "lagraph", "3.23", ""},
		{"u22", "pg_mooncake", "3.24", ""},
		{"u24", "graphblas", "3.23", ""},
		{"u22", "other", "3.24", ""},
		{"u22", "graphblas", "3.23", "satisfy"},
		{"u22", "graphblas", "3.23", "helper"},
	} {
		t.Run(tc.os+"/"+tc.pkg+"/"+tc.failure, func(t *testing.T) {
			e := testEngine(t)
			e.Platform.OS = tc.os
			j := &Job{Package: tc.pkg, Recipe: filepath.Join(e.Options.SpecDir, tc.pkg, "Makefile")}
			put(t, j.Recipe, "all:\n\t@../kitware-cmake\n")
			control := "Source: example\nBuild-Depends: cmake (>= " + tc.constraint + "), pkg-config, libgraphblas-dev (>= 10.5.2)\n"
			put(t, filepath.Join(filepath.Dir(j.Recipe), "debian", "control"), control)
			put(t, filepath.Join(e.Options.SpecDir, "kitware-cmake"), "#!/bin/sh\nprintf 'private cmake\\n'\n")
			f := &fakeExecutor{run: func(c Command) error {
				if tc.failure == "satisfy" && c.Args[1] == "satisfy" || tc.failure == "helper" && c.Args[0] == "sh" {
					return errors.New("injected failure")
				}
				return nil
			}}
			e.Exec = f
			err := e.dep(context.Background(), j, io.Discard)
			if (err != nil) != (tc.failure != "") {
				t.Fatal(err)
			}
			wantDeps := "cmake (>= " + tc.constraint + "), pkg-config, libgraphblas-dev (>= 10.5.2)"
			private := tc.os == "u22" && tc.pkg != "other"
			if private {
				wantDeps = "cmake, pkg-config, libgraphblas-dev (>= 10.5.2)"
			}
			if f.commands[1].Args[4] != wantDeps {
				t.Fatal("dependency constraints changed", f.commands[1])
			}
			wantCommands := 2
			if private && tc.failure != "satisfy" {
				wantCommands++
			}
			if len(f.commands) != wantCommands {
				t.Fatal(f.commands)
			}
			if tc.failure != "" || tc.pkg == "other" {
				return
			}
			// Run just the helper from the isolated recipe, never a package build.
			e.Exec = NativeExecutor{}
			var output bytes.Buffer
			_, err = e.makeDEB(context.Background(), j, 0, t.TempDir(), t.TempDir(), &Result{}, &output)
			if err != nil || output.String() != "private cmake\n" {
				t.Fatal(err, output.String())
			}
		})
	}
}
