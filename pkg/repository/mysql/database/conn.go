package database

import (
	"context"
	"crypto/config"
	"crypto/pkg/util/logger"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
)

type DBClient struct {
	clientMaster *gorm.DB
	clientSlave  *gorm.DB
}

func NewDBClient() (*DBClient, error) {
	masterCfg := config.GetMasterDBConfig()
	dbMaster = connectDB(masterCfg)
	logger.SysLog().Info(context.Background(), fmt.Sprintf("Master database [%s] connected", masterCfg.Address))

	dbSlave = dbMaster
	if slaveCfg := config.GetSlaveDBConfig(); slaveCfg != nil {
		dbSlave = connectDB(*slaveCfg)
		logger.SysLog().Info(context.Background(), fmt.Sprintf("Slave database [%s] connected", masterCfg.Address))
	}

	return &DBClient{clientMaster: dbMaster, clientSlave: dbSlave}, nil
}

// Session creates an original gorm.DB session.
func (d *DBClient) Session() *gorm.DB {
	return d.clientMaster
}

// SessionSlave creates an original gorm.DB session.
func (d *DBClient) SessionSlave() *gorm.DB {
	return d.clientSlave
}

func connectDB(cfg config.DBConfig) *gorm.DB {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Database,
	)
	var err error
	client, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		logger.SysLog().Error(context.Background(), fmt.Sprintf("Connection Fail ERR :%v", err))
	}
	if cfg.LogMode {
		client = client.Debug()
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	_client, err := client.DB()
	if err != nil {
		panic(err)
	}

	_client.SetMaxIdleConns(cfg.MaxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	_client.SetMaxOpenConns(cfg.MaxOpen)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	_client.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeMin) * time.Minute)

	return client
}
