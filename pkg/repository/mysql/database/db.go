package database

func NewDBRepository(dbClient *DBClient) *DBRepository {
	return &DBRepository{
		DBClient:           dbClient,
		MemberRepository:   newMemberRepository(),
		WalletRepository:   newWalletRepository(),
		CurrencyRepository: newCurrencyRepository(),

		DepositRepository:     newDepositRepository(),
		WithdrawRepository:    newWithdrawRepository(),
		TransferRepository:    newTransferRepository(),
		TransactionRepository: newTransactionRepository(),
	}
}

type DBRepository struct {
	*DBClient
	MemberRepository   MemberRepositoryInterface
	WalletRepository   WalletRepositoryInterface
	CurrencyRepository CurrencyRepositoryInterface

	DepositRepository     DepositRepositoryInterface
	WithdrawRepository    WithdrawRepositoryInterface
	TransferRepository    TransferRepositoryInterface
	TransactionRepository TransactionRepositoryInterface
}
