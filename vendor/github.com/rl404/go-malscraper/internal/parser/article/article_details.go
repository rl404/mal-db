package article

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.Article
}

// GetDetails to get featured article details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Article {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setID()
	d.setTitle()
	d.setSummary()
	d.setContent()
	d.setUsername()
	d.setView()
	d.setDate()
	d.setSpoiler()
	d.setAdvertorial()
	d.setTags()
	d.setRelated()
	return &d.data
}

func (d *detail) setID() {
	link, _ := d.area.Find(".breadcrumb").First().Find(".di-ib:last-child a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (d *detail) setTitle() {
	d.data.Title = strings.TrimSpace(d.area.Find("h1.title").Text())
}

func (d *detail) setSummary() {
	d.data.Summary = d.area.Find("p.summary").Text()
}

func (d *detail) setContent() {
	content, _ := d.area.Find(".news-container .content").Html()
	d.data.Content = strings.TrimSpace(content)
}

func (d *detail) setDate() {
	dateArea := d.area.Find(".information").First()
	dateArea.Find("a").Remove()
	dateArea.Find("b").Remove()
	date := strings.Replace(dateArea.Text(), "by", "", -1)
	date = strings.Replace(date, "|", "", -1)
	date = strings.Replace(date, "views", "", -1)
	t, _ := utils.StrToTime(date)
	d.data.Date = t
}

func (d *detail) setUsername() {
	user, _ := d.area.Find(".information").First().Find("a").First().Attr("href")
	d.data.Username = utils.GetValueFromSplit(user, "/", 4)
}

func (d *detail) setView() {
	d.data.View = utils.StrToNum(d.area.Find(".news-info-block .information b").Text())
}

func (d *detail) setSpoiler() {
	d.data.IsSpoiler = d.area.Find(".news-container .tag-color-feature-spoiler").Text() != ""
}

func (d *detail) setAdvertorial() {
	d.data.IsAdvertorial = d.area.Find(".news-container .tag-color-feature-advertorial").Text() != ""
}

func (d *detail) setTags() {
	tags := []string{}
	d.area.Find(".news-container .tags").Find("a.tag").Each(func(i int, tag *goquery.Selection) {
		link, _ := tag.Attr("href")
		tags = append(tags, utils.GetValueFromSplit(link, "/", 5))
	})
	d.data.Tags = tags
}

func (d *detail) setRelated() {
	d.area.Find(".mr8 h2").EachWithBreak(func(i int, rel *goquery.Selection) bool {
		if goquery.NodeName(rel.Next()) != "table" {
			return true
		}
		rel.Next().Find("tr").Each(func(i int, relArea *goquery.Selection) {
			relType := strings.TrimSpace(relArea.Find("td").First().Text())
			if relType == "Anime:" {
				d.data.Related.Anime = d.getRelatedDetail(relArea)
			} else if relType == "Manga:" {
				d.data.Related.Manga = d.getRelatedDetail(relArea)
			} else if relType == "People:" {
				d.data.Related.People = d.getRelatedDetail(relArea)
			} else if relType == "Characters:" {
				d.data.Related.Character = d.getRelatedDetail(relArea)
			}
		})
		return true
	})
}

func (d *detail) getRelatedDetail(relArea *goquery.Selection) []model.Item {
	data := []model.Item{}
	relArea.Find("a").Each(func(i int, d *goquery.Selection) {
		id, _ := d.Attr("href")
		data = append(data, model.Item{
			ID:   utils.StrToNum(utils.GetValueFromSplit(id, "/", 4)),
			Name: d.Text(),
		})
	})
	return data
}
