package top

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type anime struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.TopAnime
}

// GetAnime to get top anime list.
func (p *parser) GetAnime(a *goquery.Selection) []model.TopAnime {
	v := anime{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (a *anime) setDetail() {
	topList := []model.TopAnime{}
	a.area.Find("table").Find("tr.ranking-list").Each(func(j int, eachTop *goquery.Selection) {
		nameArea := eachTop.Find("td .detail")
		infoArea, _ := nameArea.Find("div.information").Html()
		parsedInfo := strings.Split(infoArea, "<br/>")
		topList = append(topList, model.TopAnime{
			Rank:      a.getRank(eachTop),
			Image:     a.getImage(eachTop),
			ID:        a.getID(nameArea),
			Title:     a.getTitle(nameArea),
			Type:      a.getType(parsedInfo),
			Episode:   a.getEpCh(parsedInfo),
			StartDate: a.getDate(parsedInfo, 0),
			EndDate:   a.getDate(parsedInfo, 1),
			Member:    a.getMember(parsedInfo),
			Score:     a.getScore(eachTop),
		})
	})
	a.data = topList
}

func (a *anime) getRank(eachTop *goquery.Selection) int {
	return utils.StrToNum(eachTop.Find("td").First().Find("span").Text())
}

func (a *anime) getImage(eachTop *goquery.Selection) string {
	image, _ := eachTop.Find("td:nth-of-type(2) a img").Attr("data-src")
	return utils.URLCleaner(image, "image", a.cleanImg)
}

func (a *anime) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div").First().Attr("id")
	return utils.StrToNum(strings.Replace(id, "area", "", -1))
}

func (a *anime) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

func (a *anime) getType(parsedInfo []string) string {
	return strings.Split(strings.TrimSpace(parsedInfo[0]), " ")[0]
}

func (a *anime) getEpCh(parsedInfo []string) int {
	splitEpCh := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	return utils.StrToNum(splitEpCh[1][1:])
}

func (a *anime) getDate(parsedInfo []string, t int) model.Date {
	splitDate := strings.Split(strings.TrimSpace(parsedInfo[1]), "-")
	y, m, d := utils.StrToDate(splitDate[t])
	return model.Date{Year: y, Month: m, Day: d}
}

func (a *anime) getMember(parsedInfo []string) int {
	member := strings.TrimSpace(parsedInfo[2])
	member = strings.Replace(member, "members", "", -1)
	member = strings.Replace(member, "favorites", "", -1)
	return utils.StrToNum(member)
}

func (a *anime) getScore(eachTop *goquery.Selection) float64 {
	return utils.StrToFloat(eachTop.Find("td:nth-of-type(3)").Text())
}
