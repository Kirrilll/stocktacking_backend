package rest

import (
	"github.com/sarulabs/di/v2"
	"stocktacking_backend/pkg/configurations/plugin/actions"
	"stocktacking_backend/pkg/configurations/rest/handlers"
	"stocktacking_backend/pkg/configurations/rest/services"
)

var Definitions = []di.Def{
	{
		Name: "focus.configurations.handlers.configurations",
		Build: func(ctn di.Container) (interface{}, error) {
			configurations := ctn.Get("focus.configurations.actions.configurations").(*actions.Configurations)
			validator := ctn.Get("focus.validator").(services.Validator)
			return handlers.NewConfigurationsHandler(configurations, validator), nil
		},
	},
	{
		Name: "focus.configurations.handlers.options",
		Build: func(ctn di.Container) (interface{}, error) {
			options := ctn.Get("focus.configurations.actions.options").(*actions.Options)
			validator := ctn.Get("focus.validator").(services.Validator)
			return handlers.NewOptionsHandler(options, validator), nil
		},
	},
	{
		Name: "focus.configurations.router",
		Build: func(ctn di.Container) (interface{}, error) {
			confHandler := ctn.Get("focus.configurations.handlers.configurations").(*handlers.ConfigurationsHandler)
			optHandler := ctn.Get("focus.configurations.handlers.options").(*handlers.OptionsHandler)
			errorHandler := ctn.Get("focus.errorHandler").(services.ErrorHandler)
			return NewRouter(confHandler, optHandler, errorHandler), nil
		},
	},
}
