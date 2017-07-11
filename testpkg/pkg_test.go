package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test0" {
		t.Fatal("test failed")
	}
}
