/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

IO command - generates Hugo/Docsy content for pigsty.io (English only)
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"pgext/cli"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ioOutputDir string
	ioStubDir   string
)

// ioCmd represents the io command
var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "Generate Hugo/Docsy content for pigsty.io",
	Long: `Generate Hugo/Docsy-compatible markdown files for pigsty.io English documentation site.

This command generates native Markdown content without Hextra shortcodes,
suitable for the Hugo Docsy theme used by pigsty.io.`,
	Example: `  pgext io page       # Generate extension detail pages
  pgext io list       # Generate all list pages
  pgext io os         # Generate OS-specific availability pages
  pgext io all        # Generate all content`,
}

// ioPageCmd generates individual extension detail pages
var ioPageCmd = &cobra.Command{
	Use:   "page [extension-names...]",
	Short: "Generate extension detail pages for pigsty.io",
	Long: `Generate Hugo/Docsy markdown detail pages for PostgreSQL extensions.
If no extension names are provided, generates pages for all extensions.`,
	Example: `  pgext io page              # Generate pages for all extensions
  pgext io page pgvector     # Generate page for pgvector extension
  pgext io page pgvector postgis  # Generate pages for specific extensions`,
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		generator := cli.NewIOPageGenerator(cache, ioOutputDir, ioStubDir)

		var extensionsToGenerate []*cli.Extension
		if len(args) == 0 {
			extensionsToGenerate = cache.ReadyExtensions()
			logrus.Infof("Generating pages for all %d extensions...", len(extensionsToGenerate))
		} else {
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

		successCount := 0
		var failedExtensions []string
		for i, ext := range extensionsToGenerate {
			if len(extensionsToGenerate) > 10 && (i+1)%100 == 0 {
				logrus.Infof("Progress: %d/%d extensions processed", i+1, len(extensionsToGenerate))
			}
			if err := generator.GenerateExtensionPage(ctx, ext); err != nil {
				logrus.Errorf("Failed to generate page for %s: %v", ext.Name, err)
				failedExtensions = append(failedExtensions, ext.Name)
			} else {
				successCount++
				if len(args) > 0 {
					logrus.Infof("Generated: %s/e/%s.md", ioOutputDir, ext.Name)
				}
			}
		}

		logrus.Infof("Successfully generated %d/%d extension pages", successCount, len(extensionsToGenerate))
		if len(failedExtensions) > 0 {
			logrus.Warnf("Failed to generate pages for: %v", failedExtensions)
		}
		return nil
	}),
}

// ioListCmd generates various list pages
var ioListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Generate list pages for pigsty.io",
	Long:    `Generate list pages including extension list, package list, category list, language list, and license list.`,
	Example: `  pgext io list    # Generate all list pages`,
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		logrus.Info("Generating list pages...")
		if err := cli.NewIOListGenerator(cache, ioOutputDir).GenerateAllLists(); err != nil {
			return fmt.Errorf("failed to generate list pages: %w", err)
		}
		logrus.Info("Successfully generated all list pages")
		return nil
	}),
}

// ioOSCmd generates OS-specific availability pages
var ioOSCmd = &cobra.Command{
	Use:   "os [os-name]",
	Short: "Generate OS-specific availability pages for pigsty.io",
	Long: `Generate markdown pages showing extension availability for specific OS distributions.
If no argument is provided, generates pages for all active OS distributions.`,
	Example: `  pgext io os               # Generate pages for all active OS distributions
  pgext io os el9.x86_64    # Generate page for RHEL 9 x86_64`,
	Args: cobra.MaximumNArgs(1),
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		generator := cli.NewIOOSGenerator(cache, ioOutputDir)
		if len(args) == 1 {
			logrus.Infof("Generating OS page for %s...", args[0])
			if err := generator.GenerateOSPage(ctx, args[0]); err != nil {
				return fmt.Errorf("failed to generate OS page for %s: %w", args[0], err)
			}
			logrus.Infof("Successfully generated OS page for %s", args[0])
		} else {
			if err := generator.GenerateAllOSPages(ctx); err != nil {
				return fmt.Errorf("failed to generate OS pages: %w", err)
			}
		}
		return nil
	}),
}

// ioCateCmd generates per-category PKG index pages
var ioCateCmd = &cobra.Command{
	Use:   "cate [category-name]",
	Short: "Generate per-category PKG index pages for pigsty.io",
	Long: `Generate markdown pages showing PKG-centric index for each extension category.
If no argument is provided, generates pages for all 16 categories.`,
	Example: `  pgext io cate              # Generate pages for all categories
  pgext io cate fts          # Generate page for FTS category`,
	Args: cobra.MaximumNArgs(1),
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		generator := cli.NewIOCateGenerator(cache, ioOutputDir)
		if len(args) == 1 {
			catName := strings.ToUpper(args[0])
			cat, exists := cache.CateMap[catName]
			if !exists {
				return fmt.Errorf("category '%s' not found", args[0])
			}
			logrus.Infof("Generating category page for %s...", catName)
			if err := generator.GenerateCategoryPage(ctx, cat); err != nil {
				return fmt.Errorf("failed to generate category page for %s: %w", catName, err)
			}
			logrus.Infof("Successfully generated category page for %s", catName)
		} else {
			if err := generator.GenerateAllCategoryPages(ctx); err != nil {
				return fmt.Errorf("failed to generate category pages: %w", err)
			}
		}
		return nil
	}),
}

// ioAttrCmd generates extension attribute pages
var ioAttrCmd = &cobra.Command{
	Use:     "attr",
	Short:   "Generate extension attribute pages for pigsty.io",
	Long:    `Generate attribute pages including load, headless, deps, multi, and fork sub-pages.`,
	Example: `  pgext io attr    # Generate all attribute pages`,
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		logrus.Info("Generating attribute pages...")
		if err := cli.NewIOAttrGenerator(cache, ioOutputDir).GenerateAllAttrPages(); err != nil {
			return fmt.Errorf("failed to generate attribute pages: %w", err)
		}
		logrus.Info("Successfully generated all attribute pages")
		return nil
	}),
}

// ioAllCmd generates all content
var ioAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Generate all content for pigsty.io",
	Long:    `Generate all content including extension pages, list pages, and OS availability pages.`,
	Example: `  pgext io all    # Generate all content`,
	RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
		var errors []error
		var wg sync.WaitGroup
		var mu sync.Mutex

		runTask := func(name string, fn func() error) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				logrus.Infof("Generating %s...", name)
				if err := fn(); err != nil {
					mu.Lock()
					errors = append(errors, fmt.Errorf("%s: %w", name, err))
					mu.Unlock()
				} else {
					logrus.Infof("%s generated successfully", name)
				}
			}()
		}

		runTask("list pages", func() error { return cli.NewIOListGenerator(cache, ioOutputDir).GenerateAllLists() })
		runTask("OS pages", func() error { return cli.NewIOOSGenerator(cache, ioOutputDir).GenerateAllOSPages(ctx) })
		runTask("category pages", func() error {
			return cli.NewIOCateGenerator(cache, ioOutputDir).GenerateAllCategoryPages(ctx)
		})
		runTask("attribute pages", func() error { return cli.NewIOAttrGenerator(cache, ioOutputDir).GenerateAllAttrPages() })

		wg.Wait()

		// Generate extension pages sequentially
		logrus.Info("Generating extension pages...")
		pageGenerator := cli.NewIOPageGenerator(cache, ioOutputDir, ioStubDir)
		successCount := 0
		var failedExtensions []string

		for i, ext := range cache.ReadyExtensions() {
			if (i+1)%100 == 0 {
				logrus.Infof("Progress: %d extensions processed", i+1)
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

		if len(errors) > 0 {
			for _, err := range errors {
				logrus.Error(err)
			}
			return fmt.Errorf("generation completed with %d errors", len(errors))
		}

		logrus.Info("All content generated successfully!")
		return nil
	}),
}

func init() {
	rootCmd.AddCommand(ioCmd)
	ioCmd.AddCommand(ioPageCmd, ioListCmd, ioOSCmd, ioCateCmd, ioAttrCmd, ioAllCmd)

	defaultOutput := filepath.Join(os.Getenv("HOME"), "pgsty", "pigsty.io", "content", "ext")
	ioCmd.PersistentFlags().StringVarP(&ioOutputDir, "output", "o", defaultOutput,
		"Output directory for generated files")
	ioCmd.PersistentFlags().StringVar(&ioStubDir, "stub-dir", "stub",
		"Directory containing stub markdown files for extension pages")
}
