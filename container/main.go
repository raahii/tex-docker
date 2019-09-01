package main

import (
	"bytes"
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

	cmd := exec.Command(args.BuildCommand)
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(stderr.String())
		return
	}
	fmt.Println(string(out))
}
