package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/hatchify/queue"
	"github.com/hatchify/scribe"
)

// DefaultConfigLocation is the default configuration location
const DefaultConfigLocation = "./config.toml"

var (
	v   vpm
	out *scribe.Scribe

	q = queue.New(runtime.NumCPU(), 32)
)

func main() {
	configLocation := os.Getenv("VROOMY_CONFIG")
	if len(configLocation) == 0 {
		configLocation = DefaultConfigLocation
	}

	outW := scribe.NewStdout()
	outW.SetTypePrefix(scribe.TypeNotification, "")
	out = scribe.NewWithWriter(outW, "")
	out.Notification(":: Vroomy package manager ::")

	var err error
	if _, err = toml.DecodeFile(configLocation, &v.cfg); err != nil {
		handleError(err)
	}

	flag.Parse()
	cmd := flag.Arg(0)

	// Get any potential trailing args
	args := flag.Args()
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	switch cmd {
	case "update":
		if len(args) > 0 {
			out.Notificationf("Updating %v", args)
		} else {
			out.Notification("Updating packages")
		}
		if err = v.updatePlugins(args...); err != nil {
			handleError(err)
		}

		out.Success("Update complete")

	case "build":
		if len(args) > 0 {
			out.Notificationf("Building %v", args)
		} else {
			out.Notification("Building packages")
		}
		if err = v.buildPlugins(args...); err != nil {
			handleError(err)
		}

		out.Success("Build complete")

	case "test":
		// TODO: Finish this
		out.Error("Test not yet implemented")
	case "list":
		if len(args) > 0 {
			out.Notificationf("Listing %v", args)
		} else {
			out.Notification("Listing packages")
		}
		v.listPlugins(args...)

	case "help":
		// TODO: Use parg?
		out.Notification("Supported commands are: update, build, list, and help.")

	default:
		err = fmt.Errorf("invalid command, \"%s\" is not supported", cmd)
		handleError(err)
	}

}
