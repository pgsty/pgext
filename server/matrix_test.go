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
		state   byte
	}{
		{"3.6.3@G", "3.6.3", 'G'},
		{"0.8.5@P", "0.8.5", 'P'},
		{"@M", "", 'M'},
		{"@N", "", 'N'},
		{"@G", "", 'G'},
		{"1.2-3.el9@P", "1.2-3.el9", 'P'},
		{"", "", 'N'},
		{"x@Z", "x", 'N'}, // unknown state degrades to N/A
	}
	for _, tt := range tests {
		version, state := parseMatrixCell(tt.value)
		if version != tt.version || state != tt.state {
			t.Errorf("parseMatrixCell(%q) = (%q,%c), want (%q,%c)",
				tt.value, version, state, tt.version, tt.state)
		}
	}
}

func TestMatrixViews(t *testing.T) {
	if matrixViews[0].name != "" {
		t.Fatalf("the global view must lead the rebuild order, got %q", matrixViews[0].name)
	}
	for _, mv := range matrixViews {
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
		"el8i.17": "1.0@G",
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
