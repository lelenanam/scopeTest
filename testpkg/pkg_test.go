package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test101" {
		t.Fatal("test failed")
	}
}
