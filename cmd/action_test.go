package cmd

import (
	"strings"
	"testing"
)

func TestRootCommandIncludesRescan(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"rescan"})
	if err != nil {
		t.Fatalf("find rescan command: %v", err)
	}
	if cmd == nil || cmd.Use != "rescan" {
		t.Fatalf("expected rescan command, got %#v", cmd)
	}
}

func TestRescanCommandDocumentsLocalPipeline(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"rescan"})
	if err != nil {
		t.Fatalf("find rescan command: %v", err)
	}
	if cmd == nil {
		t.Fatal("rescan command not found")
	}

	for _, want := range []string{
		"pgext scan",
		"pgext parse",
		"pgext recap",
	} {
		if !strings.Contains(cmd.Long, want) {
			t.Fatalf("rescan long help missing %q: %s", want, cmd.Long)
		}
	}
}
