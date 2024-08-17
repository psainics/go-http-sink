package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route Groups
	
	e.GET("/", s.HelloWorldHandler)
	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
