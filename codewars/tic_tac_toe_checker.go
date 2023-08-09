package codewars

func IsSolved(board [3][3]int) int {

	isFinished := true

	cnt := 0

	for cnt < 3 {

		elem := board[cnt][cnt]

		if elem == 0 {
			isFinished = false
			cnt++
			continue
		}

		var n1, n2 int

		if cnt == 0 {

			n1, n2 = board[cnt][cnt+1], board[cnt][cnt+2]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}

			n1, n2 = board[cnt+1][cnt], board[cnt+2][cnt]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}

			n1, n2 = board[cnt+1][cnt+1], board[cnt+2][cnt+2]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}
		}

		if cnt == 1 {

			n1, n2 = board[cnt-1][cnt], board[cnt+1][cnt]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}

			n1, n2 = board[cnt][cnt-1], board[cnt][cnt+1]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}

			n1, n2 = board[cnt-1][cnt+1], board[cnt+1][cnt-1]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}
		}

		if cnt == 2 {
			n1, n2 = board[cnt][cnt-1], board[cnt][cnt-2]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}

			n1, n2 = board[cnt-1][cnt], board[cnt-2][cnt]

			if n1 == 0 || n2 == 0 {
				isFinished = false
			}

			if elem == n1 && elem == n2 {
				return elem
			}
		}

		cnt++
	}

	if isFinished {
		return 0
	}

	return -1
}
