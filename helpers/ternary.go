package helpers

func If[T any](condition bool, right T, left T) T {
	if condition {
		return right
	}

	return left
}

func IfOrNil[T any](condition bool, right T) *T {
	if condition {
		return &right
	}

	return nil
}
