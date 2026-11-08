package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pgext/cli/build/proxy"
)

func newEnvProxyCommand() *cobra.Command {
	var jsonOutput bool
	c := &cobra.Command{Use: "proxy [id@remote] [local]", Short: "install or configure the build proxy", Args: cobra.MaximumNArgs(2), SilenceUsage: true}
	c.PersistentFlags().BoolVar(&jsonOutput, "json", false, "emit a credential-free JSON result")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		if len(args) == 0 {
			err = proxy.InstallProxy()
		} else {
			listen := proxy.DefaultProxyListen
			if len(args) > 1 {
				listen = args[1]
			}
			err = proxy.SetupProxy(args[0], listen)
		}
		if err != nil {
			return err
		}
		return printBuildValue(cmd, map[string]any{"action": "env proxy", "success": true}, jsonOutput)
	}
	var client proxy.ProxyClientOptions
	var cp bool
	cc := &cobra.Command{Use: "client [vless-uri]", Short: "configure an Xray HTTP/SOCKS client", Args: cobra.MaximumNArgs(1), SilenceUsage: true}
	cc.RunE = func(cmd *cobra.Command, args []string) error {
		client.URI = ""
		if len(args) > 0 {
			client.URI = args[0]
		}
		client.ListenSet = cmd.Flags().Changed("listen")
		if cp {
			p, err := proxy.PlanXrayClient(client, cmd.InOrStdin())
			if err != nil {
				return err
			}
			return printBuildValue(cmd, p, jsonOutput)
		}
		r, err := proxy.SetupXrayClient(client, cmd.InOrStdin())
		if err != nil {
			return err
		}
		return printBuildValue(cmd, r, jsonOutput)
	}
	f := cc.Flags()
	f.StringVar(&client.From, "from", "", "read URI from a file or '-' for stdin")
	f.StringVar(&client.Server, "server", "", "server host:port")
	f.StringVar(&client.ID, "id", "", "client UUID")
	f.StringVar(&client.SNI, "sni", "", "REALITY server name")
	f.StringVar(&client.PublicKey, "public-key", "", "client authentication key")
	f.StringVar(&client.ShortID, "short-id", "", "REALITY short ID")
	f.StringVar(&client.PQV, "pqv", "", "ML-DSA-65 verification key")
	f.StringVar(&client.Listen, "listen", proxy.DefaultProxyListen, "HTTP/SOCKS listener")
	f.BoolVar(&client.Replace, "replace", false, "replace a different client configuration")
	f.BoolVarP(&client.Yes, "yes", "y", false, "confirm replacement")
	f.BoolVar(&cp, "plan", false, "preview without changes")
	var server proxy.ProxyServerOptions
	var sp bool
	sc := &cobra.Command{Use: "server", Short: "configure or export an Xray REALITY server", Args: cobra.NoArgs, SilenceUsage: true}
	sc.RunE = func(cmd *cobra.Command, args []string) error {
		server.ListenSet = cmd.Flags().Changed("listen")
		server.ProxyProtocolSet = cmd.Flags().Changed("proxy-protocol")
		if server.Export != "" && jsonOutput {
			return fmt.Errorf("credential export requires text output; omit --json")
		}
		if sp {
			p, err := proxy.PlanXrayServer(server)
			if err != nil {
				return err
			}
			return printBuildValue(cmd, p, jsonOutput)
		}
		r, err := proxy.SetupXrayServer(server, cmd.OutOrStdout())
		if err != nil {
			return err
		}
		if server.Export != "" {
			return nil
		}
		return printBuildValue(cmd, r, jsonOutput)
	}
	f = sc.Flags()
	f.StringVar(&server.Host, "host", "", "advertised server domain/IP")
	f.IntVar(&server.Port, "port", 443, "public server port")
	f.StringVar(&server.Target, "target", "", "camouflage TLS endpoint")
	f.StringVar(&server.SNI, "sni", "", "allowed server name")
	f.StringVar(&server.Listen, "listen", "", "bind endpoint")
	f.BoolVar(&server.ProxyProtocol, "proxy-protocol", false, "PROXY protocol from an existing trusted frontend")
	f.StringVar(&server.Export, "export", "", "export URI to a protected file or '-' for stdout")
	f.BoolVar(&server.ExportOnly, "export-only", false, "read and export existing server without changing it")
	f.BoolVar(&sp, "plan", false, "preview without changes")
	c.AddCommand(cc, sc)
	return c
}
