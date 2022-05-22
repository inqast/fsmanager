package sender

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Sender) SendMessage(chatId int64, response fmt.Stringer, disableNotification bool) {
	msg := tgbotapi.NewMessage(chatId, response.String())
	msg.DisableNotification = disableNotification
	s.bot.Send(msg)
}
