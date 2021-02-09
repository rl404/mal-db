package manga

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type article struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ArticleItem
}

// GetArticle to get manga featured article list.
func (p *parser) GetArticle(a *goquery.Selection) []model.ArticleItem {
	v := article{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (a *article) setDetail() {
	articleList := []model.ArticleItem{}
	a.area.Find("div.news-unit").Each(func(i int, eachArticle *goquery.Selection) {
		articleList = append(articleList, model.ArticleItem{
			ID:            a.getID(eachArticle),
			Title:         a.getTitle(eachArticle),
			Image:         a.getImage(eachArticle),
			Summary:       a.getSummary(eachArticle),
			Username:      a.getUsername(eachArticle),
			View:          a.getView(eachArticle),
			IsSpoiler:     a.getSpoiler(eachArticle),
			IsAdvertorial: a.getAdvertorial(eachArticle),
			Tags:          a.getTags(eachArticle),
		})
	})
	a.data = articleList
}

func (a *article) getID(eachArticle *goquery.Selection) int {
	id, _ := eachArticle.Find(".title a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (a *article) getTitle(eachArticle *goquery.Selection) string {
	return eachArticle.Find(".title a").Text()
}

func (a *article) getImage(eachArticle *goquery.Selection) string {
	image, _ := eachArticle.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", a.cleanImg)
}

func (a *article) getSummary(eachArticle *goquery.Selection) string {
	return strings.TrimSpace(eachArticle.Find(".text").Text())
}

func (a *article) getUsername(eachArticle *goquery.Selection) string {
	user, _ := eachArticle.Find(".information").Find("a").First().Attr("href")
	return utils.GetValueFromSplit(user, "/", 4)
}

func (a *article) getView(eachArticle *goquery.Selection) int {
	return utils.StrToNum(eachArticle.Find(".information p b").Text())
}

func (a *article) getSpoiler(eachArticle *goquery.Selection) bool {
	return eachArticle.Find(".tag-color-feature-spoiler").Text() != ""
}

func (a *article) getAdvertorial(eachArticle *goquery.Selection) bool {
	return eachArticle.Find(".tag-color-feature-advertorial").Text() != ""
}

func (a *article) getTags(eachArticle *goquery.Selection) []string {
	tagList := []string{}
	eachArticle.Find(".information .tags").Find("a.tag").Each(func(i int, tag *goquery.Selection) {
		t, _ := tag.Attr("href")
		tagList = append(tagList, utils.GetValueFromSplit(t, "/", 5))
	})
	return tagList
}
