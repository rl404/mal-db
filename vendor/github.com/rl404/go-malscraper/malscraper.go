package malscraper

import (
	"time"

	"github.com/rl404/go-malscraper/internal/cacher"
	"github.com/rl404/go-malscraper/internal/parser"
	"github.com/rl404/go-malscraper/internal/validator"
	"github.com/rl404/go-malscraper/service"
	"github.com/rl404/mal-plugin/cache/nocache"
)

// Malscraper is malscraper instance which contains all
// methods to parse MyAnimeList web page.
type Malscraper struct {
	api    service.API
	cacher service.Cacher
}

// New to create new malscraper with config.
func New(cfg Config) (*Malscraper, error) {
	// Init config.
	if err := cfg.init(); err != nil {
		return nil, err
	}

	// Init the core of malscraper which access and parse
	// MyAnimeList web.
	api := parser.New(cfg.CleanImageURL, cfg.CleanVideoURL, cfg.Logger)

	// Init cacher which intercepts request to check to
	// cache first before actually access and parse MyAnimeList.
	api = cacher.New(api, cfg.Cacher, cfg.Logger)

	// Init validator which validates requested params
	// before processing the request.
	api = validator.New(api, cfg.Cacher, cfg.Logger)

	return &Malscraper{
		api:    api,
		cacher: cfg.Cacher,
	}, nil
}

// NewDefault to quickly create new malscraper with
// default config. Will cache for 1 day as default.
func NewDefault() (*Malscraper, error) {
	return New(Config{
		CacheTime:     24 * time.Hour,
		CleanImageURL: true,
		CleanVideoURL: true,
		LogColor:      true,
		LogLevel:      LevelDefault,
	})
}

// NewNoCache to create new malscraper without caching.
func NewNoCache() (*Malscraper, error) {
	c, _ := nocache.New()
	return New(Config{
		Cacher:        c,
		CleanImageURL: true,
		CleanVideoURL: true,
		LogColor:      true,
		LogLevel:      LevelDefault,
	})
}

// Close to close cache connection if exists.
func (m *Malscraper) Close() error {
	if m.cacher == nil {
		return nil
	}
	return m.cacher.Close()
}
