package producermagazine

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type magazines struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ItemCount
}

// GetMagazines to get manga magazine/serialization list.
func (p *parser) GetMagazines(a *goquery.Selection) []model.ItemCount {
	v := magazines{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (m *magazines) setDetail() {
	magazines := []model.ItemCount{}
	m.area.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		magazines = append(magazines, model.ItemCount{
			ID:    m.getID(area),
			Name:  m.getName(area),
			Count: m.getCount(area),
		})
	})
	m.data = magazines
}

func (m *magazines) getID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 3))
}

func (m *magazines) getName(area *goquery.Selection) string {
	r := regexp.MustCompile(`\([0-9,-]+\)`)
	return strings.TrimSpace(r.ReplaceAllString(area.Text(), ""))
}

func (m *magazines) getCount(area *goquery.Selection) int {
	r := regexp.MustCompile(`\([0-9,-]+\)`)
	count := r.FindString(area.Text())
	return utils.StrToNum(count[1 : len(count)-1])
}
