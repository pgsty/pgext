package build

import (
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestPG19BuildTargetsAndPG18ToolDefault(t *testing.T) {
	for _, family := range []string{"rpm", "deb"} {
		t.Run(family, func(t *testing.T) {
			e := testEngine(t)
			e.Platform.Family = family
			e.Options.PG = ""
			cmd, err := e.toolCommand("full")
			if err != nil {
				t.Fatal(err)
			}
			defaultArgs := strings.Join(cmd.Args, " ")
			if !strings.Contains(defaultArgs, "18") || strings.Contains(defaultArgs, "19") {
				t.Fatalf("unexpected tool default: %s", defaultArgs)
			}
			e.Options.PG = "19"
			cmd, err = e.toolCommand("full")
			if err != nil || !strings.Contains(strings.Join(cmd.Args, " "), "19") {
				t.Fatalf("explicit PG19 tools: %+v, %v", cmd, err)
			}
		})
	}
	got, err := ParsePG("19,18,15,19")
	if err != nil || !reflect.DeepEqual(got, []int{19, 18, 15}) {
		t.Fatalf("ParsePG: %v, %v", got, err)
	}
	for _, family := range []string{"rpm", "deb"} {
		root := t.TempDir()
		recipe := []byte("build19 build18:\n\t./build\n")
		if family == "rpm" {
			put(t, filepath.Join(root, "Makefile"), "pgvector:\n\t./build pgvector 19 18\n")
		}
		got := recipePG(family, root, "pgvector", recipe)
		if len(got) != 2 || !slices.Contains(got, 19) {
			t.Fatalf("%s recipe lost PG19: %v", family, got)
		}
	}
}

func TestRetiredTargetsAreRejectedAndEL8ExcludesPG19(t *testing.T) {
	for _, pg := range []string{"14", "14,19", "19,14"} {
		if _, err := ParsePG(pg); err == nil {
			t.Fatalf("accepted retired target %q", pg)
		}
	}
	for _, arch := range []string{"x86_64", "aarch64", "amd64", "arm64"} {
		if _, err := PlanPlatform("el8." + arch); err != nil {
			t.Fatalf("expected active el8.%s builder: %v", arch, err)
		}
	}
	if _, err := ParsePlatform("ID=rocky\nVERSION_ID=8.10\n", "amd64"); err != nil {
		t.Fatalf("native EL8 detection must remain supported for PG15–18: %v", err)
	}
	if _, err := ParsePlatform("ID=centos\nVERSION_ID=7.9\n", "amd64"); err == nil {
		t.Fatal("native EL7 detection must reject retired builder")
	}

	el8 := testEngine(t)
	el8.Platform = Platform{OS: "el8", Family: "rpm", Arch: "amd64", Version: "8.10"}
	el8.Options.Plan = true
	repoText, err := el8.repoContent("pgdg")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(repoText, "pgdg19") || !strings.Contains(repoText, "pgdg18") {
		t.Fatalf("EL8 PGDG repo must include PG18 and omit PG19:\n%s", repoText)
	}
	el9 := testEngine(t)
	el9.Platform = Platform{OS: "el9", Family: "rpm", Arch: "amd64", Version: "9.6"}
	el9.Options.Plan = true
	el9RepoText, err := el9.repoContent("pgdg")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(el9RepoText, "pgdg19") {
		t.Fatalf("EL9 PGDG repo must include PG19:\n%s", el9RepoText)
	}
	el8.Options.PG = "19"
	if _, err := el8.toolCommand("full"); err == nil {
		t.Fatal("EL8 toolCommand must reject PG19")
	}
}

func TestRetiredRecipeDoesNotFallBackToCatalog(t *testing.T) {
	for _, family := range []string{"rpm", "deb"} {
		for _, requested := range []string{"", "19"} {
			e := testEngine(t)
			e.Platform.Family, e.Options.PG, e.Options.Plan = family, requested, true
			j := &Job{Package: "pgvector", Extension: e.Catalog.Resolve("vector"), Recipe: filepath.Join(e.Options.SpecDir, "pgvector", "Makefile")}
			j.Extension.PG = []int{19, 18, 15, 14}
			recipe := []byte("PG_VERSIONS := 14\nbuild14:\n\t./build\n")
			if family == "rpm" {
				put(t, filepath.Join(e.Options.SpecDir, "Makefile"), "pgvector:\n\t./build pgvector 14\n")
			}
			if pgs, err := e.jobPG(j, recipe); err == nil {
				t.Fatalf("%s --pg=%q scheduled retired recipe as %v", family, requested, pgs)
			}
		}
	}
}

func TestCatalogEmptyPGRange(t *testing.T) {
	for _, scenario := range []struct {
		name, catalogPG, requested, recipe string
		want                               []int
	}{
		{"unknown uses platform", "", "", "", []int{19, 18, 17, 16, 15}},
		{"unknown allows explicit major", "", "18", "", []int{18}},
		{"empty has no default", "{}", "", "", nil},
		{"empty rejects explicit major", "{}", "18", "", nil},
		{"declared recipe takes precedence", "{}", "", "PG_VERSIONS := 17\n", []int{17}},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			engine := testEngine(t)
			catalog, err := ParseCatalog([]byte("name,pkg,source,pg_ver,deb_pkg\nexample,example,example.tar.gz," + scenario.catalogPG + ",postgresql-$v-example\n"))
			if err != nil {
				t.Fatal(err)
			}
			engine.Catalog = catalog
			engine.Options.PG = scenario.requested
			engine.Options.Plan = true
			put(t, filepath.Join(engine.Options.SpecDir, "example", "Makefile"), scenario.recipe)
			plan, err := engine.Prepare("make", []string{"example"})
			if scenario.want == nil {
				if err == nil {
					t.Fatalf("scheduled empty catalog range: %v", plan.Jobs[0].PG)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := plan.Jobs[0].PG; !slices.Equal(got, scenario.want) {
				t.Fatalf("planned %v, want %v", got, scenario.want)
			}
		})
	}
}
