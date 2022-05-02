//+build wireinject

//The build tag makes sure the stub is not built in the final build.

package migrate

import (
	"crypto/config"
	"crypto/pkg/repository/mysql/database"
	"crypto/pkg/wireset"

	"github.com/google/wire"
)

func Initialize(configPath string) (Application, error) {
	wire.Build(
		newApplication,
		database.NewDBClient,
		wireset.InitLogger,
		config.New,
	)
	return Application{}, nil
}
