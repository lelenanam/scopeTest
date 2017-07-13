package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test12" {
		t.Fatal("test failed")
	}
}
