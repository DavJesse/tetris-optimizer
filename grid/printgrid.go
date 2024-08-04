package grid

import (
	"tetris/misc"
)

// Prints solved grid
func PrintGrid(size int, tetroSlc [][]string) {
	// Create empty grid of specified size
	grid := make([][]byte, size)

	for i := range grid {
		grid[i] = make([]byte, size)
	}
	placeTetromino(grid, tetroSlc, 0, size) // Place tetrominos

	// Print grid
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				misc.Print(".") // Replace blanks with '.'s
			} else {
				misc.Print(string(grid[i][j])) // Print alphabetic character
			}
		}
		misc.PrintLine("") // Print newline to move on to next line
	}
}
