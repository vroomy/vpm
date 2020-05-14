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
	var strippedKey string
	strippedKey = removeBranchHash(key)
	for _, name := range pluginNames {
		if strings.HasSuffix(key, name) || strings.HasSuffix(strippedKey, name) {
			return true
		}
	}

	return false
}

func removeBranchHash(gitURL string) (out string) {
	out = strings.Split(gitURL, "#")[0]
	out = strings.Split(out, "@")[0]
	return
}
