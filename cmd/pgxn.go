/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"fmt"
	"pgext/cli"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	pgxnOutputDir string
	pgxnCatalog   string
)

var pgxnCmd = &cobra.Command{
	Use:   "pgxn",
	Short: "crawl PGXN metadata into standalone tables",
	Long: `Crawl the full PGXN distribution catalog without touching the current
pgext database tables or db/*.csv source files.

The command:
1. Discovers all distributions from api.pgxn.org/src/
2. Fetches dist.json and latest META.json for each distribution
3. Matches PGXN distributions against db/extension.csv
4. Writes standalone CSV/JSONL/Markdown reports into a separate output directory`,
	Example: `
  pgext pgxn
  pgext pgxn -p 24
  pgext pgxn --output research/pgxn
  pgext pgxn --catalog db/extension.csv
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cli.ExportPGXNCatalog(cmd.Context(), cli.PgxnExportOptions{
			OutputDir:   pgxnOutputDir,
			CatalogPath: pgxnCatalog,
			Workers:     workers,
			Retry:       retry,
		})
		if err != nil {
			return err
		}

		logrus.Infof("PGXN export complete: %d dist, %d matched, %d missing, %d failed",
			result.DistCount, result.MatchedCount, result.MissingCount, result.FailedCount)
		for label, path := range result.Files {
			logrus.Infof("  %s => %s", label, path)
		}
		if result.FailedCount > 0 {
			logrus.Warnf("some distributions failed; inspect %s for details", result.Files["summary_json"])
		}
		return nil
	},
}

func init() {
	pgxnCmd.Flags().StringVarP(&pgxnOutputDir, "output", "o", "research/pgxn", "output directory for generated PGXN artifacts")
	pgxnCmd.Flags().StringVar(&pgxnCatalog, "catalog", "db/extension.csv", "existing pgext extension catalog CSV used for matching")
	pgxnCmd.Flags().IntVarP(&workers, "parallel", "p", 16, "number of parallel workers")
	pgxnCmd.Flags().IntVar(&retry, "retry", 2, "number of retry attempts per PGXN request")

	rootCmd.AddCommand(pgxnCmd)
	rootCmd.Example += fmt.Sprintf("  %s\n", "pgext pgxn                    # crawl PGXN dist metadata into standalone tables")
}
