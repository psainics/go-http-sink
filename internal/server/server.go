package server

import (
	"fmt"
	"httpsink/internal/config"
	"net/http"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	newServer := &Server{port: config.SERVER_PORT}
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
