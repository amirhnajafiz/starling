package grpc

import (
	"net"

	"github.com/amirhnajafiz/starling/internal/ports/grpc/service"
	"github.com/amirhnajafiz/starling/pkg/raftpb"

	"google.golang.org/grpc"
)

// NewServer creates a new gRPC server instance and registers the RaftServer implementation.
func NewServer() *grpc.Server {
	// create a new gRPC server instance
	server := grpc.NewServer()

	// register the RaftServer implementation with the gRPC server
	raftpb.RegisterRaftServer(server, &service.RaftServerImpl{})

	return server
}

// StartServer starts the gRPC server on the specified address.
func StartServer(server *grpc.Server, address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// start the gRPC server
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
