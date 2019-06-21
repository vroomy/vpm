package main

import (
	"os"

	"github.com/hatchify/output"
)

func handleError(err error) {
	output.Error(err.Error())
	os.Exit(1)
}
