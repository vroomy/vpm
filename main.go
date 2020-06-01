package main

import (
	"os"
	"runtime"

	flag "github.com/hatchify/parg"
	"github.com/hatchify/queue"
	"github.com/hatchify/scribe"
	"github.com/vroomy/config"
)

// DefaultConfigLocation is the default configuration location
const DefaultConfigLocation = "./config.toml"

var (
	v    vpm
	out  *scribe.Scribe
	outW *scribe.Stdout

	q = queue.New(runtime.NumCPU(), 32)
)

func main() {
	var err error
	outW = scribe.NewStdout()
	outW.SetTypePrefix(scribe.TypeNotification, ":: vpm :: ")
	out = scribe.NewWithWriter(outW, "")

	configLocation := os.Getenv("VROOMY_CONFIG")
	if len(configLocation) == 0 {
		configLocation = DefaultConfigLocation
	}

	var cmd *flag.Command
	if cmd, err = commandFromArgs(); err != nil {
		help(cmd)
		handleError(err)
	}

	if customCfg := cmd.StringFrom("config"); customCfg != "" {
		configLocation = customCfg
	}

	switch cmd.Action {
	case "help", "version", "upgrade":
		// No config needed
	default:
		// Parse config
		if v.cfg, err = config.NewConfig(configLocation); err != nil {
			handleError(err)
		}
	}

	if err = cmd.Exec(); err != nil {
		handleError(err)
	}
}
