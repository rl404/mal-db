package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) saveRelated(tx *gorm.DB, t string, id int, data model.Related) error {
	var args []interface{}

	count := len(data.Sequel)
	for _, r := range data.Sequel {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 1)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.Prequel)
	for _, r := range data.Prequel {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 2)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.AltSetting)
	for _, r := range data.AltSetting {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 3)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.AltVersion)
	for _, r := range data.AltVersion {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 4)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.SideStory)
	for _, r := range data.SideStory {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 6)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.Summary)
	for _, r := range data.Summary {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 7)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.FullStory)
	for _, r := range data.FullStory {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 8)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.ParentStory)
	for _, r := range data.ParentStory {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 9)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.SpinOff)
	for _, r := range data.SpinOff {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 10)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.Adaptation)
	for _, r := range data.Adaptation {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 11)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.Character)
	for _, r := range data.Character {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 12)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	count += len(data.Other)
	for _, r := range data.Other {
		args = append(args, id)
		args = append(args, t)
		args = append(args, 13)
		args = append(args, r.ID)
		args = append(args, r.Type)

		if err := a.enqueue(r.Type, r.ID); err != nil {
			return err
		}
	}

	if count == 0 {
		return nil
	}

	// Delete existing data.
	if err := tx.Where("media_id = ? and media_type = ?", id, t).Delete(raw.MediaRelated{}).Error; err != nil {
		return err
	}

	// Create new relation.
	query := utils.BatchInsertQuery(raw.MediaRelated{}.TableName(), count, 5)

	return tx.Exec(query, args...).Error
}
