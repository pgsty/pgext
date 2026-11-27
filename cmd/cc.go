/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

`gen cc` and `gen io` commands - unified OINK documentation site generator for
pigsty.cc (Chinese) and pigsty.io (English).
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

type sharedCacheKey struct{}

// withSharedCache attaches an already-loaded ExtensionCache to ctx so compound
// commands like `gen site` reuse a single DB pool and cache load.
func withSharedCache(ctx context.Context, cache *cli.ExtensionCache) context.Context {
	return context.WithValue(ctx, sharedCacheKey{}, cache)
}

// runWithCache wraps common boilerplate: InitDB, LoadExtensionCache, then call fn.
// When an outer command already populated sharedCacheKey{}, it reuses that pool and cache.
func runWithCache(fn func(ctx context.Context, cache *cli.ExtensionCache, args []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if cache, ok := ctx.Value(sharedCacheKey{}).(*cli.ExtensionCache); ok && cache != nil {
			return fn(ctx, cache, args)
		}
		if err := cli.InitDB(dbURL); err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer cli.CloseDB()
		logrus.Info("Loading extension data...")
		cache, err := cli.LoadExtensionCache(ctx)
		if err != nil {
			return fmt.Errorf("failed to load extension cache: %w", err)
		}
		return fn(ctx, cache, args)
	}
}

// selectExtensions resolves explicit CLI extension names or defaults to all ready extensions.
func selectExtensions(cache *cli.ExtensionCache, args []string) ([]*cli.Extension, error) {
	if len(args) == 0 {
		exts := cache.ReadyExtensions()
		logrus.Infof("Generating pages for all %d extensions...", len(exts))
		return exts, nil
	}
	var exts []*cli.Extension
	for _, name := range args {
		ext, exists := cache.ExtMap[name]
		if !exists {
			logrus.Warnf("Extension '%s' not found in cache, skipping", name)
			continue
		}
		exts = append(exts, ext)
	}
	if len(exts) == 0 {
		return nil, fmt.Errorf("no valid extensions found from provided names")
	}
	logrus.Infof("Generating pages for %d specified extensions...", len(exts))
	return exts, nil
}

// generateExtensionPages validates stubs and runs genFn across exts with progress tracking.
func generateExtensionPages(ctx context.Context, stubDir string, exts []*cli.Extension, logEach bool, outDir string, genFn func(context.Context, *cli.Extension) error) error {
	if err := cli.ValidateExtensionStubs(stubDir, exts); err != nil {
		return err
	}
	successCount := 0
	var failed []string
	for i, ext := range exts {
		if len(exts) > 10 && (i+1)%100 == 0 {
			logrus.Infof("Progress: %d/%d extensions processed", i+1, len(exts))
		}
		if err := genFn(ctx, ext); err != nil {
			logrus.Errorf("Failed to generate page for %s: %v", ext.Name, err)
			failed = append(failed, ext.Name)
		} else {
			successCount++
			if logEach {
				logrus.Infof("Generated: %s/%s.md", outDir, ext.Name)
			}
		}
	}
	logrus.Infof("Successfully generated %d/%d extension pages", successCount, len(exts))
	if len(failed) > 0 {
		logrus.Warnf("Failed to generate pages for: %v", failed)
		return fmt.Errorf("failed to generate %d extension pages: %s", len(failed), strings.Join(failed, ", "))
	}
	return nil
}

type docsSiteSpec struct {
	use        string
	site       string
	langLabel  string
	locale     cli.DocsLocale
	stubSubdir string
	repoDir    string
}

func newDocsSiteCommand(spec docsSiteSpec) (parent, all *cobra.Command) {
	var outputDir, stubDir string

	newPageGen := func(cache *cli.ExtensionCache, resolvedStub string) *cli.CCPageGenerator {
		if spec.locale == cli.DocsEN {
			return cli.NewIOPageGenerator(cache, outputDir, resolvedStub)
		}
		return cli.NewCCPageGenerator(cache, outputDir, resolvedStub)
	}
	newListGen := func(cache *cli.ExtensionCache) *cli.CCListGenerator {
		if spec.locale == cli.DocsEN {
			return cli.NewIOListGenerator(cache, outputDir)
		}
		return cli.NewCCListGenerator(cache, outputDir)
	}
	newOSGen := func(cache *cli.ExtensionCache) *cli.CCOSGenerator {
		if spec.locale == cli.DocsEN {
			return cli.NewIOOSGenerator(cache, outputDir)
		}
		return cli.NewCCOSGenerator(cache, outputDir)
	}
	newCateGen := func(cache *cli.ExtensionCache) *cli.CCCateGenerator {
		if spec.locale == cli.DocsEN {
			return cli.NewIOCateGenerator(cache, outputDir)
		}
		return cli.NewCCCateGenerator(cache, outputDir)
	}
	newAttrGen := func(cache *cli.ExtensionCache) *cli.CCAttrGenerator {
		if spec.locale == cli.DocsEN {
			return cli.NewIOAttrGenerator(cache, outputDir)
		}
		return cli.NewCCAttrGenerator(cache, outputDir)
	}

	pageCmd := &cobra.Command{
		Use:     "page [extension-names...]",
		Short:   "Generate extension detail pages for " + spec.site,
		Long:    "Generate OINK Markdown detail pages for PostgreSQL extensions.\nIf no extension names are provided, generates pages for all extensions.",
		Example: fmt.Sprintf("  pgext gen %s page\n  pgext gen %s page vector postgis", spec.use, spec.use),
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			resolvedStub, err := cli.SourceStubDir(stubDir, spec.stubSubdir)
			if err != nil {
				return err
			}
			exts, err := selectExtensions(cache, args)
			if err != nil {
				return err
			}
			gen := newPageGen(cache, resolvedStub)
			if err := generateExtensionPages(ctx, resolvedStub, exts, len(args) > 0, filepath.Join(outputDir, "e"), gen.GenerateExtensionPage); err != nil {
				return err
			}
			if len(args) == 0 {
				removed, err := gen.PruneStaleExtensionPages(exts)
				if err != nil {
					return fmt.Errorf("failed to prune stale extension pages: %w", err)
				}
				if len(removed) > 0 {
					logrus.Infof("Removed %d stale extension pages: %s", len(removed), strings.Join(removed, ", "))
				}
			}
			return nil
		}),
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Generate list pages for " + spec.site,
		Args:  cobra.NoArgs,
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, _ []string) error {
			logrus.Info("Generating list pages...")
			if err := newListGen(cache).GenerateAllLists(); err != nil {
				return fmt.Errorf("failed to generate list pages: %w", err)
			}
			logrus.Info("Successfully generated all list pages")
			return nil
		}),
	}

	osCmd := &cobra.Command{
		Use:   "os [os-name]",
		Short: "Generate OS-specific availability pages for " + spec.site,
		Args:  cobra.MaximumNArgs(1),
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			gen := newOSGen(cache)
			if len(args) == 1 {
				logrus.Infof("Generating OS page for %s...", args[0])
				if err := gen.GenerateOSPage(ctx, args[0]); err != nil {
					return fmt.Errorf("failed to generate OS page for %s: %w", args[0], err)
				}
				logrus.Infof("Successfully generated OS page for %s", args[0])
				return nil
			}
			return gen.GenerateAllOSPages(ctx)
		}),
	}

	cateCmd := &cobra.Command{
		Use:   "cate [category-name]",
		Short: "Generate per-category PKG index pages for " + spec.site,
		Args:  cobra.MaximumNArgs(1),
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, args []string) error {
			gen := newCateGen(cache)
			if len(args) == 1 {
				catName := strings.ToUpper(args[0])
				cat, exists := cache.CateMap[catName]
				if !exists {
					return fmt.Errorf("category '%s' not found", args[0])
				}
				logrus.Infof("Generating category page for %s...", catName)
				if err := gen.GenerateCategoryPage(ctx, cat); err != nil {
					return fmt.Errorf("failed to generate category page for %s: %w", catName, err)
				}
				logrus.Infof("Successfully generated category page for %s", catName)
				return nil
			}
			return gen.GenerateAllCategoryPages(ctx)
		}),
	}

	attrCmd := &cobra.Command{
		Use:   "attr",
		Short: "Generate extension attribute pages for " + spec.site,
		Args:  cobra.NoArgs,
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, _ []string) error {
			logrus.Info("Generating attribute pages...")
			if err := newAttrGen(cache).GenerateAllAttrPages(); err != nil {
				return fmt.Errorf("failed to generate attribute pages: %w", err)
			}
			logrus.Info("Successfully generated all attribute pages")
			return nil
		}),
	}

	all = &cobra.Command{
		Use:   "all",
		Short: "Generate all content for " + spec.site,
		Args:  cobra.NoArgs,
		RunE: runWithCache(func(ctx context.Context, cache *cli.ExtensionCache, _ []string) error {
			resolvedStub, err := cli.SourceStubDir(stubDir, spec.stubSubdir)
			if err != nil {
				return err
			}
			exts := cache.ReadyExtensions()
			if err := cli.ValidateExtensionStubs(resolvedStub, exts); err != nil {
				return err
			}

			var errs []error
			var wg sync.WaitGroup
			var mu sync.Mutex
			runTask := func(name string, fn func() error) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					logrus.Infof("Generating %s...", name)
					if err := fn(); err != nil {
						mu.Lock()
						errs = append(errs, fmt.Errorf("%s: %w", name, err))
						mu.Unlock()
					} else {
						logrus.Infof("%s generated successfully", name)
					}
				}()
			}

			runTask("list pages", func() error { return newListGen(cache).GenerateAllLists() })
			runTask("OS pages", func() error { return newOSGen(cache).GenerateAllOSPages(ctx) })
			runTask("category pages", func() error { return newCateGen(cache).GenerateAllCategoryPages(ctx) })
			runTask("attribute pages", func() error { return newAttrGen(cache).GenerateAllAttrPages() })
			wg.Wait()

			logrus.Info("Generating extension pages...")
			pageGen := newPageGen(cache, resolvedStub)
			if err := generateExtensionPages(ctx, resolvedStub, exts, false, filepath.Join(outputDir, "e"), pageGen.GenerateExtensionPage); err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				for _, e := range errs {
					logrus.Error(e)
				}
				return fmt.Errorf("generation completed with %d errors", len(errs))
			}

			removed, err := pageGen.PruneStaleExtensionPages(exts)
			if err != nil {
				return fmt.Errorf("failed to prune stale extension pages: %w", err)
			}
			if len(removed) > 0 {
				logrus.Infof("Removed %d stale extension pages: %s", len(removed), strings.Join(removed, ", "))
			}
			logrus.Info("All content generated successfully!")
			return nil
		}),
	}

	parent = &cobra.Command{
		Use:   spec.use,
		Short: "Generate OINK content for " + spec.site,
		Long: fmt.Sprintf(`Generate OINK-native Markdown files for %s %s documentation site.

Run without a subcommand to generate all %s extension content.`, spec.site, spec.langLabel, spec.site),
		Example: fmt.Sprintf("  pgext gen %s page\n  pgext gen %s list\n  pgext gen %s os\n  pgext gen %s all", spec.use, spec.use, spec.use, spec.use),
		Args:    cobra.NoArgs,
		RunE:    defaultToSubcommand(all),
	}
	parent.AddCommand(pageCmd, listCmd, osCmd, cateCmd, attrCmd, all)

	defaultOutput := filepath.Join(os.Getenv("HOME"), "pgsty", spec.repoDir, "content", "ext")
	parent.PersistentFlags().StringVarP(&outputDir, "output", "o", defaultOutput, "Output directory for generated files")
	parent.PersistentFlags().StringVar(&stubDir, "stub-dir", "", fmt.Sprintf("Usage sources (default: %s/ in the pgext source checkout)", spec.stubSubdir))
	return parent, all
}

var (
	ccCmd, ccAllCmd = newDocsSiteCommand(docsSiteSpec{
		use: "cc", site: "pigsty.cc", langLabel: "Chinese",
		locale: cli.DocsZH, stubSubdir: "stub-zh", repoDir: "pigsty.cc",
	})
	ioCmd, ioAllCmd = newDocsSiteCommand(docsSiteSpec{
		use: "io", site: "pigsty.io", langLabel: "English",
		locale: cli.DocsEN, stubSubdir: "stub", repoDir: "pigsty.io",
	})
)

func init() {
	genCmd.AddCommand(ccCmd, ioCmd)
}
