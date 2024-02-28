package helpers

func If[T any](condition bool, right T, left T) T {
	if condition {
		return right
	}

	return left
}
