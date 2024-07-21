package misc

import (
	"os"
	"testing"
)

func TestCheckExtension(t *testing.T) {
	subject := "test.txt"
	found := CheckExtension(subject) // Save function output to found
	expected := true

	// Compare found & Expected
	// Print error messages incase of inconsistencies
	if found != expected {
		t.Errorf("Found: %t", found)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestCheckExtension Failed!")
	}
}

func TestReadFile(t *testing.T) {
	subject := "test.txt"
	found := ReadFile(subject)             // Save function output to found
	expected, _ := os.ReadFile("test.txt") // Read file directly from os.ReadFile

	// Compare found & Expected
	// Print error messages incase of inconsistencies
	if found != string(expected) {
		t.Errorf("Found: %s", found)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestReadFile Failed!")
	}
}
