package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

type StartHandler struct {
	bot *tgbotapi.BotAPI
	DB  *gorm.DB
}

func NewStartHandler(bot *tgbotapi.BotAPI, db *gorm.DB) *StartHandler {
	return &StartHandler{bot: bot, DB: db}
}

func (h *StartHandler) Handle(ctx context.Context, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the bot! Type /help to see available commands.")
	h.bot.Send(msg)
}
