package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Name string
}

var App Config = Config{Name: "Sample site"}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/", indexPage)

	fmt.Println("site running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("views/*")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := struct{ App Config }{App}
	tmp.ExecuteTemplate(w, "index.html", data)
}
