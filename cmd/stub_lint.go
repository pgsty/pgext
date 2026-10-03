package cmd

import (
	"pgext/cli"

	"github.com/spf13/cobra"
)

var stubLintCmd = &cobra.Command{
	Use:   "lint [stub-files-or-directories...]",
	Short: "Check stub Markdown before generating site pages (no database required)",
	Example: `  pgext gen lint
  pgext gen lint stub/acl.md stub-zh/acl.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			args = []string{"stub", "stub-zh"}
		}
		count, err := cli.LintStubPaths(args)
		if err != nil {
			return err
		}
		cmd.Printf("Stub Markdown check passed: %d files\n", count)
		return nil
	},
}

func init() {
	genCmd.AddCommand(stubLintCmd)
}
