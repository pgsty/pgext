package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestGenCommandIncludesAllSubcommand(t *testing.T) {
	t.Helper()

	for _, subcmd := range genCmd.Commands() {
		if subcmd == genAllCmd {
			return
		}
	}

	t.Fatal("gen command does not include all subcommand")
}

func TestDefaultCommandsRunAllSubcommand(t *testing.T) {
	tests := []struct {
		name   string
		parent *cobra.Command
		target *cobra.Command
	}{
		{
			name:   "gen",
			parent: genCmd,
			target: genAllCmd,
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
			tt.target.RunE = func(cmd *cobra.Command, args []string) error {
				called = true
				if len(args) != 0 {
					t.Fatalf("expected no args, got %v", args)
				}
				return nil
			}
			defer func() { tt.target.RunE = originalRunE }()

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
