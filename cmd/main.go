package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-flaskr/internal/config"
	"github.com/fouched/go-flaskr/internal/handlers"
	"github.com/fouched/go-flaskr/internal/helpers"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"log"
	"net/http"
	"time"
)

const port = ":9080"

var app config.AppConfig
var session *scs.SessionManager

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

	db, err := repo.CreateDb("./flaskr.db", false)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	} else {
		log.Println("Connected to database!")
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
