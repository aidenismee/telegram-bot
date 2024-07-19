package router

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/user"
	dbPkg "github.com/nekizz/telegram-bot/internal/pkg/db"
	"github.com/nekizz/telegram-bot/pkg/telegram"
	"gorm.io/gorm"
)

type service struct {
	db          *gorm.DB
	telegramSvc telegram.Service
	userSvc     *user.Service
}

func NewService(cfg *configs.Configuration) *service {
	//external service
	db := dbPkg.New(cfg.DbPsn, cfg.DbLog)
	telegramSvc := telegram.New(cfg.TelegramBotToken, cfg.TelegramChatID, false)

	//internal service
	userSvc := user.NewService(db)

	return &service{
		db:          db,
		telegramSvc: telegramSvc,
		userSvc:     userSvc,
	}
}
