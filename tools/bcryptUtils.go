package tools

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	} else if err != nil {
		panic(err)
	}
	return nil
}
