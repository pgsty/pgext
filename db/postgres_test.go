package db

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestPostgreSQLTargetsMatchEmbeddedCatalog(t *testing.T) {
	data, err := ReadFile("pg.csv")
	if err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	var active []int
	for _, row := range rows[1:] {
		if row[1] != "t" {
			continue
		}
		pg, err := strconv.Atoi(row[0])
		if err != nil {
			t.Fatal(err)
		}
		active = append(active, pg)
	}
	if !reflect.DeepEqual(active, SupportedPGVersions()) {
		t.Fatalf("embedded PostgreSQL targets %v differ from build targets %v", active, SupportedPGVersions())
	}
	if DefaultPGVersion != 18 || SupportsPGVersion(14) || !SupportsPGVersion(19) || SupportsPGVersion(20) {
		t.Fatal("build targets must be PG15–19 with the PG18 default")
	}
}

func TestBuildPlatformsMatchEmbeddedCatalog(t *testing.T) {
	data, err := ReadFile("os.csv")
	if err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	var active []string
	for _, row := range rows[1:] {
		if row[1] == "t" {
			active = append(active, row[0])
		}
	}
	want := SupportedOSPlatforms()
	slices.Sort(active)
	slices.Sort(want)
	if !slices.Equal(active, want) || len(want) != 16 {
		t.Fatalf("catalog platforms %v differ from build platforms %v", active, want)
	}
	for _, platform := range []string{"el8.x86_64", "el8.aarch64"} {
		if !SupportsOSPlatform(platform) {
			t.Fatalf("expected active platform %s", platform)
		}
	}
	if !SupportsOSPG("el8", 18) || SupportsOSPG("el8", 19) || !SupportsOSPG("el9", 19) {
		t.Fatal("EL8 must support PG15–18 and mark PG19 unsupported")
	}
	if !reflect.DeepEqual(SupportedPGVersionsForOS("el8"), []int{18, 17, 16, 15}) {
		t.Fatalf("unexpected EL8 PG versions: %v", SupportedPGVersionsForOS("el8"))
	}
	for _, platform := range []string{"el7.x86_64", "d11.x86_64", "u20.x86_64", "el9.ppc64le"} {
		if SupportsOSPlatform(platform) {
			t.Fatalf("unexpected build platform %s", platform)
		}
	}
}
