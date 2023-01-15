package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandManager struct {
	userState map[int64]string
	ctx       context.Context
	b         *bot.Bot
	update    *models.Update
}

func (cm *CommandManager) createProfile() {
	lastUserCommand := cm.userState[cm.update.Message.Chat.ID]
	switch lastUserCommand {
	case "/start":
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите имя",
		})
	}
	if lastUserCommand == "/start" {

	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите имя",
		//Text:   update.Message.Text,
	})

}
