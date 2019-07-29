package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hatchify/output"
)

// DefaultConfigLocation is the default configuration location
const DefaultConfigLocation = "./config.toml"

var v vpm

func main() {
	configLocation := os.Getenv("config")
	if len(configLocation) == 0 {
		configLocation = DefaultConfigLocation
	}

	var err error
	if _, err = toml.DecodeFile(configLocation, &v.cfg); err != nil {
		handleError(err)
	}

	output.Print(":: Vroomy package manager ::")
	flag.Parse()
	cmd := flag.Arg(0)

	switch cmd {
	case "update":
		output.Print("Updating packages")
		if err = v.updatePlugins(); err != nil {
			handleError(err)
		}

		output.Success("Update complete")

	case "build":
		output.Print("Building packages")
		if err = v.buildPlugins(); err != nil {
			handleError(err)
		}

		output.Success("Build complete")

	case "list":
		// TODO: Finish this

	case "help":
		output.Print("Supported commands are: update, list, and help.")

	default:
		err = fmt.Errorf("invalid command, \"%s\" is not supported", cmd)
		handleError(err)
	}

}
