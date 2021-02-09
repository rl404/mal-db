package people

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type character struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.PeopleCharacter
}

// GetCharacters to get people anime character list.
func (p *parser) GetCharacters(a *goquery.Selection) []model.PeopleCharacter {
	v := character{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *character) setDetail() {
	actors := []model.PeopleCharacter{}
	area := c.area.Find("#content table tr td").Next()
	area = area.Find(".normal_header").First().Next()
	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, va *goquery.Selection) {
			animeImageArea := va.Find("td:nth-of-type(1)")
			animeArea := va.Find("td:nth-of-type(2)")
			charImageArea := va.Find("td:nth-of-type(4)")
			charArea := va.Find("td:nth-of-type(3)")
			actors = append(actors, model.PeopleCharacter{
				Anime: model.Role{
					ID:    c.getID(animeArea),
					Name:  c.getTitle(animeArea),
					Image: c.getImage(animeImageArea),
					Role:  c.getRole(charArea),
				},
				Character: model.Role{
					ID:    c.getID(charArea),
					Name:  c.getTitle(charArea),
					Role:  c.getRole(charArea),
					Image: c.getImage(charImageArea),
				},
			})
		})
	}
	c.data = actors
}

func (c *character) getID(animeArea *goquery.Selection) int {
	animeID, _ := animeArea.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(animeID, "/", 4))
}

func (c *character) getImage(animeImageArea *goquery.Selection) string {
	animeImage, _ := animeImageArea.Find("img").Attr("data-src")
	return utils.URLCleaner(animeImage, "image", c.cleanImg)
}

func (c *character) getTitle(animeArea *goquery.Selection) string {
	return animeArea.Find("a").First().Text()
}

func (c *character) getRole(animeArea *goquery.Selection) string {
	return strings.TrimSpace(animeArea.Find("div").Text())
}
