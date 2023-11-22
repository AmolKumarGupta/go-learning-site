package controller

import (
	"amol/sample-site/config"
	"amol/sample-site/core"
	"html/template"
	"net/http"
	"time"
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
	session, err := core.Store.Get(r, "sessions")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	session.Options.MaxAge = -1

	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	if err := core.Store.Save(r, w, session); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
