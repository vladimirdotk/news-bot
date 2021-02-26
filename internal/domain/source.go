package domain

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

// Source is a source of news.
type Source struct {
	// URL is a source location.
	URL string `json:"url"`
	// Type is source type. Like RSS.
	Type SourceType `json:"type"`
	// LastSeenID is the last seen ID of the source item.
	LastSeenID string `json:"last_seen_id"`
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

// UserSourceKey returns key for user source storage.
func UserSourceKey(userID, source string) string {
	hashSum := md5.Sum([]byte(source))
	return fmt.Sprintf("%s|%s", userID, string(hashSum[:]))
}

// UserSourcesKey returns key for searching user sources.
func UserSourcesSearchKey(userID string) string {
	return fmt.Sprintf("%s|*", userID)
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
