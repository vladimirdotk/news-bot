package source

import (
	"context"

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

func (d *Detector) Detect(ctx context.Context, sourceURL string) domain.SourceType {
	_, err := d.rssParser.ParseURL(sourceURL)
	if err != nil {
		return domain.SourceTypeUnknown
	}

	return domain.SourceTypeRSS
}
