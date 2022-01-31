package auth

import "golang.org/x/crypto/bcrypt"

// Hash plaintext to bcrypt
func Hash(plaintText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintText), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyHash compares plaintext to hash
// TODO: maybe the err should be handled better?
func VerifyHash(hash, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	return err == nil
}
