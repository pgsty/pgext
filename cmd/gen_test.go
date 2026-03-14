package cmd

import "testing"

func TestGenCommandIncludesAllSubcommand(t *testing.T) {
	t.Helper()

	for _, subcmd := range genCmd.Commands() {
		if subcmd == genAllCmd {
			return
		}
	}

	t.Fatal("gen command does not include all subcommand")
}
