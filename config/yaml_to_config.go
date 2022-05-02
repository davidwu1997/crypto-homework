package config

import "crypto/pkg/logger"

type LogConfig struct {
	Name   string           `mapstructure:"name"`
	Env    string           `mapstructure:"env"`
	Level  string           `mapstructure:"level"`
	Format logger.LogFormat `mapstructure:"format"`
}

type GinConfig struct {
	Address string `mapstructure:"address"`
	Port    uint16 `mapstructure:"port"`
}

type DBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`

	LogMode        bool `mapstructure:"log_mode"`
	MaxIdle        int  `mapstructure:"max_idle"`
	MaxOpen        int  `mapstructure:"max_open"`
	ConnMaxLifeMin int  `mapstructure:"conn_max_life_min"`
}

func GetLogConfig() LogConfig {
	return config.LogConfig
}

func GetGinConfig() GinConfig {
	return config.GinConfig
}

func GetMasterDBConfig() DBConfig {
	return config.DBMasterConfig
}

func GetSlaveDBConfig() *DBConfig {
	return config.DBSlaveConfig
}
