package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) refreshList(c *gin.Context) {

}

func (h *Handler) retrieveInfo(c *gin.Context) {
	list, err := h.services.retrieveInfo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)

}
