package news

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type tag struct {
	area *goquery.Selection
	data model.NewsTag
}

// GetTags to get news tag list.
func (p *parser) GetTags(a *goquery.Selection) *model.NewsTag {
	v := tag{area: a}
	v.setAnime()
	v.setManga()
	v.setPeople()
	v.setMusic()
	v.setEvent()
	v.setIndustry()
	return &v.data
}

func (t *tag) setAnime() {
	t.data.Anime = t.setTags("#cat-anime")
}

func (t *tag) setManga() {
	t.data.Manga = t.setTags("#cat-manga")
}

func (t *tag) setPeople() {
	t.data.People = t.setTags("#cat-people")
}

func (t *tag) setMusic() {
	t.data.Music = t.setTags("#cat-music")
}

func (t *tag) setEvent() {
	t.data.Event = t.setTags("#cat-events")
}

func (t *tag) setIndustry() {
	t.data.Industry = t.setTags("#cat-industry")
}

func (t *tag) setTags(category string) []model.NewsTagItem {
	tagList := []model.NewsTagItem{}
	t.area.Find(category).Find("table").Find("tr").Each(func(i int, eachTag *goquery.Selection) {
		tagList = append(tagList, model.NewsTagItem{
			Name:        t.getName(eachTag),
			Tag:         t.getTag(eachTag),
			Description: t.getDescription(eachTag),
		})
	})
	return tagList
}

func (t *tag) getName(eachTag *goquery.Selection) string {
	return eachTag.Find(".tag-name a").Text()
}

func (t *tag) getTag(eachTag *goquery.Selection) string {
	link, _ := eachTag.Find(".tag-name a").Attr("href")
	return utils.GetValueFromSplit(link, "/", 5)
}

func (t *tag) getDescription(eachTag *goquery.Selection) string {
	return eachTag.Find(".tag-description span").Text()
}
