package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"pgext/cli/build"
	"syscall"

	"github.com/spf13/cobra"
)

func buildContext(cmd *cobra.Command) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
}

func newBuildCommand(action string) *cobra.Command {
	var opts build.Options
	var jsonOutput bool
	short := map[string]string{"get": "download extension source archives", "dep": "install recipe build dependencies", "make": "build extension packages from prepared inputs", "pkg": "fetch sources, install dependencies, and build packages"}[action]
	c := &cobra.Command{Use: action + " <extension|package...>", Short: short, Args: cobra.MinimumNArgs(1), SilenceUsage: true}
	if action == "make" {
		c.Aliases = []string{"m", "mk", "mak"}
		c.Example = "  pgext mak pgvector --pg 18\n  pgext mak vector --pg 17,18 --json"
	} else {
		c.Example = "  pgext " + action + " pgvector --pg 18"
	}
	if action == "pkg" {
		c.Long = "Run get → dep → make, stopping on failure. The former read-only pkg query is now pgext matrix <name>."
	}
	buildFlags(c, &opts, &jsonOutput)
	c.RunE = func(cmd *cobra.Command, args []string) error {
		e, err := build.New(opts, cmd.ErrOrStderr())
		if err != nil {
			return err
		}
		e.Version = Version + " (" + Commit + ")"
		p, err := e.Prepare(action, args)
		if err != nil {
			return err
		}
		if opts.Plan {
			return printBuildValue(cmd, p, jsonOutput)
		}
		ctx, stop := buildContext(cmd)
		defer stop()
		r, runErr := e.Run(ctx, p)
		if r != nil {
			if r.ToolVersions == nil {
				r.ToolVersions = map[string]string{}
			}
			r.ToolVersions["pgext"] = Version + " (" + Commit + ")"
			if err := printBuildValue(cmd, r, jsonOutput); err != nil {
				return err
			}
		}
		return runErr
	}
	return c
}

func buildFlags(c *cobra.Command, o *build.Options, jsonOutput *bool) {
	f := c.Flags()
	f.StringVarP(&o.PG, "pg", "p", "", "PostgreSQL majors, comma-separated (default: installed recipe targets)")
	if c.Name() == "get" || c.Name() == "dep" || c.Name() == "make" || c.Name() == "pkg" {
		f.StringVar(&o.CatalogFile, "catalog", "", "catalog CSV (default: XDG pgext/extension.csv, then embedded)")
	}
	f.StringVar(&o.SpecDir, "spec-dir", "", "recipe root (default: ~/rpmbuild or ~/debbuild)")
	f.StringVar(&o.Root, "work-dir", "", "build workspace (default: ~/ext)")
	f.StringVar(&o.Platform, "platform", "", "OS.ARCH for preview, e.g. el9.arm64 or u24.amd64")
	f.BoolVar(&o.Plan, "plan", false, "preview without downloads, installs, or builds")
	f.BoolVar(jsonOutput, "json", false, "emit one JSON result on stdout")
	switch c.Name() {
	case "get", "pkg", "spec", "rust":
		f.BoolVarP(&o.Mirror, "mirror", "m", false, "prefer the China mirror")
		f.BoolVarP(&o.Force, "force", "f", false, "refresh cached downloads (rust: back up and replace Cargo mirror configuration)")
	}
	if c.Name() == "get" || c.Name() == "pkg" {
		f.StringVar(&o.SourceURL, "source-url", "", "source archive base URL")
	}
	if c.Name() == "get" || c.Name() == "pkg" || c.Name() == "make" {
		f.StringVar(&o.Checksums, "checksums", "", "JSON filename-to-SHA256 source manifest")
	}
	if c.Name() == "make" || c.Name() == "pkg" {
		f.IntVarP(&o.Jobs, "jobs", "j", 0, "compiler parallelism (0: recipe default)")
		f.BoolVar(&o.NoDebug, "no-debug", false, "explicitly omit debug packages (default: preserve debug information)")
	}
}

func printBuildValue(cmd *cobra.Command, v any, asJSON bool) error {
	if asJSON {
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
	switch x := v.(type) {
	case *build.Plan:
		fmt.Fprintf(cmd.OutOrStdout(), "%s on %s.%s; catalog %s\n", x.Action, x.Platform.OS, x.Platform.Arch, x.Catalog.Path)
		for _, j := range x.Jobs {
			fmt.Fprintf(cmd.OutOrStdout(), "  %s PG%v: %v\n    recipe: %s\n    sources: %v\n", j.Package, j.PG, j.Stages, j.Recipe, j.Sources)
		}
	case *build.Result:
		fmt.Fprintf(cmd.OutOrStdout(), "%s success=%t; %d artifacts; manifest: %s\n", x.Action, x.Success, len(x.Artifacts), x.Manifest)
	case *build.EnvResult:
		fmt.Fprintf(cmd.OutOrStdout(), "%s on %s.%s success=%t plan=%t\n", x.Action, x.Platform.OS, x.Platform.Arch, x.Success, x.Plan)
		for _, p := range x.Paths {
			fmt.Fprintln(cmd.OutOrStdout(), p)
		}
		if x.Plan {
			for _, c := range x.Commands {
				fmt.Fprintln(cmd.OutOrStdout(), c.Args)
			}
		}
	default:
		return json.NewEncoder(cmd.OutOrStdout()).Encode(v)
	}
	return nil
}

func newEnvCommand() *cobra.Command {
	c := &cobra.Command{Use: "env", Short: "inspect or prepare the extension build environment", Args: cobra.NoArgs, SilenceUsage: true}
	var jsonStatus bool
	c.Flags().BoolVar(&jsonStatus, "json", false, "emit JSON")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		e, err := build.NewEnvironment(build.Options{}, cmd.ErrOrStderr())
		if err != nil {
			return err
		}
		r, err := e.Environment(cmd.Context(), "status", build.EnvOptions{})
		if err != nil {
			return err
		}
		return printBuildValue(cmd, r, jsonStatus)
	}
	for _, action := range []string{"tool", "spec", "rust", "pgrx"} {
		var opts build.EnvOptions
		var asJSON bool
		sub := &cobra.Command{Use: action, Short: "prepare build " + action, Args: cobra.NoArgs, SilenceUsage: true}
		buildFlags(sub, &opts.Options, &asJSON)
		switch action {
		case "tool":
			sub.Use = "tool [full|mini]"
			sub.Args = cobra.MaximumNArgs(1)
		case "spec":
			sub.Flags().StringVar(&opts.From, "from", "", "local recipe directory or tar.gz archive")
		case "pgrx":
			sub.Flags().StringVarP(&opts.PgrxVersion, "pgrx", "v", build.DefaultPgrxVersion, "cargo-pgrx version")
		}
		sub.RunE = func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.Mode = args[0]
			}
			e, err := build.NewEnvironment(opts.Options, cmd.ErrOrStderr())
			if err != nil {
				return err
			}
			ctx, stop := buildContext(cmd)
			defer stop()
			r, runErr := e.Environment(ctx, action, opts)
			if err := printBuildValue(cmd, r, asJSON); err != nil {
				return err
			}
			return runErr
		}
		c.AddCommand(sub)
	}
	c.AddCommand(newEnvRepoCommand(), newEnvProxyCommand())
	return c
}
