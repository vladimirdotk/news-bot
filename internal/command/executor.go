package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Executor struct {
	redisClient    *redis.Client
	responseSender ResponseSender
	sourceDetector SourceDetector
}

func NewExecutor(redisClient *redis.Client, responseSender ResponseSender, sourceDetector SourceDetector) *Executor {
	return &Executor{
		redisClient:    redisClient,
		responseSender: responseSender,
		sourceDetector: sourceDetector,
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

// addSource validates and adds source to user's list (if valid).
func (e *Executor) addSource(message domain.IncomingMessage) error {
	commandArgs, err := getCommandArgs(message.Text)
	if err != nil {
		return fmt.Errorf("get command args: %v", err)
	}

	sourceType := e.sourceDetector.Detect(commandArgs[1])
	if sourceType == domain.SourceTypeUnknown {
		return fmt.Errorf("unknown source by URL: %v", commandArgs[1])
	}

	sourceJSON, err := sourceToJSON(&domain.Source{
		URL:  commandArgs[1],
		Type: sourceType,
	})
	if err != nil {
		return fmt.Errorf("source to JSON: %v", err)
	}

	if err := e.redisClient.SAdd(message.UserID, sourceJSON).Err(); err != nil {
		return fmt.Errorf("sadd, key: %s, value: %s, err: %v", message.UserID, string(sourceJSON), err)
	}

	outgoingMessage := toOutgoingMessage(message, "Источник добавлен")

	if err := e.responseSender.Send(outgoingMessage); err != nil {
		return fmt.Errorf("send response: %v", err)
	}

	return nil
}

// listSources sends user's sources list if any.
func (e *Executor) listSources(message domain.IncomingMessage) error {
	sources, err := e.redisClient.SMembers(message.UserID).Result()
	if err != nil {
		return fmt.Errorf("smembers, key: %s, err: %v", message.UserID, err)
	}

	sourcesURLs := make([]string, len(sources))
	for i, source := range sources {
		domainSource, err := sourceFromJSON(source)
		if err != nil {
			return fmt.Errorf("source from JSON: %v", err)
		}
		sourcesURLs[i] = domainSource.URL
	}

	outgoingMessage := toOutgoingMessage(message, strings.Join(sourcesURLs, "\n"))
	if len(sources) == 0 {
		outgoingMessage.Text = "Источники не найдены"
	}
	if err := e.responseSender.Send(outgoingMessage); err != nil {
		return fmt.Errorf("send response: %v", err)
	}

	return nil
}

// getCommandArgs returns slice of command args
// TODO: make more clear and reusable
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

func sourceToJSON(src *domain.Source) ([]byte, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("marshal source: %v", err)
	}

	return b, nil
}

func sourceFromJSON(src string) (*domain.Source, error) {
	var source domain.Source
	if err := json.Unmarshal([]byte(src), &source); err != nil {
		return nil, fmt.Errorf("unmarshal source: %v", err)
	}

	return &source, nil
}
