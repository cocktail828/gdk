package main

import (
	"fmt"

	router "github.com/cocktail828/gdk/v1/command"
)

var (
	GitTag         string
	BuildTime      string
	GitCommitLog   string
	BuildGoVersion string
)

func init() {
	fmt.Println("-> GitTag:", GitTag)
	fmt.Println("-> BuildTime:", BuildTime)
	fmt.Println("-> GitCommitLog:", GitCommitLog)
	fmt.Println("-> BuildGoVersion:", BuildGoVersion)
}

func main() {
	router.Execute()
}
