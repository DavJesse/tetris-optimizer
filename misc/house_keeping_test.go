package misc

import "testing"

func TestCheckExtension(t *testing.T) {
	subject := "test.txt"
	found := CheckExtension(subject)
	expected := true

	if found != expected {
		t.Errorf("Found: %t", found)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestCheckExtension Failed!")
	}
}
