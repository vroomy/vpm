package main

import (
	"os"
)

func handleError(err error) {
	out.Error(err.Error())
	os.Exit(1)
}
