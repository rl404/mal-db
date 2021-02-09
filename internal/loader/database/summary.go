package database

import (
	"net/http"

	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
)

// GetEntryCount to get all entry count.
func (d *Database) GetEntryCount() (*model.Total, map[string]interface{}, int, error) {
	var animeCount, mangaCount, charCount, peopleCount int64
	if err := d.db.Model(&raw.Anime{}).Count(&animeCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.Manga{}).Count(&mangaCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.Character{}).Count(&charCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.People{}).Count(&peopleCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	return &model.Total{
		Anime:     int(animeCount),
		Manga:     int(mangaCount),
		Character: int(charCount),
		People:    int(peopleCount),
	}, nil, http.StatusOK, nil
}
