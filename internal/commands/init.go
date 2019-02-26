package commands

import (
	"io"
	"io/ioutil"
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
	if err := ioutil.WriteFile(filename, []byte(template), os.ModePerm); err != nil {
		return err
	}

	return nil
}
