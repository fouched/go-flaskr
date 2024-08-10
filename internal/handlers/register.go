package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/forms"
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
)

func (a *HandlerConfig) RegisterGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Register(&models.TemplateData{
		Form: forms.New(nil),
	})
	_ = render.Template(w, r, component)
}

func (a *HandlerConfig) RegisterPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		component := templates.Register(&models.TemplateData{
			Form: form,
		})
		_ = render.Template(w, r, component)
		return
	}

	// forms passed persist the form
	user := models.User{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	err = repo.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		form.Errors.Add("heading", "Unexpected error, please try again later.")
		component := templates.Register(&models.TemplateData{
			Form: form,
		})
		_ = render.Template(w, r, component)
		return
	}

	component := templates.Home(&models.TemplateData{})
	_ = render.Template(w, r, component)
}
