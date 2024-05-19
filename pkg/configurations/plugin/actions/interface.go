package actions

import (
	"context"
	"stocktacking_backend/pkg/configurations/plugin/entity"
)

// ConfigurationsRepository интерфейс репозитория конфигураций
type ConfigurationsRepository interface {
	Get(ctx context.Context, id int) (conf *entity.Configuration, err error)
	Has(ctx context.Context, id int) bool
	HasByCode(ctx context.Context, code string) bool
	Create(ctx context.Context, conf ...entity.Configuration) error
	Update(ctx context.Context, conf *entity.Configuration) error
	FindByCode(ctx context.Context, code string) (conf *entity.Configuration, err error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, dto ConfigurationsListFilter) ([]entity.Configuration, error)
	Count(ctx context.Context, filter ConfigurationsFilter) (int, error)
}

type ConfigurationsListFilter struct {
	Offset int
	Limit  int
	Sort   string
	Order  string
	Filter ConfigurationsFilter
}

type ConfigurationsFilter struct {
	Query string
}

// OptionsRepository интерфейс репозитория настроек
type OptionsRepository interface {
	Has(ctx context.Context, confId int, optId int) bool
	HasByCode(ctx context.Context, confId int, optCodes ...string) bool
	Get(ctx context.Context, id int) (*entity.Option, error)
	Create(ctx context.Context, opt ...entity.Option) error
	Update(ctx context.Context, opt ...entity.Option) error
	Delete(ctx context.Context, optId int) error
	List(ctx context.Context, filter OptionsListFilter) ([]entity.Option, error)
	ListPreviews(ctx context.Context, filter OptionsListFilter) (optShorts []entity.OptionShort, err error)
	Count(ctx context.Context, filter OptionsFilter) (int, error)
}

type OptionsListFilter struct {
	Offset int
	Limit  int
	Sort   string
	Order  string
	Filter OptionsFilter
}

type OptionsFilter struct {
	ConfId   int
	ConfCode string
	OptCodes []string
}

// MediaService интерфейс сервиса медиа
type MediaService interface {
	CheckIds(ctx context.Context, ids ...int) error
}
