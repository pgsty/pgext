/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import "testing"

func TestCleanVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// User-provided test cases
		{"0.2.0+git20211101.d7d10f2", "0.2.0"},
		{"5.5.0+debpgdg", "5.5.0"},
		{"1.2_20240606", "1.2"},
		{"2.22.0+dfsg", "2.22.0"},
		{"1.3.0~alpha", "1.3.0"},

		// Epoch prefix handling
		{"1:2.0.12-3.el8", "2.0.12"},
		{"2:1.2.3", "1.2.3"},

		// Various delimiters
		{"1.0.0-beta", "1.0.0"},
		{"2.5.1~rc1", "2.5.1"},
		{"3.0.0_alpha", "3.0.0"},

		// Edge cases
		{"1.2.3", "1.2.3"},
		{"10", "10"},
		{"1.", "1"},
		{"1.0.", "1.0"},

		// Complex versions
		{"1.2.3.citus-1", "1.2.3"},
		{"2.0.0.pgdg120+1", "2.0.0"},
		{"1.5.0-1PGDG.rhel8", "1.5.0"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := cleanVersion(tt.input)
			if result != tt.expected {
				t.Errorf("cleanVersion(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
