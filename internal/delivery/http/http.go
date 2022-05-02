package http

import (
	"crypto/internal/delivery/handler"
	"crypto/pkg/service/controller/restctl"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*gin.Engine
}

func (server *HttpServer) setRouter(handler *handler.RestService) {
	publicRouteGroup := server.Group("")

	// 設定 middleware
	publicRouteGroup.Use(
		handler.ResponseMiddleware.Handle,
	)

	currencyRouteGroup := publicRouteGroup.Group("/currency")
	currencyRouteGroup.POST("/create", handler.CurrencyController.Create)

	memberRouteGroup := publicRouteGroup.Group("/member-api")

	memberRouteGroup.POST("/create", handler.MemberController.Create)
	memberRouteGroup.GET("/wallet/list", handler.MemberController.WalletList)
	memberRouteGroup.POST("/wallet/deposit", handler.MemberController.Deposit)
	memberRouteGroup.POST("/wallet/withdraw", handler.MemberController.Withdraw)
	memberRouteGroup.POST("/wallet/transfer", handler.MemberController.Transfer)

	memberRouteGroup.GET("/report/deposit-log", handler.MemberController.DepositLog)
	memberRouteGroup.GET("/report/withdraw-log", handler.MemberController.WithdrawLog)
	memberRouteGroup.GET("/report/transfer-log", handler.MemberController.TransferLog)
	memberRouteGroup.GET("/report/transaction-log", handler.MemberController.TransactionLog)
}

func NewHttpServer(controller *restctl.RestCtl) *HttpServer {
	httpServer := &HttpServer{
		Engine: gin.Default(),
	}

	restHandler := handler.NewRestService(controller)

	//httpServer.Use(middleware)

	httpServer.setRouter(restHandler)

	return httpServer
}
