package main

import (
	"fmt"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

type arguments struct {
	BuildCommand string `long:"command" required:"true"`
	WatchFileExp string `long:"watch"`
}

func main() {
	var args arguments
	if _, err := flags.Parse(&args); err != nil {
		return
	}

	command := args.BuildCommand
	err := exec.Command(command).Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
