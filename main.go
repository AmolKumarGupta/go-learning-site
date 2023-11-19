package main

import (
	"amol/sample-site/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/login", controller.Login).Methods("GET")

	fmt.Println("site running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
