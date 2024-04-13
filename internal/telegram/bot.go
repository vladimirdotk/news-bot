package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	updatesChan tgbotapi.UpdatesChannel

	messageHandler MessageHandler

	log *slog.Logger
}

func NewBot(token string, messageHandler MessageHandler, debug bool, log *slog.Logger) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("create bot: %v", err)
	}

	botAPI.Debug = debug

	bot := Bot{
		bot:            botAPI,
		messageHandler: messageHandler,
		log: slog.With(
			slog.Group("telegram bot"),
		),
	}

	bot.log.Debug(
		"Authorized",
		slog.String("account", bot.bot.Self.UserName),
	)

	bot.updatesChan = botAPI.GetUpdatesChan(
		tgbotapi.UpdateConfig{
			Timeout: 60,
		},
	)

	return &bot, nil
}

func (b *Bot) Run() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	for update := range b.updatesChan {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		b.log.Debug(
			"Received message",
			slog.String("username", update.Message.From.UserName),
			slog.String("text", update.Message.Text),
		)

		incomingMessage := incomingMessageToDomain(update.Message)
		if incomingMessage == nil {
			continue
		}

		if err := b.messageHandler.Handle(ctx, incomingMessage); err != nil {
			b.log.Error(
				"Handle message",
				slog.String("error", err.Error()),
			)
			b.reply(update.Message.Chat.ID, update.Message.MessageID, "Error happend")
			continue
		}
	}
}

func (b *Bot) reply(chatID int64, replyToMessageID int, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyToMessageID = replyToMessageID
	b.send(message)
}

func (b *Bot) send(message tgbotapi.Chattable) {
	if _, err := b.bot.Send(message); err != nil {
		b.log.Error(
			"Send message",
			slog.String("error", err.Error()),
		)
	}
}

func incomingMessageToDomain(src *tgbotapi.Message) *domain.IncomingMessage {
	if src == nil {
		return nil
	}

	return &domain.IncomingMessage{
		ID:       strconv.Itoa(src.MessageID),
		UserID:   strconv.FormatInt(src.From.ID, 10),
		Username: src.From.UserName,
		Text:     src.Text,
		Source:   domain.MessageSystemTelegram,
	}
}
