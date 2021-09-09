package anime

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type character struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.CharacterItem
}

// GetCharacters to get anime character list.
func (p *parser) GetCharacters(a *goquery.Selection) []model.CharacterItem {
	c := character{area: a, cleanImg: p.cleanImg}
	c.setDetail()
	return c.data
}

func (c *character) setDetail() {
	data := []model.CharacterItem{}
	c.area.Find(".js-anime-character-table").Each(func(i int, area *goquery.Selection) {
		charNameArea := area.Find("td:nth-of-type(2)")
		vaArea := area.Find("td:nth-of-type(3)")
		data = append(data, model.CharacterItem{
			ID:          c.getID(charNameArea),
			Image:       c.getImage(area),
			Name:        c.getName(charNameArea),
			Role:        c.getRole(charNameArea),
			VoiceActors: c.getVa(vaArea),
		})
	})
	c.data = data
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

func (c *character) getVa(vaArea *goquery.Selection) []model.Role {
	vaList := []model.Role{}
	vaArea.Find("table").Find("tr").Each(func(i int, eachVa *goquery.Selection) {
		vaNameArea := eachVa.Find("td").First()
		vaList = append(vaList, model.Role{
			ID:    c.getID(vaNameArea),
			Image: c.getImage(eachVa),
			Name:  c.getName(vaNameArea),
			Role:  c.getRole(vaNameArea),
		})
	})
	return vaList
}
