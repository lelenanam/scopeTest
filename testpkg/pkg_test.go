package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test xx" {
		t.Fatal("test failed")
	}
}
