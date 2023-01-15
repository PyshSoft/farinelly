package telegram

import (
	"context"
	"github.com/PyshSoft/farinelly/entity"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

type CommandManager struct {
	userState map[int64]UserState
	ctx       context.Context
	b         *bot.Bot
	update    *models.Update
	mu        sync.Mutex
}

type UserState struct {
	lastUserMessage string
	lastBotMessage  string
	profile         *entity.Profile
}

func (cm *CommandManager) createProfile() {
	userState := cm.userState[cm.update.Message.Chat.ID]
	switch userState.lastUserMessage {
	case "/start":
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите имя",
		})
		return
	default:
	}

	switch userState.lastBotMessage {
	case "Введите имя":
		userState.profile.UserName = cm.update.Message.Text
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите возраст",
		})
		return
	case "Введите возраст":
		var err error
		userState.profile.UserAge, err = strconv.Atoi(cm.update.Message.Text)
		if err != nil {
			panic(err)
		}
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите пол",
		})
		return
	case "Введите пол":
		sex := cm.update.Message.Text
		if sex == "Мужской" {
			userState.profile.Sex = entity.Male
		} else if sex == "Женский" {
			userState.profile.Sex = entity.Female
		}
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите город",
		})
		return
	case "Введите город":
		userState.profile.Location = cm.update.Message.Text
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите информацию о себе",
		})
		return
	case "Введите информацию о себе":
		userState.profile.About = cm.update.Message.Text
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите свои интересы",
		})
		return
	case "Введите свои интересы":
		userState.profile.UserName = cm.update.Message.Text
		cm.b.SendMessage(cm.ctx, &bot.SendMessageParams{
			ChatID: cm.update.Message.Chat.ID,
			Text:   "Введите возраст",
		})
		return
	}

}
