package handlers

import (
	"fmt"
	"github.com/fouched/go-flaskr/internal/forms"
	"github.com/fouched/go-flaskr/internal/models"
	"github.com/fouched/go-flaskr/internal/render"
	"github.com/fouched/go-flaskr/internal/templates"
	"net/http"
)

func (a *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["HeaderFoo"] = "HeaderBar"
	stringMap["PageFoo"] = "PageBar"
	fmt.Println(stringMap["Foo"])
	td := models.TemplateData{
		Form: forms.New(nil),
	}

	if td.Form.Errors.Get("foobar") != "" {
		fmt.Println(td.Form.Errors.Get("foobar"))
	} else {
		fmt.Println("No Errors")
	}

	component := templates.Home(&models.TemplateData{
		StringMap: stringMap,
	})
	_ = render.Template(w, r, component)
}
