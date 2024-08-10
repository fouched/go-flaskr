package render

import (
	"github.com/fouched/go-flaskr/internal/config"
	"net/http"

	"github.com/a-h/templ"
)

var App *config.AppConfig

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	App = a
}

func Template(w http.ResponseWriter, r *http.Request, template templ.Component) error {

	return template.Render(r.Context(), w)
}
