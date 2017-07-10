package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test 2" {
		t.Fatal("test failed")
	}
}
