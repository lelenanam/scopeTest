package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test 3" {
		t.Fatal("test failed")
	}
}
