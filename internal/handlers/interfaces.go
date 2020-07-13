package handlers

type QueueService interface {
	Publish(topic string, data interface{}) error
}
