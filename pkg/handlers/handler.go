package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

// 100% info that I need... -> Handler syntax shit

func NewHandler(services *service.Service) {

	return &Handler{services: services}

}
/* 
robbery of the century:
1. handlers get everything
2. solve my old problems!! 
3. postgres solve suka 
4. services get full suka 
5. rob and just rob suka. just rob everything... 


*/

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/getAll", h.retrieveInfo)
		api.GET("/refreshAll", h.refreshAll)
		api.GET("/:slug", h.retrieveBySlug)
		api.GET("/:slug/refresh", h.refreshSlug)
		api.POST("/:id", h.returnStatus)
	}
	return router
}
