package misc

import (
	"os"
	"testing"

	"tetris/strung"
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

func TestSplitString(t *testing.T) {
	subject := "abcabcabc"
	found := strung.Split(subject, "abc")
	expected := []string{"", "", "", ""}

	// Check if lengths of found and expected are equal...
	// ...or if their elements match
	for i := 0; i < len(expected); i++ {
		if len(found) != len(expected) || found[i] != expected[i] {
			// When lengths are different, print the following
			if len(found) != len(expected) {
				t.Errorf("Found has %d elements", len(found))
				t.Errorf("Expected has %d elements", len(expected))
				t.Errorf("TestSplitString Failed!")
				t.FailNow()

				// When elements don't match, print the following
			} else {
				t.Errorf("Found: %s", found)
				t.Errorf("Expected: %s", expected)
				t.Errorf("TestSplitString Failed!")

			}
		}
	}
}

func TestTwoD(t *testing.T) {
	subject, _ := os.ReadFile("test.txt")
	found := TwoD(string(subject))
	expected := [][]string{
		{"#...", "#...", "#...", "#..."},
		{"....", "....", "..##", "..##"},
	}

	for i := 0; i < len(expected); i++ {
		for j := 0; j < 4; j++ {
			if len(found) != len(expected) || found[i][j] != expected[i][j] {
				// When lengths are different, print the following
				if len(found) != len(expected) {
					t.Errorf("Found has %d elements", len(found))
					t.Errorf("Expected has %d elements", len(expected))
					t.Errorf("TestTwoD Failed!")
					t.FailNow()

					// When elements don't match, print the following
				} else {
					t.Errorf("Found: %s", found)
					t.Errorf("Expected: %s", expected)
					t.Errorf("TestTwoD Failed!")

				}
			}
		}
	}
}
