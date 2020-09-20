package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Source is a source of news.
type Source struct {
	// URL is a source location.
	URL string `json:"url"`
	// Type is source type. Like RSS.
	Type SourceType `json:"type"`
	// LastNewsPublishTime indicates time of last collected source data.
	LastNewsPublishTime time.Time
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

// UserSourcesKey returns key fot user sources storage.
func UserSourcesKey(userID string) string {
	return fmt.Sprintf("user_sources|%s", userID)
}

// ExtractUserFromKey returns userID extracted from user source key.
func ExtractUserFromKey(sourceKey string) (string, error) {
	keyParts := strings.Split(sourceKey, "|")
	if len(keyParts) != 2 {
		return "", fmt.Errorf("invalid source key: %s", sourceKey)
	}

	return keyParts[1], nil
}

func SourceToJSON(src *Source) ([]byte, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("marshal source: %v", err)
	}

	return b, nil
}

func SourceFromJSON(src string) (*Source, error) {
	var source Source
	if err := json.Unmarshal([]byte(src), &source); err != nil {
		return nil, fmt.Errorf("unmarshal source: %v", err)
	}

	return &source, nil
}
