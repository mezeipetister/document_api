package user

import "golang.org/x/crypto/bcrypt"

// TODO: How fast is it? Please check it.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// TODO: Why is it soo slow? Check it!
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
