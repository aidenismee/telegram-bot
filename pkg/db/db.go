package db

import (
	"time"

	"gorm.io/gorm"
)

// New creates new database connection to the database server
func New(dialect gorm.Dialector, enableLog bool) (*gorm.DB, error) {
	orm, err := gorm.Open(dialect, &gorm.Config{
		AllowGlobalUpdate:    true,
		FullSaveAssociations: true,
	})
	if nil != err {
		return nil, err
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	// TODO: use config for these values
	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	return orm, nil
}
