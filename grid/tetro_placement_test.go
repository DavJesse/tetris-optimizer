package grid

import "testing"

func TestPlaceTetromino(t *testing.T) {
		test := "TestPlaceTetromino"
		grid := make([][]byte, 4)
		for i := range grid {
			grid[i] = make([]byte, 4)
		}
		subject :=
		got :=  // Retrieve output for comparison
		expected :=
	
		if got != expected {
			t.Errorf("Got: %v", got)
			t.Errorf("Expected: %v", expected)
			t.Errorf("%s Failed!", test)
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
