package models

import "time"

type Post struct {
	ID       int
	AuthorID int
	Created  time.Time
	Title    string
	Body     string
}
