package database

import (
	"net/http"

	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
)

// GetProducerMagazine to get all producer/magazine list.
func (d *Database) GetProducerMagazine(t string) (data []model.Item, meta map[string]interface{}, code int, err error) {
	if err := d.db.Model(raw.ProducerMagazine{}).Where("type = ?", t).Order("name asc").Find(&data).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta = map[string]interface{}{
		"count": len(data),
	}
	return data, meta, http.StatusOK, nil
}
