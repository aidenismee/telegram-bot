package router

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/pkg/telegram"
)

type service struct {
	telegramSvc telegram.Service
}

func NewService(cfg *configs.Configuration) *service {
	telegramSvc := telegram.New("6873033629:AAG1XJfuvM0704IIrwro3G9Q4bEgaG2_UgE", -1002248978143, false)

	return &service{
		telegramSvc: telegramSvc,
	}
}
