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

// GetAnimeCharacter to get anime character and their voice actor.
func (d *Database) GetAnimeCharacter(id int, page int, limit int) ([]model.AnimeCharacter, map[string]interface{}, int, error) {
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
	subQuery := d.db.Table(fmt.Sprintf("%s as ac", raw.AnimeCharacter{}.TableName())).
		Select("c.id").
		Joins(fmt.Sprintf("left join %s as c on c.id = ac.character_id", raw.Character{}.TableName())).
		Where("ac.anime_id = ?", id).
		Order("ac.role asc, c.name asc").
		Group("c.id, ac.role").
		Limit(limit).
		Offset(limit * (page - 1))

	rows, err := d.db.Table("(?) as aa", subQuery).
		Select("ac.character_id as c_id, c.name as c_name, c.image_url as c_image, ac.role, ac.language_id, ac.people_id as p_id, p.name as p_name, p.image_url as p_image").
		Joins(fmt.Sprintf("left join %s as ac on aa.id = ac.character_id", raw.AnimeCharacter{}.TableName())).
		Joins(fmt.Sprintf("left join %s as c on c.id = ac.character_id", raw.Character{}.TableName())).
		Joins(fmt.Sprintf("left join %s as p on p.id = ac.people_id", raw.People{}.TableName())).
		Where("ac.anime_id = ?", id).
		Order("ac.role asc, c.name asc, ac.language_id asc").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	chars := []model.AnimeCharacter{}
	charMap := make(map[int]int)
	vaMap := make(map[int][]model.Role)
	for rows.Next() {
		var tmp join.AnimeCharacter
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		charMap[tmp.CID]++

		if tmp.PID != 0 {
			vaMap[tmp.CID] = append(vaMap[tmp.CID], model.Role{
				ID:    tmp.PID,
				Name:  tmp.PName,
				Image: tmp.PImage,
				Role:  constant.Languages[tmp.LanguageID],
			})
		}

		if charMap[tmp.CID] == 1 {
			chars = append(chars, model.AnimeCharacter{
				ID:          tmp.CID,
				Name:        tmp.CName,
				Image:       tmp.CImage,
				Role:        tmp.Role,
				VoiceActors: []model.Role{},
			})
		}
	}

	for i, c := range chars {
		if len(vaMap[c.ID]) > 0 {
			chars[i].VoiceActors = vaMap[c.ID]
		}
	}

	// Prepare meta.
	var count int64
	if err = subQuery.Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta := map[string]interface{}{
		"count": count,
	}

	return chars, meta, http.StatusOK, nil
}
