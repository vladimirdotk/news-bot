package domain

type Source struct {
	URL  string     `json:"url"`
	Type SourceType `json:"type"`
}

type SourceType string

func (t SourceType) String() string {
	return string(t)
}

const (
	SourceTypeUnknown SourceType = "UNKNOWN"
	SourceTypeRSS     SourceType = "RSS"
)
