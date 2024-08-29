package manager

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/telegram"
	"github.com/nekizz/telegram-bot/internal/api/user"
	"github.com/nekizz/telegram-bot/pkg/db"
	telegramPkg "github.com/nekizz/telegram-bot/pkg/telegram"
)

type Manager interface {
	UserHandler() *user.Handler
	TelegramHandler() *telegram.Handler
}

type manager struct {
	telegramHandler *telegram.Handler
	userHandler     *user.Handler
}

func NewManager(cfg *configs.Configuration) Manager {
	db := db.New(cfg.DbPsn, cfg.DbType)
	telegramSvc := telegramPkg.New(cfg.TelegramBotToken, cfg.TelegramChatID, false)

	userSvc := user.NewService(db.DB())
	teleSvc := telegram.NewService(db.DB(), telegramSvc)

	return &manager{
		userHandler:     user.NewHandler(userSvc),
		telegramHandler: telegram.NewHandler(teleSvc),
	}
}

func (m *manager) UserHandler() *user.Handler {
	return m.userHandler
}

func (m *manager) TelegramHandler() *telegram.Handler {
	return m.telegramHandler
}
