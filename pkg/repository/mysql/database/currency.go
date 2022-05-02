package database

import (
	"context"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type CurrencyRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Currency) error
	Find(ctx context.Context, db *gorm.DB) ([]*po.Currency, error)
}

type currencyRepository struct{}

func newCurrencyRepository() CurrencyRepositoryInterface {
	return &currencyRepository{}
}

func (repo *currencyRepository) Create(ctx context.Context, db *gorm.DB, data *po.Currency) error {
	return db.Create(data).Error
}

func (repo *currencyRepository) Find(ctx context.Context, db *gorm.DB) ([]*po.Currency, error) {
	result := make([]*po.Currency, 0)

	if err := db.Model(po.Currency{}).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
