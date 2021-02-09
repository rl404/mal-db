package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetNews to get news detail information.
//
// Example: https://myanimelist.net/news/34036779.
func (m *Malscraper) GetNews(id int) (*model.News, int, error) {
	return m.api.GetNews(id)
}

// GetNewsList to get news list.
//
// Tag should be from `GetNewsTag()`.
//
// Example: https://myanimelist.net/news.
func (m *Malscraper) GetNewsList(pageTag ...interface{}) ([]model.NewsItem, int, error) {
	page, tag := 1, ""
	for i, param := range pageTag {
		switch i {
		case 0:
			if v, ok := param.(int); ok {
				page = v
			}
		case 1:
			if v, ok := param.(string); ok {
				tag = v
			}
		}
	}
	return m.api.GetNewsList(page, tag)
}

// GetNewsTag to get news tag list.
//
// Example: https://myanimelist.net/news/tag.
func (m *Malscraper) GetNewsTag() (*model.NewsTag, int, error) {
	return m.api.GetNewsTag()
}
