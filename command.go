package main

import (
	"strings"

	flag "github.com/hatchify/parg"
)

func commandFromArgs() (cmd *flag.Command, err error) {
	var p *flag.Parg
	p = flag.New()

	p.AddHandler("", help, "Manages vroomy packages.\n  To learn more, run `vpm help` or `vpm help <command>`")

	p.AddHandler("help", help, "Prints available commands and flags.\n  Use `vpm help <command>` to get more specific info.")
	p.AddHandler("version", printVersion, "Prints current version of vpm installation.\n  Use `vpm version`")
	p.AddHandler("upgrade", upgrade, "Upgrades vpm installation itself.\n  Skips if version is up to date.\n  Use `vpm upgrade` or `vpm upgrade <branch>`")

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
		msg = "Plugins matching: " + strings.Join(args, ", ")
	}

	return
}

func help(cmd *flag.Command) (err error) {
	if cmd == nil {
		out.Notification("Usage ::\n\n# Vroomy Package Manager\n" + flag.Help(true))
		return
	}

	out.Notification("Usage ::\n\n# Vroomy Package Manager\n" + cmd.Help(true))
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
