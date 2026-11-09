package build

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLegacyRPMBuilderRepositories(t *testing.T) {
	// Repository definitions from pgsty/el9a:build, Rocky 9.8.
	for _, name := range []string{"infra.repo", "pgsql.repo"} {
		t.Run(name, func(t *testing.T) {
			b, err := os.ReadFile(filepath.Join("testdata", "legacy-pig-rpm", name))
			if err != nil {
				t.Fatal(err)
			}
			content := string(b)
			for _, tc := range []struct {
				name, content string
				want          bool
			}{
				{"original", content, true},
				{"signature-check-enabled", strings.ReplaceAll(content, "gpgcheck=0", "gpgcheck=1"), true},
				{"key-whitespace", strings.ReplaceAll(content, "gpgkey=", "gpgkey = "), true},
				{"custom-baseurl", strings.ReplaceAll(content, "https://repo.pigsty.cc/", "https://private.example/"), false},
				{"lookalike-host", strings.ReplaceAll(content, "https://repo.pigsty.cc/", "https://repo.pigsty.cc.evil.example/"), false},
				{"custom-section", content + "\n[custom]\nbaseurl=https://private.example/repo\n", false},
				{"mirrorlist", content + "mirrorlist=https://private.example/list\n", false},
				{"exclude", content + "exclude=postgresql18*\n", false},
				{"includepkgs", content + "includepkgs=postgresql18*\n", false},
			} {
				t.Run(tc.name, func(t *testing.T) {
					if got := recognizedLegacyRepos(tc.content, "rpm"); got != tc.want {
						t.Fatalf("recognizedLegacyRepos() = %t, want %t", got, tc.want)
					}
				})
			}
		})
	}
}

func TestLegacyRPMPGDGSigningKey(t *testing.T) {
	// The signing-key field is also standard metadata on PGDG repositories.
	content := "[pgdg18]\n" +
		"baseurl=https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9.8-$basearch\n" +
		"gpgcheck=1\n" +
		"gpgkey=file:///etc/pki/rpm-gpg/PGDG-RPM-GPG-KEY-AARCH64-RHEL\n"
	if !recognizedLegacyRepos(content, "rpm") {
		t.Fatal("rejected PGDG repository with a signing-key declaration")
	}
}
