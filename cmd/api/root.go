package main

import (
	"net/http"

	"github.com/Light2Dark/death/internal/templates"
)

func (app application) rootHandler(w http.ResponseWriter, r *http.Request) {
	templates.Root().Render(r.Context(), w)
}
