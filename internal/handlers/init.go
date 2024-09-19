package handlers

import (
	"github.com/fouched/go-flaskr/internal/config"
	"github.com/fouched/go-flaskr/internal/forms"
	"github.com/fouched/go-flaskr/internal/helpers"
	"github.com/fouched/go-flaskr/internal/models"
	"net/http"
)

var Instance *HandlerConfig

type HandlerConfig struct {
	App *config.AppConfig
}

func NewConfig(a *config.AppConfig) *HandlerConfig {
	return &HandlerConfig{
		App: a,
	}
}

func NewHandlers(h *HandlerConfig) {
	Instance = h
}

func DefaultTemplateData(r *http.Request) *models.TemplateData {
	return &models.TemplateData{
		IsAuthenticated: helpers.IsAuthenticated(r),
		Form:            forms.New(nil),
	}
}
