package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestPipelineCommandsRejectUnexpectedArgs(t *testing.T) {
	for _, command := range []*cobraCommandAdapter{
		{name: "init", args: initCmd.Args, cmd: initCmd},
		{name: "schema", args: schemaCmd.Args, cmd: schemaCmd},
		{name: "purge", args: purgeCmd.Args, cmd: purgeCmd},
		{name: "reload", args: reloadCmd.Args, cmd: reloadCmd},
		{name: "rescan", args: rescanCmd.Args, cmd: rescanCmd},
		{name: "fetch", args: fetchCmd.Args, cmd: fetchCmd},
		{name: "scan", args: scanCmd.Args, cmd: scanCmd},
		{name: "parse", args: parseCmd.Args, cmd: parseCmd},
		{name: "recap", args: recapCmd.Args, cmd: recapCmd},
	} {
		t.Run(command.name, func(t *testing.T) {
			if err := command.args(command.cmd, []string{"unexpected"}); err == nil {
				t.Fatalf("%s accepted an unexpected positional argument", command.name)
			}
		})
	}
}

func TestPipelineBestEffortFlags(t *testing.T) {
	for _, command := range []struct {
		name string
		cmd  flagLookup
	}{
		{"init", initCmd},
		{"reload", reloadCmd},
		{"rescan", rescanCmd},
		{"fetch", fetchCmd},
		{"scan", scanCmd},
		{"parse", parseCmd},
	} {
		flag := command.cmd.Flags().Lookup("best-effort")
		if flag == nil {
			t.Errorf("%s does not register --best-effort", command.name)
		} else if flag.DefValue != "false" {
			t.Errorf("%s --best-effort default = %q, want false", command.name, flag.DefValue)
		}
	}
}

func TestParseNoPkgFlag(t *testing.T) {
	flag := parseCmd.Flags().Lookup("no-pkg")
	if flag == nil {
		t.Fatal("parse does not register --no-pkg")
	}
	if flag.DefValue != "false" {
		t.Fatalf("parse --no-pkg default = %q, want false", flag.DefValue)
	}
	if flag.Shorthand != "N" {
		t.Fatalf("parse --no-pkg shorthand = %q, want N", flag.Shorthand)
	}
}

func TestKeepTempCompatibilityFlags(t *testing.T) {
	for _, command := range []struct {
		name string
		cmd  flagLookup
	}{
		{"init", initCmd},
		{"reload", reloadCmd},
		{"rescan", rescanCmd},
		{"parse", parseCmd},
	} {
		keepTemp := command.cmd.Flags().Lookup("keep-temp")
		if keepTemp == nil || keepTemp.Shorthand != "k" {
			t.Errorf("%s --keep-temp/-k is not registered correctly", command.name)
		}
		legacy := command.cmd.Flags().Lookup("keep")
		if legacy == nil || legacy.Deprecated == "" {
			t.Errorf("%s --keep is not a deprecated compatibility alias", command.name)
		}
	}
}

// Small local interfaces keep the table-driven tests focused on Cobra's
// command contract without constructing another root command.
type flagLookup interface {
	Flags() *pflag.FlagSet
}

type cobraCommandAdapter struct {
	name string
	cmd  *cobra.Command
	args cobra.PositionalArgs
}
