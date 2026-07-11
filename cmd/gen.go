/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pgext/cli"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const defaultCatalogOutputDir = "content"

var (
	genListOutputDir   = defaultCatalogOutputDir
	genOSOutputDir     = defaultCatalogOutputDir
	genMatrixOutputDir = defaultCatalogOutputDir
	genMatrixStaticDir = "static"
	genPageOutputDir   = defaultCatalogOutputDir
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate extension catalogs, documentation, matrices, and configuration",
	Long: `Generate PostgreSQL extension artifacts from the pgext database.

Use cc and io for the Chinese and English Pigsty documentation sites. The
page, list, os, and matrix commands generate the PGEXT catalog site, while
conf generates Pigsty package configuration. Run pgext gen or pgext gen all
to execute every generator.`,
	Example: `  pgext gen cc          # Generate all pigsty.cc extension content
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

// genListCmd generates various list pages
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
	Args:      cobra.OnlyValidArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// Initialize database connection
		if err := cli.InitDB(dbURL); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer cli.CloseDB()

		// Load extension cache
		logrus.Info("Loading extension data...")
		cache, err := cli.LoadExtensionCache(ctx)
		if err != nil {
			return fmt.Errorf("failed to load extension cache: %w", err)
		}

		// Prepare output directory
		listDir := filepath.Join(genListOutputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Prepare e directory for ext index pages
		extDir := filepath.Join(genListOutputDir, "e")
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create extension directory: %w", err)
		}

		// Determine which lists to generate
		typesToGenerate := make(map[string]bool)
		if len(args) == 0 {
			// No arguments, generate all types
			typesToGenerate["ext"] = true
			typesToGenerate["pkg"] = true
			typesToGenerate["cate"] = true
			typesToGenerate["lang"] = true
			typesToGenerate["license"] = true
			typesToGenerate["catalog"] = true
			logrus.Info("Generating all list types...")
		} else {
			// Generate specified types
			for _, arg := range args {
				typesToGenerate[strings.ToLower(arg)] = true
			}
			logrus.Infof("Generating list types: %v", args)
		}

		// Initialize generators
		listGenerator := cli.NewListGenerator(cache, listDir)
		extGenerator := &cli.ExtensionGenerator{
			DB:    cli.DB,
			Cache: cache,
		}

		// Track results
		var errors []error
		successCount := 0

		// Generate ext list pages
		if typesToGenerate["ext"] {
			logrus.Info("Generating extension list pages...")

			// Generate extension list pages (in list directory)
			enListPath := filepath.Join(listDir, "ext.md")
			if err := extGenerator.GenerateExtensionList("en", enListPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English extension list: %w", err))
			} else {
				logrus.Infof("Generated: %s", enListPath)
				successCount++
			}

			zhListPath := filepath.Join(listDir, "ext.zh.md")
			if err := extGenerator.GenerateExtensionList("zh", zhListPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese extension list: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhListPath)
				successCount++
			}

			// Generate extension index pages (in e directory)
			enIndexPath := filepath.Join(extDir, "_index.md")
			zhIndexPath := filepath.Join(extDir, "_index.zh.md")
			if err := extGenerator.GenerateExtensionIndexPages(enIndexPath, zhIndexPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate extension index pages: %w", err))
			} else {
				logrus.Infof("Generated: %s", enIndexPath)
				logrus.Infof("Generated: %s", zhIndexPath)
				successCount += 2
			}
		}

		// Generate pkg list pages
		if typesToGenerate["pkg"] {
			logrus.Info("Generating package list pages...")

			enPath := filepath.Join(listDir, "pkg.md")
			if err := extGenerator.GeneratePackageList("en", enPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English package list: %w", err))
			} else {
				logrus.Infof("Generated: %s", enPath)
				successCount++
			}

			zhPath := filepath.Join(listDir, "pkg.zh.md")
			if err := extGenerator.GeneratePackageList("zh", zhPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese package list: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhPath)
				successCount++
			}
		}

		// Generate cate list pages
		if typesToGenerate["cate"] {
			logrus.Info("Generating category list pages...")

			enPath := filepath.Join(listDir, "cate.md")
			if err := listGenerator.GenerateCategoryList("en", enPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English category list: %w", err))
			} else {
				logrus.Infof("Generated: %s", enPath)
				successCount++
			}

			zhPath := filepath.Join(listDir, "cate.zh.md")
			if err := listGenerator.GenerateCategoryList("zh", zhPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese category list: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhPath)
				successCount++
			}
		}

		// Generate lang list pages
		if typesToGenerate["lang"] {
			logrus.Info("Generating language list pages...")

			enPath := filepath.Join(listDir, "lang.md")
			if err := listGenerator.GenerateLanguageList("en", enPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English language list: %w", err))
			} else {
				logrus.Infof("Generated: %s", enPath)
				successCount++
			}

			zhPath := filepath.Join(listDir, "lang.zh.md")
			if err := listGenerator.GenerateLanguageList("zh", zhPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese language list: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhPath)
				successCount++
			}
		}

		// Generate license list pages
		if typesToGenerate["license"] {
			logrus.Info("Generating license list pages...")

			enPath := filepath.Join(listDir, "license.md")
			if err := listGenerator.GenerateLicenseList("en", enPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English license list: %w", err))
			} else {
				logrus.Infof("Generated: %s", enPath)
				successCount++
			}

			zhPath := filepath.Join(listDir, "license.zh.md")
			if err := listGenerator.GenerateLicenseList("zh", zhPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese license list: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhPath)
				successCount++
			}
		}

		// Generate catalog list pages
		if typesToGenerate["catalog"] {
			logrus.Info("Generating catalog list pages...")

			enPath := filepath.Join(listDir, "_index.md")
			if err := extGenerator.GenerateCatalogPage("en", enPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate English catalog: %w", err))
			} else {
				logrus.Infof("Generated: %s", enPath)
				successCount++
			}

			zhPath := filepath.Join(listDir, "_index.zh.md")
			if err := extGenerator.GenerateCatalogPage("zh", zhPath); err != nil {
				errors = append(errors, fmt.Errorf("failed to generate Chinese catalog: %w", err))
			} else {
				logrus.Infof("Generated: %s", zhPath)
				successCount++
			}
		}

		// Report results
		logrus.Infof("List generation completed: %d files generated", successCount)

		if len(errors) > 0 {
			logrus.Errorf("Encountered %d errors during generation:", len(errors))
			for _, err := range errors {
				logrus.Error(err)
			}
			return fmt.Errorf("list generation completed with %d errors", len(errors))
		}

		return nil
	},
}

// genOSCmd generates OS-specific availability pages
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
		ctx := cmd.Context()

		// Initialize database connection
		if err := cli.InitDB(dbURL); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer cli.CloseDB()

		// Load extension cache
		logrus.Info("Loading extension data...")
		cache, err := cli.LoadExtensionCache(ctx)
		if err != nil {
			return fmt.Errorf("failed to load extension cache: %w", err)
		}

		// Prepare output directory
		listDir := filepath.Join(genOSOutputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Prepare os directory
		osDir := filepath.Join(genOSOutputDir, "os")
		if err := os.MkdirAll(osDir, 0755); err != nil {
			return fmt.Errorf("failed to create os output directory: %w", err)
		}
		generator := cli.NewOSGenerator(cache, osDir)

		// Determine which OS pages to generate
		var osList []string
		if len(args) == 0 {
			// No argument provided, get all active OS
			logrus.Info("Getting all active OS distributions...")
			osList, err = cli.GetActiveOS(ctx)
			if err != nil {
				return fmt.Errorf("failed to get active OS list: %w", err)
			}
			if len(osList) == 0 {
				return fmt.Errorf("no active OS distributions found")
			}
			logrus.Infof("Found %d active OS distributions to generate", len(osList))
		} else {
			// Single OS specified
			osList = []string{args[0]}
		}

		// Generate pages (in parallel for multiple OS)
		if len(osList) == 1 {
			// Single OS, generate directly
			osName := osList[0]
			logrus.Infof("Generating OS page for %s...", osName)
			if err := generator.GenerateOSPage(ctx, osName); err != nil {
				return fmt.Errorf("failed to generate OS page for %s: %w", osName, err)
			}
			logrus.Infof("Successfully generated OS page: %s.md", osName)
			return nil
		}

		// Multiple OS, use parallel generation
		type result struct {
			os      string
			success bool
			err     error
		}

		// Create channels for coordination
		results := make(chan result, len(osList))
		// Use 8 parallel workers by default for generation
		const maxParallel = 8
		semaphore := make(chan struct{}, maxParallel)
		var wg sync.WaitGroup

		// Launch goroutines to generate pages
		for _, osName := range osList {
			wg.Add(1)
			go func(os string) {
				defer wg.Done()
				semaphore <- struct{}{}        // Acquire semaphore
				defer func() { <-semaphore }() // Release semaphore

				logrus.Infof("Generating OS page for %s...", os)
				err := generator.GenerateOSPage(ctx, os)
				if err != nil {
					logrus.Errorf("Failed to generate OS page for %s: %v", os, err)
					results <- result{os: os, success: false, err: err}
				} else {
					logrus.Infof("Successfully generated OS page: %s.md", os)
					results <- result{os: os, success: true, err: nil}
				}
			}(osName)
		}

		// Wait for all goroutines to complete
		go func() {
			wg.Wait()
			close(results)
		}()

		// Collect results
		successCount := 0
		failedOS := []string{}
		for res := range results {
			if res.success {
				successCount++
			} else {
				failedOS = append(failedOS, res.os)
			}
		}

		logrus.Infof("Successfully generated %d/%d OS pages", successCount, len(osList))

		if len(failedOS) > 0 {
			logrus.Warnf("Failed to generate pages for: %v", failedOS)
			return fmt.Errorf("failed to generate %d/%d OS pages", len(failedOS), len(osList))
		}
		return nil
	},
}

// genMatrixCmd generates the global OS x PG package availability matrix.
var genMatrixCmd = &cobra.Command{
	Use:   "matrix",
	Short: "Generate global OS/PG availability matrix",
	Long: `Generate a package-oriented global availability matrix across all active
operating systems and PostgreSQL major versions. Hugo pages are written below
--output, while CSV/JSON/HTML assets are written below --static-output.`,
	Example: `  pgext gen matrix
  pgext gen matrix -d data
  pgext gen matrix -o /tmp/content --static-output /tmp/static`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		if err := cli.InitDB(dbURL); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer cli.CloseDB()

		logrus.Info("Loading extension data...")
		cache, err := cli.LoadExtensionCache(ctx)
		if err != nil {
			return fmt.Errorf("failed to load extension cache: %w", err)
		}

		generator := cli.NewGlobalMatrixGenerator(cache, genMatrixOutputDir, genMatrixStaticDir)
		if err := generator.Generate(ctx); err != nil {
			return fmt.Errorf("failed to generate global matrix: %w", err)
		}

		logrus.Infof("Generated global matrix under %s and %s", filepath.Join(genMatrixOutputDir, "os"), filepath.Join(genMatrixStaticDir, "matrix"))
		return nil
	},
}

// genPageCmd generates individual extension detail pages
var genPageCmd = &cobra.Command{
	Use:   "page [extension-names...]",
	Short: "Generate individual extension detail pages",
	Long: `Generate Hugo markdown detail pages for PostgreSQL extensions.
If no extension names are provided, generates pages for all extensions.`,
	Example: `  pgext gen page              # Generate pages for all extensions
  pgext gen page vector       # Generate page for pgvector's vector extension
  pgext gen page vector postgis pg_stat_statements  # Generate pages for specific extensions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// Initialize database connection
		if err := cli.InitDB(dbURL); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer cli.CloseDB()

		// Load extension cache
		logrus.Info("Loading extension data...")
		cache, err := cli.LoadExtensionCache(ctx)
		if err != nil {
			return fmt.Errorf("failed to load extension cache: %w", err)
		}

		// Prepare output directory
		extDir := filepath.Join(genPageOutputDir, "e")
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Initialize generator
		generator := cli.NewExtensionGenerator(cache, extDir)

		// Determine which extensions to generate
		var extensionsToGenerate []*cli.Extension
		if len(args) == 0 {
			// No arguments provided, generate all extensions (except not-ready)
			for _, ext := range cache.Extensions {
				if ext.IsReady() {
					extensionsToGenerate = append(extensionsToGenerate, ext)
				}
			}
			logrus.Infof("Generating pages for all %d extensions...", len(extensionsToGenerate))
		} else {
			// Specific extensions requested
			for _, name := range args {
				ext, exists := cache.ExtMap[name]
				if !exists {
					logrus.Warnf("Extension '%s' not found in cache, skipping", name)
					continue
				}
				extensionsToGenerate = append(extensionsToGenerate, ext)
			}
			if len(extensionsToGenerate) == 0 {
				return fmt.Errorf("no valid extensions found from provided names")
			}
			logrus.Infof("Generating pages for %d specified extensions...", len(extensionsToGenerate))
		}

		// Generate pages with progress tracking
		successCount := 0
		var failedExtensions []string

		for i, ext := range extensionsToGenerate {
			// Progress logging for batch generation
			if len(extensionsToGenerate) > 10 && (i+1)%100 == 0 {
				logrus.Infof("Progress: %d/%d extensions processed", i+1, len(extensionsToGenerate))
			}

			if err := generator.GenerateExtensionPage(ctx, ext); err != nil {
				logrus.Errorf("Failed to generate page for %s: %v", ext.Name, err)
				failedExtensions = append(failedExtensions, ext.Name)
			} else {
				successCount++
				if len(args) > 0 {
					// Only log individual successes when specific extensions are requested
					logrus.Infof("Generated: %s/%s.md", extDir, ext.Name)
				}
			}
		}

		// Final summary
		logrus.Infof("Successfully generated %d/%d extension pages", successCount, len(extensionsToGenerate))
		if len(failedExtensions) > 0 {
			logrus.Warnf("Failed to generate pages for: %v", failedExtensions)
			return fmt.Errorf("failed to generate %d extension pages: %s", len(failedExtensions), strings.Join(failedExtensions, ", "))
		}

		return nil
	},
}

// allGeneratorCommands defines the complete generation pipeline in execution order.
var allGeneratorCommands = []*cobra.Command{
	ccCmd,
	ioCmd,
	genPageCmd,
	genListCmd,
	genOSCmd,
	genMatrixCmd,
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
  pgext gen page
  pgext gen list
  pgext gen os
  pgext gen matrix
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

// genConfCmd generates Pigsty configuration files
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
		// Get flags
		configOutputDir, _ := cmd.Flags().GetString("output")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Set default output directory if not specified
		if configOutputDir == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			configOutputDir = filepath.Join(homeDir, "pigsty", "roles", "node_id", "vars")
		}

		// Determine which OS to generate
		var osList []string
		if len(args) > 0 {
			// Use provided OS arguments
			osList = args
		} else {
			// Default: generate for all supported OS
			osList = []string{
				"el8.x86_64", "el8.aarch64",
				"el9.x86_64", "el9.aarch64",
				"el10.x86_64", "el10.aarch64",
				"d12.x86_64", "d12.aarch64",
				"d13.x86_64", "d13.aarch64",
				"u22.x86_64", "u22.aarch64",
				"u24.x86_64", "u24.aarch64",
				"u26.x86_64", "u26.aarch64",
			}
			if !verbose {
				fmt.Printf("Generating configuration for all %d supported OS...\n", len(osList))
			}
		}

		// Generate configuration for each OS
		var failedOS []string
		for _, osName := range osList {
			if verbose {
				logrus.Infof("Generating configuration for %s...", osName)
			}

			err := cli.GeneratePigstyConfig(osName, configOutputDir, dryRun, verbose)
			if err != nil {
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

	genCmd.AddCommand(genPageCmd, genListCmd, genOSCmd, genMatrixCmd, genConfCmd, allCmd)

	genPageCmd.Flags().StringVarP(&genPageOutputDir, "output", "o", defaultCatalogOutputDir,
		"Output directory for generated files")
	genListCmd.Flags().StringVarP(&genListOutputDir, "output", "o", defaultCatalogOutputDir,
		"Output directory for generated files")
	genOSCmd.Flags().StringVarP(&genOSOutputDir, "output", "o", defaultCatalogOutputDir,
		"Output directory for generated files")
	genMatrixCmd.Flags().StringVarP(&genMatrixOutputDir, "output", "o", defaultCatalogOutputDir,
		"Hugo content output directory")
	genMatrixCmd.Flags().StringVar(&genMatrixStaticDir, "static-output", "static",
		"Static matrix asset output directory")

	// Add flags specific to gen conf command
	genConfCmd.Flags().StringP("output", "o", "", "Output directory for configuration files (default: ~/pigsty/roles/node_id/vars)")
	genConfCmd.Flags().BoolP("dry-run", "n", false, "Perform a dry run without writing files")
	genConfCmd.Flags().BoolP("verbose", "v", false, "Show verbose output")
}
