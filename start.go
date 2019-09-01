package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type arguments struct {
	Path string `short:"p" long:"path" required:"true" description:"latex source directory to compile"`
}

func startDocker(args *arguments) {
}

func main() {
	var args arguments
	if _, err := flags.Parse(&args); err != nil {
		os.Exit(1)
	}

	startDocker(&args)
}
