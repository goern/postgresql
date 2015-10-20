package main

import (
	"os"

	"github.com/docker/docker/builder/parser"
)

func main() {
	var dockerfile *os.File
	var err error

	// lets get the Dockerfile
	dockerfile, err = os.Open("Dockerfile.rhel7")
	if err != nil {
		panic(err)
	}

	// and parse the Dockerfile
	ast, err := parser.Parse(dockerfile)
	if err != nil {
		panic(err)
	} else { // and figure out all the LABELs that we need

	}
}
