package news

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type newsList struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.NewsItem
}

// GetList to get news list.
func (p *parser) GetList(a *goquery.Selection) []model.NewsItem {
	v := newsList{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (nl *newsList) setDetail() {
	list := []model.NewsItem{}
	nl.area.Find("div.news-list").Find(".news-unit").EachWithBreak(func(i int, eachNews *goquery.Selection) bool {
		if eachNews.Find(".title").Text() == "" {
			return true
		}
		list = append(list, model.NewsItem{
			ID:       nl.getID(eachNews),
			Title:    nl.getTitle(eachNews),
			Image:    nl.getImage(eachNews),
			Content:  nl.getContent(eachNews),
			Username: nl.getUsername(eachNews),
			ForumID:  nl.getForumID(eachNews),
			Comment:  nl.getComment(eachNews),
			Date:     nl.getDate(eachNews),
		})
		return true
	})
	nl.data = list
}

func (nl *newsList) getID(eachNews *goquery.Selection) int {
	id, _ := eachNews.Find(".title a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (nl *newsList) getTitle(eachNews *goquery.Selection) string {
	return eachNews.Find(".title a").Text()
}

func (nl *newsList) getImage(eachNews *goquery.Selection) string {
	image, _ := eachNews.Find("img").Attr("src")
	return utils.URLCleaner(image, "image", nl.cleanImg)
}

func (nl *newsList) getContent(eachNews *goquery.Selection) string {
	return strings.TrimSpace(eachNews.Find(".text").Text())
}

func (nl *newsList) getDate(eachNews *goquery.Selection) time.Time {
	dateArea := eachNews.Find(".information").Find("p").First()
	dateArea.Find("a").Remove()
	date := strings.Replace(dateArea.Text(), "by", "", -1)
	t, _ := utils.StrToTime(date)
	return t
}

func (nl *newsList) getUsername(eachNews *goquery.Selection) string {
	user, _ := eachNews.Find(".information").Find("p").First().Find("a").First().Attr("href")
	return utils.GetValueFromSplit(user, "/", 4)
}

func (nl *newsList) getForumID(eachNews *goquery.Selection) int {
	id, _ := eachNews.Find(".information .comment").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "topicid=", 1))
}

func (nl *newsList) getComment(eachNews *goquery.Selection) int {
	comment := eachNews.Find(".information .comment").Text()
	comment = strings.Replace(comment, "Comments", "", -1)
	comment = strings.Replace(comment, "Comment", "", -1)
	return utils.StrToNum(comment)
}
