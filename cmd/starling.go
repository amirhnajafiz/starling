package cmd

// Exec starts all starling servers.
func Exec() {
	// start both gRPC and HTTP servers concurrently
	go startGRPCServer()
	go startHTTPServer()
}
