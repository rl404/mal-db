package article

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type tag struct {
	area *goquery.Selection
	data []model.ArticleTagItem
}

// GetTags to get featured article tag list.
func (p *parser) GetTags(a *goquery.Selection) []model.ArticleTagItem {
	v := tag{area: a}
	v.setDetails()
	return v.data
}

func (t *tag) setDetails() {
	tagList := []model.ArticleTagItem{}
	t.area.Find("div.tag-list").Each(func(i int, eachTag *goquery.Selection) {
		link, _ := eachTag.Find("a").Attr("href")
		tagList = append(tagList, model.ArticleTagItem{
			Name: eachTag.Find("span").Text(),
			Tag:  utils.GetValueFromSplit(link, "/", 5),
		})
	})
	t.data = tagList
}
