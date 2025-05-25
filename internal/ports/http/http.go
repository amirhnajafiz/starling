package http

import (
	"net/http"
	"time"

	"github.com/amirhnajafiz/starling/internal/ports/http/handler"

	"github.com/gorilla/mux"
)

// NewServer initializes a new HTTP server with the specified routes.
func NewServer() *http.Server {
	// initialize a new mux router
	app := mux.NewRouter()

	// define the routes
	handler.Register(app)

	return &http.Server{
		Handler:      app,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
