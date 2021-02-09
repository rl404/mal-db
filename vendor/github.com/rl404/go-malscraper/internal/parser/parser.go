package parser

import (
	"net/http"
	"time"

	"github.com/rl404/go-malscraper/internal/parser/anime"
	"github.com/rl404/go-malscraper/internal/parser/article"
	"github.com/rl404/go-malscraper/internal/parser/character"
	"github.com/rl404/go-malscraper/internal/parser/club"
	"github.com/rl404/go-malscraper/internal/parser/genre"
	"github.com/rl404/go-malscraper/internal/parser/manga"
	"github.com/rl404/go-malscraper/internal/parser/news"
	"github.com/rl404/go-malscraper/internal/parser/people"
	producermagazine "github.com/rl404/go-malscraper/internal/parser/producer_magazine"
	"github.com/rl404/go-malscraper/internal/parser/recommendation"
	"github.com/rl404/go-malscraper/internal/parser/review"
	"github.com/rl404/go-malscraper/internal/parser/search"
	"github.com/rl404/go-malscraper/internal/parser/top"
	"github.com/rl404/go-malscraper/internal/parser/user"
	"github.com/rl404/go-malscraper/service"
)

// Requester is mockable http client.
type Requester interface {
	Do(*http.Request) (*http.Response, error)
}

// Parser parse MyAnimeList web amd convert to easy-to-use data.
type Parser struct {
	anime          anime.Parser
	manga          manga.Parser
	character      character.Parser
	people         people.Parser
	producer       producermagazine.Parser
	genre          genre.Parser
	review         review.Parser
	recommendation recommendation.Parser
	news           news.Parser
	article        article.Parser
	club           club.Parser
	top            top.Parser
	user           user.Parser
	search         search.Parser
	logger         service.Logger
	http           Requester
}

// New to create new parser.
func New(cleanImg, cleanVid bool, l service.Logger) service.API {
	return &Parser{
		anime:          anime.New(cleanImg, cleanVid),
		manga:          manga.New(cleanImg),
		character:      character.New(cleanImg),
		people:         people.New(cleanImg),
		producer:       producermagazine.New(cleanImg),
		genre:          genre.New(cleanImg),
		review:         review.New(cleanImg),
		recommendation: recommendation.New(cleanImg),
		news:           news.New(cleanImg),
		article:        article.New(cleanImg),
		club:           club.New(cleanImg),
		top:            top.New(cleanImg),
		user:           user.New(cleanImg),
		search:         search.New(cleanImg),
		logger:         l,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
