package command

import (
	"github.com/cocktail828/gdk/v1/must"
	"github.com/cocktail828/gdk/v1/servers/httpd"
	"github.com/spf13/cobra"
)

var ProxyCmd = &cobra.Command{
	Use:   "http",
	Short: "gdk http server",
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		httpd.Start(address)
	},
}

func init() {
	RootCmd.AddCommand(ProxyCmd)
	ProxyCmd.Flags().StringP("address", "a", ":8080", "HTTP address, host:port")
	must.Must(ProxyCmd.MarkFlagRequired("address"))
}
