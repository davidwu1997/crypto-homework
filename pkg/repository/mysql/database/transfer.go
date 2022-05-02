package database

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type TransferRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Transfer) error
	Find(ctx context.Context, db *gorm.DB, cond *bo.TransferCond) ([]*po.Transfer, error)
}

type transferRepository struct{}

func newTransferRepository() TransferRepositoryInterface {
	return &transferRepository{}
}

func (repo *transferRepository) Create(ctx context.Context, db *gorm.DB, data *po.Transfer) error {
	return db.Create(data).Error
}

func (repo *transferRepository) Find(ctx context.Context, db *gorm.DB, cond *bo.TransferCond) ([]*po.Transfer, error) {
	result := make([]*po.Transfer, 0)

	db = db.Model(po.Transfer{}).Joins("Inner Join `wallet` ON `wallet`.`id` = `transfer`.`from_wallet_id`")
	if cond.MemberLogin != nil {
		db = db.Where("`wallet`.`member_login` = ?", cond.MemberLogin)
	}

	if cond.StartTime != nil {
		db = db.Where("`transfer`.`created_at` >= ?", cond.StartTime)
	}

	if cond.EndTime != nil {
		db = db.Where("`transfer`.`created_at` = ?", cond.EndTime)
	}

	if err := db.Order("`transfer`.`created_at` desc").Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
