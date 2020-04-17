package main

import (
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

	// TODO: Use parg for command parsing?
	cmd, args, msg := parse()

	switch cmd {
	case "update":
		out.Notificationf("Updating %s...", msg)

		if err = v.updatePlugins(args...); err != nil {
			handleError(err)
		}

		out.Success("Update complete!")

	case "build":
		out.Notificationf("Building %s...", msg)

		if err = v.buildPlugins(args...); err != nil {
			handleError(err)
		}

		out.Success("Build complete!")

	case "test":
		out.Notificationf("Testing %s...", msg)

		if err = v.testPlugins(args...); err != nil {
			handleError(err)
		}

		out.Success("Test complete!")

	case "list":
		out.Notificationf("Listing %s...", msg)

		v.listPlugins(args...)

	case "help":
		// TODO: Use parg for help docs?
		out.Notification("Supported commands are: update, build, list, and help.")

	default:
		err = fmt.Errorf("invalid command, \"%s\" is not supported", cmd)
		handleError(err)
	}

}
