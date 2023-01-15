package telegram

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func StartPolling() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(persistLastCommand),
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("5850142897:AAGvE8t02Ndrg0leux3nZKacNYjBZmFV_Og", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func persistLastCommand(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {

			log.Printf("%d say: %s", update.Message.From.ID, update.Message.Text)
		}
		next(ctx, b, update)
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message.Text == "Создать анкету" {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Hello",
			//Text:   update.Message.Text,
		})
	}
}
