package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandState struct {
	ChatId      int64
	LastCommand string
}

func createProfile(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите имя",
		//Text:   update.Message.Text,
	})

}
