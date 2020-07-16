package domain

type IncomingMessage struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Source   System `json:"system"`
}

const (
	MessageCommandAddSource  = "/add"
	MessageCommandListSource = "/list"
)

type System string

func (m System) String() string {
	return string(m)
}

const SystemTelegram System = "telegram"

type OutgoingMessage struct {
	UserID      string
	Text        string
	Destination System
}
