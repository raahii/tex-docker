package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

var (
	dockerImage = "raahii/tex-docker:latest"
)

type arguments struct {
	Path          string `short:"p" long:"path" required:"true" description:"Latex source path to compile (execute latexmk)"`
	ContainerName string `short:"n" long:"container-name" description:"Docker container name"`
	WatchExp      string `short:"w" long:"watch" description:"Process any events whose filename matches the specified POSIX extended regular expression"`
	Recursive     bool   `short:"r" long:"recursive" description:"Watch all subdirectories of the --path directory"`
}

func buildCommand(args *arguments) *exec.Cmd {
	// build commands
	cmds := []string{"docker", "run", "--rm"}

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

	// command in docker container
	cmds = append(cmds, "tex-docker")
	cmds = append(cmds, "--command")
	cmds = append(cmds, "latexmk")

	if args.WatchExp != "" {
		cmds = append(cmds, "--watch")
		cmds = append(cmds, args.WatchExp)
	}

	if args.Recursive {
		cmds = append(cmds, "--recursive")
	}

	// build command
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func main() {
	var args arguments
	if _, err := flags.Parse(&args); err != nil {
		return
	}

	cmd := buildCommand(&args)
	cmd.Start()

	cmd.Wait()
}
