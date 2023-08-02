package leetcode

func gameOfLife(board [][]int) {

	tmpBoard := make([][]int, len(board))
	for i := range board {
		tmpBoard[i] = make([]int, len(board[i]))
		copy(tmpBoard[i], board[i])
	}

	for i, a := range board {

		for j := range a {
			cell := isAlive(tmpBoard, i, j)
			board[i][j] = cell
		}

	}
}

func isAlive(board [][]int, row int, col int) int {

	value := board[row][col]

	dead := 0
	alive := 0

	// check vertical

	// up
	if (row - 1) > -1 {
		if board[row-1][col] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	// down
	if (row + 1) < len(board) {
		if board[row+1][col] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	// check horizontal

	// left
	if (col - 1) > -1 {
		if board[row][col-1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	// right
	if (col + 1) < len(board[row]) {
		if board[row][col+1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	// check diagonals
	if (row-1) > -1 && (col-1) > -1 {
		if board[row-1][col-1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	if (row-1) > -1 && (col+1) < len(board[row]) {
		if board[row-1][col+1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	if (row+1) < len(board) && (col+1) < len(board[row]) {
		if board[row+1][col+1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	if (row+1) < len(board) && (col-1) > -1 {
		if (board)[row+1][col-1] == 0 {
			dead += 1
		} else {
			alive += 1
		}
	}

	// check if cell state

	if value == 1 && alive < 2 {
		return 0
	} else if value == 1 && alive == 2|3 {
		return 1
	} else if value == 1 && alive > 3 {
		return 0
	} else if value == 0 && alive == 3 {
		return 1
	}

	return value
}
