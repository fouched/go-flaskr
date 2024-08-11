package repo

import (
	"context"
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

	_, err = stmt.Exec(u.Username, hashedPwd)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllPosts() ([]*models.Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT id, author_id, created, title, body FROM post"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Created,
			&post.Title,
			&post.Body,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func Authenticate(u models.User) (int, error) {
	stmt, _ := db.Prepare("SELECT id, password FROM user WHERE username=?")
	defer stmt.Close()

	var id int
	var hashedPassword string

	row := stmt.QueryRow(u.Username)
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
