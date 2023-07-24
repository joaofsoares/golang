package stringhelper

func FindString(target string, strs *[]string) string {

	for i, j := 0, len(*strs)-1; i <= j; i, j = i+1, j-1 {

		if (*strs)[i] == target {
			return (*strs)[i]
		}

		if (*strs)[j] == target {
			return (*strs)[j]
		}

	}

	return ""

}
