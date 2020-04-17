package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Parse Command Line Arguments
func parse() (cmd string, args []string, msg string) {
	flag.Parse()

	// Get vpm Command
	cmd = flag.Arg(0)

	// Get any potential trailing args
	if len(flag.Arg(1)) > 0 {
		args = flag.Args()[1:]
		msg = fmt.Sprintf("%v", args)
	} else {
		msg = "plugins"
	}

	return
}

func handleError(err error) {
	out.Error(err.Error())
	os.Exit(1)
}

func keyHasSuffixInAny(key string, pluginNames ...string) bool {
	for _, name := range pluginNames {
		if strings.HasSuffix(name, key) {
			return true
		}
	}

	return false
}
