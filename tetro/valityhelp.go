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
							// Check for block directly above
							count += checkAbove(i, j, count, h, file)
							// 	// Check for block on left
							if file[i][j-1] == h {
								count++
							}
							// Check for block on right
							if file[i][j+1] == h {
								count++
							}
							// Check for block directly below
							if file[i+1][j] == h {
								count++
							}
						}
					}

				}
				if i != len(file)-1 && j != 0 && j != len(file[1])-1 {

				}
			}
		}
	}

	return conect
}

func checkAbove(i, j, count int, h byte, file []string) int {
	if file[i-1][j] == h {
		count++
	}
	return count
}
