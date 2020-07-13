package natsserver

import nats "github.com/nats-io/nats.go"

type QueueService struct {
	natsClient *nats.EncodedConn
}

func NewQueueService(natsClient *nats.EncodedConn) *QueueService {
	return &QueueService{natsClient: natsClient}
}

func (q *QueueService) Publish(topic string, data interface{}) error {
	return q.natsClient.Publish(topic, data)
}
