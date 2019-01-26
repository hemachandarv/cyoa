package adventure

import (
	"net/http"
)

/*
GameEngine contains the whole story and possible transitions.
It satifies the http.Handler interface.
*/
type GameEngine struct {
}

/*
ServeHTTP will be responsible for state transitions
in the adventure. It will serve the corresponding HTML
to the client.
*/
func (game *GameEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
