package tetro

import "tetris/strung"

// Check each tetromino for validity
func CheckValidity(file [][]string) ([][]string, string) {
	var err string
	var trimmed []string
	var tetro [][]string

	if len(file) > 26 {
		err = Errors()
	} else {
		// Range through file testing individual tetrominos
		for i, tet := range file {
			// Ensure file only contains '#' and '.' characters
			if !strung.IsHashDot(tet) {
				err = Errors()
				tetro = [][]string{}
				break
			}

			// Check dimensions of tetromino
			if !fourByFour(tet) {
				err = Errors()
				tetro = [][]string{}
				break
			}

			// Check if each tetromino has four '#'
			if !fourHashes(tet) {
				err = Errors()
				tetro = [][]string{}
				break
			}

			// Check for valid number of connections
			if !validConnections(tet) {
				err = Errors()
				tetro = [][]string{}
			}

			trimmed = TetroTrim(tet) // Trim tetromino for efficiency

			trimmed = replaceHash(i, trimmed) // Replace '#' with unique alphabetic letter

			tetro = append(tetro, trimmed) // Trim and append valid tetominos to tetro

		}
	}

	return tetro, err
}
