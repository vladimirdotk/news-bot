package domain

// Task is an assignment for collecting news from specific source for a user.
type Task struct {
	// UserID is an ID of user who asks for news.
	UserID string
	// Source is a source of news.
	Source *Source
}
