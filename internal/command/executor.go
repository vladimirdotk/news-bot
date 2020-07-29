package command

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Executor struct {
	redisClient    *redis.Client
	responseSender ResponseSender
}

func NewExecutor(redisClient *redis.Client, responseSender ResponseSender) *Executor {
	return &Executor{
		redisClient:    redisClient,
		responseSender: responseSender,
	}
}

func (e *Executor) Exec(message domain.IncomingMessage) error {
	if strings.HasPrefix(message.Text, domain.MessageCommandAddSource) {
		return e.addSource(message)
	}

	if strings.HasPrefix(message.Text, domain.MessageCommandListSource) {
		return e.listSources(message)
	}

	return fmt.Errorf("executor not found for command: %s", message)
}

func (e *Executor) addSource(message domain.IncomingMessage) error {
	commandArgs, err := getCommandArgs(message.Text)
	if err != nil {
		return fmt.Errorf("get command args: %v", err)
	}

	// TODO: validate source
	if err := e.redisClient.SAdd(message.UserID, commandArgs[1]).Err(); err != nil {
		return fmt.Errorf("sadd, key: %s, value: %s, err: %v", message.UserID, commandArgs[1], err)
	}

	outgoingMessage := toOutgoingMessage(message, "Источник добавлен")

	if err == nil {
		return e.responseSender.Send(outgoingMessage)
	}
	return nil
}

func (e *Executor) listSources(message domain.IncomingMessage) error {
	sources, err := e.redisClient.SMembers(message.UserID).Result()
	if err != nil {
		return fmt.Errorf("smembers, key: %s, err: %v", message.UserID, err)
	}

	outgoingMessage := toOutgoingMessage(message, strings.Join(sources, "\n"))
	if len(sources) == 0 {
		outgoingMessage.Text = "Источники не найдены"
	}
	if err := e.responseSender.Send(outgoingMessage); err != nil {
		return fmt.Errorf("send response: %v", err)
	}

	return nil
}

func getCommandArgs(message string) ([]string, error) {
	messageParts := strings.Split(message, " ")
	if len(messageParts) < 2 {
		return nil, fmt.Errorf("wrong command, expected 2 arguments, found: %v", messageParts)
	}

	return messageParts, nil
}

func toOutgoingMessage(src domain.IncomingMessage, text string) domain.OutgoingMessage {
	return domain.OutgoingMessage{
		UserID:      src.UserID,
		Text:        text,
		Destination: src.Source,
	}
}
