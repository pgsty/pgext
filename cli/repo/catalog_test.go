package repo

import (
	"slices"
	"strings"
	"testing"
)

func testPlatform(t *testing.T, code, version string) Platform {
	t.Helper()
	p, err := ParsePlatform(code, version)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func testCatalog(t *testing.T) []Repository {
	t.Helper()
	catalog, err := DefaultCatalog("")
	if err != nil {
		t.Fatal(err)
	}
	return catalog
}

func TestCatalogSelectionAndPlatformFiltering(t *testing.T) {
	catalog := testCatalog(t)
	for _, tt := range []struct {
		platform string
		selector []string
		want     []string
		absent   []string
	}{
		{"u24.arm64", []string{"node"}, []string{"base", "updates", "backports", "security"}, nil},
		{"d12.amd64", []string{"node"}, []string{"base", "updates", "security"}, []string{"backports"}},
		{"el8.x86_64", []string{"node"}, []string{"baseos", "appstream", "extras", "powertools", "epel"}, []string{"crb"}},
		{"el9.aarch64", []string{"pgdg,pigsty", "pgdg18"}, []string{"pigsty-infra", "pigsty-pgsql", "pgdg-common", "pgdg14", "pgdg15", "pgdg16", "pgdg17", "pgdg18"}, []string{"pgdg-beta", "pgdg18-nonfree"}},
		{"u24.amd64", []string{"all"}, []string{"pigsty-pgsql", "pigsty-infra", "base", "updates", "backports", "security", "pgdg"}, []string{"pigsty-local", "docker-ce", "pgdg-beta"}},
	} {
		t.Run(tt.platform+strings.Join(tt.selector, "-"), func(t *testing.T) {
			selected, err := Select(catalog, tt.selector, testPlatform(t, tt.platform, ""))
			if err != nil {
				t.Fatal(err)
			}
			var names []string
			for _, r := range selected {
				names = append(names, r.Name)
			}
			if !slices.Equal(names, tt.want) {
				t.Fatalf("names = %v, want %v", names, tt.want)
			}
			for _, absent := range tt.absent {
				if slices.Contains(names, absent) {
					t.Fatal("unexpected", absent)
				}
			}
		})
	}
	for _, args := range [][]string{{"pgdg", "typo"}, {"pgdg,"}, {"pgdg18-nonfree"}, {"crb"}} {
		if _, err := Select(catalog, args, testPlatform(t, "el8.arm64", "")); err == nil {
			t.Fatalf("accepted invalid/unavailable selection %v", args)
		}
	}
}

func TestRenderMatchesPigstyPlatformRules(t *testing.T) {
	for _, tt := range []struct {
		platform, version, name, region, want string
	}{
		{"el9.arm64", "9.6", "pgdg18", "default", "rhel-9.6-aarch64"},
		{"el9.amd64", "9.10.2", "pgdg18", "default", "rhel-9.10-x86_64"},
		{"el9.amd64", "9.5", "pgdg18", "default", "rhel-9-x86_64"},
		{"el10.arm64", "10.1", "pgdg18", "default", "rhel-10.1-aarch64"},
		{"el10.amd64", "10.1", "epel", "default", "/10z/Everything/x86_64/"},
		{"el9.arm64", "9.6", "pigsty-pgsql", "china", "https://repo.pigsty.cc/yum/pgsql/el9.aarch64"},
		{"el9.arm64", "9.6", "pgdg18", "china", "https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-9.6-aarch64"},
		{"u24.arm64", "", "base", "default", "http://ports.ubuntu.com/ubuntu-ports/ noble"},
		{"u24.amd64", "", "base", "default", "https://mirrors.edge.kernel.org/ubuntu/ noble"},
		{"d13.arm64", "", "docker", "default", "https://download.docker.com/linux/debian trixie stable"},
		{"u26.amd64", "", "pgdg", "china", "https://mirrors.cloud.tencent.com/postgresql/repos/apt/ resolute-pgdg main"},
		{"d12.amd64", "", "pigsty", "missing-region", "https://repo.pigsty.io/apt/pgsql/bookworm"},
	} {
		t.Run(tt.platform+tt.version+tt.name+tt.region, func(t *testing.T) {
			p := testPlatform(t, tt.platform, tt.version)
			rs, err := Select(testCatalog(t), []string{tt.name}, p)
			if err != nil {
				t.Fatal(err)
			}
			content, err := Render(rs[0], p, RenderOptions{Region: tt.region})
			if err != nil || !strings.Contains(content, tt.want) {
				t.Fatalf("render = %q, %v; want %q", content, err, tt.want)
			}
		})
	}
}

func TestRenderEverySupportedCoordinate(t *testing.T) {
	catalog := testCatalog(t)
	for code := range codenames {
		for _, arch := range []string{"x86_64", "aarch64"} {
			t.Run(code+"."+arch, func(t *testing.T) {
				version := ""
				if code == "el9" {
					version = "9.6"
				} else if code == "el10" {
					version = "10.1"
				}
				p := testPlatform(t, code+"."+arch, version)
				selected, err := Select(catalog, nil, p)
				if err != nil || len(selected) == 0 {
					t.Fatalf("select: %v", err)
				}
				for _, r := range selected {
					for _, region := range []string{"default", "china", "europe"} {
						text, err := Render(r, p, RenderOptions{Region: region, AdminIP: "10.10.10.10:8080"})
						if err != nil || strings.Contains(text, "$") {
							t.Fatalf("%s %s: %s %v", r.Name, region, text, err)
						}
					}
				}
			})
		}
	}
}

func TestCatalogOverridesAndTrustPolicy(t *testing.T) {
	data := `repo_upstream_default:
  - {name: ignored, module: custom, releases: [24], arch: [amd64], baseurl: {default: 'https://ignored.invalid noble main'}}
repo_upstream:
  - name: custom
    module: custom
    releases: [24]
    arch: [arm64]
    baseurl: {default: 'https://example.org/apt ${distro_codename} main'}
    meta: {signed-by: /etc/apt/keyrings/custom.asc}
unrelated_pigsty_parameter: true
`
	entries, err := ParseCatalog([]byte(data), "deb")
	if err != nil || len(entries) != 1 || entries[0].Name != "custom" {
		t.Fatalf("override: %v %v", entries, err)
	}
	content, err := Render(entries[0], testPlatform(t, "u24.arm64", ""), RenderOptions{})
	if err != nil || strings.Contains(content, "trusted=yes") || !strings.Contains(content, "arch=arm64 signed-by=/etc/apt/keyrings/custom.asc") {
		t.Fatalf("signature policy: %q %v", content, err)
	}
	entries, err = ParseCatalog([]byte(data+"\n"), "rpm")
	entries[0].Releases = []int{9}
	entries[0].BaseURL["default"] = "https://example.org/rpm/$releasever/$basearch"
	entries[0].Meta = map[string]string{"gpgcheck": "1", "gpgkey": "https://example.org/key.asc", "enabled": "0"}
	content, err = Render(entries[0], testPlatform(t, "el9.arm64", "9.6"), RenderOptions{})
	if err != nil || !strings.Contains(content, "gpgcheck=1") || !strings.Contains(content, "enabled=0") {
		t.Fatalf("RPM policy: %q %v", content, err)
	}
	if entries, err := ParseCatalog([]byte("repo_upstream: []\n"), "deb"); err != nil || len(entries) != 0 {
		t.Fatalf("explicit empty catalog: %v %v", entries, err)
	}
	for _, data := range []string{
		"unrelated: true", "repo_upstream: [",
		`repo_upstream: [{name: ../bad, module: node, releases: [24], arch: [arm64], baseurl: {default: 'https://example.org noble main'}}]`,
		`repo_upstream: [{name: bad, module: node, releases: [24], arch: [arm64], baseurl: {default: "https://example.org\ninjected"}}]`,
	} {
		if _, err := ParseCatalog([]byte(data), "deb"); err == nil {
			t.Fatalf("accepted invalid catalog %q", data)
		}
	}
}

func TestPlatformAndLocalValidation(t *testing.T) {
	p, err := ParseOSRelease("ID=almalinux\nVERSION_ID=\"9.6\"\n", "amd64")
	if err != nil || p.OS != "el9" || p.Version != "9.6" || p.Arch != "x86_64" {
		t.Fatal(p, err)
	}
	for _, tt := range [][2]string{{"el9.arm64", "10.1"}, {"u20.amd64", ""}, {"u24.bad", ""}, {"u24.amd64", "24;evil"}} {
		if _, err := ParsePlatform(tt[0], tt[1]); err == nil {
			t.Fatal("accepted", tt)
		}
	}
	p = testPlatform(t, "u24.arm64", "")
	rs, _ := Select(testCatalog(t), []string{"local"}, p)
	for _, admin := range []string{"", "host/evil", "host\ninjected", "user:pass@host"} {
		if _, err := Render(rs[0], p, RenderOptions{AdminIP: admin}); err == nil {
			t.Fatal("accepted admin host", admin)
		}
	}
	p = testPlatform(t, "el10.arm64", "")
	rs, _ = Select(testCatalog(t), []string{"pgdg"}, p)
	if _, err := Render(rs[0], p, RenderOptions{}); err == nil {
		t.Fatal("accepted unknown PGDG EL10 minor version")
	}
}
