package handler

import (
	"net/http"

	"github.com/JaspalSingh1998/url-shortener-api/internal/service"
	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *service.LinkService
	baseURL string
}

func NewLinkHandler(service *service.LinkService, baseURL string) *LinkHandler {
	return &LinkHandler{
		service: service,
		baseURL: baseURL,
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
