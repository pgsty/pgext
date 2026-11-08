package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"pgext/cli/repo"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

var repoCmd = newRepoCommand()

func newRepoCommand() *cobra.Command {
	var platform, version, catalogPath, root string
	var opts repo.Options
	var asJSON, mirror bool
	c := &cobra.Command{
		Use: "repo", Short: "manage Linux package repositories from the confgen catalog", Aliases: []string{"r"},
		Args: cobra.NoArgs, SilenceUsage: true,
		Long: `Manage Linux repositories using the same definitions as Pigsty confgen.
Select repository names or modules; comma-separated selectors are accepted.
all = node + pgsql + infra; pgdg and pigsty select their stable repositories.
add/rm/set manage pgext's own files. set --replace also backs up and replaces
existing system sources. Each change has a guarded restore point.
These commands work without a metadata database; only summary needs one.`,
		Example: `  pgext repo list
  pgext repo list all
  pgext repo show pgsql --platform u24.arm64 --mirror
  pgext repo add pgdg pigsty --dry-run
  sudo pgext repo add pgdg pigsty --update
  sudo pgext repo set all --replace --mirror
  pgext repo status
  sudo pgext repo rm docker
  sudo pgext repo restore <backup-id>`,
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	f := c.PersistentFlags()
	f.StringVar(&platform, "platform", "", "target OS.ARCH for listing, preview or staging (default: detected Linux)")
	f.StringVar(&version, "os-version", "", "full OS version, e.g. 9.6 or 10.1")
	f.StringVar(&catalogPath, "catalog", "", "Pigsty YAML with repo_upstream or repo_upstream_default (default: embedded confgen catalog)")
	f.StringVar(&root, "root", "/", "existing staging root; never runs the host package manager")
	f.StringVar(&opts.Region, "region", "default", "mirror region; missing mirrors fall back to default")
	f.BoolVarP(&mirror, "mirror", "m", false, "use China mirrors (same as --region=china)")
	f.StringVar(&opts.AdminIP, "admin-ip", "", "host[:port] for the local Pigsty repository")
	f.BoolVar(&opts.DryRun, "dry-run", false, "preview changes without writing files or updating caches")
	f.BoolVar(&opts.DryRun, "plan", false, "alias for --dry-run")
	f.BoolVar(&asJSON, "json", false, "emit one JSON result on stdout")
	load := func(cmd *cobra.Command) (repo.Platform, []repo.Repository, string, error) {
		var p repo.Platform
		var err error
		if platform == "" {
			p, err = repo.DetectPlatform()
			if err == nil && version != "" {
				return p, nil, "", fmt.Errorf("--os-version requires --platform")
			}
		} else {
			p, err = repo.ParsePlatform(platform, version)
		}
		if err != nil {
			return p, nil, "", err
		}
		if mirror {
			if cmd.Flags().Changed("region") && opts.Region != "china" {
				return p, nil, "", fmt.Errorf("--mirror conflicts with --region=%s", opts.Region)
			}
			opts.Region = "china"
		}
		source := "embedded confgen catalog"
		var catalog []repo.Repository
		if catalogPath == "" {
			catalog, err = repo.DefaultCatalog("")
		} else {
			source = catalogPath
			var data []byte
			data, err = os.ReadFile(catalogPath)
			if err == nil {
				catalog, err = repo.ParseCatalog(data, p.Family)
			}
		}
		return p, catalog, source, err
	}
	for _, action := range []string{"list", "info", "show"} {
		var showAll bool
		sub := &cobra.Command{Use: action + " [repo|module...]", Args: cobra.ArbitraryArgs, SilenceUsage: true}
		switch action {
		case "list":
			sub.Short, sub.Aliases = "list available repositories", []string{"ls", "l"}
			sub.Flags().BoolVar(&showAll, "all", false, "include every OS and architecture")
		case "info":
			sub.Short, sub.Aliases = "show catalog definitions, mirrors and metadata", []string{"i"}
		case "show":
			sub.Short = "render the exact repository configuration without writing it"
		}
		sub.RunE = func(cmd *cobra.Command, args []string) error {
			p, catalog, source, err := load(cmd)
			if err != nil {
				return err
			}
			if action == "list" && len(args) == 1 && args[0] == "all" {
				showAll, args = true, nil
			}
			filter := p
			if showAll {
				filter = repo.Platform{}
			}
			if action == "show" && len(args) == 0 {
				args = []string{"all"}
			}
			selected, err := repo.Select(catalog, args, filter)
			if err != nil {
				return err
			}
			contents := map[string]string{}
			if action == "show" {
				if p.OS == "" {
					return fmt.Errorf("show requires --platform when not running on supported Linux")
				}
				for _, r := range selected {
					contents[r.Name], err = repo.Render(r, p, opts.RenderOptions)
					if err != nil {
						return err
					}
				}
			}
			if asJSON {
				return writeRepoJSON(cmd, map[string]any{"action": action, "success": true, "platform": filter, "catalog": source, "repositories": selected, "content": contents})
			}
			if action == "show" {
				for _, r := range selected {
					fmt.Fprintln(cmd.OutOrStdout(), contents[r.Name])
				}
				return nil
			}
			if action == "info" {
				enc := yaml.NewEncoder(cmd.OutOrStdout())
				defer enc.Close()
				return enc.Encode(selected)
			}
			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 4, 2, ' ', 0)
			fmt.Fprintln(w, "NAME\tMODULE\tTYPE\tRELEASES\tARCH\tDESCRIPTION")
			for _, r := range selected {
				fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%s\t%s\n", r.Name, r.Module, r.Family, r.Releases, strings.Join(r.Arch, ","), r.Description)
			}
			return w.Flush()
		}
		c.AddCommand(sub)
	}
	for _, action := range []string{"status", "add", "set", "rm", "update", "restore"} {
		sub := &cobra.Command{Use: action, Args: cobra.NoArgs, SilenceUsage: true}
		var update, replace, all bool
		switch action {
		case "status":
			sub.Short = "show installed repository files and pgext ownership"
		case "add", "set":
			sub.Use += " [repo|module...]"
			sub.Args = cobra.ArbitraryArgs
			sub.Short = action + " repositories (default: node, pgsql, infra)"
			sub.Flags().BoolVar(&replace, "replace", false, "back up and replace ALL existing system repository definitions")
		case "rm":
			sub.Use += " [repo|module...]"
			sub.Args = cobra.ArbitraryArgs
			sub.Short, sub.Aliases = "remove selected pgext-managed repositories", []string{"remove"}
			sub.Flags().BoolVar(&all, "all", false, "remove every pgext-managed repository, including custom modules")
		case "update":
			sub.Short = "refresh APT/DNF package metadata"
		case "restore":
			sub.Use += " <backup-id>"
			sub.Args = cobra.ExactArgs(1)
			sub.Short = "restore a change if the affected files have not changed since"
		}
		if action != "update" && action != "status" {
			sub.Flags().BoolVarP(&update, "update", "u", action == "set", "refresh package metadata after changes")
		}
		sub.RunE = func(cmd *cobra.Command, args []string) error {
			p, catalog, source, err := load(cmd)
			if err != nil {
				return err
			}
			m, err := repo.NewManager(root, p, catalog, cmd.ErrOrStderr())
			if err != nil {
				return err
			}
			opts.Update, opts.Replace, opts.All = update, replace, all
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()
			result, runErr := m.Execute(ctx, action, args, opts)
			result.Catalog = source
			if asJSON {
				if err := writeRepoJSON(cmd, result); err != nil {
					return err
				}
			} else {
				printRepoResult(cmd, result)
			}
			return runErr
		}
		c.AddCommand(sub)
	}
	// Keep database hooks on this subcommand, never on the management parent.
	c.AddCommand(newRepoSummaryCommand())
	return c
}

func writeRepoJSON(cmd *cobra.Command, value any) error {
	enc := json.NewEncoder(cmd.OutOrStdout())
	enc.SetIndent("", "  ")
	return enc.Encode(value)
}

func printRepoResult(cmd *cobra.Command, r *repo.Result) {
	w := cmd.OutOrStdout()
	if r.Action == "status" {
		for _, file := range r.Installed {
			owner := "external"
			if file.Managed {
				owner = "pgext"
			}
			fmt.Fprintf(w, "%-8s %s\n", owner, file.Path)
		}
		if r.Success && len(r.Installed) == 0 {
			fmt.Fprintln(w, "No repository files installed.")
		}
		return
	}
	label := "applied"
	if r.DryRun {
		label = "planned"
	} else if !r.Success {
		label = "failed"
	}
	fmt.Fprintf(w, "%s: %s on %s (%d file changes)\n", r.Action, label, r.Platform, len(r.Changes))
	for _, file := range r.Changes {
		op := "write"
		if file.After == nil {
			op = "remove"
		}
		fmt.Fprintf(w, "  %-6s %s\n", op, filepath.Join(r.Root, file.Path))
		if r.DryRun && file.After != nil {
			fmt.Fprintln(w, file.After.Content)
		}
	}
	for _, command := range r.Commands {
		fmt.Fprintf(w, "  cache: %s\n", strings.Join(command, " "))
	}
	if r.Backup != "" {
		fmt.Fprintf(w, "Backup: %s/var/lib/pgext/repo/backups/%s\n", strings.TrimRight(r.Root, "/"), r.Backup)
		fmt.Fprintf(w, "Restore: pgext repo restore %s\n", r.Backup)
	}
	if r.RolledBack {
		fmt.Fprintln(w, "Repository file changes were rolled back.")
	}
}
