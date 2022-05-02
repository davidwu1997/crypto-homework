package database

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type WalletRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Wallet) error
	Updated(ctx context.Context, db *gorm.DB, data *po.Wallet) error

	Get(ctx context.Context, db *gorm.DB, id uint64) (*po.Wallet, error)
	FindByMember(ctx context.Context, db *gorm.DB, cond *bo.MemberWalletCond) ([]*po.Wallet, error)
}

type walletRepository struct{}

func newWalletRepository() WalletRepositoryInterface {
	return &walletRepository{}
}

func (repo *walletRepository) Create(ctx context.Context, db *gorm.DB, data *po.Wallet) error {
	return db.Create(data).Error
}

func (repo *walletRepository) Updated(ctx context.Context, db *gorm.DB, data *po.Wallet) error {

	return db.Updates(data).Error
}

func (repo *walletRepository) FindByMember(ctx context.Context, db *gorm.DB, cond *bo.MemberWalletCond) ([]*po.Wallet, error) {
	result := make([]*po.Wallet, 0)

	db = db.Table("wallet").Where("`member_login` = ?", cond.MemberLogin)

	if cond.CurrencyCode != nil {
		db = db.Where("`currency_code` = ?", cond.CurrencyCode)
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *walletRepository) Get(ctx context.Context, db *gorm.DB, id uint64) (*po.Wallet, error) {
	wallet := &po.Wallet{ID: po.ID{ID: id}}
	if err := db.First(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}
