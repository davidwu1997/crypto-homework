package dto

import (
	"github.com/shopspring/decimal"
	"time"
)

type CreateMemberReq struct {
	Login    string `json:"login"`
	PassWord string `json:"pass_word"`
	Name     string `json:"name"`
}

type MemberWalletReq struct {
	MemberLogin  string  `form:"member_login"`
	CurrencyCode *string `form:"currency_code"`
}

type MemberWalletResp struct {
	WalletID     string          `json:"wallet_id"`
	CurrencyCode string          `json:"currency_code"`
	Balance      decimal.Decimal `json:"balance"`
}

type MemberDepositReq struct {
	WalletID       string          `json:"wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
}

type MemberWithdrawReq struct {
	WalletID       string          `json:"wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
}

type MemberTransferReq struct {
	FromWalletID   string          `json:"from_wallet_id"`
	ToWalletID     string          `json:"to_wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
}

type MemberDepositLogReq struct {
	MemberLogin string     `form:"member_login"`
	StartTime   *time.Time `from:"start_time"`
	EndTime     *time.Time `from:"end_time"`
}

type MemberWithdrawLogReq struct {
	MemberLogin string     `form:"member_login"`
	StartTime   *time.Time `from:"start_time"`
	EndTime     *time.Time `from:"end_time"`
}

type MemberTransferLogReq struct {
	MemberLogin string     `form:"member_login"`
	StartTime   *time.Time `from:"start_time"`
	EndTime     *time.Time `from:"end_time"`
}

type MemberTransactionLogReq struct {
	MemberLogin string     `form:"member_login"`
	StartTime   *time.Time `from:"start_time"`
	EndTime     *time.Time `from:"end_time"`
}

type MemberDepositLogResp struct {
	ID             string          `json:"id"`
	WalletID       string          `json:"wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
	CreatedAt      *time.Time      `json:"created_at"`
}
type MemberWithdrawLogResp struct {
	ID             string          `json:"id"`
	WalletID       string          `json:"wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
	CreatedAt      *time.Time      `json:"created_at"`
}
type MemberTransferLogResp struct {
	ID             string          `json:"id"`
	FromWalletID   string          `json:"from_wallet_id"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
	ToWalletID     string          `json:"to_wallet_id"`
	CreatedAt      *time.Time      `json:"created_at"`
}
type MemberTransactionLogResp struct {
	ID             string          `json:"id"`
	WalletID       string          `json:"wallet_id"`
	Method         string          `json:"method"`
	ActionID       string          `json:"action_id"`
	TransferBefore decimal.Decimal `json:"transfer_before"`
	TransferAmount decimal.Decimal `json:"transfer_amount"`
	TransferAfter  decimal.Decimal `json:"transfer_after"`
	CreatedAt      *time.Time      `json:"created_at"`
}
