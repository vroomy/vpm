package main

import (
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	flag "github.com/hatchify/parg"
	"github.com/hatchify/queue"
	"github.com/hatchify/scribe"
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
	configLocation := os.Getenv("VROOMY_CONFIG")
	if len(configLocation) == 0 {
		configLocation = DefaultConfigLocation
	}

	outW = scribe.NewStdout()
	outW.SetTypePrefix(scribe.TypeNotification, ":: vpm :: ")
	out = scribe.NewWithWriter(outW, "")
	out.Notification("Vroomy Package Manager ::")

	var err error
	if _, err = toml.DecodeFile(configLocation, &v.cfg); err != nil {
		handleError(err)
	}

	var cmd *flag.Command
	if cmd, err = commandFromArgs(); err != nil {
		help(cmd)
		handleError(err)
	}

	if err = cmd.Exec(); err != nil {
		handleError(err)
	}
}
