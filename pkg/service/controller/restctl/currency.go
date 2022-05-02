package restctl

import (
	"crypto/pkg/model/bo"
	"crypto/pkg/model/dto"
	"crypto/pkg/service/core/rest"
	"crypto/pkg/util/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyControllerInterface interface {
	Create(ctx *gin.Context)
}

func newCurrencyController(in *rest.UseCase, logger logger.LoggerInterface) CurrencyControllerInterface {
	return &currencyController{
		logger: logger,
		in:     in,
	}
}

type currencyController struct {
	logger logger.LoggerInterface

	in *rest.UseCase
}

func (ctl currencyController) Create(ctx *gin.Context) {
	req := &dto.CreateCurrencyReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusBadRequest, "0", err.Error())
		return
	}

	if err := ctl.in.CurrencyUseCase.Create(ctx, &bo.Currency{Code: req.Code, Name: req.Name}); err != nil {
		ctl.logger.Error(ctx, err.Error())
		SetResp(ctx, http.StatusInternalServerError, "0", err.Error())
		return
	}

	SetResp(ctx, http.StatusOK, "0", "success")
}
