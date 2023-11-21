package main

import (
	"log"

	"github.com/cocktail828/gdk/v1/command"
)

var (
	GitRemote string
	GitTag    string
	BuildTime string
	GitCommit string
	GoVersion string
)

func init() {
	log.Println("-> GitRemote:", GitRemote)
	log.Println("-> GitTag:", GitTag)
	log.Println("-> GitCommit:", GitCommit)
	log.Println("-> BuildTime:", BuildTime)
	log.Println("-> GoVersion:", GoVersion)
}

func main() {
	command.Execute()
}
