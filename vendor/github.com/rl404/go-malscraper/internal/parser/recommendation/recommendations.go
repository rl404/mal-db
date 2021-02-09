package recommendation

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type recommendations struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Recommendation
}

// GetRecommendations to get anime/manga recommendation list.
func (p *parser) GetRecommendations(a *goquery.Selection) []model.Recommendation {
	v := recommendations{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (r *recommendations) setDetail() {
	recommends := []model.Recommendation{}
	r.area.Find("div[class=\"spaceit borderClass\"]").Each(func(i int, eachRecom *goquery.Selection) {
		recommends = append(recommends, model.Recommendation{
			Source:      r.getSource(eachRecom),
			Recommended: r.getRecommended(eachRecom),
			Users:       r.getUser(eachRecom),
		})
	})
	r.data = recommends
}

func (r *recommendations) getUser(eachRecom *goquery.Selection) (recs []model.RecommendationUser) {
	return append(recs, model.RecommendationUser{
		Username: r.getUsername(eachRecom),
		Content:  r.getContent(eachRecom),
	})
}

func (r *recommendations) getSource(eachRecom *goquery.Selection) model.Source {
	area := eachRecom.Find("table tr td:nth-of-type(1)")
	return model.Source{
		ID:    r.getSourceID(area),
		Title: r.getSourceTitle(area),
		Type:  r.getSourceType(area),
		Image: r.getSourceImage(area),
	}
}

func (r *recommendations) getRecommended(eachRecom *goquery.Selection) model.Source {
	area := eachRecom.Find("table tr td:nth-of-type(2)")
	return model.Source{
		ID:    r.getSourceID(area),
		Title: r.getSourceTitle(area),
		Type:  r.getSourceType(area),
		Image: r.getSourceImage(area),
	}
}

func (r *recommendations) getSourceID(area *goquery.Selection) int {
	id, _ := area.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (r *recommendations) getSourceTitle(area *goquery.Selection) string {
	return area.Find("strong").First().Text()
}

func (r *recommendations) getSourceType(area *goquery.Selection) string {
	id, _ := area.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(id, "/", 3)
}

func (r *recommendations) getSourceImage(area *goquery.Selection) string {
	image, _ := area.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *recommendations) getUsername(eachRecom *goquery.Selection) string {
	eachRecom.Find("table").First().Next().Next().Find("a").First().Remove()
	return eachRecom.Find("table").First().Next().Next().Find("a").Text()
}

func (r *recommendations) getContent(eachRecom *goquery.Selection) string {
	return eachRecom.Find(".recommendations-user-recs-text").Text()
}
