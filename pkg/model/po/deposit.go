package po

import (
	"github.com/shopspring/decimal"
)

type Deposit struct {
	ID
	WalletID uint64 `gorm:"<-:create;column:wallet_id"`

	TransferAmount decimal.Decimal `gorm:"<-:create;column:transfer_amount"`

	CreatedAt
	UpdatedAt
}

func (t *Deposit) TableName() string {
	return "deposit"
}
