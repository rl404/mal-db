package search

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type manga struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.MangaSearch
}

// GetManga to search manga.
func (p *parser) GetManga(a *goquery.Selection) []model.MangaSearch {
	v := manga{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (m *manga) setDetail() {
	searchList := []model.MangaSearch{}
	m.area.Find("table").Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if eachSearch.Find(".picSurround").Text() == "" {
			return true
		}
		nameArea := eachSearch.Find("td:nth-of-type(2)")
		searchList = append(searchList, model.MangaSearch{
			Image:     m.getImage(eachSearch),
			ID:        m.getID(nameArea),
			Title:     m.getTitle(nameArea),
			Summary:   m.getSummary(nameArea),
			Type:      m.getType(eachSearch),
			Volume:    m.getVolume(eachSearch),
			Chapter:   m.getChapter(eachSearch),
			Score:     m.getScore(eachSearch),
			StartDate: m.getStartDate(eachSearch),
			EndDate:   m.getEndDate(eachSearch),
			Member:    m.getMember(eachSearch),
		})
		return true
	})
	m.data = searchList
}

func (m *manga) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

func (m *manga) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div[id^=sarea]").Attr("id")
	return utils.StrToNum(strings.Replace(id, "sarea", "", -1))
}

func (m *manga) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

func (m *manga) getSummary(nameArea *goquery.Selection) string {
	return strings.Replace(nameArea.Find(".pt4").Text(), "read more.", "", -1)
}

func (m *manga) getType(eachSearch *goquery.Selection) string {
	return strings.TrimSpace(strings.Replace(eachSearch.Find("td:nth-of-type(3)").Text(), "Unknown", "", -1))
}

func (m *manga) getVolume(eachSearch *goquery.Selection) int {
	return utils.StrToNum(eachSearch.Find("td:nth-of-type(4)").Text())
}

func (m *manga) getChapter(eachSearch *goquery.Selection) int {
	return utils.StrToNum(eachSearch.Find("td:nth-of-type(5)").Text())
}

func (m *manga) getScore(eachSearch *goquery.Selection) float64 {
	return utils.StrToFloat(eachSearch.Find("td:nth-of-type(6)").Text())
}

func (m *manga) getStartDate(eachSearch *goquery.Selection) model.Date {
	y1, m1, d1 := utils.StrToDate(eachSearch.Find("td:nth-of-type(7)").Text())
	return model.Date{Year: y1, Month: m1, Day: d1}
}

func (m *manga) getEndDate(eachSearch *goquery.Selection) model.Date {
	y1, m1, d1 := utils.StrToDate(eachSearch.Find("td:nth-of-type(8)").Text())
	return model.Date{Year: y1, Month: m1, Day: d1}
}

func (m *manga) getMember(eachSearch *goquery.Selection) int {
	return utils.StrToNum(eachSearch.Find("td:nth-of-type(9)").Text())
}
