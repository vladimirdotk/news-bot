package command

import (
	"testing"

	"github.com/elliotchance/redismock/v7"
	"github.com/go-redis/redis/v7"
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
					Expect("https://habr.com/ru/rss/all/all/").
					Return(domain.SourceTypeRSS)

				sourceJSON, err := domain.SourceToJSON(&domain.Source{
					URL:  "https://habr.com/ru/rss/all/all/",
					Type: domain.SourceTypeRSS,
				})
				require.NoError(t, err)

				key := domain.UserSourceKey("u1", "https://habr.com/ru/rss/all/all/")

				redisClient := redismock.NewMock()
				redisClient.
					On("SAdd", key, []interface{}{sourceJSON}).
					Return(redis.NewIntCmd())

				redisClient.
					On("Exists", []string{key}).
					Return(redis.NewIntResult(0, nil))

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(domain.OutgoingMessage{
					UserID:      "u1",
					Text:        "Источник добавлен",
					Destination: domain.MessageSystemTelegram,
				}).
					Return(nil)

				return &Executor{
					redisClient:    redisClient,
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

				redisClient := redismock.NewMock()
				redisClient.
					On("SAdd", key, []interface{}{sourceJSON}).
					Return(redis.NewIntCmd())

				redisClient.
					On("Exists", []string{key}).
					Return(redis.NewIntResult(1, nil))

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(domain.OutgoingMessage{
					UserID:      "u1",
					Text:        "Источник добавлен",
					Destination: domain.MessageSystemTelegram,
				}).
					Return(nil)

				return &Executor{
					redisClient:    redisClient,
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
				redisClient := redismock.NewMock()
				redisClient.
					On("SMembers", domain.UserSourcesSearchKey("u1")).
					Return(redis.NewStringSliceResult([]string{
						`{"url":"https://news.yandex.ru/health.rss","type":"RSS"}`,
						`{"url":"https://habr.com/ru/rss/all/all/","type":"RSS"}`,
					}, nil))

				responseSender := mocks.NewResponseSenderMock(t)
				responseSender.SendMock.Expect(domain.OutgoingMessage{
					UserID:      "u1",
					Text:        "https://news.yandex.ru/health.rss\nhttps://habr.com/ru/rss/all/all/",
					Destination: domain.MessageSystemTelegram,
				}).
					Return(nil)

				return &Executor{
					redisClient:    redisClient,
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
			defer mc.Finish()

			handler := tc.executorFunc(mc, t)
			err := handler.Exec(tc.incomingMessage)
			assert.Equal(t, tc.err, err)
		})
	}
}
