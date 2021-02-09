package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetAnime to get anime details.
func (p *Parser) GetAnime(id int) (*model.Anime, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetDetails(doc), http.StatusOK, nil
}

// GetAnimeCharacter to get anime charater list.
func (p *Parser) GetAnimeCharacter(id int) ([]model.CharacterItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "characters"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetCharacters(doc), http.StatusOK, nil
}

// GetAnimeStaff to get anime staff list.
func (p *Parser) GetAnimeStaff(id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "characters"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetStaff(doc), http.StatusOK, nil
}

// GetAnimeVideo to get anime video list.
func (p *Parser) GetAnimeVideo(id int, page int) (*model.Video, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "anime", id, "a", "video"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetVideos(doc), http.StatusOK, nil
}

// GetAnimeEpisode to get anime episode list.
func (p *Parser) GetAnimeEpisode(id int, page int) ([]model.Episode, int, error) {
	q := map[string]interface{}{"offset": 100 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "anime", id, "a", "episode"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetEpisodes(doc), http.StatusOK, nil
}

// GetAnimeStats to get anime stats.
func (p *Parser) GetAnimeStats(id int) (*model.Stats, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "stats"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetStats(doc), http.StatusOK, nil
}

// GetAnimeReview to get anime review list.
func (p *Parser) GetAnimeReview(id int, page int) ([]model.Review, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "anime", id, "a", "reviews"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetReviews(doc), http.StatusOK, nil
}

// GetAnimeRecommendation to get anime recommendation list.
func (p *Parser) GetAnimeRecommendation(id int) ([]model.Recommendation, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "userrecs"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetRecommendations(doc), http.StatusOK, nil
}

// GetAnimeNews to get anime recommendation list.
func (p *Parser) GetAnimeNews(id int) ([]model.NewsItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "news"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetNews(doc), http.StatusOK, nil
}

// GetAnimeArticle to get anime featured article list.
func (p *Parser) GetAnimeArticle(id int) ([]model.ArticleItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "featured"), ".news-list")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetArticle(doc), http.StatusOK, nil
}

// GetAnimeClub to get anime club list.
func (p *Parser) GetAnimeClub(id int) ([]model.ClubItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "clubs"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetClubs(doc), http.StatusOK, nil
}

// GetAnimePicture to get anime picture list.
func (p *Parser) GetAnimePicture(id int) ([]string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "pics"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return nil, code, err
	}
	return p.anime.GetPictures(doc), http.StatusOK, nil
}

// GetAnimeMoreInfo to get anime more info.
func (p *Parser) GetAnimeMoreInfo(id int) (string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", id, "a", "moreinfo"), ".js-scrollfix-bottom-rel")
	if err != nil {
		return "", code, err
	}
	return p.anime.GetMoreInfo(doc), http.StatusOK, nil
}
