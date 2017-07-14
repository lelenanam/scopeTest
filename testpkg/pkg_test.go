package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test12345" {
		t.Fatal("test failed")
	}
}
