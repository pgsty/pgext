/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"pgext/cli"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	outputDir string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Hugo markdown content from database",
	Long:  `Generate Hugo/Hextra-compatible markdown files for PostgreSQL extensions`,
	Example: `  pgext gen page       # Generate extension detail pages
  pgext gen ext        # Generate extension list pages with stats
  pgext gen pkg        # Generate package list pages with stats
  pgext gen index      # Generate extension index pages
  pgext gen cate       # Generate category list pages
  pgext gen lang       # Generate language list pages
  pgext gen license    # Generate license list pages
  pgext gen catalog    # Generate catalog pages
  pgext gen os         # Generate OS-specific availability page
  pgext gen conf       # Generate Pigsty configuration files`,
}

// genIndexCmd generates extension index pages
var genIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "Generate extension index pages",
	Long:  `Generate _index.md and _index.zh.md for extension listing`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		extDir := filepath.Join(outputDir, "e")
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Generate index pages
		generator := cli.NewExtensionGenerator(cache, extDir)

		// Generate English index
		enPath := filepath.Join(extDir, "_index.md")
		if err := generator.GenerateIndexPage("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English index: %w", err)
		}
		logrus.Infof("Generated: %s", enPath)

		// Generate Chinese index
		zhPath := filepath.Join(extDir, "_index.zh.md")
		if err := generator.GenerateIndexPage("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese index: %w", err)
		}
		logrus.Infof("Generated: %s", zhPath)

		logrus.Infof("Extensions: %d, Packages: %d", len(cache.Extensions), len(cache.PkgMap))
		return nil
	},
}

// genCateCmd generates category list pages
var genCateCmd = &cobra.Command{
	Use:   "cate",
	Short: "Generate category list pages",
	Long:  `Generate cate.md and cate.zh.md for category-based listing`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Generate category list pages
		generator := cli.NewListGenerator(cache, listDir)

		// Generate English version
		enPath := filepath.Join(listDir, "cate.md")
		if err := generator.GenerateCategoryList("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English category list: %w", err)
		}
		logrus.Infof("Generated: %s", enPath)

		// Generate Chinese version
		zhPath := filepath.Join(listDir, "cate.zh.md")
		if err := generator.GenerateCategoryList("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese category list: %w", err)
		}
		logrus.Infof("Generated: %s", zhPath)

		// Count unique packages
		pkgCount := len(cache.PkgMap)
		logrus.Infof("Categories: %d, Extensions: %d, Packages: %d",
			len(cache.Categories), len(cache.Extensions), pkgCount)
		return nil
	},
}

// genLangCmd generates language list pages
var genLangCmd = &cobra.Command{
	Use:   "lang",
	Short: "Generate language list pages",
	Long:  `Generate lang.md and lang.zh.md for language-based listing`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Generate language list pages
		generator := cli.NewListGenerator(cache, listDir)

		// Generate English version
		enPath := filepath.Join(listDir, "lang.md")
		if err := generator.GenerateLanguageList("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English language list: %w", err)
		}
		logrus.Infof("Generated: %s", enPath)

		// Generate Chinese version
		zhPath := filepath.Join(listDir, "lang.zh.md")
		if err := generator.GenerateLanguageList("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese language list: %w", err)
		}
		logrus.Infof("Generated: %s", zhPath)

		return nil
	},
}

// genLicenseCmd generates license list pages
var genLicenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Generate license list pages",
	Long:  `Generate license.md and license.zh.md for license-based listing`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Generate license list pages
		generator := cli.NewListGenerator(cache, listDir)

		// Generate English version
		enPath := filepath.Join(listDir, "license.md")
		if err := generator.GenerateLicenseList("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English license list: %w", err)
		}
		logrus.Infof("Generated: %s", enPath)

		// Generate Chinese version
		zhPath := filepath.Join(listDir, "license.zh.md")
		if err := generator.GenerateLicenseList("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese license list: %w", err)
		}
		logrus.Infof("Generated: %s", zhPath)

		return nil
	},
}

var genCatalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Generate catalog list pages",
	Long:  `Generate catalog list pages (_index.md and _index.zh.md) with statistics and categories`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Initialize generator with database from cli package
		generator := &cli.ExtensionGenerator{
			DB:    cli.DB,
			Cache: cache,
		}

		// Generate English version
		enPath := filepath.Join(listDir, "_index.md")
		if err := generator.GenerateCatalogPage("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English catalog: %w", err)
		}
		logrus.Info("Generated English catalog", "path", enPath)

		// Generate Chinese version
		zhPath := filepath.Join(listDir, "_index.zh.md")
		if err := generator.GenerateCatalogPage("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese catalog: %w", err)
		}
		logrus.Info("Generated Chinese catalog", "path", zhPath)

		logrus.Info("Catalog generation completed", "output", outputDir)
		return nil
	},
}

// genExtCmd generates extension list pages
var genExtCmd = &cobra.Command{
	Use:   "ext",
	Short: "Generate extension list pages and index",
	Long: `Generate ext.md, ext.zh.md for comprehensive extension listing with statistics, categories, and full index.
Also generates _index.md and _index.zh.md for extension index pages.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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

		// Prepare output directories
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create list directory: %w", err)
		}

		extDir := filepath.Join(outputDir, "e")
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create extension directory: %w", err)
		}

		// Initialize generator with database from cli package
		generator := &cli.ExtensionGenerator{
			DB:    cli.DB,
			Cache: cache,
		}

		// Generate extension list pages (in list directory)
		enListPath := filepath.Join(listDir, "ext.md")
		if err := generator.GenerateExtensionList("en", enListPath); err != nil {
			return fmt.Errorf("failed to generate English extension list: %w", err)
		}
		logrus.Infof("Generated: %s", enListPath)

		zhListPath := filepath.Join(listDir, "ext.zh.md")
		if err := generator.GenerateExtensionList("zh", zhListPath); err != nil {
			return fmt.Errorf("failed to generate Chinese extension list: %w", err)
		}
		logrus.Infof("Generated: %s", zhListPath)

		// Generate extension index pages (in e directory)
		enIndexPath := filepath.Join(extDir, "_index.md")
		zhIndexPath := filepath.Join(extDir, "_index.zh.md")
		if err := generator.GenerateExtensionIndexPages(enIndexPath, zhIndexPath); err != nil {
			return fmt.Errorf("failed to generate extension index pages: %w", err)
		}
		logrus.Infof("Generated: %s", enIndexPath)
		logrus.Infof("Generated: %s", zhIndexPath)

		logrus.Infof("Extension list and index generation completed")
		return nil
	},
}

// genPackageCmd generates package list pages
var genPkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Generate package list pages",
	Long:  `Generate pkg.md and pkg.zh.md for comprehensive package listing with statistics, categories, and full index`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Initialize generator with database from cli package
		generator := &cli.ExtensionGenerator{
			DB:    cli.DB,
			Cache: cache,
		}

		// Generate English version
		enPath := filepath.Join(listDir, "pkg.md")
		if err := generator.GeneratePackageList("en", enPath); err != nil {
			return fmt.Errorf("failed to generate English package list: %w", err)
		}
		logrus.Infof("Generated: %s", enPath)

		// Generate Chinese version
		zhPath := filepath.Join(listDir, "pkg.zh.md")
		if err := generator.GeneratePackageList("zh", zhPath); err != nil {
			return fmt.Errorf("failed to generate Chinese package list: %w", err)
		}
		logrus.Infof("Generated: %s", zhPath)

		logrus.Infof("Package list generation completed")
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
		ctx := context.Background()

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
		listDir := filepath.Join(outputDir, "list")
		if err := os.MkdirAll(listDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Prepare os directory
		osDir := filepath.Join(outputDir, "os")
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
		}

		if successCount == 0 {
			return fmt.Errorf("failed to generate any OS pages")
		}
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
  pgext gen page pgvector     # Generate page for pgvector extension
  pgext gen page pgvector postgis pg_stat_statements  # Generate pages for specific extensions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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
		extDir := filepath.Join(outputDir, "e")
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
				if !ext.State.Valid || ext.State.String != "not-ready" {
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
		}

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
	PreRunE:  initDatabase,
	PostRun:  closeDatabase,
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
			// Default: generate for all 14 supported OS
			osList = []string{
				"el8.x86_64", "el8.aarch64",
				"el9.x86_64", "el9.aarch64",
				"el10.x86_64", "el10.aarch64",
				"d12.x86_64", "d12.aarch64",
				"d13.x86_64", "d13.aarch64",
				"u22.x86_64", "u22.aarch64",
				"u24.x86_64", "u24.aarch64",
			}
			if !verbose {
				fmt.Printf("Generating configuration for all %d supported OS...\n", len(osList))
			}
		}

		// Generate configuration for each OS
		for _, osName := range osList {
			if verbose {
				logrus.Infof("Generating configuration for %s...", osName)
			}

			err := cli.GeneratePigstyConfig(osName, configOutputDir, dryRun, verbose)
			if err != nil {
				logrus.Errorf("Failed to generate config for %s: %v", osName, err)
				// Continue with other OS even if one fails
				continue
			}

			if verbose {
				logrus.Infof("✓ Generated configuration for %s", osName)
			} else {
				fmt.Printf("Generated configuration for %s\n", osName)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.AddCommand(genExtCmd, genPkgCmd, genPageCmd, genIndexCmd, genCateCmd, genLangCmd, genLicenseCmd, genCatalogCmd, genOSCmd, genConfCmd)

	genCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "content",
		"Output directory for generated files")

	// Add flags specific to gen conf command
	genConfCmd.Flags().StringP("output", "o", "", "Output directory for configuration files (default: ~/pigsty/roles/node_id/vars)")
	genConfCmd.Flags().BoolP("dry-run", "n", false, "Perform a dry run without writing files")
	genConfCmd.Flags().BoolP("verbose", "v", false, "Show verbose output")
}
