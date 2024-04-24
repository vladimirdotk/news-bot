package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	incomingChan tgbotapi.UpdatesChannel
	outgoingChan <-chan domain.OutgoingMessage

	messageHandler MessageHandler

	log *slog.Logger
}

func NewBot(
	token string,
	messageHandler MessageHandler,
	outgoingChan <-chan domain.OutgoingMessage,
	debug bool,
	log *slog.Logger,
) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("create bot: %v", err)
	}

	botAPI.Debug = debug

	bot := Bot{
		bot:            botAPI,
		messageHandler: messageHandler,
		outgoingChan:   outgoingChan,
		log: slog.With(
			slog.Group("telegram bot"),
		),
	}

	bot.log.Debug(
		"Authorized",
		slog.String("account", bot.bot.Self.UserName),
	)

	bot.incomingChan = botAPI.GetUpdatesChan(
		tgbotapi.UpdateConfig{
			Timeout: 60,
		},
	)

	return &bot, nil
}

func (b *Bot) Run(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go b.handleIncoming(ctx, &wg)
	go b.handleOutgoing(ctx, &wg)

	wg.Wait()
}

func (b *Bot) handleIncoming(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			break
		case update := <-b.incomingChan:
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
}

func (b *Bot) handleOutgoing(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			break
		case messsage := <-b.outgoingChan:
			userID, err := strconv.ParseInt(messsage.UserID, 10, 64)
			if err != nil {
				b.log.Error(
					"Converting userID to int",
					slog.String("error", err.Error()),
				)
			}

			b.send(
				tgbotapi.NewMessage(userID, messsage.Text),
			)
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
