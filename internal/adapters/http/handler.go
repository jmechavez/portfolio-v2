package http

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {

	var err error
	templates, err = template.ParseFiles(
		"templates/index.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("template loaded")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
