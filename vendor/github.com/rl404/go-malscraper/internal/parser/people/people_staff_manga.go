package people

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type staffManga struct {
	area     *goquery.Selection
	_type    string
	cleanImg bool
	data     []model.Role
}

// GetStaffManga to get people anime staff or published manga list.
func (p *parser) GetStaffManga(a *goquery.Selection, t string) []model.Role {
	v := staffManga{area: a, cleanImg: p.cleanImg, _type: t}
	v.setDetail()
	return v.data
}

func (sm *staffManga) setDetail() {
	list := []model.Role{}

	area := sm.area.Find("#content table tr td").Next()
	area.Find(".normal_header").First().Remove()

	if sm._type == "manga" {
		area.Find(".normal_header").First().Remove()
	}

	area = area.Find(".normal_header").First().Next()

	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, staff *goquery.Selection) {
			animeImageArea := staff.Find("td:nth-of-type(1)")
			stArea := staff.Find("td:nth-of-type(2)")

			list = append(list, model.Role{
				ID:    sm.getID(stArea),
				Name:  sm.getTitle(stArea),
				Image: sm.getImage(animeImageArea),
				Role:  sm.getRole(stArea),
			})

		})
	}
	sm.data = list
}

func (sm *staffManga) getID(animeArea *goquery.Selection) int {
	animeID, _ := animeArea.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(animeID, "/", 4))
}

func (sm *staffManga) getImage(animeImageArea *goquery.Selection) string {
	animeImage, _ := animeImageArea.Find("img").Attr("data-src")
	return utils.URLCleaner(animeImage, "image", sm.cleanImg)
}

func (sm *staffManga) getTitle(animeArea *goquery.Selection) string {
	return animeArea.Find("a").First().Text()
}

func (sm *staffManga) getRole(stArea *goquery.Selection) string {
	return strings.TrimSpace(stArea.Find("small").Text())
}
