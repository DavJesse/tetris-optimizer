package grid

func placeTetromino(grid [][]byte, tetrominoes [][]string, index, size int) bool {
	if index == len(tetrominoes) {
		return true
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if canPlace(grid, tetrominoes[index], i, j, size) {
				place(grid, tetrominoes[index], i, j)
				if placeTetromino(grid, tetrominoes, index+1, size) {
					return true
				}
				remove(grid, tetrominoes[index], i, j)
			}
		}
	}

	return false
}

func canPlace(grid [][]byte, tetromino []string, row, col, size int) bool {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] != '.' {
				if row+i >= size || col+j >= size || grid[row+i][col+j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func place(grid [][]byte, tetromino []string, row, col int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] != '.' {
				grid[row+i][col+j] = tetromino[i][j]
			}
		}
	}
}

func remove(grid [][]byte, tetromino []string, row, col int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] != '.' {
				grid[row+i][col+j] = 0
			}
		}
	}
}
