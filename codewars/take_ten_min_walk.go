package codewars

func IsValidWalk(walk []rune) bool {

	if len(walk) != 10 {
		return false
	}

	coordX, coordY := 0, 0

	for _, w := range walk {

		switch w {
		case 'n':
			coordY += 1
		case 's':
			coordY -= 1
		case 'e':
			coordX += 1
		case 'w':
			coordX -= 1
		}

	}

	return coordX == 0 && coordY == 0
}
