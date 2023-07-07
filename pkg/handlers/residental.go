package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) refreshList(c *gin.Context) {
	err := h.services.Resident.Refresh()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

}

func (h *Handler) retrieveInfo(c *gin.Context) {
	list, err := h.services.Resident.Get()
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} 
	c.JSON(http.StatusOK, list) 

}
