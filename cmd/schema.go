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

// schemaCmd initializes the pgext schema
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "initialize pgext schema",
	Long: `Initialize the pgext schema in the target PostgreSQL database.

This command will:
1. Create the pgext schema
2. Load initial CSV data (pg, os, category, repository, extension)

If the schema already exists, this command will skip initialization.
Use --force to drop and recreate the schema.
Requires the semver extension to be available in the database.`,
	Example: `
  pgext schema                  # initialize schema
  pgext schema -f               # force drop and recreate
  pgext schema -d vonng         # use specific database
  PGURL=meta pgext schema       # use env variable
`,
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
	Long: `Perform complete initialization from scratch:
1. Initialize pgext schema with force
2. Fetch repository metadata
3. Parse repository data into pgext.bin
4. Generate availability info pgext.pkg

This is equivalent to running:
  pgext schema --force
  pgext fetch
  pgext parse
  pgext recap
`,
	Example: `
  pgext init                    # initialize everything
  pgext init -d vonng           # use specific database
  pgext init -p 4               # use 4 parallel workers
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("starting complete initialization...")

		// Step 1: Initialize schema with force
		logrus.Info("step 1/4: initializing schema...")
		if err := cli.InitSchema(true); err != nil {
			return fmt.Errorf("failed to initialize schema: %w", err)
		}

		// Step 2: Fetch repository metadata
		logrus.Info("step 2/4: fetching repository metadata...")
		fetcher := cli.NewFetcher(cli.FetchOptions{
			Force:    true,
			Region:   region,
			Parallel: workers,
			Retry:    retry,
		})
		if err := fetcher.FetchAll(cmd.Context()); err != nil {
			return fmt.Errorf("failed to fetch: %w", err)
		}

		// Step 3: Parse repository data
		logrus.Info("step 3/4: parsing repository data...")
		parser := cli.NewParser(cli.ParseOptions{
			Parallel: workers,
			Keep:     keep,
		})
		if err := parser.ParseAll(); err != nil {
			return fmt.Errorf("failed to parse: %w", err)
		}

		// Step 4: Generate package matrix
		logrus.Info("step 4/4: generating package matrix...")
		if err := cli.RecapMatrix(); err != nil {
			return fmt.Errorf("failed to recap: %w", err)
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
	initCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	initCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	initCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts")
	initCmd.Flags().BoolVarP(&keep, "keep", "k", false, "keep temporary SQLite files")

	// reload command flags
	reloadCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-fetch all repositories")
	reloadCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
	reloadCmd.Flags().IntVarP(&workers, "parallel", "p", 8, "number of parallel workers")
	reloadCmd.Flags().IntVar(&retry, "retry", 1, "number of retry attempts")
	reloadCmd.Flags().BoolVarP(&keep, "keep", "k", false, "keep temporary SQLite files")
}
