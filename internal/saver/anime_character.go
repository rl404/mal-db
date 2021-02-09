package saver

import (
	"strings"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveAnimeCharacter(tx *gorm.DB, data *model.Anime) error {
	// Parse anime character.
	chars, _, err := a.mal.GetAnimeCharacter(data.ID)
	if err != nil {
		return err
	}

	if len(chars) == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("anime_id = ?", data.ID).Delete(raw.AnimeCharacter{}).Error; err != nil {
		return err
	}

	// Prepare raw insert batch query.
	var args []interface{}
	count := 0
	for _, c := range chars {
		// Enqueue related character.
		if err = a.enqueue(constant.CharacterType, c.ID); err != nil {
			return err
		}

		if len(c.VoiceActors) > 0 {
			for _, v := range c.VoiceActors {
				args = append(args, data.ID)
				args = append(args, c.ID)
				args = append(args, v.ID)
				args = append(args, strings.ToLower(c.Role))
				args = append(args, utils.GetMapKey(constant.Languages, v.Role))
				count++

				// Enqueue related people (voice actor).
				if err = a.enqueue(constant.PeopleType, v.ID); err != nil {
					return err
				}
			}
		} else {
			args = append(args, data.ID)
			args = append(args, c.ID)
			args = append(args, 0)
			args = append(args, strings.ToLower(c.Role))
			args = append(args, 0)
			count++
		}
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(raw.AnimeCharacter{}.TableName(), count, 5)

	return tx.Exec(query, args...).Error
}
