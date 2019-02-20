package b

import "testing"

func TestValue(t *testing.T) {
	actual := Value()
	if actual != "b" {
		t.Fail()
	}
}
