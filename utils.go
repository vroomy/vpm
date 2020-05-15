package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/hatchify/parg"
)

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

func parseArgs() (cmd *flag.Command, args []string, msg string) {
	var p *flag.Parg
	p = flag.New()

	p.AddAction("", "Manages vroomy packages.\n  To learn more, run `vpm help` or `vpm help <command>`")
	p.AddAction("help", "Prints available commands and flags.\n  Use `vpm help <command>` to get more specific info.")

	p.AddAction("doc", "Outputs postman docs for specified plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm doc` for all plugins, or `vpm doc <plugin> <plugin>`")

	p.AddAction("update", "Loads specified version or latest channel from config and builds plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm update` for all plugins, or `vpm update <plugin> <plugin>`")
	p.AddAction("build", "Builds the currently checked out version of plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm build` for all plugins, or `vpm build <plugin> <plugin>`")
	p.AddAction("list", "Lists the plugin(s) and associated key/alias.\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm list` for all plugins, or `vpm list <plugin> <plugin>`")
	//p.AddAction("test", "Tests the currently checked out version of plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm test` for all plugins, or `vpm test <plugin> <plugin>`")

	cmd, err := flag.Validate()
	if err != nil {
		showHelp(cmd)
		handleError(err)
	}

	args = cmd.Args()
	if len(args) == 0 {
		msg = "all Plugins"
	} else {
		msg = strings.Join(args, ", ")
	}

	return
}

func showHelp(cmd *flag.Command) {
	if cmd == nil {
		fmt.Println(flag.Help())
	} else {
		fmt.Println(cmd.Help())
	}
}
