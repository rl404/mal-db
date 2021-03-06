package saver

import (
	"errors"
	"fmt"
	"time"

	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-db/internal/cacher"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pubsub"
	"github.com/rl404/mal-plugin/log/elasticsearch"
	"gorm.io/gorm"
)

// API is saver api.
type API struct {
	logger logger.Logger
	cacher cacher.Cacher
	db     *gorm.DB
	mal    *malscraper.Malscraper
	pubsub pubsub.PubSub
	es     *elasticsearch.Client
}

// New to create new saver.
func New(l logger.Logger, c cacher.Cacher, db *gorm.DB, mal *malscraper.Malscraper, ps pubsub.PubSub, es *elasticsearch.Client) *API {
	return &API{
		logger: l,
		cacher: c,
		db:     db,
		mal:    mal,
		pubsub: ps,
		es:     es,
	}
}

type queueLog struct {
	Type      string    `json:"type"`
	ID        int       `json:"id"`
	Error     error     `json:"error"`
	CreatedAt time.Time `json:"created_at"`
}

// Parse to parse entry.
func (a *API) Parse(t string, id int) (err error) {
	switch t {
	case constant.AnimeType:
		err = a.parseAnime(id)
	case constant.MangaType:
		err = a.parseManga(id)
	case constant.CharacterType:
		err = a.parseCharacter(id)
	case constant.PeopleType:
		err = a.parsePeople(id)
	default:
		err = fmt.Errorf("invalid type [%s:%v]", t, id)
	}

	if err != nil {
		// Re-queue if error.
		if errQ := a.enqueue(t, id); errQ != nil {
			a.logger.Error(errQ.Error())
		}
	}

	if a.es == nil {
		return err
	}

	if errEs := a.es.Send("mal-db-parse", queueLog{
		Type:      t,
		ID:        id,
		Error:     err,
		CreatedAt: time.Now(),
	}); errEs != nil {
		a.logger.Error(errEs.Error())
	}

	return err
}

func (a *API) isEntryExist(t string, id int) bool {
	switch t {
	case constant.AnimeType:
		if !errors.Is(a.db.Select("id").Where("id = ?", id).First(&raw.Anime{}).Error, gorm.ErrRecordNotFound) {
			return true
		}
	case constant.MangaType:
		if !errors.Is(a.db.Select("id").Where("id = ?", id).First(&raw.Manga{}).Error, gorm.ErrRecordNotFound) {
			return true
		}
	case constant.CharacterType:
		if !errors.Is(a.db.Select("id").Where("id = ?", id).First(&raw.Character{}).Error, gorm.ErrRecordNotFound) {
			return true
		}
	case constant.PeopleType:
		if !errors.Is(a.db.Select("id").Where("id = ?", id).First(&raw.People{}).Error, gorm.ErrRecordNotFound) {
			return true
		}
	}
	return !errors.Is(a.db.Where("type = ? and id = ?", t, id).First(&raw.EmptyID{}).Error, gorm.ErrRecordNotFound)
}

func (a *API) enqueue(t string, id int) error {
	if a.isEntryExist(t, id) {
		return nil
	}
	err := a.pubsub.Publish(constant.PubSubTopic, pubsub.Message{
		Type: t,
		ID:   id,
	})
	if a.es == nil {
		return err
	}
	if errEs := a.es.Send("mal-db-queue", queueLog{
		Type:      t,
		ID:        id,
		Error:     err,
		CreatedAt: time.Now(),
	}); errEs != nil {
		a.logger.Error(errEs.Error())
	}
	return err
}
