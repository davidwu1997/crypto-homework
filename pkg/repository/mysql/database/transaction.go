package database

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Transaction) error
	Find(ctx context.Context, db *gorm.DB, cond *bo.TransactionCond) ([]*po.Transaction, error)
}

type transactionRepository struct{}

func newTransactionRepository() TransactionRepositoryInterface {
	return &transactionRepository{}
}

func (repo *transactionRepository) Create(ctx context.Context, db *gorm.DB, data *po.Transaction) error {
	return db.Create(data).Error
}

func (repo *transactionRepository) Find(ctx context.Context, db *gorm.DB, cond *bo.TransactionCond) ([]*po.Transaction, error) {
	result := make([]*po.Transaction, 0)

	db = db.Model(po.Transaction{}).Joins("Inner Join `wallet` ON `wallet`.`id` = `transaction`.`wallet_id`")
	if cond.MemberLogin != nil {
		db = db.Where("`wallet`.`member_login` = ?", cond.MemberLogin)
	}

	if cond.StartTime != nil {
		db = db.Where("`transaction`.`created_at` >= ?", cond.StartTime)
	}

	if cond.EndTime != nil {
		db = db.Where("`transaction`.`created_at` = ?", cond.EndTime)
	}

	if err := db.Order("`transaction`.`created_at` desc").Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
