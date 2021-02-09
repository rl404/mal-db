package saver

import (
	"strings"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveAnimeStaff(tx *gorm.DB, data *model.Anime) error {
	// Parse anime staff.
	staff, _, err := a.mal.GetAnimeStaff(data.ID)
	if err != nil {
		return err
	}

	if len(staff) == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("anime_id = ?", data.ID).Delete(raw.AnimeStaff{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	count := 0
	for _, p := range staff {
		// Enqueue related people (staff).
		if err = a.enqueue(constant.PeopleType, p.ID); err != nil {
			return err
		}

		positions := strings.Split(p.Role, ",")
		for _, v := range positions {
			args = append(args, data.ID)
			args = append(args, p.ID)
			args = append(args, utils.GetMapKey(constant.Positions, strings.TrimSpace(v)))
			count++
		}
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(raw.AnimeStaff{}.TableName(), count, 3)

	return tx.Exec(query, args...).Error
}
