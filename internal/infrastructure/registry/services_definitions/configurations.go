package services_definitions

import (
	"github.com/sarulabs/di/v2"
	"stocktacking_backend/pkg/configurations/plugin"
	"stocktacking_backend/pkg/configurations/plugin/actions"
	"stocktacking_backend/pkg/configurations/postgres"
	"stocktacking_backend/pkg/configurations/rest"
	"stocktacking_backend/pkg/configurations/rest/handlers"
	"stocktacking_backend/pkg/services/validation"
)

var ConfigurationsDefinitions = appendArr(
	[]di.Def{
		{
			Name: "optionsHandler",
			Build: func(ctn di.Container) (interface{}, error) {
				options := ctn.Get("focus.configurations.actions.options").(*actions.Options)
				validator := ctn.Get("focus.validator").(*validation.Validator)
				return handlers.NewOptionsHandler(options, validator), nil
			},
		},
	}, plugin.Definitions, postgres.Definitions, rest.Definitions,
)

func appendArr[T any](defs ...[]T) []T {
	var res []T
	for _, def := range defs {
		res = append(res, def...)
	}

	return res
}
