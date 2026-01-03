package routes

import (
	"github.com/JaspalSingh1998/url-shortener-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, linkHandler *handler.LinkHandler) {

	// Public redirect (NO /v1)
	router.GET("/:shortCode", linkHandler.Redirect)

	v1 := router.Group("/v1")
	{
		v1.POST("/links", linkHandler.Create)
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
