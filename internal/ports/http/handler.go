package http

import (
	"github.com/amirhnajafiz/starling/internal/ports/http/routes"

	"github.com/gorilla/mux"
)

// NewHandler initializes a new HTTP handler with the specified routes.
func NewHandler() *mux.Router {
	// initialize a new mux router
	app := mux.NewRouter()

	// define the routes
	router := routes.NewRouter()

	// register the routes with the router
	app.HandleFunc("/healthz", router.Health).Methods("GET")
	app.HandleFunc("/states", router.GetState).Methods("GET")
	app.HandleFunc("/states", router.AddState).Methods("POST")

	return app
}
