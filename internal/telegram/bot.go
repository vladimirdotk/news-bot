package telegram

import (
	"fmt"
	"log"

	"github.com/vladimirdotk/news-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	updatesChan tgbotapi.UpdatesChannel

	messageHandler MessageHandler
}

func NewBot(token string, messageHandler MessageHandler, debug bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("create bot: %v", err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, fmt.Errorf("create updates channel: %v", err)
	}

	return &Bot{
		bot:            bot,
		updatesChan:    updates,
		messageHandler: messageHandler,
	}, nil
}

func (b *Bot) Run() {
	for update := range b.updatesChan {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		incomingMessage := incomingMessageToDomain(update.Message)
		if err := b.messageHandler.Handle(incomingMessage); err != nil {
			// log error
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка")
			msg.ReplyToMessageID = update.Message.MessageID
			b.bot.Send(msg)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принято")
		msg.ReplyToMessageID = update.Message.MessageID

		b.bot.Send(msg)
	}
}

func incomingMessageToDomain(src *tgbotapi.Message) *domain.IncomingMessage {
	if src == nil {
		return nil
	}

	return &domain.IncomingMessage{
		ID:       string(src.MessageID),
		UserID:   string(src.From.ID),
		Username: src.From.UserName,
		Text:     src.Text,
	}
}
