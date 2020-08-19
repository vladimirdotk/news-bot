package domain

import "fmt"

// Source is a source of news.
type Source struct {
	// URL is a source location.
	URL string `json:"url"`
	// Type is source type. Like RSS.
	Type SourceType `json:"type"`
}

// SourceType is a type of news source.
type SourceType string

func (t SourceType) String() string {
	return string(t)
}

const (
	SourceTypeUnknown SourceType = "UNKNOWN"
	SourceTypeRSS     SourceType = "RSS"
)

// UserSourcesKey returns key to store user sources in a storage.
func UserSourcesKey(userID string) string {
	return fmt.Sprintf("user_sources|%s", userID)
}
