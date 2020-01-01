package main

import (
	"log"
	"os"

	version "distance_calc"
	"github.com/mitchellh/cli"
)

var errPrefix = "\x1b[31merror\x1b[0m"

func main() {
	c := cli.NewCLI("timeline", version.Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"distance": func() (cli.Command, error) {
			return &distance{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Printf(errPrefix+" %#v", err)
	}

	os.Exit(exitStatus)
}
