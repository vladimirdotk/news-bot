package command

import (
	"github.com/vladimirdotk/news-bot/internal/domain"
)

//go:generate minimock -g -i github.com/vladimirdotk/news-bot/internal/command.ResponseSender -o ./mocks -s "_mock.go"

// ResponseSender describes message sender.
type ResponseSender interface {
	// Send sends message, returns error if any.
	Send(message domain.OutgoingMessage) error
}

//go:generate minimock -g -i github.com/vladimirdotk/news-bot/internal/command.QueueService -o ./mocks -s "_mock.go"

// QueueService describes a service for working with queue.
type QueueService interface {
	// Publish sets message with certain topic to queue.
	Publish(topic domain.QueueTopic, data interface{}) error
}

//go:generate minimock -g -i github.com/vladimirdotk/news-bot/internal/command.SourceDetector -o ./mocks -s "_mock.go"

// SourceDetector describes a service that detects different sources types.
type SourceDetector interface {
	// Detect finds out and return source type.
	Detect(sourceURL string) domain.SourceType
}
