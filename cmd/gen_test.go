package cmd

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

type defaultCommandContextKey struct{}

func TestGeneratorCommandMounts(t *testing.T) {
	for _, tt := range []struct {
		name   string
		cmd    *cobra.Command
		parent *cobra.Command
	}{
		{name: "gen cc", cmd: ccCmd, parent: genCmd},
		{name: "gen io", cmd: ioCmd, parent: genCmd},
		{name: "gen all", cmd: allCmd, parent: genCmd},
	} {
		if got := tt.cmd.Parent(); got != tt.parent {
			t.Errorf("%s parent = %v, want %s", tt.name, got, tt.parent.CommandPath())
		}
	}

	if genCmd.RunE == nil {
		t.Error("gen should delegate to all when invoked without a subcommand")
	}
	for _, legacyArg := range []string{"typo"} {
		if err := genCmd.Args(genCmd, []string{legacyArg}); err == nil {
			t.Errorf("gen accepted unexpected argument %q", legacyArg)
		}
	}
	if err := allCmd.Args(allCmd, []string{"typo"}); err == nil {
		t.Error("gen all accepted an unexpected positional argument")
	}
	for _, cmd := range rootCmd.Commands() {
		if cmd == ccCmd || cmd == ioCmd || cmd == allCmd {
			t.Errorf("%s remains mounted directly under root", cmd.Name())
		}
	}
}

func TestCatalogOutputFlagsAreScopedToCatalogGenerators(t *testing.T) {
	for _, cmd := range []*cobra.Command{genCmd, allCmd} {
		if flag := cmd.Flags().Lookup("output"); flag != nil {
			t.Errorf("%s unexpectedly exposes --output", cmd.CommandPath())
		}
		if flag := cmd.PersistentFlags().Lookup("output"); flag != nil {
			t.Errorf("%s unexpectedly defines persistent --output", cmd.CommandPath())
		}
	}

	for _, cmd := range []*cobra.Command{genPageCmd, genListCmd, genOSCmd, genMatrixCmd} {
		flag := cmd.Flags().Lookup("output")
		if flag == nil {
			t.Errorf("%s does not expose --output", cmd.CommandPath())
			continue
		}
		if flag.Shorthand != "o" {
			t.Errorf("%s --output shorthand = %q, want o", cmd.CommandPath(), flag.Shorthand)
		}
		if flag.DefValue != defaultCatalogOutputDir {
			t.Errorf("%s --output default = %q, want %q", cmd.CommandPath(), flag.DefValue, defaultCatalogOutputDir)
		}
	}
	staticOutput := genMatrixCmd.Flags().Lookup("static-output")
	if staticOutput == nil {
		t.Fatal("gen matrix does not expose --static-output")
	}
	if staticOutput.DefValue != "" {
		t.Fatalf("gen matrix --static-output default = %q, want empty derived default", staticOutput.DefValue)
	}
}

func TestResolveMatrixStaticOutputDir(t *testing.T) {
	for _, tt := range []struct {
		name     string
		content  string
		explicit string
		want     string
	}{
		{name: "repository defaults", content: "content", want: "static"},
		{name: "custom content root", content: "/tmp/site/content", want: "/tmp/site/static"},
		{name: "explicit override", content: "/tmp/site/content", explicit: "/tmp/assets", want: "/tmp/assets"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveMatrixStaticOutputDir(tt.content, tt.explicit); got != tt.want {
				t.Fatalf("resolveMatrixStaticOutputDir(%q, %q) = %q, want %q", tt.content, tt.explicit, got, tt.want)
			}
		})
	}
}

func TestGenListArgsAreCaseInsensitive(t *testing.T) {
	if err := genListCmd.Args(genListCmd, []string{"EXT", "License"}); err != nil {
		t.Fatalf("uppercase list types rejected: %v", err)
	}
	if err := genListCmd.Args(genListCmd, []string{"unknown"}); err == nil {
		t.Fatal("unknown list type accepted")
	}
}

func TestRunSubcommandLifecycle(t *testing.T) {
	caller := &cobra.Command{Use: "caller"}
	target := &cobra.Command{Use: "target"}
	caller.SetContext(context.WithValue(context.Background(), defaultCommandContextKey{}, "caller"))
	target.SetContext(context.WithValue(context.Background(), defaultCommandContextKey{}, "target"))

	var order []string
	checkContext := func(cmd *cobra.Command) {
		if got := cmd.Context().Value(defaultCommandContextKey{}); got != "caller" {
			t.Fatalf("context value = %v, want caller", got)
		}
	}
	target.Args = func(cmd *cobra.Command, args []string) error {
		checkContext(cmd)
		order = append(order, "args")
		return nil
	}
	target.PreRunE = func(cmd *cobra.Command, args []string) error {
		checkContext(cmd)
		order = append(order, "pre")
		return nil
	}
	target.RunE = func(cmd *cobra.Command, args []string) error {
		checkContext(cmd)
		order = append(order, "run")
		return nil
	}
	target.PostRunE = func(cmd *cobra.Command, args []string) error {
		checkContext(cmd)
		order = append(order, "post")
		return nil
	}

	if err := runSubcommand(caller, target, nil); err != nil {
		t.Fatalf("runSubcommand failed: %v", err)
	}
	if got, want := strings.Join(order, ","), "args,pre,run,post"; got != want {
		t.Fatalf("hook order = %q, want %q", got, want)
	}
	if got := target.Context().Value(defaultCommandContextKey{}); got != "target" {
		t.Fatalf("target context after run = %v, want target", got)
	}
}

func TestAllCommandRunsEveryGenerator(t *testing.T) {
	originalRunner := runAllGeneratorCommand
	t.Cleanup(func() { runAllGeneratorCommand = originalRunner })

	sentinel := errors.New("page failed")
	var got []*cobra.Command
	runAllGeneratorCommand = func(caller, target *cobra.Command, args []string) error {
		got = append(got, target)
		if target == genPageCmd {
			return sentinel
		}
		return nil
	}

	err := allCmd.RunE(allCmd, nil)
	if !errors.Is(err, sentinel) {
		t.Fatalf("all error = %v, want wrapped page error", err)
	}

	want := []*cobra.Command{
		ccCmd,
		ioCmd,
		genPageCmd,
		genListCmd,
		genOSCmd,
		genMatrixCmd,
		genConfCmd,
	}
	if len(got) != len(want) {
		t.Fatalf("all ran %d commands, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("all command %d = %s, want %s", i, got[i].CommandPath(), want[i].CommandPath())
		}
	}
}

func TestGeneratorExamplesUseMountedPaths(t *testing.T) {
	var visit func(*cobra.Command)
	visit = func(cmd *cobra.Command) {
		if strings.Contains(cmd.Example, "pgext cc") {
			t.Errorf("%s example contains legacy pgext cc path", cmd.CommandPath())
		}
		if strings.Contains(cmd.Example, "pgext io") {
			t.Errorf("%s example contains legacy pgext io path", cmd.CommandPath())
		}
		if strings.Contains(cmd.Example, "pgext all") {
			t.Errorf("%s example contains invalid root-level pgext all path", cmd.CommandPath())
		}
		for _, child := range cmd.Commands() {
			visit(child)
		}
	}
	visit(rootCmd)
}

func TestGeneratorCommandsRunAllSubcommandByDefault(t *testing.T) {
	tests := []struct {
		name   string
		parent *cobra.Command
		target *cobra.Command
	}{
		{
			name:   "gen",
			parent: genCmd,
			target: allCmd,
		},
		{
			name:   "cc",
			parent: ccCmd,
			target: ccAllCmd,
		},
		{
			name:   "io",
			parent: ioCmd,
			target: ioAllCmd,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			originalRunE := tt.target.RunE
			originalParentContext := tt.parent.Context()
			originalTargetContext := tt.target.Context()
			tt.parent.SetContext(context.WithValue(context.Background(), defaultCommandContextKey{}, tt.name))
			tt.target.RunE = func(cmd *cobra.Command, args []string) error {
				called = true
				if got := cmd.Context().Value(defaultCommandContextKey{}); got != tt.name {
					t.Fatalf("context value = %v, want %q", got, tt.name)
				}
				if len(args) != 0 {
					t.Fatalf("expected no args, got %v", args)
				}
				return nil
			}
			defer func() {
				tt.target.RunE = originalRunE
				tt.parent.SetContext(originalParentContext)
				tt.target.SetContext(originalTargetContext)
			}()

			if tt.parent.RunE == nil {
				t.Fatalf("%s command has no default RunE", tt.name)
			}
			if err := tt.parent.RunE(tt.parent, nil); err != nil {
				t.Fatalf("%s default RunE failed: %v", tt.name, err)
			}
			if !called {
				t.Fatalf("%s did not delegate to all subcommand", tt.name)
			}
		})
	}
}
