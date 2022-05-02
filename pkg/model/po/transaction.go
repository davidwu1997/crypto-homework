package po

import (
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID

	WalletID       uint64          `gorm:"<-:create;column:wallet_id"`
	Method         string          `gorm:"<-:create;column:method"`
	ActionID       uint64          `gorm:"<-:create;column:action_id"`
	TransferBefore decimal.Decimal `gorm:"<-:create;column:transfer_before"`
	TransferAmount decimal.Decimal `gorm:"<-:create;column:transfer_amount"`
	TransferAfter  decimal.Decimal `gorm:"<-:create;column:transfer_after"`

	CreatedAt
}

func (t *Transaction) TableName() string {
	return "transaction"
}
