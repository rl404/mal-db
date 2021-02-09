package database

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetAnimeStaff to get anime staff list.
func (d *Database) GetAnimeStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(constant.AnimeType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	if errors.Is(d.db.Where("id = ?", id).First(&raw.Anime{}).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(constant.AnimeType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Prepare query.
	rows, err := d.db.Table(fmt.Sprintf("%s as st", raw.AnimeStaff{}.TableName())).
		Select("st.people_id as id, p.name, p.image_url as image, p2.position as role").
		Joins(fmt.Sprintf("left join %s as p on p.id = st.people_id", raw.People{}.TableName())).
		Joins(fmt.Sprintf("left join %s as p2 on p2.id = st.position_id", raw.Position{}.TableName())).
		Where("st.anime_id = ?", id).
		Order("st.position_id asc, p.name asc").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	// Loop the result.
	staff := []model.Role{}
	staffMap, positionMap := make(map[int]int), make(map[int][]string)
	for rows.Next() {
		var tmp model.Role
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		if staffMap[tmp.ID] == 0 {
			staffMap[tmp.ID] = 1
			staff = append(staff, model.Role{
				ID:    tmp.ID,
				Name:  tmp.Name,
				Image: tmp.Image,
			})
		}

		positionMap[tmp.ID] = append(positionMap[tmp.ID], tmp.Role)
	}

	for i, s := range staff {
		staff[i].Role = strings.Join(positionMap[s.ID], ", ")
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"count": len(staff),
	}

	return staff, meta, http.StatusOK, nil
}
