package manga

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type character struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Role
}

// GetCharacters to get manga character list.
func (p *parser) GetCharacters(a *goquery.Selection) []model.Role {
	v := character{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *character) setDetail() {
	characters := []model.Role{}
	c.area.Find(".js-anime-character-table").Each(func(i int, area *goquery.Selection) {
		charNameArea := area.Find("td:nth-of-type(2)")
		characters = append(characters, model.Role{
			ID:    c.getID(charNameArea),
			Image: c.getImage(area),
			Name:  c.getName(charNameArea),
			Role:  c.getRole(charNameArea),
		})
	})
	c.data = characters
}

func (c *character) getID(charNameArea *goquery.Selection) int {
	id, _ := charNameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (c *character) getImage(charArea *goquery.Selection) string {
	image, _ := charArea.Find("td .picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", c.cleanImg)
}

func (c *character) getName(charNameArea *goquery.Selection) string {
	return charNameArea.Find("a").First().Text()
}

func (c *character) getRole(charNameArea *goquery.Selection) string {
	return strings.TrimSpace(charNameArea.Find(".spaceit_pad").First().Next().Text())
}
