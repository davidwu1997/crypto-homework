package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type DepositCond struct {
	StartTime *time.Time
	EndTime   *time.Time

	MemberLogin *string
}

type Deposit struct {
	ID        uint64
	WalletID  uint64
	Amount    decimal.Decimal
	CreatedAt *time.Time
}
