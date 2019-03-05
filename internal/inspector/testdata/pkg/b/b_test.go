package b_test

import (
	"testing"

	"example.com/foo/pkg/b"
)

func TestValue(t *testing.T) {
	actual := b.Value()
	if actual != "b" {
		t.Fail()
	}
}
