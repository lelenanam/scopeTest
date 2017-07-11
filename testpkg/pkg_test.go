package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test 11" {
		t.Fatal("test failed")
	}
}
