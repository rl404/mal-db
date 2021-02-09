package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetManga to get manga details.
func (p *Parser) GetManga(id int) (*model.Manga, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetDetails(doc), http.StatusOK, nil
}

// GetMangaReview to get manga review list.
func (p *Parser) GetMangaReview(id int, page int) ([]model.Review, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "manga", id, "a", "reviews"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetReviews(doc), http.StatusOK, nil
}

// GetMangaRecommendation to get manga recommendation list.
func (p *Parser) GetMangaRecommendation(id int) ([]model.Recommendation, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "userrecs"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetRecommendations(doc), http.StatusOK, nil
}

// GetMangaStats to get manga stats list.
func (p *Parser) GetMangaStats(id int) (*model.Stats, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "stats"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetStats(doc), http.StatusOK, nil
}

// GetMangaCharacter to get manga character list.
func (p *Parser) GetMangaCharacter(id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "characters"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetCharacters(doc), http.StatusOK, nil
}

// GetMangaNews to get manga news list.
func (p *Parser) GetMangaNews(id int) ([]model.NewsItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "news"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetNews(doc), http.StatusOK, nil
}

// GetMangaArticle to get manga featured article list.
func (p *Parser) GetMangaArticle(id int) ([]model.ArticleItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "featured"), ".news-list")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetArticle(doc), http.StatusOK, nil
}

// GetMangaClub to get manga club list.
func (p *Parser) GetMangaClub(id int) ([]model.ClubItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "clubs"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetClubs(doc), http.StatusOK, nil
}

// GetMangaPicture to get manga picture list.
func (p *Parser) GetMangaPicture(id int) ([]string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "pics"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.manga.GetPictures(doc), http.StatusOK, nil
}

// GetMangaMoreInfo to get manga more info.
func (p *Parser) GetMangaMoreInfo(id int) (string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", id, "a", "moreinfo"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return "", code, err
	}
	return p.manga.GetMoreInfo(doc), http.StatusOK, nil
}
