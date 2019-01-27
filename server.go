package main

import (
	"log"

	"github.com/hemv/cyoa/adventure"
)

const (
	adventureFile = "adventure/adventure.json"
)

func main() {
	adventure, err := adventure.ParseJSON(adventureFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("%v", adventure)
}
