package cli

import (
	"bytes"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"text/template"
)

func TestPackagePatternUsesExtensionTemplates(t *testing.T) {
	g := &PigstyConfigGenerator{}
	ext := &ExtensionData{
		Name:   "babelfishpg_common",
		RPMPkg: "babelfish_$v",
		DEBPkg: "babelfishpg-$v-babelfish",
	}

	if got := g.getRPMPackagePattern(ext); got != "babelfish_$v" {
		t.Fatalf("getRPMPackagePattern() = %q, want %q", got, "babelfish_$v")
	}
	if got := g.getDEBPackagePattern(ext); got != "babelfishpg-$v-babelfish" {
		t.Fatalf("getDEBPackagePattern() = %q, want %q", got, "babelfishpg-$v-babelfish")
	}
}

func TestGetPackageNameForCategoryUsesTemplate(t *testing.T) {
	rpmGen := &PigstyConfigGenerator{osCode: "el9"}
	debGen := &PigstyConfigGenerator{osCode: "u24"}
	ext := &ExtensionData{
		Name:   "babelfishpg_common",
		RPMPkg: "babelfish_$v",
		DEBPkg: "babelfishpg-$v-babelfish",
	}

	if got := rpmGen.getPackageNameForCategory(ext, 17); got != "babelfish_17" {
		t.Fatalf("RPM package = %q, want %q", got, "babelfish_17")
	}
	if got := debGen.getPackageNameForCategory(ext, 17); got != "babelfishpg-17-babelfish" {
		t.Fatalf("DEB package = %q, want %q", got, "babelfishpg-17-babelfish")
	}
}

func TestGenerateExtensionMappingsPreferAvailable(t *testing.T) {
	g := &PigstyConfigGenerator{osCode: "el9"}
	extensions := map[string]*ExtensionData{
		"babelfishpg_common": {
			ID:         1,
			Name:       "babelfishpg_common",
			Alias:      "babelfishpg_common",
			Category:   "SIM",
			CategoryID: 1,
			RPMPkg:     "babelfish_$v",
			RPMPG:      []int{17},
			Available: map[int]bool{
				18: true,
				17: true,
			},
		},
	}

	mappings := g.generateExtensionMappings(extensions)
	if len(mappings) != 1 {
		t.Fatalf("expected 1 mapping, got %d", len(mappings))
	}
	if !reflect.DeepEqual(mappings[0].PGVers, []int{18, 17}) {
		t.Fatalf("mapping versions = %v, want %v", mappings[0].PGVers, []int{18, 17})
	}
}

func TestGenerateExtensionMappingsFallbackToDeclaredPG(t *testing.T) {
	tests := []struct {
		name     string
		gen      *PigstyConfigGenerator
		ext      *ExtensionData
		expected []int
	}{
		{
			name: "rpm",
			gen:  &PigstyConfigGenerator{osCode: "el9"},
			ext: &ExtensionData{
				ID:         1,
				Name:       "babelfishpg_common",
				Alias:      "babelfishpg_common",
				Category:   "SIM",
				CategoryID: 1,
				RPMPkg:     "babelfish_$v",
				RPMPG:      []int{17},
				Available:  map[int]bool{},
			},
			expected: []int{17},
		},
		{
			name: "deb",
			gen:  &PigstyConfigGenerator{osCode: "u24"},
			ext: &ExtensionData{
				ID:         1,
				Name:       "babelfishpg_common",
				Alias:      "babelfishpg_common",
				Category:   "SIM",
				CategoryID: 1,
				DEBPkg:     "babelfishpg-$v-babelfish",
				DEBPG:      []int{17},
				Available:  map[int]bool{},
			},
			expected: []int{17},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappings := tt.gen.generateExtensionMappings(map[string]*ExtensionData{
				tt.ext.Name: tt.ext,
			})
			if len(mappings) != 1 {
				t.Fatalf("expected 1 mapping, got %d", len(mappings))
			}
			if !reflect.DeepEqual(mappings[0].PGVers, tt.expected) {
				t.Fatalf("mapping versions = %v, want %v", mappings[0].PGVers, tt.expected)
			}
		})
	}
}

func TestEL9ARMPatroniVersionLockUsesCurrentVersion(t *testing.T) {
	g := &PigstyConfigGenerator{
		osCode:    "el9",
		arch:      "aarch64",
		constants: GetConfigConstants(),
	}
	funcs := g.getFuncMap()

	pgsqlUtility := funcs["getPgsqlUtility"].(func() string)()
	wantPackages := []string{
		"patroni-" + el9ARMPatroniVersion,
		"patroni-etcd-" + el9ARMPatroniVersion,
	}
	for _, pkg := range wantPackages {
		if !strings.Contains(pgsqlUtility, pkg) {
			t.Fatalf("pgsql utility package list missing %q: %q", pkg, pgsqlUtility)
		}
	}
	for _, patch := range []string{"0", "1"} {
		stalePackage := "patroni-4.1." + patch
		if strings.Contains(pgsqlUtility, stalePackage) {
			t.Fatalf("pgsql utility package list contains stale patroni version: %q", pgsqlUtility)
		}
	}

	getUtilPkg := funcs["getUtilPkg"].(func(string, string) string)
	if got := getUtilPkg("patroni", "patroni patroni-etcd"); got != "patroni-"+el9ARMPatroniVersion+" patroni-etcd-"+el9ARMPatroniVersion {
		t.Fatalf("patroni util package list = %q", got)
	}
	if got := getUtilPkg("pgsql-common", "patroni patroni-etcd pgbouncer"); !strings.Contains(got, "patroni-"+el9ARMPatroniVersion+" patroni-etcd-"+el9ARMPatroniVersion) {
		t.Fatalf("pgsql-common util package list = %q", got)
	}
}

func TestDebDNSPackageUsesBind9DNSUtils(t *testing.T) {
	for _, osCode := range []string{"d12", "d13", "u22", "u24", "u26"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osCode:    osCode,
				constants: GetConfigConstants(),
			}
			funcs := g.getFuncMap()

			if got := funcs["getDNSPackage"].(func() string)(); got != "bind9-dnsutils" {
				t.Fatalf("getDNSPackage() = %q, want %q", got, "bind9-dnsutils")
			}

			packages := strings.Fields(funcs["getNodePackage2"].(func() string)())
			hasBind9DNSUtils := false
			for _, pkg := range packages {
				if pkg == "bind9-dnsutils" {
					hasBind9DNSUtils = true
				}
				if pkg == "dnsutils" {
					t.Fatalf("node-package2 contains legacy dnsutils package: %q", strings.Join(packages, " "))
				}
			}
			if !hasBind9DNSUtils {
				t.Fatalf("node-package2 missing bind9-dnsutils: %q", strings.Join(packages, " "))
			}
		})
	}
}

func TestRPMDNSPackageUsesBindUtils(t *testing.T) {
	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osCode:    osCode,
				constants: GetConfigConstants(),
			}
			funcs := g.getFuncMap()

			if got := funcs["getDNSPackage"].(func() string)(); got != "bind-utils" {
				t.Fatalf("getDNSPackage() = %q, want %q", got, "bind-utils")
			}

			packages := strings.Fields(funcs["getNodePackage2"].(func() string)())
			hasBindUtils := false
			for _, pkg := range packages {
				if pkg == "bind-utils" {
					hasBindUtils = true
				}
				if pkg == "bind9-utils" || pkg == "bind9.18-utils" {
					t.Fatalf("node-package2 contains nonstandard RPM DNS utilities package: %q", strings.Join(packages, " "))
				}
			}
			if !hasBindUtils {
				t.Fatalf("node-package2 missing bind-utils: %q", strings.Join(packages, " "))
			}
		})
	}
}

func TestDebTemplateRendersBind9DNSUtils(t *testing.T) {
	g := &PigstyConfigGenerator{
		osName:    "u24.x86_64",
		osCode:    "u24",
		arch:      "x86_64",
		constants: GetConfigConstants(),
	}
	tmpl := template.Must(template.New("deb").Funcs(g.getFuncMap()).Parse(debTemplate))
	data := map[string]interface{}{
		"OSName":       g.osName,
		"OSCode":       g.osCode,
		"Arch":         g.arch,
		"Constants":    g.constants,
		"Extensions":   map[string]*ExtensionData{},
		"PGVersions":   []int{18, 17, 16, 15, 14},
		"CategoryExts": map[string]CategoryPkgList{},
		"Categories":   []string{},
		"ExtMappings":  []ExtensionMapping{},
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatal(err)
	}
	rendered := buf.String()

	required := []string{
		"zlib1g,acl,bind9-dnsutils,libreadline-dev",
		"node-package2:           \"netcat-openbsd socat net-tools ipvsadm bind9-dnsutils",
	}
	for _, fragment := range required {
		if !strings.Contains(rendered, fragment) {
			t.Fatalf("rendered DEB config missing expected fragment: %q", fragment)
		}
	}
	for _, line := range strings.Split(rendered, "\n") {
		if strings.Contains(line, ",dnsutils,") || strings.Contains(line, " dnsutils ") {
			t.Fatalf("rendered DEB config contains legacy dnsutils token: %q", line)
		}
	}
}

func TestConfigConstantsIncludeBabelfishAlias(t *testing.T) {
	constants := GetConfigConstants()
	for _, mapping := range constants.PGSQLExoticMap {
		if mapping.Key != "babelfish" {
			continue
		}
		if mapping.RPM != "babelfishpg-17 babelfishpg-17-babelfish" {
			t.Fatalf("babelfish RPM alias = %q", mapping.RPM)
		}
		if mapping.DEB != "babelfishpg-17 babelfishpg-17-babelfish" {
			t.Fatalf("babelfish DEB alias = %q", mapping.DEB)
		}
		return
	}
	t.Fatal("babelfish mapping not found in PGSQLExoticMap")
}

func TestConfigConstantsIncludeAgensAndPgEdgeAlias(t *testing.T) {
	constants := GetConfigConstants()
	expected := map[string]PackageMapping{
		"agens": {
			Key: "agens",
			RPM: "agensgraph_$v",
			DEB: "agensgraph-$v",
		},
		"pgedge": {
			Key: "pgedge",
			RPM: "pgedge_$v spock_$v lolor_$v snowflake_$v",
			DEB: "pgedge-$v pgedge-$v-spock pgedge-$v-lolor pgedge-$v-snowflake",
		},
	}

	found := make(map[string]PackageMapping)
	for _, mapping := range constants.PGSQLExoticMap {
		if _, ok := expected[mapping.Key]; ok {
			found[mapping.Key] = mapping
		}
	}

	for key, want := range expected {
		got, ok := found[key]
		if !ok {
			t.Fatalf("%s mapping not found in PGSQLExoticMap", key)
		}
		if got.RPM != want.RPM {
			t.Fatalf("%s RPM alias = %q, want %q", key, got.RPM, want.RPM)
		}
		if got.DEB != want.DEB {
			t.Fatalf("%s DEB alias = %q, want %q", key, got.DEB, want.DEB)
		}
	}
}

func TestTemplatesIncludeAgensAndPgEdgeHomeMap(t *testing.T) {
	required := []string{
		"agens:  '/usr/agens-$v'",
		"pgedge: '/usr/pgedge-$v'",
	}
	for _, fragment := range required {
		if !strings.Contains(rpmTemplate, fragment) {
			t.Fatalf("rpmTemplate missing expected home map fragment: %q", fragment)
		}
		if !strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate missing expected home map fragment: %q", fragment)
		}
	}
}

func TestDebTemplateUbuntuRepoSuffixesMatchRepoName(t *testing.T) {
	required := []string{
		"name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         22,24,26] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-updates",
		"name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         22,24,26] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-backports",
		"name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         22,24,26] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-security",
		"name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         22,24,26] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-updates",
		"name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         22,24,26] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-backports",
		"name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         22,24,26] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-security",
	}
	for _, fragment := range required {
		if !strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate missing expected fragment: %q", fragment)
		}
	}
}

func TestDebTemplateRepoReleaseCompatibility(t *testing.T) {
	required := []string{
		"name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [11,12,13,22,24,26]",
		"name: pigsty-pgsql   ,description: 'Pigsty PgSQL'       ,module: pgsql   ,releases: [11,12,13,22,24,26]",
		"name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         22,24,26] ,arch: [x86_64",
		"name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         22,24,26] ,arch: [        aarch64]",
		"name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [11,12,13,22,24,26]",
		"name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta    ,releases: [11,12,13,22,24,26]",
		"name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [11,12,13,22,24   ]",
		"name: citus          ,description: 'Citus'              ,module: extra   ,releases: [11,12,   22      ]",
		"name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [   12,13,22,24   ]",
		"name: groonga        ,description: 'Groonga Ubuntu'     ,module: groonga ,releases: [         22,24   ]",
		"name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [11,12,   22,24   ] ,arch: [x86_64         ]",
		"name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [   12,   22,24   ]",
		"name: redis          ,description: 'Redis'              ,module: redis   ,releases: [11,12,   22,24   ]",
		"name: llvm           ,description: 'LLVM'               ,module: llvm    ,releases: [11,12,13,22,24   ]",
		"name: haproxyd       ,description: 'Haproxy Debian'     ,module: haproxy ,releases: [   12            ]",
		"${distro_codename}-backports-3.2 main",
		"name: haproxyu       ,description: 'Haproxy Ubuntu'     ,module: haproxy ,releases: [            24,26]",
		"vbernat/haproxy-3.2/ubuntu/",
		"name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [11,12,13,22,24,26]",
		"name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [11,12,13,22,24,26]",
		"name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [11,12,13,22,24   ]",
		"name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [11,12,13,22,24   ]",
		"name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [11,12,13,22,24,26]",
	}
	for _, fragment := range required {
		if !strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate missing expected compatibility fragment: %q", fragment)
		}
	}

	forbidden := []string{
		"name: wiltondb",
		"repo.percona.com/ppg-18.3",
		"vbernat/haproxy-3.1",
		"${distro_codename}-backports-3.1",
	}
	for _, fragment := range forbidden {
		if strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate still contains incompatible repo fragment: %q", fragment)
		}
	}
}

func TestDebTemplateUbuntu20Removed(t *testing.T) {
	releaseWith20 := regexp.MustCompile(`releases:\s*\[[^]]*\b20\b`)
	if match := releaseWith20.FindString(debTemplate); match != "" {
		t.Fatalf("debTemplate still includes Ubuntu 20.04 in repo releases: %q", match)
	}

	if _, ok := GetConfigConstants().DistroAdhocPkg["u20"]; ok {
		t.Fatal("DistroAdhocPkg still includes u20 bootstrap packages")
	}
}

func TestDebTemplateDebianChinaComponents(t *testing.T) {
	required := []string{
		"name: base           ,description: 'Debian Basic'       ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename} main non-free-firmware'                                  ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main non-free-firmware' }",
		"name: updates        ,description: 'Debian Updates'     ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename}-updates main non-free-firmware'                          ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main non-free-firmware' }",
		"name: security       ,description: 'Debian Security'    ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://security.debian.org/debian-security ${distro_codename}-security main non-free-firmware'            ,china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main non-free-firmware' }",
	}
	for _, fragment := range required {
		if !strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate missing expected Debian china fragment: %q", fragment)
		}
	}

	forbidden := []string{
		"china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main restricted universe multiverse'",
		"china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main restricted universe multiverse'",
		"china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main restricted universe multiverse'",
		"china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main contrib non-free non-free-firmware'",
		"china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main contrib non-free non-free-firmware'",
		"china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main contrib non-free non-free-firmware'",
	}
	for _, fragment := range forbidden {
		if strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate still contains invalid Debian/Ubuntu mixed components: %q", fragment)
		}
	}
}
