package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test2" {
		t.Fatal("test failed")
	}
}
