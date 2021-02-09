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

// GetPeopleStaff to get people anime staff role list.
func (d *Database) GetPeopleStaff(id int, page int, limit int) (roles []model.Role, meta map[string]interface{}, code int, err error) {
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
	err = d.db.Table(fmt.Sprintf("%s as st", raw.AnimeStaff{}.TableName())).
		Select("st.anime_id as id, a.title as name, a.image_url as image, p.position as role").
		Joins(fmt.Sprintf("left join %s as a on a.id = st.anime_id", raw.Anime{}.TableName())).
		Joins(fmt.Sprintf("left join %s as p on p.id = st.position_id", raw.Position{}.TableName())).
		Where("st.people_id = ?", id).
		Order("a.title asc, st.position_id asc").
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
