package grid

import (
	"bytes"
	"testing"
)

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
	got := placeTetromino(grid, tetroSlc, index, size) // Retrieve outputs for comparison
	expected := true

	if got != expected {
		t.Errorf("Got: %v", got)
		t.Errorf("Expected: %v", expected)
		t.Errorf("%s Failed!", test)
	}
}

func TestCanPlace(t *testing.T) {
	test := "TestCanPlace"
	grid := make([][]byte, 4)
	for i := range grid {
		grid[i] = make([]byte, 4)
	}
	tetromino := []string{
		"A",
		"A",
		"A",
		"A",
	}
	row, col, size := 0, 0, 4
	got := canPlace(grid, tetromino, row, col, size) // Retrieve outputs for comparison
	expected := true

	if got != expected {
		t.Errorf("Got: %v", got)
		t.Errorf("Expected: %v", expected)
		t.Errorf("%s Failed!", test)
	}
}
