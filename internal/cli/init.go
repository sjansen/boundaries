package cli

import (
	"os"

	"github.com/sjansen/boundaries/internal/commands"
)

type initCmd struct{}

func (cmd *initCmd) Run() error {
	init := &commands.InitCmd{}
	return init.Run(os.Stdout, os.Stderr)
}
