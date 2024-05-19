package services_definitions

import (
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stocktacking_backend/internal/infrastructure/registry/services_definitions/translations"
	et "stocktacking_backend/pkg/services/error-translator"
	middleware "stocktacking_backend/pkg/services/gin_middleware"
	"stocktacking_backend/pkg/services/validation"
)

var PkgDefinitions = []di.Def{
	{
		Name: "focus.logger",
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get("logger").(*zap.SugaredLogger)
			return logger, nil
		},
	},
	{
		Name: "focus.validator",
		Build: func(ctn di.Container) (interface{}, error) {
			universalTranslator := ctn.Get("focus.universalTranslator").(*ut.UniversalTranslator)
			return validation.NewValidator(universalTranslator), nil
		},
	},
	{
		Name: "focus.errorHandler",
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get("logger").(*zap.SugaredLogger)
			errTrans := ctn.Get("focus.errorTranslator").(*et.Translator)
			errHandler := middleware.NewErrorHandler(logger).SetTranslator(errTrans)

			return errHandler, nil
		},
	},
	//{
	//	Name: "focus.pageHandler",
	//	Build: func(ctn di.Container) (interface{}, error) {
	//		logger := ctn.Get("logger").(*zap.SugaredLogger)
	//		errTrans := ctn.Get("focus.errorTranslator").(*et.Translator)
	//		errHandler := middleware.NewErrorHandler(logger).SetTranslator(errTrans)
	//
	//		return errHandler, nil
	//	},
	//},
	{
		Name: "focus.db",
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get("db").(*gorm.DB)
			return db, nil
		},
	},
	{
		Name: "focus.universalTranslator",
		Build: func(ctn di.Container) (interface{}, error) {
			russian := ru.New()
			utTranslator := ut.New(russian, russian)

			return utTranslator, nil
		},
	},
	{
		Name: "focus.errorTranslator",
		Build: func(ctn di.Container) (interface{}, error) {
			uTranslator := ctn.Get("focus.universalTranslator").(*ut.UniversalTranslator)
			ru, _ := uTranslator.GetTranslator("ru")
			etTranslator := et.New(ru)

			if err := etTranslator.AddTranslation(translations.ErrTranslations...); err != nil {
				return nil, err
			}

			return etTranslator, nil
		},
	},
	{
		Name: "focus.awsS3.client",
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get("awsS3.client"), nil
		},
	},
	{
		Name: "focus.awsS3.bucketName",
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get("awsS3.bucketName"), nil
		},
	},
	//{
	//	Name: "focus.callbacks.test",
	//	Build: func(ctn di.Container) (interface{}, error) {
	//		logger := ctn.Get("focus.logger").(*zap.SugaredLogger)
	//		notifier := ctn.Get("notifier.websockets").(notify.Notifier)
	//
	//		return func(plugin, entity string) callbacks.Callbacks {
	//			fun := func(callback string, ids ...uuid.UUID) {
	//				notification := map[string]any{"plugin": plugin, "entity": entity, "callback": callback, "ids": ids}
	//				if err := notifier.NotifyAll(notification); err != nil {
	//					logger.Warnw("error notify", "err", err, "notification", notification)
	//				}
	//			}
	//
	//			return callbacks.Callbacks{
	//				AfterCreate: func(ids ...uuid.UUID) { fun("afterCreate", ids...) },
	//				AfterUpdate: func(ids ...uuid.UUID) { fun("afterUpdate", ids...) },
	//				AfterDelete: func(ids ...uuid.UUID) { fun("afterDelete", ids...) },
	//			}
	//		}, nil
	//	},
	//},
}
