/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"testing"
)

// TestCleanVersion tests the version cleaning logic
func TestCleanVersion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic versions
		{"simple version", "1.2.3", "1.2.3"},
		{"version with patch", "1.2.3.4", "1.2.3.4"},

		// Epoch prefixes
		{"epoch prefix single digit", "1:2.0.12", "2.0.12"},
		{"epoch prefix double digit", "10:2.0.12", "2.0.12"},
		{"no epoch colon in middle", "1.2:3", "1.2:3"},

		// Special suffixes
		{"citus suffix", "12.1.6.citus", "12.1.6"},
		{"citus suffix with release", "12.1.6.citus-1", "12.1.6"},
		{"pgdg suffix", "3.0.6.pgdg", "3.0.6"},
		{"PGDG suffix", "3.0.6.PGDG", "3.0.6"},

		// Delimiters
		{"dash delimiter", "13.2.0-9PIGSTY.el9", "13.2.0"},
		{"tilde delimiter", "1.2.3~beta1", "1.2.3"},
		{"plus delimiter", "1.2.3+dfsg", "1.2.3"},
		{"underscore with date", "1.2_20240606", "1.2"},

		// Complex cases
		{"epoch and suffix", "1:12.1.6.citus-1", "12.1.6"},
		{"multiple delimiters", "1.2.3-4~beta+dfsg", "1.2.3"},
		{"real package version", "13.2.0-9PIGSTY.el9", "13.2.0"},
		{"pgdg version", "13.2.0-1PGDG.rhel9", "13.2.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanVersion(tt.input)
			if got != tt.expected {
				t.Errorf("cleanVersion(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestCompareSemanticVersions tests semantic version comparison
func TestCompareSemanticVersions(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int // 1 if a > b, -1 if a < b, 0 if equal
	}{
		// Equal versions
		{"equal simple", "1.2.3", "1.2.3", 0},
		{"equal with patch", "1.2.3.4", "1.2.3.4", 0},

		// Major version differences
		{"major greater", "2.0.0", "1.9.9", 1},
		{"major lesser", "1.9.9", "2.0.0", -1},

		// Minor version differences
		{"minor greater", "1.3.0", "1.2.9", 1},
		{"minor lesser", "1.2.9", "1.3.0", -1},

		// Patch version differences
		{"patch greater", "1.2.3", "1.2.2", 1},
		{"patch lesser", "1.2.2", "1.2.3", -1},

		// Different component counts
		{"more components greater", "1.2.3.4", "1.2.3", 1},
		{"fewer components lesser", "1.2.3", "1.2.3.4", -1},

		// Real version examples
		{"citus versions", "13.2.0", "13.1.0", 1},
		{"citus versions 2", "13.1.0", "13.0.4", 1},
		{"citus versions 3", "12.1.6", "12.1.5", 1},

		// Edge cases
		{"empty vs non-empty", "", "1.0.0", -1},
		{"non-empty vs empty", "1.0.0", "", 1},
		{"both empty", "", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareSemanticVersions(tt.a, tt.b)
			// Normalize result to -1, 0, or 1
			if got > 0 {
				got = 1
			} else if got < 0 {
				got = -1
			}
			if got != tt.expected {
				t.Errorf("compareSemanticVersions(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestCompareVersions tests the full version comparison logic
func TestCompareVersions(t *testing.T) {
	g := &BinGenerator{}

	tests := []struct {
		name     string
		a        BinaryPackage
		b        BinaryPackage
		expected int // positive if a > b (a should come first), negative if a < b
	}{
		// Different semantic versions
		{
			"different major versions",
			BinaryPackage{Ver: "2.0.0-1", Version: "2.0.0"},
			BinaryPackage{Ver: "1.9.9-1", Version: "1.9.9"},
			1, // 2.0.0 > 1.9.9
		},
		{
			"citus 13.2.0 vs 13.1.0",
			BinaryPackage{Ver: "13.2.0-1PGDG.rhel9", Version: "13.2.0"},
			BinaryPackage{Ver: "13.1.0-1PGDG.rhel9", Version: "13.1.0"},
			1, // 13.2.0 > 13.1.0
		},

		// Same semantic version, different full strings
		{
			"same version, 9PIGSTY vs 1PGDG",
			BinaryPackage{Ver: "13.1.0-9PIGSTY.el9", Version: "13.1.0"},
			BinaryPackage{Ver: "13.1.0-1PGDG.rhel9", Version: "13.1.0"},
			1, // "13.1.0-9PIGSTY" > "13.1.0-1PGDG" (string comparison)
		},
		{
			"same version, 1PIGSTY vs 1PGDG",
			BinaryPackage{Ver: "13.1.0-1PIGSTY.el9", Version: "13.1.0"},
			BinaryPackage{Ver: "13.1.0-1PGDG.rhel9", Version: "13.1.0"},
			1, // "13.1.0-1PIGSTY" > "13.1.0-1PGDG" (string comparison, PI > PG)
		},
		{
			"same version, 11PIGSTY vs 5PGDG",
			BinaryPackage{Ver: "13.2.0-11PIGSTY.el9", Version: "13.2.0"},
			BinaryPackage{Ver: "13.2.0-5PGDG.rhel9", Version: "13.2.0"},
			-1, // "13.2.0-11PIGSTY" < "13.2.0-5PGDG" (string comparison, "11" < "5")
		},

		// Equal versions
		{
			"identical versions",
			BinaryPackage{Ver: "1.2.3-1", Version: "1.2.3"},
			BinaryPackage{Ver: "1.2.3-1", Version: "1.2.3"},
			0,
		},

		// Complex real-world examples
		{
			"timescaledb versions",
			BinaryPackage{Ver: "2.18.1-1PIGSTY.el9", Version: "2.18.1"},
			BinaryPackage{Ver: "2.18.0-1PGDG.rhel9", Version: "2.18.0"},
			1, // 2.18.1 > 2.18.0
		},
		{
			"same base, different suffixes",
			BinaryPackage{Ver: "3.0.6", Version: "3.0.6"},
			BinaryPackage{Ver: "3.0.6-1PIGSTY", Version: "3.0.6"},
			-1, // "3.0.6" < "3.0.6-1PIGSTY" in string comparison
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := g.compareVersions(tt.a, tt.b)
			// Check sign matches expected
			if (got > 0 && tt.expected <= 0) || (got < 0 && tt.expected >= 0) || (got == 0 && tt.expected != 0) {
				t.Errorf("compareVersions(%v, %v) = %d, want sign %d", tt.a.Ver, tt.b.Ver, got, tt.expected)
			}
		})
	}
}

// TestSortingOrder tests that packages are sorted correctly for ID assignment
func TestSortingOrder(t *testing.T) {
	// Simulate packages for the same pg/os/name combination
	packages := []BinaryPackage{
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "13.0.0-1PGDG.rhel9", Version: "13.0.0"},
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "13.1.0-1PGDG.rhel9", Version: "13.1.0"},
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "13.1.0-9PIGSTY.el9", Version: "13.1.0"},
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "13.2.0-1PGDG.rhel9", Version: "13.2.0"},
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "13.2.0-9PIGSTY.el9", Version: "13.2.0"},
		{PG: 16, OS: "el9.x86_64", Name: "citus_16", Ver: "12.1.6-1PGDG.rhel9", Version: "12.1.6"},
	}

	g := &BinGenerator{}
	g.sortPackages(packages)

	// After sorting, packages should be in descending version order
	// With same semantic versions sorted by string comparison
	expectedOrder := []string{
		"13.2.0-9PIGSTY.el9",  // 13.2.0, "9PIGSTY" > "1PGDG"
		"13.2.0-1PGDG.rhel9",  // 13.2.0
		"13.1.0-9PIGSTY.el9",  // 13.1.0, "9PIGSTY" > "1PGDG"
		"13.1.0-1PGDG.rhel9",  // 13.1.0
		"13.0.0-1PGDG.rhel9",  // 13.0.0
		"12.1.6-1PGDG.rhel9",  // 12.1.6
	}

	for i, expected := range expectedOrder {
		if packages[i].Ver != expected {
			t.Errorf("Position %d: got %s, want %s", i, packages[i].Ver, expected)
		}
	}
}

// TestDatabaseConsistency performs integration tests with actual database data
func TestDatabaseConsistency(t *testing.T) {
	// Skip if no database connection
	if DB == nil {
		t.Skip("Skipping database test - no connection")
	}

	// Test case 1: Verify Citus 13.1.0 ordering
	t.Run("Citus 13.1.0 versions", func(t *testing.T) {
		query := `
			SELECT ver, id
			FROM pgext.bin
			WHERE name = 'citus_16' AND pg = 16 AND os = 'el9.x86_64' AND version = '13.1.0'
			ORDER BY id
		`
		rows, err := Query(query)
		if err != nil {
			t.Fatalf("Query failed: %v", err)
		}
		defer rows.Close()

		var versions []struct {
			Ver string
			ID  int
		}
		for rows.Next() {
			var v struct {
				Ver string
				ID  int
			}
			if err := rows.Scan(&v.Ver, &v.ID); err != nil {
				t.Fatalf("Scan failed: %v", err)
			}
			versions = append(versions, v)
		}

		// Verify ordering: 9PIGSTY should come before 1PGDG (smaller ID)
		if len(versions) >= 2 {
			if versions[0].Ver == "13.1.0-1PGDG.rhel9" && versions[1].Ver == "13.1.0-9PIGSTY.el9" {
				t.Errorf("Wrong order: 1PGDG (ID:%d) before 9PIGSTY (ID:%d)", versions[0].ID, versions[1].ID)
			}
		}
	})

	// Test case 2: Verify version descending order
	t.Run("Version descending order", func(t *testing.T) {
		query := `
			SELECT version, MIN(id) as min_id
			FROM pgext.bin
			WHERE name = 'citus_16' AND pg = 16 AND os = 'el9.x86_64'
			GROUP BY version
			ORDER BY min_id
			LIMIT 5
		`
		rows, err := Query(query)
		if err != nil {
			t.Fatalf("Query failed: %v", err)
		}
		defer rows.Close()

		var prevVersion string
		first := true
		for rows.Next() {
			var version string
			var minID int
			if err := rows.Scan(&version, &minID); err != nil {
				t.Fatalf("Scan failed: %v", err)
			}

			if !first {
				// Each version should be less than or equal to the previous
				cmp := compareSemanticVersions(prevVersion, version)
				if cmp < 0 {
					t.Errorf("Version %s (ID:%d) came before %s but is smaller", prevVersion, minID, version)
				}
			}
			prevVersion = version
			first = false
		}
	})
}