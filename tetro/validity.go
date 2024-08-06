package tetro

// Check each tetromino for validity
func CheckValidy(file [][]string) ([][]string, string) {
	var err string
	var trimmed []string
	var tetro [][]string

	if len(file) > 26 {
		err = errors()
	} else {
		// Range through file testing individual tetrominos
		for i, tet := range file {
			// Check dimensions of tetromino
			if !fourByFour(tet) {
				err = errors()
				tetro = [][]string{}
				break
			}

			// Trim tetromino for efficiency
			trimmed = TetroTrim(tet)

			// Check if each tetromino has four '#'
			if !fourHashes(trimmed) {
				err = errors()
				tetro = [][]string{}
				break
			}

			// Check for valid number of connections
			if !validConnections(trimmed) {
				err = errors()
				tetro = [][]string{}
			}

			trimmed = replaceHash(i, trimmed) // Replace '#' with unique alphabetic letter

			tetro = append(tetro, trimmed) // Trim and append valid tetominos to tetro

		}
	}

	return tetro, err
}
