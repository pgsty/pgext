/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pgext/cli"
	"pgext/db"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	genDir         string
	genPageStubDir string
)

// genCmd represents the gen command.
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate extension catalogs, documentation, matrices, and configuration",
	Long: `Generate PostgreSQL extension artifacts from the pgext database.

Use cc and io for the Chinese and English Pigsty documentation sites. The
page, list, os, and matrix commands generate the PGEXT catalog site, while
conf generates Pigsty package configuration. Run pgext gen site to refresh the
static catalog, or pgext gen all to execute every generator.`,
	Example: `  pgext gen site        # Refresh the Hugo catalog at ~/pgsty/pgext.cloud
  pgext gen cc          # Generate all pigsty.cc extension content
  pgext gen io          # Generate all pigsty.io extension content
  pgext gen page        # Generate PGEXT extension detail pages
  pgext gen list        # Generate PGEXT list pages
  pgext gen os          # Generate PGEXT OS availability pages
  pgext gen matrix      # Generate the global OS/PG availability matrix
  pgext gen conf        # Generate Pigsty configuration files
  pgext gen all         # Run every generator`,
	Args: cobra.NoArgs,
	RunE: defaultToSubcommand(allCmd),
}

// genListCmd generates various list pages.
var genListCmd = &cobra.Command{
	Use:   "list [types...]",
	Short: "Generate list pages for extensions",
	Long: `Generate list pages for extensions. If no types are specified, generates all types.
Available types: ext, pkg, cate, lang, license, catalog`,
	Example: `  pgext gen list              # Generate all list pages
  pgext gen list ext          # Generate extension list pages only
  pgext gen list ext pkg      # Generate extension and package list pages
  pgext gen list cate lang    # Generate category and language list pages`,
	ValidArgs: []string{"ext", "pkg", "cate", "lang", "license", "catalog"},
	Args:      onlyValidArgsCaseInsensitive,
	RunE: func(cmd *cobra.Command, args []string) error {
		contentDir, err := catalogContentDir(cmd)
		if err != nil {
			return err
		}
		return runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			listDir := filepath.Join(contentDir, "list")
			extDir := filepath.Join(contentDir, "e")
			for _, dir := range []string{listDir, extDir} {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return fmt.Errorf("failed to create output directory: %w", err)
				}
			}

			selected := make(map[string]bool)
			if len(args) == 0 {
				for _, t := range cmd.ValidArgs {
					selected[t] = true
				}
				logrus.Info("Generating all list types...")
			} else {
				for _, arg := range args {
					selected[strings.ToLower(arg)] = true
				}
				logrus.Infof("Generating list types: %v", args)
			}

			listGen := cli.NewListGenerator(cache, listDir)
			extGen := &cli.ExtensionGenerator{DB: cli.DB, Cache: cache}

			type listTask struct {
				key   string
				label string
				stem  string
				gen   func(locale, path string) error
			}
			tasks := []listTask{
				{"ext", "extension list", "ext", extGen.GenerateExtensionList},
				{"pkg", "package list", "pkg", extGen.GeneratePackageList},
				{"cate", "category list", "cate", listGen.GenerateCategoryList},
				{"lang", "language list", "lang", listGen.GenerateLanguageList},
				{"license", "license list", "license", listGen.GenerateLicenseList},
				{"catalog", "catalog", "_index", extGen.GenerateCatalogPage},
			}

			var errs []error
			successCount := 0
			for _, task := range tasks {
				if !selected[task.key] {
					continue
				}
				logrus.Infof("Generating %s pages...", task.label)
				for _, loc := range []struct{ code, suffix, name string }{
					{"en", ".md", "English"},
					{"zh", ".zh.md", "Chinese"},
				} {
					p := filepath.Join(listDir, task.stem+loc.suffix)
					if err := task.gen(loc.code, p); err != nil {
						errs = append(errs, fmt.Errorf("failed to generate %s %s: %w", loc.name, task.label, err))
					} else {
						logrus.Infof("Generated: %s", p)
						successCount++
					}
				}
				if task.key == "ext" {
					enIdx, zhIdx := filepath.Join(extDir, "_index.md"), filepath.Join(extDir, "_index.zh.md")
					if err := extGen.GenerateExtensionIndexPages(enIdx, zhIdx); err != nil {
						errs = append(errs, fmt.Errorf("failed to generate extension index pages: %w", err))
					} else {
						logrus.Infof("Generated: %s", enIdx)
						logrus.Infof("Generated: %s", zhIdx)
						successCount += 2
					}
				}
			}

			logrus.Infof("List generation completed: %d files generated", successCount)
			if len(errs) > 0 {
				logrus.Errorf("Encountered %d errors during generation:", len(errs))
				for _, e := range errs {
					logrus.Error(e)
				}
				return fmt.Errorf("list generation completed with %d errors", len(errs))
			}
			return nil
		})(cmd, args)
	},
}

// genOSCmd generates OS-specific availability pages.
var genOSCmd = &cobra.Command{
	Use:   "os [os-name]",
	Short: "Generate OS-specific availability page(s)",
	Long: `Generate markdown page(s) showing extension availability for specific OS distribution(s).
If no argument is provided, generates pages for all active OS distributions.`,
	Example: `  pgext gen os               # Generate pages for all active OS distributions
  pgext gen os el9.x86_64    # Generate page for RHEL 9 x86_64
  pgext gen os u24.x86_64    # Generate page for Ubuntu 24.04 x86_64
  pgext gen os el10.aarch64  # Generate page for RHEL 10 ARM64`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		contentDir, err := catalogContentDir(cmd)
		if err != nil {
			return err
		}
		return runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			osDir := filepath.Join(contentDir, "os")
			if err := os.MkdirAll(osDir, 0755); err != nil {
				return fmt.Errorf("failed to create os output directory: %w", err)
			}
			generator := cli.NewOSGenerator(cache, osDir)
			if len(args) == 1 {
				logrus.Infof("Generating OS page for %s...", args[0])
				if err := generator.GenerateOSPage(ctx, args[0]); err != nil {
					return fmt.Errorf("failed to generate OS page for %s: %w", args[0], err)
				}
				logrus.Infof("Successfully generated OS page: %s.md", args[0])
				return nil
			}
			return generator.GenerateAllOSPages(ctx)
		})(cmd, args)
	},
}

// genMatrixCmd generates the global OS x PG package availability matrix.
var genMatrixCmd = &cobra.Command{
	Use:   "matrix",
	Short: "Generate global OS/PG availability matrix",
	Long: `Generate a package-oriented global availability matrix across all active
operating systems and PostgreSQL major versions. Hugo pages are written below
--output, while CSV/JSON/HTML assets are written below --static-output. When
--static-output is omitted, it defaults to a static directory alongside the
content output directory.`,
	Example: `  pgext gen matrix
  pgext gen matrix -d data
  pgext gen matrix -o /tmp/content --static-output /tmp/static`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		contentDir, err := catalogContentDir(cmd)
		if err != nil {
			return err
		}
		staticOverride, _ := cmd.Flags().GetString("static-output")
		staticOutputDir := resolveMatrixStaticOutputDir(contentDir, staticOverride)
		return runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, _ []string) error {
			generator := cli.NewGlobalMatrixGenerator(cache, contentDir, staticOutputDir)
			if err := generator.Generate(ctx); err != nil {
				return fmt.Errorf("failed to generate global matrix: %w", err)
			}
			logrus.Infof("Generated global matrix under %s and %s", filepath.Join(contentDir, "os"), filepath.Join(staticOutputDir, "matrix"))
			return nil
		})(cmd, args)
	},
}

func resolveMatrixStaticOutputDir(contentOutputDir, explicitStaticOutputDir string) string {
	if explicitStaticOutputDir != "" {
		return explicitStaticOutputDir
	}
	return filepath.Join(filepath.Dir(contentOutputDir), "static")
}

func onlyValidArgsCaseInsensitive(cmd *cobra.Command, args []string) error {
	normalized := make([]string, len(args))
	for i, arg := range args {
		normalized[i] = strings.ToLower(arg)
	}
	return cobra.OnlyValidArgs(cmd, normalized)
}

// genPageCmd generates individual extension detail pages.
var genPageCmd = &cobra.Command{
	Use:   "page [extension-names...]",
	Short: "Generate individual extension detail pages",
	Long: `Generate Hugo markdown detail pages for PostgreSQL extensions.
If no extension names are provided, generates pages for all extensions.`,
	Example: `  pgext gen page              # Generate pages for all extensions
  pgext gen page vector       # Generate page for pgvector's vector extension
  pgext gen page vector postgis pg_stat_statements  # Generate pages for specific extensions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		contentDir, err := catalogContentDir(cmd)
		if err != nil {
			return err
		}
		return runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			extDir := filepath.Join(contentDir, "e")
			if err := os.MkdirAll(extDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}
			stubDir, err := cli.SourceStubDir(genPageStubDir, "stub")
			if err != nil {
				return err
			}
			exts, err := selectExtensions(cache, args)
			if err != nil {
				return err
			}
			generator := cli.NewExtensionGenerator(cache, extDir)
			generator.StubDir = stubDir
			return generateExtensionPages(ctx, stubDir, exts, len(args) > 0, extDir, generator.GenerateExtensionPage)
		})(cmd, args)
	},
}

// allGeneratorCommands defines the complete generation pipeline in execution order.
var allGeneratorCommands = []*cobra.Command{
	ccCmd,
	ioCmd,
	genSiteCmd,
	genConfCmd,
}

// runAllGeneratorCommand is replaceable in tests so the orchestration contract
// can be verified without running database-backed generators.
var runAllGeneratorCommand = runSubcommand

// allCmd runs every generator under the gen command.
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run every content and configuration generator",
	Long: `Run the complete generation pipeline in this order:

  pgext gen cc
  pgext gen io
  pgext gen site (page, list, os, matrix and catalog CSVs)
  pgext gen conf

Each generator uses its normal default output directory. All generators are
attempted, and the command returns a combined error if any step fails.`,
	Example: `  pgext gen all           # Run every generator
  pgext gen all -d data   # Use the local data database`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var runErrors []error
		for _, target := range allGeneratorCommands {
			path := target.CommandPath()
			logrus.Infof("Running %s...", path)
			if err := runAllGeneratorCommand(cmd, target, nil); err != nil {
				wrapped := fmt.Errorf("%s: %w", path, err)
				logrus.Error(wrapped)
				runErrors = append(runErrors, wrapped)
				continue
			}
			logrus.Infof("%s completed successfully", path)
		}
		if len(runErrors) > 0 {
			return fmt.Errorf("%d generator command(s) failed: %w", len(runErrors), errors.Join(runErrors...))
		}
		logrus.Info("All generators completed successfully")
		return nil
	},
}

// genConfCmd generates Pigsty configuration files.
var genConfCmd = &cobra.Command{
	Use:   "conf [os...]",
	Short: "Generate Pigsty configuration files",
	Long: `Generate Pigsty configuration YAML files for different OS distributions.
This command generates OS-specific package mappings and extension lists
based on the availability data in the pgext database.`,
	Example: `  # Generate configuration for all supported OS (default)
  pgext gen conf

  # Generate configuration for specific OS
  pgext gen conf el9.x86_64
  pgext gen conf el9.x86_64 d12.aarch64

  # Generate to a specific output directory
  pgext gen conf --output ./output el9.x86_64`,
	PreRunE: initDatabase,
	PostRun: closeDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		configOutputDir, _ := cmd.Flags().GetString("output")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		verbose, _ := cmd.Flags().GetBool("verbose")

		if configOutputDir == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			configOutputDir = filepath.Join(homeDir, "pigsty", "roles", "node_id", "vars")
		}

		osList := args
		if len(osList) == 0 {
			osList = db.SupportedOSPlatforms()
			if !verbose {
				fmt.Printf("Generating configuration for all %d supported OS...\n", len(osList))
			}
		}

		var failedOS []string
		for _, osName := range osList {
			if verbose {
				logrus.Infof("Generating configuration for %s...", osName)
			}
			if err := cli.GeneratePigstyConfig(osName, configOutputDir, dryRun, verbose); err != nil {
				logrus.Errorf("Failed to generate config for %s: %v", osName, err)
				failedOS = append(failedOS, osName)
				continue
			}
			if verbose {
				logrus.Infof("✓ Generated configuration for %s", osName)
			} else {
				fmt.Printf("Generated configuration for %s\n", osName)
			}
		}
		if len(failedOS) > 0 {
			return fmt.Errorf("failed to generate configuration for %d/%d OS targets: %s", len(failedOS), len(osList), strings.Join(failedOS, ", "))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.AddCommand(genPageCmd, genListCmd, genOSCmd, genMatrixCmd, genSiteCmd, genConfCmd, allCmd)
	genCmd.PersistentFlags().StringVarP(&genDir, "dir", "D", "", "Hugo catalog root (PGEXT_SITE_DIR or ~/pgsty/pgext.cloud)")
	for _, catalogCmd := range []*cobra.Command{genPageCmd, genListCmd, genOSCmd, genMatrixCmd} {
		catalogCmd.Flags().StringP("output", "o", "", "Hugo content output directory (default: <dir>/content)")
	}
	genMatrixCmd.Flags().String("static-output", "", "Static asset directory (default: alongside the content output)")
	genPageCmd.Flags().StringVar(&genPageStubDir, "stub-dir", "", "English usage sources (default: stub/ in the pgext source checkout)")

	genConfCmd.Flags().StringP("output", "o", "", "Output directory for configuration files (default: ~/pigsty/roles/node_id/vars)")
	genConfCmd.Flags().BoolP("dry-run", "n", false, "Perform a dry run without writing files")
	genConfCmd.Flags().BoolP("verbose", "v", false, "Show verbose output")
}
