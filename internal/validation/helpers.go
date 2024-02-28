package validation

import "net/mail"

func IsEmpty(value string) bool {
	return value == ""
}

func IsLower(value string, min int) bool {
	return len(value) < min
}

func IsLowerOrEqual(value string, min int) bool {
	return len(value) <= min
}

func IsGreater(value string, max int) bool {
	return len(value) > max
}

func IsGreaterOrEqual(value string, max int) bool {
	return len(value) >= max
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
