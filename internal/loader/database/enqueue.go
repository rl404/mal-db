package database

import (
	"errors"
	"net/http"
	"time"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// Enqueue to enqueue to be re-parsed.
func (d *Database) Enqueue(t string, id int) (int, error) {
	if d.isEntryNew(t, id) {
		return http.StatusBadRequest, _errors.ErrNewData
	}
	if err := d.enqueue(t, id); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func (d *Database) isEntryNew(t string, id int) bool {
	today := time.Now()
	switch t {
	case constant.AnimeType:
		var anime raw.Anime
		if errors.Is(d.db.Select("id, updated_at").Where("id = ?", id).First(&anime).Error, gorm.ErrRecordNotFound) {
			return false
		}
		return today.Sub(anime.UpdatedAt) < d.ageLimit
	case constant.MangaType:
		var manga raw.Manga
		if errors.Is(d.db.Select("id, updated_at").Where("id = ?", id).First(&manga).Error, gorm.ErrRecordNotFound) {
			return false
		}
		return today.Sub(manga.UpdatedAt) < d.ageLimit
	case constant.CharacterType:
		var char raw.Character
		if errors.Is(d.db.Select("id, updated_at").Where("id = ?", id).First(&char).Error, gorm.ErrRecordNotFound) {
			return false
		}
		return today.Sub(char.UpdatedAt) < d.ageLimit
	case constant.PeopleType:
		var people raw.People
		if errors.Is(d.db.Select("id, updated_at").Where("id = ?", id).First(&people).Error, gorm.ErrRecordNotFound) {
			return false
		}
		return today.Sub(people.UpdatedAt) < d.ageLimit
	default:
		return true
	}
}
