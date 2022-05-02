package config

import (
	"crypto/pkg/util/logger"
	"strings"

	"github.com/spf13/viper"
)

type EnvType = string

const (
	EnvTypeLocal EnvType = "local"
	EnvTypeDev   EnvType = "dev"
	EnvTypeProd  EnvType = "prod"

	defaultEnv EnvType = EnvTypeLocal
)

// Config ...
var (
	config *ConfigSetup
)

func New() (*ConfigSetup, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		configPath = viper.GetString("PROJ_DIR") + "/deployment/config"
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "default"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	if err := logger.InitSysLog(
		GetLogConfig().Name,
		GetLogConfig().Level); err != nil {

		panic(err)
	}
	return config, nil
}

type ConfigSetup struct {
	LogConfig      LogConfig `mapstructure:"logger"`
	GinConfig      GinConfig `mapstructure:"gin_config"`
	DBMasterConfig DBConfig  `mapstructure:"db_master_config"`
	DBSlaveConfig  *DBConfig `mapstructure:"db_slave_config"`
}
