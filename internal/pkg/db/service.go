package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekizz/telegram-bot/pkg/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New creates new database connection to a postgres database
func New(dbPsn string, enableLog bool) *gorm.DB {
	db, _ := db.New(mysql.Open(dbPsn), enableLog)

	return db
}
