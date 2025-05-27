package cmd

import "github.com/amirhnajafiz/starling/internal/ports/grpc"

func startGRPCServer() {
	// initialize the gRPC server
	grpcServer := grpc.NewServer()

	// start the gRPC server on the specified address
	if err := grpc.StartServer(grpcServer, ":50051"); err != nil {
		panic(err)
	}
}
