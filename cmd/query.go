/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"fmt"
	"pgext/cli"

	"github.com/spf13/cobra"
)

// statusCmd shows metadata status
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show metadata status",
	Long: `Display current metadata status including:
- Active PostgreSQL versions
- Active Linux distributions
- Table record counts
- Last update timestamps`,
	Example: `
  pgext status                  # show status
  pgext status -d vonng         # check specific database
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.ShowStatus(); err != nil {
			return fmt.Errorf("failed to show status: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// pkgCmd shows package availability matrix
var pkgCmd = &cobra.Command{
	Use:   "pkg <name>",
	Short: "show package availability matrix",
	Long: `Display package availability matrix for a given extension.

Shows which PostgreSQL versions and operating systems have the extension
available, along with version and repository information.`,
	Example: `
  pgext pkg pgvector             # show pgvector availability
  pgext pkg timescaledb          # show timescaledb availability
`,
	Args:    cobra.ExactArgs(1),
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := cli.ShowPackage(name); err != nil {
			return fmt.Errorf("failed to show package: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// binCmd shows binary package information
var binCmd = &cobra.Command{
	Use:   "bin <name>",
	Short: "show binary package information with download URLs",
	Long: `Display binary package information including download URLs.

Can be filtered by PostgreSQL version and operating system.
Supports region-specific URLs (default or china/mirror).`,
	Example: `
  pgext bin pgvector              # show all pgvector packages
  pgext bin pgvector -p 17        # only PG 17 packages
  pgext bin pgvector -o el9       # only EL9 packages
  pgext bin pgvector -p 17 -o el9 # specific combination
  pgext bin pgvector -r china     # use China mirror URLs
`,
	Args:    cobra.ExactArgs(1),
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := cli.ShowBin(name, pgVer, osFilter, region); err != nil {
			return fmt.Errorf("failed to show bin: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// repoCmd shows repository summary
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "show repository summary",
	Long: `Display repository summary including:
- Repository ID and type
- Operating system
- Organization
- Data size
- Status (has data or missing)`,
	Example: `
  pgext repo                    # show all repositories
`,
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cli.ShowRepositories(); err != nil {
			return fmt.Errorf("failed to show repositories: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

// extCmd shows extension information
var extCmd = &cobra.Command{
	Use:   "ext <name>",
	Short: "show extension information",
	Long: `Display detailed extension information including:
- Package name
- Category
- Repository
- Version
- Description
- Requirements
- URLs (website, documentation, source)`,
	Example: `
  pgext ext pgvector              # show pgvector info
  pgext ext timescaledb           # show timescaledb info
`,
	Args:    cobra.ExactArgs(1),
	PreRunE: initDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := cli.ShowExt(name); err != nil {
			return fmt.Errorf("failed to show extension: %w", err)
		}
		return nil
	},
	PostRun: closeDatabase,
}

func init() {
	// bin command flags
	binCmd.Flags().IntVarP(&pgVer, "pg", "p", 0, "PostgreSQL major version (e.g., 17)")
	binCmd.Flags().StringVarP(&osFilter, "os", "o", "", "OS filter (e.g., el9, u24)")
	binCmd.Flags().StringVarP(&region, "region", "r", "", "region: default or china/mirror")
}
