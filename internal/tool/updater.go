package tool

import (
	"errors"
	"strconv"
	"time"

	"github.com/rl404/go-malscraper/pkg/utils"
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
	lastWeek := now.AddDate(0, 0, -7)
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	season, seasonStart, seasonEnd := utils.GetCurrentSeason(), 0, 0
	switch season {
	case "winter":
		seasonStart, seasonEnd = 1, 4
	case "spring":
		seasonStart, seasonEnd = 4, 7
	case "summer":
		seasonStart, seasonEnd = 7, 10
	case "fall":
		seasonStart, seasonEnd = 10, 13
	}

	// Weekly airing or current season anime update.
	var anime raw.Anime
	if !errors.Is(m.db.Select("id").Where("(((premiered = '' and start_year = ? and start_month >= ? and start_month < ?) or (premiered != '' and split_part(premiered, ' ', 1) = ? and split_part(premiered, ' ', 2) = ?)) or anime_status_id = ?) and (updated_at < ? or updated_at is null)", now.Year(), seasonStart, seasonEnd, season, strconv.Itoa(now.Year()), 1, lastWeek).First(&anime).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.AnimeType, anime.ID)
	}

	// Monthly data update.
	if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&anime).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.AnimeType, anime.ID)
	}

	var manga raw.Manga
	if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&manga).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.MangaType, manga.ID)
	}

	var people raw.People
	if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&people).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.PeopleType, people.ID)
	}

	// Fill missing data.
	for _, t := range constant.MainTypes {
		var maxID int
		if err := m.db.Table(t).Select("max(id)").Row().Scan(&maxID); err != nil {
			return err
		}

		entryMap := make(map[int]int)
		if err := m.setEntryMap(t, entryMap); err != nil {
			return err
		}
		if err := m.setEmptyMap(t, entryMap); err != nil {
			return err
		}

		for id := 1; id < maxID; id++ {
			if entryMap[id] == 0 {
				return m.saver.Parse(t, id)
			}
		}
	}

	// Update character (low priority).
	var char raw.Character
	if !errors.Is(m.db.Select("id").Where("updated_at < ? or updated_at is null", currentMonth).First(&char).Error, gorm.ErrRecordNotFound) {
		return m.saver.Parse(constant.CharacterType, char.ID)
	}

	return nil
}

func (m *Updater) setEntryMap(t string, em map[int]int) error {
	rows, err := m.db.Table(t).Select("id").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp struct {
			ID int
		}
		if err = m.db.ScanRows(rows, &tmp); err != nil {
			return err
		}
		em[tmp.ID] = tmp.ID
	}
	return nil
}

func (m *Updater) setEmptyMap(t string, em map[int]int) error {
	rows, err := m.db.Model(&raw.EmptyID{}).Where("type = ?", t).Select("id").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp struct {
			ID int
		}
		if err = m.db.ScanRows(rows, &tmp); err != nil {
			return err
		}
		em[tmp.ID] = tmp.ID
	}
	return nil
}
