package command

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Executor struct {
	redisClient    redis.Cmdable
	responseSender ResponseSender
	sourceDetector SourceDetector
	log            *slog.Logger
}

func NewExecutor(
	redisClient redis.Cmdable,
	responseSender ResponseSender,
	sourceDetector SourceDetector,
	log *slog.Logger,
) *Executor {
	return &Executor{
		redisClient:    redisClient,
		responseSender: responseSender,
		sourceDetector: sourceDetector,
		log: slog.With(
			slog.Group("command executor"),
		),
	}
}

func (e *Executor) Exec(ctx context.Context, message domain.IncomingMessage) error {
	if strings.HasPrefix(message.Text, domain.MessageCommandAddSource) {
		return e.addSource(ctx, message)
	}

	if strings.HasPrefix(message.Text, domain.MessageCommandListSource) {
		return e.listSources(ctx, message)
	}

	return fmt.Errorf("executor not found for command: %s", message)
}

// addSource validates and adds source to user's list (if valid).
func (e *Executor) addSource(ctx context.Context, message domain.IncomingMessage) error {
	commandArgs, err := getCommandArgs(message.Text, 2)
	if err != nil {
		return fmt.Errorf("get command args: %v", err)
	}

	url := commandArgs[1]

	key := domain.UserSourceKey(message.UserID, url)
	exits, err := e.keyExists(ctx, key)
	if err != nil {
		return fmt.Errorf("key exists: %v", err)
	}
	if exits {
		return e.sendSuccessMessage(ctx, message)
	}

	sourceType := e.sourceDetector.Detect(ctx, url)
	if sourceType == domain.SourceTypeUnknown {
		return fmt.Errorf("unknown source by URL: %v", url)
	}

	sourceJSON, err := domain.SourceToJSON(&domain.Source{
		URL:  url,
		Type: sourceType,
	})
	if err != nil {
		return fmt.Errorf("source to JSON: %v", err)
	}

	if err := e.redisClient.SAdd(ctx, key, sourceJSON).Err(); err != nil {
		return fmt.Errorf("sadd, key: %s, value: %s, err: %v", message.UserID, string(sourceJSON), err)
	}

	return e.sendSuccessMessage(ctx, message)
}

// listSources sends user's sources list if any.
func (e *Executor) listSources(ctx context.Context, message domain.IncomingMessage) error {
	key := domain.UserSourcesSearchKey(message.UserID)

	sources, err := e.redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("smembers, key: %s, err: %v", key, err)
	}

	sourcesURLs := make([]string, len(sources))
	for i, source := range sources {
		domainSource, err := domain.SourceFromJSON(source)
		if err != nil {
			return fmt.Errorf("source from JSON: %v", err)
		}
		sourcesURLs[i] = domainSource.URL
	}

	outgoingMessage := toOutgoingMessage(message, strings.Join(sourcesURLs, "\n"))
	if len(sources) == 0 {
		outgoingMessage.Text = "Sources not found"
	}

	return e.responseSender.Send(ctx, outgoingMessage)
}

// getCommandArgs returns slice of command args
func getCommandArgs(message string, argsCount int) ([]string, error) {
	messageParts := strings.Split(message, " ")
	if len(messageParts) != argsCount {
		return nil, fmt.Errorf("wrong command, expected %d arguments, found: %v", argsCount, messageParts)
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

func (e *Executor) sendSuccessMessage(ctx context.Context, message domain.IncomingMessage) error {
	outgoingMessage := toOutgoingMessage(message, "Source added successfully")

	if err := e.responseSender.Send(ctx, outgoingMessage); err != nil {
		return fmt.Errorf("send response: %v", err)
	}

	return nil
}

func (e *Executor) keyExists(ctx context.Context, key string) (bool, error) {
	res := e.redisClient.Exists(ctx, key)

	if res.Err() != nil {
		return false, fmt.Errorf("exists, key: %s, err: %v", key, res.Err())
	}
	if res.Val() == 1 {
		return true, nil
	}

	return false, nil
}
