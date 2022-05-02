package bo

import "github.com/shopspring/decimal"

type MemberCond struct {
	Login string `json:"login"`
}

type Member struct {
	Login    string
	PassWord string
	Name     string
}

type MemberDeposit struct {
	WalletID       uint64
	TransferAmount decimal.Decimal
}

type MemberWithdraw struct {
	WalletID       uint64
	TransferAmount decimal.Decimal
}

type MemberTransfer struct {
	FromWalletID   uint64
	ToWalletID     uint64
	TransferAmount decimal.Decimal
}
