package tetro

import "testing"

// func Test(t *testing.T) {
// 	test := "Test"
// 	subject :=
// 	got := // Retrieve output for comparison
// 	expected :=

// 	if got != expected {
// 		t.Errorf("Got: %v", got)
// 		t.Errorf("Expected: %v", expected)
// 		t.Errorf("%s Failed!", test)
// 	}
// }

func TestTetroTrim(t *testing.T) {
	test := "TestTetroTrim"
	subject := []string{"....", "....", "..##", "..##"}
	got := TetroTrim(subject) // Retrieve output for comparison
	expected := []string{"##", "##"}

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

func TestEmptColumn(t *testing.T) {
	test := "TestEmptColumn"
	subject := []string{"#...", "#...", "#...", "#..."}
	got := occupyCol(subject) // Retrieve output for comparison
	expected := []int{0}

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

func TestRemoveRow(t *testing.T) {
	test := "TestRemoveRow"
	subject := []string{"....", "....", "..##", "..##"}
	got := removeRow(subject) // Retrieve output for comparison
	expected := []string{"..##", "..##"}

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
