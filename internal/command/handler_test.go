package command

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vladimirdotk/news-bot/internal/command/mocks"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

func TestHandler_Handle(t *testing.T) {
	testCases := []struct {
		name            string
		handlerFunc     func(mc minimock.MockController, t *testing.T) *Handler
		incomingMessage *domain.IncomingMessage
		err             error
	}{
		{
			name: "Simple success case",
			handlerFunc: func(mc minimock.MockController, t *testing.T) *Handler {
				commandBytes, err := json.Marshal(&domain.IncomingMessage{
					ID:       "1",
					UserID:   "u1",
					Username: "uname1",
					Text:     "Hello World",
					Source:   domain.MessageSystemTelegram,
				})
				require.NoError(t, err)
				queueService := mocks.NewQueueServiceMock(t)
				queueService.PublishMock.
					Expect(context.TODO(), "incoming_command", commandBytes).
					Return(nil)

				return &Handler{queueService: queueService}
			},
			incomingMessage: &domain.IncomingMessage{
				ID:       "1",
				UserID:   "u1",
				Username: "uname1",
				Text:     "Hello World",
				Source:   domain.MessageSystemTelegram,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mc := minimock.NewController(t)

			handler := tc.handlerFunc(mc, t)
			err := handler.Handle(context.TODO(), tc.incomingMessage)
			assert.Equal(t, tc.err, err)
		})
	}

}
