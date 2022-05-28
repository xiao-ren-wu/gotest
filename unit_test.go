package gotest

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 1
	expected := 2
	actual := Add(a, b)

	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}
