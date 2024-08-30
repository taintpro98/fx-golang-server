package handlers

import (
	"context"
	"fx-golang-server/module/blockchain"
	"fx-golang-server/pkg/utility"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EthHandler struct {
	bot       *tgbotapi.BotAPI
	DB        *gorm.DB
	ethClient blockchain.IEthClient
}

func NewEthHandler(bot *tgbotapi.BotAPI, db *gorm.DB, ethClient blockchain.IEthClient) *EthHandler {
	return &EthHandler{
		bot:       bot,
		DB:        db,
		ethClient: ethClient,
	}
}

func (v *EthHandler) Balance(ctx context.Context, message *tgbotapi.Message) {
	args := message.CommandArguments()

	// Kiểm tra nếu không có tham số được truyền vào
	if args == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Please provide a message to echo!")
		v.bot.Send(msg)
		return
	}

	address := args
	balance, err := v.ethClient.GetBalance(ctx, address)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("get balance error")
		msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		v.bot.Send(msg)
	}

	// Trả lại tham số đã nhận được
	msg := tgbotapi.NewMessage(message.Chat.ID, utility.GetETHValue(balance).String())
	v.bot.Send(msg)
}
