package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"stocktacking_backend/internal/adapters/postgres"
	"stocktacking_backend/internal/adapters/rest"
	"stocktacking_backend/internal/adapters/rest/handlers"
	"stocktacking_backend/internal/infrastructure/env"
	"stocktacking_backend/internal/infrastructure/registry/services_definitions"
	"stocktacking_backend/internal/service"
	"stocktacking_backend/internal/service/fixtures"
	rest2 "stocktacking_backend/pkg/configurations/rest"
	"time"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err = builder.Add(definitions...); err != nil {
		return nil, err
	}
	if err = builder.Add(services_definitions.ConfigurationsDefinitions...); err != nil {
		return nil, err
	}
	//if err = builder.Add(services_definitions.RestHandlerDefinitions...); err != nil {
	//	return nil, err
	//}
	//if err = builder.Add(services_definitions.UseCaseDefinitions...); err != nil {
	//	return nil, err
	//}
	//if err = builder.Add(services_definitions.RepositoryDefinitions...); err != nil {
	//	return nil, err
	//}
	//
	//if err = builder.Add(services_definitions.UseCaseDefinitions...); err != nil {
	//	return nil, err
	//}
	//if err = builder.Add(services_definitions.RepositoryDefinitions...); err != nil {
	//	return nil, err
	//}
	//if err = builder.Add(services_definitions.RestHandlerDefinitions...); err != nil {
	//	return nil, err
	//}
	//
	//if err = builder.Add(services_definitions.MediaDefinitions...); err != nil {
	//	return nil, err
	//}
	//
	//if err = builder.Add(services_definitions.MenuDefinitions...); err != nil {
	//	return nil, err
	//}
	if err = builder.Add(services_definitions.PkgDefinitions...); err != nil {
		return nil, err
	}

	//if err = builder.Add(services_definitions.MailTemplatesDefinitions...); err != nil {
	//	return nil, err
	//}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

var definitions = []di.Def{
	{
		Name: "logger",
		Build: func(ctn di.Container) (interface{}, error) {
			rawJSON := []byte(`{
				  "level": "` + env.LogLevel + `",
				  "encoding": "json",
				  "outputPaths": ["stdout"],
				  "errorOutputPaths": ["stderr"],
				  "encoderConfig": {
					"messageKey": "message",
					"levelKey": "level",
					"levelEncoder": "lowercase"
				  }
				}`)
			var cfg zap.Config
			if err := json.Unmarshal(rawJSON, &cfg); err != nil {
				return nil, err
			}
			logger, err := cfg.Build()
			if err != nil {
				return nil, err
			}
			logger.Debug("logger construction succeeded")
			sugar := logger.Sugar()
			return sugar, nil
		},
	},
	{
		Name: "copier_service",
		Build: func(ctn di.Container) (interface{}, error) {
			return service.Copier{}, nil
		},
	},
	{
		Name: "focus.db",
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get("db").(*gorm.DB)
			return db, nil
		},
	},
	{
		Name: "db", // Database
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get("logger").(*zap.SugaredLogger)

			conn := fmt.Sprintf(
				"host=%s port=%s dbname=%s user=%s sslmode=%s",
				env.DbHost, env.DbPort, env.DbName, env.DbUser, env.DbSslMode,
			)

			c1 := make(chan *gorm.DB, 1)
			go func() {
				// Получение лог-сервиса из контейнера
				logger.Info("Try to connect to database")

				_, err := os.ReadFile(env.DbSslCertPath)
				if env.DbSslMode != "disable" && env.DbSslCertPath != "" && err != nil {
					logger.Fatal("cannot open cert file")
				}

				// dsn
				dsn := conn + " sslrootcert=" + env.DbSslCertPath + " password=" + env.DbPassword
				dbConn, err := postgres.NewDatabase(dsn)
				if err != nil {
					logger.Error(fmt.Sprintf("DB error connection. Details: %s", err))
					logger.Fatal(fmt.Sprintf("DB error connection. Connection data: %s", conn))
					c1 <- nil
					return
				} else if dbConn != nil {
					logger.Info("DB connection has  been successfully created")
				}
				c1 <- dbConn
			}()

			// Listen on our channel AND a timeout channel - which ever happens first.
			select {
			case db := <-c1:
				_ = db.AutoMigrate(
				// todo
				)
				return db, nil
			case <-time.After(30 * time.Second):
				err := errors.New("5 seconds timeout error")
				logger.Error(fmt.Sprintf("DB error connection. Details: %s", err))
				logger.Error(fmt.Sprintf("DB error connection. Connection data: %s", conn))
				return nil, err
			}
		},
	},
	{
		Name: "router",
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get("logger").(*zap.SugaredLogger)
			fixturesHandler := ctn.Get("fixturesHandler").(*handlers.FixturesHandler)
			configurationsRouter := ctn.Get("focus.configurations.router").(*rest2.Router)
			logger.Info("building router")

			router := rest.NewRouter(
				env.GinMode,
				rest.APISettings{
					Root:        env.HTTPApiRoot,
					Version:     env.HTTPApiVersion,
					ServiceName: env.HTTPApiServiceName,
					PkgPath:     env.HTTPApiPkgPath,
				},
				fixturesHandler,
				configurationsRouter,
			)
			logger.Info("router has built")
			return router, nil
		},
	},
	{
		Name: "fixtures",
		Build: func(ctn di.Container) (interface{}, error) {
			return []fixtures.Fixture{
				//fixtures.ExternalTagFixture{},
				//fixtures.InternalTagFixture{},

				fixtures.ConfigurationFixture{},
				fixtures.OrganizationFixture{},
				fixtures.BranchFixture{},
				fixtures.CategoryFixture{},
				fixtures.OptionFixture{},
				fixtures.UserFixture{},
				fixtures.StorageFixture{},
				fixtures.ItemFixture{},
				fixtures.ReportFixture{},
			}, nil
		},
	},
	{
		Name: "fixturesService",
		Build: func(ctn di.Container) (interface{}, error) {
			f := ctn.Get("fixtures").([]fixtures.Fixture)
			fr := ctn.Get("fixtureRepository").(fixtures.FixtureRepository)
			return fixtures.NewFixtureService(f, fr), nil
		},
	},
	{
		Name: "fixtureRepository",
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get("db").(*gorm.DB)
			return postgres.NewFixtureRepository(db), nil
		},
	},
	{
		Name: "fixturesHandler",
		Build: func(ctn di.Container) (interface{}, error) {
			fixtureService := ctn.Get("fixturesService").(*fixtures.FixtureService)
			return handlers.NewFixturesHandler(fixtureService), nil
		},
	},
}
