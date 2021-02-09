package saver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

func (a *API) parsePeople(id int) error {
	// Check old entry.
	if a.isEntryNew(constant.PeopleType, id) {
		return _errors.ErrNewData
	}

	// Delete existing empty id.
	if err := a.deleteEmptyID(constant.PeopleType, id); err != nil {
		return err
	}

	// Parse people.
	data, code, err := a.mal.GetPeople(id)

	// Empty id.
	if code == http.StatusNotFound {
		if err = a.insertEmptyID(constant.PeopleType, id); err != nil {
			return err
		}
		a.logger.Debug("empty people %v", id)
		return nil
	}

	// Other error.
	if err != nil {
		return err
	}

	// Begin transaction.
	tx := a.db.Begin()
	defer tx.Rollback()

	// Get existing people db.
	var people raw.People
	err = tx.Where("id = ?", id).First(&people).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Update data.
	people.ID = data.ID
	people.Name = data.Name
	people.GivenName = data.GivenName
	people.FamilyName = data.FamilyName
	people.AlternativeName = strings.Join(data.AlternativeNames, ", ")
	people.ImageURL = data.Image
	people.BirthdayYear = data.Birthday.Year
	people.BirthdayMonth = data.Birthday.Month
	people.BirthdayDay = data.Birthday.Day
	people.Website = data.Website
	people.Favorite = data.Favorite
	people.More = data.More

	// Save people.
	if err = tx.Save(&people).Error; err != nil {
		return err
	}

	// Save history.
	if err = a.savePeopleHistory(tx, people); err != nil {
		return err
	}

	// Enqueue anime character.
	if err = a.enqueuePeopleCharacter(data.ID); err != nil {
		return err
	}

	// Enqueue anime staff.
	if err = a.enqueuePeopleStaff(data.ID); err != nil {
		return err
	}

	// Save manga.
	if err = a.savePeopleManga(tx, data.ID); err != nil {
		return err
	}

	// Commit.
	if err = tx.Commit().Error; err != nil {
		return err
	}

	// Delete cache.
	return a.cacher.Delete(constant.GetKey(constant.KeyPeople, id))
}

func (a *API) enqueuePeopleCharacter(id int) error {
	data, _, err := a.mal.GetPeopleCharacter(id)
	if err != nil {
		return err
	}
	for _, d := range data {
		if err = a.enqueue(constant.AnimeType, d.Anime.ID); err != nil {
			return err
		}
	}
	return nil
}

func (a *API) enqueuePeopleStaff(id int) error {
	data, _, err := a.mal.GetPeopleStaff(id)
	if err != nil {
		return err
	}
	for _, d := range data {
		if err = a.enqueue(constant.AnimeType, d.ID); err != nil {
			return err
		}
	}
	return nil
}
