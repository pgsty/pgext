package cli

import (
	"reflect"
	"strings"
	"testing"
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

func TestDebTemplateUbuntuRepoSuffixesMatchRepoName(t *testing.T) {
	required := []string{
		"name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-updates",
		"name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-backports",
		"name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-security",
		"name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-updates",
		"name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-backports",
		"name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-security",
	}
	for _, fragment := range required {
		if !strings.Contains(debTemplate, fragment) {
			t.Fatalf("debTemplate missing expected fragment: %q", fragment)
		}
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
