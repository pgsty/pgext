package cmd

import (
	"strings"
	"testing"
)

func TestInitCommandForceFlag(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"init"})
	if err != nil {
		t.Fatalf("find init command: %v", err)
	}
	if cmd != initCmd {
		t.Fatalf("expected init command, got %#v", cmd)
	}

	flag := cmd.Flags().Lookup("force")
	if flag == nil {
		t.Fatal("init command does not register --force")
	}
	if flag.Shorthand != "f" {
		t.Fatalf("--force shorthand = %q, want %q", flag.Shorthand, "f")
	}
	if flag.DefValue != "false" {
		t.Fatalf("--force default = %q, want false", flag.DefValue)
	}
	if flag.NoOptDefVal != "true" {
		t.Fatalf("bare --force value = %q, want true", flag.NoOptDefVal)
	}

	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		t.Fatalf("read --force: %v", err)
	}
	if force {
		t.Fatal("init command must preserve the existing schema by default")
	}
}

func TestInitForceFlagDoesNotMutateOtherCommands(t *testing.T) {
	flag := initCmd.Flags().Lookup("force")
	if flag == nil {
		t.Fatal("init command does not register --force")
	}

	originalInitForce := initForce
	originalForce := force
	originalValue := flag.Value.String()
	originalChanged := flag.Changed
	t.Cleanup(func() {
		_ = flag.Value.Set(originalValue)
		flag.Changed = originalChanged
		initForce = originalInitForce
		force = originalForce
	})

	initForce = false
	force = false
	if err := flag.Value.Set("true"); err != nil {
		t.Fatalf("set init --force: %v", err)
	}
	if !initForce {
		t.Fatal("init --force is not bound to initForce")
	}
	if force {
		t.Fatal("init --force must not mutate the force flag used by other commands")
	}
}

func TestInitCommandDocumentsForceBehavior(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"init"})
	if err != nil {
		t.Fatalf("find init command: %v", err)
	}

	flag := cmd.Flags().Lookup("force")
	if flag == nil {
		t.Fatal("init command does not register --force")
	}
	for _, want := range []string{"force", "drop", "recreate"} {
		if !strings.Contains(strings.ToLower(flag.Usage), want) {
			t.Fatalf("--force help missing %q: %s", want, flag.Usage)
		}
	}

	for _, want := range []string{
		"preserved by default",
		"Use --force to drop and recreate",
		"pgext schema [--force]",
	} {
		if !strings.Contains(cmd.Long, want) {
			t.Fatalf("init long help missing %q: %s", want, cmd.Long)
		}
	}

	for _, want := range []string{
		"pgext init                    # initialize everything",
		"pgext init -f                 # drop and recreate schema, then reload",
	} {
		if !strings.Contains(cmd.Example, want) {
			t.Fatalf("init example missing %q: %s", want, cmd.Example)
		}
	}
}
