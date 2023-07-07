package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kazakhattila/kapster/pkg/services"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {

	return &Handler{services: services}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/getAll", h.retrieveInfo)
		api.GET("/refreshAll", h.refreshList)
		api.GET("/slug/:slug", h.retrieveBySlug)
		api.GET("/refresh/:slug", h.refreshSlug)
		//	api.POST("/:id", h.returnStatus)
	}
	return router
}
