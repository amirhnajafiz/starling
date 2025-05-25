package routes

import "github.com/gorilla/mux"

// Router is a struct that represents the router for the HTTP server.
type Router struct{}

// NewRouter creates a new instance of the Router struct.
func NewRouter(app *mux.Router) {
	// initialize a new router
	router := &Router{}

	// register the routes with the router
	app.HandleFunc("/healthz", router.Health).Methods("GET")
	app.HandleFunc("/states", router.GetState).Methods("GET")
	app.HandleFunc("/states", router.AddState).Methods("POST")
}
