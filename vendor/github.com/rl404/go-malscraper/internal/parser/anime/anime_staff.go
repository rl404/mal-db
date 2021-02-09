package anime

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type staff struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Role
}

// GetStaff to get anime staff list.
func (p *parser) GetStaff(a *goquery.Selection) []model.Role {
	c := staff{area: a, cleanImg: p.cleanImg}
	c.setDetail()
	return c.data
}

func (s *staff) setDetail() {
	staffList := []model.Role{}
	s.area.Find("article").Remove()
	s.area.Find("h2").First().Remove()
	staffArea := s.area.Find("h2").First().Parent().Next()
	for goquery.NodeName(staffArea) == "table" {
		staffNameArea := staffArea.Find("td:nth-of-type(2)")
		staffList = append(staffList, model.Role{
			ID:    s.getID(staffNameArea),
			Image: s.getImage(staffArea),
			Name:  s.getName(staffNameArea),
			Role:  s.getRole(staffNameArea),
		})
		staffArea = staffArea.Next()
	}
	s.data = staffList
}

func (s *staff) getID(charNameArea *goquery.Selection) int {
	id, _ := charNameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (s *staff) getImage(charArea *goquery.Selection) string {
	image, _ := charArea.Find("td .picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", s.cleanImg)
}

func (s *staff) getName(charNameArea *goquery.Selection) string {
	return charNameArea.Find("a").First().Text()
}

func (s *staff) getRole(charNameArea *goquery.Selection) string {
	return charNameArea.Find("small").First().Text()
}
