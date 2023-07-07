package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) retrieveBySlug(c *gin.Context) {
	slug := c.Param("slug")

	if slug == `` {
		newErrorResponse(c, http.StatusInternalServerError, "Your slug is empty!")
	}
	list, err := h.services.ResidentSlug.Get(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) refreshSlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == `` {
		newErrorResponse(c, http.StatusInternalServerError, "Your slug is empty!")
	}
	err := h.services.ResidentSlug.Refresh(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, nil)
}
