package handler

import (
	"context"
	"net/http"

	"github.com/JaspalSingh1998/url-shortener-api/internal/model"
	"github.com/JaspalSingh1998/url-shortener-api/internal/service"
	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service      *service.LinkService
	clickService *service.ClickService
	baseURL      string
}

func NewLinkHandler(
	linkService *service.LinkService,
	clickService *service.ClickService,
	baseURL string,
) *LinkHandler {
	return &LinkHandler{
		service:      linkService,
		clickService: clickService,
		baseURL:      baseURL,
	}
}

func (h *LinkHandler) Create(c *gin.Context) {
	var req CreateLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	link, err := h.service.CreateLink(
		c.Request.Context(),
		req.OriginalURL,
		req.CustomAlias,
		req.ExpiresAt,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := CreateLinkResponse{
		ID:          link.ID,
		ShortCode:   link.ShortCode,
		ShortURL:    h.baseURL + "/" + link.ShortCode,
		OriginalURL: link.OriginalURL,
		ExpiresAt:   link.ExpiresAt,
		CreatedAt:   link.CreatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{"data": resp})

}

func (h *LinkHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")

	link, err := h.service.ResolveLink(c.Request.Context(), shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "link not found or expired",
		})
		return
	}

	// Async Click tracking
	go h.trackClick(c, link)

	c.Redirect(http.StatusFound, link.OriginalURL)
}

func (h *LinkHandler) trackClick(c *gin.Context, link *model.Link) {
	e := &model.ClickEvent{
		LinkID:    link.ID,
		ShortCode: link.ShortCode,
		IPAddress: c.ClientIP(),
		UserAgent: c.GetHeader("User-Agent"),
		Referrer:  c.GetHeader("Referer"),
	}

	h.clickService.Track(context.Background(), e)
}
