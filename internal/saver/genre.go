package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveGenre(tx *gorm.DB, t string, id int, data []model.Item) error {
	if len(data) == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("type = ? and media_id = ?", t, id).Delete(raw.MediaGenre{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.MediaGenre{}.TableName(), len(data), 3)
	for _, g := range data {
		args = append(args, t)
		args = append(args, id)
		args = append(args, g.ID)
	}

	return tx.Exec(query, args...).Error
}
