package grid

import "testing"

func TestPlaceTetromino(t *testing.T) {
	test := "TestPlaceTetromino"
	grid := make([][]byte, 4)
	for i := range grid {
		grid[i] = make([]byte, 4)
	}
	tetroSlc := [][]string{
		{
			"A",
			"A",
			"A",
			"A",
		},
	}
	index, size := 0, 4
	got := placeTetromino(grid, tetroSlc, index, size) // Retrieve output for comparison
	expected := true

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
