package routes

import (
	"github.com/JaspalSingh1998/url-shortener-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, linkHandler *handler.LinkHandler, analyticsHandler *handler.AnalyticsHandler) {

	// Public redirect (NO /v1)
	router.GET("/:shortCode", linkHandler.Redirect)

	v1 := router.Group("/v1")
	{
		v1.POST("/links", linkHandler.Create)
		v1.GET("/links/:id/analytics/daily", analyticsHandler.Daily)
		v1.GET("/links/:id/analytics/hourly", analyticsHandler.Hourly)
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
