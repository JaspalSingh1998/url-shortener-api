package routes

import (
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/handler"
	"github.com/JaspalSingh1998/url-shortener-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(
	router *gin.Engine,
	linkHandler *handler.LinkHandler,
	analyticsHandler *handler.AnalyticsHandler,
	auth gin.HandlerFunc,
	rateLimiter *middleware.RateLimiter,
) {

	// 🌍 Public redirect (IP-based limit)
	router.GET(
		"/:shortCode",
		rateLimiter.LimitByIP(100, time.Minute),
		linkHandler.Redirect,
	)

	// 🔐 Protected APIs
	v1 := router.Group("/v1")
	v1.Use(auth)
	{
		v1.POST(
			"/links",
			rateLimiter.LimitByOrg(20, time.Minute),
			linkHandler.Create,
		)

		v1.GET(
			"/links/:id/analytics/daily",
			rateLimiter.LimitByOrg(60, time.Minute),
			middleware.RequireScope("analytics:read"),
			analyticsHandler.Daily,
		)

		v1.GET(
			"/links/:id/analytics/hourly",
			rateLimiter.LimitByOrg(60, time.Minute),
			middleware.RequireScope("analytics:read"),
			analyticsHandler.Hourly,
		)
	}
}
