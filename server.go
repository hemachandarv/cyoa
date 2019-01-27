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
	gameEngine, err := adventure.GetGameEngine(adventureData, gameTemplate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/cyoa/", gameEngine)
	http.ListenAndServe(":8000", mux)
}
