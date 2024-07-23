package tetro

import "testing"

func TestFourByFour(t *testing.T) {
	subject := [][]string{
		{"#...", "#..", "#...", "#..."},
		{"....", "....", "..##", "..##"},
	}
	got := fourByFour(subject)
	expected := false

	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestFourByFour Failed!")
	}
}

func TestFourChars(t *testing.T) {
	subject
}
