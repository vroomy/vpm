package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hatchify/scribe"
)

// DefaultConfigLocation is the default configuration location
const DefaultConfigLocation = "./config.toml"

var (
	v   vpm
	out *scribe.Scribe
)

func main() {
	configLocation := os.Getenv("VROOMY_CONFIG")
	if len(configLocation) == 0 {
		configLocation = DefaultConfigLocation
	}

	var err error
	if _, err = toml.DecodeFile(configLocation, &v.cfg); err != nil {
		handleError(err)
	}

	outW := scribe.NewStdout()
	outW.SetTypePrefix(scribe.TypeNotification, "")
	out = scribe.NewWithWriter(outW, "")
	out.Notification(":: Vroomy package manager ::")

	flag.Parse()
	cmd := flag.Arg(0)

	switch cmd {
	case "update":
		out.Notification("Updating packages")
		if err = v.updatePlugins(); err != nil {
			handleError(err)
		}

		out.Success("Update complete")

	case "build":
		out.Notification("Building packages")
		if err = v.buildPlugins(); err != nil {
			handleError(err)
		}

		out.Success("Build complete")

	case "list":
		// TODO: Finish this

	case "help":
		out.Notification("Supported commands are: update, list, and help.")

	default:
		err = fmt.Errorf("invalid command, \"%s\" is not supported", cmd)
		handleError(err)
	}

}
