package anime

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type moreInfo struct {
	area *goquery.Selection
	data string
}

// GetMoreInfo to get anime more info.
func (p *parser) GetMoreInfo(a *goquery.Selection) string {
	c := moreInfo{area: a}
	c.setDetail()
	return c.data
}

func (mi *moreInfo) setDetail() {
	mi.area.Find("div").Remove()
	mi.area.Find("a").Remove()
	mi.area.Find("h2").Remove()
	mi.data = strings.TrimSpace(mi.area.Text())
}
