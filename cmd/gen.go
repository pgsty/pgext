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
	Example: `  pgext gen ext      # Generate extension detail pages
  pgext gen index    # Generate extension index pages
  pgext gen cate     # Generate category list pages
  pgext gen lang     # Generate language list pages
  pgext gen license  # Generate license list pages
  pgext gen os el9.x86_64  # Generate OS-specific availability page`,
}

// genExtCmd generates extension detail pages
var genExtCmd = &cobra.Command{
	Use:   "ext",
	Short: "Generate extension detail pages",
	Long:  `Generate individual markdown pages for each extension in content/e/`,
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

		// Generate extension pages
		generator := cli.NewExtensionGenerator(cache, extDir)
		successCount, totalCount := 0, len(cache.Extensions)

		logrus.Infof("Generating %d extension pages...", totalCount)
		for _, ext := range cache.Extensions {
			if err := generator.GenerateExtensionPage(ctx, ext); err != nil {
				logrus.Errorf("Failed to generate page for %s: %v", ext.Name, err)
			} else {
				successCount++
				logrus.Debugf("Generated: %s.md", ext.Name)
			}
		}

		logrus.Infof("Successfully generated %d/%d extension pages", successCount, totalCount)
		return nil
	},
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

// genOSCmd generates OS-specific availability pages
var genOSCmd = &cobra.Command{
	Use:   "os <os-name>",
	Short: "Generate OS-specific availability page",
	Long:  `Generate a markdown page showing extension availability for a specific OS distribution`,
	Example: `  pgext gen os el9.x86_64    # Generate page for RHEL 9 x86_64
  pgext gen os u24.x86_64    # Generate page for Ubuntu 24.04 x86_64
  pgext gen os el10.aarch64  # Generate page for RHEL 10 ARM64`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		osName := args[0]

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

		// Generate OS page
		generator := cli.NewOSGenerator(cache, listDir)
		if err := generator.GenerateOSPage(ctx, osName); err != nil {
			return fmt.Errorf("failed to generate OS page for %s: %w", osName, err)
		}

		logrus.Infof("Successfully generated OS page: %s.md", osName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.AddCommand(genExtCmd, genIndexCmd, genCateCmd, genLangCmd, genLicenseCmd, genCatalogCmd, genOSCmd)

	genCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "content",
		"Output directory for generated files")
}
