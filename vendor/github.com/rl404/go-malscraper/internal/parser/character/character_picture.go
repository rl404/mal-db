package character

import "github.com/PuerkitoBio/goquery"

type picture struct {
	area     *goquery.Selection
	cleanImg bool
	data     []string
}

// GetPictures to get character picture list.
func (p *parser) GetPictures(a *goquery.Selection) []string {
	d := picture{area: a, cleanImg: p.cleanImg}
	d.setDetail()
	return d.data
}

func (p *picture) setDetail() {
	pictures := []string{}
	area := p.area.Find("table tr td").Next().Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})
	p.data = pictures
}
