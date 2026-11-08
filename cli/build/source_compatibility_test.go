package build

import (
	"slices"
	"strings"
	"testing"
)

func TestSourceFallbackDoesNotOverrideCatalog(t *testing.T) {
	e := testEngine(t)
	j := &Job{Package: "rdkit", Extension: &Extension{Source: "rdkit-future.tar.gz future-dependency.tar.gz"}}
	for _, family := range []string{"", "rpm", "deb"} {
		e.Platform.Family = family
		if got := e.sourceNames(j, nil); !slices.Equal(got, strings.Fields(j.Extension.Source)) {
			t.Fatalf("%s: hardcoded sources override catalog: %v", family, got)
		}
	}
	if got := e.sourceNames(j, []byte("TARBALL := rdkit-recipe.tar.gz\n")); !slices.Equal(got, []string{"rdkit-recipe.tar.gz"}) {
		t.Fatal("local recipe must remain authoritative", got)
	}
	j.Extension = nil
	if got := e.sourceNames(j, nil); !slices.Equal(got, SpecialSourceMapping["rdkit"]) {
		t.Fatal("missing fallback", got)
	}
}
