package top

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type character struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.TopCharacter
}

// GetCharacter to get top character list.
func (p *parser) GetCharacter(a *goquery.Selection) []model.TopCharacter {
	v := character{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *character) setDetail() {
	characters := []model.TopCharacter{}
	c.area.Find(".characters-favorites-ranking-table").Find("tr.ranking-list").Each(func(i int, eachChar *goquery.Selection) {
		nameArea := eachChar.Find(".people")
		characters = append(characters, model.TopCharacter{
			Rank:         c.getRank(eachChar),
			ID:           c.getID(nameArea),
			Name:         c.getName(nameArea),
			JapaneseName: c.getJapaneseName(nameArea),
			Image:        c.getImage(nameArea),
			Favorite:     c.getFavorite(eachChar),
		})
	})
	c.data = characters
}

func (c *character) getRank(eachChar *goquery.Selection) int {
	return utils.StrToNum(eachChar.Find("td").First().Find("span").Text())
}

func (c *character) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (c *character) getName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

func (c *character) getJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()
	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}
	return japName
}

func (c *character) getImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", c.cleanImg)
}

func (c *character) getFavorite(eachChar *goquery.Selection) int {
	return utils.StrToNum(eachChar.Find(".favorites").Text())
}
