package telegram

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimirdotk/news-bot/internal/domain"
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
		if incomingMessage == nil {
			continue
		}

		if err := b.messageHandler.Handle(incomingMessage); err != nil {
			log.Printf("handle message: %v", err)
			b.answer(update.Message.Chat.ID, update.Message.MessageID, "Произошла ошибка")
			continue
		}
	}
}

func (b *Bot) answer(chatID int64, replyToMessageID int, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyToMessageID = replyToMessageID
	b.send(message)
}

func (b *Bot) send(message tgbotapi.Chattable) {
	if _, err := b.bot.Send(message); err != nil {
		log.Printf("send message: %v", err)
	}
}

func incomingMessageToDomain(src *tgbotapi.Message) *domain.IncomingMessage {
	if src == nil {
		return nil
	}

	return &domain.IncomingMessage{
		ID:       strconv.Itoa(src.MessageID),
		UserID:   strconv.Itoa(src.From.ID),
		Username: src.From.UserName,
		Text:     src.Text,
		Source:   domain.MessageSystemTelegram,
	}
}
