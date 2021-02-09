package user

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type recommendation struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Recommendation
}

// GetRecommendations to get user recommendation list.
func (p *parser) GetRecommendations(a *goquery.Selection) []model.Recommendation {
	v := recommendation{area: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (r *recommendation) setDetails() {
	recommendations := []model.Recommendation{}
	username := r.getUsername()
	r.area.Find("div[class=\"spaceit borderClass\"]").EachWithBreak(func(i int, eachRecom *goquery.Selection) bool {
		if eachRecom.Find("table").Text() == "" {
			return true
		}
		recommendations = append(recommendations, model.Recommendation{
			Source:      r.getSource(eachRecom),
			Recommended: r.getSource(eachRecom),
			Users: []model.RecommendationUser{
				{
					Username: username,
					Content:  r.getContent(eachRecom),
				},
			},
		})
		return true
	})
	r.data = recommendations
}

func (r *recommendation) getUsername() string {
	user := strings.TrimSpace(r.area.Parent().Parent().Parent().Find("h1.h1 span").Text())
	return user[:len(user)-10]
}

func (r *recommendation) getSource(eachRecom *goquery.Selection) (s model.Source) {
	area := eachRecom.Find("table tr").Find("td").First()
	s = model.Source{
		ID:    r.getSourceID(area),
		Title: r.getSourceTitle(area),
		Type:  r.getSourceType(area),
		Image: r.getSourceImage(area),
	}
	area.Remove()
	return s
}

func (r *recommendation) getSourceID(area *goquery.Selection) int {
	id, _ := area.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (r *recommendation) getSourceTitle(area *goquery.Selection) string {
	return area.Find("strong").First().Text()
}

func (r *recommendation) getSourceType(area *goquery.Selection) string {
	t, _ := area.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(t, "/", 3)
}

func (r *recommendation) getSourceImage(area *goquery.Selection) string {
	image, _ := area.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *recommendation) getContent(eachRecom *goquery.Selection) string {
	return eachRecom.Find(".profile-user-recs-text").Text()
}
