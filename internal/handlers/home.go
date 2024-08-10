package handlers

import (
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
)

func (a *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	component := templates.Home(&models.TemplateData{})
	_ = render.Template(w, r, component)
}
