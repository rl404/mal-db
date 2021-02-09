package recommendation

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.Recommendation
}

// GetDetails to get recommendation details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Recommendation {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setSource()
	d.setUser()
	return &d.data
}

func (d *detail) setSource() {
	sourceArea := d.area.Find(".borderDark table tr")
	likedArea := sourceArea.Find("td").First()
	d.data.Source = model.Source{
		ID:    d.getSourceID(likedArea),
		Title: d.getSourceTitle(likedArea),
		Type:  d.getSourceType(likedArea),
		Image: d.getSourceImage(likedArea),
	}

	likedArea.Remove()
	recommendationArea := sourceArea.Find("td").First()

	d.data.Recommended = model.Source{
		ID:    d.getSourceID(recommendationArea),
		Title: d.getSourceTitle(recommendationArea),
		Type:  d.getSourceType(recommendationArea),
		Image: d.getSourceImage(recommendationArea),
	}
}

func (d *detail) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (d *detail) getSourceTitle(sourceArea *goquery.Selection) string {
	return sourceArea.Find("strong").Text()
}

func (d *detail) getSourceType(sourceArea *goquery.Selection) string {
	link, _ := sourceArea.Find("a").Attr("href")
	return utils.GetValueFromSplit(link, "/", 3)
}

func (d *detail) getSourceImage(sourceArea *goquery.Selection) string {
	image, _ := sourceArea.Find("img").Attr("src")
	return utils.URLCleaner(image, "image", d.cleanImg)
}

func (d *detail) setUser() {
	contents := []model.RecommendationUser{}
	d.area.Find(".borderDark").Find(".borderClass").Each(func(i int, eachRecom *goquery.Selection) {
		contents = append(contents, model.RecommendationUser{
			Username: d.getRecomUser(eachRecom),
			Content:  d.getRecomText(eachRecom),
		})
	})
	d.data.Users = contents
}

func (d *detail) getRecomUser(eachRecom *goquery.Selection) string {
	return eachRecom.Find("a[href*=\"/profile/\"]").Text()
}

func (d *detail) getRecomText(eachRecom *goquery.Selection) string {
	eachRecom.Find("a").Remove()
	content := strings.TrimSpace(eachRecom.Find("div").First().Text())
	return regexp.MustCompile(`(\n\n\s*)`).ReplaceAllString(content, " ")
}
