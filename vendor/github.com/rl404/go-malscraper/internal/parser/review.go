package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetReview to get review details.
func (p *Parser) GetReview(id int) (*model.Review, int, error) {
	q := map[string]interface{}{"id": id}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "reviews.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.review.GetDetails(doc), http.StatusOK, nil
}

// GetReviews to get anime/manga/best review list.
func (p *Parser) GetReviews(t string, page int) ([]model.Review, int, error) {
	q := map[string]interface{}{"p": page}
	if t == "bestvoted" {
		q["st"] = t
	} else {
		q["t"] = t
	}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "reviews.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.review.GetReviews(doc), http.StatusOK, nil
}
