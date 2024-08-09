package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"github.com/fouched/go-flaskr/templates"
	"net/http"
)

func RegisterGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Register()
	_ = render.Template(w, r, component)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	user := models.User{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	err = repo.InsertUser(user)
	if err != nil {
		fmt.Println(err)
	}

}
