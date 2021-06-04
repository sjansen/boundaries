package cli

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Help    helpCmd    `kong:"cmd,help='Show this help text'"`
	Init    initCmd    `kong:"cmd,help='Create a minimal project config or reinitialize an existing one'"`
	Version versionCmd `kong:"cmd,help='Show the current boundaries version'"`
}

// ParseAndRun parses command line arguments then runs the matching command.
func ParseAndRun(version string) {
	ctx := kong.Parse(&cli,
		kong.Description("Check and enforce code organization"),
		kong.UsageOnError(),
	)
	cli.Help.ctx = ctx
	cli.Version.version = version

	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
