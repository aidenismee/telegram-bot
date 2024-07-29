package migration

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/migration/versions"
	"github.com/nekizz/telegram-bot/pkg/db"
	"github.com/nekizz/telegram-bot/pkg/migration"
)

func Run(cfg *configs.Configuration) (respErr error) {
	db := db.New(cfg.DbPsn, cfg.DbType).DB()

	defer func() {
		sqlDb, err := db.DB()
		if sqlDb != nil && err == nil {
			sqlDb.Close()
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				respErr = fmt.Errorf("%s", x)
			case error:
				respErr = x
			default:
				respErr = fmt.Errorf("unknown error: %+v", x)
			}
		}
	}()

	migration.Run(db, []*gormigrate.Migration{
		{
			ID:       "20230808080810",
			Migrate:  versions.Version20230808080810,
			Rollback: versions.Rollback20230808080810,
		},
	})

	return nil
}
