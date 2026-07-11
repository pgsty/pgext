package cmd

import "testing"

func TestExtJSONFlagContract(t *testing.T) {
	flag := extCmd.Flags().Lookup("json")
	if flag == nil {
		t.Fatal("ext --json flag is not registered")
	}
	if flag.DefValue != "false" {
		t.Fatalf("ext --json default = %q, want false", flag.DefValue)
	}
}
