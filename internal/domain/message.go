package domain

// IncomingMessage describes message received by system.
type IncomingMessage struct {
	// ID is a message identity.
	ID string `json:"id"`
	// UserID is an ID of user who sent message.
	UserID string `json:"user_id"`
	// Username is a name of user who sent message.
	Username string `json:"username"`
	// Text is a message core.
	Text string `json:"text"`
	// Source is a message system from where message was received.
	Source MessageSystem `json:"source"`
}

const (
	MessageCommandAddSource  = "/add"
	MessageCommandListSource = "/list"
)

// MessageSystem is a system for sending/receiving messages.
type MessageSystem string

func (m MessageSystem) String() string {
	return string(m)
}

const MessageSystemTelegram MessageSystem = "telegram"

// OutgoingMessage describes message that will be send.
type OutgoingMessage struct {
	// UserID is an ID of user to whom message will be send.
	UserID string
	// Text is a message core.
	Text string
	// Source is a message system to where message will be send.
	Destination MessageSystem
}
