package handlers 


import( 

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

)


type Handler struct {
	services *service.Service
}
// 100% info that I need... -> Handler syntax shit

func NewHandler(services  *service.Service){ 
			
			return &Handler{services: services}

}

// 100% info that I need... 
//  gin shit. why did they make this decision and why did they divide? hui ego znaet...
// stoit li ono togo chtobi izi chitalos'??? 
// kakie ewe tut mozhno tradi sdelat' nahui??? 
// kakie otsyuda mozhno poluchit' trade ups??? 
// voobshe v celom was -> absurd or relax or pohui or new stories??? which one...

// goal right now -> Handlers full do the shit... 
// then also, another goal - > new schemes and so on
// another shit -> repos, services -> do everything, reach the testing stage. 
// make the schemes. 
// plan other Robberies, which means find other repositories like that!!! 


// now -> try to understand why i couldn't make such decisions or speak this languages or get these schemes? 
// absurd questions and absurd positions point by point is needed. 
// 

func (h *Handler) InitRoutes() *gin.Engine{ 
		router := gin.New() 
		router.GET("/swagger/*any",  ginSwagger.WrapHandler(swaggerFiles.Handler))
		api:= router.Group("/api")
		{ 
			api.GET("/getAll", h.retrieveInfo) 
			api.GET("/refreshAll", h.refreshAll)
			api.GET("/:slug", h.retrieveBySlug) 
			api.GET("/:slug/refresh", h.refreshSlug)
			api.POST("/:id", h.returnStatus)
		}

}