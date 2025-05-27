package cmd

import (
	"github.com/amirhnajafiz/starling/internal/ports/grpc"
	"github.com/amirhnajafiz/starling/internal/ports/http"
)

func Start() {
	// initialize the gRPC server
	grpcServer := grpc.NewServer()

	go func() {
		// start the gRPC server on the specified address
		if err := grpc.StartServer(grpcServer, ":50051"); err != nil {
			panic(err)
		}
	}()

	// initialize the HTTP server
	server := http.NewServer()

	// start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
