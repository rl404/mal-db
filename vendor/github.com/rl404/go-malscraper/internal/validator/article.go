package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetArticle to get featured article detail information.
func (v *Validator) GetArticle(id int) (*model.Article, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyArticle, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetArticle(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetArticles to get featured article list.
func (v *Validator) GetArticles(page int, tag string) ([]model.ArticleItem, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isArticleTagValid(tag) {
		return nil, http.StatusBadRequest, errors.ErrInvalidTag
	}
	return v.api.GetArticles(page, tag)
}

// GetArticleTag to get featured article tag list.
func (v *Validator) GetArticleTag() ([]model.ArticleTagItem, int, error) {
	return v.api.GetArticleTag()
}
