package misc

import (
	"os"
	"testing"
)

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

func TestReadFile(t *testing.T) {
	subject := "test.txt"
	found := ReadFile(subject)
	expected, _ := os.ReadFile("test.txt")

	if found != string(expected) {
		t.Errorf("Found: %s", found)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestReadFile Failed!")
	}
}
