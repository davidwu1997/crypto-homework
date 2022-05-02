package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type WithdrawCond struct {
	StartTime *time.Time
	EndTime   *time.Time

	MemberLogin *string
}

type Withdraw struct {
	ID        uint64
	WalletID  uint64
	Amount    decimal.Decimal
	CreatedAt *time.Time
}
