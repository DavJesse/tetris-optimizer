package tetro

import "testing"

func TestFourByFour(t *testing.T) {
	subject := []string{"#...", "#..", "#...", "#..."}
	got := fourByFour(subject)
	expected := false

	// Compare got and expected
	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestFourByFour Failed!")
	}
}

func TestFourHashes(t *testing.T) {
	subject := []string{"#...", "....", "#...", "#..."}
	got := fourHashes(subject)
	expected := false

	// Compare got and expected
	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestFourChars Failed!")
	}
}

func TestValidConnections(t *testing.T) {
	subject := []string{
		"....",
		"....",
		"..##",
		"..##",
	}
	got := validConnections(subject)
	expected := true

	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestConnection Failed!")
	}
}

func TestReplaceHash(t *testing.T) {
	test := "TestReplaceHash"
	subject := []string{"#...", "#...", "#...", "#..."}
	got := replaceHash(25, subject) // Retrieve output for comparison
	expected := []string{"Z...", "Z...", "Z...", "Z..."}

	// Compare lengths, Failing in case of descrepancy
	if len(got) != len(expected) {
		t.Errorf("Got: %v", got)
		t.Errorf("Expected: %v", expected)
		t.Errorf("%s Failed!", test)
		t.Fail()
	} else {
		// Compare elements, Failing in case of descrepancy
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("Got: %v", got)
				t.Errorf("Expected: %v", expected)
				t.Errorf("%s Failed!", test)
			}
		}
	}
}
