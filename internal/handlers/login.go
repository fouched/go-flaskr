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

func (a *HandlerConfig) LoginGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Login(DefaultTemplateData(r))
	_ = render.Template(w, r, component)
}

func (a *HandlerConfig) LoginPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		component := templates.Login(&models.TemplateData{
			Form: form,
		})
		_ = render.Template(w, r, component)
		return
	}

	user := models.User{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}
	id, err := repo.Authenticate(user)
	if err != nil {
		form.Errors.Add("heading", "Invalid credentials")
		component := templates.Login(&models.TemplateData{
			Form: form,
		})
		_ = render.Template(w, r, component)
		return
	}

	a.App.Session.Put(r.Context(), "userId", id)
	// Good practice: prevent a post re-submit with a http redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// Logout logs a user out
func (a *HandlerConfig) Logout(w http.ResponseWriter, r *http.Request) {

	_ = a.App.Session.Destroy(r.Context())
	_ = a.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
