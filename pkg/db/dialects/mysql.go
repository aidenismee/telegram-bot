package dialects

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dialectMysql(dbPsn string) gorm.Dialector {
	return mysql.Open(dbPsn)
}
