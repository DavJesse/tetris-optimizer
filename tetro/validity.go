package tetro

func CheckValidy(file [][]string) ([][]string, string) {
	var err string
	var tetro [][]string

	for _, tet := range file {
		if !fourByFour(tet) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		if !fourHashes(tet) {
			err = validityError()
			tetro = [][]string{}
			break
		}

		tetro = append(tetro, TetroTrim(tet))

	}
	return tetro, err
}
