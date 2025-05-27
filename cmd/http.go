package cmd

import "github.com/amirhnajafiz/starling/internal/ports/http"

func startHTTPServer() {
	// initialize the HTTP server
	server := http.NewServer()

	// start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
