package domain

type QueueTopic string

func (qt QueueTopic) String() string {
	return string(qt)
}

const QueueTopicIncomingCommand QueueTopic = "incoming_command"
