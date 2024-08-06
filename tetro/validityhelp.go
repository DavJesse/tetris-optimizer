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

	return count == 4
}

func validConnections(tet []string) bool {
	var count int
	tag := byte('#')
	row, col := len(tet), len(tet[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if tet[i][j] == tag {
				// Check above
				if i > 0 && tet[i-1][j] == tag {
					count++
				}
				// Check below
				if i < row-1 && tet[i+1][j] == tag {
					count++
				}
				// Check left
				if j > 0 && tet[i][j-1] == tag {
					count++
				}
				// Check right
				if j < col-1 && tet[i][j+1] == tag {
					count++
				}

			}
		}
	}
	// Return true should 'count' be 6 or 8
	return count == 6 || count == 8
}

func replaceHash(index int, tet []string) []string {
	var result []string
	var lnSlc []rune
	rep := 'A' + int32(index) // Establish alphabetic charater to use

	for _, line := range tet {
		lnSlc = []rune(line)
		for i := 0; i < len(lnSlc); i++ {
			// Replace '#' with alphabetic character
			if lnSlc[i] == '#' {
				lnSlc[i] = rep
			}
			// Append adjusted line to result
			if i == len(lnSlc)-1 {
				result = append(result, string(lnSlc))
			}
		}
	}
	return result
}

// Outputs validity error message
func errors() string {
	return "ERROR"
}
