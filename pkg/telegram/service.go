package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
)

type Service interface {
	SendMessage(message string) error
	SendHTMLMessage(message string) error
	SendMedia(files []interface{}) error
	CommandHandler() error
}

type service struct {
	chatID int64
	client *tgbotapi.BotAPI
}

func New(telegramBotToken string, chatID int64, debugMode bool) *service {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return nil
	}
	bot.Debug = debugMode

	return &service{
		chatID: chatID,
		client: bot,
	}
}

func (s *service) SendMessage(message string) error {
	messageCfg := tgbotapi.NewMessage(s.chatID, message)

	if _, err := s.client.Send(messageCfg); err != nil {
		return err
	}

	return nil
}

func (s *service) SendHTMLMessage(message string) error {
	messageCfg := tgbotapi.NewMessage(s.chatID, message)
	messageCfg.ParseMode = tgbotapi.ModeHTML

	if _, err := s.client.Send(messageCfg); err != nil {
		return err
	}

	return nil
}

func (s *service) SendMedia(files []interface{}) error {
	mediaGroupCfg := tgbotapi.NewMediaGroup(s.chatID, files)

	if _, err := s.client.SendMediaGroup(mediaGroupCfg); err != nil {
		return err
	}

	return nil
}

func (s *service) CommandHandler() error {
	var (
		err error
	)

	cmd := newCommander(s)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.client.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "help":
			err = cmd.helpCmd()
		case "hi":
			err = cmd.hiCmd(update.Message.From.UserName)
		case "status":
			err = cmd.statusCmd()
		case "birthday":
			err = cmd.birthdayCmd()
		default:
			err = cmd.unknownCmd()
		}
		if err != nil {
			log.Print(err)
		}
	}

	return nil
}
