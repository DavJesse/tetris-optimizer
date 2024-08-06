package tetro

import "testing"

func TestCheckValidity(t *testing.T) {
	test := "TestCheckValidity"
	subject := [][]string{
		{"#...", "#...", "#...", "#..."},
		{"....", "....", "##..", "##.."},
	}
	gotTetro, gotErr := CheckValidity(subject) // Retrieve output for comparison
	expectedTetro, exErr := [][]string{
		{"A", "A", "A", "A"},
		{"BB", "BB"},
	}, ""

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
		} else {
			// Compare lines of expected and got tetrominos
			for i := 0; i < len(expectedTetro); i++ {
				for j := 0; j < len(expectedTetro[i]); j++ {
					if gotTetro[i][j] != expectedTetro[i][j] {
						t.Errorf("Got: %v", gotTetro)
						t.Errorf("Expected: %v", expectedTetro)
						t.Errorf("%s Failed!", test)
					}
				}
			}
		}
	}
}
