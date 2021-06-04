package cli

import "fmt"

type versionCmd struct {
	version string
}

func (cmd *versionCmd) Run() error {
	fmt.Println(cmd.version)
	return nil
}
