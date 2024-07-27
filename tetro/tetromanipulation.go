package tetro

import "tetris/strung"

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
	for i, row := range tet {
		if !strung.Contains(row, "#") {
			if i == 0 {
				tet = tet[1:]
				i++
			} else if i == len(tet)-1 {
				tet = tet[:i]
				break
			} else {
				tet = append(tet[:i], tet[i+1:]...)
			}
		}
	}

	for i := 0; i < len(tet); i++ {
		if !strung.Contains(tet[i], "#") {
			tet = removeRow(tet)
		}
	}
	return tet
}