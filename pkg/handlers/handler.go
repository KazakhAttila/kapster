package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/kazakhattila/kapster/pkg/services"

)

type Handler struct {
	services *services.Service
}

// 100% info that I need... -> Handler syntax shit

func NewHandler(services *services.Service) (*Handler){

	return &Handler{services: services}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/getAll", h.retrieveInfo)
		api.GET("/refreshAll", h.refreshList)
		api.GET("/:slug", h.retrieveBySlug)
		api.GET("/:slug/refresh", h.refreshSlug)
		api.POST("/:id", h.returnStatus)
	}
	return router
}
