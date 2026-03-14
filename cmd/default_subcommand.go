package cmd

import "github.com/spf13/cobra"

// defaultToSubcommand executes the target child command when the parent is
// invoked without an explicit subcommand.
func defaultToSubcommand(target *cobra.Command) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if target.Args != nil {
			if err := target.Args(target, args); err != nil {
				return err
			}
		}
		if target.PreRunE != nil {
			if err := target.PreRunE(target, args); err != nil {
				return err
			}
		} else if target.PreRun != nil {
			target.PreRun(target, args)
		}
		if target.RunE != nil {
			return target.RunE(target, args)
		}
		if target.Run != nil {
			target.Run(target, args)
		}
		return nil
	}
}
