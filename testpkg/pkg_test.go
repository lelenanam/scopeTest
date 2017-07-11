package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test 111" {
		t.Fatal("test failed")
	}
}
