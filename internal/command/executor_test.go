package command

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vladimirdotk/news-bot/internal/command/mocks"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

func TestExecutor_Exec(t *testing.T) {
	testCases := []struct {
		name            string
		executorFunc    func(mc minimock.MockController, t *testing.T) *Executor
		incomingMessage domain.IncomingMessage
		err             error
	}{
		{
			name: "Exec add source cmd, source not exists, will be created",
			executorFunc: func(mc minimock.MockController, t *testing.T) *Executor {
				sourceDetector := mocks.NewSourceDetectorMock(t)
				sourceDetector.DetectMock.
					Expect(context.TODO(), "https://habr.com/ru/rss/all/all/").
					Return(domain.SourceTypeRSS)

				sourceJSON, err := domain.SourceToJSON(&domain.Source{
					URL:  "https://habr.com/ru/rss/all/all/",
					Type: domain.SourceTypeRSS,
				})
				require.NoError(t, err)

				key := domain.UserSourceKey("u1", "https://habr.com/ru/rss/all/all/")

				client, mock := redismock.NewClientMock()

				mock.ExpectExists(key).SetVal(0)
				mock.ExpectSAdd(key, sourceJSON).SetVal(0)

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(
					context.TODO(),
					domain.OutgoingMessage{
						UserID:      "u1",
						Text:        "Source added successfully",
						Destination: domain.MessageSystemTelegram,
					},
				).
					Return(nil)

				return &Executor{
					redisClient:    client,
					responseSender: responseSender,
					sourceDetector: sourceDetector,
				}
			},
			incomingMessage: domain.IncomingMessage{
				ID:       "1",
				UserID:   "u1",
				Username: "uname1",
				Text:     "/add https://habr.com/ru/rss/all/all/",
				Source:   domain.MessageSystemTelegram,
			},
		},
		{
			name: "Exec add source cmd, source exists, will not be created",
			executorFunc: func(mc minimock.MockController, t *testing.T) *Executor {
				sourceJSON, err := domain.SourceToJSON(&domain.Source{
					URL:  "https://habr.com/ru/rss/all/all/",
					Type: domain.SourceTypeRSS,
				})
				require.NoError(t, err)

				key := domain.UserSourceKey("u1", "https://habr.com/ru/rss/all/all/")

				client, mock := redismock.NewClientMock()
				mock.ExpectExists(key).SetVal(1)
				mock.ExpectSAdd(key, sourceJSON).SetVal(0)

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(
					context.TODO(),
					domain.OutgoingMessage{
						UserID:      "u1",
						Text:        "Source added successfully",
						Destination: domain.MessageSystemTelegram,
					},
				).
					Return(nil)

				return &Executor{
					redisClient:    client,
					responseSender: responseSender,
				}
			},
			incomingMessage: domain.IncomingMessage{
				ID:       "1",
				UserID:   "u1",
				Username: "uname1",
				Text:     "/add https://habr.com/ru/rss/all/all/",
				Source:   domain.MessageSystemTelegram,
			},
		},
		{
			name: "Exec list sources cmd",
			executorFunc: func(mc minimock.MockController, t *testing.T) *Executor {
				client, mock := redismock.NewClientMock()
				mock.
					ExpectSMembers(domain.UserSourcesSearchKey("u1")).
					SetVal(
						[]string{
							`{"url":"https://news.yandex.ru/health.rss","type":"RSS"}`,
							`{"url":"https://habr.com/ru/rss/all/all/","type":"RSS"}`,
						},
					)

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(
					context.TODO(),
					domain.OutgoingMessage{
						UserID:      "u1",
						Text:        "https://news.yandex.ru/health.rss\nhttps://habr.com/ru/rss/all/all/",
						Destination: domain.MessageSystemTelegram,
					},
				).
					Return(nil)

				return &Executor{
					redisClient:    client,
					responseSender: responseSender,
				}
			},
			incomingMessage: domain.IncomingMessage{
				ID:       "1",
				UserID:   "u1",
				Username: "uname1",
				Text:     "/list",
				Source:   domain.MessageSystemTelegram,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mc := minimock.NewController(t)

			handler := tc.executorFunc(mc, t)
			err := handler.Exec(context.TODO(), tc.incomingMessage)
			assert.Equal(t, tc.err, err)
		})
	}
}
