package httpserver

import (
	"net/http"
)

type Config struct {
	Port string `default:"8080" envconfig:"HTTP_PORT"`
}

type Server struct {
	server *http.Server
}

func New(handler http.Handler, c Config) *Server {
	// Настраиваем порт и таймауты

	return &Server{}
}

func (s *Server) Close() {
	// Shutdown
}
