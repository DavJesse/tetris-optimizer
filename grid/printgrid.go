package grid

import (
	"fmt"

	"tetris/misc"
)

func PrintGrid(size int, tetroSlc [][]string) {
	grid := make([][]byte, size)

	for i := range grid {
		grid[i] = make([]byte, size)
	}
	placeTetromino(grid, tetroSlc, 0, size)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				misc.Print(".")
			} else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		misc.PrintLine("")
	}
}
