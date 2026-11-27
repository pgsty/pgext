package cmd

import (
	"path/filepath"

	"pgext/cli"

	"github.com/spf13/cobra"
)

func newStubLintCommand() *cobra.Command {
	var (
		pair      bool
		stubDir   string
		zhStubDir string
	)
	cmd := &cobra.Command{
		Use:   "lint [stub-files-or-directories...]",
		Short: "Check stub Markdown or bilingual stub pairs (no database required)",
		Example: `  pgext gen lint
  pgext gen lint stub/acl.md stub-zh/acl.md
  pgext gen lint --pair pgvector postgis`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if pair {
				if stubDir == "" || zhStubDir == "" {
					root, err := cli.SourceDir()
					if err != nil {
						return err
					}
					if stubDir == "" {
						stubDir = filepath.Join(root, "stub")
					}
					if zhStubDir == "" {
						zhStubDir = filepath.Join(root, "stub-zh")
					}
				}
				count, err := cli.LintStubPairs(stubDir, zhStubDir, args)
				if err != nil {
					return err
				}
				cmd.Printf("Bilingual stub pair check passed: %d pairs\n", count)
				return nil
			}
			if len(args) == 0 {
				root, err := cli.SourceDir()
				if err != nil {
					return err
				}
				args = []string{filepath.Join(root, "stub"), filepath.Join(root, "stub-zh")}
			}
			count, err := cli.LintStubPaths(args)
			if err != nil {
				return err
			}
			cmd.Printf("Stub Markdown check passed: %d files\n", count)
			return nil
		},
	}
	cmd.Flags().BoolVar(&pair, "pair", false, "validate bilingual stub/<ext>.md and stub-zh/<ext>.md pairs")
	cmd.Flags().StringVar(&stubDir, "stub-dir", "", "English stub directory for --pair (default: stub/)")
	cmd.Flags().StringVar(&zhStubDir, "zh-stub-dir", "", "Chinese stub directory for --pair (default: stub-zh/)")
	return cmd
}

var stubLintCmd = newStubLintCommand()

func init() {
	genCmd.AddCommand(stubLintCmd)
}
