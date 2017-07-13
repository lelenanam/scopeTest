package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test123" {
		t.Fatal("test failed")
	}
}
