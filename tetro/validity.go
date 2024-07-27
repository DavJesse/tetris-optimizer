package tetro

func CheckValidy(file [][]string) ([][]string, string) {
	var err string
	var tetro []string

	for _, tet := range file {
		if !fourByFour(tet) {
			err = validityError()
			break
		}

		tetro = TetroTrim(tet)

		if !fourHashes(tet) {
			err = validityError()
			break
		}

	}
	return tetro, err
}
