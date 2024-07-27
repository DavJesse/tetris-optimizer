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

func fourChars(file [][]string) bool {
	var count int
	valid := true

	// i ranges though tetrominoes in file
	// j ranges though lines of tetromino
	// k ranges through character in each line
	for i := 0; i < len(file); i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {

				// Count number of '#' in tetromino
				if string(file[i][j][k]) == "#" {
					count++
				}

				// Check for irregular '#' count at end of tetromino
				if j == 3 && k == 3 {
					if count == 4 {
						count = 0 // Reset coint for next round of tetromino
					} else {
						valid = false // Set valid to false at irregular count
						break
					}
				}

			}
			// When invalid tetromino is discovered...
			// Break out of loop for efficiency
			if !valid {
				break
			}
		}
		// Break for efficiency
		if !valid {
			break
		}
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
