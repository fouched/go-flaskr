package handlers

import (
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	component := templates.Home()
	_ = render.Template(w, r, component)
}
