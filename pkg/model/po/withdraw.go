package po

import (
	"github.com/shopspring/decimal"
)

type Withdraw struct {
	ID
	WalletID       uint64          `gorm:"<-:create;column:wallet_id"`
	TransferAmount decimal.Decimal `gorm:"<-:create;column:transfer_amount"`

	CreatedAt
	UpdatedAt
}

func (t *Withdraw) TableName() string {
	return "withdraw"
}
