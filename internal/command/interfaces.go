package command

import "github.com/vladimirdotk/news-bot/internal/domain"

type ResponseSender interface {
	Send(message domain.OutgoingMessage) error
}

type QueueService interface {
	Publish(topic string, data interface{}) error
}

type SourceDetector interface {
	Detect(sourceURL string) domain.SourceType
}
