package http

import (
	"net/http"
	"time"

	"github.com/amirhnajafiz/starling/internal/ports/http/routes"

	"github.com/gorilla/mux"
)

// NewHandler initializes a new HTTP handler with the specified routes.
func NewHandler() *http.Server {
	// initialize a new mux router
	app := mux.NewRouter()

	// define the routes
	routes.NewRouter(app)

	return &http.Server{
		Handler:      app,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
