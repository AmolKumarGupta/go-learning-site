package core

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("go-sample-site12")
	Store = sessions.NewCookieStore(key)
)

func AuthCheck(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "sessions")

	if name, ok := session.Values["name"].(string); !ok || name == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
}

func Generate(w http.ResponseWriter, r *http.Request, data map[string]any) {
	session, _ := Store.Get(r, "sessions")

	for str, val := range data {
		session.Values[str] = val
	}

	session.Save(r, w)
}
