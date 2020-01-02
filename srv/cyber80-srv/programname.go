package main

import (
	"path/filepath"
	"os"
)

var (
	programName string = "<([unknown-program-name])>"
)

func init() {
	if 1 > len(os.Args) {
		return
	}

	programName = filepath.Base(os.Args[0])
}
