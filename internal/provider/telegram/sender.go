package telegram

import (
	"bytes"
	"context"
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

func (s *Sender) Send(ctx context.Context, message domain.OutgoingMessage) error {
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

	req, err := http.NewRequest("POST", s.sendURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("create http reqeust: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := s.httpClient.Do(
		req.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("send message to telegram: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status after sending message to telegram: %v", res.Status)
	}

	return nil
}
