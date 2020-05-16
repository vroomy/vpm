package main

import (
	"os"
	"strings"
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
