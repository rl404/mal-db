package user

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type history struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.UserHistory
}

// GetHistory to get user history list.
func (p *parser) GetHistory(a *goquery.Selection) []model.UserHistory {
	v := history{area: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (h *history) setDetails() {
	historyList := []model.UserHistory{}
	h.area.Find("table").Find("tr").EachWithBreak(func(i int, history *goquery.Selection) bool {
		historyClass, _ := history.Find("td").First().Attr("class")
		if historyClass != "borderClass" {
			return true
		}
		nameArea := history.Find("td").First()
		historyList = append(historyList, model.UserHistory{
			ID:       h.getID(nameArea),
			Title:    h.getTitle(nameArea),
			Type:     h.getType(nameArea),
			Progress: h.getProgress(nameArea),
			Date:     h.getDate(history),
		})
		return true
	})
	h.data = historyList
}

func (h *history) getID(nameArea *goquery.Selection) int {
	idLink, _ := nameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(idLink, "=", 1))
}

func (h *history) getTitle(nameArea *goquery.Selection) string {
	return strings.TrimSpace(nameArea.Find("a").First().Text())
}

func (h *history) getType(nameArea *goquery.Selection) string {
	typeLink, _ := nameArea.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(typeLink, ".php", 0)[1:]
}

func (h *history) getProgress(nameArea *goquery.Selection) int {
	return utils.StrToNum(nameArea.Find("strong").First().Text())
}

func (h *history) getDate(history *goquery.Selection) time.Time {
	date := history.Find("td:nth-of-type(2)")
	date.Find("a").Remove()
	t, _ := utils.StrToTime(date.Text())
	return t
}
