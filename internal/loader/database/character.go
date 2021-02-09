package database

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rl404/go-malscraper/pkg/utils"
	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetCharacter to get character.
func (d *Database) GetCharacter(id int) (*model.Character, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(constant.CharacterType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	var charRaw raw.Character
	if errors.Is(d.db.Where("id = ?", id).First(&charRaw).Error, gorm.ErrRecordNotFound) {
		if err := d.enqueue(constant.CharacterType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Fill data.
	char := &model.Character{
		ID:           charRaw.ID,
		Name:         charRaw.Name,
		Nicknames:    []string{},
		JapaneseName: charRaw.JapaneseName,
		Image:        charRaw.ImageURL,
		Favorite:     charRaw.Favorite,
		About:        charRaw.About,
	}

	nickNames := utils.ArrayFilter(strings.Split(charRaw.Nickname, ", "))
	if len(nickNames) != 0 {
		char.Nicknames = nickNames
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"parsedAt": charRaw.UpdatedAt,
	}

	return char, meta, http.StatusOK, nil
}
