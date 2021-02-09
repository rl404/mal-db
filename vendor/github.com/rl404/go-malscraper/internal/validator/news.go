package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetNews to get news detail information.
func (v *Validator) GetNews(id int) (*model.News, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyNews, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetNews(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetNewsList to get news list.
func (v *Validator) GetNewsList(page int, tag string) ([]model.NewsItem, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isNewsTagValid(tag) {
		return nil, http.StatusBadRequest, errors.ErrInvalidTag
	}
	return v.api.GetNewsList(page, tag)
}

// GetNewsTag to get news tag list.
func (v *Validator) GetNewsTag() (*model.NewsTag, int, error) {
	return v.api.GetNewsTag()
}
