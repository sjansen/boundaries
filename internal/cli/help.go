package cli

import (
	"github.com/alecthomas/kong"
)

type helpCmd struct {
	ctx *kong.Context
}

func (cmd *helpCmd) Run() error {
	cmd.ctx.Args = []string{}
	cmd.ctx.Path = []*kong.Path{}
	_ = cmd.ctx.PrintUsage(false)
	return nil
}
