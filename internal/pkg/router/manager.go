package router

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/user"
	"github.com/nekizz/telegram-bot/pkg/db"
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
	db := db.New(cfg.DbPsn, cfg.DbType)
	telegramSvc := telegram.New(cfg.TelegramBotToken, cfg.TelegramChatID, false)

	//internal service
	userSvc := user.NewService(db.DB())

	return &service{
		db:          db.DB(),
		telegramSvc: telegramSvc,
		userSvc:     userSvc,
	}
}
