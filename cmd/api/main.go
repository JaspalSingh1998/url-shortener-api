package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/JaspalSingh1998/url-shortener-api/internal/app"
	"github.com/JaspalSingh1998/url-shortener-api/internal/config"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	// Build application (DI + router + server)
	srv, cleanup, err := app.Build(cfg)
	if err != nil {
		log.Fatalf("failed to build app: %v", err)
	}
	defer cleanup()

	// Run server in goroutine
	go func() {
		log.Printf("Server started on :%s", cfg.ServerPort)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
