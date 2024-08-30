package telebot

import (
	"context"
	"fx-golang-server/module/blockchain"
	"fx-golang-server/module/telebot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type TelegramClient struct {
	bot          *tgbotapi.BotAPI
	startHandler *handlers.StartHandler
	ethHandler   *handlers.EthHandler
}

func NewTelegramClient(
	bot *tgbotapi.BotAPI,
	db *gorm.DB,
	ethClient blockchain.IEthClient,
) *TelegramClient {
	startHandler := handlers.NewStartHandler(bot, db)
	ethHandler := handlers.NewEthHandler(bot, db, ethClient)
	return &TelegramClient{
		bot:          bot,
		startHandler: startHandler,
		ethHandler:   ethHandler,
	}
}

func (b *TelegramClient) Handle(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Info().Ctx(ctx).
				Interface("text", update.Message.Text).
				Interface("from", update.Message.From).
				Msg("Incoming message")
				
			switch update.Message.Command() {
			case "start":
				b.startHandler.Handle(ctx, update.Message)
			case "balance":
				b.ethHandler.Balance(ctx, update.Message)
			}
		}
	}
	return nil
}
