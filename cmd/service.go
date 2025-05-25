package cmd

import "github.com/amirhnajafiz/starling/internal/ports/http"

func Start() {
	// initialize the HTTP server
	server := http.NewHandler()

	// start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
