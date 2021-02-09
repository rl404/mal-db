package saver

import (
	"errors"
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

func (a *API) parseCharacter(id int) error {
	// Check old entry.
	if a.isEntryNew(constant.CharacterType, id) {
		return _errors.ErrNewData
	}

	// Delete existing empty id.
	if err := a.deleteEmptyID(constant.CharacterType, id); err != nil {
		return err
	}

	// Parse character.
	data, code, err := a.mal.GetCharacter(id)

	// Empty id.
	if code == http.StatusNotFound {
		if err = a.insertEmptyID(constant.CharacterType, id); err != nil {
			return err
		}
		a.logger.Debug("empty character %v", id)
		return nil
	}

	// Other error.
	if err != nil {
		return err
	}

	// Begin transaction.
	tx := a.db.Begin()
	defer tx.Rollback()

	// Get existing anime db.
	var character raw.Character
	err = tx.Where("id = ?", id).First(&character).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Update data.
	character.ID = data.ID
	character.Name = data.Name
	character.Nickname = data.Nickname
	character.JapaneseName = data.JapaneseName
	character.ImageURL = data.Image
	character.Favorite = data.Favorite
	character.About = data.About

	// Save character.
	if err = tx.Save(&character).Error; err != nil {
		return err
	}

	// Save history.
	if err = a.saveCharacterHistory(tx, character); err != nil {
		return err
	}

	// Enqueue anime.
	if err = a.enqueueCharacterOgraphy(constant.AnimeType, id); err != nil {
		return err
	}

	// Enqueue manga.
	if err = a.enqueueCharacterOgraphy(constant.MangaType, id); err != nil {
		return err
	}

	// Commit.
	if err = tx.Commit().Error; err != nil {
		return err
	}

	// Delete cache.
	return a.cacher.Delete(constant.GetKey(constant.KeyCharacter, id))
}

func (a *API) enqueueCharacterOgraphy(t string, id int) (err error) {
	var data []model.Role
	switch t {
	case constant.AnimeType:
		data, _, err = a.mal.GetCharacterAnime(id)
	case constant.MangaType:
		data, _, err = a.mal.GetCharacterManga(id)
	}
	if err != nil {
		return err
	}
	for _, v := range data {
		if err = a.enqueue(constant.AnimeType, v.ID); err != nil {
			return err
		}
	}
	return nil
}
