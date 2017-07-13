package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "test1234" {
		t.Fatal("test failed")
	}
}
