package command

import (
	"github.com/cocktail828/gdk/v1/logger"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gdk",
	Short: "general development kit",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Default().Panicln(err)
	}
}
