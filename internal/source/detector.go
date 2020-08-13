package source

import (
	"github.com/mmcdole/gofeed"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

// TODO: make more abstract
type Detector struct {
	rssParser *gofeed.Parser
}

func NewDetector() *Detector {
	return &Detector{
		rssParser: gofeed.NewParser(),
	}
}

func (d *Detector) Detect(sourceURL string) domain.SourceType {
	_, err := d.rssParser.ParseURL(sourceURL)
	if err != nil {
		return domain.SourceTypeUnknown
	}

	return domain.SourceTypeRSS
}
