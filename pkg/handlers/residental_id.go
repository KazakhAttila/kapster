package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type slugg struct {
	Slug string `json:"slug"`
}

func (h *Handler) retrieveBySlug(c *gin.Context) {
	var slug slugg
	err := c.BindJSON(&slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Your slug is empty!")
	}
	list, err := h.services.ResidentSlug.Get(slug.Slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) refreshSlug(c *gin.Context) {
	var slug slugg
	err := c.BindJSON(&slug)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Your slug is empty!")
	}
	err = h.services.ResidentSlug.Refresh(slug.Slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, nil)
}
