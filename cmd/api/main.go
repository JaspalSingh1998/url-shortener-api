package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/config"
	"github.com/JaspalSingh1998/url-shortener-api/internal/routes"
	"github.com/JaspalSingh1998/url-shortener-api/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	routes.Register(router)

	srv := server.New(router, cfg.ServerPort)

	// Run Server in goroutine

	go func() {
		log.Printf("Server started on :%s", cfg.ServerPort)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Listen for shutdown signals

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received")

	// Context with timeout for graceful shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
