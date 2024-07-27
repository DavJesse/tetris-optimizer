package tetro

// Check each tetromino for validity
func CheckValidy(file [][]string) ([][]string, string) {
	var err string
	var tetro [][]string

	// Range through file testing individual tetrominos
	for _, tet := range file {
		// Check dimensions of tetromino
		if !fourByFour(tet) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		// Check if each tetromino has four '#'
		if !fourHashes(tet) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		// Trim and append valid tetominos to tetro
		tetro = append(tetro, TetroTrim(tet))

	}
	return tetro, err
}
