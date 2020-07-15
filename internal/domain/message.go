package domain

type IncomingMessage struct {
	ID       string        `json:"id"`
	UserID   string        `json:"user_id"`
	Username string        `json:"username"`
	Text     string        `json:"text"`
	Source   MessageSource `json:"source"`
}

type MessageSource string

func (m MessageSource) String() string {
	return string(m)
}

const MessageSourceTelegram MessageSource = "telegram"
