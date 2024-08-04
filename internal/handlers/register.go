package handlers

import (
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/templates"
	"net/http"
)

func RegisterGet(w http.ResponseWriter, r *http.Request) {

	component := templates.Register()
	_ = render.Template(w, r, component)
}
