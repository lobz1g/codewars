package src

func IsSolved(board [3][3]int) int {
	var isNotFinished bool
	var winner int
	for i := 0; i < len(board); i++ {
		switch {

		case board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != 0:
			winner = board[i][0]

		case board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != 0:
			winner = board[0][i]

		case board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[1][1] != 0,
			board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[1][1] != 0:
			winner = board[1][1]

		case board[i][0] == 0,
			board[i][1] == 0,
			board[i][2] == 0,
			board[0][i] == 0,
			board[1][i] == 0,
			board[2][i] == 0:
			isNotFinished = true
		}

		if winner != 0 {
			break
		}
	}

	switch {
	case winner != 0:
		return winner
	case isNotFinished:
		return -1
	default:
		return 0
	}
}
