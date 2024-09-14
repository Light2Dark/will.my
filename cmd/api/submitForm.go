package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Light2Dark/death/internal/templates"
)

type FormData struct {
	name  string
	email string
	will  string
}

func (app application) submitFormHandler(w http.ResponseWriter, r *http.Request) {
	var message string = "Submitted! We will contact you via email"
	r.ParseForm()

	formData := FormData{
		name:  r.FormValue("name"),
		email: r.FormValue("email"),
		will:  r.FormValue("will"),
	}

	if strings.Trim(formData.name, "") == "" {
		message = "Empty name, please try again."
	} else {
		message = fmt.Sprintf("%s %s.", message, formData.name)
	}

	templates.JobResult(message).Render(r.Context(), w)
}
