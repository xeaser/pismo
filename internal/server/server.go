package server

import (
	"fmt"
	"net/http"

	"github.com/xeaser/pismo/config"
	"github.com/xeaser/pismo/internal/account"
)

// Server represents the HTTP server
type Server struct {
	port string
}

// NewServer creates a new instance of Server
func NewServer(cfg *config.Server) *Server {
	return &Server{
		port: cfg.Port,
	}
}

// Start initializes and starts the HTTP server and Register api handlers
func (s *Server) Start() error {
	mux := http.NewServeMux()
	account.RegisterHandler(mux)
	fmt.Printf("Server is running on port %s\n", s.port)
	fmt.Println()
	return http.ListenAndServe(":"+s.port, mux)
}
