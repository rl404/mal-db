package news

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.News
}

// GetDetails to get news details.
func (p *parser) GetDetails(a *goquery.Selection) *model.News {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setID()
	d.setTitle()
	d.setContent()
	d.setUsername()
	d.setForumID()
	d.setComment()
	d.setDate()
	d.setTags()
	d.setRelated()
	return &d.data
}

func (d *detail) setID() {
	link, _ := d.area.Find(".breadcrumb").First().Find(".di-ib:last-child a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (d *detail) setTitle() {
	d.data.Title = strings.TrimSpace(d.area.Find(".title").Text())
}

func (d *detail) setContent() {
	content, _ := d.area.Find(".news-container .content").Html()
	d.data.Content = strings.TrimSpace(content)
}

func (d *detail) setDate() {
	dateArea := d.area.Find(".information").First()
	dateArea.Find("a").Remove()
	date := strings.Replace(dateArea.Text(), "by", "", -1)
	date = strings.Replace(date, "|", "", -1)
	t, _ := utils.StrToTime(strings.TrimSpace(date))
	d.data.Date = t
}

func (d *detail) setUsername() {
	user, _ := d.area.Find(".information").First().Find("a").First().Attr("href")
	d.data.Username = utils.GetValueFromSplit(user, "/", 4)
}

func (d *detail) setComment() {
	comment := d.area.Find(".information .comment").Text()
	comment = strings.Replace(comment, "Comments", "", -1)
	comment = strings.Replace(comment, "Comment", "", -1)
	d.data.Comment = utils.StrToNum(comment)
}

func (d *detail) setForumID() {
	id, _ := d.area.Find(".information .comment").Attr("href")
	d.data.ForumID = utils.StrToNum(utils.GetValueFromSplit(id, "topicid=", 1))
}

func (d *detail) setTags() {
	tags := []string{}
	d.area.Find(".tags").Find(".tag").Each(func(i int, tag *goquery.Selection) {
		link, _ := tag.Attr("href")
		tags = append(tags, utils.GetValueFromSplit(link, "/", 5))
	})
	d.data.Tags = tags
}

func (d *detail) setRelated() {
	d.area.Find("table.news-related-database").Find("tr").Each(func(i int, relArea *goquery.Selection) {
		relType := strings.TrimSpace(relArea.Find("td").First().Text())
		if relType == "Anime:" {
			d.data.Related.Anime = d.getRelatedDetail(relArea)
		} else if relType == "Manga:" {
			d.data.Related.Manga = d.getRelatedDetail(relArea)
		} else if relType == "People:" {
			d.data.Related.People = d.getRelatedDetail(relArea)
		}
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
