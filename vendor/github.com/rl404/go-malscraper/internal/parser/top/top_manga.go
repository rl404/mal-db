package top

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type manga struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.TopManga
}

// GetManga to get top manga list.
func (p *parser) GetManga(a *goquery.Selection) []model.TopManga {
	v := manga{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (m *manga) setDetail() {
	topList := []model.TopManga{}
	m.area.Find("table").Find("tr.ranking-list").Each(func(j int, eachTop *goquery.Selection) {
		nameArea := eachTop.Find("td .detail")
		infoArea, _ := nameArea.Find("div.information").Html()
		parsedInfo := strings.Split(infoArea, "<br/>")
		topList = append(topList, model.TopManga{
			Rank:      m.getRank(eachTop),
			Image:     m.getImage(eachTop),
			ID:        m.getID(nameArea),
			Title:     m.getTitle(nameArea),
			Type:      m.getType(parsedInfo),
			Volume:    m.getEpCh(parsedInfo),
			StartDate: m.getDate(parsedInfo, 0),
			EndDate:   m.getDate(parsedInfo, 1),
			Member:    m.getMember(parsedInfo),
			Score:     m.getScore(eachTop),
		})
	})
	m.data = topList
}

func (m *manga) getRank(eachTop *goquery.Selection) int {
	return utils.StrToNum(eachTop.Find("td").First().Find("span").Text())
}

func (m *manga) getImage(eachTop *goquery.Selection) string {
	image, _ := eachTop.Find("td:nth-of-type(2) a img").Attr("data-src")
	return utils.URLCleaner(image, "image", m.cleanImg)
}

func (m *manga) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div").First().Attr("id")
	return utils.StrToNum(strings.Replace(id, "area", "", -1))
}

func (m *manga) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

func (m *manga) getType(parsedInfo []string) string {
	return strings.Split(strings.TrimSpace(parsedInfo[0]), " ")[0]
}

func (m *manga) getEpCh(parsedInfo []string) int {
	splitEpCh := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	return utils.StrToNum(splitEpCh[1][1:])
}

func (m *manga) getDate(parsedInfo []string, t int) model.Date {
	splitDate := strings.Split(strings.TrimSpace(parsedInfo[1]), "-")
	y1, m1, d1 := utils.StrToDate(splitDate[t])
	return model.Date{Year: y1, Month: m1, Day: d1}
}

func (m *manga) getMember(parsedInfo []string) int {
	member := strings.TrimSpace(parsedInfo[2])
	member = strings.Replace(member, "members", "", -1)
	member = strings.Replace(member, "favorites", "", -1)
	return utils.StrToNum(member)
}

func (m *manga) getScore(eachTop *goquery.Selection) float64 {
	return utils.StrToFloat(eachTop.Find("td:nth-of-type(3)").Text())
}
