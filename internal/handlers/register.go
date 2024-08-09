package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"github.com/fouched/go-flaskr/internal/templates"
	"github.com/fouched/go-flaskr/internal/validation"
	"net/http"
)

func (a *HandlerConfig) RegisterGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Register(validation.Form{})
	_ = render.Template(w, r, component)
}

func (a *HandlerConfig) RegisterPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	form := validation.New(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		component := templates.Register(*form)
		_ = render.Template(w, r, component)
		return
	}

	// validation passed persist the form
	user := models.User{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	err = repo.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		form.Errors.Add("heading", "Unexpected error, please try again later.")
		component := templates.Register(*form)
		_ = render.Template(w, r, component)
		return
	}

	component := templates.Home()
	_ = render.Template(w, r, component)
}
