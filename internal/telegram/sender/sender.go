package sender

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender struct {
	bot *tgbotapi.BotAPI
}

func New(
	bot *tgbotapi.BotAPI,
) *Sender {
	return &Sender{
		bot: bot,
	}
}
