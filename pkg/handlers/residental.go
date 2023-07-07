package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

1. true banality of these problems
2. i am over reacting.
3. math is much better than this shit.
4. I need a good map.

*/

func (h *Handler) refreshList(c *gin.Context) {
	err := h.services.Resident.Refresh()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

}

func (h *Handler) retrieveInfo(c *gin.Context) {
	list, err := h.services.Resident.Get()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)

}
