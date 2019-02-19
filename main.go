package main

import (
	"fmt"
	"os"

	"github.com/sjansen/boundaries/internal/cli"
)

var build string // set by goreleaser

func main() {
	if build == "" {
		build = version
	}
	parser := cli.RegisterCommands(build)

	cmd, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cmd.Run(os.Stdout, os.Stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
