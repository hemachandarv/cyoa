package adventure

import (
	"html/template"
	"net/http"
	"strings"
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
to the client by rendering the parsed template with context data.
*/
func (game *GameEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := currentArc(r.URL.Path)
	tmplData, ok := game.adventureData[arc]
	if ok {
		game.tmpl.Execute(w, tmplData)
	} else {
		http.NotFound(w, r)
	}
}

func currentArc(path string) (arc string) {
	p := strings.Split(path, "/")
	if len(p) == 0 {
		return
	}
	arc = p[len(p)-1]
	return
}
