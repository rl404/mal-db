package search

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type anime struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.AnimeSearch
}

// GetAnime to search anime.
func (p *parser) GetAnime(a *goquery.Selection) []model.AnimeSearch {
	v := anime{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (a *anime) setDetail() {
	searchList := []model.AnimeSearch{}
	a.area.Find("table").Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if eachSearch.Find(".picSurround").Text() == "" {
			return true
		}
		nameArea := eachSearch.Find("td:nth-of-type(2)")
		searchList = append(searchList, model.AnimeSearch{
			Image:     a.getImage(eachSearch),
			ID:        a.getID(nameArea),
			Title:     a.getTitle(nameArea),
			Summary:   a.getSummary(nameArea),
			Type:      a.getType(eachSearch),
			Episode:   a.getProgress(eachSearch),
			Score:     a.getScore(eachSearch),
			StartDate: a.getStartDate(eachSearch),
			EndDate:   a.getEndDate(eachSearch),
			Member:    a.getMember(eachSearch),
			Rated:     a.getRated(eachSearch),
		})
		return true
	})
	a.data = searchList
}

func (a *anime) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

func (a *anime) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div[id^=sarea]").Attr("id")
	return utils.StrToNum(strings.Replace(id, "sarea", "", -1))
}

func (a *anime) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

func (a *anime) getSummary(nameArea *goquery.Selection) string {
	return strings.Replace(nameArea.Find(".pt4").Text(), "read more.", "", -1)
}

func (a *anime) getType(eachSearch *goquery.Selection) string {
	return strings.TrimSpace(strings.Replace(eachSearch.Find("td:nth-of-type(3)").Text(), "Unknown", "", -1))
}

func (a *anime) getProgress(eachSearch *goquery.Selection) int {
	return utils.StrToNum(eachSearch.Find("td:nth-of-type(4)").Text())
}

func (a *anime) getScore(eachSearch *goquery.Selection) float64 {
	return utils.StrToFloat(eachSearch.Find("td:nth-of-type(5)").Text())
}

func (a *anime) getStartDate(eachSearch *goquery.Selection) model.Date {
	y, m, d := utils.StrToDate(eachSearch.Find("td:nth-of-type(6)").Text())
	return model.Date{Year: y, Month: m, Day: d}
}

func (a *anime) getEndDate(eachSearch *goquery.Selection) model.Date {
	y, m, d := utils.StrToDate(eachSearch.Find("td:nth-of-type(7)").Text())
	return model.Date{Year: y, Month: m, Day: d}
}

func (a *anime) getMember(eachSearch *goquery.Selection) int {
	return utils.StrToNum(eachSearch.Find("td:nth-of-type(8)").Text())
}

func (a *anime) getRated(eachSearch *goquery.Selection) string {
	rated := strings.TrimSpace(eachSearch.Find("td:nth-of-type(9)").Text())
	if rated == "-" {
		return ""
	}
	return rated
}
