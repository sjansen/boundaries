package inspector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTagsFromFilename(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Filename string
		Expected string
	}{{
		Filename: "README.md",
	}, {
		Filename: "foo.go",
	}, {
		Filename: "foo_arm64.go",
		Expected: "arm64",
	}, {
		Filename: "foo_android.go",
		Expected: "android",
	}, {
		Filename: "foo_darwin_amd64.go",
		Expected: "amd64,darwin",
	}, {
		Filename: "foo_test.go",
	}, {
		Filename: "foo_linux_test.go",
		Expected: "linux",
	}, {
		Filename: "foo_linux_test.go",
		Expected: "linux",
	}, {
		Filename: "foo_linux_ppc64_test.go",
		Expected: "ppc64,linux",
	}} {
		actual := buildTagsFromFilename(tc.Filename)
		require.Equal(tc.Expected, actual, tc.Filename)
	}
}
