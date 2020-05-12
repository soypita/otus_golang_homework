package main

import (
	"log"
	"os"
)

const minimalArgLen = 2

func main() {
	// Get args from command line
	inputArgs := os.Args

	if len(inputArgs) <= minimalArgLen {
		log.Fatal(`Wrong params length. Usage: go-envdir /path/to/env/dir command arg1 arg2`)
	}

	pathToEnvDir := inputArgs[1]
	commands := inputArgs[2:]

	enviromentParams, err := ReadDir(pathToEnvDir)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(RunCmd(commands, enviromentParams))
}
