package cmd

import (
	"github.com/cocktail828/gdk/v1/cmd/entry"
	"github.com/cocktail828/gdk/v1/cmd/status"
	"github.com/cocktail828/go-tools/z"
	"github.com/spf13/cobra"
)

func newProxyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "gdk proxy",
		Run: func(cmd *cobra.Command, args []string) {
			inMode, _ := cmd.Flags().GetInt("mode")
			inCfg, _ := cmd.Flags().GetString("cfg")
			inGroup, _ := cmd.Flags().GetString("group")
			inService, _ := cmd.Flags().GetString("service")
			inRegistry, _ := cmd.Flags().GetString("companionUrl")
			status.AppVersion, _ = cmd.Flags().GetString("version")
			status.ApiVersion, _ = cmd.Flags().GetString("apiVersion")
			entry.Start(inMode, inCfg, inGroup, inService, inRegistry)
		},
	}
	cmd.Flags().StringP("version", "v", status.AppVersion, "app(config) version")
	cmd.Flags().IntP("mode", "m", 1, "config load mode")
	cmd.Flags().StringP("cfg", "c", "", "cfgName")
	cmd.Flags().StringP("group", "g", "", "group")
	cmd.Flags().StringP("service", "s", "", "service")
	cmd.Flags().StringP("registry", "r", "", "registry url")
	cmd.Flags().StringP("apiVersion", "V", "1.0.0", "apiVersion")

	z.Must(cmd.MarkFlagRequired("cfg"))
	z.Must(cmd.MarkFlagRequired("group"))
	z.Must(cmd.MarkFlagRequired("service"))
	z.Must(cmd.MarkFlagRequired("registry"))
	return cmd
}
