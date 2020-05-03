package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

// New inicialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
