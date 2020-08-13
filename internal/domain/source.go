package domain

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
