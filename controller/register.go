package controller

import (
	"amol/sample-site/config"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := struct{ App config.Config }{config.App}
	tpl.ExecuteTemplate(w, "register.html", data)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
