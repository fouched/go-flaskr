package repo

import (
	"github.com/fouched/go-flaskr/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(u models.User) error {

	stmt, _ := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	hashedPwd, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPwd)
	if err != nil {
		return err
	}

	return nil
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
