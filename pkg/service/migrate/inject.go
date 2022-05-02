package migrate

import (
	"crypto/deployment/migration"
	"crypto/pkg/repository/mysql/database"

	"crypto/config"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog"
)

type Application struct {
	config *config.ConfigSetup
	logger zerolog.Logger
	mysql  *database.DBClient
}

func (application Application) Start() error {

	m := gormigrate.New(application.mysql.Session(), gormigrate.DefaultOptions, migration.Migrations)
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}

func newApplication(config *config.ConfigSetup, db *database.DBClient, logger zerolog.Logger) Application {
	return Application{
		config: config,
		logger: logger,
		mysql:  db,
	}
}
