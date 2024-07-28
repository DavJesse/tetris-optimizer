package grid

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
