package tetro

func TetroTrim(tet []string) []string {

}

func emptColumn(tet []string) []int {
	var index []int
	empt := true
	tag := byte('#')
	row := len(tet)
	col := len(tet[0])

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if tet[j][i] == tag {
				empt = false
				break
			} else if empt == true && j == row-1 {
				index = append(index, j)
			}
		}
	}
	return index
}