package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test 1" {
		t.Fatal("test failed")
	}
}
