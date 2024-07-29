package grid

func Solve(tetroSlc [][]string) int {

	// initalize the smallest possible grid,
	// placing and backtracking tetrominoes,
	// and incrimentally increasing grid size
	for size := 2; ; size++ {
		grid := make([][]byte, size) // Create outer slice

		for i := range grid {
			grid[i] = make([]byte, size) // Adjust size of inner slice
		}

		if placeTetromino(grid, tetroSlc, 0, size) {
			return size
		}
	}
}
