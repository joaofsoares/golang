package codewars

import (
	"strconv"
	"strings"
)

func IpsBetween(start string, end string) int {

	ranges := 256 * 256 * 256
	sum := 0

	splitStart := strings.Split(start, ".")
	splitEnd := strings.Split(end, ".")

	for cnt := 0; cnt < 4; cnt++ {

		numStart, err := strconv.Atoi(splitStart[cnt])
		if err != nil {
			return 0
		}

		numEnd, err := strconv.Atoi(splitEnd[cnt])
		if err != nil {
			return 0
		}

		tmp := numEnd - numStart

		sum += tmp * ranges
		ranges = ranges / 256
	}

	return sum
}
