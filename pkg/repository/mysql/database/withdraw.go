package database

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type WithdrawRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Withdraw) error
	Find(ctx context.Context, db *gorm.DB, cond *bo.WithdrawCond) ([]*po.Withdraw, error)
}

type withdrawRepository struct {
}

func newWithdrawRepository() WithdrawRepositoryInterface {
	return &withdrawRepository{}
}

func (repo *withdrawRepository) Create(ctx context.Context, db *gorm.DB, data *po.Withdraw) error {
	return db.Create(data).Error
}

func (repo *withdrawRepository) Find(ctx context.Context, db *gorm.DB, cond *bo.WithdrawCond) ([]*po.Withdraw, error) {
	result := make([]*po.Withdraw, 0)

	db = db.Model(po.Withdraw{}).Joins("Inner Join `wallet` ON `wallet`.`id` = `withdraw`.`wallet_id`")
	if cond.MemberLogin != nil {
		db = db.Where("`wallet`.`member_login` = ?", cond.MemberLogin)
	}

	if cond.StartTime != nil {
		db = db.Where("`withdraw`.`created_at` >= ?", cond.StartTime)
	}

	if cond.EndTime != nil {
		db = db.Where("`withdraw`.`created_at` = ?", cond.EndTime)
	}

	if err := db.Order("`withdraw`.`created_at` desc").Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
