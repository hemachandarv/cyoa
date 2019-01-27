package main

import (
	"log"
	"net/http"

	"github.com/hemv/cyoa/adventure"
)

const (
	adventureFile = "adventure/adventure.json"
	gameTemplate  = "tmpl/index.gohtml"
)

func main() {
	adventureData, err := adventure.ParseJSON(adventureFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	engine, err := adventure.GetGameEngine(adventureData, gameTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	mux := defaultMux()
	mux.Handle("/cyoa", engine)
	http.ListenAndServe(":8000", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	http.HandleFunc("/", notFoundHandler)
	return mux
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Error! Not Found!"))
}
