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

//
func (d *Database) GetPeople(id int) (*model.People, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(constant.PeopleType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	var peopleRaw raw.People
	if errors.Is(d.db.Where("id = ?", id).First(&peopleRaw).Error, gorm.ErrRecordNotFound) {
		if err := d.enqueue(constant.PeopleType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Fill data.
	people := &model.People{
		ID:               peopleRaw.ID,
		Name:             peopleRaw.Name,
		Image:            peopleRaw.ImageURL,
		GivenName:        peopleRaw.GivenName,
		FamilyName:       peopleRaw.FamilyName,
		AlternativeNames: []string{},
		Birthday: model.Date{
			Year:  peopleRaw.BirthdayYear,
			Month: peopleRaw.BirthdayMonth,
			Day:   peopleRaw.BirthdayDay,
		},
		Website:  peopleRaw.Website,
		Favorite: peopleRaw.Favorite,
		More:     peopleRaw.More,
	}

	altNames := utils.ArrayFilter(strings.Split(peopleRaw.AlternativeName, ", "))
	if len(altNames) != 0 {
		people.AlternativeNames = altNames
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"parsedAt": peopleRaw.UpdatedAt,
	}

	return people, meta, http.StatusOK, nil
}
