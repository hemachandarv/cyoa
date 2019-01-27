package adventure

import (
	"encoding/json"
	"os"
)

/*
Option contains the next arc and its description
from a specific arc.
*/
type Option struct {
	Arc         string `json:"arc"`
	Description string `json:"text"`
}

/*
Story contains the details like arc-name,
definition of the arc and possible paths
from this arc.
*/
type Story struct {
	Title   string   `json:"title"`
	Content []string `json:"story"`
	Paths   []Option `json:"options"`
}

/*
ParseJSON parses the stringified adventure "data" and
returns a map where Key is the arc-name and value is the
corresponsing Story struct that contains the
details from input data.
*/
func ParseJSON(filename string) (adventure map[string]Story, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return
	}
	d := json.NewDecoder(f)
	err = d.Decode(&adventure)
	return
}
