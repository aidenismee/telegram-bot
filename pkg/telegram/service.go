package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service interface {
	Client() *tgbotapi.BotAPI
	SendMessage(message string) error
	SendHTMLMessage(message string) error
	SendMedia(files []interface{}) error
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

func (s *service) Client() *tgbotapi.BotAPI {
	return s.client
}
