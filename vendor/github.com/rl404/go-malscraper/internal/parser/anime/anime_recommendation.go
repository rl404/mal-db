package anime

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

// GetRecommendations to get anime recommendation list.
func (p *parser) GetRecommendations(a *goquery.Selection) []model.Recommendation {
	v := recommendation{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (r *recommendation) setDetail() {
	recommendations := []model.Recommendation{}

	source := model.Source{
		ID:    r.getSourceID(),
		Title: r.getSourceTitle(),
		Image: r.getSourceImage(),
		Type:  "anime",
	}

	r.area.Find("div.borderClass table").Each(func(i int, eachRecommendation *goquery.Selection) {
		contentArea := eachRecommendation.Find("td:nth-of-type(2)")
		recommendations = append(recommendations, model.Recommendation{
			Source: source,
			Recommended: model.Source{
				ID:    r.getID(contentArea),
				Title: r.getTitle(contentArea),
				Image: r.getImage(eachRecommendation),
				Type:  "anime",
			},
			Users: r.getUsers(eachRecommendation),
		})
	})
	r.data = recommendations
}

func (r *recommendation) getSourceID() int {
	area := r.area.Find("#horiznav_nav").Find("li").First()
	link, _ := area.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (r *recommendation) getSourceImage() string {
	img, _ := r.area.Find("img.ac").Attr("data-src")
	return utils.URLCleaner(img, "img", r.cleanImg)
}

func (r *recommendation) getSourceTitle() string {
	titleArea := r.area.Parent().Find("h1.title-name").First()
	titleArea.Find("span.title-english").Remove()
	return titleArea.Text()
}

func (r *recommendation) getID(contentArea *goquery.Selection) int {
	id, _ := contentArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (r *recommendation) getTitle(contentArea *goquery.Selection) string {
	return contentArea.Find("strong").First().Text()
}

func (r *recommendation) getImage(eachRecommendation *goquery.Selection) string {
	image, _ := eachRecommendation.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *recommendation) getUsers(eachRecommendation *goquery.Selection) []model.RecommendationUser {
	contents := []model.RecommendationUser{}

	contentArea := eachRecommendation.Find("td:nth-of-type(2)")
	contents = append(contents, model.RecommendationUser{
		Username: r.getUsername(contentArea),
		Content:  r.getContent(contentArea),
	})

	otherArea := eachRecommendation.Find("div[id^=simaid]")
	otherArea.Find(".borderClass").Each(func(i int, eachOther *goquery.Selection) {
		contents = append(contents, model.RecommendationUser{
			Username: r.getUsername(eachOther),
			Content:  r.getContent(eachOther),
		})
	})

	return contents
}

func (r *recommendation) getUsername(contentArea *goquery.Selection) string {
	userArea := contentArea.Find(".borderClass .spaceit_pad:nth-of-type(2)").First()
	userArea.Find("a").First().Remove()
	return userArea.Find("a").Text()
}

func (r *recommendation) getContent(contentArea *goquery.Selection) string {
	content := contentArea.Find(".borderClass .spaceit_pad").First()
	content.Find("a").Remove()
	return strings.TrimSpace(content.Text())
}
