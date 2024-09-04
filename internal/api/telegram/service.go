package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	errorUtils "github.com/nekizz/telegram-bot/internal/errors"
	"github.com/nekizz/telegram-bot/internal/repository"
	logger "github.com/nekizz/telegram-bot/pkg/log"
	"github.com/nekizz/telegram-bot/pkg/telegram"
	"gorm.io/gorm"
	"log"
	"time"
)

type service struct {
	db       *gorm.DB
	logger   *logger.Logger
	telegram telegram.Service
	userRepo repository.UserRepository
}

func NewService(db *gorm.DB, telegramSvc telegram.Service) *service {
	return &service{
		db:       db,
		telegram: telegramSvc,
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *service) handleCommand() error {
	var err error

	defer func() {
		if err := recover(); err != nil {
			log.Print("recovered from panic")
		}
	}()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.telegram.Client().GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "help":
			err = s.helpCmd()
		case "hi":
			err = s.hiCmd(update.Message.From.UserName)
		case "status":
			err = s.statusCmd()
		case "birthday":
			err = s.birthdayCmd()
		case "abuse":
			err = s.abuseCmd(update.Message.From.UserName)
		case "images":
			err = s.imagesCmd()
		default:
			err = s.unknownCmd()
		}
		if err != nil {
			log.Print(err)
		}
	}

	return nil
}

func (s *service) checkBirthdays(c echo.Context) error {
	logFields := map[string]any{
		"service": "checkBirthdays",
	}
	user, err := s.userRepo.ReadByCondition(
		fmt.Sprintf("date_of_birth = %s",
			time.Now().Format("02-01-2006")))
	if err != nil {
		s.logger.FromContext(c).WithFields(logFields).WithErr(err).Error("get user info failed")
		return errorUtils.ErrInvalidRequest.SetInternal(err)
	}

	if err := s.telegram.SendHTMLMessage(toCheckBirthdayMessage(user)); err != nil {
		s.logger.FromContext(c).WithFields(logFields).WithErr(err).Error("send telegram message failed")
		return errorUtils.ErrInvalidRequest.SetInternal(err)
	}

	return nil
}

func (s *service) alertJob(c echo.Context) error {
	logFields := map[string]any{
		"service": "alertJob",
	}

	t := time.Now()

	if t.Hour() == 8 && (t.Minute() >= 30 || t.Minute() <= 59) {
		if err := s.telegram.SendMessage("Di lam thoi cac eimm oiii!"); err != nil {
			s.logger.FromContext(c).WithFields(logFields).WithErr(err).Error("send telegram message failed")
			return errorUtils.ErrInvalidRequest.SetInternal(err)
		}
	}

	if t.Hour() == 18 && (t.Minute() >= 30 || t.Minute() <= 59) {
		if err := s.telegram.SendMessage("Di lam thoi cac eimm oiii!"); err != nil {
			s.logger.FromContext(c).WithFields(logFields).WithErr(err).Error("send telegram message failed")
			return errorUtils.ErrInvalidRequest.SetInternal(err)
		}
	}

	return nil
}
