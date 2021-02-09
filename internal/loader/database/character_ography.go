package database

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetCharacterOgraphy to get character anime/manga ography list..
func (d *Database) GetCharacterOgraphy(id int, t string, page int, limit int) (roles []model.Role, meta map[string]interface{}, code int, err error) {
	// Is empty.
	if d.isEntryEmpty(constant.CharacterType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	if errors.Is(d.db.Where("id = ?", id).First(&raw.Character{}).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(constant.CharacterType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	roles = []model.Role{}
	switch t {
	case constant.AnimeType:
		roles, code, err = d.getCharacterAnime(id)
	case constant.MangaType:
		roles, code, err = d.getCharacterManga(id)
	}
	if err != nil {
		return nil, nil, code, err
	}

	// Prepare meta.
	meta = map[string]interface{}{
		"count": len(roles),
	}

	return roles, meta, http.StatusOK, nil
}

func (d *Database) getCharacterAnime(id int) (roles []model.Role, code int, err error) {
	err = d.db.Table(fmt.Sprintf("%s as ac", raw.AnimeCharacter{}.TableName())).
		Select("ac.anime_id as id, a.title as name, a.image_url as image, ac.role").
		Joins(fmt.Sprintf("left join %s as a on a.id = ac.anime_id", raw.Anime{}.TableName())).
		Where("ac.character_id = ?", id).
		Group("ac.anime_id, a.title, a.image_url, ac.role").
		Order("ac.role asc, ac.anime_id asc").
		Find(&roles).Error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return roles, http.StatusOK, nil
}

func (d *Database) getCharacterManga(id int) (roles []model.Role, code int, err error) {
	err = d.db.Table(fmt.Sprintf("%s as mc", raw.MangaCharacter{}.TableName())).
		Select("mc.manga_id as id, m.title as name, m.image_url as image, mc.role").
		Joins(fmt.Sprintf("left join %s as m on m.id = mc.manga_id", raw.Manga{}.TableName())).
		Where("mc.character_id = ?", id).
		Group("mc.manga_id, m.title, m.image_url, mc.role").
		Order("mc.role asc, mc.manga_id asc").
		Find(&roles).Error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return roles, http.StatusOK, nil
}
