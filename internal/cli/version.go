package cli

import "github.com/sjansen/boundaries/internal/commands"

func registerVersion(p *ArgParser, build string) {
	c := &commands.VersionCmd{
		App:   "boundaries",
		Build: build,
	}
	p.addCommand(c, "version", "Print boundaries's version")
}
