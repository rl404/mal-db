package club

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type picture struct {
	area     *goquery.Selection
	cleanImg bool
	data     []string
}

// GetPictures to get club picture list.
func (p *parser) GetPictures(a *goquery.Selection) []string {
	v := picture{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *picture) setDetail() {
	pics := []string{}
	p.area.Find("table").First().Find("td.borderClass").Each(func(i int, eachPic *goquery.Selection) {
		image, _ := eachPic.Find(".picSurround img").Attr("data-src")
		pics = append(pics, utils.URLCleaner(image, "image", p.cleanImg))
	})
	p.data = pics
}
