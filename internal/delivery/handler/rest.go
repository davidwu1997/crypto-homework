package handler

import (
	"crypto/pkg/service/controller/restctl"
)

func NewRestService(ctl *restctl.RestCtl) *RestService {
	return &RestService{
		ResponseMiddleware: ctl.ResponseMiddleware,
		MemberController:   ctl.MemberController,
		CurrencyController: ctl.CurrencyController,
	}
}

type RestService struct {
	ResponseMiddleware restctl.ResponseMiddlewareInterface
	MemberController   restctl.MemberControllerInterface
	CurrencyController restctl.CurrencyControllerInterface
}
