package controller

import (
	"amol/sample-site/config"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := struct {
		App      config.Config
		AuthName string
	}{config.App, ""}
	tmp.ExecuteTemplate(w, "login.html", data)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, r.FormValue("name"))
}
