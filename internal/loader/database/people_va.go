package database

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/join"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetPeopleVA to get people voice actor role list.
func (d *Database) GetPeopleVA(id int, page int, limit int) ([]model.VoiceActor, map[string]interface{}, int, error) {
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
	rows, err := d.db.Table(fmt.Sprintf("%s as ac", raw.AnimeCharacter{}.TableName())).
		Select("ac.anime_id as a_id, a.title as a_title, a.image_url as a_image, ac.role, l.language, ac.character_id as c_id, c.name as c_name, c.image_url as c_image").
		Joins(fmt.Sprintf("left join %s as a on a.id = ac.anime_id", raw.Anime{}.TableName())).
		Joins(fmt.Sprintf("left join %s as c on c.id = ac.character_id", raw.Character{}.TableName())).
		Joins(fmt.Sprintf("left join %s as l on l.id = ac.language_id", raw.Language{}.TableName())).
		Where("ac.people_id = ?", id).
		Order("c.name asc, a.title asc").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	roles := []model.VoiceActor{}
	charMap := make(map[int]int)
	animeMap := make(map[int][]model.Role)
	for rows.Next() {
		var tmp join.PeopleVA
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		charMap[tmp.CID]++

		if charMap[tmp.CID] == 1 {
			roles = append(roles, model.VoiceActor{
				ID:    tmp.CID,
				Name:  tmp.CName,
				Image: tmp.CImage,
				Role:  tmp.Language,
				Anime: []model.Role{},
			})
		}

		animeMap[tmp.CID] = append(animeMap[tmp.CID], model.Role{
			ID:    tmp.AID,
			Name:  tmp.ATitle,
			Image: tmp.AImage,
			Role:  tmp.Role,
		})
	}

	for i, r := range roles {
		if len(animeMap[r.ID]) != 0 {
			roles[i].Anime = animeMap[r.ID]
		}
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"count": len(roles),
	}

	return roles, meta, http.StatusOK, nil
}
