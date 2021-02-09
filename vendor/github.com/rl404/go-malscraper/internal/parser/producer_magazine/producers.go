package producermagazine

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type producers struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ItemCount
}

// GetProducers to get anime producer/studio/licensor list.
func (p *parser) GetProducers(a *goquery.Selection) []model.ItemCount {
	v := producers{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *producers) setDetail() {
	producers := []model.ItemCount{}
	p.area.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		producers = append(producers, model.ItemCount{
			ID:    p.getID(area),
			Name:  p.getName(area),
			Count: p.getCount(area),
		})
	})
	p.data = producers
}

func (p *producers) getID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 3))
}

func (p *producers) getName(area *goquery.Selection) string {
	r := regexp.MustCompile(`\([0-9,-]+\)`)
	return strings.TrimSpace(r.ReplaceAllString(area.Text(), ""))
}

func (p *producers) getCount(area *goquery.Selection) int {
	r := regexp.MustCompile(`\([0-9,-]+\)`)
	count := r.FindString(area.Text())
	return utils.StrToNum(count[1 : len(count)-1])
}
