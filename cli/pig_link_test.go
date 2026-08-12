/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"strings"
	"testing"
)

func TestPageGeneratorsLinkPigToDedicatedSite(t *testing.T) {
	cache := &ExtensionCache{PGVersions: []int{18}}
	ext := &Extension{Name: "vector", Pkg: "vector", PgVer: []string{"18"}}

	cases := []struct {
		name string
		got  string
		want string
	}{
		{"english", NewIOPageGenerator(cache, "", "").generateInstall(ext), "[**pig**](https://pig.pgsty.com)"},
		{"chinese", NewCCPageGenerator(cache, "", "").generateInstall(ext), "[**pig**](https://pig.pgsty.com/zh)"},
		{"catalog", NewExtensionGenerator(cache, "").generateInstallSection(ext), "[**pig**](https://pig.pgsty.com)"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if !strings.Contains(tc.got, tc.want) {
				t.Fatalf("generated install documentation missing %q", tc.want)
			}
			for _, stale := range []string{"(/docs/pig)", "](/pig)"} {
				if strings.Contains(tc.got, stale) {
					t.Fatalf("generated install documentation contains stale pig link %q", stale)
				}
			}
		})
	}
}
