package crypto

import (
	"crypto/config"
	"crypto/internal/delivery/http"
	"fmt"

	"github.com/rs/zerolog"
)

type Application struct {
	logger     zerolog.Logger
	config     *config.ConfigSetup
	httpServer *http.HttpServer
}

func (application Application) Start() error {
	application.logger.Info().Msgf("http server listen :%d", application.config.GinConfig.Port)
	return application.httpServer.Run(fmt.Sprintf(":%d", application.config.GinConfig.Port))
}

func newApplication(
	logger zerolog.Logger,
	config *config.ConfigSetup,
	httpServer *http.HttpServer,
) Application {
	return Application{
		logger:     logger,
		config:     config,
		httpServer: httpServer,
	}
}
