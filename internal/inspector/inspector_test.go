package inspector_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/boundaries/internal/inspector"
)

func TestInspector(t *testing.T) {
	require := require.New(t)

	data, err := os.ReadFile("testdata/files.json")
	require.NoError(err)

	files := map[string]*inspector.File{}
	err = json.Unmarshal(data, &files)
	require.NoError(err)

	ins, err := inspector.New("testdata")
	require.NoError(err)

	os.Chdir("testdata")
	err = filepath.Walk(".", func(filename string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case strings.HasSuffix(filename, "testdata"):
			return filepath.SkipDir
		case strings.HasSuffix(filename, "vendor"):
			return filepath.SkipDir
		case info.IsDir():
			return nil
		case !strings.HasSuffix(filename, ".go"):
			return nil
		}

		expected, ok := files[filename]
		if !ok {
			require.Failf("unexpected file", "filename=%q", filename)
		}
		delete(files, filename)

		actual, err := ins.Inspect(filename)
		if expected == nil {
			require.Error(err, filename)
			require.Nil(actual, filename)
		} else {
			require.NoError(err, filename)
			require.Equal(expected, actual, filename)
		}

		return nil
	})
	require.NoError(err)
	require.Empty(files)
}
