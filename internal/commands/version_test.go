package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	require := require.New(t)

	expected := "pgmagic 1.0-test\n"
	cmd := &VersionCmd{
		App:   "pgmagic",
		Build: "1.0-test",
	}

	var stdout, stderr bytes.Buffer
	err := cmd.Run(&stdout, &stderr)
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())
}
