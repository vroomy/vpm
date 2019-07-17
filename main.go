package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/hatchify/output"
)

var v vpm

func main() {
	var err error
	if _, err = toml.DecodeFile("./config.toml", &v.cfg); err != nil {
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
