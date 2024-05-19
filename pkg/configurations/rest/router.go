package rest

import (
	"github.com/gin-gonic/gin"
	"stocktacking_backend/pkg/configurations/rest/handlers"
	"stocktacking_backend/pkg/configurations/rest/services"
)

type Router struct {
	confHandler  *handlers.ConfigurationsHandler
	optHandler   *handlers.OptionsHandler
	errorHandler services.ErrorHandler
}

func NewRouter(
	confHandler *handlers.ConfigurationsHandler,
	optHandler *handlers.OptionsHandler,
	errorHandler services.ErrorHandler,
) *Router {
	return &Router{
		confHandler:  confHandler,
		optHandler:   optHandler,
		errorHandler: errorHandler,
	}
}

func (r *Router) SetRoutes(routerGroup *gin.RouterGroup) {
	// ----------------- CONFIGURATIONS -----------------
	configurations := routerGroup.Group("configurations")
	configurations.Use(r.errorHandler.Handle) // отлов ошибок

	configurations.GET("", r.confHandler.List)
	configurations.POST("", r.confHandler.Create)

	configuration := configurations.Group(":" + handlers.ConfigurationIdParam)
	configuration.GET("", r.confHandler.Get)
	configuration.PUT("", r.confHandler.Update)
	configuration.DELETE("", r.confHandler.Delete)

	// ----------------- OPTIONS -----------------
	options := configuration.Group("options")
	options.GET("", r.optHandler.List)
	options.POST("", r.optHandler.Create)
	options.PUT("", r.optHandler.UpdateList)

	option := options.Group(":" + handlers.OptionIdParam)
	option.GET("", r.optHandler.Get)
	option.PUT("", r.optHandler.Update)
	option.DELETE("", r.optHandler.Delete)
}
