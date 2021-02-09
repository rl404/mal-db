package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
)

// GetProducers to get anime producer/studio/licensor list.
func (v *Validator) GetProducers() ([]model.ItemCount, int, error) {
	return v.api.GetProducers()
}

// GetProducer to get producer anime list.
func (v *Validator) GetProducer(id int, page int) ([]model.AnimeItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isProducerValid(id) {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetProducer(id, page)
}

// GetMagazines to get manga magazine/serialization list.
func (v *Validator) GetMagazines() ([]model.ItemCount, int, error) {
	return v.api.GetMagazines()
}

// GetMagazine to get magazine manga list.
func (v *Validator) GetMagazine(id int, page int) ([]model.MangaItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isMagazineValid(id) {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetMagazine(id, page)
}
