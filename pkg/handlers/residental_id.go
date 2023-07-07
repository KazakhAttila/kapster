package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) retrieveBySlug(c *gin.Context) {
		list, err := h.services.ResidentSlug.Get()
		if err!=nil{
			
		}
		c.JSON(list)
}

func (h *Handler) refreshSlug(c *gin.Context) {
		list, err := 
}
