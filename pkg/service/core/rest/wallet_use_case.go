package rest

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/repository/mysql/database"
	"crypto/pkg/util/logger"
)

type WalletUseCaseInterface interface {
	GetListByMember(ctx context.Context, cond *bo.MemberWalletCond) ([]*bo.Wallet, error)
}

func newWalletUseCase(in *database.DBRepository, logger logger.LoggerInterface) WalletUseCaseInterface {
	return &walletUseCase{
		in:     in,
		logger: logger,
	}
}

type walletUseCase struct {
	logger logger.LoggerInterface

	in *database.DBRepository
}

func (uc *walletUseCase) GetListByMember(ctx context.Context, cond *bo.MemberWalletCond) ([]*bo.Wallet, error) {
	result := []*bo.Wallet{}

	wallets, err := uc.in.WalletRepository.FindByMember(ctx, uc.in.DBClient.Session(), cond)
	if err != nil {
		return result, err
	}

	for _, wallet := range wallets {
		result = append(result, &bo.Wallet{
			MemberLogin:  wallet.MemberLogin,
			CurrencyCode: wallet.CurrencyCode,
			Balance:      wallet.Balance,
		})
	}

	return result, nil
}
