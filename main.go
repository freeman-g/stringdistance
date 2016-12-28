package main

import (
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/index.html",
))

func init() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	index.ExecuteTemplate(w, "index.html", nil)
}
