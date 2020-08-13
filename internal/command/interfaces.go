package command

import "github.com/vladimirdotk/news-bot/internal/domain"

// ResponseSender describes message sender.
type ResponseSender interface {
	// Send sends message, returns error if any.
	Send(message domain.OutgoingMessage) error
}

// QueueService describes a service for working with queue.
type QueueService interface {
	// Publish sets message with certain topic to queue.
	Publish(topic string, data interface{}) error
}

// SourceDetector describes a service that detects different sources types.
type SourceDetector interface {
	// Detect finds out and return source type.
	Detect(sourceURL string) domain.SourceType
}
