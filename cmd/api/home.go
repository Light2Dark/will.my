package main

import (
	"net/http"

	"github.com/Light2Dark/death/internal/templates"
)

func (app application) homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}

func (app application) explanationHandler(w http.ResponseWriter, r *http.Request) {
	templates.Explanation().Render(r.Context(), w)
}

func (app application) mainInfoHandler(w http.ResponseWriter, r *http.Request) {
	templates.MainInfo().Render(r.Context(), w)
}
