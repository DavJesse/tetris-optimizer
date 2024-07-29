package grid

func Solve(tetrominoes [][]string) int {
	for size := 2; ; size++ {
		grid := make([][]byte, size)
		for i := range grid {
			grid[i] = make([]byte, size)
		}

		if placeTetromino(grid, tetrominoes, 0, size) {
			return size
		}
	}
}
