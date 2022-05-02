package wireset

import (
	"crypto/config"
	"crypto/pkg/logger"

	"github.com/rs/zerolog"
)

func InitLogger(config *config.ConfigSetup) (zerolog.Logger, error) {

	return logger.NewLogger(config.LogConfig.Level, config.LogConfig.Format)
}
