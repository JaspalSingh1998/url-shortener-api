package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JaspalSingh1998/url-shortener-api/internal/config"
	"github.com/JaspalSingh1998/url-shortener-api/internal/handler"
	"github.com/JaspalSingh1998/url-shortener-api/internal/routes"
	"github.com/JaspalSingh1998/url-shortener-api/internal/server"
	"github.com/JaspalSingh1998/url-shortener-api/internal/service"
	"github.com/JaspalSingh1998/url-shortener-api/internal/store"
)

func Build(cfg *config.Config) (*server.Server, func(), error) {
	// DB
	db, err := pgxpool.New(context.Background(), cfg.DBURL())
	if err != nil {
		return nil, nil, err
	}

	// Dependencies
	linkStore := store.NewLinkStore(db)
	linkService := service.NewLinkService(linkStore)
	linkHandler := handler.NewLinkHandler(linkService, cfg.BaseURL)

	// Router
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	routes.Register(router, linkHandler)

	// Server
	srv := server.New(router, cfg.ServerPort)

	cleanup := func() {
		db.Close()
	}

	return srv, cleanup, nil
}
