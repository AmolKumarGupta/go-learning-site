package controller

import (
	"amol/sample-site/config"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := struct{ App config.Config }{config.App}
	tmp.ExecuteTemplate(w, "login.html", data)
}
