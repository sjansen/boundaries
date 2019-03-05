package c_test

import (
	"testing"

	"example.com/foo/pkg/c"
)

func TestValue(t *testing.T) {
	expected := "c"
	actual := c.Value()
	if actual != expected {
		t.Fail()
	}
}
