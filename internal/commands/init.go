package commands

import (
	"io"
	"os"
	"path/filepath"
)

const template = `[core]
config_version = 0
`

type InitCmd struct{}

func (c *InitCmd) Run(stdout, stderr io.Writer) error {
	if err := os.MkdirAll(".boundaries", os.ModePerm); err != nil {
		return err
	}

	filename := filepath.Join(".boundaries", "config")
	if err := os.WriteFile(filename, []byte(template), os.ModePerm); err != nil {
		return err
	}

	return nil
}
