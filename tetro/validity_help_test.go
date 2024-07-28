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

func TestFourChars(t *testing.T) {
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

// func Test(t *testing.T) {
// 	test := "Test"
// 	subject :=
// 	got :=  // Retrieve output for comparison
// 	expected :=

// 	if got != expected {
// 		t.Errorf("Got: %v", got)
// 		t.Errorf("Expected: %v", expected)
// 		t.Errorf("%s Failed!", test)
// 	}
// }
