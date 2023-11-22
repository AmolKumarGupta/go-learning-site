package controller

import (
	"amol/sample-site/config"
	"amol/sample-site/core"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var session *sessions.Session
	var flashMsg string

	tmp, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if session, err = core.Store.Get(r, "sessions"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	flashes := session.Flashes("error")
	core.Store.Save(r, w, session)
	if len(flashes) != 0 {
		flashMsg = flashes[0].(string)
	}

	data := struct {
		App      config.Config
		AuthName string
		Error    string
	}{config.App, "", flashMsg}

	tmp.ExecuteTemplate(w, "login.html", data)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	var (
		session *sessions.Session
		err     error
	)

	if session, err = core.Store.Get(r, "sessions"); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := core.FindByEmail(email)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if user == nil {
		session.AddFlash("User not Exists", "error")

		if err := core.Store.Save(r, w, session); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if password != user.Password {
		session.AddFlash("Passwords do not match", "error")

		if err := core.Store.Save(r, w, session); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	core.Generate(w, r, map[string]any{"name": user.Name, "email": user.Email})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
