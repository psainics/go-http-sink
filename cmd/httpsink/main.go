package main

import (
	"fmt"
	"httpsink/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("Shutting down...")
		server.Close()
	}()

	fmt.Println("Server is ready to handle requests at", server.Addr)
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
