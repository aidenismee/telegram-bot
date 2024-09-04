package user

import (
	"github.com/nekizz/telegram-bot/internal/repository"
	"github.com/nekizz/telegram-bot/pkg/telegram"
	"gorm.io/gorm"
)

type service struct {
	db       *gorm.DB
	telegram telegram.Service
	userRepo repository.UserRepository
}

func NewService(db *gorm.DB) *service {
	return &service{
		db:       db,
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *service) Hello() error {
	return nil
}
