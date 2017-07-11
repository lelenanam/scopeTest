package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test7" {
		t.Fatal("test failed")
	}
}
