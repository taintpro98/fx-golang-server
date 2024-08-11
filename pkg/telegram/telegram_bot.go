package telegram

import (
	"context"
	"fmt"
	"fx-golang-server/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

func TelegramBotProvider(cnf *config.Config) ITelegramBot {
	bot, err := NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		log.Error().Err(err).Msg("cannot init telegram bot")
	}
	return bot
}

type ITelegramBot interface {
	SendMessage(ctx context.Context, text string) error
	GetMessages(ctx context.Context) error
}

type telegramBot struct {
	cnf config.TelegramBot
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(cnf config.TelegramBot) (ITelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cnf.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return &telegramBot{
		cnf: cnf,
		bot: bot,
	}, nil
}

func (b *telegramBot) SendMessage(ctx context.Context, text string) error {
	msg := tgbotapi.NewMessage(b.cnf.ChatID, text)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("send message error")
	}
	return err
}

func (b *telegramBot) GetMessages(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			fmt.Printf("Chat ID: %d, username: %s, message: %s", update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			b.bot.Send(msg)
		}
	}
	return nil
}