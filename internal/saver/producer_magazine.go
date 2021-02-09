package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveAnimeProducer(tx *gorm.DB, data *model.Anime) error {
	count := len(data.Producers) + len(data.Licensors) + len(data.Studios)
	if count == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("anime_id = ?", data.ID).Delete(raw.AnimeProducer{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.AnimeProducer{}.TableName(), count, 4)

	// Producer.
	for _, p := range data.Producers {
		args = append(args, data.ID)
		args = append(args, p.ID)
		args = append(args, false)
		args = append(args, false)
	}

	// Licensor.
	for _, p := range data.Licensors {
		args = append(args, data.ID)
		args = append(args, p.ID)
		args = append(args, true)
		args = append(args, false)
	}

	// Studio.
	for _, p := range data.Studios {
		args = append(args, data.ID)
		args = append(args, p.ID)
		args = append(args, false)
		args = append(args, true)
	}

	return tx.Exec(query, args...).Error
}

func (a *API) saveMangaMagazine(tx *gorm.DB, data *model.Manga) error {
	count := len(data.Serializations)
	if count == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("manga_id = ?", data.ID).Delete(raw.MangaMagazine{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.MangaMagazine{}.TableName(), count, 2)

	// Serialization.
	for _, s := range data.Serializations {
		args = append(args, data.ID)
		args = append(args, s.ID)
	}

	return tx.Exec(query, args...).Error
}
