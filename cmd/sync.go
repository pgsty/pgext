package cmd

import (
	"fmt"
	"time"

	"pgext/cli"

	"github.com/spf13/cobra"
)

func newSyncCommand() *cobra.Command {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Synchronize external upstream catalog metadata",
	}

	var opts cli.GitHubSyncOptions
	var delaySec float64
	var refreshAfter string

	ghCmd := &cobra.Command{
		Use:   "github",
		Short: "Refresh GitHub repository metrics in pgext.gh_repo",
		Args:  cobra.NoArgs,
		Example: `  pgext sync github --ext vector --dry-run
  pgext sync github --stale-days 7 --backfill
  pgext sync github --backfill-only --refresh-after 2026-11-27T00:00:00Z`,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, flag := range []string{"ext", "url", "status"} {
				if cmd.Flags().Changed(flag) {
					values, _ := cmd.Flags().GetStringSlice(flag)
					if len(values) == 0 {
						return fmt.Errorf("--%s requires at least one value", flag)
					}
				}
			}
			if cmd.Flags().Changed("refresh-after") {
				cutoff, err := time.Parse(time.RFC3339, refreshAfter)
				if err != nil {
					return fmt.Errorf("--refresh-after: %w", err)
				}
				opts.RefreshAfter = cutoff
			}
			if err := initDatabase(cmd, args); err != nil {
				return err
			}
			defer closeDatabase(cmd, args)
			opts.Delay = time.Duration(delaySec * float64(time.Second))
			summary, err := cli.SyncGitHub(cmd.Context(), cli.DB, opts)
			if err != nil {
				return err
			}
			cmd.Printf("GitHub sync complete: targets=%d fetched=%d blocked=%d errors=%d universe_updated=%d ext_star_updated=%d\n",
				summary.Targets, summary.Fetched, summary.Blocked, summary.Errors,
				summary.UniverseUpdated, summary.ExtensionUpdated)
			if summary.StoppedOnReserve {
				return fmt.Errorf("stopped early on GitHub rate-limit reserve (core=%d, graphql=%d)",
					summary.CoreRemaining, summary.GraphQLRemaining)
			}
			return nil
		},
	}

	ghCmd.Flags().StringSliceVar(&opts.Extensions, "ext", nil, "limit to specific extension names")
	ghCmd.Flags().StringSliceVar(&opts.URLs, "url", nil, "limit to specific GitHub repository URLs")
	ghCmd.Flags().StringSliceVar(&opts.Statuses, "status", nil, "limit to specific gh_repo statuses (missing, pending, fetched, error, blocked)")
	ghCmd.Flags().IntVar(&opts.StaleDays, "stale-days", 0, "skip fetched repositories updated within N days")
	ghCmd.Flags().IntVar(&opts.Limit, "limit", 0, "maximum number of repositories to fetch")
	ghCmd.Flags().IntVar(&opts.Reserve, "reserve", 200, "stop when REST or GraphQL remaining budget drops to N")
	ghCmd.Flags().Float64Var(&delaySec, "delay", 0.1, "delay in seconds between repositories")
	ghCmd.Flags().StringVar(&opts.Token, "token", "", "GitHub token (default: GITHUB_TOKEN, GH_TOKEN, or gh auth token)")
	ghCmd.Flags().StringVar(&opts.Proxy, "proxy", "", "HTTP/HTTPS proxy URL (default: environment proxy)")
	ghCmd.Flags().BoolVar(&opts.Backfill, "backfill", false, "backfill metrics for repositories fetched by this run")
	ghCmd.Flags().BoolVar(&opts.BackfillOnly, "backfill-only", false, "backfill from pgext.gh_repo without GitHub API calls (requires --refresh-after)")
	ghCmd.Flags().StringVar(&refreshAfter, "refresh-after", "", "earliest accepted fetch time for --backfill-only (RFC3339)")
	ghCmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "preview selected targets without calling GitHub or writing to DB")

	syncCmd.AddCommand(ghCmd)
	return syncCmd
}
