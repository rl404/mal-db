package character

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type va struct {
	area *goquery.Selection
	cleanImg bool
	data []model.Role
}

// GetVA to get character voice actor list.
func (p *parser) GetVA(a *goquery.Selection) []model.Role {
	v := va{area: a, cleanImg: p.cleanImg}

	if !v.isValid() {
		return nil
	}

	v.setDetail()
	return v.data
}

func (v *va) isValid() bool {
	return v.area.Find(".badresult").Text() == ""
}

func (v *va) setDetail() {
	vas := []model.Role{}
	vaArea := v.area.Find("#content table tr td").Next()
	vaArea.Find("table[width=\"100%\"]").Each(func(i int, va *goquery.Selection) {
		vaNameArea := va.Find("td:nth-of-type(2)")
		vas = append(vas, model.Role{
			ID:    v.getID(vaNameArea),
			Name:  v.getName(vaNameArea),
			Role:  v.getRole(vaNameArea),
			Image: v.getImage(va),
		})
	})
	v.data = vas
}

func (v *va) getID(vaArea *goquery.Selection) int {
	id, _ := vaArea.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (v *va) getName(vaArea *goquery.Selection) string {
	name, _ := vaArea.Find("a:nth-of-type(1)").Html()
	return name
}

func (v *va) getImage(vaArea *goquery.Selection) string {
	image, _ := vaArea.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", v.cleanImg)
}

func (v *va) getRole(vaArea *goquery.Selection) string {
	return vaArea.Find("div small").Text()
}
