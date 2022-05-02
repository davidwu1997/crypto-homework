package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransactionCond struct {
	StartTime *time.Time
	EndTime   *time.Time

	MemberLogin *string
}

type Transaction struct {
	ID uint64

	WalletID       uint64
	Method         string
	ActionID       uint64
	TransferBefore decimal.Decimal
	TransferAmount decimal.Decimal
	TransferAfter  decimal.Decimal

	CreatedAt *time.Time
}
