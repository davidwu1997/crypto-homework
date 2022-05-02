package po

import (
	"github.com/shopspring/decimal"
)

type Transfer struct {
	ID
	FromWalletID   uint64          `gorm:"<-:create;column:from_wallet_id"`
	TransferAmount decimal.Decimal `gorm:"<-:create;column:transfer_amount"`
	ToWalletID     uint64          `gorm:"<-:create;column:to_wallet_id"`

	CreatedAt
	UpdatedAt 
}

func (t *Transfer) TableName() string {
	return "transfer"
}
