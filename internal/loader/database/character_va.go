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

// GetCharacterVA to get character voice actor list.
func (d *Database) GetCharacterVA(id int, page int, limit int) (roles []model.Role, meta map[string]interface{}, code int, err error) {
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

	// Prepare query.
	err = d.db.Table(fmt.Sprintf("%s as ac", raw.AnimeCharacter{}.TableName())).
		Select("ac.people_id as id, p.name, p.image_url as image, l.language as role").
		Joins(fmt.Sprintf("left join %s as p on p.id = ac.people_id", raw.People{}.TableName())).
		Joins(fmt.Sprintf("left join %s as l on l.id = ac.language_id", raw.Language{}.TableName())).
		Where("ac.character_id = ? and p.id != 0", id).
		Group("ac.people_id, p.name, p.image_url, ac.language_id, l.language").
		Order("ac.language_id asc, ac.people_id asc").
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
