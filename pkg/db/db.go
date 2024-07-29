package db

import (
	"time"

	"gorm.io/gorm"

	"github.com/nekizz/telegram-bot/pkg/db/dialects"
)

type DB interface {
	Ping() error
	DB() *gorm.DB
}

type dbConnection struct {
	client *gorm.DB
}

// New creates new database connection to the database server
func New(dbPsn, dialectorType string) *dbConnection {
	orm, err := gorm.Open(dialects.NewDialector(dbPsn, dialectorType), &gorm.Config{
		AllowGlobalUpdate:    false,
		FullSaveAssociations: true,
	})
	if nil != err {
		return nil
	}

	sqlDB, err := orm.DB()
	if nil != err {
		return nil
	}

	// TODO: use config for these values
	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	return &dbConnection{client: orm}
}

func (db *dbConnection) DB() *gorm.DB {
	return db.client
}

func (db *dbConnection) Ping() error {
	sqlDB, err := db.client.DB()
	if nil != err {
		return err
	}

	return sqlDB.Ping()
}
