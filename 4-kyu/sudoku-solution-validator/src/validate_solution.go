package src

func ValidateSolution(mat [][]int) bool {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 || !valid(mat, i, j) {
				return false
			}
		}
	}

	return true
}

func valid(mat [][]int, x, y int) bool {
	for i := 0; i < 9; i++ {
		if i != x && mat[i][y] == mat[x][y] {
			return false
		}

		if i != y && mat[x][i] == mat[x][y] {
			return false
		}

		xx := (x/3)*3 + i/3
		yy := (y/3)*3 + i%3
		if xx != x && y != yy && mat[xx][yy] == mat[x][y] {
			return false
		}
	}

	return true
}
