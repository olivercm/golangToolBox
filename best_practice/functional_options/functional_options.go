package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr string
	Port int
	Protocol string
	Timeout time.Duration
	Maxconns int
	TLS *tls.Config
}

type Option func(server *Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server *Server)) (*Server, error) {
	srv := Server{
		Addr: addr,
		Port: port,
		Protocol: "tcp",
		Timeout: 20* time.Second,
		Maxconns: 1000,
		TLS: nil,
	}
	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main() {
	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v", s3)
}

