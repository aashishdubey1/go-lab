package encapsulationcomposition

import (
	"fmt"
	"time"
)

// -> -> ->  Immutable Config via Options Pattern
// Build a Server struct configured through functional options:
// func WithPort(p int) Option,
// func WithTimeout(d time.Duration) Option,
// combined in NewServer(opts ...Option) *Server.
// This is the idiomatic replacement for constructor overloading, which Go doesn't support.

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

type Option func(*Server)

func WithPort(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}
func WithTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.Timeout = d
	}
}

func NewServer(host string, opts ...Option) *Server {
	server := &Server{
		Host:    host,
		Timeout: 30 * time.Second,
		Port:    8080,
	}

	for _, opt := range opts {
		opt(server)
	}
	return server
}

func (s *Server) Start() {
	protocol := "http"
	fmt.Printf("Starting server on %s://%s:%d (Timeout: %v)\n",
		protocol, s.Host, s.Port, s.Timeout)
}

func RunEx3() {
	server1 := NewServer("localhost")
	server1.Start()

	server2 := NewServer("0.0.0.0", WithPort(9000), WithTimeout(10*time.Second))
	server2.Start()

}
