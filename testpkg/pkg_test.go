package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test xxx" {
		t.Fatal("test failed")
	}
}
