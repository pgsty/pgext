package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"pgext/cli"

	"github.com/spf13/cobra"
)

func newDocCommand() *cobra.Command {
	docCmd := &cobra.Command{
		Use:   "doc",
		Short: "Load or dump bilingual extension stubs in pgext.doc",
	}

	resolveStubDirs := func(en, zh string) (string, string, error) {
		if en == "" || zh == "" {
			root, err := cli.SourceDir()
			if err != nil {
				return "", "", err
			}
			if en == "" {
				en = filepath.Join(root, "stub")
			}
			if zh == "" {
				zh = filepath.Join(root, "stub-zh")
			}
		}
		return en, zh, nil
	}

	connectDocDB := func(explicit string) error {
		if explicit == "" {
			explicit = dbURL
		}
		return cli.InitDB(explicit)
	}

	var (
		loadPGURL  string
		loadEN     string
		loadZH     string
		loadDryRun bool
	)
	loadDocCmd := &cobra.Command{
		Use:   "load",
		Short: "Atomically sync stub/ and stub-zh/ into pgext.doc against pgext.universe",
		Example: `  pgext doc load --dry-run
  pgext doc load --pgurl 'host=/tmp port=5432 dbname=data user=postgres'`,
		RunE: func(cmd *cobra.Command, args []string) error {
			enDir, zhDir, err := resolveStubDirs(loadEN, loadZH)
			if err != nil {
				return err
			}
			if err := connectDocDB(loadPGURL); err != nil {
				return fmt.Errorf("failed to connect to database: %w", err)
			}
			defer cli.CloseDB()

			stats, err := cli.LoadStubDocs(cmd.Context(), cli.DB, enDir, zhDir, loadDryRun)
			if err != nil {
				return err
			}
			if len(stats.ExtraEN) > 0 {
				cmd.PrintErrf("warning: ignoring %d extra EN stubs not in pgext.universe: %s\n",
					len(stats.ExtraEN), strings.Join(stats.ExtraEN, ", "))
			}
			if len(stats.ExtraZH) > 0 {
				cmd.PrintErrf("warning: ignoring %d extra ZH stubs not in pgext.universe: %s\n",
					len(stats.ExtraZH), strings.Join(stats.ExtraZH, ", "))
			}
			mode := "loaded"
			if stats.DryRun {
				mode = "dry-run verified"
			}
			cmd.Printf("pgext.doc %s: %d/%d universe rows (EN files=%d, ZH files=%d)\n",
				mode, stats.LoadedCount, stats.UniverseCount, stats.ENFiles, stats.ZHFiles)
			return nil
		},
	}
	loadDocCmd.Flags().StringVar(&loadPGURL, "pgurl", "", "PostgreSQL connection URL (overrides -d and PGURL)")
	loadDocCmd.Flags().StringVar(&loadEN, "stub-dir", "", "English stub directory (default: stub/)")
	loadDocCmd.Flags().StringVar(&loadZH, "zh-stub-dir", "", "Chinese stub directory (default: stub-zh/)")
	loadDocCmd.Flags().BoolVar(&loadDryRun, "dry-run", false, "validate and stage inside a rolled-back transaction")

	var (
		dumpPGURL string
		dumpEN    string
		dumpZH    string
	)
	dumpDocCmd := &cobra.Command{
		Use:   "dump",
		Short: "Export en_doc and zh_doc from pgext.doc into stub/ and stub-zh/",
		RunE: func(cmd *cobra.Command, args []string) error {
			enDir, zhDir, err := resolveStubDirs(dumpEN, dumpZH)
			if err != nil {
				return err
			}
			if err := connectDocDB(dumpPGURL); err != nil {
				return fmt.Errorf("failed to connect to database: %w", err)
			}
			defer cli.CloseDB()

			enCount, zhCount, err := cli.DumpStubDocs(cmd.Context(), cli.DB, enDir, zhDir)
			if err != nil {
				return err
			}
			cmd.Printf("Dumped pgext.doc: %d EN stubs -> %s, %d ZH stubs -> %s\n", enCount, enDir, zhCount, zhDir)
			return nil
		},
	}
	dumpDocCmd.Flags().StringVar(&dumpPGURL, "pgurl", "", "PostgreSQL connection URL (overrides -d and PGURL)")
	dumpDocCmd.Flags().StringVar(&dumpEN, "stub-dir", "", "English stub output directory (default: stub/)")
	dumpDocCmd.Flags().StringVar(&dumpZH, "zh-stub-dir", "", "Chinese stub output directory (default: stub-zh/)")

	docCmd.AddCommand(loadDocCmd, dumpDocCmd)
	return docCmd
}
