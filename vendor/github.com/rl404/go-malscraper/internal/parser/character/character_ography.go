package character

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type ography struct {
	area     *goquery.Selection
	_type    string
	cleanImg bool
	data     []model.Role
}

// GetOgraphy to get character animeography/mangaography list.
func (p *parser) GetOgraphy(a *goquery.Selection, t string) []model.Role {
	d := ography{area: a, cleanImg: p.cleanImg, _type: t}

	if !d.isValid() {
		return nil
	}

	d.setDetail()
	return d.data
}

func (o *ography) isValid() bool {
	return o.area.Find(".badresult").Text() == ""
}

func (o *ography) setDetail() {
	medias := []model.Role{}
	area := o.area.Find("#content table tr td:nth-of-type(1)")

	if o._type == "anime" {
		area = area.Find("table:nth-of-type(1)")
	} else {
		area = area.Find("table:nth-of-type(2)")
	}

	area.Find("tr").Each(func(i int, media *goquery.Selection) {
		mediaImage := media.Find("td:nth-of-type(1)")
		eachArea := media.Find("td:nth-of-type(2)")
		medias = append(medias, model.Role{
			ID:    o.getID(eachArea),
			Name:  o.getName(eachArea),
			Image: o.getImage(mediaImage),
			Role:  o.getRole(eachArea),
		})
	})
	o.data = medias
}

func (o *ography) getID(vaArea *goquery.Selection) int {
	id, _ := vaArea.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (o *ography) getName(vaArea *goquery.Selection) string {
	name, _ := vaArea.Find("a:nth-of-type(1)").Html()
	return name
}

func (o *ography) getImage(vaArea *goquery.Selection) string {
	image, _ := vaArea.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", o.cleanImg)
}

func (o *ography) getRole(vaArea *goquery.Selection) string {
	return vaArea.Find("div small").Text()
}
