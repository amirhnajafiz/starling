package handler

import "github.com/gorilla/mux"

// Handler is a struct that represents the router for the HTTP server.
type Handler struct{}

// Register is a function that registers the routes with the given router.
func Register(app *mux.Router) {
	// initialize a new handler
	h := &Handler{}

	// register the routes with the handler methods
	app.HandleFunc("/healthz", h.health).Methods("GET")
	app.HandleFunc("/states", h.getState).Methods("GET")
	app.HandleFunc("/states", h.addState).Methods("POST")
}
