package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	data := "Adaman"
	tmpl, _ := template.ParseFiles("assets/template/index.html")
	tmpl.Execute(w, data)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", index)
	http.ListenAndServe(":8080", r)
}
