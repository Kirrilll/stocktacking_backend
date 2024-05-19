package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stocktacking_backend/internal/adapters/rest/handlers"
)

type Router struct {
	ginMode         string
	apiSettings     APISettings
	fixturesHandler *handlers.FixturesHandler
	//mediaHandler *
}

type APISettings struct {
	Root        string
	Version     string
	ServiceName string
	//PkgPath     string
}

func NewRouter(
	ginMode string,
	apiSettings APISettings,
	fixturesHandler *handlers.FixturesHandler,
) *Router {
	return &Router{
		ginMode:         ginMode,
		apiSettings:     apiSettings,
		fixturesHandler: fixturesHandler,
	}
}

func (r Router) Router() *gin.Engine {
	gin.SetMode(r.ginMode)
	router := gin.New()
	// обработка ошибок page_repository not found и method not allowed
	router.NoRoute(
		func(c *gin.Context) {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"applicationErrorCode": http.StatusText(http.StatusNotFound),
					"url":                  c.Request.URL.Path,
				},
			)
		},
	)
	router.NoMethod(
		func(c *gin.Context) {
			c.JSON(
				http.StatusMethodNotAllowed, gin.H{
					"applicationErrorCode": http.StatusText(http.StatusMethodNotAllowed),
					"message":              "method not allowed",
				},
			)
		},
	)

	healthCheck := func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) }

	api := router.Group(r.apiSettings.Root)
	v1 := api.Group(r.apiSettings.Version)

	// internal
	service := v1.Group(r.apiSettings.ServiceName)
	//service.Use(r.errorHandler.Handle)
	//service.Use(logger.Logger())
	service.Group("health").Group("check").GET("", healthCheck)

	fixtures := service.Group("fixtures")
	fixtures.POST("run", r.fixturesHandler.RunFixtures)

	return router
}

func (r Router) GetRouter() *gin.Engine {
	return r.Router()
}
