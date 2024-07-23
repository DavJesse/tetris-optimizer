package tetro

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
