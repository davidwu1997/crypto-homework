//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package crypto

import (
	"crypto/config"
	"crypto/internal/delivery/http"
	"crypto/pkg/repository/mysql/database"
	"crypto/pkg/service/controller/restctl"
	"crypto/pkg/service/core/rest"
	"crypto/pkg/wireset"

	"github.com/google/wire"
)

func Initialize() (Application, error) {
	wire.Build(
		newApplication,
		config.New,
		wireset.InitLogger,
		rest.NewUseCase,
		restctl.NewRestCtl,
		http.NewHttpServer,
		database.NewDBClient,
		database.NewDBRepository,
	)
	return Application{}, nil
}
