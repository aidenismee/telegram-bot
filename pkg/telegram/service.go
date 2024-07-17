package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service interface {
	SendMessage(message string) error
	NewMessage(message string) tgbotapi.MessageConfig
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

func (s *service) NewMessage(message string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(s.chatID, message)
}

func (s *service) SendMessage(message string) error {
	if _, err := s.client.Send(s.NewMessage(message)); err != nil {
		return err
	}

	return nil
}

func (s *service) CommandHandler() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.client.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "minha":
			msg.Text = "Anh Minh A rat tuyet voi la rat tuyet voi"
		case "hieule":
			msg.Text = "Anh Hieu Le rat la ngu, rat la ngu"
		case "minhduc":
			msg.Text = "Anh Minh Duc cung la rat la nguuu"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := s.client.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
