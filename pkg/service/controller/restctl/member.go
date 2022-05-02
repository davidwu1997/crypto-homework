package restctl

import (
	"crypto/pkg/model/bo"
	"crypto/pkg/model/dto"
	"crypto/pkg/service/core/rest"
	"crypto/pkg/util/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberControllerInterface interface {
	Create(ctx *gin.Context)
	WalletList(ctx *gin.Context)

	Deposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
	Transfer(ctx *gin.Context)

	DepositLog(ctx *gin.Context)
	WithdrawLog(ctx *gin.Context)
	TransferLog(ctx *gin.Context)
	TransactionLog(ctx *gin.Context)
}

func newMemberController(in *rest.UseCase, logger logger.LoggerInterface) MemberControllerInterface {
	return &memberController{
		logger: logger,
		in:     in,
	}
}

type memberController struct {
	logger logger.LoggerInterface

	in *rest.UseCase
}

func (ctl memberController) Create(ctx *gin.Context) {
	req := &dto.CreateMemberReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	if err := ctl.in.MemberUseCase.Create(ctx, &bo.Member{Login: req.Login, PassWord: req.PassWord, Name: req.Name}); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	SetResp(ctx, http.StatusOK, "0", "success")
}

func (ctl memberController) WalletList(ctx *gin.Context) {
	req := &dto.MemberWalletReq{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	wallets, err := ctl.in.MemberUseCase.GetWalletList(ctx, &bo.MemberWalletCond{
		MemberLogin: req.MemberLogin,
	})

	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	result := make([]*dto.MemberWalletResp, len(wallets))
	for _, wallet := range wallets {
		result = append(result, &dto.MemberWalletResp{
			WalletID:     strconv.FormatUint(wallet.ID, 10),
			CurrencyCode: wallet.CurrencyCode,
			Balance:      wallet.Balance,
		})
	}

	fmt.Println(result)

	SetResp(ctx, http.StatusOK, "0", result)
}

func (ctl memberController) Deposit(ctx *gin.Context) {
	req := &dto.MemberDepositReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	walletID, err := strconv.ParseUint(req.WalletID, 10, 64)
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	if err = ctl.in.MemberUseCase.Deposit(ctx, &bo.MemberDeposit{WalletID: walletID, TransferAmount: req.TransferAmount}); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	SetResp(ctx, http.StatusOK, "0", "success")
}

func (ctl memberController) Withdraw(ctx *gin.Context) {
	req := &dto.MemberWithdrawReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	walletID, err := strconv.ParseUint(req.WalletID, 10, 64)
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	if err = ctl.in.MemberUseCase.Withdraw(ctx, &bo.MemberWithdraw{WalletID: walletID, TransferAmount: req.TransferAmount}); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	SetResp(ctx, http.StatusOK, "0", "success")
}

func (ctl memberController) Transfer(ctx *gin.Context) {
	req := &dto.MemberTransferReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	fromWalletID, err := strconv.ParseUint(req.FromWalletID, 10, 64)
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}
	toWalletID, err := strconv.ParseUint(req.ToWalletID, 10, 64)
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	if err = ctl.in.MemberUseCase.Transfer(ctx, &bo.MemberTransfer{FromWalletID: fromWalletID, TransferAmount: req.TransferAmount, ToWalletID: toWalletID}); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	SetResp(ctx, http.StatusOK, "0", "success")
}

func (ctl memberController) DepositLog(ctx *gin.Context) {
	req := &dto.MemberDepositLogReq{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	logs, err := ctl.in.MemberUseCase.DepositLog(ctx, &bo.DepositCond{MemberLogin: &req.MemberLogin, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	result := make([]*dto.MemberDepositLogResp, len(logs))
	for _, log := range logs {
		result = append(result, &dto.MemberDepositLogResp{
			ID:             strconv.FormatUint(log.ID, 10),
			WalletID:       strconv.FormatUint(log.WalletID, 10),
			TransferAmount: log.Amount,
			CreatedAt:      log.CreatedAt,
		})
	}

	SetResp(ctx, http.StatusOK, "0", result)
}

func (ctl memberController) WithdrawLog(ctx *gin.Context) {
	req := &dto.MemberWithdrawLogReq{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	logs, err := ctl.in.MemberUseCase.WithdrawLog(ctx, &bo.WithdrawCond{MemberLogin: &req.MemberLogin, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	result := make([]*dto.MemberWithdrawLogResp, len(logs))
	for _, log := range logs {
		result = append(result, &dto.MemberWithdrawLogResp{
			ID:             strconv.FormatUint(log.ID, 10),
			WalletID:       strconv.FormatUint(log.WalletID, 10),
			TransferAmount: log.Amount,
			CreatedAt:      log.CreatedAt,
		})
	}

	SetResp(ctx, http.StatusOK, "0", result)
}

func (ctl memberController) TransferLog(ctx *gin.Context) {
	req := &dto.MemberTransferLogReq{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	logs, err := ctl.in.MemberUseCase.TransferLog(ctx, &bo.TransferCond{MemberLogin: &req.MemberLogin, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	result := make([]*dto.MemberTransferLogResp, len(logs))
	for _, log := range logs {
		result = append(result, &dto.MemberTransferLogResp{
			ID:             strconv.FormatUint(log.ID, 10),
			FromWalletID:   strconv.FormatUint(log.FromWalletID, 10),
			TransferAmount: log.Amount,
			ToWalletID:     strconv.FormatUint(log.ToWalletID, 10),
			CreatedAt:      log.CreatedAt,
		})
	}

	SetResp(ctx, http.StatusOK, "0", result)
}

func (ctl memberController) TransactionLog(ctx *gin.Context) {
	req := &dto.MemberTransactionLogReq{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	logs, err := ctl.in.MemberUseCase.TransactionLog(ctx, &bo.TransactionCond{MemberLogin: &req.MemberLogin, StartTime: req.StartTime, EndTime: req.EndTime})
	if err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	result := make([]*dto.MemberTransactionLogResp, len(logs))
	for _, log := range logs {
		result = append(result, &dto.MemberTransactionLogResp{
			ID:             strconv.FormatUint(log.ID, 10),
			WalletID:       strconv.FormatUint(log.WalletID, 10),
			Method:         log.Method,
			ActionID:       strconv.FormatUint(log.ActionID, 10),
			TransferBefore: log.TransferBefore,
			TransferAmount: log.TransferAmount,
			TransferAfter:  log.TransferAfter,
			CreatedAt:      log.CreatedAt,
		})
	}

	SetResp(ctx, http.StatusOK, "0", result)
}
