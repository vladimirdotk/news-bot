package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

const (
	telegramURL    = "https://api.telegram.org"
	sendURLPattern = "%s/bot%s/sendMessage"
)

type Sender struct {
	httpClient *http.Client
	sendURL    string
}

func NewSender(httpClient *http.Client, token string) *Sender {
	return &Sender{
		httpClient: httpClient,
		sendURL:    fmt.Sprintf(sendURLPattern, telegramURL, token),
	}
}

func (s *Sender) Send(message domain.OutgoingMessage) error {
	chatID, err := strconv.ParseInt(message.UserID, 10, 64)
	if err != nil {
		return fmt.Errorf("convert userID to chatID: %v", err)
	}

	reqBody := struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: chatID,
		Text:   message.Text,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("encode telegram request body: %v", err)
	}

	res, err := s.httpClient.Post(s.sendURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("send message to telegram: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status after sending message to telegram: %v", res.Status)
	}

	return nil
}
