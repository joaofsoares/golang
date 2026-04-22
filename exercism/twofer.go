package exercism

import "fmt"

func ShareWith(name string) string {
	person := "you"

	if len(name) > 0 {
		person = name
	}

	return fmt.Sprintf("One for %s, one for me.", person)
}
