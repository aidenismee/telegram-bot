package router

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/pkg/telegram"
)

type service struct {
	telegramSvc telegram.Service
}

func NewService(cfg *configs.Configuration) *service {
	telegramSvc := telegram.New(cfg.TelegramBotToken, cfg.TelegramChatID, false)

	return &service{
		telegramSvc: telegramSvc,
	}
}
