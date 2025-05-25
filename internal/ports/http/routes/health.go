package routes

import "net/http"

// health is a handler function that responds to health check requests.
func (r *Router) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
