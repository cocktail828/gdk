package main

import (
	"fmt"

	"github.com/cocktail828/gdk/v1/command"
)

var (
	GitTag       string
	BuildTime    string
	GitCommitLog string
	GoVersion    string
)

func init() {
	fmt.Println("-> GitTag:", GitTag)
	fmt.Println("-> GitCommitLog:", GitCommitLog)
	fmt.Println("-> BuildTime:", BuildTime)
	fmt.Println("-> GoVersion:", GoVersion)
}

func main() {
	command.Execute()
}
