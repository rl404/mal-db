package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveSong(tx *gorm.DB, data *model.Anime) error {
	count := len(data.Song.Opening) + len(data.Song.Ending)
	if count == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("anime_id = ?", data.ID).Delete(raw.Song{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.Song{}.TableName(), count, 3, []string{"anime_id", "type", "song"})

	// Opening.
	for _, s := range data.Song.Opening {
		args = append(args, data.ID)
		args = append(args, constant.OpeningSong)
		args = append(args, s)
	}

	// Ending.
	for _, s := range data.Song.Ending {
		args = append(args, data.ID)
		args = append(args, constant.EndingSong)
		args = append(args, s)
	}

	return tx.Exec(query, args...).Error
}
