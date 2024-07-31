package manager

import (
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/telegram"
	"github.com/nekizz/telegram-bot/internal/api/user"
	"github.com/nekizz/telegram-bot/pkg/db"
	telegramPkg "github.com/nekizz/telegram-bot/pkg/telegram"
	"gorm.io/gorm"
)

type Manager interface {
	UserService() *user.Service
	TeleService() *telegram.Service
}

type manager struct {
	db          *gorm.DB
	telegramSvc telegramPkg.Service
	userSvc     *user.Service
	teleSvc     *telegram.Service
}

func NewManager(cfg *configs.Configuration) Manager {
	db := db.New(cfg.DbPsn, cfg.DbType)
	telegramSvc := telegramPkg.New(cfg.TelegramBotToken, cfg.TelegramChatID, false)

	userSvc := user.NewService(db.DB())
	teleSvc := telegram.NewService(db.DB(), telegramSvc)

	return &manager{
		db:          db.DB(),
		userSvc:     userSvc,
		telegramSvc: telegramSvc,
		teleSvc:     teleSvc,
	}
}

func (m *manager) UserService() *user.Service {
	return m.userSvc
}

func (m *manager) TeleService() *telegram.Service {
	return m.teleSvc
}
