package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test1" {
		t.Fatal("test failed")
	}
}
