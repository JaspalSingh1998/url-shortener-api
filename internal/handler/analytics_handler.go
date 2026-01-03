package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JaspalSingh1998/url-shortener-api/internal/service"
	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	service *service.AnalyticsService
}

func NewAnalyticsHandler(service *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) Daily(c *gin.Context) {
	linkID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	from, _ := time.Parse("2006-01-02", c.Query("from"))
	to, _ := time.Parse("2006-01-02", c.Query("to"))

	stats, err := h.service.Daily(c.Request.Context(), linkID, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"link_id":     linkID,
			"granularity": "daily",
			"items":       stats,
		},
	})
}

func (h *AnalyticsHandler) Hourly(c *gin.Context) {
	linkID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	from, _ := time.Parse(time.RFC3339, c.Query("from"))
	to, _ := time.Parse(time.RFC3339, c.Query("to"))

	stats, err := h.service.Hourly(c.Request.Context(), linkID, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"link_id":     linkID,
			"granularity": "hourly",
			"items":       stats,
		},
	})
}
