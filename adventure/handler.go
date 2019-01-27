package adventure

import (
	"html/template"
	"net/http"
)

/*
GameEngine contains the whole story and possible transitions.
It satifies the http.Handler interface.
*/
type GameEngine struct {
	adventureData map[string]*Story
	tmpl          *template.Template
}

/*
GetGameEngine accepts the adventure data and returns a game engine
that contains the HTML template for rendering
*/
func GetGameEngine(adventureData map[string]*Story, tmplFile string) (engine *GameEngine, err error) {
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		return
	}
	engine = &GameEngine{
		adventureData: adventureData,
		tmpl:          t,
	}
	return
}

/*
ServeHTTP will be responsible for state transitions
in the adventure. It will serve the corresponding HTML
to the client.
*/
func (game *GameEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
