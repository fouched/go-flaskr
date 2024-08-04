package render

import (
	"net/http"

	"github.com/a-h/templ"
)

func Template(w http.ResponseWriter, r *http.Request, template templ.Component) error {

	return template.Render(r.Context(), w)
}
