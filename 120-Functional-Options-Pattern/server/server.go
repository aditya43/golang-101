package server

import (
	"log"
	"time"
)

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(option ...Option) *Server {
	s := &Server{}
	for _, o := range option {
		o(s)
	}
	return s
}

func (s *Server) Start() {
	// ...
	log.Printf("server started %#v\n", s)
}

func (s *Server) Stop() {
	// ...
	log.Printf("server stopped %#v\n", s)
}
