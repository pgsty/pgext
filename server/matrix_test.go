/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"reflect"
	"testing"
)

func TestParseMatrixCell(t *testing.T) {
	tests := []struct {
		value   string
		version string
		hidden  bool
		state   byte
	}{
		{"3.6.3@G", "3.6.3", false, 'G'},
		{"0.8.5#P", "0.8.5", true, 'P'},
		{"@M", "", false, 'M'},
		{"#N", "", true, 'N'},
		{"@G", "", false, 'G'},
		{"1.2-3.el9@P", "1.2-3.el9", false, 'P'},
		{"", "", false, 'N'},
		{"x@Z", "x", false, 'N'}, // unknown state degrades to N/A
	}
	for _, tt := range tests {
		version, hidden, state := parseMatrixCell(tt.value)
		if version != tt.version || hidden != tt.hidden || state != tt.state {
			t.Errorf("parseMatrixCell(%q) = (%q,%v,%c), want (%q,%v,%c)",
				tt.value, version, hidden, state, tt.version, tt.hidden, tt.state)
		}
	}
}

func TestMatrixViews(t *testing.T) {
	if matrixViews[0].name != "" {
		t.Fatalf("the global view must lead the rebuild order, got %q", matrixViews[0].name)
	}
	seen := map[int]string{}
	for _, mv := range matrixViews {
		if prev, dup := seen[mv.id]; dup {
			t.Fatalf("cache id %d shared by %q and %q", mv.id, prev, mv.name)
		}
		seen[mv.id] = mv.name
		if mv.id < 1 || mv.id > 3 {
			t.Fatalf("view %q cache id %d outside the matrix_cache_id_range check", mv.name, mv.id)
		}
		if mv.name == "" && mv.letter != "" {
			t.Fatalf("global view must not carry a status letter, got %q", mv.letter)
		}
		if mv.name != "" && mv.letter != "G" && mv.letter != "P" {
			t.Fatalf("view %q letter %q is not a known AVAIL status byte", mv.name, mv.letter)
		}
		if got, ok := matrixViewOf(mv.name); !ok || got != mv {
			t.Fatalf("matrixViewOf(%q) = %+v, %v", mv.name, got, ok)
		}
	}
	if _, ok := matrixViewOf("bogus"); ok {
		t.Fatal("matrixViewOf must reject unknown views")
	}
}

func TestMatrixColumnKeys(t *testing.T) {
	keys := matrixColumnKeys([]string{"el8.x86_64", "el8.aarch64", "d12.x86_64"}, []int{18, 17})
	want := []string{"el8i.18", "el8i.17", "el8a.18", "el8a.17", "d12i.18", "d12i.17"}
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf("matrixColumnKeys = %v, want %v", keys, want)
	}
}

func TestLayoutMatrixRow(t *testing.T) {
	keys := matrixColumnKeys([]string{"el8.x86_64", "el8.aarch64"}, []int{18, 17})
	cells := map[string]string{
		"el8i.18": "1.0@G",
		"el8i.17": "1.0#G", // hidden flag is ignored: same state, same version slot
		"el8a.18": "2.0@P",
		"el8a.17": "@M",
	}
	counts := map[string]int{}
	row := layoutMatrixRow(cells, keys, counts)
	if row.Codes != "GGPM" {
		t.Fatalf("codes = %q, want GGPM", row.Codes)
	}
	if !reflect.DeepEqual(row.Versions, []string{"1.0", "2.0"}) {
		t.Fatalf("versions = %v", row.Versions)
	}
	if row.VIdx != "001." {
		t.Fatalf("vidx = %q, want 001.", row.VIdx)
	}
	if counts["G"] != 2 || counts["P"] != 1 || counts["M"] != 1 {
		t.Fatalf("counts = %v", counts)
	}

	// absent keys become N/A without a version slot
	row = layoutMatrixRow(map[string]string{}, keys, map[string]int{})
	if row.Codes != "NNNN" || row.VIdx != "" || row.Versions != nil {
		t.Fatalf("empty layout = %+v", row)
	}
}
