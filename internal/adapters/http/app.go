package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// Static files
	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Root path handler
	router.HandleFunc("/", IndexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
