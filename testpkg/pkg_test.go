package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test xxx 2" {
		t.Fatal("test failed")
	}
}
