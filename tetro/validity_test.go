package tetro

import "testing"

func TestCheckValidity(t *testing.T) {
	test := "TestCheckValidity"
	subject := [][]string{
		{"#...", "#..", "#...", "#..."},
		{"#...", "....", "#...", "#..."},
	}
	gotTetro, gotErr := CheckValidy(subject) // Retrieve output for comparison
	expectedTetro, exErr := [][]string{}, " contains invalid tetromino"

	// Compare error lengths, Failing in case of descrepancy
	if len(gotErr) != len(exErr) {
		t.Errorf("Got: %v", gotErr)
		t.Errorf("Expected: %v", exErr)
		t.Errorf("%s Failed!", test)
		t.Fail()
	} else {
		// Compare lengths, Failing in case of descrepancy
		if len(gotTetro) != len(expectedTetro) {
			t.Errorf("Got: %v", gotTetro)
			t.Errorf("Expected: %v", expectedTetro)
			t.Errorf("%s Failed!", test)
		}
	}
}
