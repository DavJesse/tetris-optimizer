package tetro

// Check each tetromino for validity
func CheckValidy(file [][]string) ([][]string, string) {
	var err string
	var trimmed []string
	var tetro [][]string

	// Range through file testing individual tetrominos
	for _, tet := range file {
		// Check dimensions of tetromino
		if !fourByFour(tet) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		trimmed = TetroTrim(tet)

		// Check if each tetromino has four '#'
		if !fourHashes(trimmed) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		// Check for valid number of connections
		if !validConnections(trimmed) {
			err = validityError()
			tetro = [][]string{}
		}

		// Trim and append valid tetominos to tetro
		tetro = append(tetro, TetroTrim(tet))

	}
	return tetro, err
}
