package user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	errorUtils "github.com/nekizz/telegram-bot/internal/errors"
	"github.com/nekizz/telegram-bot/internal/repository"
	"github.com/nekizz/telegram-bot/pkg/telegram"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	db       *gorm.DB
	telegram telegram.Service
	userRepo repository.UserRepository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:       db,
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *Service) checkBirthdays(c echo.Context) error {
	user, err := s.userRepo.ReadByCondition(
		fmt.Sprintf("date_of_birth = %s",
			time.Now().Format("02-01-2006")))
	if err != nil {
		return errorUtils.ErrInvalidRequest.SetInternal(err)
	}

	if err := s.telegram.SendHTMLMessage(toCheckBirthdayMessage(user)); err != nil {
		return errorUtils.ErrInvalidRequest.SetInternal(err)
	}

	return nil
}

func (s *Service) alertJob(c echo.Context) error {
	t := time.Now()

	if t.Hour() == 8 && (t.Minute() >= 30 || t.Minute() <= 59) {
		if err := s.telegram.SendMessage("Di lam thoi cac eimm oiii!"); err != nil {
			return errorUtils.ErrInvalidRequest.SetInternal(err)
		}
	}

	if t.Hour() == 18 && (t.Minute() >= 30 || t.Minute() <= 59) {
		if err := s.telegram.SendMessage("Di lam thoi cac eimm oiii!"); err != nil {
			return errorUtils.ErrInvalidRequest.SetInternal(err)
		}
	}

	return nil
}
