package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetArticle to get featured article detail information.
//
// Example: https://myanimelist.net/featured/2321/Free_Manga_Service__Update__New_Anime_Titles.
func (m *Malscraper) GetArticle(id int) (*model.Article, int, error) {
	return m.api.GetArticle(id)
}

// GetArticles to get featured article list.
//
// Tag should be from `GetArticleTag()`.
//
// Example: https://myanimelist.net/featured.
func (m *Malscraper) GetArticles(pageTag ...interface{}) ([]model.ArticleItem, int, error) {
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
	return m.api.GetArticles(page, tag)
}

// GetArticleTag to get featured article tag list.
//
// Example: https://myanimelist.net/featured/tag.
func (m *Malscraper) GetArticleTag() ([]model.ArticleTagItem, int, error) {
	return m.api.GetArticleTag()
}
