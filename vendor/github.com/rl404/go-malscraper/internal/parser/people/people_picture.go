package people

import "github.com/PuerkitoBio/goquery"

type picture struct {
	area     *goquery.Selection
	cleanImg bool
	data     []string
}

// GetPictures to get people picture list.
func (p *parser) GetPictures(a *goquery.Selection) []string {
	v := picture{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *picture) setDetail() {
	pictures := []string{}
	p.area.Next().Find("table").First().Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})
	p.data = pictures
}
