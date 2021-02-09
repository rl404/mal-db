package search

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type char struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.CharacterSearch
}

// GetCharacter to search character.
func (p *parser) GetCharacter(a *goquery.Selection) []model.CharacterSearch {
	v := char{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *char) setDetail() {
	characters := []model.CharacterSearch{}
	c.area.Find("table").Find("tr").Each(func(i int, eachSearch *goquery.Selection) {
		nameArea := eachSearch.Find("td:nth-of-type(2)")
		characters = append(characters, model.CharacterSearch{
			Image:    c.getImage(eachSearch),
			ID:       c.getID(nameArea),
			Name:     c.getName(nameArea),
			Nickname: c.getNickname(nameArea),
		})
	})
	c.data = characters
}

func (c *char) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td div.picSurround a img").Attr("data-src")
	return utils.URLCleaner(image, "image", c.cleanImg)
}

func (c *char) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (c *char) getName(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

func (c *char) getNickname(nameArea *goquery.Selection) string {
	nick := nameArea.Find("small").First().Text()
	if nick != "" {
		return nick[1 : len(nick)-1]
	}
	return nick
}
