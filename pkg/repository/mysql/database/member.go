package database

import (
	"context"
	"crypto/pkg/model/po"

	"gorm.io/gorm"
)

type MemberRepositoryInterface interface {
	Create(ctx context.Context, db *gorm.DB, data *po.Member) error
}

type memberRepository struct{}

func newMemberRepository() MemberRepositoryInterface {
	return &memberRepository{}
}

func (repo *memberRepository) Create(ctx context.Context, db *gorm.DB, data *po.Member) error {
	return db.Create(data).Error
}
