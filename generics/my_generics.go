package generics

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var empty T
		return empty
	}
	return s[len(s)-1]
}
