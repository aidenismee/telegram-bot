package telegram

import (
	"fmt"
	"github.com/nekizz/telegram-bot/internal/model"
	"strings"
)

func toCheckBirthdayMessage(user *model.User) string {
	var msgText strings.Builder

	msgText.WriteString(
		fmt.Sprintf("Count not the candles… see the lights they give. Count not the years, but the life you live. Wishing you a wonderful time ahead. Happy birthday %s", user.FirstName+user.LastName))

	return msgText.String()
}
