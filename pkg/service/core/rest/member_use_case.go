package rest

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"
	"crypto/pkg/repository/mysql/database"
	"crypto/pkg/service/constant"
	"crypto/pkg/util/logger"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MemberUseCaseInterface interface {
	Create(ctx context.Context, data *bo.Member) error

	GetWalletList(ctx context.Context, cond *bo.MemberWalletCond) ([]*bo.Wallet, error)

	Deposit(ctx context.Context, data *bo.MemberDeposit) error
	Withdraw(ctx context.Context, data *bo.MemberWithdraw) error
	Transfer(ctx context.Context, data *bo.MemberTransfer) error

	DepositLog(ctx context.Context, cond *bo.DepositCond) ([]*bo.Deposit, error)
	WithdrawLog(ctx context.Context, cond *bo.WithdrawCond) ([]*bo.Withdraw, error)
	TransferLog(ctx context.Context, cond *bo.TransferCond) ([]*bo.Transfer, error)
	TransactionLog(ctx context.Context, cond *bo.TransactionCond) ([]*bo.Transaction, error)
}

func newMemberUseCase(in *database.DBRepository, logger logger.LoggerInterface) MemberUseCaseInterface {
	return &memberUseCase{
		db:     in,
		logger: logger,
	}
}

type memberUseCase struct {
	logger logger.LoggerInterface
	db     *database.DBRepository
}

func (uc *memberUseCase) Create(ctx context.Context, data *bo.Member) error {
	db := uc.db.DBClient.Session()

	err := uc.db.MemberRepository.Create(ctx, db, &po.Member{Login: data.Login, PassWord: data.PassWord, Name: data.Name})
	if err != nil {
		return err
	}

	return nil
}

func (uc *memberUseCase) GetWalletList(ctx context.Context, cond *bo.MemberWalletCond) ([]*bo.Wallet, error) {
	result := []*bo.Wallet{}
	db := uc.db.DBClient.Session()

	currencies, err := uc.db.CurrencyRepository.Find(ctx, db)
	wallets, err := uc.db.WalletRepository.FindByMember(ctx, db, cond)
	if err != nil {
		return result, err
	}

	currenciesMap := map[string]bool{}
	for _, currency := range currencies {
		currenciesMap[currency.Code] = false
	}

	for _, wallet := range wallets {
		result = append(result, &bo.Wallet{
			ID:           wallet.ID.ID,
			MemberLogin:  wallet.MemberLogin,
			CurrencyCode: wallet.CurrencyCode,
			Balance:      wallet.Balance,
		})

		currenciesMap[wallet.CurrencyCode] = true
	}

	for currencyCode, isExist := range currenciesMap {
		if !isExist {
			wallet := &po.Wallet{MemberLogin: cond.MemberLogin, CurrencyCode: currencyCode}
			if err = uc.db.WalletRepository.Create(ctx, db, wallet); err != nil {
				return result, err
			}

			result = append(result, &bo.Wallet{
				ID:           wallet.ID.ID,
				MemberLogin:  wallet.MemberLogin,
				CurrencyCode: wallet.CurrencyCode,
				Balance:      wallet.Balance,
			})
		}
	}

	return result, nil
}

func (uc *memberUseCase) Deposit(ctx context.Context, data *bo.MemberDeposit) error {
	db := uc.db.DBClient.Session()

	return db.Transaction(func(tx *gorm.DB) error {
		wallet, err := uc.db.WalletRepository.Get(ctx, tx, data.WalletID)
		if err != nil {
			return err
		}
		transferBefore := wallet.Balance
		wallet.Balance = wallet.Balance.Add(data.TransferAmount)
		transferAfter := wallet.Balance

		if err = uc.db.WalletRepository.Updated(ctx, tx, wallet); err != nil {
			return err
		}

		deposit := &po.Deposit{WalletID: data.WalletID, TransferAmount: data.TransferAmount}
		if err = uc.db.DepositRepository.Create(ctx, uc.db.DBClient.Session(), deposit); err != nil {
			return err
		}

		if err = uc.db.TransactionRepository.Create(ctx, tx, &po.Transaction{
			WalletID:       data.WalletID,
			Method:         "deposit",
			ActionID:       deposit.ID.ID,
			TransferBefore: transferBefore,
			TransferAmount: data.TransferAmount,
			TransferAfter:  transferAfter,
		}); err != nil {
			return err
		}

		return nil
	})

}

func (uc *memberUseCase) Withdraw(ctx context.Context, data *bo.MemberWithdraw) error {

	db := uc.db.DBClient.Session()

	return db.Transaction(func(tx *gorm.DB) error {
		wallet, err := uc.db.WalletRepository.Get(ctx, tx, data.WalletID)
		if err != nil {
			return err
		}
		transferBefore := wallet.Balance
		wallet.Balance = wallet.Balance.Sub(data.TransferAmount)
		transferAfter := wallet.Balance

		if transferAfter.LessThan(decimal.NewFromFloat(0.00)) {
			return constant.BalanceInsufficient_Error
		}

		if err = uc.db.WalletRepository.Updated(ctx, tx, wallet); err != nil {
			return err
		}

		withdraw := &po.Withdraw{WalletID: data.WalletID, TransferAmount: data.TransferAmount}
		if err = uc.db.WithdrawRepository.Create(ctx, uc.db.DBClient.Session(), withdraw); err != nil {
			return err
		}

		if err = uc.db.TransactionRepository.Create(ctx, tx, &po.Transaction{
			WalletID:       data.WalletID,
			Method:         "withdraw",
			ActionID:       withdraw.ID.ID,
			TransferBefore: transferBefore,
			TransferAmount: data.TransferAmount,
			TransferAfter:  transferAfter,
		}); err != nil {
			return err
		}

		return nil
	})
}

func (uc *memberUseCase) Transfer(ctx context.Context, data *bo.MemberTransfer) error {

	db := uc.db.DBClient.Session()

	return db.Transaction(func(tx *gorm.DB) error {
		fromWallet, err := uc.db.WalletRepository.Get(ctx, tx, data.FromWalletID)
		if err != nil {
			return err
		}

		toWallet, err := uc.db.WalletRepository.Get(ctx, tx, data.ToWalletID)
		if err != nil {
			return err
		}

		//如果幣別不同則無法轉換
		if fromWallet.CurrencyCode != toWallet.CurrencyCode {
			return constant.WalletTransfer_CurrencyDiff_Error
		}

		fromWalletTransferBefore := fromWallet.Balance
		fromWallet.Balance = fromWallet.Balance.Sub(data.TransferAmount)
		fromWalletTransferAfter := fromWallet.Balance

		toWalletTransferBefore := toWallet.Balance
		toWallet.Balance = toWallet.Balance.Add(data.TransferAmount)
		toWalletTransferAfter := toWallet.Balance

		// 如果扣款後小於0，則失敗
		if fromWalletTransferAfter.LessThan(decimal.NewFromFloat(0.00)) {
			return constant.BalanceInsufficient_Error
		}

		// 更新來源錢包餘額
		if err = uc.db.WalletRepository.Updated(ctx, tx, fromWallet); err != nil {
			return err
		}

		// 更新目標錢包
		if err = uc.db.WalletRepository.Updated(ctx, tx, toWallet); err != nil {
			return err
		}

		// 建立紀錄
		transfer := &po.Transfer{FromWalletID: data.FromWalletID, TransferAmount: data.TransferAmount, ToWalletID: data.ToWalletID}
		if err = uc.db.TransferRepository.Create(ctx, tx, transfer); err != nil {
			return err
		}

		// 建立來源錢包流水紀錄
		if err = uc.db.TransactionRepository.Create(ctx, tx, &po.Transaction{
			WalletID:       data.FromWalletID,
			Method:         "transfer",
			ActionID:       transfer.ID.ID,
			TransferBefore: fromWalletTransferBefore,
			TransferAmount: data.TransferAmount,
			TransferAfter:  fromWalletTransferAfter,
		}); err != nil {
			return err
		}

		// 建立目標錢包流水紀錄
		if err = uc.db.TransactionRepository.Create(ctx, tx, &po.Transaction{
			WalletID:       data.ToWalletID,
			Method:         "transfer",
			ActionID:       transfer.ID.ID,
			TransferBefore: toWalletTransferBefore,
			TransferAmount: data.TransferAmount,
			TransferAfter:  toWalletTransferAfter,
		}); err != nil {
			return err
		}

		return nil
	})
}

func (uc *memberUseCase) DepositLog(ctx context.Context, cond *bo.DepositCond) ([]*bo.Deposit, error) {
	result := make([]*bo.Deposit, 0)
	db := uc.db.DBClient.Session()

	data, err := uc.db.DepositRepository.Find(ctx, db, cond)
	if err != nil {
		return result, err
	}

	for _, item := range data {
		result = append(result, &bo.Deposit{
			ID:        item.ID.ID,
			WalletID:  item.WalletID,
			Amount:    item.TransferAmount,
			CreatedAt: item.CreatedAt.CreatedAt,
		})
	}

	return result, nil
}

func (uc *memberUseCase) WithdrawLog(ctx context.Context, cond *bo.WithdrawCond) ([]*bo.Withdraw, error) {
	result := make([]*bo.Withdraw, 0)
	db := uc.db.DBClient.Session()

	data, err := uc.db.WithdrawRepository.Find(ctx, db, cond)
	if err != nil {
		return result, err
	}

	for _, item := range data {
		result = append(result, &bo.Withdraw{
			ID:        item.ID.ID,
			WalletID:  item.WalletID,
			Amount:    item.TransferAmount,
			CreatedAt: item.CreatedAt.CreatedAt,
		})
	}

	return result, nil
}

func (uc *memberUseCase) TransferLog(ctx context.Context, cond *bo.TransferCond) ([]*bo.Transfer, error) {
	result := make([]*bo.Transfer, 0)
	db := uc.db.DBClient.Session()

	data, err := uc.db.TransferRepository.Find(ctx, db, cond)
	if err != nil {
		return result, err
	}

	for _, item := range data {
		result = append(result, &bo.Transfer{
			ID:           item.ID.ID,
			FromWalletID: item.FromWalletID,
			Amount:       item.TransferAmount,
			ToWalletID:   item.ToWalletID,
			CreatedAt:    item.CreatedAt.CreatedAt,
		})
	}

	return result, nil
}

func (uc *memberUseCase) TransactionLog(ctx context.Context, cond *bo.TransactionCond) ([]*bo.Transaction, error) {
	result := make([]*bo.Transaction, 0)
	db := uc.db.DBClient.Session()

	data, err := uc.db.TransactionRepository.Find(ctx, db, cond)
	if err != nil {
		return result, err
	}

	for _, item := range data {
		result = append(result, &bo.Transaction{
			ID:             item.ID.ID,
			WalletID:       item.WalletID,
			Method:         item.Method,
			ActionID:       item.ActionID,
			TransferBefore: item.TransferBefore,
			TransferAmount: item.TransferAmount,
			TransferAfter:  item.TransferAfter,
			CreatedAt:      item.CreatedAt.CreatedAt,
		})
	}

	return result, nil
}
