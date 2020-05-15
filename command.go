package main

import (
	"fmt"
	"strings"

	flag "github.com/hatchify/parg"
)

func commandFromArgs() (cmd *flag.Command, err error) {
	var p *flag.Parg
	p = flag.New()

	p.AddHandler("", help, "Manages vroomy packages.\n  To learn more, run `vpm help` or `vpm help <command>`")
	p.AddHandler("help", help, "Prints available commands and flags.\n  Use `vpm help <command>` to get more specific info.")

	p.AddHandler("doc", doc, "Outputs postman docs for specified plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm doc` for all plugins, or `vpm doc <plugin> <plugin>`")

	p.AddHandler("update", update, "Loads specified version or latest channel from config and builds plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm update` for all plugins, or `vpm update <plugin> <plugin>`")
	p.AddHandler("build", build, "Builds the currently checked out version of plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm build` for all plugins, or `vpm build <plugin> <plugin>`")

	p.AddHandler("list", list, "Lists the plugin(s) and associated key/alias.\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm list` for all plugins, or `vpm list <plugin> <plugin>`")
	p.AddHandler("test", test, "Tests the currently checked out version of plugin(s).\n  Accepts filtered trailing args to target specific plugins.\n  Use `vpm test` for all plugins, or `vpm test <plugin> <plugin>`")

	cmd, err = flag.Validate()
	return
}

func commandParams(cmd *flag.Command) (args []string, msg string) {
	args = cmd.Args()

	msg = "all Plugins"
	if len(args) > 0 {
		msg = strings.Join(args, ", ")
	}

	return
}

func help(cmd *flag.Command) (err error) {
	if cmd == nil {
		fmt.Println(flag.Help())
		return
	}

	fmt.Println(cmd.Help())
	return
}

func doc(cmd *flag.Command) (err error) {
	_, msg := commandParams(cmd)
	out.Notificationf("Documenting %s...", msg)

	/*
		if err := v.documentPlugins(args...); err != nil {
			handleError(err)
		}
	*/
	return
}

func update(cmd *flag.Command) (err error) {
	args, msg := commandParams(cmd)
	out.Notificationf("Updating %s...", msg)

	if err := v.updatePlugins(args...); err != nil {
		handleError(err)
	}

	out.Success("Update complete!")
	return
}

func build(cmd *flag.Command) (err error) {
	args, msg := commandParams(cmd)
	out.Notificationf("Building %s...", msg)

	if err := v.buildPlugins(args...); err != nil {
		handleError(err)
	}

	out.Success("Build complete!")
	return
}

func test(cmd *flag.Command) (err error) {
	args, msg := commandParams(cmd)
	out.Notificationf("Testing %s...", msg)

	if err := v.testPlugins(args...); err != nil {
		handleError(err)
	}

	out.Success("Test complete!")
	return
}

func list(cmd *flag.Command) (err error) {
	args, msg := commandParams(cmd)
	out.Notificationf("Listing %s...", msg)

	v.listPlugins(args...)
	return
}
