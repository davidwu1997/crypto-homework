package po

import (
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID
	MemberLogin  string          `gorm:"<-:create;column:member_login"`
	CurrencyCode string          `gorm:"<-:create;column:currency_code"`
	Balance      decimal.Decimal `gorm:"column:balance; default:0"`

	CreatedAt
	UpdatedAt
}

func (t *Wallet) TableName() string {
	return "wallet"
}
