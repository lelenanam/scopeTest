package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test" {
		t.Fatal("test failed")
	}
}
