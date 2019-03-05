// +build integration

package c_test

import (
	"io/ioutil"
	"os"
	"testing"

	"example.com/foo/pkg/c"
)

func TestBuildTags(t *testing.T) {
	r, err := os.Open("testdata/expected.txt")
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	expected := string(b)
	actual := c.Value()
	if actual != expected {
		t.Fail()
	}
}
