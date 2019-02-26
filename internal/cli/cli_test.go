package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/boundaries/internal/cli"
	"github.com/sjansen/boundaries/internal/commands"
)

func TestArgParser(t *testing.T) {
	require := require.New(t)

	parser := cli.RegisterCommands("test")
	for _, tc := range []struct {
		args        []string
		expected    cli.Command
		expectError bool
	}{{
		args:     []string{"init"},
		expected: &commands.InitCmd{},
	}, {
		args: []string{"version"},
		expected: &commands.VersionCmd{
			App:   "boundaries",
			Build: "test",
		},
	}} {
		actual, err := parser.Parse(tc.args)
		if tc.expectError {
			require.Error(err)
		} else {
			require.NoError(err)
			require.Equal(tc.expected, actual)
		}
	}
}
