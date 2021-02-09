package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) enqueueMangaPeople(data *model.Manga) error {
	for _, p := range data.Authors {
		if err := a.enqueue(constant.PeopleType, p.ID); err != nil {
			return err
		}
	}
	return nil
}

func (a *API) savePeopleManga(tx *gorm.DB, id int) error {
	// Parse people manga.
	data, _, err := a.mal.GetPeopleManga(id)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("people_id = ?", id).Delete(raw.PeopleManga{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.PeopleManga{}.TableName(), len(data), 3)

	for _, m := range data {
		args = append(args, id)
		args = append(args, m.ID)
		args = append(args, utils.GetMapKey(constant.Positions, m.Role))

		// Enqueue related published manga.
		if err = a.enqueue(constant.MangaType, m.ID); err != nil {
			return err
		}
	}

	return tx.Exec(query, args...).Error
}
