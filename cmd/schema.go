/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"fmt"
	"pgext/cli"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	initForce      bool
	initBestEffort bool
	initKeepTemp   bool
)

// schemaCmd initializes the pgext schema
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "initialize pgext schema",
	Long: `Initialize the pgext schema in the target PostgreSQL database.

This command will:
1. Create the pgext schema
2. Load initial CSV data (pg, os, category, repository, extension, universe)

If the schema already exists, this command preserves its data and applies
supported non-destructive schema additions (currently pgext.universe).
Use --force to drop and recreate the entire schema.
Requires the semver extension to be available in the database.`,
	Example: `
  pgext schema                  # initialize schema
  pgext schema -f               # force drop and recreate
  pgext schema -d vonng         # use specific database
  PGURL=meta pgext schema       # use env variable
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.InitSchema(force); err != nil {
			return fmt.Errorf("failed to initialize schema: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// initCmd performs complete initialization: schema + reload
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "complete setup: schema + reload",
	Long: `Perform complete initialization:
1. Initialize the pgext schema if it does not exist
2. Fetch repository metadata
3. Parse repository data into pgext.apt, pgext.dnf, and pgext.bin
4. Generate availability info pgext.pkg

Existing schema data is preserved by default.
Use --force to drop and recreate the pgext schema before reloading.

This runs the same logical stages as:
  pgext schema [--force]
  pgext fetch
  pgext parse

The parsed package tables and availability matrix are published atomically.
`,
	Example: `
  pgext init                    # initialize everything
  pgext init -f                 # drop and recreate schema, then reload
  pgext init -d vonng           # use specific database
  pgext init -p 4               # use 4 parallel workers
  pgext init --best-effort      # allow a partial catalog if some repos fail
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("starting complete initialization...")

		// Step 1: Initialize schema, preserving existing data unless forced
		logrus.Info("step 1/4: initializing schema...")
		if err := cli.InitSchema(initForce); err != nil {
			return fmt.Errorf("failed to initialize schema: %w", err)
		}

		// Step 2: Fetch repository metadata
		logrus.Info("step 2/4: fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:      true,
			Region:     region,
			Parallel:   workers,
			Retry:      retry,
			BestEffort: initBestEffort,
		})
		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		// Steps 3-4: Parse repository data and build the package matrix
		logrus.Info("steps 3-4/4: parsing repository data and generating package matrix...")
		parser := cli.NewParserContext(cmd.Context(), cli.ParseOptions{
			Parallel:   workers,
			KeepTemp:   initKeepTemp,
			BestEffort: initBestEffort,
		})
		if err := parser.ParseAndRecap(); err != nil {
			return fmt.Errorf("failed to build package catalog: %w", err)
		}

		logrus.Info("initialization completed successfully!")
		return nil
	},
	PostRun: closeDatabase,
}

// purgeCmd removes the pgext schema
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "remove pgext schema completely",
	Long: `Remove the pgext schema and all its data from the database.

WARNING: This operation is destructive and cannot be undone!
All extension metadata, package information, and cached repository
data will be permanently deleted.`,
	Example: `
  pgext purge                   # remove pgext schema
`,
	Args:    cobra.NoArgs,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("⚠️  WARNING: This will permanently delete the pgext schema and all data!")
		fmt.Print("Type 'yes' to confirm: ")

		var confirm string
		fmt.Scanln(&confirm)

		if confirm != "yes" {
			fmt.Println("Purge cancelled.")
			return nil
		}

		if err := cli.PurgeSchema(); err != nil {
			return fmt.Errorf("failed to purge schema: %w", err)
		}

		fmt.Println("✅ pgext schema purged successfully")
		return nil
	},
	PostRun: closeDatabase,
}

func init() {
	// schema command flags
	schemaCmd.Flags().BoolVarP(&force, "force", "f", false, "force drop and recreate schema")

	// init command flags
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "force drop and recreate schema before reload")
	initCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	initCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	initCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts")
	initCmd.Flags().BoolVarP(&initKeepTemp, "keep-temp", "k", false, "keep temporary DNF SQLite files")
	initCmd.Flags().BoolVar(&initKeepTemp, "keep", false, "deprecated alias for --keep-temp")
	_ = initCmd.Flags().MarkDeprecated("keep", "use --keep-temp")
	initCmd.Flags().BoolVar(&initBestEffort, "best-effort", false, "publish a partial catalog when some repositories fail")

	// reload command flags
	reloadCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-fetch all repositories")
	reloadCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	reloadCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	reloadCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts")
	reloadCmd.Flags().BoolVarP(&reloadKeepTemp, "keep-temp", "k", false, "keep temporary DNF SQLite files")
	reloadCmd.Flags().BoolVar(&reloadKeepTemp, "keep", false, "deprecated alias for --keep-temp")
	_ = reloadCmd.Flags().MarkDeprecated("keep", "use --keep-temp")
	reloadCmd.Flags().BoolVar(&reloadBestEffort, "best-effort", false, "publish a partial catalog when some repositories fail")
}
