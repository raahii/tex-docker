package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

var (
	dockerImage string = "tex-docker"
)

type arguments struct {
	Path          string `short:"p" long:"path" required:"true" description:"latex source directory to compile"`
	ContainerName string `short:"n" long:"container-name" description:"docker container name"`
}

func startDocker(args *arguments) (string, error) {
	// build commands
	cmds := []string{"docker", "run", "-d"}

	// volume option
	volume := fmt.Sprintf("%s:/home/work", args.Path)
	cmds = append(cmds, "--volume")
	cmds = append(cmds, volume)

	// name option
	if len(args.ContainerName) > 0 {
		cmds = append(cmds, "--name")
		cmds = append(cmds, args.ContainerName)
	}

	// docker image name
	cmds = append(cmds, dockerImage)

	// build command
	cmd := exec.Command(cmds[0], cmds[1:]...)
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	out, err := cmd.Output()
	if err != nil {
		return string(out), errors.New(stderr.String())
	}
	return string(out), nil
}

func main() {
	var args arguments
	if _, err := flags.Parse(&args); err != nil {
		return
	}

	out, err := startDocker(&args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
