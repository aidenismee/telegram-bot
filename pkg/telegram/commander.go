package telegram

import (
	"fmt"
	"strings"
)

type cmdFunc func() error

type commander struct {
	telegramService Service
}

func newCommander(telegramSvc *service) *commander {
	return &commander{
		telegramService: telegramSvc,
	}
}

func (c *commander) helpCmd() error {
	var msgText strings.Builder

	msgText.WriteString("I can help you create and manage <b>telegram bot</b>.")
	msgText.WriteString("If you're new here, check out this list of commands you can use to interact with the bot.")
	msgText.WriteString("\n\n")
	msgText.WriteString("Here is the available command list: \n")
	msgText.WriteString("\n/hi - execute hi command")
	msgText.WriteString("\n/status - retrieve the current bot status")
	msgText.WriteString("\n/birthday - retrieve the member's birthday")
	msgText.WriteString("\n/help - return bot's list of command")

	text := msgText.String()

	return c.telegramService.SendHTMLMessage(text)
}

func (c *commander) hiCmd(userName string) error {
	msgText := fmt.Sprintf("Hello welcome, %s", userName)

	return c.telegramService.SendMessage(msgText)
}

func (c *commander) statusCmd() error {
	msgText := fmt.Sprintf("OK!")

	return c.telegramService.SendMessage(msgText)
}

func (c *commander) unknownCmd() error {
	msgText := fmt.Sprintf("I don't know that command")

	return c.telegramService.SendMessage(msgText)
}

func (c *commander) birthdayCmd() error {
	var msgText strings.Builder

	msgText.WriteString("*Member's birthday*\n")
	msgText.WriteString("\n<b>Minh A</b> - 05/11/2000")
	msgText.WriteString("\n<b>Hieu Le</b> - 15/08/2000")
	msgText.WriteString("\n<b>Minh Duc</b> - 10/11/2000")

	text := msgText.String()

	return c.telegramService.SendHTMLMessage(text)
}
