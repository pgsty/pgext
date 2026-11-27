package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// runSubcommand executes a command through its local argument, pre-run, run,
// and post-run hooks while preserving the caller's context. Once PreRun/PreRunE
// succeeds, PostRun/PostRunE always runs even if RunE returns an error.
func runSubcommand(caller, target *cobra.Command, args []string) (retErr error) {
	previousContext := target.Context()
	target.SetContext(caller.Context())
	defer target.SetContext(previousContext)

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

	defer func() {
		if target.PostRunE != nil {
			if postErr := target.PostRunE(target, args); postErr != nil {
				retErr = errors.Join(retErr, postErr)
			}
		} else if target.PostRun != nil {
			target.PostRun(target, args)
		}
	}()

	if target.RunE != nil {
		return target.RunE(target, args)
	}
	if target.Run != nil {
		target.Run(target, args)
	}
	return nil
}

// defaultToSubcommand executes the target child command when the parent is
// invoked without an explicit subcommand.
func defaultToSubcommand(target *cobra.Command) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return runSubcommand(cmd, target, args)
	}
}
