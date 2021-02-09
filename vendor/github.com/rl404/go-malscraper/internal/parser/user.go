package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetUser to get user details.
func (p *Parser) GetUser(user string) (*model.User, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "profile", user), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetDetails(doc), http.StatusOK, nil
}

// GetUserStats to get user stats details.
func (p *Parser) GetUserStats(user string) (*model.UserStats, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "profile", user), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetStats(doc), http.StatusOK, nil
}

// GetUserFavorite to get user favorite list.
func (p *Parser) GetUserFavorite(user string) (*model.UserFavorite, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "profile", user), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetFavorites(doc), http.StatusOK, nil
}

// GetUserFriend to get user friend list.
func (p *Parser) GetUserFriend(user string, page int) ([]model.UserFriend, int, error) {
	q := map[string]interface{}{"offset": 100 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "profile", user, "friends"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetFriends(doc), http.StatusOK, nil
}

// GetUserHistory to get user history list.
func (p *Parser) GetUserHistory(user string, t string) ([]model.UserHistory, int, error) {
	dir := []interface{}{"history", user}
	if t != "" {
		dir = append(dir, t)
	}
	doc, code, err := p.getDoc(utils.BuildURL(malURL, dir...), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetHistory(doc), http.StatusOK, nil
}

// GetUserReview to get user review list.
func (p *Parser) GetUserReview(user string, page int) ([]model.Review, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "profile", user, "reviews"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetReviews(doc), http.StatusOK, nil
}

// GetUserRecommendation to get user recommendation list.
func (p *Parser) GetUserRecommendation(user string, page int) ([]model.Recommendation, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "profile", user, "recommendations"), ".container-right")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetRecommendations(doc), http.StatusOK, nil
}

// GetUserClub to get user club list.
func (p *Parser) GetUserClub(user string) ([]model.Item, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "profile", user, "clubs"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.user.GetClubs(doc), http.StatusOK, nil
}

// Testable function.
var readBody func(io.Reader) ([]byte, error) = ioutil.ReadAll
var decodeJSON func([]byte, interface{}) error = json.Unmarshal

// GetUserAnime to get user anime list.
func (p *Parser) GetUserAnime(query model.UserListQuery) ([]model.UserAnime, int, error) {
	data := []model.UserAnime{}

	q := map[string]interface{}{
		"status": query.Status,
		"order":  query.Order,
		"tag":    query.Tag,
	}
	offset := 300 * (query.Page - 1)
	if query.Page == -1 {
		offset = 0
	}

	for {
		// Get body response.
		q["offset"] = offset
		body, code, err := p.getBody(utils.BuildURLWithQuery(q, malURL, "animelist", query.Username, "load.json"))
		if err != nil {
			return nil, code, err
		}

		// Read body.
		resp, err := readBody(body)
		if err != nil {
			p.logger.Error("failed reading body: %s", err.Error())
			return nil, http.StatusInternalServerError, errors.ErrParseBody
		}

		// Decode to arrays.
		var raw []model.UserRawAnime
		if err = decodeJSON(resp, &raw); err != nil {
			p.logger.Error("failed decode JSON: %s", err.Error())
			return nil, http.StatusInternalServerError, errors.ErrDecodeJSON
		}

		tmp := p.user.GetAnime(raw)
		data = append(data, tmp...)

		if len(tmp) == 0 || len(tmp) < 300 || query.Page != -1 {
			break
		}

		offset += 300
	}

	return data, http.StatusOK, nil
}

// GetUserManga to get user manga list.
func (p *Parser) GetUserManga(query model.UserListQuery) ([]model.UserManga, int, error) {
	data := []model.UserManga{}

	q := map[string]interface{}{
		"status": query.Status,
		"order":  query.Order,
		"tag":    query.Tag,
	}
	offset := 300 * (query.Page - 1)
	if query.Page == -1 {
		offset = 0
	}

	for {
		// Get body response.
		q["offset"] = offset
		body, code, err := p.getBody(utils.BuildURLWithQuery(q, malURL, "mangalist", query.Username, "load.json"))
		if err != nil {
			return nil, code, err
		}

		// Read body.
		resp, err := readBody(body)
		if err != nil {
			p.logger.Error("failed reading body: %s", err.Error())
			return nil, http.StatusInternalServerError, errors.ErrParseBody
		}

		// Decode to arrays.
		var raw []model.UserRawManga
		if err = decodeJSON(resp, &raw); err != nil {
			p.logger.Error("failed decode JSON: %s", err.Error())
			return nil, http.StatusInternalServerError, errors.ErrDecodeJSON
		}

		tmp := p.user.GetManga(raw)
		data = append(data, tmp...)

		if len(tmp) == 0 || len(tmp) < 300 || query.Page != -1 {
			break
		}

		offset += 300
	}

	return data, http.StatusOK, nil
}
