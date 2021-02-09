package parser

import (
	"net/http"
	"strconv"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetRecommendation to get recommendation details.
func (p *Parser) GetRecommendation(rType string, id1, id2 int) (*model.Recommendation, int, error) {
	id := strconv.Itoa(id1) + "-" + strconv.Itoa(id2)
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "recommendations", rType, id), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.recommendation.GetDetails(doc), http.StatusOK, nil
}

// GetRecommendations to get anime/manga recommendation list.
func (p *Parser) GetRecommendations(t string, page int) ([]model.Recommendation, int, error) {
	q := map[string]interface{}{"s": "recentrecs", "t": t, "show": 100 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "recommendations.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.recommendation.GetRecommendations(doc), http.StatusOK, nil
}
