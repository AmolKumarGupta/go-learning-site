package controller

import (
	"amol/sample-site/config"
	"amol/sample-site/core"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := struct {
		App      config.Config
		AuthName string
	}{config.App, ""}
	tpl.ExecuteTemplate(w, "register.html", data)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	core.NewUser(name, email, password)
	core.Generate(w, r, map[string]any{"name": name, "email": email})

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := core.Store.Get(r, "sessions")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
