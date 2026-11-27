package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"

	"pgext/cli"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func catalogRoot() (string, error) {
	dir := genDir
	if dir == "" {
		dir = os.Getenv("PGEXT_SITE_DIR")
	}
	if dir == "" {
		dir = "~/pgsty/pgext.cloud"
	}
	return cli.ExpandDir(dir)
}

// A full refresh builds its four existing generators into a private directory
// before installing the completed set into the Hugo checkout.
type catalogStageKey struct{}

func catalogContentDir(cmd *cobra.Command) (string, error) {
	if stage, ok := cmd.Context().Value(catalogStageKey{}).(string); ok {
		return filepath.Join(stage, "content"), nil
	}
	if output, _ := cmd.Flags().GetString("output"); output != "" {
		return cli.ExpandDir(output)
	}
	root, err := catalogRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "content"), nil
}

var genSiteCmd = &cobra.Command{
	Use:   "site",
	Short: "Refresh the standalone Hugo catalog",
	Long:  "Generate extension, list and OS pages, the availability matrix and catalog CSV snapshots in the Hugo checkout.",
	Example: `  pgext gen site
  pgext gen site -D ../pgext-preview
  PGEXT_SITE_DIR=/tmp/preview pgext gen site -d data`,
	Args: cobra.NoArgs,
	RunE: generateCatalogSite,
}

func generateCatalogSite(cmd *cobra.Command, _ []string) error {
	root, err := catalogRoot()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}
	work, err := os.MkdirTemp(root, ".pgext-gen-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(work)
	stage := filepath.Join(work, "site")

	if err := cli.InitDB(dbURL); err != nil {
		return err
	}
	defer cli.CloseDB()

	ctx := cmd.Context()
	logrus.Info("Loading extension data...")
	cache, err := cli.LoadExtensionCache(ctx)
	if err != nil {
		return fmt.Errorf("failed to load extension cache: %w", err)
	}

	stageCtx := withSharedCache(context.WithValue(ctx, catalogStageKey{}, stage), cache)
	cmd.SetContext(stageCtx)
	defer cmd.SetContext(ctx)
	var failures []error
	for _, target := range []*cobra.Command{genPageCmd, genListCmd, genOSCmd, genMatrixCmd} {
		logrus.Infof("Running %s...", target.CommandPath())
		if err := runSubcommand(cmd, target, nil); err != nil {
			failures = append(failures, fmt.Errorf("%s: %w", target.CommandPath(), err))
		}
	}
	if err := errors.Join(failures...); err != nil {
		return err
	}
	if err := cli.ExportSiteCSVs(ctx, filepath.Join(stage, "assets", "db")); err != nil {
		return err
	}
	if err := publishCatalogSite(stage, root); err != nil {
		return err
	}
	logrus.Infof("Refreshed Hugo catalog in %s", root)
	return nil
}

// Promote only generated paths. An invalid target or a failed generator leaves
// existing pages intact. Roll back already installed files on a rename error.
func publishCatalogSite(stage, root string) error {
	dir, err := os.Open(root)
	if err != nil {
		return err
	}
	defer dir.Close()
	if err := syscall.Flock(int(dir.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(dir.Fd()), syscall.LOCK_UN)

	backup, err := os.MkdirTemp(filepath.Dir(stage), ".previous-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(backup)
	var changed []string
	err = filepath.WalkDir(stage, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		rel, err := filepath.Rel(stage, path)
		if err != nil {
			return err
		}
		dst := filepath.Join(root, rel)
		old, err := os.ReadFile(dst)
		if err == nil {
			fresh, err := os.ReadFile(path)
			if err != nil || bytes.Equal(old, fresh) {
				return err
			}
			saved := filepath.Join(backup, rel)
			if err := os.MkdirAll(filepath.Dir(saved), 0755); err != nil {
				return err
			}
			if err := os.Link(dst, saved); err != nil {
				return err
			}
		} else if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
			return err
		}
		changed = append(changed, rel)
		return nil
	})
	if err != nil {
		return err
	}
	for i, rel := range changed {
		if err := os.Rename(filepath.Join(stage, rel), filepath.Join(root, rel)); err != nil {
			failures := []error{err}
			for _, previous := range changed[:i] {
				dst := filepath.Join(root, previous)
				err := os.Rename(filepath.Join(backup, previous), dst)
				if errors.Is(err, fs.ErrNotExist) {
					err = os.Remove(dst)
				}
				failures = append(failures, err)
			}
			return errors.Join(failures...)
		}
	}
	return nil
}
