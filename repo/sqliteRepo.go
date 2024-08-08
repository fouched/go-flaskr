package repo

import (
	"github.com/fouched/go-flaskr/models"
)

func InsertUser(u models.User) error {

	stmt, _ := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	_, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}
