package tetro

import (
	"tetris/strung"
)

func TetroTrim(tet []string) []string {
	var trimmed []string
	var line string
	index := occupyCol(tet)
	row := len(tet)

	for i := 0; i < row; i++ {
		for j, ind := range index {
			if j != len(index)-1 {
				line += string(tet[i][ind])
			} else {
				line += string(tet[i][ind])
				trimmed = append(trimmed, line)
				line = ""
			}
		}
	}

	trimmed = removeRow(trimmed)

	return trimmed
}

func occupyCol(tet []string) []int {
	var index []int
	// empt := false
	tag := byte('#')
	row, col := len(tet), len(tet[0])

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if tet[j][i] == tag {
				// empt = true
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
