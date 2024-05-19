package plugin

import (
	"github.com/sarulabs/di/v2"
	"stocktacking_backend/pkg/configurations/plugin/actions"
)

var Definitions = []di.Def{
	{
		Name: "focus.configurations.actions.configurations",
		Build: func(ctn di.Container) (interface{}, error) {
			confRepository := ctn.Get("focus.configurations.repositories.configurations").(actions.ConfigurationsRepository)

			return actions.NewConfigurations(confRepository), nil
		},
	},
	{
		Name: "focus.configurations.actions.options",
		Build: func(ctn di.Container) (interface{}, error) {
			confRepository := ctn.Get("focus.configurations.repositories.configurations").(actions.ConfigurationsRepository)
			optRepository := ctn.Get("focus.configurations.repositories.options").(actions.OptionsRepository)

			var mediaService actions.MediaService
			if mediaServiceInterface, err := ctn.SafeGet("focus.media.actions.media"); err == nil {
				mediaService = mediaServiceInterface.(actions.MediaService)
			}

			return actions.NewOptions(confRepository, optRepository, mediaService), nil
		},
	},
}
