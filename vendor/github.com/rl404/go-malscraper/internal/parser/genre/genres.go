package genre

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type genres struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ItemCount
}

// GetGenres to get anime/manga genre list.
func (p *parser) GetGenres(a *goquery.Selection) []model.ItemCount {
	v := genres{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (g *genres) setDetail() {
	genres := []model.ItemCount{}
	g.area.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		genres = append(genres, model.ItemCount{
			ID:    g.getID(area),
			Name:  g.getName(area),
			Count: g.getCount(area),
		})
	})
	g.data = genres
}

func (g *genres) getID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 3))
}

func (g *genres) getName(area *goquery.Selection) string {
	r := regexp.MustCompile(`\([0-9,]+\)`)
	return strings.TrimSpace(r.ReplaceAllString(area.Text(), ""))
}

func (g *genres) getCount(area *goquery.Selection) int {
	r, _ := regexp.Compile(`\([0-9,]+\)`)
	count := r.FindString(area.Text())
	return utils.StrToNum(count[1 : len(count)-1])
}
