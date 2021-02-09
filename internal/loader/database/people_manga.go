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

// GetPeopleManga to get people published manga list.
func (d *Database) GetPeopleManga(id int, page int, limit int) (roles []model.Role, meta map[string]interface{}, code int, err error) {
	// Is empty.
	if d.isEntryEmpty(constant.PeopleType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	if errors.Is(d.db.Where("id = ?", id).First(&raw.People{}).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(constant.PeopleType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Prepare query.
	err = d.db.Table(fmt.Sprintf("%s as pm", raw.PeopleManga{}.TableName())).
		Select("pm.manga_id as id, m.title as name, m.image_url as image, p.position as role").
		Joins(fmt.Sprintf("left join %s as m on m.id = pm.manga_id", raw.Manga{}.TableName())).
		Joins(fmt.Sprintf("left join %s as p on p.id = pm.position_id", raw.Position{}.TableName())).
		Where("pm.people_id = ?", id).
		Order("m.title asc, pm.position_id asc").
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
