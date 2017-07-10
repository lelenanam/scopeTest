package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test10" {
		t.Fatal("test failed")
	}
}
