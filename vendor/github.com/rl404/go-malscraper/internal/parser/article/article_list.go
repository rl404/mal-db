package article

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type articles struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ArticleItem
}

// GetList to get featured article list.
func (p *parser) GetList(a *goquery.Selection) []model.ArticleItem {
	v := articles{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (a *articles) setDetail() {
	list := []model.ArticleItem{}
	a.area.Find("div.featured-pickup-unit").Each(func(i int, eachF *goquery.Selection) {
		list = append(list, model.ArticleItem{
			ID:            a.getIDPinned(eachF),
			Title:         a.getTitlePinned(eachF),
			Image:         a.getImagePinned(eachF),
			Summary:       a.getSummary(eachF),
			Username:      a.getUsername(eachF),
			View:          a.getView(eachF),
			IsSpoiler:     a.getSpoiler(eachF),
			IsAdvertorial: a.getAdvertorial(eachF),
			Tags:          a.getTags(eachF),
		})
	})
	a.area.Find(".news-unit").Each(func(i int, eachF *goquery.Selection) {
		list = append(list, model.ArticleItem{
			ID:            a.getID(eachF),
			Title:         a.getTitle(eachF),
			Image:         a.getImage(eachF),
			Summary:       a.getSummary(eachF),
			Username:      a.getUsername(eachF),
			View:          a.getView(eachF),
			IsSpoiler:     a.getSpoiler(eachF),
			IsAdvertorial: a.getAdvertorial(eachF),
			Tags:          a.getTags(eachF),
		})
	})
	a.data = list
}

func (a *articles) getIDPinned(eachF *goquery.Selection) int {
	id, _ := eachF.Find("a.title").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (a *articles) getTitlePinned(eachF *goquery.Selection) string {
	return eachF.Find("a.title").Text()
}

func (a *articles) getImagePinned(eachF *goquery.Selection) string {
	image, _ := eachF.Find("a.image").Attr("data-bg")
	return utils.URLCleaner(image, "image", a.cleanImg)
}

func (a *articles) getID(eachF *goquery.Selection) int {
	id, _ := eachF.Find(".title a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (a *articles) getTitle(eachF *goquery.Selection) string {
	return strings.TrimSpace(eachF.Find(".title a").Text())
}

func (a *articles) getImage(eachF *goquery.Selection) string {
	image, _ := eachF.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", a.cleanImg)
}

func (a *articles) getSummary(eachF *goquery.Selection) string {
	return strings.TrimSpace(eachF.Find(".text").Text())
}

func (a *articles) getUsername(eachF *goquery.Selection) string {
	user, _ := eachF.Find(".information").Find("p").First().Find("a").First().Attr("href")
	return utils.GetValueFromSplit(user, "/", 4)
}

func (a *articles) getView(eachF *goquery.Selection) int {
	return utils.StrToNum(eachF.Find(".information p b").Text())
}

func (a *articles) getSpoiler(eachF *goquery.Selection) bool {
	return eachF.Find(".tag-color-feature-spoiler").Text() != ""
}

func (a *articles) getAdvertorial(eachF *goquery.Selection) bool {
	return eachF.Find(".tag-color-feature-advertorial").Text() != ""
}

func (a *articles) getTags(eachF *goquery.Selection) []string {
	tagList := []string{}
	eachF.Find(".information .tags").Find("a.tag").Each(func(i int, tag *goquery.Selection) {
		t, _ := tag.Attr("href")
		tagList = append(tagList, utils.GetValueFromSplit(t, "/", 5))
	})
	return tagList
}
