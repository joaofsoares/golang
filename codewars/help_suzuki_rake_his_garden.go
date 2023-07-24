package codewars

import "strings"

func mapFunc(words *[]string) {
	for i, v := range *words {
		if v != "rock" && v != "gravel" {
			(*words)[i] = "gravel"
		}
	}
}

func RakeGarden(garden string) string {
	words := strings.Fields(garden)

	mapFunc(&words)

	return strings.Join(words, " ")
}
