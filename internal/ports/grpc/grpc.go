package grpc

import (
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
