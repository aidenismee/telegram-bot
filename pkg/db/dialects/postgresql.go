package dialects

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dialectPostgreSql(dbPsn string) gorm.Dialector {
	return postgres.Open(dbPsn)
}
