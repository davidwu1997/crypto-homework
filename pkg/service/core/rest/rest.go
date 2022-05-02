package rest

import (
	"crypto/pkg/util/logger"

	repo "crypto/pkg/repository/mysql/database"
)

func NewUseCase(db *repo.DBRepository) *UseCase {
	return &UseCase{
		MemberUseCase:   newMemberUseCase(db, logger.SysLog()),
		CurrencyUseCase: newCurrencyUseCase(db, logger.SysLog()),
	}
}

//
//type Repository struct {
//	DBClient           *repo.DBClient
//	WalletRepository   repo.WalletRepositoryInterface
//	MemberRepository   repo.MemberRepositoryInterface
//	CurrencyRepository repo.CurrencyRepositoryInterface
//
//	DepositRepository     repo.DepositRepositoryInterface
//	WithdrawRepository    repo.WithdrawRepositoryInterface
//	TransferRepository    repo.TransferRepositoryInterface
//	TransactionRepository repo.TransactionRepositoryInterface
//}

type UseCase struct {
	MemberUseCase   MemberUseCaseInterface
	CurrencyUseCase CurrencyUseCaseInterface
}
