package cli

import "github.com/sjansen/boundaries/internal/commands"

func registerInit(p *ArgParser) {
	c := &commands.InitCmd{}
	p.addCommand(c, "init", "Create a minimal project config or reinitialize an existing one")
}
