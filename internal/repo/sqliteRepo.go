package repo

import (
	"errors"
	"github.com/fouched/go-flaskr/internal/models"
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

func Authenticate(u models.User) (int, error) {
	stmt, _ := db.Prepare("SELECT id, password FROM user WHERE username=?")
	defer stmt.Close()

	var id int
	var hashedPassword string

	row := stmt.QueryRow(u.Email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, err
	}

	if VerifyPassword(u.Password, hashedPassword) {
		return id, nil
	} else {
		return 0, errors.New("invalid password")
	}
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
