package postgres

import (
	"context"
	"gorm.io/gorm"
)

type FixtureRepository struct {
	db *gorm.DB
}

func NewFixtureRepository(db *gorm.DB) *FixtureRepository {
	return &FixtureRepository{db: db}
}

func (r FixtureRepository) Create(ctx context.Context, entity any) error {
	//switch entity.(type) {
	//case []product_entity.ProductsMedia:
	//	err := r.db.WithContext(ctx).Table("products_media").Create(entity).Error
	//	return err
	//case []user_entity.Reviews:
	//	err := r.db.WithContext(ctx).Table("reviews").Create(entity).Error
	//	return err
	//case []shop_entity.ShopsMedia:
	//	err := r.db.WithContext(ctx).Table("shops_media").Create(entity).Error
	//	return err
	//case []product_entity.ProductItem:
	//	err := r.db.WithContext(ctx).Table("product_items").Create(entity).Error
	//	return err
	//}
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r FixtureRepository) DeleteAll(ctx context.Context, entity any) error {

	return r.db.WithContext(ctx).Where("TRUE").Delete(entity).Error
}
