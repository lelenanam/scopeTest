package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test11" {
		t.Fatal("test failed")
	}
}
