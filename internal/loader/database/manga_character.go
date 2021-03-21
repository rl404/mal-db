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

// GetMangaCharacter to get manga character list.
func (d *Database) GetMangaCharacter(id int, page int, limit int) (roles []model.Role, meta map[string]interface{}, code int, err error) {
	// Is empty.
	if d.isEntryEmpty(constant.MangaType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	if errors.Is(d.db.Where("id = ?", id).First(&raw.Manga{}).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(constant.MangaType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Prepare query.
	err = d.db.Table(fmt.Sprintf("%s as mc", raw.MangaCharacter{}.TableName())).
		Select("mc.character_id as id, c.name, c.image_url as image, mc.role").
		Joins(fmt.Sprintf("left join %s as c on c.id = mc.character_id", raw.Character{}.TableName())).
		Where("mc.manga_id = ?", id).
		Order("mc.role asc, c.name asc").
		Limit(limit).
		Offset(limit * (page - 1)).
		Find(&roles).Error
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	// Prepare meta.
	meta = map[string]interface{}{
		"count": len(roles),
	}

	return roles, meta, http.StatusOK, nil
}
