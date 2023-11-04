package command

import (
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gdk",
	Short: "general development kit",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
