package manga

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type news struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.NewsItem
}

// GetNews to get manga news list.
func (p *parser) GetNews(a *goquery.Selection) []model.NewsItem {
	v := news{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (n *news) setDetail() {
	newsList := []model.NewsItem{}
	n.area.Find("div.clearfix").EachWithBreak(func(i int, eachNews *goquery.Selection) bool {
		style, _ := eachNews.Attr("style")
		if style != "" {
			return true
		}
		newsList = append(newsList, model.NewsItem{
			ID:       n.getID(eachNews),
			Title:    n.getTitle(eachNews),
			Image:    n.getImage(eachNews),
			Content:  n.getContent(eachNews),
			Username: n.getUsername(eachNews),
			ForumID:  n.getForumID(eachNews),
			Comment:  n.getComment(eachNews),
			Date:     n.getDate(eachNews),
		})
		return true
	})
	n.data = newsList
}

func (n *news) getID(eachNews *goquery.Selection) int {
	id, _ := eachNews.Find(".spaceit a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 2))
}

func (n *news) getTitle(eachNews *goquery.Selection) string {
	return eachNews.Find(".spaceit strong").Text()
}

func (n *news) getImage(eachNews *goquery.Selection) string {
	image, _ := eachNews.Find(".picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", n.cleanImg)
}

func (n *news) getContent(eachNews *goquery.Selection) string {
	contentArea := eachNews.Find(".clearfix").Find("p")
	contentArea.Find("a").Remove()
	return contentArea.Text()
}

func (n *news) getDate(eachNews *goquery.Selection) time.Time {
	date := eachNews.Find("p.lightLink")
	date.Find("a").Remove()
	dateStr := strings.Replace(date.Text(), "by", "", -1)
	dateStr = strings.Replace(dateStr, "|", "", -1)
	t, _ := utils.StrToTime(dateStr)
	return t
}

func (n *news) getUsername(eachNews *goquery.Selection) string {
	user, _ := eachNews.Find("p.lightLink").Find("a").First().Attr("href")
	return utils.GetValueFromSplit(user, "/", 2)
}

func (n *news) getForumID(eachNews *goquery.Selection) int {
	id, _ := eachNews.Find("p.lightLink").Find("a:nth-of-type(2)").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "topicid=", 1))
}

func (n *news) getComment(eachNews *goquery.Selection) int {
	comment := eachNews.Find("p.lightLink").Find("a:nth-of-type(2)").Text()
	comment = strings.Replace(comment, "Discuss (", "", -1)
	comment = strings.Replace(comment, "comments)", "", -1)
	comment = strings.Replace(comment, "comment)", "", -1)
	return utils.StrToNum(comment)
}
