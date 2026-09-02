package network

import (
	"encoding/json"
	"net/http"
)

type HTTPServer struct {
	mux *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	mux := http.NewServeMux()

	server := &HTTPServer{
		mux: mux,
	}

	mux.HandleFunc("GET /health", server.handleHealth)

	return server
}

func (s *HTTPServer) Handler() http.Handler {
	return s.mux
}

func (s *HTTPServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
