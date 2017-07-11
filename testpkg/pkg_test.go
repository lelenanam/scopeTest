package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test x" {
		t.Fatal("test failed")
	}
}
