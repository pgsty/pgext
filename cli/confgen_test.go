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

func TestGenerateExtensionMappingsSkipsForkBoundExtensions(t *testing.T) {
	g := &PigstyConfigGenerator{osCode: "el9"}
	extensions := map[string]*ExtensionData{
		"orioledb": {
			ID:         1,
			Name:       "orioledb",
			Alias:      "orioledb",
			Category:   "FEAT",
			CategoryID: 1,
			RPMPkg:     "orioledb_$v",
			RPMPG:      []int{18, 17, 16},
			Available:  map[int]bool{},
			ForkBound:  true,
		},
		"pg_graphql": {
			ID:         2,
			Name:       "pg_graphql",
			Alias:      "pg_graphql",
			Category:   "FEAT",
			CategoryID: 1,
			RPMPkg:     "pg_graphql_$v",
			RPMPG:      []int{18},
			Available:  map[int]bool{},
		},
	}

	mappings := g.generateExtensionMappings(extensions)
	if len(mappings) != 1 {
		t.Fatalf("expected 1 non-fork mapping, got %d: %#v", len(mappings), mappings)
	}
	if mappings[0].Name != "pg_graphql" {
		t.Fatalf("mapping name = %q, want pg_graphql", mappings[0].Name)
	}
}

func TestEL9ARMPatroniUsesNoarchPackageSelectors(t *testing.T) {
	g := &PigstyConfigGenerator{
		osCode:    "el9",
		arch:      "aarch64",
		constants: GetConfigConstants(),
	}
	funcs := g.getFuncMap()

	pgsqlUtility := funcs["getPgsqlUtility"].(func() string)()
	wantPackages := []string{
		"patroni.noarch",
		"patroni-etcd.noarch",
	}
	for _, pkg := range wantPackages {
		if !strings.Contains(pgsqlUtility, pkg) {
			t.Fatalf("pgsql utility package list missing %q: %q", pkg, pgsqlUtility)
		}
	}
	for _, stalePackage := range []string{"patroni-4.1.2", "patroni-etcd-4.1.2", "patroni-4.1.3", "patroni-etcd-4.1.3"} {
		if strings.Contains(pgsqlUtility, stalePackage) {
			t.Fatalf("pgsql utility package list contains stale patroni version: %q", pgsqlUtility)
		}
	}

	getUtilPkg := funcs["getUtilPkg"].(func(string, string) string)
	if got := getUtilPkg("patroni", "patroni patroni-etcd"); got != "patroni.noarch patroni-etcd.noarch" {
		t.Fatalf("patroni util package list = %q", got)
	}
	wantCommon := "patroni.noarch patroni-etcd.noarch pgbouncer pgbackrest pg-exporter pgbackrest-exporter vip-manager"
	if got := getUtilPkg("pgsql-common", "patroni patroni-etcd pgbouncer"); got != wantCommon {
		t.Fatalf("pgsql-common util package list = %q, want %q", got, wantCommon)
	}
}

func TestDebNodePackageAliasesUseBind9DNSUtils(t *testing.T) {
	for _, osCode := range []string{"d12", "d13", "u22", "u24", "u26"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osCode:    osCode,
				constants: GetConfigConstants(),
			}
			funcs := g.getFuncMap()

			packages := fieldPackages(
				funcs["getDebNodePackage1"].(func() string)(),
				funcs["getNodePackage2"].(func() string)(),
				funcs["getNodePackage3"].(func() string)(),
			)
			hasBind9DNSUtils := false
			for _, pkg := range packages {
				if pkg == "bind9-dnsutils" {
					hasBind9DNSUtils = true
				}
				if pkg == "dnsutils" {
					t.Fatalf("node package aliases contain legacy dnsutils package: %q", strings.Join(packages, " "))
				}
			}
			if !hasBind9DNSUtils {
				t.Fatalf("node package aliases missing bind9-dnsutils: %q", strings.Join(packages, " "))
			}
		})
	}
}

func TestRPMNodePackageAliasesUseBindUtils(t *testing.T) {
	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osCode:    osCode,
				constants: GetConfigConstants(),
			}
			funcs := g.getFuncMap()

			packages := fieldPackages(
				funcs["getNodePackage1"].(func() string)(),
				funcs["getNodePackage2"].(func() string)(),
				funcs["getNodePackage3"].(func() string)(),
			)
			hasBindUtils := false
			for _, pkg := range packages {
				if pkg == "bind-utils" {
					hasBindUtils = true
				}
				if pkg == "bind9-utils" || pkg == "bind9.18-utils" {
					t.Fatalf("node package aliases contain nonstandard RPM DNS utilities package: %q", strings.Join(packages, " "))
				}
			}
			if !hasBindUtils {
				t.Fatalf("node package aliases missing bind-utils: %q", strings.Join(packages, " "))
			}
		})
	}
}

func TestRPMNodePackage1IncludesACL(t *testing.T) {
	constants := GetConfigConstants()
	if !containsToken(constants.RPMCommonPkg[3], "acl") {
		t.Fatalf("RPMCommonPkg node-package1 missing acl: %q", constants.RPMCommonPkg[3])
	}

	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osName:    osCode + ".x86_64",
				osCode:    osCode,
				arch:      "x86_64",
				constants: constants,
			}
			funcs := g.getFuncMap()

			nodePackage1 := funcs["getNodePackage1"].(func() string)()
			if !containsToken(nodePackage1, "acl") {
				t.Fatalf("node-package1 missing acl: %q", nodePackage1)
			}
		})
	}

	g := &PigstyConfigGenerator{
		osName:    "el9.x86_64",
		osCode:    "el9",
		arch:      "x86_64",
		constants: constants,
	}
	tmpl := template.Must(template.New("rpm").Funcs(g.getFuncMap()).Parse(rpmTemplate))
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
		"bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd",
		"node-package1:           \"bash python3 sudo acl ca-certificates openssl curl wget lz4 zstd",
		"node-package3:           \"zlib readline xz glibc-langpack-en",
	}
	for _, fragment := range required {
		if !strings.Contains(rendered, fragment) {
			t.Fatalf("rendered RPM config missing expected fragment: %q", fragment)
		}
	}
}

func TestNodePackagesDefaultRenderedAsThreeAlignedGroups(t *testing.T) {
	tests := []struct {
		name     string
		rendered string
		want     []string
	}{
		{
			name:     "rpm",
			rendered: renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
			want: []string{
				"bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd,unzip,bzip2,gzip,tar,tzdata,chrony,openssh-server,util-linux,rsync,psmisc,logrotate",
				"pv,jq,git,make,patch,lsof,less,ncdu,htop,iotop,socat,net-tools,telnet,ipvsadm,tuned,numactl,nvme-cli,sysstat,keepalived,etcd,haproxy,vector,pig,uv",
				"zlib,readline,xz,glibc-langpack-en,cronie,openssh-clients,node-exporter,bind-utils,iproute,iputils,nmap-ncat,procps-ng,vim-minimal,yum,audit,grubby,chkconfig",
			},
		},
		{
			name:     "deb",
			rendered: renderDEBTemplateForTest(t, "u24.x86_64", "u24"),
			want: []string{
				"bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd,unzip,bzip2,gzip,tar,tzdata,chrony,openssh-server,util-linux,rsync,psmisc,logrotate",
				"pv,jq,git,make,patch,lsof,less,ncdu,htop,iotop,socat,net-tools,telnet,ipvsadm,tuned,numactl,nvme-cli,sysstat,keepalived,etcd,haproxy,vector,pig,uv",
				"zlib1g,libreadline-dev,xz-utils,locales,cron,openssh-client,node-exporter,bind9-dnsutils,iproute2,iputils-ping,netcat-openbsd,procps,vim-tiny",
			},
		},
	}

	var commonLines []string
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nodePackagesDefaultLines(tt.rendered)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("node_packages_default lines = %#v, want %#v", got, tt.want)
			}
			if !strings.Contains(tt.rendered, "node-package3:") {
				t.Fatalf("rendered config missing node-package3 alias")
			}
			if commonLines == nil {
				commonLines = got[:2]
			} else if !reflect.DeepEqual(got[:2], commonLines) {
				t.Fatalf("first two node package lines not aligned: got %#v, want %#v", got[:2], commonLines)
			}
		})
	}
}

func TestRPMTemplateAppStreamExcludesDistroPostgreSQLProviders(t *testing.T) {
	const wantMeta = "meta: { excludepkgs: 'postgresql* libpq*' }"

	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			rendered := renderRPMTemplateForTest(t, osCode+".x86_64", osCode)
			appstream := renderedRepoUpstreamLine(rendered, "appstream")
			if appstream == "" {
				t.Fatal("rendered RPM config missing appstream repo")
			}
			if !strings.Contains(appstream, "releases: [8,9,10]") {
				t.Fatalf("appstream repo no longer covers EL8/9/10: %q", appstream)
			}
			if !strings.Contains(appstream, wantMeta) {
				t.Fatalf("appstream repo missing distro PostgreSQL exclude meta: %q", appstream)
			}

			for _, repo := range []string{"baseos", "extras", "powertools", "crb", "epel"} {
				line := renderedRepoUpstreamLine(rendered, repo)
				if line == "" {
					t.Fatalf("rendered RPM config missing %s repo", repo)
				}
				if strings.Contains(line, "excludepkgs") {
					t.Fatalf("%s repo should not carry PostgreSQL exclude meta: %q", repo, line)
				}
			}
		})
	}
}

func TestNodePackagesDefaultPreservesLegacyPackagesExceptTcpdump(t *testing.T) {
	tests := []struct {
		name     string
		rendered string
		legacy   []string
	}{
		{
			name:     "rpm",
			rendered: renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
			legacy: commaPackages(
				"lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync,acl,tcpdump",
				"python3,socat,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,cronie,pig,uv",
				"zlib,yum,audit,bind-utils,readline,vim-minimal,node-exporter,grubby,openssh-server,openssh-clients,chkconfig,vector",
			),
		},
		{
			name:     "deb",
			rendered: renderDEBTemplateForTest(t, "u22.x86_64", "u22"),
			legacy: commaPackages(
				"lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync,tcpdump",
				"python3,socat,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,cron,pig,uv",
				"zlib1g,acl,bind9-dnsutils,libreadline-dev,vim-tiny,node-exporter,openssh-server,openssh-client,vector",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			current := commaPackages(nodePackagesDefaultLines(tt.rendered)...)
			assertNoToken(t, current, "tcpdump")
			assertContainsAllExcept(t, current, tt.legacy, "tcpdump")
		})
	}
}

func TestNodePackageAliasesPreserveLegacyPackagesExceptTcpdump(t *testing.T) {
	constants := GetConfigConstants()
	oldRPMNodePackage1 := "lz4 unzip bzip2 zlib yum pv jq git ncdu make patch bash lsof wget tuned nvme-cli numactl grubby sysstat iotop htop rsync acl tcpdump perf"
	oldRPMNodePackage2 := "netcat socat ftp net-tools ipvsadm bind-utils telnet audit ca-certificates readline vim-minimal keepalived openssl openssh-server openssh-clients chrony cronie"
	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osName:    osCode + ".x86_64",
				osCode:    osCode,
				arch:      "x86_64",
				constants: constants,
			}
			funcs := g.getFuncMap()

			legacyNodePackage1 := oldRPMNodePackage1 + " chkconfig uv"
			if osCode != "el10" {
				legacyNodePackage1 = oldRPMNodePackage1 + " flamegraph chkconfig uv"
			}
			current := fieldPackages(
				funcs["getNodePackage1"].(func() string)(),
				funcs["getNodePackage2"].(func() string)(),
				funcs["getNodePackage3"].(func() string)(),
			)

			assertNoToken(t, current, "tcpdump")
			assertContainsAllExcept(t, current, fieldPackages(legacyNodePackage1, oldRPMNodePackage2), "tcpdump")
		})
	}

	oldDEBNodePackage1 := "lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget tuned nvme-cli numactl sysstat iotop htop rsync tcpdump acl chrony cron uv"
	oldU24NodePackage1 := "lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget tuned nvme-cli numactl sysstat iotop htop rsync acl chrony cron uv"
	oldDEBNodePackage2 := "netcat-openbsd socat net-tools ipvsadm bind9-dnsutils telnet ca-certificates libreadline-dev vim-tiny keepalived openssl openssh-server openssh-client"
	for _, osCode := range []string{"d12", "d13", "u22", "u24", "u26"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osName:    osCode + ".x86_64",
				osCode:    osCode,
				arch:      "x86_64",
				constants: constants,
			}
			funcs := g.getFuncMap()

			legacyNodePackage1 := oldDEBNodePackage1
			if osCode == "u24" {
				legacyNodePackage1 = oldU24NodePackage1
			}
			current := fieldPackages(
				funcs["getDebNodePackage1"].(func() string)(),
				funcs["getNodePackage2"].(func() string)(),
				funcs["getNodePackage3"].(func() string)(),
			)

			assertNoToken(t, current, "tcpdump")
			assertContainsAllExcept(t, current, fieldPackages(legacyNodePackage1, oldDEBNodePackage2), "tcpdump")
		})
	}
}

func TestExtraModulesDownloadAliasExcludesOpenCode(t *testing.T) {
	constants := GetConfigConstants()
	for name, packages := range map[string]string{
		"rpm": constants.RPMCommonPkg[2],
		"deb": constants.DEBCommonPkg[2],
	} {
		t.Run(name, func(t *testing.T) {
			if containsToken(packages, "opencode") {
				t.Fatalf("extra-modules download alias contains opencode: %q", packages)
			}
		})
	}

	for name, rendered := range map[string]string{
		"rpm": renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
		"deb": renderDEBTemplateForTest(t, "u24.x86_64", "u24"),
	} {
		t.Run(name+"-rendered", func(t *testing.T) {
			for _, line := range strings.Split(rendered, "\n") {
				if strings.Contains(line, "extra-modules:") && strings.Contains(line, "opencode") {
					t.Fatalf("rendered extra-modules download alias contains opencode: %q", line)
				}
			}
		})
	}
}

func TestKafkaStackAliasIncludesMonitoringPackages(t *testing.T) {
	tests := []struct {
		name     string
		rendered string
		want     string
	}{
		{
			name:     "rpm",
			rendered: renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
			want:     "kafka kafka-exporter jmx-exporter",
		},
		{
			name:     "deb",
			rendered: renderDEBTemplateForTest(t, "u24.x86_64", "u24"),
			want:     "kafka kafka-exporter jmx-exporter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderedPackageAliasValue(tt.rendered, "kafka-stack"); got != tt.want {
				t.Fatalf("kafka-stack alias = %q, want %q", got, tt.want)
			}
			if got := renderedPackageAliasValue(tt.rendered, "kafka"); got != "" {
				t.Fatalf("rendered config still exposes legacy kafka alias: %q", got)
			}
			if got := renderedPackageAliasValue(tt.rendered, "jmx-exporter"); got != "" {
				t.Fatalf("rendered config exposes standalone jmx-exporter alias: %q", got)
			}
		})
	}
}

func TestMySQL84RepoAndPackageContract(t *testing.T) {
	rpm := renderRPMTemplateForTest(t, "el9.x86_64", "el9")
	deb := renderDEBTemplateForTest(t, "u24.x86_64", "u24")

	rpmMySQL := renderedRepoUpstreamLine(rpm, "mysql")
	for _, fragment := range []string{
		"description: 'MySQL 8.4 LTS'",
		"mysql-8.4-community",
		"mirrors.ustc.edu.cn/mysql-repo",
	} {
		if !strings.Contains(rpmMySQL, fragment) {
			t.Fatalf("RPM mysql repo missing %q: %q", fragment, rpmMySQL)
		}
	}
	for _, repo := range []string{"mysql-tools", "pxb84"} {
		line := renderedRepoUpstreamLine(rpm, repo)
		if line == "" {
			t.Fatalf("RPM config missing %s repository", repo)
		}
		if strings.Contains(line, "meta:") {
			t.Fatalf("RPM %s repo should not carry meta: %q", repo, line)
		}
	}
	if strings.Contains(rpmMySQL, "meta:") {
		t.Fatalf("RPM mysql repo should not carry meta: %q", rpmMySQL)
	}

	debMySQL := renderedRepoUpstreamLine(deb, "mysql")
	for _, fragment := range []string{
		"description: 'MySQL 8.4 LTS'",
		"releases: [   12,13,22,24   ]",
		"mysql-8.4-lts",
		"mirrors.ustc.edu.cn/mysql-repo",
	} {
		if !strings.Contains(debMySQL, fragment) {
			t.Fatalf("DEB mysql repo missing %q: %q", fragment, debMySQL)
		}
	}
	if strings.Contains(debMySQL, "meta:") {
		t.Fatalf("DEB mysql repo should not carry meta: %q", debMySQL)
	}
	if strings.Contains(debMySQL, "mysql-tools") || strings.Contains(debMySQL, "mysql-8.0") {
		t.Fatalf("DEB mysql repo still enables a non-8.4 server/tools component: %q", debMySQL)
	}
	debPXB := renderedRepoUpstreamLine(deb, "pxb84")
	if !strings.Contains(debPXB, "repo.percona.com/pxb-84-lts/apt") {
		t.Fatalf("DEB pxb84 repo missing upstream URL: %q", debPXB)
	}
	if strings.Contains(debPXB, "meta:") {
		t.Fatalf("DEB pxb84 repo should not carry meta: %q", debPXB)
	}

	if got := renderedPackageAliasValue(rpm, "mysql"); got != "mysql-community-server mysql-community-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter" {
		t.Fatalf("RPM mysql alias = %q", got)
	}
	if got := renderedPackageAliasValue(deb, "mysql"); got != "mysql-server mysql-client mysql-shell mysql-router-community percona-xtrabackup-84 mysqld-exporter" {
		t.Fatalf("DEB mysql alias = %q", got)
	}
}

func TestInfraPackageNamesAndLegacyAliases(t *testing.T) {
	legacyAliases := map[string]string{
		"blackbox_exporter":   "blackbox-exporter",
		"kafka_exporter":      "kafka-exporter",
		"keepalived_exporter": "keepalived-exporter",
		"mongodb_exporter":    "mongodb-exporter",
		"mysqld_exporter":     "mysqld-exporter",
		"nginx_exporter":      "nginx-exporter",
		"node_exporter":       "node-exporter",
		"pg_exporter":         "pg-exporter",
		"pgbackrest_exporter": "pgbackrest-exporter",
		"redis_exporter":      "redis-exporter",
		"zfs_exporter":        "zfs-exporter",
	}

	for name, rendered := range map[string]string{
		"rpm": renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
		"deb": renderDEBTemplateForTest(t, "u24.x86_64", "u24"),
	} {
		t.Run(name, func(t *testing.T) {
			for oldName, newName := range legacyAliases {
				if got := renderedPackageAliasValue(rendered, oldName); got != newName {
					t.Errorf("legacy alias %s = %q, want %q", oldName, got, newName)
				}
			}
		})
	}
}

func TestRPMNodePackageBaselineIncludesMinimalNodePackages(t *testing.T) {
	constants := GetConfigConstants()
	for _, osCode := range []string{"el8", "el9", "el10"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osName:    osCode + ".x86_64",
				osCode:    osCode,
				arch:      "x86_64",
				constants: constants,
			}
			funcs := g.getFuncMap()

			nodePackage1 := funcs["getNodePackage1"].(func() string)()
			nodePackage2 := funcs["getNodePackage2"].(func() string)()
			nodePackage3 := funcs["getNodePackage3"].(func() string)()

			for _, pkg := range []string{
				"sudo", "curl", "zstd", "tzdata", "glibc-langpack-en", "less", "iproute",
				"iputils", "nmap-ncat", "procps-ng", "util-linux", "tar", "gzip", "xz",
				"psmisc", "logrotate",
			} {
				if !containsToken(nodePackage1, pkg) && !containsToken(nodePackage2, pkg) && !containsToken(nodePackage3, pkg) {
					t.Fatalf("EL %s node package aliases missing %q:\nnode-package1=%q\nnode-package2=%q\nnode-package3=%q", osCode, pkg, nodePackage1, nodePackage2, nodePackage3)
				}
			}
		})
	}
}

func TestEL7NodePackage2IncludesIPRoute(t *testing.T) {
	g := &PigstyConfigGenerator{
		osName:    "el7.x86_64",
		osCode:    "el7",
		arch:      "x86_64",
		constants: GetConfigConstants(),
	}
	funcs := g.getFuncMap()

	nodePackage2 := funcs["getNodePackage2"].(func() string)()
	if !containsToken(nodePackage2, "iproute") {
		t.Fatalf("EL7 node-package2 missing iproute: %q", nodePackage2)
	}

	rendered := renderRPMTemplateForTest(t, "el7.x86_64", "el7")
	renderedNodePackage2 := renderedPackageAliasValue(rendered, "node-package2")
	if !containsToken(renderedNodePackage2, "iproute") {
		t.Fatalf("rendered EL7 node-package2 missing iproute: %q", renderedNodePackage2)
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
		"bash,python3,sudo,acl,ca-certificates,openssl,curl,wget,lz4,zstd",
		"node-package2:           \"pv jq git make patch lsof less ncdu",
		"node-package3:           \"zlib1g libreadline-dev xz-utils locales",
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

func TestDebNodePackageBaselineIncludesMinimalNodePackages(t *testing.T) {
	constants := GetConfigConstants()
	for _, osCode := range []string{"d12", "d13", "u22", "u24", "u26"} {
		t.Run(osCode, func(t *testing.T) {
			g := &PigstyConfigGenerator{
				osName:    osCode + ".x86_64",
				osCode:    osCode,
				arch:      "x86_64",
				constants: constants,
			}
			funcs := g.getFuncMap()

			nodePackage1 := funcs["getDebNodePackage1"].(func() string)()
			nodePackage2 := funcs["getNodePackage2"].(func() string)()
			nodePackage3 := funcs["getNodePackage3"].(func() string)()

			for _, pkg := range []string{
				"sudo", "curl", "zstd", "tzdata", "locales", "less", "iproute2",
				"iputils-ping", "netcat-openbsd", "procps", "util-linux", "tar", "gzip",
				"xz-utils", "psmisc", "logrotate",
			} {
				if !containsToken(nodePackage1, pkg) && !containsToken(nodePackage2, pkg) && !containsToken(nodePackage3, pkg) {
					t.Fatalf("DEB %s node package aliases missing %q:\nnode-package1=%q\nnode-package2=%q\nnode-package3=%q", osCode, pkg, nodePackage1, nodePackage2, nodePackage3)
				}
			}
		})
	}
}

func renderRPMTemplateForTest(t *testing.T, osName, osCode string) string {
	t.Helper()
	g := &PigstyConfigGenerator{
		osName:    osName,
		osCode:    osCode,
		arch:      "x86_64",
		constants: GetConfigConstants(),
	}
	tmpl := template.Must(template.New("rpm").Funcs(g.getFuncMap()).Parse(rpmTemplate))
	return renderConfigTemplateForTest(t, tmpl, g)
}

func renderDEBTemplateForTest(t *testing.T, osName, osCode string) string {
	t.Helper()
	g := &PigstyConfigGenerator{
		osName:    osName,
		osCode:    osCode,
		arch:      "x86_64",
		constants: GetConfigConstants(),
	}
	tmpl := template.Must(template.New("deb").Funcs(g.getFuncMap()).Parse(debTemplate))
	return renderConfigTemplateForTest(t, tmpl, g)
}

func renderConfigTemplateForTest(t *testing.T, tmpl *template.Template, g *PigstyConfigGenerator) string {
	t.Helper()
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
	return buf.String()
}

func nodePackagesDefaultLines(rendered string) []string {
	lines := strings.Split(rendered, "\n")
	for i, line := range lines {
		if line != "node_packages_default:" {
			continue
		}
		var packages []string
		for _, candidate := range lines[i+1:] {
			if !strings.HasPrefix(candidate, "  - ") {
				break
			}
			packages = append(packages, strings.TrimPrefix(candidate, "  - "))
		}
		return packages
	}
	return nil
}

func commaPackages(lines ...string) []string {
	var tokens []string
	for _, line := range lines {
		for _, token := range strings.Split(line, ",") {
			token = strings.TrimSpace(token)
			if token != "" {
				tokens = append(tokens, token)
			}
		}
	}
	return tokens
}

func fieldPackages(groups ...string) []string {
	var tokens []string
	for _, group := range groups {
		tokens = append(tokens, strings.Fields(group)...)
	}
	return tokens
}

func renderedPackageAliasValue(rendered, alias string) string {
	prefix := alias + ":"
	for _, line := range strings.Split(rendered, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			parts := strings.SplitN(line, "\"", 3)
			if len(parts) >= 3 {
				return parts[1]
			}
		}
	}
	return ""
}

func renderedRepoUpstreamLine(rendered, name string) string {
	prefix := "- { name: " + name
	for _, line := range strings.Split(rendered, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, prefix) {
			return line
		}
	}
	return ""
}

func assertContainsAllExcept(t *testing.T, current, legacy []string, except string) {
	t.Helper()
	currentSet := tokenSet(current)
	for _, token := range legacy {
		if token == except {
			continue
		}
		if !currentSet[token] {
			t.Fatalf("package list removed legacy package %q; current=%q", token, strings.Join(current, " "))
		}
	}
}

func assertNoToken(t *testing.T, current []string, token string) {
	t.Helper()
	if tokenSet(current)[token] {
		t.Fatalf("package list still contains removable package %q; current=%q", token, strings.Join(current, " "))
	}
}

func tokenSet(tokens []string) map[string]bool {
	set := make(map[string]bool, len(tokens))
	for _, token := range tokens {
		set[token] = true
	}
	return set
}

func containsToken(s, token string) bool {
	for _, field := range strings.Fields(s) {
		if field == token {
			return true
		}
	}
	return false
}

func TestConfigConstantsIncludeBabelfishAlias(t *testing.T) {
	constants := GetConfigConstants()
	for _, mapping := range constants.PGSQLExoticMap {
		if mapping.Key != "babelfish" {
			continue
		}
		if mapping.RPM != "babelfish-$v" {
			t.Fatalf("babelfish RPM alias = %q", mapping.RPM)
		}
		if mapping.DEB != "babelfish-$v" {
			t.Fatalf("babelfish DEB alias = %q", mapping.DEB)
		}
		return
	}
	t.Fatal("babelfish mapping not found in PGSQLExoticMap")
}

func TestConfigConstantsIncludeAgensAndPgEdgeAlias(t *testing.T) {
	constants := GetConfigConstants()
	expected := map[string]PackageMapping{
		"agensgraph": {
			Key: "agensgraph",
			RPM: "agensgraph-$v",
			DEB: "agensgraph-$v",
		},
		"pgedge": {
			Key: "pgedge",
			RPM: "pgedge-$v",
			DEB: "pgedge-$v",
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

func TestConfigConstantsIncludePolarDBAlias(t *testing.T) {
	constants := GetConfigConstants()
	expected := map[string]PackageMapping{
		"polardb": {
			Key: "polardb",
			RPM: "polardb-$v",
			DEB: "polardb-$v",
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

func TestConfigConstantsIncludeOrioleVersionTemplate(t *testing.T) {
	constants := GetConfigConstants()
	for _, mapping := range constants.PGSQLExoticMap {
		if mapping.Key != "orioledb" {
			continue
		}
		if mapping.RPM != "orioledb-$v" {
			t.Fatalf("oriole RPM alias = %q", mapping.RPM)
		}
		if mapping.DEB != "orioledb-$v" {
			t.Fatalf("oriole DEB alias = %q", mapping.DEB)
		}
		return
	}
	t.Fatal("oriole mapping not found in PGSQLExoticMap")
}

func TestConfigConstantsIncludeCloudberryFullPackageAlias(t *testing.T) {
	constants := GetConfigConstants()
	for _, mapping := range constants.PGSQLExoticMap {
		if mapping.Key != "cloudberry" {
			continue
		}
		want := "cloudberry cloudberry-backup cloudberry-pxf"
		if mapping.RPM != want {
			t.Fatalf("cloudberry RPM alias = %q, want %q", mapping.RPM, want)
		}
		if mapping.DEB != want {
			t.Fatalf("cloudberry DEB alias = %q, want %q", mapping.DEB, want)
		}
		return
	}
	t.Fatal("cloudberry mapping not found in PGSQLExoticMap")
}

func TestConfigConstantsIncludePgHardstorageToolAlias(t *testing.T) {
	constants := GetConfigConstants()
	expected := map[string]PackageMapping{
		"pg-hardstorage": {
			Key: "pg-hardstorage",
			RPM: "pg-hardstorage",
			DEB: "pg-hardstorage",
		},
		"pg_hardstorage": {
			Key: "pg_hardstorage",
			RPM: "pg-hardstorage",
			DEB: "pg-hardstorage",
		},
	}

	found := make(map[string]PackageMapping)
	for _, mapping := range constants.PGSQLUtilMap {
		if _, ok := expected[mapping.Key]; ok {
			found[mapping.Key] = mapping
		}
	}

	for key, want := range expected {
		got, ok := found[key]
		if !ok {
			t.Fatalf("%s mapping not found in PGSQLUtilMap", key)
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
		"gpsql:  '/usr/cloudberry'",
		"polar:  '/usr/polar-$v'",
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

func TestRenderedTemplatesIncludePolarDBPackageMap(t *testing.T) {
	for name, rendered := range map[string]string{
		"rpm": renderRPMTemplateForTest(t, "el9.x86_64", "el9"),
		"deb": renderDEBTemplateForTest(t, "u24.x86_64", "u24"),
	} {
		for _, fragment := range []string{
			`polardb:                 "polardb-$v"`,
			`polar:  '/usr/polar-$v'`,
		} {
			if !strings.Contains(rendered, fragment) {
				t.Fatalf("%s rendered template missing expected PolarDB fragment: %q", name, fragment)
			}
		}
		for _, fragment := range []string{
			`polardb-for-postgresql`,
			`PolarDB"`,
			`polar:  '/u01/polardb_pg_17'`,
			`polar:  '/usr/polar-17'`,
			`polar:                   "polardb-$v"`,
			`oracle:`,
			`'/u01/polardb_pg'`,
		} {
			if strings.Contains(rendered, fragment) {
				t.Fatalf("%s rendered template still contains old PolarDB fragment: %q", name, fragment)
			}
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
		"name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [11,12,13,22,24,26]",
		"name: citus          ,description: 'Citus'              ,module: extra   ,releases: [11,12,   22      ]",
		"name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [   12,13,22,24,26]",
		"name: groonga        ,description: 'Groonga Ubuntu'     ,module: groonga ,releases: [         22,24   ]",
		"name: mysql          ,description: 'MySQL 8.4 LTS'      ,module: mysql   ,releases: [   12,13,22,24   ] ,arch: [x86_64         ]",
		"name: pxb84          ,description: 'Percona XtraBackup' ,module: mysql   ,releases: [   12,13,22,24   ]",
		"name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [   12,   22,24   ]",
		"name: redis          ,description: 'Redis'              ,module: redis   ,releases: [11,12,   22,24,26]",
		"name: llvm           ,description: 'LLVM'               ,module: llvm    ,releases: [11,12,13,22,24,26]",
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

func TestRPMTemplateEL10RepoReleaseCompatibility(t *testing.T) {
	required := []string{
		"name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: mysql          ,description: 'MySQL 8.4 LTS'      ,module: mysql   ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: mysql-tools    ,description: 'MySQL 8.4 Tools'    ,module: mysql   ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: pxb84          ,description: 'Percona XtraBackup' ,module: mysql   ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
		"name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [8,9,10] ,arch: [x86_64, aarch64]",
	}
	for _, fragment := range required {
		if !strings.Contains(rpmTemplate, fragment) {
			t.Fatalf("rpmTemplate missing expected EL10 repo fragment: %q", fragment)
		}
	}

	forbidden := []string{
		"name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [8,9   ] ,arch: [x86_64, aarch64]",
		"name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [   10] ,arch: [x86_64         ]",
	}
	for _, fragment := range forbidden {
		if strings.Contains(rpmTemplate, fragment) {
			t.Fatalf("rpmTemplate still advertises unsupported EL10 TimescaleDB arch: %q", fragment)
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
