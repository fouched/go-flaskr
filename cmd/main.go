package main

import (
	"database/sql"
	"fmt"
	"github.com/fouched/go-flaskr/internal/repo"
	"log"
	"net/http"
)

const port = ":9080"

func main() {

	db, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	// we have database connectivity, close it after app stops
	defer db.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	fmt.Printf("Starting application on %s\n", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}

func initApp() (*sql.DB, error) {

	db, err := repo.CreateDb("./flaskr.db")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	} else {
		log.Println("Connected to database!")
	}

	return db, nil
}
