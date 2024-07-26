package tetro

import "tetris/strung"

func fourByFour(file [][]string) bool {
	valid := true

	// Determine if each tetromino is stacked four lines high
	for _, height := range file {
		if len(height) != 4 {
			valid = false
			break
		}

		// Determine if each tetromino is four characters wide
		for _, width := range height {
			if len(width) != 4 {
				valid = false
				break
			}
		}

		// In case of invalidity, break for efficiency
		if !valid {
			break
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

func connections(file []string) int {
	var count int
	h := byte('#')

	for i := 0; i < len(file); i++ {
		if strung.Contains(file[i], "#") {
			for j := 0; j < len(file[i]); j++ {
				if i != 0 && i != len(file)-1 {
					if j != 0 && j != len(file[i])-1 {
						if file[i][j] == h {

							count += checkAbove(i, j, count, h, file) // Check for block directly above
							count += checkLeft(i, j, count, h, file)  // Check for block on left
							count += checkRight(i, j, count, h, file) // Check for block on right
							count += checkBelow(i, j, count, h, file) // Check for block directly below

						}
					}

					if j != 0 && file[i][j] == h {
						count += checkAbove(i, j, count, h, file) // Check for block directly above
						count += checkLeft(i, j, count, h, file)  // Check for block on left
						count += checkBelow(i, j, count, h, file) // Check for block directly below
					}

					if j != len(file[i])-1 && file[i][j] == h {
						count += checkAbove(i, j, count, h, file) // Check for block directly above
						count += checkRight(i, j, count, h, file) // Check for block on right
						count += checkBelow(i, j, count, h, file) // Check for block directly below
					}

				}
				if i != len(file)-1 && j != 0 && j != len(file[i])-1 {
					if file[i][j] == h {
						count += checkAbove(i, j, count, h, file) // Check for block directly above
						count += checkLeft(i, j, count, h, file)  // Check for block on left
						count += checkRight(i, j, count, h, file) // Check for block on right

					}
				}
				if i != 0 && j != 0 && j != len(file[1])-1 {
					if file[i][j] == h {
						count += checkLeft(i, j, count, h, file)  // Check for block on left
						count += checkRight(i, j, count, h, file) // Check for block on right
						count += checkBelow(i, j, count, h, file) // Check for block directly below
					}
				}

			}
		}
	}

	return count
}

func checkAbove(i, j, count int, h byte, file []string) int {
	if file[i-1][j] == h {
		count++
	}
	return count
}

func checkLeft(i, j, count int, h byte, file []string) int {
	if file[i][j-1] == h {
		count++
	}
	return count
}

func checkRight(i, j, count int, h byte, file []string) int {
	if file[i][j+1] == h {
		count++
	}
	return count
}

func checkBelow(i, j, count int, h byte, file []string) int {
	if file[i+1][j] == h {
		count++
	}
	return count
}

func TrimTetro(tet []string) []string {
	row := len(tet)
	col := len(tet[0])
	pil := 0
	tag := byte('#')

	for i := 0; i < row; i++ {
		if !strung.Contains(tet[i], "#") {
			if i < row-1 {
				tet = append(tet[:i], tet[i+1:]...)
				row--
			} else {
				tet = tet[:i]
				row--
			}
			
		}
		for j := 0; j < col; j++ {
			for k := 0; k < row; k++ {
				if tet[i+k][j] == tag {
					break
				} else {
					pil++
					if pil = row

				}
			}

		}
	}

}