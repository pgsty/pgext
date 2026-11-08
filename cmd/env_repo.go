package cmd

import (
	"github.com/spf13/cobra"
	"pgext/cli/build"
)

func newEnvRepoCommand() *cobra.Command {
	var opts build.Options
	var asJSON, update bool
	c := &cobra.Command{Use: "repo", Short: "prepare or maintain builder package repositories", Args: cobra.NoArgs, SilenceUsage: true}
	f := c.PersistentFlags()
	f.BoolVarP(&opts.Mirror, "mirror", "m", false, "use the Pigsty China mirror")
	f.BoolVar(&opts.Plan, "plan", false, "preview without changing repositories")
	f.BoolVar(&asJSON, "json", false, "emit JSON")
	f.BoolVar(&update, "update", true, "refresh package metadata after changes")
	f.BoolVar(&opts.ReplaceRepos, "replace", false, "back up and replace recognized PIG infra/pgsql repository files")
	run := func(action string) func(*cobra.Command, []string) error {
		return func(cmd *cobra.Command, args []string) error {
			e, err := build.NewEnvironment(opts, cmd.ErrOrStderr())
			if err != nil {
				return err
			}
			ctx, stop := buildContext(cmd)
			defer stop()
			r, runErr := e.Repo(ctx, action, args, update)
			if err := printBuildValue(cmd, r, asJSON); err != nil {
				return err
			}
			return runErr
		}
	}
	c.RunE = run("add")
	for _, action := range []string{"add", "rm", "list", "status", "update"} {
		sub := &cobra.Command{Use: action, Short: action + " builder repositories", Args: cobra.NoArgs, RunE: run(action), SilenceUsage: true}
		if action == "add" || action == "rm" {
			sub.Use += " [node|pgdg|pigsty...]"
			sub.Args = cobra.ArbitraryArgs
		}
		c.AddCommand(sub)
	}
	return c
}
