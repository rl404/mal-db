package tool

import (
	"errors"
	"time"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pubsub"
	"github.com/rl404/mal-db/internal/saver"
	"gorm.io/gorm"
)

// Updater is tools to update old data.
type Updater struct {
	log   logger.Logger
	saver *saver.API
	db    *gorm.DB
	ps    pubsub.PubSub
}

// NewUpdater to create new updater tools.
func NewUpdater(l logger.Logger, s *saver.API, db *gorm.DB, ps pubsub.PubSub) *Updater {
	return &Updater{
		log:   l,
		saver: s,
		db:    db,
		ps:    ps,
	}
}

// Run to run updater tools.
func (m *Updater) Run() error {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var tmp raw.Anime
	if !errors.Is(m.db.Select("id").Where("anime_status_id = ? and (updated_at < ? or updated_at is null)", 1, today).First(&tmp).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.AnimeType, tmp.ID)
	}

	for _, t := range constant.MainTypes {
		switch t {
		case constant.AnimeType:
			var tmp raw.Anime
			if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&tmp).Error, gorm.ErrRecordNotFound) {
				return m.saver.Parse(t, tmp.ID)
			}
		case constant.MangaType:
			var tmp raw.Manga
			if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&tmp).Error, gorm.ErrRecordNotFound) {
				return m.saver.Parse(t, tmp.ID)
			}
		case constant.CharacterType:
			var tmp raw.Character
			if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&tmp).Error, gorm.ErrRecordNotFound) {
				return m.saver.Parse(t, tmp.ID)
			}
		case constant.PeopleType:
			var tmp raw.People
			if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&tmp).Error, gorm.ErrRecordNotFound) {
				return m.saver.Parse(t, tmp.ID)
			}
		}
	}

	return nil
}