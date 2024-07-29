package dialects

import "gorm.io/gorm"

func NewDialector(dbPsn, dialectorType string) gorm.Dialector {
	switch dialectorType {
	case "mysql":
		return dialectMysql(dbPsn)
	case "postgres":
		return dialectPostgreSql(dbPsn)
	default:
		return dialectMysql(dbPsn)
	}
}
