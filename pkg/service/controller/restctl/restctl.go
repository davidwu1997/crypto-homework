package restctl

import (
	"crypto/pkg/service/core/rest"
	"crypto/pkg/util/logger"
)

func NewRestCtl(uc *rest.UseCase) *RestCtl {
	return &RestCtl{
		ResponseMiddleware: newResponseMiddleware(),
		MemberController:   newMemberController(uc, logger.SysLog()),
		CurrencyController: newCurrencyController(uc, logger.SysLog()),
	}

}

type RestCtl struct {
	ResponseMiddleware ResponseMiddlewareInterface
	MemberController   MemberControllerInterface
	CurrencyController CurrencyControllerInterface
}
