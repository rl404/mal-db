package database

import (
	"fmt"
	"net/http"

	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
)

// GetGenres to get all anime/manga genre list.
func (d *Database) GetGenres(t string) ([]model.Item, map[string]interface{}, int, error) {
	rows, err := d.db.Model(&raw.Genre{}).Where("type = ?", t).Order("genre asc").Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	genres := []model.Item{}
	for rows.Next() {
		var tmp raw.Genre
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		genres = append(genres, model.Item{
			ID:   tmp.ID,
			Name: tmp.Genre,
		})
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"count": len(genres),
	}

	return genres, meta, http.StatusOK, nil
}

func (d *Database) getMediaGenre(t string, id int) (genres []model.Item) {
	err := d.db.Table(fmt.Sprintf("%s as mg", raw.MediaGenre{}.TableName())).
		Select("g.id, g.genre as name").
		Joins(fmt.Sprintf("left join %s as g on g.id = mg.genre_id", raw.Genre{}.TableName())).
		Where("mg.media_id = ? and mg.type = ? and g.type = ?", id, t, t).
		Find(&genres).Error
	if err != nil {
		d.log.Error(err.Error())
	}
	return genres
}
