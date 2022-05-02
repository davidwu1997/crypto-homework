package database

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type DepositRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Deposit) error
	Find(ctx context.Context, db *gorm.DB, cond *bo.DepositCond) ([]*po.Deposit, error)
}

type depositRepository struct{}

func newDepositRepository() DepositRepositoryInterface {
	return &depositRepository{}
}

func (repo *depositRepository) Create(ctx context.Context, db *gorm.DB, data *po.Deposit) error {
	return db.Create(data).Error
}

func (repo *depositRepository) Find(ctx context.Context, db *gorm.DB, cond *bo.DepositCond) ([]*po.Deposit, error) {
	result := make([]*po.Deposit, 0)

	db = db.Model(po.Deposit{}).Joins("Inner Join `wallet` ON `wallet`.`id` = `deposit`.`wallet_id`")
	if cond.MemberLogin != nil {
		db = db.Where("`wallet`.`member_login` = ?", cond.MemberLogin)
	}

	if cond.StartTime != nil {
		db = db.Where("`deposit`.`created_at` >= ?", cond.StartTime)
	}

	if cond.EndTime != nil {
		db = db.Where("`deposit`.`created_at` = ?", cond.EndTime)
	}

	if err := db.Order("`deposit`.`created_at` desc").Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
