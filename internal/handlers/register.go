package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/forms"
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
	"strings"
)

func (a *HandlerConfig) RegisterGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Register(DefaultTemplateData(r))
	_ = render.Template(w, r, component)
}

func (a *HandlerConfig) RegisterPost(w http.ResponseWriter, r *http.Request) {

	td := DefaultTemplateData(r)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	form := forms.New(r.PostForm)
	form.Required("username", "password")
	if !form.Valid() {
		td.Form = form
		component := templates.Register(td)
		_ = render.Template(w, r, component)
		return
	}

	// forms passed persist the form
	user := models.User{
		Username: strings.ToLower(r.Form.Get("username")),
		Password: r.Form.Get("password"),
	}

	err = repo.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "UNIQUE constraint") {
			td.Form.Errors.Add("heading", "Username "+user.Username+" already taken.")
		} else {
			td.Form.Errors.Add("heading", "Unexpected error, please try again later.")
		}
		component := templates.Register(td)
		_ = render.Template(w, r, component)
		return
	}

	// Good practice: prevent a post re-submit with a http redirect
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
