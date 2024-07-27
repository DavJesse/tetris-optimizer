package tetro

import "testing"

// func TestTetroTrim(t *testing.T) {
// 	test := ""
// 	subject :=
// 	got :=
// 	expected :=

// 	if got != expected {
// 		t.Errorf("Got: %v", got)
// 		t.Errorf("Expected: %v", expected)
// 		t.Errorf("%s Failed!", test)
// 	}
// }

func TestEmptColumn(t *testing.T) {
	test := "TestEmptColumn"
	subject := []string{"#...", "#...", "#...", "#..."}
	got := occupyCol(subject)
	expected := []int{0}

	if len(got) != len(expected) {
		t.Errorf("Got: %v", got)
		t.Errorf("Expected: %v", expected)
		t.Errorf("%s Failed!", test)
		t.Fail()
	} else {
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("Got: %v", got)
				t.Errorf("Expected: %v", expected)
				t.Errorf("%s Failed!", test)
			}
		}

	}

}
