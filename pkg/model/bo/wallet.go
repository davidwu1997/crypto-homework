package bo

import "github.com/shopspring/decimal"

type MemberWalletCond struct {
	MemberLogin  string
	CurrencyCode *string
}

type Wallet struct {
	ID           uint64
	MemberLogin  string
	CurrencyCode string
	Balance      decimal.Decimal
}
