package cryptography

import "golang.org/x/crypto/bcrypt"

func GenerateHash(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyIfHashMatches(hash, value string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
}
