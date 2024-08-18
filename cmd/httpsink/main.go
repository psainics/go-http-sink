package main

import (
	"fmt"
	"httpsink/internal/database"
	"httpsink/internal/server"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/lo"
)

func main() {
	appDb := database.GetDatabase()
	server := server.NewServer(appDb)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down...")
		lo.Must0(lo.Must(appDb.DB()).Close())
		lo.Must0(server.Close())
	}()

	slog.Info(fmt.Sprintf("Server is ready to handle requests at %s", server.Addr))
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
