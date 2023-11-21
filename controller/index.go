package controller

import (
	"amol/sample-site/config"
	"amol/sample-site/core"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	session, err := core.Store.Get(r, "sessions")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	name, _ := session.Values["name"].(string)

	data := struct {
		App      config.Config
		AuthName string
	}{
		config.App,
		name,
	}
	tmp.ExecuteTemplate(w, "index.html", data)
}
