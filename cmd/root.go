package cmd

import (
	"regexp"
	"strings"

	"github.com/cocktail828/gdk/v1/cmd/status"
	"github.com/cocktail828/go-tools/log"
	"github.com/spf13/cobra"
)

var (
	re = regexp.MustCompile(`([v]?\d+\.\d+\.\d+)`)
)

// v4.4.5 -> 4.4.5
func Version(ver string) string {
	ver = strings.ToLower(ver)
	if strings.HasPrefix(ver, "v") {
		ver = ver[1:]
	}

	match := re.FindStringSubmatch(ver)
	if len(match) > 0 {
		return match[1]
	}
	return ver
}

func Execute() {
	status.AppVersion = Version(status.AppVersion)
	rootCmd := &cobra.Command{
		Use:     "gdk",
		Short:   "general development kits",
		Version: status.AppVersion,
	}

	rootCmd.AddCommand(newProxyCmd())
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
