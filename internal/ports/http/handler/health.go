package handler

import "net/http"

// health is a handler function that responds to health check requests.
func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
