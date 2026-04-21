package exercism

import (
	"strconv"
	"strings"
)

func Convert(number int) string {

	var isDivThree, isDivFive, isDivSeven bool = (number%3 == 0), (number%5 == 0), (number%7 == 0)

	if isDivFive || isDivSeven || isDivThree {
		var result strings.Builder

		if isDivThree {
			result.WriteString("Pling")
		}

		if isDivFive {
			result.WriteString("Plang")
		}

		if isDivSeven {
			result.WriteString("Plong")
		}

		return result.String()
	}

	return strconv.Itoa(number)
}
