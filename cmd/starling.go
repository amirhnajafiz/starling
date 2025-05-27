package cmd

import "github.com/amirhnajafiz/starling/internal/core"

// Exec starts all starling servers.
func Exec() {
	// start both gRPC and HTTP servers concurrently
	go startGRPCServer()
	go startHTTPServer()

	// keep the main goroutine running
	core.StartCoreLoop()
}
