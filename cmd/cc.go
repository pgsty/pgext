/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC command - generates Hugo/Docsy content for pigsty.cc (Chinese only)
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
	ccOutputDir string
)

// ccCmd represents the cc command
var ccCmd = &cobra.Command{
	Use:   "cc",
	Short: "Generate Hugo/Docsy content for pigsty.cc",
	Long: `Generate Hugo/Docsy-compatible markdown files for pigsty.cc Chinese documentation site.

This command generates native Markdown content without Hextra shortcodes,
suitable for the Hugo Docsy theme used by pigsty.cc.`,
	Example: `  pgext cc page       # Generate extension detail pages
  pgext cc list       # Generate all list pages
  pgext cc os         # Generate OS-specific availability pages
  pgext cc all        # Generate all content`,
}

// ccPageCmd generates individual extension detail pages
var ccPageCmd = &cobra.Command{
	Use:   "page [extension-names...]",
	Short: "Generate extension detail pages for pigsty.cc",
	Long: `Generate Hugo/Docsy markdown detail pages for PostgreSQL extensions.
If no extension names are provided, generates pages for all extensions.`,
	Example: `  pgext cc page              # Generate pages for all extensions
  pgext cc page pgvector     # Generate page for pgvector extension
  pgext cc page pgvector postgis  # Generate pages for specific extensions`,
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
		extDir := filepath.Join(ccOutputDir, "e")
		if err := os.MkdirAll(extDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Initialize generator
		generator := cli.NewCCPageGenerator(cache, ccOutputDir)

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
					logrus.Infof("Generated: %s/e/%s/_index.md", ccOutputDir, ext.Name)
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

// ccListCmd generates various list pages
var ccListCmd = &cobra.Command{
	Use:   "list",
	Short: "Generate list pages for pigsty.cc",
	Long:  `Generate list pages including extension list, package list, category list, language list, and license list.`,
	Example: `  pgext cc list    # Generate all list pages`,
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

		// Initialize generator
		generator := cli.NewCCListGenerator(cache, ccOutputDir)

		// Generate all list pages
		logrus.Info("Generating list pages...")
		if err := generator.GenerateAllLists(); err != nil {
			return fmt.Errorf("failed to generate list pages: %w", err)
		}

		logrus.Info("Successfully generated all list pages")
		return nil
	},
}

// ccOSCmd generates OS-specific availability pages
var ccOSCmd = &cobra.Command{
	Use:   "os [os-name]",
	Short: "Generate OS-specific availability pages for pigsty.cc",
	Long: `Generate markdown pages showing extension availability for specific OS distributions.
If no argument is provided, generates pages for all active OS distributions.`,
	Example: `  pgext cc os               # Generate pages for all active OS distributions
  pgext cc os el9.x86_64    # Generate page for RHEL 9 x86_64`,
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

		// Initialize generator
		generator := cli.NewCCOSGenerator(cache, ccOutputDir)

		if len(args) == 1 {
			// Generate for specific OS
			osName := args[0]
			logrus.Infof("Generating OS page for %s...", osName)
			if err := generator.GenerateOSPage(ctx, osName); err != nil {
				return fmt.Errorf("failed to generate OS page for %s: %w", osName, err)
			}
			logrus.Infof("Successfully generated OS page for %s", osName)
		} else {
			// Generate for all active OS
			if err := generator.GenerateAllOSPages(ctx); err != nil {
				return fmt.Errorf("failed to generate OS pages: %w", err)
			}
		}

		return nil
	},
}

// ccAllCmd generates all content
var ccAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all content for pigsty.cc",
	Long:  `Generate all content including extension pages, list pages, and OS availability pages.`,
	Example: `  pgext cc all    # Generate all content`,
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

		// Track results
		var errors []error
		var wg sync.WaitGroup
		var mu sync.Mutex

		// 1. Generate list pages
		wg.Add(1)
		go func() {
			defer wg.Done()
			logrus.Info("Generating list pages...")
			listGenerator := cli.NewCCListGenerator(cache, ccOutputDir)
			if err := listGenerator.GenerateAllLists(); err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("list pages: %w", err))
				mu.Unlock()
			} else {
				logrus.Info("List pages generated successfully")
			}
		}()

		// 2. Generate OS pages
		wg.Add(1)
		go func() {
			defer wg.Done()
			logrus.Info("Generating OS pages...")
			osGenerator := cli.NewCCOSGenerator(cache, ccOutputDir)
			if err := osGenerator.GenerateAllOSPages(ctx); err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("OS pages: %w", err))
				mu.Unlock()
			} else {
				logrus.Info("OS pages generated successfully")
			}
		}()

		// Wait for list and OS pages to complete first
		wg.Wait()

		// 3. Generate extension pages (sequential to avoid overwhelming the system)
		logrus.Info("Generating extension pages...")
		pageGenerator := cli.NewCCPageGenerator(cache, ccOutputDir)

		successCount := 0
		var failedExtensions []string

		for i, ext := range cache.Extensions {
			if ext.State.Valid && ext.State.String == "not-ready" {
				continue
			}

			// Progress logging
			if (i+1)%100 == 0 {
				logrus.Infof("Progress: %d/%d extensions processed", i+1, len(cache.Extensions))
			}

			if err := pageGenerator.GenerateExtensionPage(ctx, ext); err != nil {
				logrus.Errorf("Failed to generate page for %s: %v", ext.Name, err)
				failedExtensions = append(failedExtensions, ext.Name)
			} else {
				successCount++
			}
		}

		logrus.Infof("Extension pages: %d generated", successCount)
		if len(failedExtensions) > 0 {
			logrus.Warnf("Failed extensions: %v", failedExtensions)
		}

		// Report any errors
		if len(errors) > 0 {
			logrus.Errorf("Encountered %d errors during generation:", len(errors))
			for _, err := range errors {
				logrus.Error(err)
			}
			return fmt.Errorf("generation completed with %d errors", len(errors))
		}

		logrus.Info("All content generated successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ccCmd)

	// Add subcommands
	ccCmd.AddCommand(ccPageCmd, ccListCmd, ccOSCmd, ccAllCmd)

	// Default output directory is ~/pgsty/pigsty.cc/content/ext
	defaultOutput := filepath.Join(os.Getenv("HOME"), "pgsty", "pigsty.cc", "content", "ext")
	ccCmd.PersistentFlags().StringVarP(&ccOutputDir, "output", "o", defaultOutput,
		"Output directory for generated files")
}
