package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stocktacking_backend/internal/adapters/rest"
	"stocktacking_backend/internal/infrastructure/env"
	"stocktacking_backend/internal/infrastructure/registry"
	"stocktacking_backend/internal/service/fixtures"
)

func main() {
	//if tracer, closer, err := tracer.SetJaegerTracer(env.TraceHeader); err == nil {
	//	opentracing.SetGlobalTracer(tracer)
	//	defer func(closer io.Closer) { _ = closer.Close() }(closer)
	//}

	// Инициализация контейнера зависимостей
	ctn, _ := registry.NewContainer()
	logger := ctn.Resolve("logger").(*zap.SugaredLogger)

	if err := migration(); err != nil {
		logger.Errorf("migration error: %s", err)
		return
	}

	server := ctn.Resolve("router").(*rest.Router)
	_ = ctn.Resolve("db").(*gorm.DB)
	router := server.Router()

	fixtureService := ctn.Resolve("fixturesService").(*fixtures.FixtureService)
	fixtureService.RunFixtures(context.Background())

	//handler := cors.Default().Handler(router)
	err := router.Run(env.HTTPPort)
	if err != nil {
		logger.Fatal(err)
	}
	_ = ctn.Clean()
}

func migration() error {
	dbDsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s&sslrootcert=%s",
		env.DbDriver,
		env.DbUser,
		env.DbPassword,
		env.DbHost,
		env.DbPort,
		env.DbName,
		env.DbSslMode,
		env.DbSslCertPath,
	)

	db, err := sql.Open(env.DbDriver, dbDsn)
	if err != nil {
		return fmt.Errorf("db init error: %w", err)
	}
	defer func() {
		_ = db.Close()
	}()

	if err := goose.SetDialect(env.DbDriver); err != nil {
		return fmt.Errorf("goose init error: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migrations up error: %w", err)
	}

	return nil
}
