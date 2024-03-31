package source

import (
	"context"
	"log/slog"

	"github.com/mmcdole/gofeed"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

// TODO: make more abstract
type Detector struct {
	rssParser *gofeed.Parser
	log       *slog.Logger
}

func NewDetector(log *slog.Logger) *Detector {
	return &Detector{
		rssParser: gofeed.NewParser(),
		log:       log,
	}
}

func (d *Detector) Detect(ctx context.Context, sourceURL string) domain.SourceType {
	_, err := d.rssParser.ParseURL(sourceURL)
	if err != nil {
		return domain.SourceTypeUnknown
	}

	return domain.SourceTypeRSS
}
