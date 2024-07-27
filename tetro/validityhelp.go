package tetro

func fourByFour(tetro []string) bool {
	valid := true

	// Determine if tetromino is stacked four lines high
	if len(tetro) != 4 {
		valid = false
	} else {
		// Determine if each line has four characters
		for _, line := range tetro {
			if len(line) != 4 {
				valid = false
				break
			}
		}
	}
	return valid
}

func fourHashes(tetro []string) bool {
	var count int
	valid := true
	row := len(tetro)
	col := len(tetro[0])

	// j ranges though lines of tetromino
	// k ranges through character in each line
	for j := 0; j < row; j++ {
		for k := 0; k < col; k++ {
			// Count number of '#' in tetromino
			if string(tetro[j][k]) == "#" {
				count++
			}
		}
	}

	if count != 4 {
		valid = false
	}

	return valid
}

// func connections(file []string) int {
// 	var count int
// 	h := byte('#')

// 	return count
// }

// func checkAbove(i, j, count int, h byte, file []string) int {
// 	if file[i-1][j] == h {
// 		count++
// 	}
// 	return count
// }

// func checkLeft(i, j, count int, h byte, file []string) int {
// 	if file[i][j-1] == h {
// 		count++
// 	}
// 	return count
// }

// func checkRight(i, j, count int, h byte, file []string) int {
// 	if file[i][j+1] == h {
// 		count++
// 	}
// 	return count
// }

// func checkBelow(i, j, count int, h byte, file []string) int {
// 	if file[i+1][j] == h {
// 		count++
// 	}
// 	return count
// }

// Outputs validity error message
func validityError() string {
	err := " contains invalid tetromino"
	return err
}
