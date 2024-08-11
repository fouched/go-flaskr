package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/repo"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
)

func (a *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	td := DefaultTemplateData(r)

	posts, err := repo.SelectAllPosts()
	if err != nil {
		fmt.Println(err)
	} else {
		if len(posts) > 0 {
			td.Data["Posts"] = posts
		}
	}

	component := templates.Home(td)
	_ = render.Template(w, r, component)
}
