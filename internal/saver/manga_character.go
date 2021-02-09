package saver

import (
	"strings"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveMangaCharacter(tx *gorm.DB, data *model.Manga) error {
	// Parse manga character.
	chars, _, err := a.mal.GetMangaCharacter(data.ID)
	if err != nil {
		return err
	}

	if len(chars) == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("manga_id = ?", data.ID).Delete(raw.MangaCharacter{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	query := utils.BatchInsertQuery(raw.MangaCharacter{}.TableName(), len(chars), 3)

	for _, c := range chars {
		args = append(args, data.ID)
		args = append(args, c.ID)
		args = append(args, strings.ToLower(c.Role))

		// Enqueue related character.
		if err = a.enqueue(constant.CharacterType, c.ID); err != nil {
			return err
		}
	}

	return tx.Exec(query, args...).Error
}
