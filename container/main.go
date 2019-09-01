package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/aybabtme/rgbterm"
	"github.com/jessevdk/go-flags"
	"github.com/radovskyb/watcher"
)

type arguments struct {
	Command   string `long:"command" required:"true"`
	WatchExp  string `long:"watch"`
	Recursive bool   `long:"recursive"`
}

func errorMessage(texts ...string) string {
	var r, g, b uint8
	r, g, b = 199, 37, 78

	coloredTexts := make([]string, len(texts))
	for i, v := range texts {
		coloredTexts[i] = rgbterm.FgString(v, r, g, b)
	}
	return strings.Join(coloredTexts, " ")
}

func infoMessage(texts ...string) string {
	var r, g, b uint8
	r, g, b = 70, 192, 31

	coloredTexts := make([]string, len(texts))
	for i, v := range texts {
		coloredTexts[i] = rgbterm.FgString(v, r, g, b)
	}
	return strings.Join(coloredTexts, " ")
}

func warnMessage(texts ...string) string {
	var r, g, b uint8
	r, g, b = 227, 192, 133

	coloredTexts := make([]string, len(texts))
	for i, v := range texts {
		coloredTexts[i] = rgbterm.FgString(v, r, g, b)
	}
	return strings.Join(coloredTexts, " ")
}

func executeCommand(command []string) {
	if len(command) == 0 {
		fmt.Println(errorMessage("argument command is empty"))
		return
	}
	fmt.Println(infoMessage("[exec]", strings.Join(command, " ")))

	var cmd *exec.Cmd
	if len(command) == 1 {
		cmd = exec.Command(command[0])
	} else {
		cmd = exec.Command(command[0], command[1:]...)
	}

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(errorMessage(stderr.String()))
		return
	}
	fmt.Println(string(out))
}

func main() {
	var args arguments
	if _, err := flags.Parse(&args); err != nil {
		fmt.Println(errorMessage(err.Error()))
		return
	}

	command := strings.Split(args.Command, " ")

	// if watcher is not needed, just compile at once.
	if args.WatchExp == "" {
		executeCommand(command)
		return
	}

	// watch specified filesand compile every time the files changes
	w := watcher.New()
	w.SetMaxEvents(1)

	r := regexp.MustCompile(args.WatchExp)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(infoMessage("[event]", event.String()))
				executeCommand(command)
				fmt.Println(infoMessage("Waiting file changes..."))
			case err := <-w.Error:
				fmt.Println(errorMessage(err.Error()))
			case <-w.Closed:
				return
			}
		}
	}()

	// watch current directory
	if args.Recursive {
		fmt.Println(infoMessage("adding subdirectories of the path, this may takes a while..."))
		fmt.Println()
		if err := w.AddRecursive("."); err != nil {
			fmt.Println(errorMessage(err.Error()))
		}
	} else {
		if err := w.Add("."); err != nil {
			fmt.Println(errorMessage(err.Error()))
		}
	}

	// print a list of all of the files and folders currently
	// being watched and their paths.
	s := fmt.Sprintf("'%s'", args.WatchExp)
	fmt.Println(warnMessage("watch only files that match", s))
	if len(w.WatchedFiles()) == 0 {
		fmt.Println(errorMessage("No files to watch !!"))
	}
	for path, _ := range w.WatchedFiles() {
		s := fmt.Sprintf("  - %s", path)
		fmt.Println(warnMessage(s))
	}
	fmt.Println()

	if err := w.Start(time.Millisecond * 500); err != nil {
		fmt.Println(errorMessage(err.Error()))
	}
}
