package rest

import (
	"context"
	"crypto/pkg/model/bo"
	"crypto/pkg/model/po"
	"crypto/pkg/repository/mysql/database"
	"crypto/pkg/util/logger"
)

type CurrencyUseCaseInterface interface {
	Create(ctx context.Context, data *bo.Currency) error
}

func newCurrencyUseCase(db *database.DBRepository, logger logger.LoggerInterface) CurrencyUseCaseInterface {
	return &currencyUseCase{
		db:     db,
		logger: logger,
	}
}

type currencyUseCase struct {
	logger logger.LoggerInterface
	db     *database.DBRepository
}

func (uc *currencyUseCase) Create(ctx context.Context, data *bo.Currency) error {

	err := uc.db.CurrencyRepository.Create(ctx, uc.db.DBClient.Session(), &po.Currency{Code: data.Code, Name: data.Name})
	if err != nil {
		return err
	}

	return nil
}
