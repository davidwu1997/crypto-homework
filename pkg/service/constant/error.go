package constant

import "golang.org/x/xerrors"

var (
	BalanceInsufficient_Error         = xerrors.New("餘額不足")
	WalletTransfer_CurrencyDiff_Error = xerrors.New("額度轉換幣別不同，請重新輸入")
)
