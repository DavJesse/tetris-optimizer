package tetro

import (
	"tetris/strung"
)

// func TetroTrim(tet []string) []string {

// }

func occupyCol(tet []string) []int {
	var index []int
	//empt := false
	tag := byte('#')
	row, col := len(tet), len(tet[0])

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if tet[j][i] == tag {
				//empt = true
				index = append(index, i)
				break
			}
		}
	}
	return index
}

func removeRow(tet []string) []string {
	var result []string
	for _, row := range tet {
		if strung.Contains(row, "#") {
			result = append(result, row)
		}
	}
	return result
}
