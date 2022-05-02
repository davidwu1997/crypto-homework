package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransferCond struct {
	StartTime *time.Time
	EndTime   *time.Time

	MemberLogin *string
}

type Transfer struct {
	ID           uint64
	FromWalletID uint64
	Amount       decimal.Decimal
	ToWalletID   uint64
	CreatedAt    *time.Time
}
