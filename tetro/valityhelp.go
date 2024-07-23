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
	}
	return valid
}

func fourChars(file [][]string) bool {
	var count int
	valid := true

	for _, tet := range file {
		for i := 0; i < 4; i++ {
			if string(tet[i]) == "#" {
				count++
				if i == 3 && count != 4 {
					valid = false
					break
				}
			}
		}
	}
	return valid
}
