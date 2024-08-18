package server

import (
	"fmt"
	"httpsink/internal/config"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Server struct {
	port int
	db   *gorm.DB
}

func NewServer(db *gorm.DB) *http.Server {
	newServer := &Server{port: config.SERVER_PORT, db: db}
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
