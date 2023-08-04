package codewars

func Rot13(msg string) string {

	if len(msg) == 0 {
		return ""
	}

	bytes := []byte(msg)

	cnt := 0

	for cnt < len(bytes) {
		if bytes[cnt] == 92 && cnt+1 < len(bytes) {
			cnt += 2
			continue
		} else if (bytes[cnt] > 96 && bytes[cnt] < 110) || (bytes[cnt] > 64 && bytes[cnt] < 78) {
			bytes[cnt] = bytes[cnt] + 13
		} else if (bytes[cnt] > 109 && bytes[cnt] < 123) || (bytes[cnt] > 77 && bytes[cnt] < 91) {
			bytes[cnt] = bytes[cnt] - 13
		}

		cnt += 1
	}

	return string(bytes)
}
